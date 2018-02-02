package cluster

import (
	"bytes"
	"encoding/csv"
	"fmt"

	"github.com/golang/glog"
	apiv1 "github.com/kubermatic/kubermatic/api/pkg/api/v1"
	"github.com/kubermatic/kubermatic/api/pkg/controller/resources"
	kubermaticv1 "github.com/kubermatic/kubermatic/api/pkg/crd/kubermatic/v1"
	kuberneteshelper "github.com/kubermatic/kubermatic/api/pkg/kubernetes"
	"github.com/kubermatic/kubermatic/api/pkg/provider"

	corev1 "k8s.io/api/core/v1"
	rbacv1beta1 "k8s.io/api/rbac/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

const (
	nodeDeletionFinalizer         = "kubermatic.io/delete-nodes"
	cloudProviderCleanupFinalizer = "kubermatic.io/cleanup-cloud-provider"
	namespaceDeletionFinalizer    = "kubermatic.io/delete-ns"

	minNodePort = 30000
	maxNodePort = 32767
)

func (cc *controller) syncPendingCluster(c *kubermaticv1.Cluster) (*kubermaticv1.Cluster, error) {
	if c.Spec.MasterVersion == "" {
		c.Spec.MasterVersion = cc.defaultMasterVersion.ID
	}

	//Every function with the prefix 'pending' *WILL* modify the cluster struct and cause an update
	//Every function with the prefix 'launching' *WONT* modify the cluster struct and should not cause an update

	// Setup required infrastructure at cloud provider
	if changedC, err := cc.pendingInitializeCloudProvider(c); err != nil || changedC != nil {
		return changedC, err
	}

	// Add finalizers
	if changedC, err := cc.pendingRegisterDefaultFinalizers(c); err != nil || changedC != nil {
		return changedC, err
	}

	// Set the hostname & url
	if changedC, err := cc.pendingCreateAddresses(c); err != nil || changedC != nil {
		return changedC, err
	}

	// Generate the kubelet and admin token
	if changedC, err := cc.pendingCreateTokens(c); err != nil || changedC != nil {
		return changedC, err
	}

	// Create the root ca
	if changedC, err := cc.pendingCreateRootCA(c); err != nil || changedC != nil {
		return changedC, err
	}

	// Create the certificates
	if changedC, err := cc.pendingCreateCertificates(c); err != nil || changedC != nil {
		return changedC, err
	}

	// Create the service account key
	if changedC, err := cc.pendingCreateServiceAccountKey(c); err != nil || changedC != nil {
		return changedC, err
	}

	// Create the ssh keys for the apiserver
	if changedC, err := cc.pendingCreateApiserverSSHKeys(c); err != nil || changedC != nil {
		return changedC, err
	}

	// Create the namespace
	if err := cc.launchingCreateNamespace(c); err != nil {
		return nil, err
	}

	// Create secret for user tokens
	if err := cc.launchingCheckTokenUsers(c); err != nil {
		return nil, err
	}

	// check that all service accounts are created
	if err := cc.launchingCheckServiceAccounts(c); err != nil {
		return nil, err
	}

	// check that all role bindings are created
	if err := cc.launchingCheckClusterRoleBindings(c); err != nil {
		return nil, err
	}

	// check that all services are available
	if err := cc.launchingCheckServices(c); err != nil {
		return nil, err
	}

	// check that all secrets are available
	if err := cc.launchingCheckSecrets(c); err != nil {
		return nil, err
	}

	// check that all configmaps are available
	if err := cc.launchingCheckConfigMaps(c); err != nil {
		return nil, err
	}

	// check that all deployments are available
	if err := cc.launchingCheckDeployments(c); err != nil {
		return nil, err
	}

	// check that the etcd-cluster cr is available
	if err := cc.launchingCheckEtcdCluster(c); err != nil {
		return nil, err
	}

	c.Status.LastTransitionTime = metav1.Now()
	c.Status.Phase = kubermaticv1.LaunchingClusterStatusPhase
	return c, nil
}

func (cc *controller) pendingInitializeCloudProvider(cluster *kubermaticv1.Cluster) (*kubermaticv1.Cluster, error) {
	_, prov, err := provider.ClusterCloudProvider(cc.cps, cluster)
	if err != nil {
		return nil, err
	}

	cloud, err := prov.InitializeCloudProvider(cluster.Spec.Cloud, cluster.Name)
	if err != nil {
		return nil, err
	}
	if cloud != nil {
		cluster.Spec.Cloud = cloud
		return cluster, nil
	}

	if !kuberneteshelper.HasFinalizer(cluster, cloudProviderCleanupFinalizer) {
		cluster.Finalizers = append(cluster.Finalizers, cloudProviderCleanupFinalizer)
		return cluster, nil
	}

	return nil, nil
}

// launchingCreateNamespace will create the cluster namespace
func (cc *controller) launchingCreateNamespace(c *kubermaticv1.Cluster) error {
	informerGroup, err := cc.clientProvider.GetInformerGroup(c.Spec.SeedDatacenterName)
	if err != nil {
		return fmt.Errorf("failed to get informer group for dc %q: %v", c.Spec.SeedDatacenterName, err)
	}

	_, err = informerGroup.NamespaceInformer.Lister().Get(c.Status.NamespaceName)
	if !errors.IsNotFound(err) {
		return err
	}

	ns := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: c.Status.NamespaceName,
		},
	}
	client, err := cc.clientProvider.GetClient(c.Spec.SeedDatacenterName)
	if err != nil {
		return fmt.Errorf("failed to get client for dc %q: %v", c.Spec.SeedDatacenterName, err)
	}
	_, err = client.CoreV1().Namespaces().Create(ns)
	return err
}

// pendingRegisterDefaultFinalizers adds all default finalizers we need
func (cc *controller) pendingRegisterDefaultFinalizers(c *kubermaticv1.Cluster) (*kubermaticv1.Cluster, error) {
	var updated bool

	finalizers := []string{
		namespaceDeletionFinalizer,
	}

	for _, f := range finalizers {
		if !kuberneteshelper.HasFinalizer(c, f) {
			c.Finalizers = append(c.Finalizers, f)
			updated = true
		}
	}

	if updated {
		glog.V(4).Infof("Added finalizers to cluster %s", c.Name)
		return c, nil
	}
	return nil, nil
}

func (cc *controller) getFreeNodePort(dc string) (int, error) {
	informerGroup, err := cc.clientProvider.GetInformerGroup(dc)
	if err != nil {
		return 0, fmt.Errorf("failed to get informer group for dc %q: %v", dc, err)
	}

	services, err := informerGroup.ServiceInformer.Lister().List(labels.Everything())
	if err != nil {
		return 0, err
	}
	allocatedPorts := map[int]struct{}{}

	for _, s := range services {
		for _, p := range s.Spec.Ports {
			if p.NodePort != 0 {
				allocatedPorts[int(p.NodePort)] = struct{}{}
			}
		}
	}

	for i := minNodePort; i < maxNodePort; i++ {
		if _, exists := allocatedPorts[i]; !exists {
			return i, nil
		}
	}

	return 0, fmt.Errorf("no free nodeport left")
}

// pendingCreateAddresses will set the cluster hostname and the url under which the apiserver will be reachable
func (cc *controller) pendingCreateAddresses(c *kubermaticv1.Cluster) (*kubermaticv1.Cluster, error) {
	var updated bool

	if c.Address.ExternalName == "" {
		c.Address.ExternalName = fmt.Sprintf("%s.%s.%s", c.Name, c.Spec.SeedDatacenterName, cc.externalURL)
		updated = true
	}

	if c.Address.ExternalPort == 0 {
		port, err := cc.getFreeNodePort(c.Spec.SeedDatacenterName)
		if err != nil {
			return nil, fmt.Errorf("failed to get nodeport: %v", err)
		}
		c.Address.ExternalPort = port
		updated = true
	}

	if c.Address.URL == "" {
		c.Address.URL = fmt.Sprintf("https://%s:%d", c.Address.ExternalName, c.Address.ExternalPort)
		updated = true
	}

	if updated {
		glog.V(4).Infof("Set address for cluster %s to %s", c.Name, c.Address.URL)
		return c, nil
	}
	return nil, nil
}

func (cc *controller) launchingCheckSecrets(c *kubermaticv1.Cluster) error {
	secrets := map[string]func(c *kubermaticv1.Cluster, app, masterResourcesPath string) (*corev1.Secret, error){
		"apiserver":          resources.LoadSecretFile,
		"controller-manager": resources.LoadSecretFile,
	}

	informerGroup, err := cc.clientProvider.GetInformerGroup(c.Spec.SeedDatacenterName)
	if err != nil {
		return fmt.Errorf("failed to get informer group for dc %q: %v", c.Spec.SeedDatacenterName, err)
	}

	ns := c.Status.NamespaceName
	for s, gen := range secrets {
		_, err = informerGroup.SecretInformer.Lister().Secrets(ns).Get(s)
		if err != nil && !errors.IsNotFound(err) {
			return err
		}
		if err == nil {
			continue
		}

		secret, err := gen(c, s, cc.masterResourcesPath)
		if err != nil {
			return fmt.Errorf("failed to generate %s: %v", s, err)
		}

		client, err := cc.clientProvider.GetClient(c.Spec.SeedDatacenterName)
		if err != nil {
			return fmt.Errorf("failed to get client for dc %q: %v", c.Spec.SeedDatacenterName, err)
		}

		_, err = client.CoreV1().Secrets(ns).Create(secret)
		if err != nil {
			return fmt.Errorf("failed to create secret for %s: %v", s, err)
		}
	}

	return nil
}

func (cc *controller) launchingCheckServices(c *kubermaticv1.Cluster) error {
	services := map[string]func(c *kubermaticv1.Cluster, app, masterResourcesPath string) (*corev1.Service, error){
		"apiserver":          resources.LoadServiceFile,
		"apiserver-external": resources.LoadServiceFile,
	}

	informerGroup, err := cc.clientProvider.GetInformerGroup(c.Spec.SeedDatacenterName)
	if err != nil {
		return fmt.Errorf("failed to get informer group for dc %q: %v", c.Spec.SeedDatacenterName, err)
	}

	ns := c.Status.NamespaceName
	for s, gen := range services {
		_, err = informerGroup.ServiceInformer.Lister().Services(ns).Get(s)
		if err != nil && !errors.IsNotFound(err) {
			return err
		}
		if err == nil {
			continue
		}

		service, err := gen(c, s, cc.masterResourcesPath)
		if err != nil {
			return fmt.Errorf("failed to generate service %s: %v", s, err)
		}

		client, err := cc.clientProvider.GetClient(c.Spec.SeedDatacenterName)
		if err != nil {
			return fmt.Errorf("failed to get client for dc %q: %v", c.Spec.SeedDatacenterName, err)
		}

		_, err = client.CoreV1().Services(ns).Create(service)
		if err != nil {
			return fmt.Errorf("failed to create service %s: %v", s, err)
		}
	}

	return nil
}

func (cc *controller) launchingCheckServiceAccounts(c *kubermaticv1.Cluster) error {
	serviceAccounts := map[string]func(app, masterResourcesPath string) (*corev1.ServiceAccount, error){
		"etcd-operator": resources.LoadServiceAccountFile,
	}

	informerGroup, err := cc.clientProvider.GetInformerGroup(c.Spec.SeedDatacenterName)
	if err != nil {
		return fmt.Errorf("failed to get informer group for dc %q: %v", c.Spec.SeedDatacenterName, err)
	}

	ns := c.Status.NamespaceName
	for s, gen := range serviceAccounts {
		_, err := informerGroup.ServiceAccountInformer.Lister().ServiceAccounts(ns).Get(s)
		if err != nil && !errors.IsNotFound(err) {
			return err
		}
		if err == nil {
			continue
		}

		sa, err := gen(s, cc.masterResourcesPath)
		if err != nil {
			return fmt.Errorf("failed to generate service account %s: %v", s, err)
		}

		client, err := cc.clientProvider.GetClient(c.Spec.SeedDatacenterName)
		if err != nil {
			return fmt.Errorf("failed to get client for dc %q: %v", c.Spec.SeedDatacenterName, err)
		}

		_, err = client.CoreV1().ServiceAccounts(ns).Create(sa)
		if err != nil {
			return fmt.Errorf("failed to create service account %s: %v", s, err)
		}
	}

	return nil
}

func (cc *controller) launchingCheckTokenUsers(c *kubermaticv1.Cluster) error {
	ns := c.Status.NamespaceName
	name := "token-users"

	informerGroup, err := cc.clientProvider.GetInformerGroup(c.Spec.SeedDatacenterName)
	if err != nil {
		return fmt.Errorf("failed to get informer group for dc %q: %v", c.Spec.SeedDatacenterName, err)
	}

	_, err = informerGroup.SecretInformer.Lister().Secrets(ns).Get(name)
	if !errors.IsNotFound(err) {
		return err
	}

	buffer := bytes.Buffer{}
	writer := csv.NewWriter(&buffer)
	if err := writer.Write([]string{c.Address.KubeletToken, "kubelet-bootstrap", "10001", "system:bootstrappers"}); err != nil {
		return err
	}
	if err := writer.Write([]string{c.Address.AdminToken, "admin", "10000", "system:masters"}); err != nil {
		return err
	}
	writer.Flush()

	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Type: corev1.SecretTypeOpaque,
		Data: map[string][]byte{
			"tokens.csv": buffer.Bytes(),
		},
	}

	client, err := cc.clientProvider.GetClient(c.Spec.SeedDatacenterName)
	if err != nil {
		return fmt.Errorf("failed to get client for dc %q: %v", c.Spec.SeedDatacenterName, err)
	}

	_, err = client.CoreV1().Secrets(ns).Create(secret)
	if err != nil {
		return fmt.Errorf("failed to create user token secret: %v", err)
	}
	return nil
}

func (cc *controller) launchingCheckClusterRoleBindings(c *kubermaticv1.Cluster) error {
	roleBindings := map[string]func(namespace, app, masterResourcesPath string) (*rbacv1beta1.ClusterRoleBinding, error){
		"etcd-operator": resources.LoadClusterRoleBindingFile,
	}

	informerGroup, err := cc.clientProvider.GetInformerGroup(c.Spec.SeedDatacenterName)
	if err != nil {
		return fmt.Errorf("failed to get informer group for dc %q: %v", c.Spec.SeedDatacenterName, err)
	}

	ns := c.Status.NamespaceName
	for s, gen := range roleBindings {
		binding, err := gen(ns, s, cc.masterResourcesPath)
		if err != nil {
			return fmt.Errorf("failed to generate cluster role binding %s: %v", s, err)
		}

		_, err = informerGroup.ClusterRoleBindingInformer.Lister().Get(binding.ObjectMeta.Name)
		if err != nil && !errors.IsNotFound(err) {
			return err
		}
		if err == nil {
			continue
		}

		client, err := cc.clientProvider.GetClient(c.Spec.SeedDatacenterName)
		if err != nil {
			return fmt.Errorf("failed to get client for dc %q: %v", c.Spec.SeedDatacenterName, err)
		}

		_, err = client.RbacV1beta1().ClusterRoleBindings().Create(binding)
		if err != nil {
			return fmt.Errorf("failed to create cluster role binding %s: %v", s, err)
		}
	}

	return nil
}

func (cc *controller) launchingCheckDeployments(c *kubermaticv1.Cluster) error {
	ns := c.Status.NamespaceName
	masterVersion, found := cc.versions[c.Spec.MasterVersion]
	if !found {
		return fmt.Errorf("unknown new cluster %q master version %q", c.Name, c.Spec.MasterVersion)
	}

	deps := map[string]string{
		"etcd-operator":      masterVersion.EtcdOperatorDeploymentYaml,
		"apiserver":          masterVersion.ApiserverDeploymentYaml,
		"controller-manager": masterVersion.ControllerDeploymentYaml,
		"scheduler":          masterVersion.SchedulerDeploymentYaml,
		"node-controller":    masterVersion.NodeControllerDeploymentYaml,
		"addon-manager":      masterVersion.AddonManagerDeploymentYaml,
		"machine-controller": masterVersion.MachineControllerDeploymentYaml,
	}

	informerGroup, err := cc.clientProvider.GetInformerGroup(c.Spec.SeedDatacenterName)
	if err != nil {
		return fmt.Errorf("failed to get informer group for dc %q: %v", c.Spec.SeedDatacenterName, err)
	}

	for name, yamlFile := range deps {
		dep, err := resources.LoadDeploymentFile(c, masterVersion, cc.masterResourcesPath, yamlFile)
		if err != nil {
			return fmt.Errorf("failed to generate deployment %q: %v", name, err)
		}

		_, err = informerGroup.DeploymentInformer.Lister().Deployments(ns).Get(name)
		if err != nil && !errors.IsNotFound(err) {
			return err
		}
		if err == nil {
			continue
		}

		client, err := cc.clientProvider.GetClient(c.Spec.SeedDatacenterName)
		if err != nil {
			return fmt.Errorf("failed to get client for dc %q: %v", c.Spec.SeedDatacenterName, err)
		}

		_, err = client.ExtensionsV1beta1().Deployments(ns).Create(dep)
		if err != nil {
			return fmt.Errorf("failed to create deployment %q: %v", name, err)
		}
	}

	return nil
}

func (cc *controller) launchingCheckConfigMaps(c *kubermaticv1.Cluster) error {
	version, found := cc.versions[c.Spec.MasterVersion]
	if !found {
		return fmt.Errorf("failed to get version %s", c.Spec.MasterVersion)
	}
	ns := c.Status.NamespaceName
	cms := map[string]func(c *kubermaticv1.Cluster, datacenter *provider.DatacenterMeta, version *apiv1.MasterVersion) (*corev1.ConfigMap, error){}
	if c.Spec.Cloud != nil {
		if c.Spec.Cloud.AWS != nil {
			cms["cloud-config"] = resources.LoadAwsCloudConfigConfigMap
		}
		if c.Spec.Cloud.Openstack != nil {
			cms["cloud-config"] = resources.LoadOpenstackCloudConfigConfigMap
		}
	}

	informerGroup, err := cc.clientProvider.GetInformerGroup(c.Spec.SeedDatacenterName)
	if err != nil {
		return fmt.Errorf("failed to get informer group for dc %q: %v", c.Spec.SeedDatacenterName, err)
	}

	for s, gen := range cms {
		_, err := informerGroup.ConfigMapInformer.Lister().ConfigMaps(ns).Get(s)
		if err != nil && !errors.IsNotFound(err) {
			return err
		}
		if err == nil {
			continue
		}

		dc := cc.dcs[c.Spec.Cloud.DatacenterName]
		cm, err := gen(c, &dc, version)
		if err != nil {
			return fmt.Errorf("failed to generate cm %s: %v", s, err)
		}

		client, err := cc.clientProvider.GetClient(c.Spec.SeedDatacenterName)
		if err != nil {
			return fmt.Errorf("failed to get client for dc %q: %v", c.Spec.SeedDatacenterName, err)
		}

		_, err = client.CoreV1().ConfigMaps(ns).Create(cm)
		if err != nil {
			return fmt.Errorf("failed to create cm %s; %v", s, err)
		}
	}

	return nil
}

func (cc *controller) launchingCheckEtcdCluster(c *kubermaticv1.Cluster) error {
	ns := c.Status.NamespaceName
	masterVersion, found := cc.versions[c.Spec.MasterVersion]
	if !found {
		return fmt.Errorf("unknown new cluster %q master version %q", c.Name, c.Spec.MasterVersion)
	}

	etcd, err := resources.LoadEtcdClusterFile(masterVersion, cc.masterResourcesPath, masterVersion.EtcdClusterYaml)
	if err != nil {
		return fmt.Errorf("failed to load etcd-cluster: %v", err)
	}

	informerGroup, err := cc.clientProvider.GetInformerGroup(c.Spec.SeedDatacenterName)
	if err != nil {
		return fmt.Errorf("failed to get informer group for dc %q: %v", c.Spec.SeedDatacenterName, err)
	}

	_, err = informerGroup.EtcdClusterInformer.Lister().EtcdClusters(ns).Get(etcd.ObjectMeta.Name)
	if !errors.IsNotFound(err) {
		return err
	}

	client, err := cc.clientProvider.GetCRDClient(c.Spec.SeedDatacenterName)
	if err != nil {
		return fmt.Errorf("failed to get client for dc %q: %v", c.Spec.SeedDatacenterName, err)
	}

	_, err = client.EtcdV1beta2().EtcdClusters(ns).Create(etcd)
	if err != nil {
		return fmt.Errorf("failed to create etcd-cluster definition (crd): %v", err)
	}

	return nil
}

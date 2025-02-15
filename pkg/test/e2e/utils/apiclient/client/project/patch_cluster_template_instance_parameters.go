// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewPatchClusterTemplateInstanceParams creates a new PatchClusterTemplateInstanceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchClusterTemplateInstanceParams() *PatchClusterTemplateInstanceParams {
	return &PatchClusterTemplateInstanceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchClusterTemplateInstanceParamsWithTimeout creates a new PatchClusterTemplateInstanceParams object
// with the ability to set a timeout on a request.
func NewPatchClusterTemplateInstanceParamsWithTimeout(timeout time.Duration) *PatchClusterTemplateInstanceParams {
	return &PatchClusterTemplateInstanceParams{
		timeout: timeout,
	}
}

// NewPatchClusterTemplateInstanceParamsWithContext creates a new PatchClusterTemplateInstanceParams object
// with the ability to set a context for a request.
func NewPatchClusterTemplateInstanceParamsWithContext(ctx context.Context) *PatchClusterTemplateInstanceParams {
	return &PatchClusterTemplateInstanceParams{
		Context: ctx,
	}
}

// NewPatchClusterTemplateInstanceParamsWithHTTPClient creates a new PatchClusterTemplateInstanceParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchClusterTemplateInstanceParamsWithHTTPClient(client *http.Client) *PatchClusterTemplateInstanceParams {
	return &PatchClusterTemplateInstanceParams{
		HTTPClient: client,
	}
}

/* PatchClusterTemplateInstanceParams contains all the parameters to send to the API endpoint
   for the patch cluster template instance operation.

   Typically these are written to a http.Request.
*/
type PatchClusterTemplateInstanceParams struct {

	// Body.
	Body PatchClusterTemplateInstanceBody

	// InstanceID.
	ClusterTemplateInstanceID string

	// ProjectID.
	ProjectID string

	// TemplateID.
	ClusterTemplateID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch cluster template instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchClusterTemplateInstanceParams) WithDefaults() *PatchClusterTemplateInstanceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch cluster template instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchClusterTemplateInstanceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch cluster template instance params
func (o *PatchClusterTemplateInstanceParams) WithTimeout(timeout time.Duration) *PatchClusterTemplateInstanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch cluster template instance params
func (o *PatchClusterTemplateInstanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch cluster template instance params
func (o *PatchClusterTemplateInstanceParams) WithContext(ctx context.Context) *PatchClusterTemplateInstanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch cluster template instance params
func (o *PatchClusterTemplateInstanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch cluster template instance params
func (o *PatchClusterTemplateInstanceParams) WithHTTPClient(client *http.Client) *PatchClusterTemplateInstanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch cluster template instance params
func (o *PatchClusterTemplateInstanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch cluster template instance params
func (o *PatchClusterTemplateInstanceParams) WithBody(body PatchClusterTemplateInstanceBody) *PatchClusterTemplateInstanceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch cluster template instance params
func (o *PatchClusterTemplateInstanceParams) SetBody(body PatchClusterTemplateInstanceBody) {
	o.Body = body
}

// WithClusterTemplateInstanceID adds the instanceID to the patch cluster template instance params
func (o *PatchClusterTemplateInstanceParams) WithClusterTemplateInstanceID(instanceID string) *PatchClusterTemplateInstanceParams {
	o.SetClusterTemplateInstanceID(instanceID)
	return o
}

// SetClusterTemplateInstanceID adds the instanceId to the patch cluster template instance params
func (o *PatchClusterTemplateInstanceParams) SetClusterTemplateInstanceID(instanceID string) {
	o.ClusterTemplateInstanceID = instanceID
}

// WithProjectID adds the projectID to the patch cluster template instance params
func (o *PatchClusterTemplateInstanceParams) WithProjectID(projectID string) *PatchClusterTemplateInstanceParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the patch cluster template instance params
func (o *PatchClusterTemplateInstanceParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WithClusterTemplateID adds the templateID to the patch cluster template instance params
func (o *PatchClusterTemplateInstanceParams) WithClusterTemplateID(templateID string) *PatchClusterTemplateInstanceParams {
	o.SetClusterTemplateID(templateID)
	return o
}

// SetClusterTemplateID adds the templateId to the patch cluster template instance params
func (o *PatchClusterTemplateInstanceParams) SetClusterTemplateID(templateID string) {
	o.ClusterTemplateID = templateID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchClusterTemplateInstanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param instance_id
	if err := r.SetPathParam("instance_id", o.ClusterTemplateInstanceID); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	// path param template_id
	if err := r.SetPathParam("template_id", o.ClusterTemplateID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

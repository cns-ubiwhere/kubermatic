// Code generated by go-swagger; DO NOT EDIT.

package openstack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListOpenstackNetworksNoCredentialsReader is a Reader for the ListOpenstackNetworksNoCredentials structure.
type ListOpenstackNetworksNoCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListOpenstackNetworksNoCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListOpenstackNetworksNoCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListOpenstackNetworksNoCredentialsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListOpenstackNetworksNoCredentialsOK creates a ListOpenstackNetworksNoCredentialsOK with default headers values
func NewListOpenstackNetworksNoCredentialsOK() *ListOpenstackNetworksNoCredentialsOK {
	return &ListOpenstackNetworksNoCredentialsOK{}
}

/* ListOpenstackNetworksNoCredentialsOK describes a response with status code 200, with default header values.

OpenstackNetwork
*/
type ListOpenstackNetworksNoCredentialsOK struct {
	Payload []*models.OpenstackNetwork
}

func (o *ListOpenstackNetworksNoCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/openstack/networks][%d] listOpenstackNetworksNoCredentialsOK  %+v", 200, o.Payload)
}
func (o *ListOpenstackNetworksNoCredentialsOK) GetPayload() []*models.OpenstackNetwork {
	return o.Payload
}

func (o *ListOpenstackNetworksNoCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOpenstackNetworksNoCredentialsDefault creates a ListOpenstackNetworksNoCredentialsDefault with default headers values
func NewListOpenstackNetworksNoCredentialsDefault(code int) *ListOpenstackNetworksNoCredentialsDefault {
	return &ListOpenstackNetworksNoCredentialsDefault{
		_statusCode: code,
	}
}

/* ListOpenstackNetworksNoCredentialsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListOpenstackNetworksNoCredentialsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list openstack networks no credentials default response
func (o *ListOpenstackNetworksNoCredentialsDefault) Code() int {
	return o._statusCode
}

func (o *ListOpenstackNetworksNoCredentialsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/openstack/networks][%d] listOpenstackNetworksNoCredentials default  %+v", o._statusCode, o.Payload)
}
func (o *ListOpenstackNetworksNoCredentialsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListOpenstackNetworksNoCredentialsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

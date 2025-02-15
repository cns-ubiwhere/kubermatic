// Code generated by go-swagger; DO NOT EDIT.

package vsphere

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListVSphereDatastoresReader is a Reader for the ListVSphereDatastores structure.
type ListVSphereDatastoresReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListVSphereDatastoresReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListVSphereDatastoresOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListVSphereDatastoresDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListVSphereDatastoresOK creates a ListVSphereDatastoresOK with default headers values
func NewListVSphereDatastoresOK() *ListVSphereDatastoresOK {
	return &ListVSphereDatastoresOK{}
}

/* ListVSphereDatastoresOK describes a response with status code 200, with default header values.

VSphereDatastoreList
*/
type ListVSphereDatastoresOK struct {
	Payload []*models.VSphereDatastoreList
}

func (o *ListVSphereDatastoresOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/vsphere/datastores][%d] listVSphereDatastoresOK  %+v", 200, o.Payload)
}
func (o *ListVSphereDatastoresOK) GetPayload() []*models.VSphereDatastoreList {
	return o.Payload
}

func (o *ListVSphereDatastoresOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListVSphereDatastoresDefault creates a ListVSphereDatastoresDefault with default headers values
func NewListVSphereDatastoresDefault(code int) *ListVSphereDatastoresDefault {
	return &ListVSphereDatastoresDefault{
		_statusCode: code,
	}
}

/* ListVSphereDatastoresDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListVSphereDatastoresDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list v sphere datastores default response
func (o *ListVSphereDatastoresDefault) Code() int {
	return o._statusCode
}

func (o *ListVSphereDatastoresDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/vsphere/datastores][%d] listVSphereDatastores default  %+v", o._statusCode, o.Payload)
}
func (o *ListVSphereDatastoresDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListVSphereDatastoresDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

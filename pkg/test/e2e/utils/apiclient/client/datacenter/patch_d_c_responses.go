// Code generated by go-swagger; DO NOT EDIT.

package datacenter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// PatchDCReader is a Reader for the PatchDC structure.
type PatchDCReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchDCReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchDCOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchDCUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchDCForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchDCDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchDCOK creates a PatchDCOK with default headers values
func NewPatchDCOK() *PatchDCOK {
	return &PatchDCOK{}
}

/* PatchDCOK describes a response with status code 200, with default header values.

Datacenter
*/
type PatchDCOK struct {
	Payload *models.Datacenter
}

func (o *PatchDCOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/seed/{seed_name}/dc/{dc}][%d] patchDCOK  %+v", 200, o.Payload)
}
func (o *PatchDCOK) GetPayload() *models.Datacenter {
	return o.Payload
}

func (o *PatchDCOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Datacenter)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchDCUnauthorized creates a PatchDCUnauthorized with default headers values
func NewPatchDCUnauthorized() *PatchDCUnauthorized {
	return &PatchDCUnauthorized{}
}

/* PatchDCUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type PatchDCUnauthorized struct {
}

func (o *PatchDCUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/seed/{seed_name}/dc/{dc}][%d] patchDCUnauthorized ", 401)
}

func (o *PatchDCUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchDCForbidden creates a PatchDCForbidden with default headers values
func NewPatchDCForbidden() *PatchDCForbidden {
	return &PatchDCForbidden{}
}

/* PatchDCForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type PatchDCForbidden struct {
}

func (o *PatchDCForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/seed/{seed_name}/dc/{dc}][%d] patchDCForbidden ", 403)
}

func (o *PatchDCForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchDCDefault creates a PatchDCDefault with default headers values
func NewPatchDCDefault(code int) *PatchDCDefault {
	return &PatchDCDefault{
		_statusCode: code,
	}
}

/* PatchDCDefault describes a response with status code -1, with default header values.

errorResponse
*/
type PatchDCDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the patch d c default response
func (o *PatchDCDefault) Code() int {
	return o._statusCode
}

func (o *PatchDCDefault) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/seed/{seed_name}/dc/{dc}][%d] patchDC default  %+v", o._statusCode, o.Payload)
}
func (o *PatchDCDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PatchDCDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

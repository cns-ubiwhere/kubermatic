// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// GetAlertmanagerReader is a Reader for the GetAlertmanager structure.
type GetAlertmanagerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlertmanagerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAlertmanagerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAlertmanagerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAlertmanagerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetAlertmanagerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAlertmanagerOK creates a GetAlertmanagerOK with default headers values
func NewGetAlertmanagerOK() *GetAlertmanagerOK {
	return &GetAlertmanagerOK{}
}

/* GetAlertmanagerOK describes a response with status code 200, with default header values.

Alertmanager
*/
type GetAlertmanagerOK struct {
	Payload *models.Alertmanager
}

func (o *GetAlertmanagerOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/alertmanager/config][%d] getAlertmanagerOK  %+v", 200, o.Payload)
}
func (o *GetAlertmanagerOK) GetPayload() *models.Alertmanager {
	return o.Payload
}

func (o *GetAlertmanagerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Alertmanager)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertmanagerUnauthorized creates a GetAlertmanagerUnauthorized with default headers values
func NewGetAlertmanagerUnauthorized() *GetAlertmanagerUnauthorized {
	return &GetAlertmanagerUnauthorized{}
}

/* GetAlertmanagerUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type GetAlertmanagerUnauthorized struct {
}

func (o *GetAlertmanagerUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/alertmanager/config][%d] getAlertmanagerUnauthorized ", 401)
}

func (o *GetAlertmanagerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAlertmanagerForbidden creates a GetAlertmanagerForbidden with default headers values
func NewGetAlertmanagerForbidden() *GetAlertmanagerForbidden {
	return &GetAlertmanagerForbidden{}
}

/* GetAlertmanagerForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type GetAlertmanagerForbidden struct {
}

func (o *GetAlertmanagerForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/alertmanager/config][%d] getAlertmanagerForbidden ", 403)
}

func (o *GetAlertmanagerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAlertmanagerDefault creates a GetAlertmanagerDefault with default headers values
func NewGetAlertmanagerDefault(code int) *GetAlertmanagerDefault {
	return &GetAlertmanagerDefault{
		_statusCode: code,
	}
}

/* GetAlertmanagerDefault describes a response with status code -1, with default header values.

errorResponse
*/
type GetAlertmanagerDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get alertmanager default response
func (o *GetAlertmanagerDefault) Code() int {
	return o._statusCode
}

func (o *GetAlertmanagerDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/alertmanager/config][%d] getAlertmanager default  %+v", o._statusCode, o.Payload)
}
func (o *GetAlertmanagerDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetAlertmanagerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

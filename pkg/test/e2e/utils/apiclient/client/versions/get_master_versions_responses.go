// Code generated by go-swagger; DO NOT EDIT.

package versions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// GetMasterVersionsReader is a Reader for the GetMasterVersions structure.
type GetMasterVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMasterVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMasterVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetMasterVersionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMasterVersionsOK creates a GetMasterVersionsOK with default headers values
func NewGetMasterVersionsOK() *GetMasterVersionsOK {
	return &GetMasterVersionsOK{}
}

/* GetMasterVersionsOK describes a response with status code 200, with default header values.

MasterVersion
*/
type GetMasterVersionsOK struct {
	Payload []*models.MasterVersion
}

func (o *GetMasterVersionsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/upgrades/cluster][%d] getMasterVersionsOK  %+v", 200, o.Payload)
}
func (o *GetMasterVersionsOK) GetPayload() []*models.MasterVersion {
	return o.Payload
}

func (o *GetMasterVersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMasterVersionsDefault creates a GetMasterVersionsDefault with default headers values
func NewGetMasterVersionsDefault(code int) *GetMasterVersionsDefault {
	return &GetMasterVersionsDefault{
		_statusCode: code,
	}
}

/* GetMasterVersionsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type GetMasterVersionsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get master versions default response
func (o *GetMasterVersionsDefault) Code() int {
	return o._statusCode
}

func (o *GetMasterVersionsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/upgrades/cluster][%d] getMasterVersions default  %+v", o._statusCode, o.Payload)
}
func (o *GetMasterVersionsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetMasterVersionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

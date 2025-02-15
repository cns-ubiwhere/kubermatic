// Code generated by go-swagger; DO NOT EDIT.

package aws

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListAWSSubnetsNoCredentialsReader is a Reader for the ListAWSSubnetsNoCredentials structure.
type ListAWSSubnetsNoCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAWSSubnetsNoCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAWSSubnetsNoCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListAWSSubnetsNoCredentialsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAWSSubnetsNoCredentialsOK creates a ListAWSSubnetsNoCredentialsOK with default headers values
func NewListAWSSubnetsNoCredentialsOK() *ListAWSSubnetsNoCredentialsOK {
	return &ListAWSSubnetsNoCredentialsOK{}
}

/* ListAWSSubnetsNoCredentialsOK describes a response with status code 200, with default header values.

AWSSubnetList
*/
type ListAWSSubnetsNoCredentialsOK struct {
	Payload models.AWSSubnetList
}

func (o *ListAWSSubnetsNoCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/aws/subnets][%d] listAWSSubnetsNoCredentialsOK  %+v", 200, o.Payload)
}
func (o *ListAWSSubnetsNoCredentialsOK) GetPayload() models.AWSSubnetList {
	return o.Payload
}

func (o *ListAWSSubnetsNoCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAWSSubnetsNoCredentialsDefault creates a ListAWSSubnetsNoCredentialsDefault with default headers values
func NewListAWSSubnetsNoCredentialsDefault(code int) *ListAWSSubnetsNoCredentialsDefault {
	return &ListAWSSubnetsNoCredentialsDefault{
		_statusCode: code,
	}
}

/* ListAWSSubnetsNoCredentialsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListAWSSubnetsNoCredentialsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list a w s subnets no credentials default response
func (o *ListAWSSubnetsNoCredentialsDefault) Code() int {
	return o._statusCode
}

func (o *ListAWSSubnetsNoCredentialsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/aws/subnets][%d] listAWSSubnetsNoCredentials default  %+v", o._statusCode, o.Payload)
}
func (o *ListAWSSubnetsNoCredentialsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListAWSSubnetsNoCredentialsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

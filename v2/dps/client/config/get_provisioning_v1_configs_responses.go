// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mariadb-corporation/skysql-api-go/v2/dps/models"
)

// GetProvisioningV1ConfigsReader is a Reader for the GetProvisioningV1Configs structure.
type GetProvisioningV1ConfigsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProvisioningV1ConfigsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProvisioningV1ConfigsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetProvisioningV1ConfigsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetProvisioningV1ConfigsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetProvisioningV1ConfigsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetProvisioningV1ConfigsOK creates a GetProvisioningV1ConfigsOK with default headers values
func NewGetProvisioningV1ConfigsOK() *GetProvisioningV1ConfigsOK {
	return &GetProvisioningV1ConfigsOK{}
}

/* GetProvisioningV1ConfigsOK describes a response with status code 200, with default header values.

OK
*/
type GetProvisioningV1ConfigsOK struct {
	Payload []*models.V1ConfigResponse
}

func (o *GetProvisioningV1ConfigsOK) Error() string {
	return fmt.Sprintf("[GET /provisioning/v1/configs][%d] getProvisioningV1ConfigsOK  %+v", 200, o.Payload)
}
func (o *GetProvisioningV1ConfigsOK) GetPayload() []*models.V1ConfigResponse {
	return o.Payload
}

func (o *GetProvisioningV1ConfigsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProvisioningV1ConfigsBadRequest creates a GetProvisioningV1ConfigsBadRequest with default headers values
func NewGetProvisioningV1ConfigsBadRequest() *GetProvisioningV1ConfigsBadRequest {
	return &GetProvisioningV1ConfigsBadRequest{}
}

/* GetProvisioningV1ConfigsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetProvisioningV1ConfigsBadRequest struct {
	Payload *models.ErrrErrorResponse
}

func (o *GetProvisioningV1ConfigsBadRequest) Error() string {
	return fmt.Sprintf("[GET /provisioning/v1/configs][%d] getProvisioningV1ConfigsBadRequest  %+v", 400, o.Payload)
}
func (o *GetProvisioningV1ConfigsBadRequest) GetPayload() *models.ErrrErrorResponse {
	return o.Payload
}

func (o *GetProvisioningV1ConfigsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrrErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProvisioningV1ConfigsUnauthorized creates a GetProvisioningV1ConfigsUnauthorized with default headers values
func NewGetProvisioningV1ConfigsUnauthorized() *GetProvisioningV1ConfigsUnauthorized {
	return &GetProvisioningV1ConfigsUnauthorized{}
}

/* GetProvisioningV1ConfigsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetProvisioningV1ConfigsUnauthorized struct {
	Payload *models.ErrrErrorResponse
}

func (o *GetProvisioningV1ConfigsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /provisioning/v1/configs][%d] getProvisioningV1ConfigsUnauthorized  %+v", 401, o.Payload)
}
func (o *GetProvisioningV1ConfigsUnauthorized) GetPayload() *models.ErrrErrorResponse {
	return o.Payload
}

func (o *GetProvisioningV1ConfigsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrrErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProvisioningV1ConfigsInternalServerError creates a GetProvisioningV1ConfigsInternalServerError with default headers values
func NewGetProvisioningV1ConfigsInternalServerError() *GetProvisioningV1ConfigsInternalServerError {
	return &GetProvisioningV1ConfigsInternalServerError{}
}

/* GetProvisioningV1ConfigsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetProvisioningV1ConfigsInternalServerError struct {
	Payload *models.ErrrErrorResponse
}

func (o *GetProvisioningV1ConfigsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /provisioning/v1/configs][%d] getProvisioningV1ConfigsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetProvisioningV1ConfigsInternalServerError) GetPayload() *models.ErrrErrorResponse {
	return o.Payload
}

func (o *GetProvisioningV1ConfigsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrrErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

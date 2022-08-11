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

// PostProvisioningV1ConfigsReader is a Reader for the PostProvisioningV1Configs structure.
type PostProvisioningV1ConfigsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostProvisioningV1ConfigsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPostProvisioningV1ConfigsAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostProvisioningV1ConfigsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostProvisioningV1ConfigsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostProvisioningV1ConfigsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostProvisioningV1ConfigsAccepted creates a PostProvisioningV1ConfigsAccepted with default headers values
func NewPostProvisioningV1ConfigsAccepted() *PostProvisioningV1ConfigsAccepted {
	return &PostProvisioningV1ConfigsAccepted{}
}

/* PostProvisioningV1ConfigsAccepted describes a response with status code 202, with default header values.

Accepted
*/
type PostProvisioningV1ConfigsAccepted struct {
	Payload *models.V1ConfigWithValuesResponse
}

func (o *PostProvisioningV1ConfigsAccepted) Error() string {
	return fmt.Sprintf("[POST /provisioning/v1/configs][%d] postProvisioningV1ConfigsAccepted  %+v", 202, o.Payload)
}
func (o *PostProvisioningV1ConfigsAccepted) GetPayload() *models.V1ConfigWithValuesResponse {
	return o.Payload
}

func (o *PostProvisioningV1ConfigsAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ConfigWithValuesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostProvisioningV1ConfigsBadRequest creates a PostProvisioningV1ConfigsBadRequest with default headers values
func NewPostProvisioningV1ConfigsBadRequest() *PostProvisioningV1ConfigsBadRequest {
	return &PostProvisioningV1ConfigsBadRequest{}
}

/* PostProvisioningV1ConfigsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostProvisioningV1ConfigsBadRequest struct {
	Payload *models.ErrrErrorResponse
}

func (o *PostProvisioningV1ConfigsBadRequest) Error() string {
	return fmt.Sprintf("[POST /provisioning/v1/configs][%d] postProvisioningV1ConfigsBadRequest  %+v", 400, o.Payload)
}
func (o *PostProvisioningV1ConfigsBadRequest) GetPayload() *models.ErrrErrorResponse {
	return o.Payload
}

func (o *PostProvisioningV1ConfigsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrrErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostProvisioningV1ConfigsUnauthorized creates a PostProvisioningV1ConfigsUnauthorized with default headers values
func NewPostProvisioningV1ConfigsUnauthorized() *PostProvisioningV1ConfigsUnauthorized {
	return &PostProvisioningV1ConfigsUnauthorized{}
}

/* PostProvisioningV1ConfigsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostProvisioningV1ConfigsUnauthorized struct {
	Payload *models.ErrrErrorResponse
}

func (o *PostProvisioningV1ConfigsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /provisioning/v1/configs][%d] postProvisioningV1ConfigsUnauthorized  %+v", 401, o.Payload)
}
func (o *PostProvisioningV1ConfigsUnauthorized) GetPayload() *models.ErrrErrorResponse {
	return o.Payload
}

func (o *PostProvisioningV1ConfigsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrrErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostProvisioningV1ConfigsInternalServerError creates a PostProvisioningV1ConfigsInternalServerError with default headers values
func NewPostProvisioningV1ConfigsInternalServerError() *PostProvisioningV1ConfigsInternalServerError {
	return &PostProvisioningV1ConfigsInternalServerError{}
}

/* PostProvisioningV1ConfigsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostProvisioningV1ConfigsInternalServerError struct {
	Payload *models.ErrrErrorResponse
}

func (o *PostProvisioningV1ConfigsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /provisioning/v1/configs][%d] postProvisioningV1ConfigsInternalServerError  %+v", 500, o.Payload)
}
func (o *PostProvisioningV1ConfigsInternalServerError) GetPayload() *models.ErrrErrorResponse {
	return o.Payload
}

func (o *PostProvisioningV1ConfigsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrrErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
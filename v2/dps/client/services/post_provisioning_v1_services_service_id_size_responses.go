// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mariadb-corporation/skysql-api-go/v2/dps/models"
)

// PostProvisioningV1ServicesServiceIDSizeReader is a Reader for the PostProvisioningV1ServicesServiceIDSize structure.
type PostProvisioningV1ServicesServiceIDSizeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostProvisioningV1ServicesServiceIDSizeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPostProvisioningV1ServicesServiceIDSizeAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostProvisioningV1ServicesServiceIDSizeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostProvisioningV1ServicesServiceIDSizeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostProvisioningV1ServicesServiceIDSizeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostProvisioningV1ServicesServiceIDSizeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostProvisioningV1ServicesServiceIDSizeAccepted creates a PostProvisioningV1ServicesServiceIDSizeAccepted with default headers values
func NewPostProvisioningV1ServicesServiceIDSizeAccepted() *PostProvisioningV1ServicesServiceIDSizeAccepted {
	return &PostProvisioningV1ServicesServiceIDSizeAccepted{}
}

/* PostProvisioningV1ServicesServiceIDSizeAccepted describes a response with status code 202, with default header values.

Accepted
*/
type PostProvisioningV1ServicesServiceIDSizeAccepted struct {
	Payload *models.V1ServiceSizeState
}

func (o *PostProvisioningV1ServicesServiceIDSizeAccepted) Error() string {
	return fmt.Sprintf("[POST /provisioning/v1/services/{service_id}/size][%d] postProvisioningV1ServicesServiceIdSizeAccepted  %+v", 202, o.Payload)
}
func (o *PostProvisioningV1ServicesServiceIDSizeAccepted) GetPayload() *models.V1ServiceSizeState {
	return o.Payload
}

func (o *PostProvisioningV1ServicesServiceIDSizeAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ServiceSizeState)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostProvisioningV1ServicesServiceIDSizeBadRequest creates a PostProvisioningV1ServicesServiceIDSizeBadRequest with default headers values
func NewPostProvisioningV1ServicesServiceIDSizeBadRequest() *PostProvisioningV1ServicesServiceIDSizeBadRequest {
	return &PostProvisioningV1ServicesServiceIDSizeBadRequest{}
}

/* PostProvisioningV1ServicesServiceIDSizeBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostProvisioningV1ServicesServiceIDSizeBadRequest struct {
	Payload *models.ErrrErrorResponse
}

func (o *PostProvisioningV1ServicesServiceIDSizeBadRequest) Error() string {
	return fmt.Sprintf("[POST /provisioning/v1/services/{service_id}/size][%d] postProvisioningV1ServicesServiceIdSizeBadRequest  %+v", 400, o.Payload)
}
func (o *PostProvisioningV1ServicesServiceIDSizeBadRequest) GetPayload() *models.ErrrErrorResponse {
	return o.Payload
}

func (o *PostProvisioningV1ServicesServiceIDSizeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrrErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostProvisioningV1ServicesServiceIDSizeUnauthorized creates a PostProvisioningV1ServicesServiceIDSizeUnauthorized with default headers values
func NewPostProvisioningV1ServicesServiceIDSizeUnauthorized() *PostProvisioningV1ServicesServiceIDSizeUnauthorized {
	return &PostProvisioningV1ServicesServiceIDSizeUnauthorized{}
}

/* PostProvisioningV1ServicesServiceIDSizeUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostProvisioningV1ServicesServiceIDSizeUnauthorized struct {
	Payload *models.ErrrErrorResponse
}

func (o *PostProvisioningV1ServicesServiceIDSizeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /provisioning/v1/services/{service_id}/size][%d] postProvisioningV1ServicesServiceIdSizeUnauthorized  %+v", 401, o.Payload)
}
func (o *PostProvisioningV1ServicesServiceIDSizeUnauthorized) GetPayload() *models.ErrrErrorResponse {
	return o.Payload
}

func (o *PostProvisioningV1ServicesServiceIDSizeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrrErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostProvisioningV1ServicesServiceIDSizeNotFound creates a PostProvisioningV1ServicesServiceIDSizeNotFound with default headers values
func NewPostProvisioningV1ServicesServiceIDSizeNotFound() *PostProvisioningV1ServicesServiceIDSizeNotFound {
	return &PostProvisioningV1ServicesServiceIDSizeNotFound{}
}

/* PostProvisioningV1ServicesServiceIDSizeNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PostProvisioningV1ServicesServiceIDSizeNotFound struct {
	Payload *models.ErrrErrorResponse
}

func (o *PostProvisioningV1ServicesServiceIDSizeNotFound) Error() string {
	return fmt.Sprintf("[POST /provisioning/v1/services/{service_id}/size][%d] postProvisioningV1ServicesServiceIdSizeNotFound  %+v", 404, o.Payload)
}
func (o *PostProvisioningV1ServicesServiceIDSizeNotFound) GetPayload() *models.ErrrErrorResponse {
	return o.Payload
}

func (o *PostProvisioningV1ServicesServiceIDSizeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrrErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostProvisioningV1ServicesServiceIDSizeInternalServerError creates a PostProvisioningV1ServicesServiceIDSizeInternalServerError with default headers values
func NewPostProvisioningV1ServicesServiceIDSizeInternalServerError() *PostProvisioningV1ServicesServiceIDSizeInternalServerError {
	return &PostProvisioningV1ServicesServiceIDSizeInternalServerError{}
}

/* PostProvisioningV1ServicesServiceIDSizeInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostProvisioningV1ServicesServiceIDSizeInternalServerError struct {
	Payload *models.ErrrErrorResponse
}

func (o *PostProvisioningV1ServicesServiceIDSizeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /provisioning/v1/services/{service_id}/size][%d] postProvisioningV1ServicesServiceIdSizeInternalServerError  %+v", 500, o.Payload)
}
func (o *PostProvisioningV1ServicesServiceIDSizeInternalServerError) GetPayload() *models.ErrrErrorResponse {
	return o.Payload
}

func (o *PostProvisioningV1ServicesServiceIDSizeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrrErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

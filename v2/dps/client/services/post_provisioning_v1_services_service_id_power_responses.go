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

// PostProvisioningV1ServicesServiceIDPowerReader is a Reader for the PostProvisioningV1ServicesServiceIDPower structure.
type PostProvisioningV1ServicesServiceIDPowerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostProvisioningV1ServicesServiceIDPowerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostProvisioningV1ServicesServiceIDPowerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostProvisioningV1ServicesServiceIDPowerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostProvisioningV1ServicesServiceIDPowerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostProvisioningV1ServicesServiceIDPowerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostProvisioningV1ServicesServiceIDPowerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostProvisioningV1ServicesServiceIDPowerOK creates a PostProvisioningV1ServicesServiceIDPowerOK with default headers values
func NewPostProvisioningV1ServicesServiceIDPowerOK() *PostProvisioningV1ServicesServiceIDPowerOK {
	return &PostProvisioningV1ServicesServiceIDPowerOK{}
}

/* PostProvisioningV1ServicesServiceIDPowerOK describes a response with status code 200, with default header values.

OK
*/
type PostProvisioningV1ServicesServiceIDPowerOK struct {
	Payload *models.V1PowerState
}

func (o *PostProvisioningV1ServicesServiceIDPowerOK) Error() string {
	return fmt.Sprintf("[POST /provisioning/v1/services/{service_id}/power][%d] postProvisioningV1ServicesServiceIdPowerOK  %+v", 200, o.Payload)
}
func (o *PostProvisioningV1ServicesServiceIDPowerOK) GetPayload() *models.V1PowerState {
	return o.Payload
}

func (o *PostProvisioningV1ServicesServiceIDPowerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1PowerState)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostProvisioningV1ServicesServiceIDPowerBadRequest creates a PostProvisioningV1ServicesServiceIDPowerBadRequest with default headers values
func NewPostProvisioningV1ServicesServiceIDPowerBadRequest() *PostProvisioningV1ServicesServiceIDPowerBadRequest {
	return &PostProvisioningV1ServicesServiceIDPowerBadRequest{}
}

/* PostProvisioningV1ServicesServiceIDPowerBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostProvisioningV1ServicesServiceIDPowerBadRequest struct {
	Payload *models.ErrrErrorResponse
}

func (o *PostProvisioningV1ServicesServiceIDPowerBadRequest) Error() string {
	return fmt.Sprintf("[POST /provisioning/v1/services/{service_id}/power][%d] postProvisioningV1ServicesServiceIdPowerBadRequest  %+v", 400, o.Payload)
}
func (o *PostProvisioningV1ServicesServiceIDPowerBadRequest) GetPayload() *models.ErrrErrorResponse {
	return o.Payload
}

func (o *PostProvisioningV1ServicesServiceIDPowerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrrErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostProvisioningV1ServicesServiceIDPowerUnauthorized creates a PostProvisioningV1ServicesServiceIDPowerUnauthorized with default headers values
func NewPostProvisioningV1ServicesServiceIDPowerUnauthorized() *PostProvisioningV1ServicesServiceIDPowerUnauthorized {
	return &PostProvisioningV1ServicesServiceIDPowerUnauthorized{}
}

/* PostProvisioningV1ServicesServiceIDPowerUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostProvisioningV1ServicesServiceIDPowerUnauthorized struct {
	Payload *models.ErrrErrorResponse
}

func (o *PostProvisioningV1ServicesServiceIDPowerUnauthorized) Error() string {
	return fmt.Sprintf("[POST /provisioning/v1/services/{service_id}/power][%d] postProvisioningV1ServicesServiceIdPowerUnauthorized  %+v", 401, o.Payload)
}
func (o *PostProvisioningV1ServicesServiceIDPowerUnauthorized) GetPayload() *models.ErrrErrorResponse {
	return o.Payload
}

func (o *PostProvisioningV1ServicesServiceIDPowerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrrErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostProvisioningV1ServicesServiceIDPowerNotFound creates a PostProvisioningV1ServicesServiceIDPowerNotFound with default headers values
func NewPostProvisioningV1ServicesServiceIDPowerNotFound() *PostProvisioningV1ServicesServiceIDPowerNotFound {
	return &PostProvisioningV1ServicesServiceIDPowerNotFound{}
}

/* PostProvisioningV1ServicesServiceIDPowerNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PostProvisioningV1ServicesServiceIDPowerNotFound struct {
	Payload *models.ErrrErrorResponse
}

func (o *PostProvisioningV1ServicesServiceIDPowerNotFound) Error() string {
	return fmt.Sprintf("[POST /provisioning/v1/services/{service_id}/power][%d] postProvisioningV1ServicesServiceIdPowerNotFound  %+v", 404, o.Payload)
}
func (o *PostProvisioningV1ServicesServiceIDPowerNotFound) GetPayload() *models.ErrrErrorResponse {
	return o.Payload
}

func (o *PostProvisioningV1ServicesServiceIDPowerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrrErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostProvisioningV1ServicesServiceIDPowerInternalServerError creates a PostProvisioningV1ServicesServiceIDPowerInternalServerError with default headers values
func NewPostProvisioningV1ServicesServiceIDPowerInternalServerError() *PostProvisioningV1ServicesServiceIDPowerInternalServerError {
	return &PostProvisioningV1ServicesServiceIDPowerInternalServerError{}
}

/* PostProvisioningV1ServicesServiceIDPowerInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostProvisioningV1ServicesServiceIDPowerInternalServerError struct {
	Payload *models.ErrrErrorResponse
}

func (o *PostProvisioningV1ServicesServiceIDPowerInternalServerError) Error() string {
	return fmt.Sprintf("[POST /provisioning/v1/services/{service_id}/power][%d] postProvisioningV1ServicesServiceIdPowerInternalServerError  %+v", 500, o.Payload)
}
func (o *PostProvisioningV1ServicesServiceIDPowerInternalServerError) GetPayload() *models.ErrrErrorResponse {
	return o.Payload
}

func (o *PostProvisioningV1ServicesServiceIDPowerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrrErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

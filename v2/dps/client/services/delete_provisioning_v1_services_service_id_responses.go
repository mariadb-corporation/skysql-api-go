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

// DeleteProvisioningV1ServicesServiceIDReader is a Reader for the DeleteProvisioningV1ServicesServiceID structure.
type DeleteProvisioningV1ServicesServiceIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProvisioningV1ServicesServiceIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteProvisioningV1ServicesServiceIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteProvisioningV1ServicesServiceIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteProvisioningV1ServicesServiceIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteProvisioningV1ServicesServiceIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteProvisioningV1ServicesServiceIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteProvisioningV1ServicesServiceIDOK creates a DeleteProvisioningV1ServicesServiceIDOK with default headers values
func NewDeleteProvisioningV1ServicesServiceIDOK() *DeleteProvisioningV1ServicesServiceIDOK {
	return &DeleteProvisioningV1ServicesServiceIDOK{}
}

/* DeleteProvisioningV1ServicesServiceIDOK describes a response with status code 200, with default header values.

OK
*/
type DeleteProvisioningV1ServicesServiceIDOK struct {
}

func (o *DeleteProvisioningV1ServicesServiceIDOK) Error() string {
	return fmt.Sprintf("[DELETE /provisioning/v1/services/{service_id}][%d] deleteProvisioningV1ServicesServiceIdOK ", 200)
}

func (o *DeleteProvisioningV1ServicesServiceIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProvisioningV1ServicesServiceIDBadRequest creates a DeleteProvisioningV1ServicesServiceIDBadRequest with default headers values
func NewDeleteProvisioningV1ServicesServiceIDBadRequest() *DeleteProvisioningV1ServicesServiceIDBadRequest {
	return &DeleteProvisioningV1ServicesServiceIDBadRequest{}
}

/* DeleteProvisioningV1ServicesServiceIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteProvisioningV1ServicesServiceIDBadRequest struct {
	Payload *models.ErrrErrorResponse
}

func (o *DeleteProvisioningV1ServicesServiceIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /provisioning/v1/services/{service_id}][%d] deleteProvisioningV1ServicesServiceIdBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteProvisioningV1ServicesServiceIDBadRequest) GetPayload() *models.ErrrErrorResponse {
	return o.Payload
}

func (o *DeleteProvisioningV1ServicesServiceIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrrErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProvisioningV1ServicesServiceIDUnauthorized creates a DeleteProvisioningV1ServicesServiceIDUnauthorized with default headers values
func NewDeleteProvisioningV1ServicesServiceIDUnauthorized() *DeleteProvisioningV1ServicesServiceIDUnauthorized {
	return &DeleteProvisioningV1ServicesServiceIDUnauthorized{}
}

/* DeleteProvisioningV1ServicesServiceIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteProvisioningV1ServicesServiceIDUnauthorized struct {
	Payload *models.ErrrErrorResponse
}

func (o *DeleteProvisioningV1ServicesServiceIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /provisioning/v1/services/{service_id}][%d] deleteProvisioningV1ServicesServiceIdUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteProvisioningV1ServicesServiceIDUnauthorized) GetPayload() *models.ErrrErrorResponse {
	return o.Payload
}

func (o *DeleteProvisioningV1ServicesServiceIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrrErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProvisioningV1ServicesServiceIDNotFound creates a DeleteProvisioningV1ServicesServiceIDNotFound with default headers values
func NewDeleteProvisioningV1ServicesServiceIDNotFound() *DeleteProvisioningV1ServicesServiceIDNotFound {
	return &DeleteProvisioningV1ServicesServiceIDNotFound{}
}

/* DeleteProvisioningV1ServicesServiceIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteProvisioningV1ServicesServiceIDNotFound struct {
	Payload *models.ErrrErrorResponse
}

func (o *DeleteProvisioningV1ServicesServiceIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /provisioning/v1/services/{service_id}][%d] deleteProvisioningV1ServicesServiceIdNotFound  %+v", 404, o.Payload)
}
func (o *DeleteProvisioningV1ServicesServiceIDNotFound) GetPayload() *models.ErrrErrorResponse {
	return o.Payload
}

func (o *DeleteProvisioningV1ServicesServiceIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrrErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProvisioningV1ServicesServiceIDInternalServerError creates a DeleteProvisioningV1ServicesServiceIDInternalServerError with default headers values
func NewDeleteProvisioningV1ServicesServiceIDInternalServerError() *DeleteProvisioningV1ServicesServiceIDInternalServerError {
	return &DeleteProvisioningV1ServicesServiceIDInternalServerError{}
}

/* DeleteProvisioningV1ServicesServiceIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteProvisioningV1ServicesServiceIDInternalServerError struct {
	Payload *models.ErrrErrorResponse
}

func (o *DeleteProvisioningV1ServicesServiceIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /provisioning/v1/services/{service_id}][%d] deleteProvisioningV1ServicesServiceIdInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteProvisioningV1ServicesServiceIDInternalServerError) GetPayload() *models.ErrrErrorResponse {
	return o.Payload
}

func (o *DeleteProvisioningV1ServicesServiceIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrrErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
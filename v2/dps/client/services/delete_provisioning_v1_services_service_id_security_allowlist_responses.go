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

// DeleteProvisioningV1ServicesServiceIDSecurityAllowlistReader is a Reader for the DeleteProvisioningV1ServicesServiceIDSecurityAllowlist structure.
type DeleteProvisioningV1ServicesServiceIDSecurityAllowlistReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProvisioningV1ServicesServiceIDSecurityAllowlistReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteProvisioningV1ServicesServiceIDSecurityAllowlistOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteProvisioningV1ServicesServiceIDSecurityAllowlistBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteProvisioningV1ServicesServiceIDSecurityAllowlistUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteProvisioningV1ServicesServiceIDSecurityAllowlistInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteProvisioningV1ServicesServiceIDSecurityAllowlistOK creates a DeleteProvisioningV1ServicesServiceIDSecurityAllowlistOK with default headers values
func NewDeleteProvisioningV1ServicesServiceIDSecurityAllowlistOK() *DeleteProvisioningV1ServicesServiceIDSecurityAllowlistOK {
	return &DeleteProvisioningV1ServicesServiceIDSecurityAllowlistOK{}
}

/* DeleteProvisioningV1ServicesServiceIDSecurityAllowlistOK describes a response with status code 200, with default header values.

OK
*/
type DeleteProvisioningV1ServicesServiceIDSecurityAllowlistOK struct {
	Payload []*models.V1AllowListEndpoint
}

func (o *DeleteProvisioningV1ServicesServiceIDSecurityAllowlistOK) Error() string {
	return fmt.Sprintf("[DELETE /provisioning/v1/services/{service_id}/security/allowlist][%d] deleteProvisioningV1ServicesServiceIdSecurityAllowlistOK  %+v", 200, o.Payload)
}
func (o *DeleteProvisioningV1ServicesServiceIDSecurityAllowlistOK) GetPayload() []*models.V1AllowListEndpoint {
	return o.Payload
}

func (o *DeleteProvisioningV1ServicesServiceIDSecurityAllowlistOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProvisioningV1ServicesServiceIDSecurityAllowlistBadRequest creates a DeleteProvisioningV1ServicesServiceIDSecurityAllowlistBadRequest with default headers values
func NewDeleteProvisioningV1ServicesServiceIDSecurityAllowlistBadRequest() *DeleteProvisioningV1ServicesServiceIDSecurityAllowlistBadRequest {
	return &DeleteProvisioningV1ServicesServiceIDSecurityAllowlistBadRequest{}
}

/* DeleteProvisioningV1ServicesServiceIDSecurityAllowlistBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteProvisioningV1ServicesServiceIDSecurityAllowlistBadRequest struct {
	Payload *models.ErrrErrorResponse
}

func (o *DeleteProvisioningV1ServicesServiceIDSecurityAllowlistBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /provisioning/v1/services/{service_id}/security/allowlist][%d] deleteProvisioningV1ServicesServiceIdSecurityAllowlistBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteProvisioningV1ServicesServiceIDSecurityAllowlistBadRequest) GetPayload() *models.ErrrErrorResponse {
	return o.Payload
}

func (o *DeleteProvisioningV1ServicesServiceIDSecurityAllowlistBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrrErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProvisioningV1ServicesServiceIDSecurityAllowlistUnauthorized creates a DeleteProvisioningV1ServicesServiceIDSecurityAllowlistUnauthorized with default headers values
func NewDeleteProvisioningV1ServicesServiceIDSecurityAllowlistUnauthorized() *DeleteProvisioningV1ServicesServiceIDSecurityAllowlistUnauthorized {
	return &DeleteProvisioningV1ServicesServiceIDSecurityAllowlistUnauthorized{}
}

/* DeleteProvisioningV1ServicesServiceIDSecurityAllowlistUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteProvisioningV1ServicesServiceIDSecurityAllowlistUnauthorized struct {
	Payload *models.ErrrErrorResponse
}

func (o *DeleteProvisioningV1ServicesServiceIDSecurityAllowlistUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /provisioning/v1/services/{service_id}/security/allowlist][%d] deleteProvisioningV1ServicesServiceIdSecurityAllowlistUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteProvisioningV1ServicesServiceIDSecurityAllowlistUnauthorized) GetPayload() *models.ErrrErrorResponse {
	return o.Payload
}

func (o *DeleteProvisioningV1ServicesServiceIDSecurityAllowlistUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrrErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProvisioningV1ServicesServiceIDSecurityAllowlistInternalServerError creates a DeleteProvisioningV1ServicesServiceIDSecurityAllowlistInternalServerError with default headers values
func NewDeleteProvisioningV1ServicesServiceIDSecurityAllowlistInternalServerError() *DeleteProvisioningV1ServicesServiceIDSecurityAllowlistInternalServerError {
	return &DeleteProvisioningV1ServicesServiceIDSecurityAllowlistInternalServerError{}
}

/* DeleteProvisioningV1ServicesServiceIDSecurityAllowlistInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteProvisioningV1ServicesServiceIDSecurityAllowlistInternalServerError struct {
	Payload *models.ErrrErrorResponse
}

func (o *DeleteProvisioningV1ServicesServiceIDSecurityAllowlistInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /provisioning/v1/services/{service_id}/security/allowlist][%d] deleteProvisioningV1ServicesServiceIdSecurityAllowlistInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteProvisioningV1ServicesServiceIDSecurityAllowlistInternalServerError) GetPayload() *models.ErrrErrorResponse {
	return o.Payload
}

func (o *DeleteProvisioningV1ServicesServiceIDSecurityAllowlistInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrrErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mariadb-corporation/skysql-api-go/v2/dps/models"
)

// GetProvisioningV1InternalAccountsOrgIDProviderReader is a Reader for the GetProvisioningV1InternalAccountsOrgIDProvider structure.
type GetProvisioningV1InternalAccountsOrgIDProviderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProvisioningV1InternalAccountsOrgIDProviderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProvisioningV1InternalAccountsOrgIDProviderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetProvisioningV1InternalAccountsOrgIDProviderBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetProvisioningV1InternalAccountsOrgIDProviderNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetProvisioningV1InternalAccountsOrgIDProviderInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetProvisioningV1InternalAccountsOrgIDProviderOK creates a GetProvisioningV1InternalAccountsOrgIDProviderOK with default headers values
func NewGetProvisioningV1InternalAccountsOrgIDProviderOK() *GetProvisioningV1InternalAccountsOrgIDProviderOK {
	return &GetProvisioningV1InternalAccountsOrgIDProviderOK{}
}

/* GetProvisioningV1InternalAccountsOrgIDProviderOK describes a response with status code 200, with default header values.

OK
*/
type GetProvisioningV1InternalAccountsOrgIDProviderOK struct {
	Payload *models.V1AccountResponse
}

func (o *GetProvisioningV1InternalAccountsOrgIDProviderOK) Error() string {
	return fmt.Sprintf("[GET /provisioning/v1/internal/accounts/{org_id}/{provider}][%d] getProvisioningV1InternalAccountsOrgIdProviderOK  %+v", 200, o.Payload)
}
func (o *GetProvisioningV1InternalAccountsOrgIDProviderOK) GetPayload() *models.V1AccountResponse {
	return o.Payload
}

func (o *GetProvisioningV1InternalAccountsOrgIDProviderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1AccountResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProvisioningV1InternalAccountsOrgIDProviderBadRequest creates a GetProvisioningV1InternalAccountsOrgIDProviderBadRequest with default headers values
func NewGetProvisioningV1InternalAccountsOrgIDProviderBadRequest() *GetProvisioningV1InternalAccountsOrgIDProviderBadRequest {
	return &GetProvisioningV1InternalAccountsOrgIDProviderBadRequest{}
}

/* GetProvisioningV1InternalAccountsOrgIDProviderBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetProvisioningV1InternalAccountsOrgIDProviderBadRequest struct {
	Payload *models.ErrrErrorResponse
}

func (o *GetProvisioningV1InternalAccountsOrgIDProviderBadRequest) Error() string {
	return fmt.Sprintf("[GET /provisioning/v1/internal/accounts/{org_id}/{provider}][%d] getProvisioningV1InternalAccountsOrgIdProviderBadRequest  %+v", 400, o.Payload)
}
func (o *GetProvisioningV1InternalAccountsOrgIDProviderBadRequest) GetPayload() *models.ErrrErrorResponse {
	return o.Payload
}

func (o *GetProvisioningV1InternalAccountsOrgIDProviderBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrrErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProvisioningV1InternalAccountsOrgIDProviderNotFound creates a GetProvisioningV1InternalAccountsOrgIDProviderNotFound with default headers values
func NewGetProvisioningV1InternalAccountsOrgIDProviderNotFound() *GetProvisioningV1InternalAccountsOrgIDProviderNotFound {
	return &GetProvisioningV1InternalAccountsOrgIDProviderNotFound{}
}

/* GetProvisioningV1InternalAccountsOrgIDProviderNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetProvisioningV1InternalAccountsOrgIDProviderNotFound struct {
	Payload *models.ErrrErrorResponse
}

func (o *GetProvisioningV1InternalAccountsOrgIDProviderNotFound) Error() string {
	return fmt.Sprintf("[GET /provisioning/v1/internal/accounts/{org_id}/{provider}][%d] getProvisioningV1InternalAccountsOrgIdProviderNotFound  %+v", 404, o.Payload)
}
func (o *GetProvisioningV1InternalAccountsOrgIDProviderNotFound) GetPayload() *models.ErrrErrorResponse {
	return o.Payload
}

func (o *GetProvisioningV1InternalAccountsOrgIDProviderNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrrErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProvisioningV1InternalAccountsOrgIDProviderInternalServerError creates a GetProvisioningV1InternalAccountsOrgIDProviderInternalServerError with default headers values
func NewGetProvisioningV1InternalAccountsOrgIDProviderInternalServerError() *GetProvisioningV1InternalAccountsOrgIDProviderInternalServerError {
	return &GetProvisioningV1InternalAccountsOrgIDProviderInternalServerError{}
}

/* GetProvisioningV1InternalAccountsOrgIDProviderInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetProvisioningV1InternalAccountsOrgIDProviderInternalServerError struct {
	Payload *models.ErrrErrorResponse
}

func (o *GetProvisioningV1InternalAccountsOrgIDProviderInternalServerError) Error() string {
	return fmt.Sprintf("[GET /provisioning/v1/internal/accounts/{org_id}/{provider}][%d] getProvisioningV1InternalAccountsOrgIdProviderInternalServerError  %+v", 500, o.Payload)
}
func (o *GetProvisioningV1InternalAccountsOrgIDProviderInternalServerError) GetPayload() *models.ErrrErrorResponse {
	return o.Payload
}

func (o *GetProvisioningV1InternalAccountsOrgIDProviderInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrrErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

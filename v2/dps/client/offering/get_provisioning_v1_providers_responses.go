// Code generated by go-swagger; DO NOT EDIT.

package offering

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mariadb-corporation/skysql-api-go/v2/dps/models"
)

// GetProvisioningV1ProvidersReader is a Reader for the GetProvisioningV1Providers structure.
type GetProvisioningV1ProvidersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProvisioningV1ProvidersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProvisioningV1ProvidersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetProvisioningV1ProvidersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetProvisioningV1ProvidersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetProvisioningV1ProvidersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetProvisioningV1ProvidersOK creates a GetProvisioningV1ProvidersOK with default headers values
func NewGetProvisioningV1ProvidersOK() *GetProvisioningV1ProvidersOK {
	return &GetProvisioningV1ProvidersOK{}
}

/* GetProvisioningV1ProvidersOK describes a response with status code 200, with default header values.

OK
*/
type GetProvisioningV1ProvidersOK struct {
	Payload []*models.V1Provider
}

func (o *GetProvisioningV1ProvidersOK) Error() string {
	return fmt.Sprintf("[GET /provisioning/v1/providers][%d] getProvisioningV1ProvidersOK  %+v", 200, o.Payload)
}
func (o *GetProvisioningV1ProvidersOK) GetPayload() []*models.V1Provider {
	return o.Payload
}

func (o *GetProvisioningV1ProvidersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProvisioningV1ProvidersBadRequest creates a GetProvisioningV1ProvidersBadRequest with default headers values
func NewGetProvisioningV1ProvidersBadRequest() *GetProvisioningV1ProvidersBadRequest {
	return &GetProvisioningV1ProvidersBadRequest{}
}

/* GetProvisioningV1ProvidersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetProvisioningV1ProvidersBadRequest struct {
	Payload *models.ErrrErrorResponse
}

func (o *GetProvisioningV1ProvidersBadRequest) Error() string {
	return fmt.Sprintf("[GET /provisioning/v1/providers][%d] getProvisioningV1ProvidersBadRequest  %+v", 400, o.Payload)
}
func (o *GetProvisioningV1ProvidersBadRequest) GetPayload() *models.ErrrErrorResponse {
	return o.Payload
}

func (o *GetProvisioningV1ProvidersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrrErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProvisioningV1ProvidersUnauthorized creates a GetProvisioningV1ProvidersUnauthorized with default headers values
func NewGetProvisioningV1ProvidersUnauthorized() *GetProvisioningV1ProvidersUnauthorized {
	return &GetProvisioningV1ProvidersUnauthorized{}
}

/* GetProvisioningV1ProvidersUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetProvisioningV1ProvidersUnauthorized struct {
	Payload *models.ErrrErrorResponse
}

func (o *GetProvisioningV1ProvidersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /provisioning/v1/providers][%d] getProvisioningV1ProvidersUnauthorized  %+v", 401, o.Payload)
}
func (o *GetProvisioningV1ProvidersUnauthorized) GetPayload() *models.ErrrErrorResponse {
	return o.Payload
}

func (o *GetProvisioningV1ProvidersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrrErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProvisioningV1ProvidersInternalServerError creates a GetProvisioningV1ProvidersInternalServerError with default headers values
func NewGetProvisioningV1ProvidersInternalServerError() *GetProvisioningV1ProvidersInternalServerError {
	return &GetProvisioningV1ProvidersInternalServerError{}
}

/* GetProvisioningV1ProvidersInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetProvisioningV1ProvidersInternalServerError struct {
	Payload *models.ErrrErrorResponse
}

func (o *GetProvisioningV1ProvidersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /provisioning/v1/providers][%d] getProvisioningV1ProvidersInternalServerError  %+v", 500, o.Payload)
}
func (o *GetProvisioningV1ProvidersInternalServerError) GetPayload() *models.ErrrErrorResponse {
	return o.Payload
}

func (o *GetProvisioningV1ProvidersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrrErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
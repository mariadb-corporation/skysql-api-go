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

// GetProvisioningV1SizesReader is a Reader for the GetProvisioningV1Sizes structure.
type GetProvisioningV1SizesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProvisioningV1SizesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProvisioningV1SizesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetProvisioningV1SizesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetProvisioningV1SizesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetProvisioningV1SizesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetProvisioningV1SizesOK creates a GetProvisioningV1SizesOK with default headers values
func NewGetProvisioningV1SizesOK() *GetProvisioningV1SizesOK {
	return &GetProvisioningV1SizesOK{}
}

/* GetProvisioningV1SizesOK describes a response with status code 200, with default header values.

OK
*/
type GetProvisioningV1SizesOK struct {
	Payload []*models.V1Size
}

func (o *GetProvisioningV1SizesOK) Error() string {
	return fmt.Sprintf("[GET /provisioning/v1/sizes][%d] getProvisioningV1SizesOK  %+v", 200, o.Payload)
}
func (o *GetProvisioningV1SizesOK) GetPayload() []*models.V1Size {
	return o.Payload
}

func (o *GetProvisioningV1SizesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProvisioningV1SizesBadRequest creates a GetProvisioningV1SizesBadRequest with default headers values
func NewGetProvisioningV1SizesBadRequest() *GetProvisioningV1SizesBadRequest {
	return &GetProvisioningV1SizesBadRequest{}
}

/* GetProvisioningV1SizesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetProvisioningV1SizesBadRequest struct {
	Payload *models.ErrrErrorResponse
}

func (o *GetProvisioningV1SizesBadRequest) Error() string {
	return fmt.Sprintf("[GET /provisioning/v1/sizes][%d] getProvisioningV1SizesBadRequest  %+v", 400, o.Payload)
}
func (o *GetProvisioningV1SizesBadRequest) GetPayload() *models.ErrrErrorResponse {
	return o.Payload
}

func (o *GetProvisioningV1SizesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrrErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProvisioningV1SizesUnauthorized creates a GetProvisioningV1SizesUnauthorized with default headers values
func NewGetProvisioningV1SizesUnauthorized() *GetProvisioningV1SizesUnauthorized {
	return &GetProvisioningV1SizesUnauthorized{}
}

/* GetProvisioningV1SizesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetProvisioningV1SizesUnauthorized struct {
	Payload *models.ErrrErrorResponse
}

func (o *GetProvisioningV1SizesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /provisioning/v1/sizes][%d] getProvisioningV1SizesUnauthorized  %+v", 401, o.Payload)
}
func (o *GetProvisioningV1SizesUnauthorized) GetPayload() *models.ErrrErrorResponse {
	return o.Payload
}

func (o *GetProvisioningV1SizesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrrErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProvisioningV1SizesInternalServerError creates a GetProvisioningV1SizesInternalServerError with default headers values
func NewGetProvisioningV1SizesInternalServerError() *GetProvisioningV1SizesInternalServerError {
	return &GetProvisioningV1SizesInternalServerError{}
}

/* GetProvisioningV1SizesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetProvisioningV1SizesInternalServerError struct {
	Payload *models.ErrrErrorResponse
}

func (o *GetProvisioningV1SizesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /provisioning/v1/sizes][%d] getProvisioningV1SizesInternalServerError  %+v", 500, o.Payload)
}
func (o *GetProvisioningV1SizesInternalServerError) GetPayload() *models.ErrrErrorResponse {
	return o.Payload
}

func (o *GetProvisioningV1SizesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrrErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

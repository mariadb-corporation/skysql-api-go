// Code generated by go-swagger; DO NOT EDIT.

package delete_channel_email

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mariadb-corporation/skysql-api-go/v2/nos/models"
)

// DeleteNotificationsV1ChannelsEmailsChannelIDReader is a Reader for the DeleteNotificationsV1ChannelsEmailsChannelID structure.
type DeleteNotificationsV1ChannelsEmailsChannelIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNotificationsV1ChannelsEmailsChannelIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteNotificationsV1ChannelsEmailsChannelIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteNotificationsV1ChannelsEmailsChannelIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteNotificationsV1ChannelsEmailsChannelIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteNotificationsV1ChannelsEmailsChannelIDOK creates a DeleteNotificationsV1ChannelsEmailsChannelIDOK with default headers values
func NewDeleteNotificationsV1ChannelsEmailsChannelIDOK() *DeleteNotificationsV1ChannelsEmailsChannelIDOK {
	return &DeleteNotificationsV1ChannelsEmailsChannelIDOK{}
}

/* DeleteNotificationsV1ChannelsEmailsChannelIDOK describes a response with status code 200, with default header values.

DeleteNotificationsV1ChannelsEmailsChannelIDOK delete notifications v1 channels emails channel Id o k
*/
type DeleteNotificationsV1ChannelsEmailsChannelIDOK struct {
}

func (o *DeleteNotificationsV1ChannelsEmailsChannelIDOK) Error() string {
	return fmt.Sprintf("[DELETE /notifications/v1/channels/emails/{channel_id}][%d] deleteNotificationsV1ChannelsEmailsChannelIdOK ", 200)
}

func (o *DeleteNotificationsV1ChannelsEmailsChannelIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNotificationsV1ChannelsEmailsChannelIDUnauthorized creates a DeleteNotificationsV1ChannelsEmailsChannelIDUnauthorized with default headers values
func NewDeleteNotificationsV1ChannelsEmailsChannelIDUnauthorized() *DeleteNotificationsV1ChannelsEmailsChannelIDUnauthorized {
	return &DeleteNotificationsV1ChannelsEmailsChannelIDUnauthorized{}
}

/* DeleteNotificationsV1ChannelsEmailsChannelIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteNotificationsV1ChannelsEmailsChannelIDUnauthorized struct {
	Payload *models.ErrrErrorResponse
}

func (o *DeleteNotificationsV1ChannelsEmailsChannelIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /notifications/v1/channels/emails/{channel_id}][%d] deleteNotificationsV1ChannelsEmailsChannelIdUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteNotificationsV1ChannelsEmailsChannelIDUnauthorized) GetPayload() *models.ErrrErrorResponse {
	return o.Payload
}

func (o *DeleteNotificationsV1ChannelsEmailsChannelIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrrErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNotificationsV1ChannelsEmailsChannelIDNotFound creates a DeleteNotificationsV1ChannelsEmailsChannelIDNotFound with default headers values
func NewDeleteNotificationsV1ChannelsEmailsChannelIDNotFound() *DeleteNotificationsV1ChannelsEmailsChannelIDNotFound {
	return &DeleteNotificationsV1ChannelsEmailsChannelIDNotFound{}
}

/* DeleteNotificationsV1ChannelsEmailsChannelIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteNotificationsV1ChannelsEmailsChannelIDNotFound struct {
	Payload *models.ErrrErrorResponse
}

func (o *DeleteNotificationsV1ChannelsEmailsChannelIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /notifications/v1/channels/emails/{channel_id}][%d] deleteNotificationsV1ChannelsEmailsChannelIdNotFound  %+v", 404, o.Payload)
}
func (o *DeleteNotificationsV1ChannelsEmailsChannelIDNotFound) GetPayload() *models.ErrrErrorResponse {
	return o.Payload
}

func (o *DeleteNotificationsV1ChannelsEmailsChannelIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrrErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

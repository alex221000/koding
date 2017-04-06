package j_user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// JUserWhoamiReader is a Reader for the JUserWhoami structure.
type JUserWhoamiReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JUserWhoamiReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewJUserWhoamiOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewJUserWhoamiUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJUserWhoamiOK creates a JUserWhoamiOK with default headers values
func NewJUserWhoamiOK() *JUserWhoamiOK {
	return &JUserWhoamiOK{}
}

/*JUserWhoamiOK handles this case with default header values.

Request processed successfully
*/
type JUserWhoamiOK struct {
	Payload *models.DefaultResponse
}

func (o *JUserWhoamiOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JUser.whoami][%d] jUserWhoamiOK  %+v", 200, o.Payload)
}

func (o *JUserWhoamiOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewJUserWhoamiUnauthorized creates a JUserWhoamiUnauthorized with default headers values
func NewJUserWhoamiUnauthorized() *JUserWhoamiUnauthorized {
	return &JUserWhoamiUnauthorized{}
}

/*JUserWhoamiUnauthorized handles this case with default header values.

Unauthorized request
*/
type JUserWhoamiUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *JUserWhoamiUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/JUser.whoami][%d] jUserWhoamiUnauthorized  %+v", 401, o.Payload)
}

func (o *JUserWhoamiUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

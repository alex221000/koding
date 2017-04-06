package j_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// JAccountSetEmailPreferencesReader is a Reader for the JAccountSetEmailPreferences structure.
type JAccountSetEmailPreferencesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JAccountSetEmailPreferencesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewJAccountSetEmailPreferencesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJAccountSetEmailPreferencesOK creates a JAccountSetEmailPreferencesOK with default headers values
func NewJAccountSetEmailPreferencesOK() *JAccountSetEmailPreferencesOK {
	return &JAccountSetEmailPreferencesOK{}
}

/*JAccountSetEmailPreferencesOK handles this case with default header values.

OK
*/
type JAccountSetEmailPreferencesOK struct {
	Payload JAccountSetEmailPreferencesOKBody
}

func (o *JAccountSetEmailPreferencesOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JAccount.setEmailPreferences/{id}][%d] jAccountSetEmailPreferencesOK  %+v", 200, o.Payload)
}

func (o *JAccountSetEmailPreferencesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*JAccountSetEmailPreferencesOKBody j account set email preferences o k body
swagger:model JAccountSetEmailPreferencesOKBody
*/
type JAccountSetEmailPreferencesOKBody struct {
	models.JAccount

	models.DefaultResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *JAccountSetEmailPreferencesOKBody) UnmarshalJSON(raw []byte) error {

	var jAccountSetEmailPreferencesOKBodyAO0 models.JAccount
	if err := swag.ReadJSON(raw, &jAccountSetEmailPreferencesOKBodyAO0); err != nil {
		return err
	}
	o.JAccount = jAccountSetEmailPreferencesOKBodyAO0

	var jAccountSetEmailPreferencesOKBodyAO1 models.DefaultResponse
	if err := swag.ReadJSON(raw, &jAccountSetEmailPreferencesOKBodyAO1); err != nil {
		return err
	}
	o.DefaultResponse = jAccountSetEmailPreferencesOKBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o JAccountSetEmailPreferencesOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	jAccountSetEmailPreferencesOKBodyAO0, err := swag.WriteJSON(o.JAccount)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jAccountSetEmailPreferencesOKBodyAO0)

	jAccountSetEmailPreferencesOKBodyAO1, err := swag.WriteJSON(o.DefaultResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jAccountSetEmailPreferencesOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this j account set email preferences o k body
func (o *JAccountSetEmailPreferencesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.JAccount.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.DefaultResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

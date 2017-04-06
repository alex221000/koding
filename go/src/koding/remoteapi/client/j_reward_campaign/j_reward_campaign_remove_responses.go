package j_reward_campaign

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

// JRewardCampaignRemoveReader is a Reader for the JRewardCampaignRemove structure.
type JRewardCampaignRemoveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JRewardCampaignRemoveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewJRewardCampaignRemoveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJRewardCampaignRemoveOK creates a JRewardCampaignRemoveOK with default headers values
func NewJRewardCampaignRemoveOK() *JRewardCampaignRemoveOK {
	return &JRewardCampaignRemoveOK{}
}

/*JRewardCampaignRemoveOK handles this case with default header values.

OK
*/
type JRewardCampaignRemoveOK struct {
	Payload JRewardCampaignRemoveOKBody
}

func (o *JRewardCampaignRemoveOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JRewardCampaign.remove/{id}][%d] jRewardCampaignRemoveOK  %+v", 200, o.Payload)
}

func (o *JRewardCampaignRemoveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*JRewardCampaignRemoveOKBody j reward campaign remove o k body
swagger:model JRewardCampaignRemoveOKBody
*/
type JRewardCampaignRemoveOKBody struct {
	models.JRewardCampaign

	models.DefaultResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *JRewardCampaignRemoveOKBody) UnmarshalJSON(raw []byte) error {

	var jRewardCampaignRemoveOKBodyAO0 models.JRewardCampaign
	if err := swag.ReadJSON(raw, &jRewardCampaignRemoveOKBodyAO0); err != nil {
		return err
	}
	o.JRewardCampaign = jRewardCampaignRemoveOKBodyAO0

	var jRewardCampaignRemoveOKBodyAO1 models.DefaultResponse
	if err := swag.ReadJSON(raw, &jRewardCampaignRemoveOKBodyAO1); err != nil {
		return err
	}
	o.DefaultResponse = jRewardCampaignRemoveOKBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o JRewardCampaignRemoveOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	jRewardCampaignRemoveOKBodyAO0, err := swag.WriteJSON(o.JRewardCampaign)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jRewardCampaignRemoveOKBodyAO0)

	jRewardCampaignRemoveOKBodyAO1, err := swag.WriteJSON(o.DefaultResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jRewardCampaignRemoveOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this j reward campaign remove o k body
func (o *JRewardCampaignRemoveOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.JRewardCampaign.Validate(formats); err != nil {
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

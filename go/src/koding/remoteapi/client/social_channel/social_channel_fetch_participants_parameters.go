package social_channel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// NewSocialChannelFetchParticipantsParams creates a new SocialChannelFetchParticipantsParams object
// with the default values initialized.
func NewSocialChannelFetchParticipantsParams() *SocialChannelFetchParticipantsParams {
	var ()
	return &SocialChannelFetchParticipantsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSocialChannelFetchParticipantsParamsWithTimeout creates a new SocialChannelFetchParticipantsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSocialChannelFetchParticipantsParamsWithTimeout(timeout time.Duration) *SocialChannelFetchParticipantsParams {
	var ()
	return &SocialChannelFetchParticipantsParams{

		timeout: timeout,
	}
}

// NewSocialChannelFetchParticipantsParamsWithContext creates a new SocialChannelFetchParticipantsParams object
// with the default values initialized, and the ability to set a context for a request
func NewSocialChannelFetchParticipantsParamsWithContext(ctx context.Context) *SocialChannelFetchParticipantsParams {
	var ()
	return &SocialChannelFetchParticipantsParams{

		Context: ctx,
	}
}

/*SocialChannelFetchParticipantsParams contains all the parameters to send to the API endpoint
for the social channel fetch participants operation typically these are written to a http.Request
*/
type SocialChannelFetchParticipantsParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the social channel fetch participants params
func (o *SocialChannelFetchParticipantsParams) WithTimeout(timeout time.Duration) *SocialChannelFetchParticipantsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the social channel fetch participants params
func (o *SocialChannelFetchParticipantsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the social channel fetch participants params
func (o *SocialChannelFetchParticipantsParams) WithContext(ctx context.Context) *SocialChannelFetchParticipantsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the social channel fetch participants params
func (o *SocialChannelFetchParticipantsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the social channel fetch participants params
func (o *SocialChannelFetchParticipantsParams) WithBody(body models.DefaultSelector) *SocialChannelFetchParticipantsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the social channel fetch participants params
func (o *SocialChannelFetchParticipantsParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SocialChannelFetchParticipantsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

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

// NewSocialChannelFetchProfileFeedCountParams creates a new SocialChannelFetchProfileFeedCountParams object
// with the default values initialized.
func NewSocialChannelFetchProfileFeedCountParams() *SocialChannelFetchProfileFeedCountParams {
	var ()
	return &SocialChannelFetchProfileFeedCountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSocialChannelFetchProfileFeedCountParamsWithTimeout creates a new SocialChannelFetchProfileFeedCountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSocialChannelFetchProfileFeedCountParamsWithTimeout(timeout time.Duration) *SocialChannelFetchProfileFeedCountParams {
	var ()
	return &SocialChannelFetchProfileFeedCountParams{

		timeout: timeout,
	}
}

// NewSocialChannelFetchProfileFeedCountParamsWithContext creates a new SocialChannelFetchProfileFeedCountParams object
// with the default values initialized, and the ability to set a context for a request
func NewSocialChannelFetchProfileFeedCountParamsWithContext(ctx context.Context) *SocialChannelFetchProfileFeedCountParams {
	var ()
	return &SocialChannelFetchProfileFeedCountParams{

		Context: ctx,
	}
}

/*SocialChannelFetchProfileFeedCountParams contains all the parameters to send to the API endpoint
for the social channel fetch profile feed count operation typically these are written to a http.Request
*/
type SocialChannelFetchProfileFeedCountParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the social channel fetch profile feed count params
func (o *SocialChannelFetchProfileFeedCountParams) WithTimeout(timeout time.Duration) *SocialChannelFetchProfileFeedCountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the social channel fetch profile feed count params
func (o *SocialChannelFetchProfileFeedCountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the social channel fetch profile feed count params
func (o *SocialChannelFetchProfileFeedCountParams) WithContext(ctx context.Context) *SocialChannelFetchProfileFeedCountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the social channel fetch profile feed count params
func (o *SocialChannelFetchProfileFeedCountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the social channel fetch profile feed count params
func (o *SocialChannelFetchProfileFeedCountParams) WithBody(body models.DefaultSelector) *SocialChannelFetchProfileFeedCountParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the social channel fetch profile feed count params
func (o *SocialChannelFetchProfileFeedCountParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SocialChannelFetchProfileFeedCountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

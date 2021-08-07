// Code generated by go-swagger; DO NOT EDIT.

package tag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openflagr/flagr/swagger_gen/models"
)

// CreateTagOKCode is the HTTP code returned for type CreateTagOK
const CreateTagOKCode int = 200

/*CreateTagOK tag just created

swagger:response createTagOK
*/
type CreateTagOK struct {

	/*
	  In: Body
	*/
	Payload *models.Tag `json:"body,omitempty"`
}

// NewCreateTagOK creates CreateTagOK with default headers values
func NewCreateTagOK() *CreateTagOK {

	return &CreateTagOK{}
}

// WithPayload adds the payload to the create tag o k response
func (o *CreateTagOK) WithPayload(payload *models.Tag) *CreateTagOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create tag o k response
func (o *CreateTagOK) SetPayload(payload *models.Tag) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTagOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreateTagDefault generic error response

swagger:response createTagDefault
*/
type CreateTagDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateTagDefault creates CreateTagDefault with default headers values
func NewCreateTagDefault(code int) *CreateTagDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateTagDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create tag default response
func (o *CreateTagDefault) WithStatusCode(code int) *CreateTagDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create tag default response
func (o *CreateTagDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create tag default response
func (o *CreateTagDefault) WithPayload(payload *models.Error) *CreateTagDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create tag default response
func (o *CreateTagDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTagDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// Code generated by go-swagger; DO NOT EDIT.

package flag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openflagr/flagr/swagger_gen/models"
)

// GetFlagEntityTypesOKCode is the HTTP code returned for type GetFlagEntityTypesOK
const GetFlagEntityTypesOKCode int = 200

/*GetFlagEntityTypesOK returns all the FlagEntityTypes

swagger:response getFlagEntityTypesOK
*/
type GetFlagEntityTypesOK struct {

	/*
	  In: Body
	*/
	Payload []string `json:"body,omitempty"`
}

// NewGetFlagEntityTypesOK creates GetFlagEntityTypesOK with default headers values
func NewGetFlagEntityTypesOK() *GetFlagEntityTypesOK {

	return &GetFlagEntityTypesOK{}
}

// WithPayload adds the payload to the get flag entity types o k response
func (o *GetFlagEntityTypesOK) WithPayload(payload []string) *GetFlagEntityTypesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get flag entity types o k response
func (o *GetFlagEntityTypesOK) SetPayload(payload []string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFlagEntityTypesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]string, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*GetFlagEntityTypesDefault generic error response

swagger:response getFlagEntityTypesDefault
*/
type GetFlagEntityTypesDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetFlagEntityTypesDefault creates GetFlagEntityTypesDefault with default headers values
func NewGetFlagEntityTypesDefault(code int) *GetFlagEntityTypesDefault {
	if code <= 0 {
		code = 500
	}

	return &GetFlagEntityTypesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get flag entity types default response
func (o *GetFlagEntityTypesDefault) WithStatusCode(code int) *GetFlagEntityTypesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get flag entity types default response
func (o *GetFlagEntityTypesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get flag entity types default response
func (o *GetFlagEntityTypesDefault) WithPayload(payload *models.Error) *GetFlagEntityTypesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get flag entity types default response
func (o *GetFlagEntityTypesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFlagEntityTypesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

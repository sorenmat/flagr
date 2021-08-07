// Code generated by go-swagger; DO NOT EDIT.

package variant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openflagr/flagr/swagger_gen/models"
)

// DeleteVariantOKCode is the HTTP code returned for type DeleteVariantOK
const DeleteVariantOKCode int = 200

/*DeleteVariantOK deleted

swagger:response deleteVariantOK
*/
type DeleteVariantOK struct {
}

// NewDeleteVariantOK creates DeleteVariantOK with default headers values
func NewDeleteVariantOK() *DeleteVariantOK {

	return &DeleteVariantOK{}
}

// WriteResponse to the client
func (o *DeleteVariantOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

/*DeleteVariantDefault generic error response

swagger:response deleteVariantDefault
*/
type DeleteVariantDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteVariantDefault creates DeleteVariantDefault with default headers values
func NewDeleteVariantDefault(code int) *DeleteVariantDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteVariantDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete variant default response
func (o *DeleteVariantDefault) WithStatusCode(code int) *DeleteVariantDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete variant default response
func (o *DeleteVariantDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete variant default response
func (o *DeleteVariantDefault) WithPayload(payload *models.Error) *DeleteVariantDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete variant default response
func (o *DeleteVariantDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteVariantDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

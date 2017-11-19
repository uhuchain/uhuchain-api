// Code generated by go-swagger; DO NOT EDIT.

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/uhuchain/uhuchain/models"
)

// AddPolicyCreatedCode is the HTTP code returned for type AddPolicyCreated
const AddPolicyCreatedCode int = 201

/*AddPolicyCreated Policy was created

swagger:response addPolicyCreated
*/
type AddPolicyCreated struct {
}

// NewAddPolicyCreated creates AddPolicyCreated with default headers values
func NewAddPolicyCreated() *AddPolicyCreated {
	return &AddPolicyCreated{}
}

// WriteResponse to the client
func (o *AddPolicyCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
}

// AddPolicyMethodNotAllowedCode is the HTTP code returned for type AddPolicyMethodNotAllowed
const AddPolicyMethodNotAllowedCode int = 405

/*AddPolicyMethodNotAllowed Invalid input

swagger:response addPolicyMethodNotAllowed
*/
type AddPolicyMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewAddPolicyMethodNotAllowed creates AddPolicyMethodNotAllowed with default headers values
func NewAddPolicyMethodNotAllowed() *AddPolicyMethodNotAllowed {
	return &AddPolicyMethodNotAllowed{}
}

// WithPayload adds the payload to the add policy method not allowed response
func (o *AddPolicyMethodNotAllowed) WithPayload(payload *models.APIResponse) *AddPolicyMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add policy method not allowed response
func (o *AddPolicyMethodNotAllowed) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddPolicyMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

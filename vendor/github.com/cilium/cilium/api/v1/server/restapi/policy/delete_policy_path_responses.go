package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/cilium/api/v1/models"
)

// DeletePolicyPathNoContentCode is the HTTP code returned for type DeletePolicyPathNoContent
const DeletePolicyPathNoContentCode int = 204

/*DeletePolicyPathNoContent Success

swagger:response deletePolicyPathNoContent
*/
type DeletePolicyPathNoContent struct {
}

// NewDeletePolicyPathNoContent creates DeletePolicyPathNoContent with default headers values
func NewDeletePolicyPathNoContent() *DeletePolicyPathNoContent {
	return &DeletePolicyPathNoContent{}
}

// WriteResponse to the client
func (o *DeletePolicyPathNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}

// DeletePolicyPathInvalidCode is the HTTP code returned for type DeletePolicyPathInvalid
const DeletePolicyPathInvalidCode int = 400

/*DeletePolicyPathInvalid Invalid request

swagger:response deletePolicyPathInvalid
*/
type DeletePolicyPathInvalid struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewDeletePolicyPathInvalid creates DeletePolicyPathInvalid with default headers values
func NewDeletePolicyPathInvalid() *DeletePolicyPathInvalid {
	return &DeletePolicyPathInvalid{}
}

// WithPayload adds the payload to the delete policy path invalid response
func (o *DeletePolicyPathInvalid) WithPayload(payload models.Error) *DeletePolicyPathInvalid {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete policy path invalid response
func (o *DeletePolicyPathInvalid) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeletePolicyPathInvalid) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// DeletePolicyPathNotFoundCode is the HTTP code returned for type DeletePolicyPathNotFound
const DeletePolicyPathNotFoundCode int = 404

/*DeletePolicyPathNotFound Policy tree not found

swagger:response deletePolicyPathNotFound
*/
type DeletePolicyPathNotFound struct {
}

// NewDeletePolicyPathNotFound creates DeletePolicyPathNotFound with default headers values
func NewDeletePolicyPathNotFound() *DeletePolicyPathNotFound {
	return &DeletePolicyPathNotFound{}
}

// WriteResponse to the client
func (o *DeletePolicyPathNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}

// DeletePolicyPathFailureCode is the HTTP code returned for type DeletePolicyPathFailure
const DeletePolicyPathFailureCode int = 500

/*DeletePolicyPathFailure Error while deleting policy

swagger:response deletePolicyPathFailure
*/
type DeletePolicyPathFailure struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewDeletePolicyPathFailure creates DeletePolicyPathFailure with default headers values
func NewDeletePolicyPathFailure() *DeletePolicyPathFailure {
	return &DeletePolicyPathFailure{}
}

// WithPayload adds the payload to the delete policy path failure response
func (o *DeletePolicyPathFailure) WithPayload(payload models.Error) *DeletePolicyPathFailure {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete policy path failure response
func (o *DeletePolicyPathFailure) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeletePolicyPathFailure) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

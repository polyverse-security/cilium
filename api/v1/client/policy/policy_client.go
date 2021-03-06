package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new policy API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for policy API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeletePolicyPath deletes a policy sub tree
*/
func (a *Client) DeletePolicyPath(params *DeletePolicyPathParams) (*DeletePolicyPathNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePolicyPathParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeletePolicyPath",
		Method:             "DELETE",
		PathPattern:        "/policy/{path}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeletePolicyPathReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeletePolicyPathNoContent), nil

}

/*
GetIdentity retrieves identity by labels
*/
func (a *Client) GetIdentity(params *GetIdentityParams) (*GetIdentityOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIdentityParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetIdentity",
		Method:             "GET",
		PathPattern:        "/identity",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetIdentityReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetIdentityOK), nil

}

/*
GetIdentityID retrieves identity
*/
func (a *Client) GetIdentityID(params *GetIdentityIDParams) (*GetIdentityIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIdentityIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetIdentityID",
		Method:             "GET",
		PathPattern:        "/identity/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetIdentityIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetIdentityIDOK), nil

}

/*
GetPolicy retrieves entire policy tree

Returns the entire policy tree with all children.

*/
func (a *Client) GetPolicy(params *GetPolicyParams) (*GetPolicyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPolicyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPolicy",
		Method:             "GET",
		PathPattern:        "/policy",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPolicyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPolicyOK), nil

}

/*
GetPolicyPath retrieves policy subtree

Returns the policy node at path with all its children

*/
func (a *Client) GetPolicyPath(params *GetPolicyPathParams) (*GetPolicyPathOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPolicyPathParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPolicyPath",
		Method:             "GET",
		PathPattern:        "/policy/{path}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPolicyPathReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPolicyPathOK), nil

}

/*
GetPolicyResolve resolves policy for an identity context
*/
func (a *Client) GetPolicyResolve(params *GetPolicyResolveParams) (*GetPolicyResolveOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPolicyResolveParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPolicyResolve",
		Method:             "GET",
		PathPattern:        "/policy/resolve",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPolicyResolveReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPolicyResolveOK), nil

}

/*
PutPolicyPath creates or update a policy sub tree
*/
func (a *Client) PutPolicyPath(params *PutPolicyPathParams) (*PutPolicyPathOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutPolicyPathParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutPolicyPath",
		Method:             "PUT",
		PathPattern:        "/policy/{path}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutPolicyPathReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutPolicyPathOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

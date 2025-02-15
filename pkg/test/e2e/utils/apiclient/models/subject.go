// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Subject Subject contains a reference to the object or user identities a role binding applies to.  This can either hold a direct API object reference,
//
// or a value for non-objects such as user and group names.
//
// swagger:model Subject
type Subject struct {

	// APIGroup holds the API group of the referenced subject.
	// Defaults to "" for ServiceAccount subjects.
	// Defaults to "rbac.authorization.k8s.io" for User and Group subjects.
	// +optional
	APIGroup string `json:"apiGroup,omitempty"`

	// Kind of object being referenced. Values defined by this API group are "User", "Group", and "ServiceAccount".
	// If the Authorizer does not recognized the kind value, the Authorizer should report an error.
	Kind string `json:"kind,omitempty"`

	// Name of the object being referenced.
	Name string `json:"name,omitempty"`

	// Namespace of the referenced object.  If the object kind is non-namespace, such as "User" or "Group", and this value is not empty
	// the Authorizer should report an error.
	// +optional
	Namespace string `json:"namespace,omitempty"`
}

// Validate validates this subject
func (m *Subject) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this subject based on context it is used
func (m *Subject) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Subject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Subject) UnmarshalBinary(b []byte) error {
	var res Subject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

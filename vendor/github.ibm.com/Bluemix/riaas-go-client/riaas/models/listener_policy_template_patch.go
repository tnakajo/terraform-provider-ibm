// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ListenerPolicyTemplatePatch ListenerPolicyTemplatePatch
// swagger:model ListenerPolicyTemplatePatch
type ListenerPolicyTemplatePatch struct {

	// The user-defined name for this policy. Names must be unique within the load balancer listener the policy resides in.
	Name string `json:"name,omitempty"`

	// Priority of the policy. Lower value indicates higher priority.
	Priority int64 `json:"priority,omitempty"`

	// target
	Target *ListenerPolicyTemplatePatchTarget `json:"target,omitempty"`
}

// Validate validates this listener policy template patch
func (m *ListenerPolicyTemplatePatch) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTarget(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListenerPolicyTemplatePatch) validateTarget(formats strfmt.Registry) error {

	if swag.IsZero(m.Target) { // not required
		return nil
	}

	if m.Target != nil {
		if err := m.Target.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListenerPolicyTemplatePatch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListenerPolicyTemplatePatch) UnmarshalBinary(b []byte) error {
	var res ListenerPolicyTemplatePatch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

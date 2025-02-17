// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DrugCategoriesResponse drug categories response
//
// swagger:model DrugCategoriesResponse
type DrugCategoriesResponse struct {

	// description
	Description string `json:"description,omitempty"`

	// id
	ID uint64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this drug categories response
func (m *DrugCategoriesResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this drug categories response based on context it is used
func (m *DrugCategoriesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DrugCategoriesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DrugCategoriesResponse) UnmarshalBinary(b []byte) error {
	var res DrugCategoriesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

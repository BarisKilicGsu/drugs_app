// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DrugsResponse drugs response
//
// swagger:model DrugsResponse
type DrugsResponse struct {

	// description
	Description string `json:"description,omitempty"`

	// drug category description
	DrugCategoryDescription string `json:"drug_category_description,omitempty"`

	// drug category id
	DrugCategoryID uint64 `json:"drug_category_id,omitempty"`

	// drug category name
	DrugCategoryName string `json:"drug_category_name,omitempty"`

	// id
	ID uint64 `json:"id,omitempty"`

	// image url
	ImageURL string `json:"image_url,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// side effects
	SideEffects string `json:"side_effects,omitempty"`

	// warnings
	Warnings string `json:"warnings,omitempty"`
}

// Validate validates this drugs response
func (m *DrugsResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this drugs response based on context it is used
func (m *DrugsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DrugsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DrugsResponse) UnmarshalBinary(b []byte) error {
	var res DrugsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

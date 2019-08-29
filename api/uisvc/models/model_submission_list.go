// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ModelSubmissionList model submission list
// swagger:model modelSubmissionList
type ModelSubmissionList []*ModelSubmissionListItems0

// Validate validates this model submission list
func (m ModelSubmissionList) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {
		if swag.IsZero(m[i]) { // not required
			continue
		}

		if m[i] != nil {
			if err := m[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ModelSubmissionListItems0 model submission list items0
// swagger:model ModelSubmissionListItems0
type ModelSubmissionListItems0 struct {

	// base ref
	BaseRef *ModelSubmissionListItems0BaseRef `json:"base_ref,omitempty"`

	// canceled
	Canceled bool `json:"canceled,omitempty"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// finished at
	// Format: date-time
	FinishedAt *strfmt.DateTime `json:"finished_at,omitempty"`

	// head ref
	HeadRef *ModelSubmissionListItems0HeadRef `json:"head_ref,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// started at
	// Format: date-time
	StartedAt *strfmt.DateTime `json:"started_at,omitempty"`

	// status
	Status *bool `json:"status,omitempty"`

	// tasks count
	TasksCount int64 `json:"tasks_count,omitempty"`

	// user
	User *ModelSubmissionListItems0User `json:"user,omitempty"`
}

// Validate validates this model submission list items0
func (m *ModelSubmissionListItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBaseRef(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFinishedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHeadRef(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelSubmissionListItems0) validateBaseRef(formats strfmt.Registry) error {

	if swag.IsZero(m.BaseRef) { // not required
		return nil
	}

	if m.BaseRef != nil {
		if err := m.BaseRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("base_ref")
			}
			return err
		}
	}

	return nil
}

func (m *ModelSubmissionListItems0) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ModelSubmissionListItems0) validateFinishedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.FinishedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("finished_at", "body", "date-time", m.FinishedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ModelSubmissionListItems0) validateHeadRef(formats strfmt.Registry) error {

	if swag.IsZero(m.HeadRef) { // not required
		return nil
	}

	if m.HeadRef != nil {
		if err := m.HeadRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("head_ref")
			}
			return err
		}
	}

	return nil
}

func (m *ModelSubmissionListItems0) validateStartedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.StartedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("started_at", "body", "date-time", m.StartedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ModelSubmissionListItems0) validateUser(formats strfmt.Registry) error {

	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelSubmissionListItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelSubmissionListItems0) UnmarshalBinary(b []byte) error {
	var res ModelSubmissionListItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ModelSubmissionListItems0BaseRef model submission list items0 base ref
// swagger:model ModelSubmissionListItems0BaseRef
type ModelSubmissionListItems0BaseRef struct {

	// id
	ID int64 `json:"id,omitempty"`

	// ref name
	RefName string `json:"ref_name,omitempty"`

	// repository
	Repository *ModelSubmissionListItems0BaseRefRepository `json:"repository,omitempty"`

	// sha
	Sha string `json:"sha,omitempty"`
}

// Validate validates this model submission list items0 base ref
func (m *ModelSubmissionListItems0BaseRef) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRepository(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelSubmissionListItems0BaseRef) validateRepository(formats strfmt.Registry) error {

	if swag.IsZero(m.Repository) { // not required
		return nil
	}

	if m.Repository != nil {
		if err := m.Repository.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("base_ref" + "." + "repository")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelSubmissionListItems0BaseRef) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelSubmissionListItems0BaseRef) UnmarshalBinary(b []byte) error {
	var res ModelSubmissionListItems0BaseRef
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ModelSubmissionListItems0BaseRefRepository model submission list items0 base ref repository
// swagger:model ModelSubmissionListItems0BaseRefRepository
type ModelSubmissionListItems0BaseRefRepository struct {

	// auto created
	AutoCreated bool `json:"auto_created,omitempty"`

	// disabled
	Disabled bool `json:"disabled,omitempty"`

	// github
	Github interface{} `json:"github,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// private
	Private bool `json:"private,omitempty"`
}

// Validate validates this model submission list items0 base ref repository
func (m *ModelSubmissionListItems0BaseRefRepository) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModelSubmissionListItems0BaseRefRepository) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelSubmissionListItems0BaseRefRepository) UnmarshalBinary(b []byte) error {
	var res ModelSubmissionListItems0BaseRefRepository
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ModelSubmissionListItems0HeadRef model submission list items0 head ref
// swagger:model ModelSubmissionListItems0HeadRef
type ModelSubmissionListItems0HeadRef struct {

	// id
	ID int64 `json:"id,omitempty"`

	// ref name
	RefName string `json:"ref_name,omitempty"`

	// repository
	Repository *ModelSubmissionListItems0HeadRefRepository `json:"repository,omitempty"`

	// sha
	Sha string `json:"sha,omitempty"`
}

// Validate validates this model submission list items0 head ref
func (m *ModelSubmissionListItems0HeadRef) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRepository(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelSubmissionListItems0HeadRef) validateRepository(formats strfmt.Registry) error {

	if swag.IsZero(m.Repository) { // not required
		return nil
	}

	if m.Repository != nil {
		if err := m.Repository.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("head_ref" + "." + "repository")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelSubmissionListItems0HeadRef) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelSubmissionListItems0HeadRef) UnmarshalBinary(b []byte) error {
	var res ModelSubmissionListItems0HeadRef
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ModelSubmissionListItems0HeadRefRepository model submission list items0 head ref repository
// swagger:model ModelSubmissionListItems0HeadRefRepository
type ModelSubmissionListItems0HeadRefRepository struct {

	// auto created
	AutoCreated bool `json:"auto_created,omitempty"`

	// disabled
	Disabled bool `json:"disabled,omitempty"`

	// github
	Github interface{} `json:"github,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// private
	Private bool `json:"private,omitempty"`
}

// Validate validates this model submission list items0 head ref repository
func (m *ModelSubmissionListItems0HeadRefRepository) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModelSubmissionListItems0HeadRefRepository) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelSubmissionListItems0HeadRefRepository) UnmarshalBinary(b []byte) error {
	var res ModelSubmissionListItems0HeadRefRepository
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ModelSubmissionListItems0User model submission list items0 user
// swagger:model ModelSubmissionListItems0User
type ModelSubmissionListItems0User struct {

	// errors
	Errors []*ModelSubmissionListItems0UserErrorsItems0 `json:"errors"`

	// id
	ID int64 `json:"id,omitempty"`

	// last scanned repos
	// Format: date-time
	LastScannedRepos *strfmt.DateTime `json:"last_scanned_repos,omitempty"`

	// token
	Token interface{} `json:"token,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this model submission list items0 user
func (m *ModelSubmissionListItems0User) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastScannedRepos(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelSubmissionListItems0User) validateErrors(formats strfmt.Registry) error {

	if swag.IsZero(m.Errors) { // not required
		return nil
	}

	for i := 0; i < len(m.Errors); i++ {
		if swag.IsZero(m.Errors[i]) { // not required
			continue
		}

		if m.Errors[i] != nil {
			if err := m.Errors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("user" + "." + "errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModelSubmissionListItems0User) validateLastScannedRepos(formats strfmt.Registry) error {

	if swag.IsZero(m.LastScannedRepos) { // not required
		return nil
	}

	if err := validate.FormatOf("user"+"."+"last_scanned_repos", "body", "date-time", m.LastScannedRepos.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelSubmissionListItems0User) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelSubmissionListItems0User) UnmarshalBinary(b []byte) error {
	var res ModelSubmissionListItems0User
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ModelSubmissionListItems0UserErrorsItems0 model submission list items0 user errors items0
// swagger:model ModelSubmissionListItems0UserErrorsItems0
type ModelSubmissionListItems0UserErrorsItems0 struct {

	// error
	Error string `json:"error,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`
}

// Validate validates this model submission list items0 user errors items0
func (m *ModelSubmissionListItems0UserErrorsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModelSubmissionListItems0UserErrorsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelSubmissionListItems0UserErrorsItems0) UnmarshalBinary(b []byte) error {
	var res ModelSubmissionListItems0UserErrorsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

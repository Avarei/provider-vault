/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/resource"
	"github.com/upbound/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this Audit
func (mg *Audit) GetTerraformResourceType() string {
	return "vault_audit"
}

// GetConnectionDetailsMapping for this Audit
func (tr *Audit) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this Audit
func (tr *Audit) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Audit
func (tr *Audit) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Audit
func (tr *Audit) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Audit
func (tr *Audit) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Audit
func (tr *Audit) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this Audit
func (tr *Audit) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this Audit using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Audit) LateInitialize(attrs []byte) (bool, error) {
	params := &AuditParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Audit) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this Mount
func (mg *Mount) GetTerraformResourceType() string {
	return "vault_mount"
}

// GetConnectionDetailsMapping for this Mount
func (tr *Mount) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this Mount
func (tr *Mount) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Mount
func (tr *Mount) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Mount
func (tr *Mount) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Mount
func (tr *Mount) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Mount
func (tr *Mount) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this Mount
func (tr *Mount) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this Mount using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Mount) LateInitialize(attrs []byte) (bool, error) {
	params := &MountParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Mount) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this VaultNamespace
func (mg *VaultNamespace) GetTerraformResourceType() string {
	return "vault_namespace"
}

// GetConnectionDetailsMapping for this VaultNamespace
func (tr *VaultNamespace) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this VaultNamespace
func (tr *VaultNamespace) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this VaultNamespace
func (tr *VaultNamespace) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this VaultNamespace
func (tr *VaultNamespace) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this VaultNamespace
func (tr *VaultNamespace) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this VaultNamespace
func (tr *VaultNamespace) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this VaultNamespace
func (tr *VaultNamespace) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this VaultNamespace using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *VaultNamespace) LateInitialize(attrs []byte) (bool, error) {
	params := &VaultNamespaceParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *VaultNamespace) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this Policy
func (mg *Policy) GetTerraformResourceType() string {
	return "vault_policy"
}

// GetConnectionDetailsMapping for this Policy
func (tr *Policy) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this Policy
func (tr *Policy) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Policy
func (tr *Policy) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Policy
func (tr *Policy) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Policy
func (tr *Policy) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Policy
func (tr *Policy) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this Policy
func (tr *Policy) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this Policy using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Policy) LateInitialize(attrs []byte) (bool, error) {
	params := &PolicyParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Policy) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this Token
func (mg *Token) GetTerraformResourceType() string {
	return "vault_token"
}

// GetConnectionDetailsMapping for this Token
func (tr *Token) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"client_token": "status.atProvider.clientToken", "wrapped_token": "status.atProvider.wrappedToken", "wrapping_accessor": "status.atProvider.wrappingAccessor"}
}

// GetObservation of this Token
func (tr *Token) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Token
func (tr *Token) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Token
func (tr *Token) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Token
func (tr *Token) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Token
func (tr *Token) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this Token
func (tr *Token) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this Token using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Token) LateInitialize(attrs []byte) (bool, error) {
	params := &TokenParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Token) GetTerraformSchemaVersion() int {
	return 0
}

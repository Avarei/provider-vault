/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type PingidInitParameters struct {

	// The mount to tie this method to for use in automatic mappings.
	// The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
	// The mount to tie this method to for use in automatic mappings. The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
	MountAccessor *string `json:"mountAccessor,omitempty" tf:"mount_accessor,omitempty"`

	// (string: <required>) – Name of the MFA method.
	// Name of the MFA method.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// A base64-encoded third-party settings file retrieved
	// from PingID's configuration page.
	// A base64-encoded third-party settings file retrieved from PingID's configuration page.
	SettingsFileBase64 *string `json:"settingsFileBase64,omitempty" tf:"settings_file_base64,omitempty"`

	// A format string for mapping Identity names to MFA method names.
	// Values to substitute should be placed in {{}}. For example, "{{alias.name}}@example.com".
	// If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
	// A format string for mapping Identity names to MFA method names. Values to substitute should be placed in `{{}}`.
	UsernameFormat *string `json:"usernameFormat,omitempty" tf:"username_format,omitempty"`
}

type PingidObservation struct {

	// (string) – Admin URL computed by Vault
	// Admin URL computed by Vault.
	AdminURL *string `json:"adminUrl,omitempty" tf:"admin_url,omitempty"`

	// (string) – Authenticator URL computed by Vault
	// Authenticator URL computed by Vault.
	AuthenticatorURL *string `json:"authenticatorUrl,omitempty" tf:"authenticator_url,omitempty"`

	// (string) – ID computed by Vault
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (string) – IDP URL computed by Vault
	// IDP URL computed by Vault.
	IdpURL *string `json:"idpUrl,omitempty" tf:"idp_url,omitempty"`

	// The mount to tie this method to for use in automatic mappings.
	// The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
	// The mount to tie this method to for use in automatic mappings. The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
	MountAccessor *string `json:"mountAccessor,omitempty" tf:"mount_accessor,omitempty"`

	// (string: <required>) – Name of the MFA method.
	// Name of the MFA method.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// (string) – Namespace ID computed by Vault
	// Namespace ID computed by Vault.
	NamespaceID *string `json:"namespaceId,omitempty" tf:"namespace_id,omitempty"`

	// (string) – Org Alias computed by Vault
	// Org Alias computed by Vault.
	OrgAlias *string `json:"orgAlias,omitempty" tf:"org_alias,omitempty"`

	// A base64-encoded third-party settings file retrieved
	// from PingID's configuration page.
	// A base64-encoded third-party settings file retrieved from PingID's configuration page.
	SettingsFileBase64 *string `json:"settingsFileBase64,omitempty" tf:"settings_file_base64,omitempty"`

	// (string) – Type of configuration computed by Vault
	// Type of configuration computed by Vault.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (string) – If set to true, enables use of PingID signature. Computed by Vault
	// If set, enables use of PingID signature. Computed by Vault
	UseSignature *bool `json:"useSignature,omitempty" tf:"use_signature,omitempty"`

	// A format string for mapping Identity names to MFA method names.
	// Values to substitute should be placed in {{}}. For example, "{{alias.name}}@example.com".
	// If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
	// A format string for mapping Identity names to MFA method names. Values to substitute should be placed in `{{}}`.
	UsernameFormat *string `json:"usernameFormat,omitempty" tf:"username_format,omitempty"`
}

type PingidParameters struct {

	// The mount to tie this method to for use in automatic mappings.
	// The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
	// The mount to tie this method to for use in automatic mappings. The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
	// +kubebuilder:validation:Optional
	MountAccessor *string `json:"mountAccessor,omitempty" tf:"mount_accessor,omitempty"`

	// (string: <required>) – Name of the MFA method.
	// Name of the MFA method.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// A base64-encoded third-party settings file retrieved
	// from PingID's configuration page.
	// A base64-encoded third-party settings file retrieved from PingID's configuration page.
	// +kubebuilder:validation:Optional
	SettingsFileBase64 *string `json:"settingsFileBase64,omitempty" tf:"settings_file_base64,omitempty"`

	// A format string for mapping Identity names to MFA method names.
	// Values to substitute should be placed in {{}}. For example, "{{alias.name}}@example.com".
	// If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
	// A format string for mapping Identity names to MFA method names. Values to substitute should be placed in `{{}}`.
	// +kubebuilder:validation:Optional
	UsernameFormat *string `json:"usernameFormat,omitempty" tf:"username_format,omitempty"`
}

// PingidSpec defines the desired state of Pingid
type PingidSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PingidParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider PingidInitParameters `json:"initProvider,omitempty"`
}

// PingidStatus defines the observed state of Pingid.
type PingidStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PingidObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Pingid is the Schema for the Pingids API. Managing the MFA PingID method configuration
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type Pingid struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.mountAccessor) || (has(self.initProvider) && has(self.initProvider.mountAccessor))",message="spec.forProvider.mountAccessor is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.settingsFileBase64) || (has(self.initProvider) && has(self.initProvider.settingsFileBase64))",message="spec.forProvider.settingsFileBase64 is a required parameter"
	Spec   PingidSpec   `json:"spec"`
	Status PingidStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PingidList contains a list of Pingids
type PingidList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Pingid `json:"items"`
}

// Repository type metadata.
var (
	Pingid_Kind             = "Pingid"
	Pingid_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Pingid_Kind}.String()
	Pingid_KindAPIVersion   = Pingid_Kind + "." + CRDGroupVersion.String()
	Pingid_GroupVersionKind = CRDGroupVersion.WithKind(Pingid_Kind)
)

func init() {
	SchemeBuilder.Register(&Pingid{}, &PingidList{})
}

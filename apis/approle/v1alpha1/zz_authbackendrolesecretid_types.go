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

type AuthBackendRoleSecretIDInitParameters struct {

	// Unique name of the auth backend to configure.
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// If set, specifies blocks of IP addresses which can
	// perform the login operation using this SecretID.
	// List of CIDR blocks that can log in using the SecretID.
	// +listType=set
	CidrList []*string `json:"cidrList,omitempty" tf:"cidr_list,omitempty"`

	// A JSON-encoded string containing metadata in
	// key-value pairs to be set on tokens issued with this SecretID.
	// JSON-encoded secret data to write.
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The name of the role to create the SecretID for.
	// Name of the role.
	RoleName *string `json:"roleName,omitempty" tf:"role_name,omitempty"`

	// The SecretID to be created. If set, uses "Push"
	// mode.  Defaults to Vault auto-generating SecretIDs.
	// The SecretID to be managed. If not specified, Vault auto-generates one.
	SecretIDSecretRef *v1.SecretKeySelector `json:"secretIdSecretRef,omitempty" tf:"-"`

	// Set to true to use the wrapped secret-id accessor as the resource ID.
	// If false (default value), a fresh secret ID will be regenerated whenever the wrapping token is expired or
	// invalidated through unwrapping.
	// Use the wrapped secret-id accessor as the id of this resource. If false, a fresh secret-id will be regenerated whenever the wrapping token is expired or invalidated through unwrapping.
	WithWrappedAccessor *bool `json:"withWrappedAccessor,omitempty" tf:"with_wrapped_accessor,omitempty"`

	// If set, the SecretID response will be
	// response-wrapped
	// and available for the duration specified. Only a single unwrapping of the
	// token is allowed.
	// The TTL duration of the wrapped SecretID.
	WrappingTTL *string `json:"wrappingTtl,omitempty" tf:"wrapping_ttl,omitempty"`
}

type AuthBackendRoleSecretIDObservation struct {

	// The unique ID for this SecretID that can be safely logged.
	// The unique ID used to access this SecretID.
	Accessor *string `json:"accessor,omitempty" tf:"accessor,omitempty"`

	// Unique name of the auth backend to configure.
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// If set, specifies blocks of IP addresses which can
	// perform the login operation using this SecretID.
	// List of CIDR blocks that can log in using the SecretID.
	// +listType=set
	CidrList []*string `json:"cidrList,omitempty" tf:"cidr_list,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A JSON-encoded string containing metadata in
	// key-value pairs to be set on tokens issued with this SecretID.
	// JSON-encoded secret data to write.
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The name of the role to create the SecretID for.
	// Name of the role.
	RoleName *string `json:"roleName,omitempty" tf:"role_name,omitempty"`

	// Set to true to use the wrapped secret-id accessor as the resource ID.
	// If false (default value), a fresh secret ID will be regenerated whenever the wrapping token is expired or
	// invalidated through unwrapping.
	// Use the wrapped secret-id accessor as the id of this resource. If false, a fresh secret-id will be regenerated whenever the wrapping token is expired or invalidated through unwrapping.
	WithWrappedAccessor *bool `json:"withWrappedAccessor,omitempty" tf:"with_wrapped_accessor,omitempty"`

	// The unique ID for the response-wrapped SecretID that can
	// be safely logged.
	// The wrapped SecretID accessor.
	WrappingAccessor *string `json:"wrappingAccessor,omitempty" tf:"wrapping_accessor,omitempty"`

	// If set, the SecretID response will be
	// response-wrapped
	// and available for the duration specified. Only a single unwrapping of the
	// token is allowed.
	// The TTL duration of the wrapped SecretID.
	WrappingTTL *string `json:"wrappingTtl,omitempty" tf:"wrapping_ttl,omitempty"`
}

type AuthBackendRoleSecretIDParameters struct {

	// Unique name of the auth backend to configure.
	// +kubebuilder:validation:Optional
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// If set, specifies blocks of IP addresses which can
	// perform the login operation using this SecretID.
	// List of CIDR blocks that can log in using the SecretID.
	// +kubebuilder:validation:Optional
	// +listType=set
	CidrList []*string `json:"cidrList,omitempty" tf:"cidr_list,omitempty"`

	// A JSON-encoded string containing metadata in
	// key-value pairs to be set on tokens issued with this SecretID.
	// JSON-encoded secret data to write.
	// +kubebuilder:validation:Optional
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The name of the role to create the SecretID for.
	// Name of the role.
	// +kubebuilder:validation:Optional
	RoleName *string `json:"roleName,omitempty" tf:"role_name,omitempty"`

	// The SecretID to be created. If set, uses "Push"
	// mode.  Defaults to Vault auto-generating SecretIDs.
	// The SecretID to be managed. If not specified, Vault auto-generates one.
	// +kubebuilder:validation:Optional
	SecretIDSecretRef *v1.SecretKeySelector `json:"secretIdSecretRef,omitempty" tf:"-"`

	// Set to true to use the wrapped secret-id accessor as the resource ID.
	// If false (default value), a fresh secret ID will be regenerated whenever the wrapping token is expired or
	// invalidated through unwrapping.
	// Use the wrapped secret-id accessor as the id of this resource. If false, a fresh secret-id will be regenerated whenever the wrapping token is expired or invalidated through unwrapping.
	// +kubebuilder:validation:Optional
	WithWrappedAccessor *bool `json:"withWrappedAccessor,omitempty" tf:"with_wrapped_accessor,omitempty"`

	// If set, the SecretID response will be
	// response-wrapped
	// and available for the duration specified. Only a single unwrapping of the
	// token is allowed.
	// The TTL duration of the wrapped SecretID.
	// +kubebuilder:validation:Optional
	WrappingTTL *string `json:"wrappingTtl,omitempty" tf:"wrapping_ttl,omitempty"`
}

// AuthBackendRoleSecretIDSpec defines the desired state of AuthBackendRoleSecretID
type AuthBackendRoleSecretIDSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AuthBackendRoleSecretIDParameters `json:"forProvider"`
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
	InitProvider AuthBackendRoleSecretIDInitParameters `json:"initProvider,omitempty"`
}

// AuthBackendRoleSecretIDStatus defines the observed state of AuthBackendRoleSecretID.
type AuthBackendRoleSecretIDStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AuthBackendRoleSecretIDObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AuthBackendRoleSecretID is the Schema for the AuthBackendRoleSecretIDs API. Manages AppRole auth backend role SecretIDs in Vault.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type AuthBackendRoleSecretID struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.roleName) || (has(self.initProvider) && has(self.initProvider.roleName))",message="spec.forProvider.roleName is a required parameter"
	Spec   AuthBackendRoleSecretIDSpec   `json:"spec"`
	Status AuthBackendRoleSecretIDStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AuthBackendRoleSecretIDList contains a list of AuthBackendRoleSecretIDs
type AuthBackendRoleSecretIDList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AuthBackendRoleSecretID `json:"items"`
}

// Repository type metadata.
var (
	AuthBackendRoleSecretID_Kind             = "AuthBackendRoleSecretID"
	AuthBackendRoleSecretID_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AuthBackendRoleSecretID_Kind}.String()
	AuthBackendRoleSecretID_KindAPIVersion   = AuthBackendRoleSecretID_Kind + "." + CRDGroupVersion.String()
	AuthBackendRoleSecretID_GroupVersionKind = CRDGroupVersion.WithKind(AuthBackendRoleSecretID_Kind)
)

func init() {
	SchemeBuilder.Register(&AuthBackendRoleSecretID{}, &AuthBackendRoleSecretIDList{})
}

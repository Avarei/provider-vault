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

type AuthBackendRoletagBlacklistInitParameters struct {

	// The path the AWS auth backend being configured was
	// mounted at.
	// Unique name of the auth backend to configure.
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// If set to true, disables the periodic
	// tidying of the roletag blacklist entries. Defaults to false.
	// If true, disables the periodic tidying of the roletag blacklist entries.
	DisablePeriodicTidy *bool `json:"disablePeriodicTidy,omitempty" tf:"disable_periodic_tidy,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The amount of extra time that must have passed
	// beyond the roletag expiration, before it is removed from the backend storage.
	// Defaults to 259,200 seconds, or 72 hours.
	// The amount of extra time that must have passed beyond the roletag expiration, before it's removed from backend storage.
	SafetyBuffer *float64 `json:"safetyBuffer,omitempty" tf:"safety_buffer,omitempty"`
}

type AuthBackendRoletagBlacklistObservation struct {

	// The path the AWS auth backend being configured was
	// mounted at.
	// Unique name of the auth backend to configure.
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// If set to true, disables the periodic
	// tidying of the roletag blacklist entries. Defaults to false.
	// If true, disables the periodic tidying of the roletag blacklist entries.
	DisablePeriodicTidy *bool `json:"disablePeriodicTidy,omitempty" tf:"disable_periodic_tidy,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The amount of extra time that must have passed
	// beyond the roletag expiration, before it is removed from the backend storage.
	// Defaults to 259,200 seconds, or 72 hours.
	// The amount of extra time that must have passed beyond the roletag expiration, before it's removed from backend storage.
	SafetyBuffer *float64 `json:"safetyBuffer,omitempty" tf:"safety_buffer,omitempty"`
}

type AuthBackendRoletagBlacklistParameters struct {

	// The path the AWS auth backend being configured was
	// mounted at.
	// Unique name of the auth backend to configure.
	// +kubebuilder:validation:Optional
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// If set to true, disables the periodic
	// tidying of the roletag blacklist entries. Defaults to false.
	// If true, disables the periodic tidying of the roletag blacklist entries.
	// +kubebuilder:validation:Optional
	DisablePeriodicTidy *bool `json:"disablePeriodicTidy,omitempty" tf:"disable_periodic_tidy,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The amount of extra time that must have passed
	// beyond the roletag expiration, before it is removed from the backend storage.
	// Defaults to 259,200 seconds, or 72 hours.
	// The amount of extra time that must have passed beyond the roletag expiration, before it's removed from backend storage.
	// +kubebuilder:validation:Optional
	SafetyBuffer *float64 `json:"safetyBuffer,omitempty" tf:"safety_buffer,omitempty"`
}

// AuthBackendRoletagBlacklistSpec defines the desired state of AuthBackendRoletagBlacklist
type AuthBackendRoletagBlacklistSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AuthBackendRoletagBlacklistParameters `json:"forProvider"`
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
	InitProvider AuthBackendRoletagBlacklistInitParameters `json:"initProvider,omitempty"`
}

// AuthBackendRoletagBlacklistStatus defines the observed state of AuthBackendRoletagBlacklist.
type AuthBackendRoletagBlacklistStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AuthBackendRoletagBlacklistObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AuthBackendRoletagBlacklist is the Schema for the AuthBackendRoletagBlacklists API. Configures the periodic tidying operation of the blacklisted role tag entries.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type AuthBackendRoletagBlacklist struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.backend) || (has(self.initProvider) && has(self.initProvider.backend))",message="spec.forProvider.backend is a required parameter"
	Spec   AuthBackendRoletagBlacklistSpec   `json:"spec"`
	Status AuthBackendRoletagBlacklistStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AuthBackendRoletagBlacklistList contains a list of AuthBackendRoletagBlacklists
type AuthBackendRoletagBlacklistList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AuthBackendRoletagBlacklist `json:"items"`
}

// Repository type metadata.
var (
	AuthBackendRoletagBlacklist_Kind             = "AuthBackendRoletagBlacklist"
	AuthBackendRoletagBlacklist_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AuthBackendRoletagBlacklist_Kind}.String()
	AuthBackendRoletagBlacklist_KindAPIVersion   = AuthBackendRoletagBlacklist_Kind + "." + CRDGroupVersion.String()
	AuthBackendRoletagBlacklist_GroupVersionKind = CRDGroupVersion.WithKind(AuthBackendRoletagBlacklist_Kind)
)

func init() {
	SchemeBuilder.Register(&AuthBackendRoletagBlacklist{}, &AuthBackendRoletagBlacklistList{})
}

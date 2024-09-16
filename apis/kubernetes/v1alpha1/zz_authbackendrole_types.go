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

type AuthBackendRoleInitParameters struct {

	// Configures how identity aliases are generated.
	// Valid choices are: serviceaccount_uid, serviceaccount_name. (vault-1.9+)
	// Configures how identity aliases are generated. Valid choices are: serviceaccount_uid, serviceaccount_name
	AliasNameSource *string `json:"aliasNameSource,omitempty" tf:"alias_name_source,omitempty"`

	// Audience claim to verify in the JWT.
	// Optional Audience claim to verify in the JWT.
	Audience *string `json:"audience,omitempty" tf:"audience,omitempty"`

	// Unique name of the kubernetes backend to configure.
	// Unique name of the kubernetes backend to configure.
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// List of service account names able to access this role. If set to ["*"] all names are allowed, both this and bound_service_account_namespaces can not be "*".
	// List of service account names able to access this role. If set to `["*"]` all names are allowed, both this and bound_service_account_namespaces can not be "*".
	// +listType=set
	BoundServiceAccountNames []*string `json:"boundServiceAccountNames,omitempty" tf:"bound_service_account_names,omitempty"`

	// List of namespaces allowed to access this role. If set to ["*"] all namespaces are allowed, both this and bound_service_account_names can not be set to "*".
	// List of namespaces allowed to access this role. If set to `["*"]` all namespaces are allowed, both this and bound_service_account_names can not be set to "*".
	// +listType=set
	BoundServiceAccountNamespaces []*string `json:"boundServiceAccountNamespaces,omitempty" tf:"bound_service_account_namespaces,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Name of the role.
	// Name of the role.
	RoleName *string `json:"roleName,omitempty" tf:"role_name,omitempty"`

	// List of CIDR blocks; if set, specifies blocks of IP
	// addresses which can authenticate successfully, and ties the resulting token to these blocks
	// as well.
	// Specifies the blocks of IP addresses which are allowed to use the generated token
	// +listType=set
	TokenBoundCidrs []*string `json:"tokenBoundCidrs,omitempty" tf:"token_bound_cidrs,omitempty"`

	// If set, will encode an
	// explicit max TTL
	// onto the token in number of seconds. This is a hard cap even if token_ttl and
	// token_max_ttl would otherwise allow a renewal.
	// Generated Token's Explicit Maximum TTL in seconds
	TokenExplicitMaxTTL *float64 `json:"tokenExplicitMaxTtl,omitempty" tf:"token_explicit_max_ttl,omitempty"`

	// The maximum lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	// The maximum lifetime of the generated token
	TokenMaxTTL *float64 `json:"tokenMaxTtl,omitempty" tf:"token_max_ttl,omitempty"`

	// If set, the default policy will not be set on
	// generated tokens; otherwise it will be added to the policies set in token_policies.
	// If true, the 'default' policy will not automatically be added to generated tokens
	TokenNoDefaultPolicy *bool `json:"tokenNoDefaultPolicy,omitempty" tf:"token_no_default_policy,omitempty"`

	// The maximum number
	// of times a generated token may be used (within its lifetime); 0 means unlimited.
	// The maximum number of times a token may be used, a value of zero means unlimited
	TokenNumUses *float64 `json:"tokenNumUses,omitempty" tf:"token_num_uses,omitempty"`

	// If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. Specified in seconds.
	// Generated Token's Period
	TokenPeriod *float64 `json:"tokenPeriod,omitempty" tf:"token_period,omitempty"`

	// List of policies to encode onto generated tokens. Depending
	// on the auth method, this list may be supplemented by user/group/other values.
	// Generated Token's Policies
	// +listType=set
	TokenPolicies []*string `json:"tokenPolicies,omitempty" tf:"token_policies,omitempty"`

	// The initial ttl of the token to generate in seconds
	TokenTTL *float64 `json:"tokenTtl,omitempty" tf:"token_ttl,omitempty"`

	// The type of token that should be generated. Can be service,
	// batch, or default to use the mount's tuned default (which unless changed will be
	// service tokens). For token store roles, there are two additional possibilities:
	// default-service and default-batch which specify the type to return unless the client
	// requests a different type at generation time.
	// The type of token to generate, service or batch
	TokenType *string `json:"tokenType,omitempty" tf:"token_type,omitempty"`
}

type AuthBackendRoleObservation struct {

	// Configures how identity aliases are generated.
	// Valid choices are: serviceaccount_uid, serviceaccount_name. (vault-1.9+)
	// Configures how identity aliases are generated. Valid choices are: serviceaccount_uid, serviceaccount_name
	AliasNameSource *string `json:"aliasNameSource,omitempty" tf:"alias_name_source,omitempty"`

	// Audience claim to verify in the JWT.
	// Optional Audience claim to verify in the JWT.
	Audience *string `json:"audience,omitempty" tf:"audience,omitempty"`

	// Unique name of the kubernetes backend to configure.
	// Unique name of the kubernetes backend to configure.
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// List of service account names able to access this role. If set to ["*"] all names are allowed, both this and bound_service_account_namespaces can not be "*".
	// List of service account names able to access this role. If set to `["*"]` all names are allowed, both this and bound_service_account_namespaces can not be "*".
	// +listType=set
	BoundServiceAccountNames []*string `json:"boundServiceAccountNames,omitempty" tf:"bound_service_account_names,omitempty"`

	// List of namespaces allowed to access this role. If set to ["*"] all namespaces are allowed, both this and bound_service_account_names can not be set to "*".
	// List of namespaces allowed to access this role. If set to `["*"]` all namespaces are allowed, both this and bound_service_account_names can not be set to "*".
	// +listType=set
	BoundServiceAccountNamespaces []*string `json:"boundServiceAccountNamespaces,omitempty" tf:"bound_service_account_namespaces,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Name of the role.
	// Name of the role.
	RoleName *string `json:"roleName,omitempty" tf:"role_name,omitempty"`

	// List of CIDR blocks; if set, specifies blocks of IP
	// addresses which can authenticate successfully, and ties the resulting token to these blocks
	// as well.
	// Specifies the blocks of IP addresses which are allowed to use the generated token
	// +listType=set
	TokenBoundCidrs []*string `json:"tokenBoundCidrs,omitempty" tf:"token_bound_cidrs,omitempty"`

	// If set, will encode an
	// explicit max TTL
	// onto the token in number of seconds. This is a hard cap even if token_ttl and
	// token_max_ttl would otherwise allow a renewal.
	// Generated Token's Explicit Maximum TTL in seconds
	TokenExplicitMaxTTL *float64 `json:"tokenExplicitMaxTtl,omitempty" tf:"token_explicit_max_ttl,omitempty"`

	// The maximum lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	// The maximum lifetime of the generated token
	TokenMaxTTL *float64 `json:"tokenMaxTtl,omitempty" tf:"token_max_ttl,omitempty"`

	// If set, the default policy will not be set on
	// generated tokens; otherwise it will be added to the policies set in token_policies.
	// If true, the 'default' policy will not automatically be added to generated tokens
	TokenNoDefaultPolicy *bool `json:"tokenNoDefaultPolicy,omitempty" tf:"token_no_default_policy,omitempty"`

	// The maximum number
	// of times a generated token may be used (within its lifetime); 0 means unlimited.
	// The maximum number of times a token may be used, a value of zero means unlimited
	TokenNumUses *float64 `json:"tokenNumUses,omitempty" tf:"token_num_uses,omitempty"`

	// If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. Specified in seconds.
	// Generated Token's Period
	TokenPeriod *float64 `json:"tokenPeriod,omitempty" tf:"token_period,omitempty"`

	// List of policies to encode onto generated tokens. Depending
	// on the auth method, this list may be supplemented by user/group/other values.
	// Generated Token's Policies
	// +listType=set
	TokenPolicies []*string `json:"tokenPolicies,omitempty" tf:"token_policies,omitempty"`

	// The initial ttl of the token to generate in seconds
	TokenTTL *float64 `json:"tokenTtl,omitempty" tf:"token_ttl,omitempty"`

	// The type of token that should be generated. Can be service,
	// batch, or default to use the mount's tuned default (which unless changed will be
	// service tokens). For token store roles, there are two additional possibilities:
	// default-service and default-batch which specify the type to return unless the client
	// requests a different type at generation time.
	// The type of token to generate, service or batch
	TokenType *string `json:"tokenType,omitempty" tf:"token_type,omitempty"`
}

type AuthBackendRoleParameters struct {

	// Configures how identity aliases are generated.
	// Valid choices are: serviceaccount_uid, serviceaccount_name. (vault-1.9+)
	// Configures how identity aliases are generated. Valid choices are: serviceaccount_uid, serviceaccount_name
	// +kubebuilder:validation:Optional
	AliasNameSource *string `json:"aliasNameSource,omitempty" tf:"alias_name_source,omitempty"`

	// Audience claim to verify in the JWT.
	// Optional Audience claim to verify in the JWT.
	// +kubebuilder:validation:Optional
	Audience *string `json:"audience,omitempty" tf:"audience,omitempty"`

	// Unique name of the kubernetes backend to configure.
	// Unique name of the kubernetes backend to configure.
	// +kubebuilder:validation:Optional
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// List of service account names able to access this role. If set to ["*"] all names are allowed, both this and bound_service_account_namespaces can not be "*".
	// List of service account names able to access this role. If set to `["*"]` all names are allowed, both this and bound_service_account_namespaces can not be "*".
	// +kubebuilder:validation:Optional
	// +listType=set
	BoundServiceAccountNames []*string `json:"boundServiceAccountNames,omitempty" tf:"bound_service_account_names,omitempty"`

	// List of namespaces allowed to access this role. If set to ["*"] all namespaces are allowed, both this and bound_service_account_names can not be set to "*".
	// List of namespaces allowed to access this role. If set to `["*"]` all namespaces are allowed, both this and bound_service_account_names can not be set to "*".
	// +kubebuilder:validation:Optional
	// +listType=set
	BoundServiceAccountNamespaces []*string `json:"boundServiceAccountNamespaces,omitempty" tf:"bound_service_account_namespaces,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Name of the role.
	// Name of the role.
	// +kubebuilder:validation:Optional
	RoleName *string `json:"roleName,omitempty" tf:"role_name,omitempty"`

	// List of CIDR blocks; if set, specifies blocks of IP
	// addresses which can authenticate successfully, and ties the resulting token to these blocks
	// as well.
	// Specifies the blocks of IP addresses which are allowed to use the generated token
	// +kubebuilder:validation:Optional
	// +listType=set
	TokenBoundCidrs []*string `json:"tokenBoundCidrs,omitempty" tf:"token_bound_cidrs,omitempty"`

	// If set, will encode an
	// explicit max TTL
	// onto the token in number of seconds. This is a hard cap even if token_ttl and
	// token_max_ttl would otherwise allow a renewal.
	// Generated Token's Explicit Maximum TTL in seconds
	// +kubebuilder:validation:Optional
	TokenExplicitMaxTTL *float64 `json:"tokenExplicitMaxTtl,omitempty" tf:"token_explicit_max_ttl,omitempty"`

	// The maximum lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	// The maximum lifetime of the generated token
	// +kubebuilder:validation:Optional
	TokenMaxTTL *float64 `json:"tokenMaxTtl,omitempty" tf:"token_max_ttl,omitempty"`

	// If set, the default policy will not be set on
	// generated tokens; otherwise it will be added to the policies set in token_policies.
	// If true, the 'default' policy will not automatically be added to generated tokens
	// +kubebuilder:validation:Optional
	TokenNoDefaultPolicy *bool `json:"tokenNoDefaultPolicy,omitempty" tf:"token_no_default_policy,omitempty"`

	// The maximum number
	// of times a generated token may be used (within its lifetime); 0 means unlimited.
	// The maximum number of times a token may be used, a value of zero means unlimited
	// +kubebuilder:validation:Optional
	TokenNumUses *float64 `json:"tokenNumUses,omitempty" tf:"token_num_uses,omitempty"`

	// If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. Specified in seconds.
	// Generated Token's Period
	// +kubebuilder:validation:Optional
	TokenPeriod *float64 `json:"tokenPeriod,omitempty" tf:"token_period,omitempty"`

	// List of policies to encode onto generated tokens. Depending
	// on the auth method, this list may be supplemented by user/group/other values.
	// Generated Token's Policies
	// +kubebuilder:validation:Optional
	// +listType=set
	TokenPolicies []*string `json:"tokenPolicies,omitempty" tf:"token_policies,omitempty"`

	// The initial ttl of the token to generate in seconds
	// +kubebuilder:validation:Optional
	TokenTTL *float64 `json:"tokenTtl,omitempty" tf:"token_ttl,omitempty"`

	// The type of token that should be generated. Can be service,
	// batch, or default to use the mount's tuned default (which unless changed will be
	// service tokens). For token store roles, there are two additional possibilities:
	// default-service and default-batch which specify the type to return unless the client
	// requests a different type at generation time.
	// The type of token to generate, service or batch
	// +kubebuilder:validation:Optional
	TokenType *string `json:"tokenType,omitempty" tf:"token_type,omitempty"`
}

// AuthBackendRoleSpec defines the desired state of AuthBackendRole
type AuthBackendRoleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AuthBackendRoleParameters `json:"forProvider"`
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
	InitProvider AuthBackendRoleInitParameters `json:"initProvider,omitempty"`
}

// AuthBackendRoleStatus defines the observed state of AuthBackendRole.
type AuthBackendRoleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AuthBackendRoleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AuthBackendRole is the Schema for the AuthBackendRoles API. Manages Kubernetes auth backend roles in Vault.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type AuthBackendRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.boundServiceAccountNames) || (has(self.initProvider) && has(self.initProvider.boundServiceAccountNames))",message="spec.forProvider.boundServiceAccountNames is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.boundServiceAccountNamespaces) || (has(self.initProvider) && has(self.initProvider.boundServiceAccountNamespaces))",message="spec.forProvider.boundServiceAccountNamespaces is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.roleName) || (has(self.initProvider) && has(self.initProvider.roleName))",message="spec.forProvider.roleName is a required parameter"
	Spec   AuthBackendRoleSpec   `json:"spec"`
	Status AuthBackendRoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AuthBackendRoleList contains a list of AuthBackendRoles
type AuthBackendRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AuthBackendRole `json:"items"`
}

// Repository type metadata.
var (
	AuthBackendRole_Kind             = "AuthBackendRole"
	AuthBackendRole_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AuthBackendRole_Kind}.String()
	AuthBackendRole_KindAPIVersion   = AuthBackendRole_Kind + "." + CRDGroupVersion.String()
	AuthBackendRole_GroupVersionKind = CRDGroupVersion.WithKind(AuthBackendRole_Kind)
)

func init() {
	SchemeBuilder.Register(&AuthBackendRole{}, &AuthBackendRoleList{})
}

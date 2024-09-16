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

type AuthBackendLoginInitParameters struct {

	// The unique name of the AWS auth backend. Defaults to
	// 'aws'.
	// AWS Auth Backend to read the token from.
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// The HTTP method used in the signed IAM
	// request.
	// The HTTP method used in the signed request.
	IAMHTTPRequestMethod *string `json:"iamHttpRequestMethod,omitempty" tf:"iam_http_request_method,omitempty"`

	// The base64-encoded body of the signed
	// request.
	// The Base64-encoded body of the signed request.
	IAMRequestBody *string `json:"iamRequestBody,omitempty" tf:"iam_request_body,omitempty"`

	// The base64-encoded, JSON serialized
	// representation of the GetCallerIdentity HTTP request headers.
	// The Base64-encoded, JSON serialized representation of the sts:GetCallerIdentity HTTP request headers.
	IAMRequestHeaders *string `json:"iamRequestHeaders,omitempty" tf:"iam_request_headers,omitempty"`

	// The base64-encoded HTTP URL used in the signed
	// request.
	// The Base64-encoded HTTP URL used in the signed request.
	IAMRequestURL *string `json:"iamRequestUrl,omitempty" tf:"iam_request_url,omitempty"`

	// The base64-encoded EC2 instance identity document to
	// authenticate with. Can be retrieved from the EC2 metadata server.
	// Base64-encoded EC2 instance identity document to authenticate with.
	Identity *string `json:"identity,omitempty" tf:"identity,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The unique nonce to be used for login requests. Can be
	// set to a user-specified value, or will contain the server-generated value
	// once a token is issued. EC2 instances can only acquire a single token until
	// the whitelist is tidied again unless they keep track of this nonce.
	// The nonce to be used for subsequent login requests.
	Nonce *string `json:"nonce,omitempty" tf:"nonce,omitempty"`

	// The PKCS#7 signature of the identity document to
	// authenticate with, with all newline characters removed. Can be retrieved from
	// the EC2 metadata server.
	// PKCS7 signature of the identity document to authenticate with, with all newline characters removed.
	Pkcs7 *string `json:"pkcs7,omitempty" tf:"pkcs7,omitempty"`

	// The name of the AWS auth backend role to create tokens
	// against.
	// AWS Auth Role to read the token from.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// The base64-encoded SHA256 RSA signature of the
	// instance identity document to authenticate with, with all newline characters
	// removed. Can be retrieved from the EC2 metadata server.
	// Base64-encoded SHA256 RSA signature of the instance identtiy document to authenticate with.
	Signature *string `json:"signature,omitempty" tf:"signature,omitempty"`
}

type AuthBackendLoginObservation struct {

	// The token's accessor.
	// The accessor returned from Vault for this token.
	Accessor *string `json:"accessor,omitempty" tf:"accessor,omitempty"`

	// The authentication type used to generate this token.
	// The auth method used to generate this token.
	AuthType *string `json:"authType,omitempty" tf:"auth_type,omitempty"`

	// The unique name of the AWS auth backend. Defaults to
	// 'aws'.
	// AWS Auth Backend to read the token from.
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// The HTTP method used in the signed IAM
	// request.
	// The HTTP method used in the signed request.
	IAMHTTPRequestMethod *string `json:"iamHttpRequestMethod,omitempty" tf:"iam_http_request_method,omitempty"`

	// The base64-encoded body of the signed
	// request.
	// The Base64-encoded body of the signed request.
	IAMRequestBody *string `json:"iamRequestBody,omitempty" tf:"iam_request_body,omitempty"`

	// The base64-encoded, JSON serialized
	// representation of the GetCallerIdentity HTTP request headers.
	// The Base64-encoded, JSON serialized representation of the sts:GetCallerIdentity HTTP request headers.
	IAMRequestHeaders *string `json:"iamRequestHeaders,omitempty" tf:"iam_request_headers,omitempty"`

	// The base64-encoded HTTP URL used in the signed
	// request.
	// The Base64-encoded HTTP URL used in the signed request.
	IAMRequestURL *string `json:"iamRequestUrl,omitempty" tf:"iam_request_url,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The base64-encoded EC2 instance identity document to
	// authenticate with. Can be retrieved from the EC2 metadata server.
	// Base64-encoded EC2 instance identity document to authenticate with.
	Identity *string `json:"identity,omitempty" tf:"identity,omitempty"`

	// The duration in seconds the token will be valid, relative
	// to the time in lease_start_time.
	// Lease duration in seconds relative to the time in lease_start_time.
	LeaseDuration *float64 `json:"leaseDuration,omitempty" tf:"lease_duration,omitempty"`

	// the approximate time at which the token was created,
	// using the clock of the system where provider was running.
	// time at which the lease was read, using the clock of the system where provider was running
	LeaseStartTime *string `json:"leaseStartTime,omitempty" tf:"lease_start_time,omitempty"`

	// A map of information returned by the Vault server about the
	// authentication used to generate this token.
	// The metadata reported by the Vault server.
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The unique nonce to be used for login requests. Can be
	// set to a user-specified value, or will contain the server-generated value
	// once a token is issued. EC2 instances can only acquire a single token until
	// the whitelist is tidied again unless they keep track of this nonce.
	// The nonce to be used for subsequent login requests.
	Nonce *string `json:"nonce,omitempty" tf:"nonce,omitempty"`

	// The PKCS#7 signature of the identity document to
	// authenticate with, with all newline characters removed. Can be retrieved from
	// the EC2 metadata server.
	// PKCS7 signature of the identity document to authenticate with, with all newline characters removed.
	Pkcs7 *string `json:"pkcs7,omitempty" tf:"pkcs7,omitempty"`

	// The Vault policies assigned to this token.
	// The policies assigned to this token.
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`

	// Set to true if the token can be extended through renewal.
	// True if the duration of this lease can be extended through renewal.
	Renewable *bool `json:"renewable,omitempty" tf:"renewable,omitempty"`

	// The name of the AWS auth backend role to create tokens
	// against.
	// AWS Auth Role to read the token from.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// The base64-encoded SHA256 RSA signature of the
	// instance identity document to authenticate with, with all newline characters
	// removed. Can be retrieved from the EC2 metadata server.
	// Base64-encoded SHA256 RSA signature of the instance identtiy document to authenticate with.
	Signature *string `json:"signature,omitempty" tf:"signature,omitempty"`
}

type AuthBackendLoginParameters struct {

	// The unique name of the AWS auth backend. Defaults to
	// 'aws'.
	// AWS Auth Backend to read the token from.
	// +kubebuilder:validation:Optional
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// The HTTP method used in the signed IAM
	// request.
	// The HTTP method used in the signed request.
	// +kubebuilder:validation:Optional
	IAMHTTPRequestMethod *string `json:"iamHttpRequestMethod,omitempty" tf:"iam_http_request_method,omitempty"`

	// The base64-encoded body of the signed
	// request.
	// The Base64-encoded body of the signed request.
	// +kubebuilder:validation:Optional
	IAMRequestBody *string `json:"iamRequestBody,omitempty" tf:"iam_request_body,omitempty"`

	// The base64-encoded, JSON serialized
	// representation of the GetCallerIdentity HTTP request headers.
	// The Base64-encoded, JSON serialized representation of the sts:GetCallerIdentity HTTP request headers.
	// +kubebuilder:validation:Optional
	IAMRequestHeaders *string `json:"iamRequestHeaders,omitempty" tf:"iam_request_headers,omitempty"`

	// The base64-encoded HTTP URL used in the signed
	// request.
	// The Base64-encoded HTTP URL used in the signed request.
	// +kubebuilder:validation:Optional
	IAMRequestURL *string `json:"iamRequestUrl,omitempty" tf:"iam_request_url,omitempty"`

	// The base64-encoded EC2 instance identity document to
	// authenticate with. Can be retrieved from the EC2 metadata server.
	// Base64-encoded EC2 instance identity document to authenticate with.
	// +kubebuilder:validation:Optional
	Identity *string `json:"identity,omitempty" tf:"identity,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The unique nonce to be used for login requests. Can be
	// set to a user-specified value, or will contain the server-generated value
	// once a token is issued. EC2 instances can only acquire a single token until
	// the whitelist is tidied again unless they keep track of this nonce.
	// The nonce to be used for subsequent login requests.
	// +kubebuilder:validation:Optional
	Nonce *string `json:"nonce,omitempty" tf:"nonce,omitempty"`

	// The PKCS#7 signature of the identity document to
	// authenticate with, with all newline characters removed. Can be retrieved from
	// the EC2 metadata server.
	// PKCS7 signature of the identity document to authenticate with, with all newline characters removed.
	// +kubebuilder:validation:Optional
	Pkcs7 *string `json:"pkcs7,omitempty" tf:"pkcs7,omitempty"`

	// The name of the AWS auth backend role to create tokens
	// against.
	// AWS Auth Role to read the token from.
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// The base64-encoded SHA256 RSA signature of the
	// instance identity document to authenticate with, with all newline characters
	// removed. Can be retrieved from the EC2 metadata server.
	// Base64-encoded SHA256 RSA signature of the instance identtiy document to authenticate with.
	// +kubebuilder:validation:Optional
	Signature *string `json:"signature,omitempty" tf:"signature,omitempty"`
}

// AuthBackendLoginSpec defines the desired state of AuthBackendLogin
type AuthBackendLoginSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AuthBackendLoginParameters `json:"forProvider"`
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
	InitProvider AuthBackendLoginInitParameters `json:"initProvider,omitempty"`
}

// AuthBackendLoginStatus defines the observed state of AuthBackendLogin.
type AuthBackendLoginStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AuthBackendLoginObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AuthBackendLogin is the Schema for the AuthBackendLogins API. Manages Vault tokens acquired using the AWS auth backend.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type AuthBackendLogin struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AuthBackendLoginSpec   `json:"spec"`
	Status            AuthBackendLoginStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AuthBackendLoginList contains a list of AuthBackendLogins
type AuthBackendLoginList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AuthBackendLogin `json:"items"`
}

// Repository type metadata.
var (
	AuthBackendLogin_Kind             = "AuthBackendLogin"
	AuthBackendLogin_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AuthBackendLogin_Kind}.String()
	AuthBackendLogin_KindAPIVersion   = AuthBackendLogin_Kind + "." + CRDGroupVersion.String()
	AuthBackendLogin_GroupVersionKind = CRDGroupVersion.WithKind(AuthBackendLogin_Kind)
)

func init() {
	SchemeBuilder.Register(&AuthBackendLogin{}, &AuthBackendLoginList{})
}

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

type AlphabetInitParameters struct {

	// A string of characters that contains the alphabet set.
	// A string of characters that contains the alphabet set.
	Alphabet *string `json:"alphabet,omitempty" tf:"alphabet,omitempty"`

	// The name of the alphabet.
	// The name of the alphabet.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Path to where the back-end is mounted within Vault.
	// The mount path for a back-end, for example, the path given in "$ vault auth enable -path=my-aws aws".
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type AlphabetObservation struct {

	// A string of characters that contains the alphabet set.
	// A string of characters that contains the alphabet set.
	Alphabet *string `json:"alphabet,omitempty" tf:"alphabet,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the alphabet.
	// The name of the alphabet.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Path to where the back-end is mounted within Vault.
	// The mount path for a back-end, for example, the path given in "$ vault auth enable -path=my-aws aws".
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type AlphabetParameters struct {

	// A string of characters that contains the alphabet set.
	// A string of characters that contains the alphabet set.
	// +kubebuilder:validation:Optional
	Alphabet *string `json:"alphabet,omitempty" tf:"alphabet,omitempty"`

	// The name of the alphabet.
	// The name of the alphabet.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Path to where the back-end is mounted within Vault.
	// The mount path for a back-end, for example, the path given in "$ vault auth enable -path=my-aws aws".
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

// AlphabetSpec defines the desired state of Alphabet
type AlphabetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AlphabetParameters `json:"forProvider"`
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
	InitProvider AlphabetInitParameters `json:"initProvider,omitempty"`
}

// AlphabetStatus defines the observed state of Alphabet.
type AlphabetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AlphabetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Alphabet is the Schema for the Alphabets API. "/transform/alphabet/{name}"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type Alphabet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.path) || (has(self.initProvider) && has(self.initProvider.path))",message="spec.forProvider.path is a required parameter"
	Spec   AlphabetSpec   `json:"spec"`
	Status AlphabetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AlphabetList contains a list of Alphabets
type AlphabetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Alphabet `json:"items"`
}

// Repository type metadata.
var (
	Alphabet_Kind             = "Alphabet"
	Alphabet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Alphabet_Kind}.String()
	Alphabet_KindAPIVersion   = Alphabet_Kind + "." + CRDGroupVersion.String()
	Alphabet_GroupVersionKind = CRDGroupVersion.WithKind(Alphabet_Kind)
)

func init() {
	SchemeBuilder.Register(&Alphabet{}, &AlphabetList{})
}

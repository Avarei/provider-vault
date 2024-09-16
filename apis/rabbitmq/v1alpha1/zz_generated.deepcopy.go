//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBackend) DeepCopyInto(out *SecretBackend) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBackend.
func (in *SecretBackend) DeepCopy() *SecretBackend {
	if in == nil {
		return nil
	}
	out := new(SecretBackend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretBackend) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBackendInitParameters) DeepCopyInto(out *SecretBackendInitParameters) {
	*out = *in
	if in.ConnectionURI != nil {
		in, out := &in.ConnectionURI, &out.ConnectionURI
		*out = new(string)
		**out = **in
	}
	if in.DefaultLeaseTTLSeconds != nil {
		in, out := &in.DefaultLeaseTTLSeconds, &out.DefaultLeaseTTLSeconds
		*out = new(float64)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisableRemount != nil {
		in, out := &in.DisableRemount, &out.DisableRemount
		*out = new(bool)
		**out = **in
	}
	if in.MaxLeaseTTLSeconds != nil {
		in, out := &in.MaxLeaseTTLSeconds, &out.MaxLeaseTTLSeconds
		*out = new(float64)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.PasswordPolicy != nil {
		in, out := &in.PasswordPolicy, &out.PasswordPolicy
		*out = new(string)
		**out = **in
	}
	out.PasswordSecretRef = in.PasswordSecretRef
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	out.UsernameSecretRef = in.UsernameSecretRef
	if in.UsernameTemplate != nil {
		in, out := &in.UsernameTemplate, &out.UsernameTemplate
		*out = new(string)
		**out = **in
	}
	if in.VerifyConnection != nil {
		in, out := &in.VerifyConnection, &out.VerifyConnection
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBackendInitParameters.
func (in *SecretBackendInitParameters) DeepCopy() *SecretBackendInitParameters {
	if in == nil {
		return nil
	}
	out := new(SecretBackendInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBackendList) DeepCopyInto(out *SecretBackendList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SecretBackend, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBackendList.
func (in *SecretBackendList) DeepCopy() *SecretBackendList {
	if in == nil {
		return nil
	}
	out := new(SecretBackendList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretBackendList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBackendObservation) DeepCopyInto(out *SecretBackendObservation) {
	*out = *in
	if in.ConnectionURI != nil {
		in, out := &in.ConnectionURI, &out.ConnectionURI
		*out = new(string)
		**out = **in
	}
	if in.DefaultLeaseTTLSeconds != nil {
		in, out := &in.DefaultLeaseTTLSeconds, &out.DefaultLeaseTTLSeconds
		*out = new(float64)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisableRemount != nil {
		in, out := &in.DisableRemount, &out.DisableRemount
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.MaxLeaseTTLSeconds != nil {
		in, out := &in.MaxLeaseTTLSeconds, &out.MaxLeaseTTLSeconds
		*out = new(float64)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.PasswordPolicy != nil {
		in, out := &in.PasswordPolicy, &out.PasswordPolicy
		*out = new(string)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	if in.UsernameTemplate != nil {
		in, out := &in.UsernameTemplate, &out.UsernameTemplate
		*out = new(string)
		**out = **in
	}
	if in.VerifyConnection != nil {
		in, out := &in.VerifyConnection, &out.VerifyConnection
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBackendObservation.
func (in *SecretBackendObservation) DeepCopy() *SecretBackendObservation {
	if in == nil {
		return nil
	}
	out := new(SecretBackendObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBackendParameters) DeepCopyInto(out *SecretBackendParameters) {
	*out = *in
	if in.ConnectionURI != nil {
		in, out := &in.ConnectionURI, &out.ConnectionURI
		*out = new(string)
		**out = **in
	}
	if in.DefaultLeaseTTLSeconds != nil {
		in, out := &in.DefaultLeaseTTLSeconds, &out.DefaultLeaseTTLSeconds
		*out = new(float64)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisableRemount != nil {
		in, out := &in.DisableRemount, &out.DisableRemount
		*out = new(bool)
		**out = **in
	}
	if in.MaxLeaseTTLSeconds != nil {
		in, out := &in.MaxLeaseTTLSeconds, &out.MaxLeaseTTLSeconds
		*out = new(float64)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.PasswordPolicy != nil {
		in, out := &in.PasswordPolicy, &out.PasswordPolicy
		*out = new(string)
		**out = **in
	}
	out.PasswordSecretRef = in.PasswordSecretRef
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	out.UsernameSecretRef = in.UsernameSecretRef
	if in.UsernameTemplate != nil {
		in, out := &in.UsernameTemplate, &out.UsernameTemplate
		*out = new(string)
		**out = **in
	}
	if in.VerifyConnection != nil {
		in, out := &in.VerifyConnection, &out.VerifyConnection
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBackendParameters.
func (in *SecretBackendParameters) DeepCopy() *SecretBackendParameters {
	if in == nil {
		return nil
	}
	out := new(SecretBackendParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBackendRole) DeepCopyInto(out *SecretBackendRole) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBackendRole.
func (in *SecretBackendRole) DeepCopy() *SecretBackendRole {
	if in == nil {
		return nil
	}
	out := new(SecretBackendRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretBackendRole) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBackendRoleInitParameters) DeepCopyInto(out *SecretBackendRoleInitParameters) {
	*out = *in
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = new(string)
		**out = **in
	}
	if in.Vhost != nil {
		in, out := &in.Vhost, &out.Vhost
		*out = make([]VhostInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VhostTopic != nil {
		in, out := &in.VhostTopic, &out.VhostTopic
		*out = make([]VhostTopicInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBackendRoleInitParameters.
func (in *SecretBackendRoleInitParameters) DeepCopy() *SecretBackendRoleInitParameters {
	if in == nil {
		return nil
	}
	out := new(SecretBackendRoleInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBackendRoleList) DeepCopyInto(out *SecretBackendRoleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SecretBackendRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBackendRoleList.
func (in *SecretBackendRoleList) DeepCopy() *SecretBackendRoleList {
	if in == nil {
		return nil
	}
	out := new(SecretBackendRoleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretBackendRoleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBackendRoleObservation) DeepCopyInto(out *SecretBackendRoleObservation) {
	*out = *in
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = new(string)
		**out = **in
	}
	if in.Vhost != nil {
		in, out := &in.Vhost, &out.Vhost
		*out = make([]VhostObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VhostTopic != nil {
		in, out := &in.VhostTopic, &out.VhostTopic
		*out = make([]VhostTopicObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBackendRoleObservation.
func (in *SecretBackendRoleObservation) DeepCopy() *SecretBackendRoleObservation {
	if in == nil {
		return nil
	}
	out := new(SecretBackendRoleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBackendRoleParameters) DeepCopyInto(out *SecretBackendRoleParameters) {
	*out = *in
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = new(string)
		**out = **in
	}
	if in.Vhost != nil {
		in, out := &in.Vhost, &out.Vhost
		*out = make([]VhostParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VhostTopic != nil {
		in, out := &in.VhostTopic, &out.VhostTopic
		*out = make([]VhostTopicParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBackendRoleParameters.
func (in *SecretBackendRoleParameters) DeepCopy() *SecretBackendRoleParameters {
	if in == nil {
		return nil
	}
	out := new(SecretBackendRoleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBackendRoleSpec) DeepCopyInto(out *SecretBackendRoleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBackendRoleSpec.
func (in *SecretBackendRoleSpec) DeepCopy() *SecretBackendRoleSpec {
	if in == nil {
		return nil
	}
	out := new(SecretBackendRoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBackendRoleStatus) DeepCopyInto(out *SecretBackendRoleStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBackendRoleStatus.
func (in *SecretBackendRoleStatus) DeepCopy() *SecretBackendRoleStatus {
	if in == nil {
		return nil
	}
	out := new(SecretBackendRoleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBackendSpec) DeepCopyInto(out *SecretBackendSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBackendSpec.
func (in *SecretBackendSpec) DeepCopy() *SecretBackendSpec {
	if in == nil {
		return nil
	}
	out := new(SecretBackendSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBackendStatus) DeepCopyInto(out *SecretBackendStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBackendStatus.
func (in *SecretBackendStatus) DeepCopy() *SecretBackendStatus {
	if in == nil {
		return nil
	}
	out := new(SecretBackendStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VhostInitParameters) DeepCopyInto(out *VhostInitParameters) {
	*out = *in
	if in.Configure != nil {
		in, out := &in.Configure, &out.Configure
		*out = new(string)
		**out = **in
	}
	if in.Host != nil {
		in, out := &in.Host, &out.Host
		*out = new(string)
		**out = **in
	}
	if in.Read != nil {
		in, out := &in.Read, &out.Read
		*out = new(string)
		**out = **in
	}
	if in.Write != nil {
		in, out := &in.Write, &out.Write
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VhostInitParameters.
func (in *VhostInitParameters) DeepCopy() *VhostInitParameters {
	if in == nil {
		return nil
	}
	out := new(VhostInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VhostObservation) DeepCopyInto(out *VhostObservation) {
	*out = *in
	if in.Configure != nil {
		in, out := &in.Configure, &out.Configure
		*out = new(string)
		**out = **in
	}
	if in.Host != nil {
		in, out := &in.Host, &out.Host
		*out = new(string)
		**out = **in
	}
	if in.Read != nil {
		in, out := &in.Read, &out.Read
		*out = new(string)
		**out = **in
	}
	if in.Write != nil {
		in, out := &in.Write, &out.Write
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VhostObservation.
func (in *VhostObservation) DeepCopy() *VhostObservation {
	if in == nil {
		return nil
	}
	out := new(VhostObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VhostParameters) DeepCopyInto(out *VhostParameters) {
	*out = *in
	if in.Configure != nil {
		in, out := &in.Configure, &out.Configure
		*out = new(string)
		**out = **in
	}
	if in.Host != nil {
		in, out := &in.Host, &out.Host
		*out = new(string)
		**out = **in
	}
	if in.Read != nil {
		in, out := &in.Read, &out.Read
		*out = new(string)
		**out = **in
	}
	if in.Write != nil {
		in, out := &in.Write, &out.Write
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VhostParameters.
func (in *VhostParameters) DeepCopy() *VhostParameters {
	if in == nil {
		return nil
	}
	out := new(VhostParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VhostTopicInitParameters) DeepCopyInto(out *VhostTopicInitParameters) {
	*out = *in
	if in.Host != nil {
		in, out := &in.Host, &out.Host
		*out = new(string)
		**out = **in
	}
	if in.Vhost != nil {
		in, out := &in.Vhost, &out.Vhost
		*out = make([]VhostTopicVhostInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VhostTopicInitParameters.
func (in *VhostTopicInitParameters) DeepCopy() *VhostTopicInitParameters {
	if in == nil {
		return nil
	}
	out := new(VhostTopicInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VhostTopicObservation) DeepCopyInto(out *VhostTopicObservation) {
	*out = *in
	if in.Host != nil {
		in, out := &in.Host, &out.Host
		*out = new(string)
		**out = **in
	}
	if in.Vhost != nil {
		in, out := &in.Vhost, &out.Vhost
		*out = make([]VhostTopicVhostObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VhostTopicObservation.
func (in *VhostTopicObservation) DeepCopy() *VhostTopicObservation {
	if in == nil {
		return nil
	}
	out := new(VhostTopicObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VhostTopicParameters) DeepCopyInto(out *VhostTopicParameters) {
	*out = *in
	if in.Host != nil {
		in, out := &in.Host, &out.Host
		*out = new(string)
		**out = **in
	}
	if in.Vhost != nil {
		in, out := &in.Vhost, &out.Vhost
		*out = make([]VhostTopicVhostParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VhostTopicParameters.
func (in *VhostTopicParameters) DeepCopy() *VhostTopicParameters {
	if in == nil {
		return nil
	}
	out := new(VhostTopicParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VhostTopicVhostInitParameters) DeepCopyInto(out *VhostTopicVhostInitParameters) {
	*out = *in
	if in.Read != nil {
		in, out := &in.Read, &out.Read
		*out = new(string)
		**out = **in
	}
	if in.Topic != nil {
		in, out := &in.Topic, &out.Topic
		*out = new(string)
		**out = **in
	}
	if in.Write != nil {
		in, out := &in.Write, &out.Write
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VhostTopicVhostInitParameters.
func (in *VhostTopicVhostInitParameters) DeepCopy() *VhostTopicVhostInitParameters {
	if in == nil {
		return nil
	}
	out := new(VhostTopicVhostInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VhostTopicVhostObservation) DeepCopyInto(out *VhostTopicVhostObservation) {
	*out = *in
	if in.Read != nil {
		in, out := &in.Read, &out.Read
		*out = new(string)
		**out = **in
	}
	if in.Topic != nil {
		in, out := &in.Topic, &out.Topic
		*out = new(string)
		**out = **in
	}
	if in.Write != nil {
		in, out := &in.Write, &out.Write
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VhostTopicVhostObservation.
func (in *VhostTopicVhostObservation) DeepCopy() *VhostTopicVhostObservation {
	if in == nil {
		return nil
	}
	out := new(VhostTopicVhostObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VhostTopicVhostParameters) DeepCopyInto(out *VhostTopicVhostParameters) {
	*out = *in
	if in.Read != nil {
		in, out := &in.Read, &out.Read
		*out = new(string)
		**out = **in
	}
	if in.Topic != nil {
		in, out := &in.Topic, &out.Topic
		*out = new(string)
		**out = **in
	}
	if in.Write != nil {
		in, out := &in.Write, &out.Write
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VhostTopicVhostParameters.
func (in *VhostTopicVhostParameters) DeepCopy() *VhostTopicVhostParameters {
	if in == nil {
		return nil
	}
	out := new(VhostTopicVhostParameters)
	in.DeepCopyInto(out)
	return out
}

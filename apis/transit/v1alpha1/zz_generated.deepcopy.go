//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBackendKey) DeepCopyInto(out *SecretBackendKey) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBackendKey.
func (in *SecretBackendKey) DeepCopy() *SecretBackendKey {
	if in == nil {
		return nil
	}
	out := new(SecretBackendKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretBackendKey) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBackendKeyInitParameters) DeepCopyInto(out *SecretBackendKeyInitParameters) {
	*out = *in
	if in.AllowPlaintextBackup != nil {
		in, out := &in.AllowPlaintextBackup, &out.AllowPlaintextBackup
		*out = new(bool)
		**out = **in
	}
	if in.AutoRotateInterval != nil {
		in, out := &in.AutoRotateInterval, &out.AutoRotateInterval
		*out = new(float64)
		**out = **in
	}
	if in.AutoRotatePeriod != nil {
		in, out := &in.AutoRotatePeriod, &out.AutoRotatePeriod
		*out = new(float64)
		**out = **in
	}
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = new(string)
		**out = **in
	}
	if in.ConvergentEncryption != nil {
		in, out := &in.ConvergentEncryption, &out.ConvergentEncryption
		*out = new(bool)
		**out = **in
	}
	if in.DeletionAllowed != nil {
		in, out := &in.DeletionAllowed, &out.DeletionAllowed
		*out = new(bool)
		**out = **in
	}
	if in.Derived != nil {
		in, out := &in.Derived, &out.Derived
		*out = new(bool)
		**out = **in
	}
	if in.Exportable != nil {
		in, out := &in.Exportable, &out.Exportable
		*out = new(bool)
		**out = **in
	}
	if in.KeySize != nil {
		in, out := &in.KeySize, &out.KeySize
		*out = new(float64)
		**out = **in
	}
	if in.MinDecryptionVersion != nil {
		in, out := &in.MinDecryptionVersion, &out.MinDecryptionVersion
		*out = new(float64)
		**out = **in
	}
	if in.MinEncryptionVersion != nil {
		in, out := &in.MinEncryptionVersion, &out.MinEncryptionVersion
		*out = new(float64)
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
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBackendKeyInitParameters.
func (in *SecretBackendKeyInitParameters) DeepCopy() *SecretBackendKeyInitParameters {
	if in == nil {
		return nil
	}
	out := new(SecretBackendKeyInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBackendKeyList) DeepCopyInto(out *SecretBackendKeyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SecretBackendKey, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBackendKeyList.
func (in *SecretBackendKeyList) DeepCopy() *SecretBackendKeyList {
	if in == nil {
		return nil
	}
	out := new(SecretBackendKeyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretBackendKeyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBackendKeyObservation) DeepCopyInto(out *SecretBackendKeyObservation) {
	*out = *in
	if in.AllowPlaintextBackup != nil {
		in, out := &in.AllowPlaintextBackup, &out.AllowPlaintextBackup
		*out = new(bool)
		**out = **in
	}
	if in.AutoRotateInterval != nil {
		in, out := &in.AutoRotateInterval, &out.AutoRotateInterval
		*out = new(float64)
		**out = **in
	}
	if in.AutoRotatePeriod != nil {
		in, out := &in.AutoRotatePeriod, &out.AutoRotatePeriod
		*out = new(float64)
		**out = **in
	}
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = new(string)
		**out = **in
	}
	if in.ConvergentEncryption != nil {
		in, out := &in.ConvergentEncryption, &out.ConvergentEncryption
		*out = new(bool)
		**out = **in
	}
	if in.DeletionAllowed != nil {
		in, out := &in.DeletionAllowed, &out.DeletionAllowed
		*out = new(bool)
		**out = **in
	}
	if in.Derived != nil {
		in, out := &in.Derived, &out.Derived
		*out = new(bool)
		**out = **in
	}
	if in.Exportable != nil {
		in, out := &in.Exportable, &out.Exportable
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.KeySize != nil {
		in, out := &in.KeySize, &out.KeySize
		*out = new(float64)
		**out = **in
	}
	if in.Keys != nil {
		in, out := &in.Keys, &out.Keys
		*out = make([]map[string]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make(map[string]*string, len(*in))
				for key, val := range *in {
					var outVal *string
					if val == nil {
						(*out)[key] = nil
					} else {
						in, out := &val, &outVal
						*out = new(string)
						**out = **in
					}
					(*out)[key] = outVal
				}
			}
		}
	}
	if in.LatestVersion != nil {
		in, out := &in.LatestVersion, &out.LatestVersion
		*out = new(float64)
		**out = **in
	}
	if in.MinAvailableVersion != nil {
		in, out := &in.MinAvailableVersion, &out.MinAvailableVersion
		*out = new(float64)
		**out = **in
	}
	if in.MinDecryptionVersion != nil {
		in, out := &in.MinDecryptionVersion, &out.MinDecryptionVersion
		*out = new(float64)
		**out = **in
	}
	if in.MinEncryptionVersion != nil {
		in, out := &in.MinEncryptionVersion, &out.MinEncryptionVersion
		*out = new(float64)
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
	if in.SupportsDecryption != nil {
		in, out := &in.SupportsDecryption, &out.SupportsDecryption
		*out = new(bool)
		**out = **in
	}
	if in.SupportsDerivation != nil {
		in, out := &in.SupportsDerivation, &out.SupportsDerivation
		*out = new(bool)
		**out = **in
	}
	if in.SupportsEncryption != nil {
		in, out := &in.SupportsEncryption, &out.SupportsEncryption
		*out = new(bool)
		**out = **in
	}
	if in.SupportsSigning != nil {
		in, out := &in.SupportsSigning, &out.SupportsSigning
		*out = new(bool)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBackendKeyObservation.
func (in *SecretBackendKeyObservation) DeepCopy() *SecretBackendKeyObservation {
	if in == nil {
		return nil
	}
	out := new(SecretBackendKeyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBackendKeyParameters) DeepCopyInto(out *SecretBackendKeyParameters) {
	*out = *in
	if in.AllowPlaintextBackup != nil {
		in, out := &in.AllowPlaintextBackup, &out.AllowPlaintextBackup
		*out = new(bool)
		**out = **in
	}
	if in.AutoRotateInterval != nil {
		in, out := &in.AutoRotateInterval, &out.AutoRotateInterval
		*out = new(float64)
		**out = **in
	}
	if in.AutoRotatePeriod != nil {
		in, out := &in.AutoRotatePeriod, &out.AutoRotatePeriod
		*out = new(float64)
		**out = **in
	}
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = new(string)
		**out = **in
	}
	if in.ConvergentEncryption != nil {
		in, out := &in.ConvergentEncryption, &out.ConvergentEncryption
		*out = new(bool)
		**out = **in
	}
	if in.DeletionAllowed != nil {
		in, out := &in.DeletionAllowed, &out.DeletionAllowed
		*out = new(bool)
		**out = **in
	}
	if in.Derived != nil {
		in, out := &in.Derived, &out.Derived
		*out = new(bool)
		**out = **in
	}
	if in.Exportable != nil {
		in, out := &in.Exportable, &out.Exportable
		*out = new(bool)
		**out = **in
	}
	if in.KeySize != nil {
		in, out := &in.KeySize, &out.KeySize
		*out = new(float64)
		**out = **in
	}
	if in.MinDecryptionVersion != nil {
		in, out := &in.MinDecryptionVersion, &out.MinDecryptionVersion
		*out = new(float64)
		**out = **in
	}
	if in.MinEncryptionVersion != nil {
		in, out := &in.MinEncryptionVersion, &out.MinEncryptionVersion
		*out = new(float64)
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
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBackendKeyParameters.
func (in *SecretBackendKeyParameters) DeepCopy() *SecretBackendKeyParameters {
	if in == nil {
		return nil
	}
	out := new(SecretBackendKeyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBackendKeySpec) DeepCopyInto(out *SecretBackendKeySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBackendKeySpec.
func (in *SecretBackendKeySpec) DeepCopy() *SecretBackendKeySpec {
	if in == nil {
		return nil
	}
	out := new(SecretBackendKeySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBackendKeyStatus) DeepCopyInto(out *SecretBackendKeyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBackendKeyStatus.
func (in *SecretBackendKeyStatus) DeepCopy() *SecretBackendKeyStatus {
	if in == nil {
		return nil
	}
	out := new(SecretBackendKeyStatus)
	in.DeepCopyInto(out)
	return out
}

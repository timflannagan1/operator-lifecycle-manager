// +build !ignore_autogenerated

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha1

import (
	json "encoding/json"
	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

// GetGeneratedDeepCopyFuncs returns the generated funcs, since we aren't registering them.
//
// Deprecated: deepcopy registration will go away when static deepcopy is fully implemented.
func GetGeneratedDeepCopyFuncs() []conversion.GeneratedDeepCopyFunc {
	return []conversion.GeneratedDeepCopyFunc{
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*NamedInstallStrategy).DeepCopyInto(out.(*NamedInstallStrategy))
			return nil
		}, InType: reflect.TypeOf(&NamedInstallStrategy{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*OperatorVersion).DeepCopyInto(out.(*OperatorVersion))
			return nil
		}, InType: reflect.TypeOf(&OperatorVersion{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*OperatorVersionList).DeepCopyInto(out.(*OperatorVersionList))
			return nil
		}, InType: reflect.TypeOf(&OperatorVersionList{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*OperatorVersionSpec).DeepCopyInto(out.(*OperatorVersionSpec))
			return nil
		}, InType: reflect.TypeOf(&OperatorVersionSpec{})},
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamedInstallStrategy) DeepCopyInto(out *NamedInstallStrategy) {
	*out = *in
	if in.StrategySpecRaw != nil {
		in, out := &in.StrategySpecRaw, &out.StrategySpecRaw
		*out = make(json.RawMessage, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamedInstallStrategy.
func (in *NamedInstallStrategy) DeepCopy() *NamedInstallStrategy {
	if in == nil {
		return nil
	}
	out := new(NamedInstallStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatorVersion) DeepCopyInto(out *OperatorVersion) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatorVersion.
func (in *OperatorVersion) DeepCopy() *OperatorVersion {
	if in == nil {
		return nil
	}
	out := new(OperatorVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OperatorVersion) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatorVersionList) DeepCopyInto(out *OperatorVersionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OperatorVersion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatorVersionList.
func (in *OperatorVersionList) DeepCopy() *OperatorVersionList {
	if in == nil {
		return nil
	}
	out := new(OperatorVersionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OperatorVersionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatorVersionSpec) DeepCopyInto(out *OperatorVersionSpec) {
	*out = *in
	in.InstallStrategy.DeepCopyInto(&out.InstallStrategy)
	out.Version = in.Version
	if in.Requirements != nil {
		in, out := &in.Requirements, &out.Requirements
		*out = make([]*unstructured.Unstructured, len(*in))
		for i := range *in {
			if (*in)[i] == nil {
				(*out)[i] = nil
			} else {
				(*out)[i] = new(unstructured.Unstructured)
				(*in)[i].DeepCopyInto((*out)[i])
			}
		}
	}
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatorVersionSpec.
func (in *OperatorVersionSpec) DeepCopy() *OperatorVersionSpec {
	if in == nil {
		return nil
	}
	out := new(OperatorVersionSpec)
	in.DeepCopyInto(out)
	return out
}

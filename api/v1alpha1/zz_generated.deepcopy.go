//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShareVolumePolicy) DeepCopyInto(out *ShareVolumePolicy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShareVolumePolicy.
func (in *ShareVolumePolicy) DeepCopy() *ShareVolumePolicy {
	if in == nil {
		return nil
	}
	out := new(ShareVolumePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SidecarContainer) DeepCopyInto(out *SidecarContainer) {
	*out = *in
	in.Container.DeepCopyInto(&out.Container)
	out.UpgradeStrategy = in.UpgradeStrategy
	out.ShareVolumePolicy = in.ShareVolumePolicy
	if in.TransferEnv != nil {
		in, out := &in.TransferEnv, &out.TransferEnv
		*out = make([]TransferEnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SidecarContainer.
func (in *SidecarContainer) DeepCopy() *SidecarContainer {
	if in == nil {
		return nil
	}
	out := new(SidecarContainer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SidecarContainerUpgradeStrategy) DeepCopyInto(out *SidecarContainerUpgradeStrategy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SidecarContainerUpgradeStrategy.
func (in *SidecarContainerUpgradeStrategy) DeepCopy() *SidecarContainerUpgradeStrategy {
	if in == nil {
		return nil
	}
	out := new(SidecarContainerUpgradeStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SidecarSet) DeepCopyInto(out *SidecarSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SidecarSet.
func (in *SidecarSet) DeepCopy() *SidecarSet {
	if in == nil {
		return nil
	}
	out := new(SidecarSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SidecarSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SidecarSetInjectionStrategy) DeepCopyInto(out *SidecarSetInjectionStrategy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SidecarSetInjectionStrategy.
func (in *SidecarSetInjectionStrategy) DeepCopy() *SidecarSetInjectionStrategy {
	if in == nil {
		return nil
	}
	out := new(SidecarSetInjectionStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SidecarSetList) DeepCopyInto(out *SidecarSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SidecarSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SidecarSetList.
func (in *SidecarSetList) DeepCopy() *SidecarSetList {
	if in == nil {
		return nil
	}
	out := new(SidecarSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SidecarSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SidecarSetSpec) DeepCopyInto(out *SidecarSetSpec) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.InitContainers != nil {
		in, out := &in.InitContainers, &out.InitContainers
		*out = make([]SidecarContainer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]SidecarContainer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]corev1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.UpdateStrategy.DeepCopyInto(&out.UpdateStrategy)
	out.InjectionStrategy = in.InjectionStrategy
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]corev1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.RevisionHistoryLimit != nil {
		in, out := &in.RevisionHistoryLimit, &out.RevisionHistoryLimit
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SidecarSetSpec.
func (in *SidecarSetSpec) DeepCopy() *SidecarSetSpec {
	if in == nil {
		return nil
	}
	out := new(SidecarSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SidecarSetStatus) DeepCopyInto(out *SidecarSetStatus) {
	*out = *in
	if in.CollisionCount != nil {
		in, out := &in.CollisionCount, &out.CollisionCount
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SidecarSetStatus.
func (in *SidecarSetStatus) DeepCopy() *SidecarSetStatus {
	if in == nil {
		return nil
	}
	out := new(SidecarSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SidecarSetUpdateStrategy) DeepCopyInto(out *SidecarSetUpdateStrategy) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Partition != nil {
		in, out := &in.Partition, &out.Partition
		*out = new(intstr.IntOrString)
		**out = **in
	}
	if in.MaxUnavailable != nil {
		in, out := &in.MaxUnavailable, &out.MaxUnavailable
		*out = new(intstr.IntOrString)
		**out = **in
	}
	if in.ScatterStrategy != nil {
		in, out := &in.ScatterStrategy, &out.ScatterStrategy
		*out = make(UpdateScatterStrategy, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SidecarSetUpdateStrategy.
func (in *SidecarSetUpdateStrategy) DeepCopy() *SidecarSetUpdateStrategy {
	if in == nil {
		return nil
	}
	out := new(SidecarSetUpdateStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceContainerNameSource) DeepCopyInto(out *SourceContainerNameSource) {
	*out = *in
	if in.FieldRef != nil {
		in, out := &in.FieldRef, &out.FieldRef
		*out = new(corev1.ObjectFieldSelector)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceContainerNameSource.
func (in *SourceContainerNameSource) DeepCopy() *SourceContainerNameSource {
	if in == nil {
		return nil
	}
	out := new(SourceContainerNameSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransferEnvVar) DeepCopyInto(out *TransferEnvVar) {
	*out = *in
	if in.SourceContainerNameFrom != nil {
		in, out := &in.SourceContainerNameFrom, &out.SourceContainerNameFrom
		*out = new(SourceContainerNameSource)
		(*in).DeepCopyInto(*out)
	}
	if in.EnvNames != nil {
		in, out := &in.EnvNames, &out.EnvNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransferEnvVar.
func (in *TransferEnvVar) DeepCopy() *TransferEnvVar {
	if in == nil {
		return nil
	}
	out := new(TransferEnvVar)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in UpdateScatterStrategy) DeepCopyInto(out *UpdateScatterStrategy) {
	{
		in := &in
		*out = make(UpdateScatterStrategy, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpdateScatterStrategy.
func (in UpdateScatterStrategy) DeepCopy() UpdateScatterStrategy {
	if in == nil {
		return nil
	}
	out := new(UpdateScatterStrategy)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpdateScatterTerm) DeepCopyInto(out *UpdateScatterTerm) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpdateScatterTerm.
func (in *UpdateScatterTerm) DeepCopy() *UpdateScatterTerm {
	if in == nil {
		return nil
	}
	out := new(UpdateScatterTerm)
	in.DeepCopyInto(out)
	return out
}

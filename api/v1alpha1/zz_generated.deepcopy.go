// +build !ignore_autogenerated

/*

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

// autogenerated by controller-gen object, do not modify manually

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwsUpgradeStrategy) DeepCopyInto(out *AwsUpgradeStrategy) {
	*out = *in
	out.CRDType = in.CRDType
	in.RollingUpgradeType.DeepCopyInto(&out.RollingUpgradeType)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsUpgradeStrategy.
func (in *AwsUpgradeStrategy) DeepCopy() *AwsUpgradeStrategy {
	if in == nil {
		return nil
	}
	out := new(AwsUpgradeStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CRDUpgradeStrategy) DeepCopyInto(out *CRDUpgradeStrategy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CRDUpgradeStrategy.
func (in *CRDUpgradeStrategy) DeepCopy() *CRDUpgradeStrategy {
	if in == nil {
		return nil
	}
	out := new(CRDUpgradeStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSCFConfiguration) DeepCopyInto(out *EKSCFConfiguration) {
	*out = *in
	if in.NodeSecurityGroups != nil {
		in, out := &in.NodeSecurityGroups, &out.NodeSecurityGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]map[string]string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make(map[string]string, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
		}
	}
	if in.ManagedPolicies != nil {
		in, out := &in.ManagedPolicies, &out.ManagedPolicies
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MetricsCollection != nil {
		in, out := &in.MetricsCollection, &out.MetricsCollection
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSCFConfiguration.
func (in *EKSCFConfiguration) DeepCopy() *EKSCFConfiguration {
	if in == nil {
		return nil
	}
	out := new(EKSCFConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSCFSpec) DeepCopyInto(out *EKSCFSpec) {
	*out = *in
	in.EKSCFConfiguration.DeepCopyInto(&out.EKSCFConfiguration)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSCFSpec.
func (in *EKSCFSpec) DeepCopy() *EKSCFSpec {
	if in == nil {
		return nil
	}
	out := new(EKSCFSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceGroup) DeepCopyInto(out *InstanceGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceGroup.
func (in *InstanceGroup) DeepCopy() *InstanceGroup {
	if in == nil {
		return nil
	}
	out := new(InstanceGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InstanceGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceGroupList) DeepCopyInto(out *InstanceGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InstanceGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceGroupList.
func (in *InstanceGroupList) DeepCopy() *InstanceGroupList {
	if in == nil {
		return nil
	}
	out := new(InstanceGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InstanceGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceGroupSpec) DeepCopyInto(out *InstanceGroupSpec) {
	*out = *in
	in.EKSCFSpec.DeepCopyInto(&out.EKSCFSpec)
	in.AwsUpgradeStrategy.DeepCopyInto(&out.AwsUpgradeStrategy)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceGroupSpec.
func (in *InstanceGroupSpec) DeepCopy() *InstanceGroupSpec {
	if in == nil {
		return nil
	}
	out := new(InstanceGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceGroupStatus) DeepCopyInto(out *InstanceGroupStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceGroupStatus.
func (in *InstanceGroupStatus) DeepCopy() *InstanceGroupStatus {
	if in == nil {
		return nil
	}
	out := new(InstanceGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RollingUpgradeStrategy) DeepCopyInto(out *RollingUpgradeStrategy) {
	*out = *in
	if in.SuspendProcesses != nil {
		in, out := &in.SuspendProcesses, &out.SuspendProcesses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RollingUpgradeStrategy.
func (in *RollingUpgradeStrategy) DeepCopy() *RollingUpgradeStrategy {
	if in == nil {
		return nil
	}
	out := new(RollingUpgradeStrategy)
	in.DeepCopyInto(out)
	return out
}

// +build !ignore_autogenerated

/*
 Copyright 2021 The Hybridnet Authors.

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

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Address) DeepCopyInto(out *Address) {
	*out = *in
	if in.NetID != nil {
		in, out := &in.NetID, &out.NetID
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Address.
func (in *Address) DeepCopy() *Address {
	if in == nil {
		return nil
	}
	out := new(Address)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddressRange) DeepCopyInto(out *AddressRange) {
	*out = *in
	if in.ReservedIPs != nil {
		in, out := &in.ReservedIPs, &out.ReservedIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExcludeIPs != nil {
		in, out := &in.ExcludeIPs, &out.ExcludeIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddressRange.
func (in *AddressRange) DeepCopy() *AddressRange {
	if in == nil {
		return nil
	}
	out := new(AddressRange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPPeer) DeepCopyInto(out *BGPPeer) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPPeer.
func (in *BGPPeer) DeepCopy() *BGPPeer {
	if in == nil {
		return nil
	}
	out := new(BGPPeer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Binding) DeepCopyInto(out *Binding) {
	*out = *in
	out.ReferredObject = in.ReferredObject
	if in.Stateful != nil {
		in, out := &in.Stateful, &out.Stateful
		*out = new(StatefulInfo)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Binding.
func (in *Binding) DeepCopy() *Binding {
	if in == nil {
		return nil
	}
	out := new(Binding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Count) DeepCopyInto(out *Count) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Count.
func (in *Count) DeepCopy() *Count {
	if in == nil {
		return nil
	}
	out := new(Count)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPInstance) DeepCopyInto(out *IPInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPInstance.
func (in *IPInstance) DeepCopy() *IPInstance {
	if in == nil {
		return nil
	}
	out := new(IPInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IPInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPInstanceList) DeepCopyInto(out *IPInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IPInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPInstanceList.
func (in *IPInstanceList) DeepCopy() *IPInstanceList {
	if in == nil {
		return nil
	}
	out := new(IPInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IPInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in IPInstancePointerSlice) DeepCopyInto(out *IPInstancePointerSlice) {
	{
		in := &in
		*out = make(IPInstancePointerSlice, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(IPInstance)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPInstancePointerSlice.
func (in IPInstancePointerSlice) DeepCopy() IPInstancePointerSlice {
	if in == nil {
		return nil
	}
	out := new(IPInstancePointerSlice)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPInstanceSpec) DeepCopyInto(out *IPInstanceSpec) {
	*out = *in
	in.Address.DeepCopyInto(&out.Address)
	in.Binding.DeepCopyInto(&out.Binding)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPInstanceSpec.
func (in *IPInstanceSpec) DeepCopy() *IPInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(IPInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPInstanceStatus) DeepCopyInto(out *IPInstanceStatus) {
	*out = *in
	in.UpdateTimestamp.DeepCopyInto(&out.UpdateTimestamp)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPInstanceStatus.
func (in *IPInstanceStatus) DeepCopy() *IPInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(IPInstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Network) DeepCopyInto(out *Network) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Network.
func (in *Network) DeepCopy() *Network {
	if in == nil {
		return nil
	}
	out := new(Network)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Network) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkConfig) DeepCopyInto(out *NetworkConfig) {
	*out = *in
	if in.BGPPeers != nil {
		in, out := &in.BGPPeers, &out.BGPPeers
		*out = make([]BGPPeer, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkConfig.
func (in *NetworkConfig) DeepCopy() *NetworkConfig {
	if in == nil {
		return nil
	}
	out := new(NetworkConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkList) DeepCopyInto(out *NetworkList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Network, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkList.
func (in *NetworkList) DeepCopy() *NetworkList {
	if in == nil {
		return nil
	}
	out := new(NetworkList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkSpec) DeepCopyInto(out *NetworkSpec) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NetID != nil {
		in, out := &in.NetID, &out.NetID
		*out = new(int32)
		**out = **in
	}
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(NetworkConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkSpec.
func (in *NetworkSpec) DeepCopy() *NetworkSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkStatus) DeepCopyInto(out *NetworkStatus) {
	*out = *in
	if in.SubnetList != nil {
		in, out := &in.SubnetList, &out.SubnetList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NodeList != nil {
		in, out := &in.NodeList, &out.NodeList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Statistics != nil {
		in, out := &in.Statistics, &out.Statistics
		*out = new(Count)
		**out = **in
	}
	if in.IPv6Statistics != nil {
		in, out := &in.IPv6Statistics, &out.IPv6Statistics
		*out = new(Count)
		**out = **in
	}
	if in.DualStackStatistics != nil {
		in, out := &in.DualStackStatistics, &out.DualStackStatistics
		*out = new(Count)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkStatus.
func (in *NetworkStatus) DeepCopy() *NetworkStatus {
	if in == nil {
		return nil
	}
	out := new(NetworkStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeInfo) DeepCopyInto(out *NodeInfo) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeInfo.
func (in *NodeInfo) DeepCopy() *NodeInfo {
	if in == nil {
		return nil
	}
	out := new(NodeInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeInfo) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeInfoList) DeepCopyInto(out *NodeInfoList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NodeInfo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeInfoList.
func (in *NodeInfoList) DeepCopy() *NodeInfoList {
	if in == nil {
		return nil
	}
	out := new(NodeInfoList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeInfoList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeInfoSpec) DeepCopyInto(out *NodeInfoSpec) {
	*out = *in
	if in.VTEPInfo != nil {
		in, out := &in.VTEPInfo, &out.VTEPInfo
		*out = new(VTEPInfo)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeInfoSpec.
func (in *NodeInfoSpec) DeepCopy() *NodeInfoSpec {
	if in == nil {
		return nil
	}
	out := new(NodeInfoSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeInfoStatus) DeepCopyInto(out *NodeInfoStatus) {
	*out = *in
	in.UpdateTimestamp.DeepCopyInto(&out.UpdateTimestamp)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeInfoStatus.
func (in *NodeInfoStatus) DeepCopy() *NodeInfoStatus {
	if in == nil {
		return nil
	}
	out := new(NodeInfoStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectMeta) DeepCopyInto(out *ObjectMeta) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectMeta.
func (in *ObjectMeta) DeepCopy() *ObjectMeta {
	if in == nil {
		return nil
	}
	out := new(ObjectMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatefulInfo) DeepCopyInto(out *StatefulInfo) {
	*out = *in
	if in.Index != nil {
		in, out := &in.Index, &out.Index
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatefulInfo.
func (in *StatefulInfo) DeepCopy() *StatefulInfo {
	if in == nil {
		return nil
	}
	out := new(StatefulInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Subnet) DeepCopyInto(out *Subnet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Subnet.
func (in *Subnet) DeepCopy() *Subnet {
	if in == nil {
		return nil
	}
	out := new(Subnet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Subnet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetConfig) DeepCopyInto(out *SubnetConfig) {
	*out = *in
	if in.AutoNatOutgoing != nil {
		in, out := &in.AutoNatOutgoing, &out.AutoNatOutgoing
		*out = new(bool)
		**out = **in
	}
	if in.Private != nil {
		in, out := &in.Private, &out.Private
		*out = new(bool)
		**out = **in
	}
	if in.AllowSubnets != nil {
		in, out := &in.AllowSubnets, &out.AllowSubnets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetConfig.
func (in *SubnetConfig) DeepCopy() *SubnetConfig {
	if in == nil {
		return nil
	}
	out := new(SubnetConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetList) DeepCopyInto(out *SubnetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Subnet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetList.
func (in *SubnetList) DeepCopy() *SubnetList {
	if in == nil {
		return nil
	}
	out := new(SubnetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SubnetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetSpec) DeepCopyInto(out *SubnetSpec) {
	*out = *in
	in.Range.DeepCopyInto(&out.Range)
	if in.NetID != nil {
		in, out := &in.NetID, &out.NetID
		*out = new(int32)
		**out = **in
	}
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(SubnetConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetSpec.
func (in *SubnetSpec) DeepCopy() *SubnetSpec {
	if in == nil {
		return nil
	}
	out := new(SubnetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetStatus) DeepCopyInto(out *SubnetStatus) {
	*out = *in
	out.Count = in.Count
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetStatus.
func (in *SubnetStatus) DeepCopy() *SubnetStatus {
	if in == nil {
		return nil
	}
	out := new(SubnetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VTEPInfo) DeepCopyInto(out *VTEPInfo) {
	*out = *in
	if in.LocalIPs != nil {
		in, out := &in.LocalIPs, &out.LocalIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VTEPInfo.
func (in *VTEPInfo) DeepCopy() *VTEPInfo {
	if in == nil {
		return nil
	}
	out := new(VTEPInfo)
	in.DeepCopyInto(out)
	return out
}

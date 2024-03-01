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

package v1beta1

import (
	memcachedv1beta1 "github.com/openstack-k8s-operators/infra-operator/apis/memcached/v1beta1"
	redisv1beta1 "github.com/openstack-k8s-operators/infra-operator/apis/redis/v1beta1"
	"github.com/openstack-k8s-operators/lib-common/modules/common/condition"
	"github.com/openstack-k8s-operators/lib-common/modules/common/route"
	"github.com/openstack-k8s-operators/lib-common/modules/storage"
	apiv1beta1 "github.com/openstack-k8s-operators/mariadb-operator/api/v1beta1"
	ovn_operatorapiv1beta1 "github.com/openstack-k8s-operators/ovn-operator/api/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BarbicanSection) DeepCopyInto(out *BarbicanSection) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	in.APIOverride.DeepCopyInto(&out.APIOverride)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BarbicanSection.
func (in *BarbicanSection) DeepCopy() *BarbicanSection {
	if in == nil {
		return nil
	}
	out := new(BarbicanSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertConfig) DeepCopyInto(out *CertConfig) {
	*out = *in
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(v1.Duration)
		**out = **in
	}
	if in.RenewBefore != nil {
		in, out := &in.RenewBefore, &out.RenewBefore
		*out = new(v1.Duration)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertConfig.
func (in *CertConfig) DeepCopy() *CertConfig {
	if in == nil {
		return nil
	}
	out := new(CertConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertSection) DeepCopyInto(out *CertSection) {
	*out = *in
	in.Cert.DeepCopyInto(&out.Cert)
	in.Ca.DeepCopyInto(&out.Ca)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertSection.
func (in *CertSection) DeepCopy() *CertSection {
	if in == nil {
		return nil
	}
	out := new(CertSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CinderSection) DeepCopyInto(out *CinderSection) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	in.APIOverride.DeepCopyInto(&out.APIOverride)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CinderSection.
func (in *CinderSection) DeepCopy() *CinderSection {
	if in == nil {
		return nil
	}
	out := new(CinderSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSMasqSection) DeepCopyInto(out *DNSMasqSection) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSMasqSection.
func (in *DNSMasqSection) DeepCopy() *DNSMasqSection {
	if in == nil {
		return nil
	}
	out := new(DNSMasqSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesignateSection) DeepCopyInto(out *DesignateSection) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	in.APIOverride.DeepCopyInto(&out.APIOverride)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesignateSection.
func (in *DesignateSection) DeepCopy() *DesignateSection {
	if in == nil {
		return nil
	}
	out := new(DesignateSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GaleraSection) DeepCopyInto(out *GaleraSection) {
	*out = *in
	if in.Templates != nil {
		in, out := &in.Templates, &out.Templates
		*out = make(map[string]apiv1beta1.GaleraSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GaleraSection.
func (in *GaleraSection) DeepCopy() *GaleraSection {
	if in == nil {
		return nil
	}
	out := new(GaleraSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlanceSection) DeepCopyInto(out *GlanceSection) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	if in.APIOverride != nil {
		in, out := &in.APIOverride, &out.APIOverride
		*out = make(map[string]Override, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlanceSection.
func (in *GlanceSection) DeepCopy() *GlanceSection {
	if in == nil {
		return nil
	}
	out := new(GlanceSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HeatSection) DeepCopyInto(out *HeatSection) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	in.APIOverride.DeepCopyInto(&out.APIOverride)
	in.CnfAPIOverride.DeepCopyInto(&out.CnfAPIOverride)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HeatSection.
func (in *HeatSection) DeepCopy() *HeatSection {
	if in == nil {
		return nil
	}
	out := new(HeatSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HorizonSection) DeepCopyInto(out *HorizonSection) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	in.APIOverride.DeepCopyInto(&out.APIOverride)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HorizonSection.
func (in *HorizonSection) DeepCopy() *HorizonSection {
	if in == nil {
		return nil
	}
	out := new(HorizonSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IronicSection) DeepCopyInto(out *IronicSection) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	in.APIOverride.DeepCopyInto(&out.APIOverride)
	in.InspectorOverride.DeepCopyInto(&out.InspectorOverride)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IronicSection.
func (in *IronicSection) DeepCopy() *IronicSection {
	if in == nil {
		return nil
	}
	out := new(IronicSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeystoneSection) DeepCopyInto(out *KeystoneSection) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	in.APIOverride.DeepCopyInto(&out.APIOverride)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeystoneSection.
func (in *KeystoneSection) DeepCopy() *KeystoneSection {
	if in == nil {
		return nil
	}
	out := new(KeystoneSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManilaSection) DeepCopyInto(out *ManilaSection) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	in.APIOverride.DeepCopyInto(&out.APIOverride)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManilaSection.
func (in *ManilaSection) DeepCopy() *ManilaSection {
	if in == nil {
		return nil
	}
	out := new(ManilaSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemcachedSection) DeepCopyInto(out *MemcachedSection) {
	*out = *in
	if in.Templates != nil {
		in, out := &in.Templates, &out.Templates
		*out = make(map[string]memcachedv1beta1.MemcachedSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemcachedSection.
func (in *MemcachedSection) DeepCopy() *MemcachedSection {
	if in == nil {
		return nil
	}
	out := new(MemcachedSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NeutronSection) DeepCopyInto(out *NeutronSection) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	in.APIOverride.DeepCopyInto(&out.APIOverride)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NeutronSection.
func (in *NeutronSection) DeepCopy() *NeutronSection {
	if in == nil {
		return nil
	}
	out := new(NeutronSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NovaCellOverrideSpec) DeepCopyInto(out *NovaCellOverrideSpec) {
	*out = *in
	in.NoVNCProxy.DeepCopyInto(&out.NoVNCProxy)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NovaCellOverrideSpec.
func (in *NovaCellOverrideSpec) DeepCopy() *NovaCellOverrideSpec {
	if in == nil {
		return nil
	}
	out := new(NovaCellOverrideSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NovaSection) DeepCopyInto(out *NovaSection) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	in.APIOverride.DeepCopyInto(&out.APIOverride)
	if in.CellOverride != nil {
		in, out := &in.CellOverride, &out.CellOverride
		*out = make(map[string]NovaCellOverrideSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NovaSection.
func (in *NovaSection) DeepCopy() *NovaSection {
	if in == nil {
		return nil
	}
	out := new(NovaSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OctaviaSection) DeepCopyInto(out *OctaviaSection) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	in.APIOverride.DeepCopyInto(&out.APIOverride)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OctaviaSection.
func (in *OctaviaSection) DeepCopy() *OctaviaSection {
	if in == nil {
		return nil
	}
	out := new(OctaviaSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackClientSection) DeepCopyInto(out *OpenStackClientSection) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackClientSection.
func (in *OpenStackClientSection) DeepCopy() *OpenStackClientSection {
	if in == nil {
		return nil
	}
	out := new(OpenStackClientSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackControlPlane) DeepCopyInto(out *OpenStackControlPlane) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackControlPlane.
func (in *OpenStackControlPlane) DeepCopy() *OpenStackControlPlane {
	if in == nil {
		return nil
	}
	out := new(OpenStackControlPlane)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackControlPlane) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackControlPlaneDefaults) DeepCopyInto(out *OpenStackControlPlaneDefaults) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackControlPlaneDefaults.
func (in *OpenStackControlPlaneDefaults) DeepCopy() *OpenStackControlPlaneDefaults {
	if in == nil {
		return nil
	}
	out := new(OpenStackControlPlaneDefaults)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackControlPlaneList) DeepCopyInto(out *OpenStackControlPlaneList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpenStackControlPlane, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackControlPlaneList.
func (in *OpenStackControlPlaneList) DeepCopy() *OpenStackControlPlaneList {
	if in == nil {
		return nil
	}
	out := new(OpenStackControlPlaneList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackControlPlaneList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackControlPlaneSpec) DeepCopyInto(out *OpenStackControlPlaneSpec) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.TLS.DeepCopyInto(&out.TLS)
	in.DNS.DeepCopyInto(&out.DNS)
	in.Keystone.DeepCopyInto(&out.Keystone)
	in.Placement.DeepCopyInto(&out.Placement)
	in.Glance.DeepCopyInto(&out.Glance)
	in.Cinder.DeepCopyInto(&out.Cinder)
	in.Galera.DeepCopyInto(&out.Galera)
	in.Rabbitmq.DeepCopyInto(&out.Rabbitmq)
	in.Memcached.DeepCopyInto(&out.Memcached)
	in.Ovn.DeepCopyInto(&out.Ovn)
	in.Neutron.DeepCopyInto(&out.Neutron)
	in.Nova.DeepCopyInto(&out.Nova)
	in.Heat.DeepCopyInto(&out.Heat)
	in.Ironic.DeepCopyInto(&out.Ironic)
	in.Manila.DeepCopyInto(&out.Manila)
	in.Horizon.DeepCopyInto(&out.Horizon)
	in.Telemetry.DeepCopyInto(&out.Telemetry)
	in.Swift.DeepCopyInto(&out.Swift)
	in.Octavia.DeepCopyInto(&out.Octavia)
	in.Designate.DeepCopyInto(&out.Designate)
	in.Barbican.DeepCopyInto(&out.Barbican)
	in.Redis.DeepCopyInto(&out.Redis)
	in.OpenStackClient.DeepCopyInto(&out.OpenStackClient)
	if in.ExtraMounts != nil {
		in, out := &in.ExtraMounts, &out.ExtraMounts
		*out = make([]OpenStackExtraVolMounts, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackControlPlaneSpec.
func (in *OpenStackControlPlaneSpec) DeepCopy() *OpenStackControlPlaneSpec {
	if in == nil {
		return nil
	}
	out := new(OpenStackControlPlaneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackControlPlaneStatus) DeepCopyInto(out *OpenStackControlPlaneStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(condition.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.TLS.DeepCopyInto(&out.TLS)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackControlPlaneStatus.
func (in *OpenStackControlPlaneStatus) DeepCopy() *OpenStackControlPlaneStatus {
	if in == nil {
		return nil
	}
	out := new(OpenStackControlPlaneStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackExtraVolMounts) DeepCopyInto(out *OpenStackExtraVolMounts) {
	*out = *in
	if in.VolMounts != nil {
		in, out := &in.VolMounts, &out.VolMounts
		*out = make([]storage.VolMounts, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackExtraVolMounts.
func (in *OpenStackExtraVolMounts) DeepCopy() *OpenStackExtraVolMounts {
	if in == nil {
		return nil
	}
	out := new(OpenStackExtraVolMounts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Override) DeepCopyInto(out *Override) {
	*out = *in
	if in.Route != nil {
		in, out := &in.Route, &out.Route
		*out = new(route.OverrideSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLSServiceOverride)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Override.
func (in *Override) DeepCopy() *Override {
	if in == nil {
		return nil
	}
	out := new(Override)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OvnResources) DeepCopyInto(out *OvnResources) {
	*out = *in
	if in.OVNDBCluster != nil {
		in, out := &in.OVNDBCluster, &out.OVNDBCluster
		*out = make(map[string]ovn_operatorapiv1beta1.OVNDBClusterSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	in.OVNNorthd.DeepCopyInto(&out.OVNNorthd)
	in.OVNController.DeepCopyInto(&out.OVNController)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OvnResources.
func (in *OvnResources) DeepCopy() *OvnResources {
	if in == nil {
		return nil
	}
	out := new(OvnResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OvnSection) DeepCopyInto(out *OvnSection) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OvnSection.
func (in *OvnSection) DeepCopy() *OvnSection {
	if in == nil {
		return nil
	}
	out := new(OvnSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlacementSection) DeepCopyInto(out *PlacementSection) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	in.APIOverride.DeepCopyInto(&out.APIOverride)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlacementSection.
func (in *PlacementSection) DeepCopy() *PlacementSection {
	if in == nil {
		return nil
	}
	out := new(PlacementSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqSection) DeepCopyInto(out *RabbitmqSection) {
	*out = *in
	if in.Templates != nil {
		in, out := &in.Templates, &out.Templates
		*out = make(map[string]RabbitmqTemplate, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqSection.
func (in *RabbitmqSection) DeepCopy() *RabbitmqSection {
	if in == nil {
		return nil
	}
	out := new(RabbitmqSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqTemplate) DeepCopyInto(out *RabbitmqTemplate) {
	*out = *in
	in.RabbitmqClusterSpec.DeepCopyInto(&out.RabbitmqClusterSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqTemplate.
func (in *RabbitmqTemplate) DeepCopy() *RabbitmqTemplate {
	if in == nil {
		return nil
	}
	out := new(RabbitmqTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisSection) DeepCopyInto(out *RedisSection) {
	*out = *in
	if in.Templates != nil {
		in, out := &in.Templates, &out.Templates
		*out = make(map[string]redisv1beta1.RedisSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisSection.
func (in *RedisSection) DeepCopy() *RedisSection {
	if in == nil {
		return nil
	}
	out := new(RedisSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwiftSection) DeepCopyInto(out *SwiftSection) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	in.ProxyOverride.DeepCopyInto(&out.ProxyOverride)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwiftSection.
func (in *SwiftSection) DeepCopy() *SwiftSection {
	if in == nil {
		return nil
	}
	out := new(SwiftSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSCAStatus) DeepCopyInto(out *TLSCAStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSCAStatus.
func (in *TLSCAStatus) DeepCopy() *TLSCAStatus {
	if in == nil {
		return nil
	}
	out := new(TLSCAStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSIngressConfig) DeepCopyInto(out *TLSIngressConfig) {
	*out = *in
	in.CertSection.DeepCopyInto(&out.CertSection)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSIngressConfig.
func (in *TLSIngressConfig) DeepCopy() *TLSIngressConfig {
	if in == nil {
		return nil
	}
	out := new(TLSIngressConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSPodLevelConfig) DeepCopyInto(out *TLSPodLevelConfig) {
	*out = *in
	in.Default.DeepCopyInto(&out.Default)
	in.Ovn.DeepCopyInto(&out.Ovn)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSPodLevelConfig.
func (in *TLSPodLevelConfig) DeepCopy() *TLSPodLevelConfig {
	if in == nil {
		return nil
	}
	out := new(TLSPodLevelConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSSection) DeepCopyInto(out *TLSSection) {
	*out = *in
	in.Ingress.DeepCopyInto(&out.Ingress)
	in.PodLevel.DeepCopyInto(&out.PodLevel)
	out.Ca = in.Ca
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSSection.
func (in *TLSSection) DeepCopy() *TLSSection {
	if in == nil {
		return nil
	}
	out := new(TLSSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSServiceOverride) DeepCopyInto(out *TLSServiceOverride) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSServiceOverride.
func (in *TLSServiceOverride) DeepCopy() *TLSServiceOverride {
	if in == nil {
		return nil
	}
	out := new(TLSServiceOverride)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSStatus) DeepCopyInto(out *TLSStatus) {
	*out = *in
	if in.CAList != nil {
		in, out := &in.CAList, &out.CAList
		*out = make([]TLSCAStatus, len(*in))
		copy(*out, *in)
	}
	out.Ca = in.Ca
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSStatus.
func (in *TLSStatus) DeepCopy() *TLSStatus {
	if in == nil {
		return nil
	}
	out := new(TLSStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TelemetrySection) DeepCopyInto(out *TelemetrySection) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TelemetrySection.
func (in *TelemetrySection) DeepCopy() *TelemetrySection {
	if in == nil {
		return nil
	}
	out := new(TelemetrySection)
	in.DeepCopyInto(out)
	return out
}

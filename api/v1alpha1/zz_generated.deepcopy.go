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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyticsEngineSpec) DeepCopyInto(out *AnalyticsEngineSpec) {
	*out = *in
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(ServiceSpec)
		(*in).DeepCopyInto(*out)
	}
	in.Deployment.DeepCopyInto(&out.Deployment)
	if in.MetricsBackend != nil {
		in, out := &in.MetricsBackend, &out.MetricsBackend
		*out = new(MetricsBackendSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyticsEngineSpec.
func (in *AnalyticsEngineSpec) DeepCopy() *AnalyticsEngineSpec {
	if in == nil {
		return nil
	}
	out := new(AnalyticsEngineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerSpec) DeepCopyInto(out *ControllerSpec) {
	*out = *in
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(ServiceSpec)
		(*in).DeepCopyInto(*out)
	}
	in.Deployment.DeepCopyInto(&out.Deployment)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerSpec.
func (in *ControllerSpec) DeepCopy() *ControllerSpec {
	if in == nil {
		return nil
	}
	out := new(ControllerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CounterMetricSpec) DeepCopyInto(out *CounterMetricSpec) {
	*out = *in
	if in.PreferredDirection != nil {
		in, out := &in.PreferredDirection, &out.PreferredDirection
		*out = new(string)
		**out = **in
	}
	if in.Units != nil {
		in, out := &in.Units, &out.Units
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CounterMetricSpec.
func (in *CounterMetricSpec) DeepCopy() *CounterMetricSpec {
	if in == nil {
		return nil
	}
	out := new(CounterMetricSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentSpec) DeepCopyInto(out *DeploymentSpec) {
	*out = *in
	if in.ReplicaCount != nil {
		in, out := &in.ReplicaCount, &out.ReplicaCount
		*out = new(int32)
		**out = **in
	}
	if in.ImagePullPolicy != nil {
		in, out := &in.ImagePullPolicy, &out.ImagePullPolicy
		*out = new(v1.PullPolicy)
		**out = **in
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentSpec.
func (in *DeploymentSpec) DeepCopy() *DeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(DeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Iter8) DeepCopyInto(out *Iter8) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Iter8.
func (in *Iter8) DeepCopy() *Iter8 {
	if in == nil {
		return nil
	}
	out := new(Iter8)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Iter8) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Iter8List) DeepCopyInto(out *Iter8List) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Iter8, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Iter8List.
func (in *Iter8List) DeepCopy() *Iter8List {
	if in == nil {
		return nil
	}
	out := new(Iter8List)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Iter8List) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Iter8Spec) DeepCopyInto(out *Iter8Spec) {
	*out = *in
	in.Controller.DeepCopyInto(&out.Controller)
	in.AnalyticsEngine.DeepCopyInto(&out.AnalyticsEngine)
	in.Metrics.DeepCopyInto(&out.Metrics)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Iter8Spec.
func (in *Iter8Spec) DeepCopy() *Iter8Spec {
	if in == nil {
		return nil
	}
	out := new(Iter8Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Iter8Status) DeepCopyInto(out *Iter8Status) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Iter8Status.
func (in *Iter8Status) DeepCopy() *Iter8Status {
	if in == nil {
		return nil
	}
	out := new(Iter8Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsBackendAuthenticationSpec) DeepCopyInto(out *MetricsBackendAuthenticationSpec) {
	*out = *in
	if in.InsecureSkipVerify != nil {
		in, out := &in.InsecureSkipVerify, &out.InsecureSkipVerify
		*out = new(bool)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsBackendAuthenticationSpec.
func (in *MetricsBackendAuthenticationSpec) DeepCopy() *MetricsBackendAuthenticationSpec {
	if in == nil {
		return nil
	}
	out := new(MetricsBackendAuthenticationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsBackendSpec) DeepCopyInto(out *MetricsBackendSpec) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
	if in.Authentication != nil {
		in, out := &in.Authentication, &out.Authentication
		*out = new(MetricsBackendAuthenticationSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsBackendSpec.
func (in *MetricsBackendSpec) DeepCopy() *MetricsBackendSpec {
	if in == nil {
		return nil
	}
	out := new(MetricsBackendSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsSpec) DeepCopyInto(out *MetricsSpec) {
	*out = *in
	if in.CounterMetrics != nil {
		in, out := &in.CounterMetrics, &out.CounterMetrics
		*out = new([]CounterMetricSpec)
		if **in != nil {
			in, out := *in, *out
			*out = make([]CounterMetricSpec, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.RatioMetrics != nil {
		in, out := &in.RatioMetrics, &out.RatioMetrics
		*out = new([]RatioMetricSpec)
		if **in != nil {
			in, out := *in, *out
			*out = make([]RatioMetricSpec, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsSpec.
func (in *MetricsSpec) DeepCopy() *MetricsSpec {
	if in == nil {
		return nil
	}
	out := new(MetricsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RatioMetricSpec) DeepCopyInto(out *RatioMetricSpec) {
	*out = *in
	if in.PreferredDirection != nil {
		in, out := &in.PreferredDirection, &out.PreferredDirection
		*out = new(string)
		**out = **in
	}
	if in.ZeroToOne != nil {
		in, out := &in.ZeroToOne, &out.ZeroToOne
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RatioMetricSpec.
func (in *RatioMetricSpec) DeepCopy() *RatioMetricSpec {
	if in == nil {
		return nil
	}
	out := new(RatioMetricSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSpec) DeepCopyInto(out *ServiceSpec) {
	*out = *in
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSpec.
func (in *ServiceSpec) DeepCopy() *ServiceSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceSpec)
	in.DeepCopyInto(out)
	return out
}

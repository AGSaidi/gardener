//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (c) SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	config "github.com/gardener/gardener/pkg/admissioncontroller/apis/config"
	v1 "k8s.io/api/rbac/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	componentbaseconfig "k8s.io/component-base/config"
	configv1alpha1 "k8s.io/component-base/config/v1alpha1"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*AdmissionControllerConfiguration)(nil), (*config.AdmissionControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_AdmissionControllerConfiguration_To_config_AdmissionControllerConfiguration(a.(*AdmissionControllerConfiguration), b.(*config.AdmissionControllerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.AdmissionControllerConfiguration)(nil), (*AdmissionControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_AdmissionControllerConfiguration_To_v1alpha1_AdmissionControllerConfiguration(a.(*config.AdmissionControllerConfiguration), b.(*AdmissionControllerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*HTTPSServer)(nil), (*config.HTTPSServer)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_HTTPSServer_To_config_HTTPSServer(a.(*HTTPSServer), b.(*config.HTTPSServer), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.HTTPSServer)(nil), (*HTTPSServer)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_HTTPSServer_To_v1alpha1_HTTPSServer(a.(*config.HTTPSServer), b.(*HTTPSServer), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ResourceAdmissionConfiguration)(nil), (*config.ResourceAdmissionConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ResourceAdmissionConfiguration_To_config_ResourceAdmissionConfiguration(a.(*ResourceAdmissionConfiguration), b.(*config.ResourceAdmissionConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ResourceAdmissionConfiguration)(nil), (*ResourceAdmissionConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ResourceAdmissionConfiguration_To_v1alpha1_ResourceAdmissionConfiguration(a.(*config.ResourceAdmissionConfiguration), b.(*ResourceAdmissionConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ResourceLimit)(nil), (*config.ResourceLimit)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ResourceLimit_To_config_ResourceLimit(a.(*ResourceLimit), b.(*config.ResourceLimit), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ResourceLimit)(nil), (*ResourceLimit)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ResourceLimit_To_v1alpha1_ResourceLimit(a.(*config.ResourceLimit), b.(*ResourceLimit), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Server)(nil), (*config.Server)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Server_To_config_Server(a.(*Server), b.(*config.Server), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.Server)(nil), (*Server)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_Server_To_v1alpha1_Server(a.(*config.Server), b.(*Server), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ServerConfiguration)(nil), (*config.ServerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ServerConfiguration_To_config_ServerConfiguration(a.(*ServerConfiguration), b.(*config.ServerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ServerConfiguration)(nil), (*ServerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ServerConfiguration_To_v1alpha1_ServerConfiguration(a.(*config.ServerConfiguration), b.(*ServerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*TLSServer)(nil), (*config.TLSServer)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_TLSServer_To_config_TLSServer(a.(*TLSServer), b.(*config.TLSServer), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.TLSServer)(nil), (*TLSServer)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_TLSServer_To_v1alpha1_TLSServer(a.(*config.TLSServer), b.(*TLSServer), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_AdmissionControllerConfiguration_To_config_AdmissionControllerConfiguration(in *AdmissionControllerConfiguration, out *config.AdmissionControllerConfiguration, s conversion.Scope) error {
	if err := configv1alpha1.Convert_v1alpha1_ClientConnectionConfiguration_To_config_ClientConnectionConfiguration(&in.GardenClientConnection, &out.GardenClientConnection, s); err != nil {
		return err
	}
	out.LogLevel = in.LogLevel
	if err := Convert_v1alpha1_ServerConfiguration_To_config_ServerConfiguration(&in.Server, &out.Server, s); err != nil {
		return err
	}
	if in.Debugging != nil {
		in, out := &in.Debugging, &out.Debugging
		*out = new(componentbaseconfig.DebuggingConfiguration)
		if err := configv1alpha1.Convert_v1alpha1_DebuggingConfiguration_To_config_DebuggingConfiguration(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Debugging = nil
	}
	return nil
}

// Convert_v1alpha1_AdmissionControllerConfiguration_To_config_AdmissionControllerConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_AdmissionControllerConfiguration_To_config_AdmissionControllerConfiguration(in *AdmissionControllerConfiguration, out *config.AdmissionControllerConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_AdmissionControllerConfiguration_To_config_AdmissionControllerConfiguration(in, out, s)
}

func autoConvert_config_AdmissionControllerConfiguration_To_v1alpha1_AdmissionControllerConfiguration(in *config.AdmissionControllerConfiguration, out *AdmissionControllerConfiguration, s conversion.Scope) error {
	if err := configv1alpha1.Convert_config_ClientConnectionConfiguration_To_v1alpha1_ClientConnectionConfiguration(&in.GardenClientConnection, &out.GardenClientConnection, s); err != nil {
		return err
	}
	out.LogLevel = in.LogLevel
	if err := Convert_config_ServerConfiguration_To_v1alpha1_ServerConfiguration(&in.Server, &out.Server, s); err != nil {
		return err
	}
	if in.Debugging != nil {
		in, out := &in.Debugging, &out.Debugging
		*out = new(configv1alpha1.DebuggingConfiguration)
		if err := configv1alpha1.Convert_config_DebuggingConfiguration_To_v1alpha1_DebuggingConfiguration(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Debugging = nil
	}
	return nil
}

// Convert_config_AdmissionControllerConfiguration_To_v1alpha1_AdmissionControllerConfiguration is an autogenerated conversion function.
func Convert_config_AdmissionControllerConfiguration_To_v1alpha1_AdmissionControllerConfiguration(in *config.AdmissionControllerConfiguration, out *AdmissionControllerConfiguration, s conversion.Scope) error {
	return autoConvert_config_AdmissionControllerConfiguration_To_v1alpha1_AdmissionControllerConfiguration(in, out, s)
}

func autoConvert_v1alpha1_HTTPSServer_To_config_HTTPSServer(in *HTTPSServer, out *config.HTTPSServer, s conversion.Scope) error {
	if err := Convert_v1alpha1_Server_To_config_Server(&in.Server, &out.Server, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_TLSServer_To_config_TLSServer(&in.TLS, &out.TLS, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_HTTPSServer_To_config_HTTPSServer is an autogenerated conversion function.
func Convert_v1alpha1_HTTPSServer_To_config_HTTPSServer(in *HTTPSServer, out *config.HTTPSServer, s conversion.Scope) error {
	return autoConvert_v1alpha1_HTTPSServer_To_config_HTTPSServer(in, out, s)
}

func autoConvert_config_HTTPSServer_To_v1alpha1_HTTPSServer(in *config.HTTPSServer, out *HTTPSServer, s conversion.Scope) error {
	if err := Convert_config_Server_To_v1alpha1_Server(&in.Server, &out.Server, s); err != nil {
		return err
	}
	if err := Convert_config_TLSServer_To_v1alpha1_TLSServer(&in.TLS, &out.TLS, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_HTTPSServer_To_v1alpha1_HTTPSServer is an autogenerated conversion function.
func Convert_config_HTTPSServer_To_v1alpha1_HTTPSServer(in *config.HTTPSServer, out *HTTPSServer, s conversion.Scope) error {
	return autoConvert_config_HTTPSServer_To_v1alpha1_HTTPSServer(in, out, s)
}

func autoConvert_v1alpha1_ResourceAdmissionConfiguration_To_config_ResourceAdmissionConfiguration(in *ResourceAdmissionConfiguration, out *config.ResourceAdmissionConfiguration, s conversion.Scope) error {
	out.Limits = *(*[]config.ResourceLimit)(unsafe.Pointer(&in.Limits))
	out.UnrestrictedSubjects = *(*[]v1.Subject)(unsafe.Pointer(&in.UnrestrictedSubjects))
	out.OperationMode = (*config.ResourceAdmissionWebhookMode)(unsafe.Pointer(in.OperationMode))
	return nil
}

// Convert_v1alpha1_ResourceAdmissionConfiguration_To_config_ResourceAdmissionConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_ResourceAdmissionConfiguration_To_config_ResourceAdmissionConfiguration(in *ResourceAdmissionConfiguration, out *config.ResourceAdmissionConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_ResourceAdmissionConfiguration_To_config_ResourceAdmissionConfiguration(in, out, s)
}

func autoConvert_config_ResourceAdmissionConfiguration_To_v1alpha1_ResourceAdmissionConfiguration(in *config.ResourceAdmissionConfiguration, out *ResourceAdmissionConfiguration, s conversion.Scope) error {
	out.Limits = *(*[]ResourceLimit)(unsafe.Pointer(&in.Limits))
	out.UnrestrictedSubjects = *(*[]v1.Subject)(unsafe.Pointer(&in.UnrestrictedSubjects))
	out.OperationMode = (*ResourceAdmissionWebhookMode)(unsafe.Pointer(in.OperationMode))
	return nil
}

// Convert_config_ResourceAdmissionConfiguration_To_v1alpha1_ResourceAdmissionConfiguration is an autogenerated conversion function.
func Convert_config_ResourceAdmissionConfiguration_To_v1alpha1_ResourceAdmissionConfiguration(in *config.ResourceAdmissionConfiguration, out *ResourceAdmissionConfiguration, s conversion.Scope) error {
	return autoConvert_config_ResourceAdmissionConfiguration_To_v1alpha1_ResourceAdmissionConfiguration(in, out, s)
}

func autoConvert_v1alpha1_ResourceLimit_To_config_ResourceLimit(in *ResourceLimit, out *config.ResourceLimit, s conversion.Scope) error {
	out.APIGroups = *(*[]string)(unsafe.Pointer(&in.APIGroups))
	out.APIVersions = *(*[]string)(unsafe.Pointer(&in.APIVersions))
	out.Resources = *(*[]string)(unsafe.Pointer(&in.Resources))
	out.Size = in.Size
	return nil
}

// Convert_v1alpha1_ResourceLimit_To_config_ResourceLimit is an autogenerated conversion function.
func Convert_v1alpha1_ResourceLimit_To_config_ResourceLimit(in *ResourceLimit, out *config.ResourceLimit, s conversion.Scope) error {
	return autoConvert_v1alpha1_ResourceLimit_To_config_ResourceLimit(in, out, s)
}

func autoConvert_config_ResourceLimit_To_v1alpha1_ResourceLimit(in *config.ResourceLimit, out *ResourceLimit, s conversion.Scope) error {
	out.APIGroups = *(*[]string)(unsafe.Pointer(&in.APIGroups))
	out.APIVersions = *(*[]string)(unsafe.Pointer(&in.APIVersions))
	out.Resources = *(*[]string)(unsafe.Pointer(&in.Resources))
	out.Size = in.Size
	return nil
}

// Convert_config_ResourceLimit_To_v1alpha1_ResourceLimit is an autogenerated conversion function.
func Convert_config_ResourceLimit_To_v1alpha1_ResourceLimit(in *config.ResourceLimit, out *ResourceLimit, s conversion.Scope) error {
	return autoConvert_config_ResourceLimit_To_v1alpha1_ResourceLimit(in, out, s)
}

func autoConvert_v1alpha1_Server_To_config_Server(in *Server, out *config.Server, s conversion.Scope) error {
	out.BindAddress = in.BindAddress
	out.Port = in.Port
	return nil
}

// Convert_v1alpha1_Server_To_config_Server is an autogenerated conversion function.
func Convert_v1alpha1_Server_To_config_Server(in *Server, out *config.Server, s conversion.Scope) error {
	return autoConvert_v1alpha1_Server_To_config_Server(in, out, s)
}

func autoConvert_config_Server_To_v1alpha1_Server(in *config.Server, out *Server, s conversion.Scope) error {
	out.BindAddress = in.BindAddress
	out.Port = in.Port
	return nil
}

// Convert_config_Server_To_v1alpha1_Server is an autogenerated conversion function.
func Convert_config_Server_To_v1alpha1_Server(in *config.Server, out *Server, s conversion.Scope) error {
	return autoConvert_config_Server_To_v1alpha1_Server(in, out, s)
}

func autoConvert_v1alpha1_ServerConfiguration_To_config_ServerConfiguration(in *ServerConfiguration, out *config.ServerConfiguration, s conversion.Scope) error {
	if err := Convert_v1alpha1_HTTPSServer_To_config_HTTPSServer(&in.HTTPS, &out.HTTPS, s); err != nil {
		return err
	}
	out.HealthProbes = (*config.Server)(unsafe.Pointer(in.HealthProbes))
	out.Metrics = (*config.Server)(unsafe.Pointer(in.Metrics))
	out.ResourceAdmissionConfiguration = (*config.ResourceAdmissionConfiguration)(unsafe.Pointer(in.ResourceAdmissionConfiguration))
	out.EnableDebugHandlers = (*bool)(unsafe.Pointer(in.EnableDebugHandlers))
	return nil
}

// Convert_v1alpha1_ServerConfiguration_To_config_ServerConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_ServerConfiguration_To_config_ServerConfiguration(in *ServerConfiguration, out *config.ServerConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_ServerConfiguration_To_config_ServerConfiguration(in, out, s)
}

func autoConvert_config_ServerConfiguration_To_v1alpha1_ServerConfiguration(in *config.ServerConfiguration, out *ServerConfiguration, s conversion.Scope) error {
	if err := Convert_config_HTTPSServer_To_v1alpha1_HTTPSServer(&in.HTTPS, &out.HTTPS, s); err != nil {
		return err
	}
	out.HealthProbes = (*Server)(unsafe.Pointer(in.HealthProbes))
	out.Metrics = (*Server)(unsafe.Pointer(in.Metrics))
	out.ResourceAdmissionConfiguration = (*ResourceAdmissionConfiguration)(unsafe.Pointer(in.ResourceAdmissionConfiguration))
	out.EnableDebugHandlers = (*bool)(unsafe.Pointer(in.EnableDebugHandlers))
	return nil
}

// Convert_config_ServerConfiguration_To_v1alpha1_ServerConfiguration is an autogenerated conversion function.
func Convert_config_ServerConfiguration_To_v1alpha1_ServerConfiguration(in *config.ServerConfiguration, out *ServerConfiguration, s conversion.Scope) error {
	return autoConvert_config_ServerConfiguration_To_v1alpha1_ServerConfiguration(in, out, s)
}

func autoConvert_v1alpha1_TLSServer_To_config_TLSServer(in *TLSServer, out *config.TLSServer, s conversion.Scope) error {
	out.ServerCertDir = in.ServerCertDir
	return nil
}

// Convert_v1alpha1_TLSServer_To_config_TLSServer is an autogenerated conversion function.
func Convert_v1alpha1_TLSServer_To_config_TLSServer(in *TLSServer, out *config.TLSServer, s conversion.Scope) error {
	return autoConvert_v1alpha1_TLSServer_To_config_TLSServer(in, out, s)
}

func autoConvert_config_TLSServer_To_v1alpha1_TLSServer(in *config.TLSServer, out *TLSServer, s conversion.Scope) error {
	out.ServerCertDir = in.ServerCertDir
	return nil
}

// Convert_config_TLSServer_To_v1alpha1_TLSServer is an autogenerated conversion function.
func Convert_config_TLSServer_To_v1alpha1_TLSServer(in *config.TLSServer, out *TLSServer, s conversion.Scope) error {
	return autoConvert_config_TLSServer_To_v1alpha1_TLSServer(in, out, s)
}

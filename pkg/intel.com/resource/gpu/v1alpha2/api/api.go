/*
 * Copyright (c) 2023, Intel Corporation.  All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package api

import (
	icrd "github.com/intel/intel-resource-drivers-for-kubernetes/pkg/intel.com/resource/gpu/v1alpha2"
)

const (
	// K8s API group name.
	APIGroupName = icrd.APIGroupName
	// K8s API group version.
	APIVersion = icrd.APIVersion
	// PLain HW GPU to be allocated.
	GpuDeviceType = icrd.GpuDeviceType
	// SR-IOV Virtual Function to be allocated.
	VfDeviceType = icrd.VfDeviceType
	// Either GPU or VF to be allocated.
	AnyDeviceType = icrd.AnyDeviceType
	// Unsupported device type.
	UnknownDeviceType = icrd.UnknownDeviceType
	// K8s API Kind string for GpuClaimParameters.
	GpuClaimParametersKind = "GpuClaimParameters"
	// K8s API Kind string for GpuClassParameters.
	GpuClassParametersKind = "GpuClassParameters"
)

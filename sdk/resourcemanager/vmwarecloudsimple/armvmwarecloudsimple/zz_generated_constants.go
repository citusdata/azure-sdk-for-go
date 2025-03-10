//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armvmwarecloudsimple

const (
	moduleName    = "armvmwarecloudsimple"
	moduleVersion = "v0.2.0"
)

// AggregationType - Metric's aggregation type for e.g. (Average, Total)
type AggregationType string

const (
	AggregationTypeAverage AggregationType = "Average"
	AggregationTypeTotal   AggregationType = "Total"
)

// PossibleAggregationTypeValues returns the possible values for the AggregationType const type.
func PossibleAggregationTypeValues() []AggregationType {
	return []AggregationType{
		AggregationTypeAverage,
		AggregationTypeTotal,
	}
}

// ToPtr returns a *AggregationType pointing to the current value.
func (c AggregationType) ToPtr() *AggregationType {
	return &c
}

// CustomizationHostNameType - Type of host name
type CustomizationHostNameType string

const (
	CustomizationHostNameTypeCUSTOMNAME         CustomizationHostNameType = "CUSTOM_NAME"
	CustomizationHostNameTypeFIXED              CustomizationHostNameType = "FIXED"
	CustomizationHostNameTypePREFIXBASED        CustomizationHostNameType = "PREFIX_BASED"
	CustomizationHostNameTypeUSERDEFINED        CustomizationHostNameType = "USER_DEFINED"
	CustomizationHostNameTypeVIRTUALMACHINENAME CustomizationHostNameType = "VIRTUAL_MACHINE_NAME"
)

// PossibleCustomizationHostNameTypeValues returns the possible values for the CustomizationHostNameType const type.
func PossibleCustomizationHostNameTypeValues() []CustomizationHostNameType {
	return []CustomizationHostNameType{
		CustomizationHostNameTypeCUSTOMNAME,
		CustomizationHostNameTypeFIXED,
		CustomizationHostNameTypePREFIXBASED,
		CustomizationHostNameTypeUSERDEFINED,
		CustomizationHostNameTypeVIRTUALMACHINENAME,
	}
}

// ToPtr returns a *CustomizationHostNameType pointing to the current value.
func (c CustomizationHostNameType) ToPtr() *CustomizationHostNameType {
	return &c
}

// CustomizationIPAddressType - Customization Specification ip type
type CustomizationIPAddressType string

const (
	CustomizationIPAddressTypeCUSTOM      CustomizationIPAddressType = "CUSTOM"
	CustomizationIPAddressTypeDHCPIP      CustomizationIPAddressType = "DHCP_IP"
	CustomizationIPAddressTypeFIXEDIP     CustomizationIPAddressType = "FIXED_IP"
	CustomizationIPAddressTypeUSERDEFINED CustomizationIPAddressType = "USER_DEFINED"
)

// PossibleCustomizationIPAddressTypeValues returns the possible values for the CustomizationIPAddressType const type.
func PossibleCustomizationIPAddressTypeValues() []CustomizationIPAddressType {
	return []CustomizationIPAddressType{
		CustomizationIPAddressTypeCUSTOM,
		CustomizationIPAddressTypeDHCPIP,
		CustomizationIPAddressTypeFIXEDIP,
		CustomizationIPAddressTypeUSERDEFINED,
	}
}

// ToPtr returns a *CustomizationIPAddressType pointing to the current value.
func (c CustomizationIPAddressType) ToPtr() *CustomizationIPAddressType {
	return &c
}

// CustomizationIdentityType - Identity type
type CustomizationIdentityType string

const (
	CustomizationIdentityTypeLINUX       CustomizationIdentityType = "LINUX"
	CustomizationIdentityTypeWINDOWS     CustomizationIdentityType = "WINDOWS"
	CustomizationIdentityTypeWINDOWSTEXT CustomizationIdentityType = "WINDOWS_TEXT"
)

// PossibleCustomizationIdentityTypeValues returns the possible values for the CustomizationIdentityType const type.
func PossibleCustomizationIdentityTypeValues() []CustomizationIdentityType {
	return []CustomizationIdentityType{
		CustomizationIdentityTypeLINUX,
		CustomizationIdentityTypeWINDOWS,
		CustomizationIdentityTypeWINDOWSTEXT,
	}
}

// ToPtr returns a *CustomizationIdentityType pointing to the current value.
func (c CustomizationIdentityType) ToPtr() *CustomizationIdentityType {
	return &c
}

// CustomizationPolicyPropertiesType - The type of customization (Linux or Windows)
type CustomizationPolicyPropertiesType string

const (
	CustomizationPolicyPropertiesTypeLINUX   CustomizationPolicyPropertiesType = "LINUX"
	CustomizationPolicyPropertiesTypeWINDOWS CustomizationPolicyPropertiesType = "WINDOWS"
)

// PossibleCustomizationPolicyPropertiesTypeValues returns the possible values for the CustomizationPolicyPropertiesType const type.
func PossibleCustomizationPolicyPropertiesTypeValues() []CustomizationPolicyPropertiesType {
	return []CustomizationPolicyPropertiesType{
		CustomizationPolicyPropertiesTypeLINUX,
		CustomizationPolicyPropertiesTypeWINDOWS,
	}
}

// ToPtr returns a *CustomizationPolicyPropertiesType pointing to the current value.
func (c CustomizationPolicyPropertiesType) ToPtr() *CustomizationPolicyPropertiesType {
	return &c
}

// DiskIndependenceMode - Disk's independence mode type
type DiskIndependenceMode string

const (
	DiskIndependenceModePersistent               DiskIndependenceMode = "persistent"
	DiskIndependenceModeIndependentPersistent    DiskIndependenceMode = "independent_persistent"
	DiskIndependenceModeIndependentNonpersistent DiskIndependenceMode = "independent_nonpersistent"
)

// PossibleDiskIndependenceModeValues returns the possible values for the DiskIndependenceMode const type.
func PossibleDiskIndependenceModeValues() []DiskIndependenceMode {
	return []DiskIndependenceMode{
		DiskIndependenceModePersistent,
		DiskIndependenceModeIndependentPersistent,
		DiskIndependenceModeIndependentNonpersistent,
	}
}

// ToPtr returns a *DiskIndependenceMode pointing to the current value.
func (c DiskIndependenceMode) ToPtr() *DiskIndependenceMode {
	return &c
}

// GuestOSNICCustomizationAllocation - IP address allocation method
type GuestOSNICCustomizationAllocation string

const (
	GuestOSNICCustomizationAllocationDynamic GuestOSNICCustomizationAllocation = "dynamic"
	GuestOSNICCustomizationAllocationStatic  GuestOSNICCustomizationAllocation = "static"
)

// PossibleGuestOSNICCustomizationAllocationValues returns the possible values for the GuestOSNICCustomizationAllocation const type.
func PossibleGuestOSNICCustomizationAllocationValues() []GuestOSNICCustomizationAllocation {
	return []GuestOSNICCustomizationAllocation{
		GuestOSNICCustomizationAllocationDynamic,
		GuestOSNICCustomizationAllocationStatic,
	}
}

// ToPtr returns a *GuestOSNICCustomizationAllocation pointing to the current value.
func (c GuestOSNICCustomizationAllocation) ToPtr() *GuestOSNICCustomizationAllocation {
	return &c
}

// GuestOSType - The Guest OS type
type GuestOSType string

const (
	GuestOSTypeLinux   GuestOSType = "linux"
	GuestOSTypeWindows GuestOSType = "windows"
	GuestOSTypeOther   GuestOSType = "other"
)

// PossibleGuestOSTypeValues returns the possible values for the GuestOSType const type.
func PossibleGuestOSTypeValues() []GuestOSType {
	return []GuestOSType{
		GuestOSTypeLinux,
		GuestOSTypeWindows,
		GuestOSTypeOther,
	}
}

// ToPtr returns a *GuestOSType pointing to the current value.
func (c GuestOSType) ToPtr() *GuestOSType {
	return &c
}

// NICType - NIC type
type NICType string

const (
	NICTypeE1000   NICType = "E1000"
	NICTypeE1000E  NICType = "E1000E"
	NICTypePCNET32 NICType = "PCNET32"
	NICTypeVMXNET  NICType = "VMXNET"
	NICTypeVMXNET2 NICType = "VMXNET2"
	NICTypeVMXNET3 NICType = "VMXNET3"
)

// PossibleNICTypeValues returns the possible values for the NICType const type.
func PossibleNICTypeValues() []NICType {
	return []NICType{
		NICTypeE1000,
		NICTypeE1000E,
		NICTypePCNET32,
		NICTypeVMXNET,
		NICTypeVMXNET2,
		NICTypeVMXNET3,
	}
}

// ToPtr returns a *NICType pointing to the current value.
func (c NICType) ToPtr() *NICType {
	return &c
}

// NodeStatus - Node status, indicates is private cloud set up on this node or not
type NodeStatus string

const (
	NodeStatusUnused NodeStatus = "unused"
	NodeStatusUsed   NodeStatus = "used"
)

// PossibleNodeStatusValues returns the possible values for the NodeStatus const type.
func PossibleNodeStatusValues() []NodeStatus {
	return []NodeStatus{
		NodeStatusUnused,
		NodeStatusUsed,
	}
}

// ToPtr returns a *NodeStatus pointing to the current value.
func (c NodeStatus) ToPtr() *NodeStatus {
	return &c
}

// OnboardingStatus - indicates whether account onboarded or not in a given region
type OnboardingStatus string

const (
	OnboardingStatusNotOnBoarded     OnboardingStatus = "notOnBoarded"
	OnboardingStatusOnBoarded        OnboardingStatus = "onBoarded"
	OnboardingStatusOnBoardingFailed OnboardingStatus = "onBoardingFailed"
	OnboardingStatusOnBoarding       OnboardingStatus = "onBoarding"
)

// PossibleOnboardingStatusValues returns the possible values for the OnboardingStatus const type.
func PossibleOnboardingStatusValues() []OnboardingStatus {
	return []OnboardingStatus{
		OnboardingStatusNotOnBoarded,
		OnboardingStatusOnBoarded,
		OnboardingStatusOnBoardingFailed,
		OnboardingStatusOnBoarding,
	}
}

// ToPtr returns a *OnboardingStatus pointing to the current value.
func (c OnboardingStatus) ToPtr() *OnboardingStatus {
	return &c
}

// OperationOrigin - The origin of operation
type OperationOrigin string

const (
	OperationOriginUser       OperationOrigin = "user"
	OperationOriginSystem     OperationOrigin = "system"
	OperationOriginUserSystem OperationOrigin = "user,system"
)

// PossibleOperationOriginValues returns the possible values for the OperationOrigin const type.
func PossibleOperationOriginValues() []OperationOrigin {
	return []OperationOrigin{
		OperationOriginUser,
		OperationOriginSystem,
		OperationOriginUserSystem,
	}
}

// ToPtr returns a *OperationOrigin pointing to the current value.
func (c OperationOrigin) ToPtr() *OperationOrigin {
	return &c
}

// StopMode - mode indicates a type of stop operation - reboot, suspend, shutdown or power-off
type StopMode string

const (
	StopModeReboot   StopMode = "reboot"
	StopModeSuspend  StopMode = "suspend"
	StopModeShutdown StopMode = "shutdown"
	StopModePoweroff StopMode = "poweroff"
)

// PossibleStopModeValues returns the possible values for the StopMode const type.
func PossibleStopModeValues() []StopMode {
	return []StopMode{
		StopModeReboot,
		StopModeSuspend,
		StopModeShutdown,
		StopModePoweroff,
	}
}

// ToPtr returns a *StopMode pointing to the current value.
func (c StopMode) ToPtr() *StopMode {
	return &c
}

// UsageCount - The usages' unit
type UsageCount string

const (
	UsageCountCount          UsageCount = "Count"
	UsageCountBytes          UsageCount = "Bytes"
	UsageCountSeconds        UsageCount = "Seconds"
	UsageCountPercent        UsageCount = "Percent"
	UsageCountCountPerSecond UsageCount = "CountPerSecond"
	UsageCountBytesPerSecond UsageCount = "BytesPerSecond"
)

// PossibleUsageCountValues returns the possible values for the UsageCount const type.
func PossibleUsageCountValues() []UsageCount {
	return []UsageCount{
		UsageCountCount,
		UsageCountBytes,
		UsageCountSeconds,
		UsageCountPercent,
		UsageCountCountPerSecond,
		UsageCountBytesPerSecond,
	}
}

// ToPtr returns a *UsageCount pointing to the current value.
func (c UsageCount) ToPtr() *UsageCount {
	return &c
}

// VirtualMachineStatus - The status of Virtual machine
type VirtualMachineStatus string

const (
	VirtualMachineStatusRunning      VirtualMachineStatus = "running"
	VirtualMachineStatusSuspended    VirtualMachineStatus = "suspended"
	VirtualMachineStatusPoweredoff   VirtualMachineStatus = "poweredoff"
	VirtualMachineStatusUpdating     VirtualMachineStatus = "updating"
	VirtualMachineStatusDeallocating VirtualMachineStatus = "deallocating"
	VirtualMachineStatusDeleting     VirtualMachineStatus = "deleting"
)

// PossibleVirtualMachineStatusValues returns the possible values for the VirtualMachineStatus const type.
func PossibleVirtualMachineStatusValues() []VirtualMachineStatus {
	return []VirtualMachineStatus{
		VirtualMachineStatusRunning,
		VirtualMachineStatusSuspended,
		VirtualMachineStatusPoweredoff,
		VirtualMachineStatusUpdating,
		VirtualMachineStatusDeallocating,
		VirtualMachineStatusDeleting,
	}
}

// ToPtr returns a *VirtualMachineStatus pointing to the current value.
func (c VirtualMachineStatus) ToPtr() *VirtualMachineStatus {
	return &c
}

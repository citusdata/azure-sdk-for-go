//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armadvisor

const (
	moduleName    = "armadvisor"
	moduleVersion = "v0.2.0"
)

// CPUThreshold - Minimum percentage threshold for Advisor low CPU utilization evaluation. Valid only for subscriptions. Valid
// values: 5 (default), 10, 15 or 20.
type CPUThreshold string

const (
	CPUThresholdFifteen CPUThreshold = "15"
	CPUThresholdFive    CPUThreshold = "5"
	CPUThresholdTen     CPUThreshold = "10"
	CPUThresholdTwenty  CPUThreshold = "20"
)

// PossibleCPUThresholdValues returns the possible values for the CPUThreshold const type.
func PossibleCPUThresholdValues() []CPUThreshold {
	return []CPUThreshold{
		CPUThresholdFifteen,
		CPUThresholdFive,
		CPUThresholdTen,
		CPUThresholdTwenty,
	}
}

// ToPtr returns a *CPUThreshold pointing to the current value.
func (c CPUThreshold) ToPtr() *CPUThreshold {
	return &c
}

type Category string

const (
	CategoryCost                  Category = "Cost"
	CategoryHighAvailability      Category = "HighAvailability"
	CategoryOperationalExcellence Category = "OperationalExcellence"
	CategoryPerformance           Category = "Performance"
	CategorySecurity              Category = "Security"
)

// PossibleCategoryValues returns the possible values for the Category const type.
func PossibleCategoryValues() []Category {
	return []Category{
		CategoryCost,
		CategoryHighAvailability,
		CategoryOperationalExcellence,
		CategoryPerformance,
		CategorySecurity,
	}
}

// ToPtr returns a *Category pointing to the current value.
func (c Category) ToPtr() *Category {
	return &c
}

type ConfigurationName string

const (
	ConfigurationNameDefault ConfigurationName = "default"
)

// PossibleConfigurationNameValues returns the possible values for the ConfigurationName const type.
func PossibleConfigurationNameValues() []ConfigurationName {
	return []ConfigurationName{
		ConfigurationNameDefault,
	}
}

// ToPtr returns a *ConfigurationName pointing to the current value.
func (c ConfigurationName) ToPtr() *ConfigurationName {
	return &c
}

// DigestConfigState - State of digest configuration.
type DigestConfigState string

const (
	DigestConfigStateActive   DigestConfigState = "Active"
	DigestConfigStateDisabled DigestConfigState = "Disabled"
)

// PossibleDigestConfigStateValues returns the possible values for the DigestConfigState const type.
func PossibleDigestConfigStateValues() []DigestConfigState {
	return []DigestConfigState{
		DigestConfigStateActive,
		DigestConfigStateDisabled,
	}
}

// ToPtr returns a *DigestConfigState pointing to the current value.
func (c DigestConfigState) ToPtr() *DigestConfigState {
	return &c
}

// Impact - The business impact of the recommendation.
type Impact string

const (
	ImpactHigh   Impact = "High"
	ImpactLow    Impact = "Low"
	ImpactMedium Impact = "Medium"
)

// PossibleImpactValues returns the possible values for the Impact const type.
func PossibleImpactValues() []Impact {
	return []Impact{
		ImpactHigh,
		ImpactLow,
		ImpactMedium,
	}
}

// ToPtr returns a *Impact pointing to the current value.
func (c Impact) ToPtr() *Impact {
	return &c
}

// Risk - The potential risk of not implementing the recommendation.
type Risk string

const (
	RiskError   Risk = "Error"
	RiskNone    Risk = "None"
	RiskWarning Risk = "Warning"
)

// PossibleRiskValues returns the possible values for the Risk const type.
func PossibleRiskValues() []Risk {
	return []Risk{
		RiskError,
		RiskNone,
		RiskWarning,
	}
}

// ToPtr returns a *Risk pointing to the current value.
func (c Risk) ToPtr() *Risk {
	return &c
}

type Scenario string

const (
	ScenarioAlerts Scenario = "Alerts"
)

// PossibleScenarioValues returns the possible values for the Scenario const type.
func PossibleScenarioValues() []Scenario {
	return []Scenario{
		ScenarioAlerts,
	}
}

// ToPtr returns a *Scenario pointing to the current value.
func (c Scenario) ToPtr() *Scenario {
	return &c
}

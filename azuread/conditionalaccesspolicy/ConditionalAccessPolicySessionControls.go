// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package conditionalaccesspolicy


type ConditionalAccessPolicySessionControls struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.42.0/docs/resources/conditional_access_policy#application_enforced_restrictions_enabled ConditionalAccessPolicy#application_enforced_restrictions_enabled}.
	ApplicationEnforcedRestrictionsEnabled interface{} `field:"optional" json:"applicationEnforcedRestrictionsEnabled" yaml:"applicationEnforcedRestrictionsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.42.0/docs/resources/conditional_access_policy#cloud_app_security_policy ConditionalAccessPolicy#cloud_app_security_policy}.
	CloudAppSecurityPolicy *string `field:"optional" json:"cloudAppSecurityPolicy" yaml:"cloudAppSecurityPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.42.0/docs/resources/conditional_access_policy#disable_resilience_defaults ConditionalAccessPolicy#disable_resilience_defaults}.
	DisableResilienceDefaults interface{} `field:"optional" json:"disableResilienceDefaults" yaml:"disableResilienceDefaults"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.42.0/docs/resources/conditional_access_policy#persistent_browser_mode ConditionalAccessPolicy#persistent_browser_mode}.
	PersistentBrowserMode *string `field:"optional" json:"persistentBrowserMode" yaml:"persistentBrowserMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.42.0/docs/resources/conditional_access_policy#sign_in_frequency ConditionalAccessPolicy#sign_in_frequency}.
	SignInFrequency *float64 `field:"optional" json:"signInFrequency" yaml:"signInFrequency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.42.0/docs/resources/conditional_access_policy#sign_in_frequency_period ConditionalAccessPolicy#sign_in_frequency_period}.
	SignInFrequencyPeriod *string `field:"optional" json:"signInFrequencyPeriod" yaml:"signInFrequencyPeriod"`
}


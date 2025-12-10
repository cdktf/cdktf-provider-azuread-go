// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package conditionalaccesspolicy


type ConditionalAccessPolicyConditionsApplications struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.7.0/docs/resources/conditional_access_policy#excluded_applications ConditionalAccessPolicy#excluded_applications}.
	ExcludedApplications *[]*string `field:"optional" json:"excludedApplications" yaml:"excludedApplications"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.7.0/docs/resources/conditional_access_policy#included_applications ConditionalAccessPolicy#included_applications}.
	IncludedApplications *[]*string `field:"optional" json:"includedApplications" yaml:"includedApplications"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.7.0/docs/resources/conditional_access_policy#included_user_actions ConditionalAccessPolicy#included_user_actions}.
	IncludedUserActions *[]*string `field:"optional" json:"includedUserActions" yaml:"includedUserActions"`
}


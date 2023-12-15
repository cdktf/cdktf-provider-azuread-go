// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package conditionalaccesspolicy


type ConditionalAccessPolicyConditionsPlatforms struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.47.0/docs/resources/conditional_access_policy#included_platforms ConditionalAccessPolicy#included_platforms}.
	IncludedPlatforms *[]*string `field:"required" json:"includedPlatforms" yaml:"includedPlatforms"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.47.0/docs/resources/conditional_access_policy#excluded_platforms ConditionalAccessPolicy#excluded_platforms}.
	ExcludedPlatforms *[]*string `field:"optional" json:"excludedPlatforms" yaml:"excludedPlatforms"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package conditionalaccesspolicy


type ConditionalAccessPolicyConditionsLocations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.6.0/docs/resources/conditional_access_policy#included_locations ConditionalAccessPolicy#included_locations}.
	IncludedLocations *[]*string `field:"required" json:"includedLocations" yaml:"includedLocations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.6.0/docs/resources/conditional_access_policy#excluded_locations ConditionalAccessPolicy#excluded_locations}.
	ExcludedLocations *[]*string `field:"optional" json:"excludedLocations" yaml:"excludedLocations"`
}


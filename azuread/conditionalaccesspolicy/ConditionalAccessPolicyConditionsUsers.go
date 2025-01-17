// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package conditionalaccesspolicy


type ConditionalAccessPolicyConditionsUsers struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.1.0/docs/resources/conditional_access_policy#excluded_groups ConditionalAccessPolicy#excluded_groups}.
	ExcludedGroups *[]*string `field:"optional" json:"excludedGroups" yaml:"excludedGroups"`
	// excluded_guests_or_external_users block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.1.0/docs/resources/conditional_access_policy#excluded_guests_or_external_users ConditionalAccessPolicy#excluded_guests_or_external_users}
	ExcludedGuestsOrExternalUsers interface{} `field:"optional" json:"excludedGuestsOrExternalUsers" yaml:"excludedGuestsOrExternalUsers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.1.0/docs/resources/conditional_access_policy#excluded_roles ConditionalAccessPolicy#excluded_roles}.
	ExcludedRoles *[]*string `field:"optional" json:"excludedRoles" yaml:"excludedRoles"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.1.0/docs/resources/conditional_access_policy#excluded_users ConditionalAccessPolicy#excluded_users}.
	ExcludedUsers *[]*string `field:"optional" json:"excludedUsers" yaml:"excludedUsers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.1.0/docs/resources/conditional_access_policy#included_groups ConditionalAccessPolicy#included_groups}.
	IncludedGroups *[]*string `field:"optional" json:"includedGroups" yaml:"includedGroups"`
	// included_guests_or_external_users block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.1.0/docs/resources/conditional_access_policy#included_guests_or_external_users ConditionalAccessPolicy#included_guests_or_external_users}
	IncludedGuestsOrExternalUsers interface{} `field:"optional" json:"includedGuestsOrExternalUsers" yaml:"includedGuestsOrExternalUsers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.1.0/docs/resources/conditional_access_policy#included_roles ConditionalAccessPolicy#included_roles}.
	IncludedRoles *[]*string `field:"optional" json:"includedRoles" yaml:"includedRoles"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.1.0/docs/resources/conditional_access_policy#included_users ConditionalAccessPolicy#included_users}.
	IncludedUsers *[]*string `field:"optional" json:"includedUsers" yaml:"includedUsers"`
}


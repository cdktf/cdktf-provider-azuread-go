// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package conditionalaccesspolicy


type ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUsersExternalTenants struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.3.0/docs/resources/conditional_access_policy#membership_kind ConditionalAccessPolicy#membership_kind}.
	MembershipKind *string `field:"required" json:"membershipKind" yaml:"membershipKind"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.3.0/docs/resources/conditional_access_policy#members ConditionalAccessPolicy#members}.
	Members *[]*string `field:"optional" json:"members" yaml:"members"`
}


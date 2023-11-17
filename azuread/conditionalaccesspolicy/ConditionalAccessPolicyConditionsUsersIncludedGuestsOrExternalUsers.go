// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package conditionalaccesspolicy


type ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsers struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.46.0/docs/resources/conditional_access_policy#guest_or_external_user_types ConditionalAccessPolicy#guest_or_external_user_types}.
	GuestOrExternalUserTypes *[]*string `field:"required" json:"guestOrExternalUserTypes" yaml:"guestOrExternalUserTypes"`
	// external_tenants block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.46.0/docs/resources/conditional_access_policy#external_tenants ConditionalAccessPolicy#external_tenants}
	ExternalTenants interface{} `field:"optional" json:"externalTenants" yaml:"externalTenants"`
}


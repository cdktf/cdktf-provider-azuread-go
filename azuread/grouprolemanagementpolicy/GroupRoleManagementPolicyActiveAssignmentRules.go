// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grouprolemanagementpolicy


type GroupRoleManagementPolicyActiveAssignmentRules struct {
	// Must the assignment have an expiry date.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.52.0/docs/resources/group_role_management_policy#expiration_required GroupRoleManagementPolicy#expiration_required}
	ExpirationRequired interface{} `field:"optional" json:"expirationRequired" yaml:"expirationRequired"`
	// The duration after which assignments expire.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.52.0/docs/resources/group_role_management_policy#expire_after GroupRoleManagementPolicy#expire_after}
	ExpireAfter *string `field:"optional" json:"expireAfter" yaml:"expireAfter"`
	// Whether a justification is required to make an assignment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.52.0/docs/resources/group_role_management_policy#require_justification GroupRoleManagementPolicy#require_justification}
	RequireJustification interface{} `field:"optional" json:"requireJustification" yaml:"requireJustification"`
	// Whether multi-factor authentication is required to make an assignment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.52.0/docs/resources/group_role_management_policy#require_multifactor_authentication GroupRoleManagementPolicy#require_multifactor_authentication}
	RequireMultifactorAuthentication interface{} `field:"optional" json:"requireMultifactorAuthentication" yaml:"requireMultifactorAuthentication"`
	// Whether ticket information is required to make an assignment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.52.0/docs/resources/group_role_management_policy#require_ticket_info GroupRoleManagementPolicy#require_ticket_info}
	RequireTicketInfo interface{} `field:"optional" json:"requireTicketInfo" yaml:"requireTicketInfo"`
}


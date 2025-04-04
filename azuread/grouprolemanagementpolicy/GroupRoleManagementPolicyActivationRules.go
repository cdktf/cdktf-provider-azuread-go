// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grouprolemanagementpolicy


type GroupRoleManagementPolicyActivationRules struct {
	// approval_stage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.3.0/docs/resources/group_role_management_policy#approval_stage GroupRoleManagementPolicy#approval_stage}
	ApprovalStage *GroupRoleManagementPolicyActivationRulesApprovalStage `field:"optional" json:"approvalStage" yaml:"approvalStage"`
	// The time after which the an activation can be valid for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.3.0/docs/resources/group_role_management_policy#maximum_duration GroupRoleManagementPolicy#maximum_duration}
	MaximumDuration *string `field:"optional" json:"maximumDuration" yaml:"maximumDuration"`
	// Whether an approval is required for activation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.3.0/docs/resources/group_role_management_policy#require_approval GroupRoleManagementPolicy#require_approval}
	RequireApproval interface{} `field:"optional" json:"requireApproval" yaml:"requireApproval"`
	// Whether a conditional access context is required during activation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.3.0/docs/resources/group_role_management_policy#required_conditional_access_authentication_context GroupRoleManagementPolicy#required_conditional_access_authentication_context}
	RequiredConditionalAccessAuthenticationContext *string `field:"optional" json:"requiredConditionalAccessAuthenticationContext" yaml:"requiredConditionalAccessAuthenticationContext"`
	// Whether a justification is required during activation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.3.0/docs/resources/group_role_management_policy#require_justification GroupRoleManagementPolicy#require_justification}
	RequireJustification interface{} `field:"optional" json:"requireJustification" yaml:"requireJustification"`
	// Whether multi-factor authentication is required during activation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.3.0/docs/resources/group_role_management_policy#require_multifactor_authentication GroupRoleManagementPolicy#require_multifactor_authentication}
	RequireMultifactorAuthentication interface{} `field:"optional" json:"requireMultifactorAuthentication" yaml:"requireMultifactorAuthentication"`
	// Whether ticket information is required during activation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.3.0/docs/resources/group_role_management_policy#require_ticket_info GroupRoleManagementPolicy#require_ticket_info}
	RequireTicketInfo interface{} `field:"optional" json:"requireTicketInfo" yaml:"requireTicketInfo"`
}


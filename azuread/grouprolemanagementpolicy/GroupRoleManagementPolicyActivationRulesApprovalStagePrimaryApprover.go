// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grouprolemanagementpolicy


type GroupRoleManagementPolicyActivationRulesApprovalStagePrimaryApprover struct {
	// The ID of the object to act as an approver.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.53.1/docs/resources/group_role_management_policy#object_id GroupRoleManagementPolicy#object_id}
	ObjectId *string `field:"required" json:"objectId" yaml:"objectId"`
	// The type of object acting as an approver.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.53.1/docs/resources/group_role_management_policy#type GroupRoleManagementPolicy#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}


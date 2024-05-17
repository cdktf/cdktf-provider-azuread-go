// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grouprolemanagementpolicy


type GroupRoleManagementPolicyActivationRulesApprovalStage struct {
	// primary_approver block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.50.0/docs/resources/group_role_management_policy#primary_approver GroupRoleManagementPolicy#primary_approver}
	PrimaryApprover interface{} `field:"required" json:"primaryApprover" yaml:"primaryApprover"`
}


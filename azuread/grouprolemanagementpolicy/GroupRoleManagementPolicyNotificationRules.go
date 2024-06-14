// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grouprolemanagementpolicy


type GroupRoleManagementPolicyNotificationRules struct {
	// active_assignments block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.52.0/docs/resources/group_role_management_policy#active_assignments GroupRoleManagementPolicy#active_assignments}
	ActiveAssignments *GroupRoleManagementPolicyNotificationRulesActiveAssignments `field:"optional" json:"activeAssignments" yaml:"activeAssignments"`
	// eligible_activations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.52.0/docs/resources/group_role_management_policy#eligible_activations GroupRoleManagementPolicy#eligible_activations}
	EligibleActivations *GroupRoleManagementPolicyNotificationRulesEligibleActivations `field:"optional" json:"eligibleActivations" yaml:"eligibleActivations"`
	// eligible_assignments block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.52.0/docs/resources/group_role_management_policy#eligible_assignments GroupRoleManagementPolicy#eligible_assignments}
	EligibleAssignments *GroupRoleManagementPolicyNotificationRulesEligibleAssignments `field:"optional" json:"eligibleAssignments" yaml:"eligibleAssignments"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grouprolemanagementpolicy


type GroupRoleManagementPolicyNotificationRulesEligibleAssignments struct {
	// admin_notifications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.2.0/docs/resources/group_role_management_policy#admin_notifications GroupRoleManagementPolicy#admin_notifications}
	AdminNotifications *GroupRoleManagementPolicyNotificationRulesEligibleAssignmentsAdminNotifications `field:"optional" json:"adminNotifications" yaml:"adminNotifications"`
	// approver_notifications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.2.0/docs/resources/group_role_management_policy#approver_notifications GroupRoleManagementPolicy#approver_notifications}
	ApproverNotifications *GroupRoleManagementPolicyNotificationRulesEligibleAssignmentsApproverNotifications `field:"optional" json:"approverNotifications" yaml:"approverNotifications"`
	// assignee_notifications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.2.0/docs/resources/group_role_management_policy#assignee_notifications GroupRoleManagementPolicy#assignee_notifications}
	AssigneeNotifications *GroupRoleManagementPolicyNotificationRulesEligibleAssignmentsAssigneeNotifications `field:"optional" json:"assigneeNotifications" yaml:"assigneeNotifications"`
}


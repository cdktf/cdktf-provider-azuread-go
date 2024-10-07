// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grouprolemanagementpolicy


type GroupRoleManagementPolicyNotificationRulesActiveAssignments struct {
	// admin_notifications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.0.2/docs/resources/group_role_management_policy#admin_notifications GroupRoleManagementPolicy#admin_notifications}
	AdminNotifications *GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAdminNotifications `field:"optional" json:"adminNotifications" yaml:"adminNotifications"`
	// approver_notifications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.0.2/docs/resources/group_role_management_policy#approver_notifications GroupRoleManagementPolicy#approver_notifications}
	ApproverNotifications *GroupRoleManagementPolicyNotificationRulesActiveAssignmentsApproverNotifications `field:"optional" json:"approverNotifications" yaml:"approverNotifications"`
	// assignee_notifications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.0.2/docs/resources/group_role_management_policy#assignee_notifications GroupRoleManagementPolicy#assignee_notifications}
	AssigneeNotifications *GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotifications `field:"optional" json:"assigneeNotifications" yaml:"assigneeNotifications"`
}


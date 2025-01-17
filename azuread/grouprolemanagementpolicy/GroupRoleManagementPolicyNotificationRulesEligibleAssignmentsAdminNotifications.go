// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grouprolemanagementpolicy


type GroupRoleManagementPolicyNotificationRulesEligibleAssignmentsAdminNotifications struct {
	// Whether the default recipients are notified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.1.0/docs/resources/group_role_management_policy#default_recipients GroupRoleManagementPolicy#default_recipients}
	DefaultRecipients interface{} `field:"required" json:"defaultRecipients" yaml:"defaultRecipients"`
	// What level of notifications are sent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.1.0/docs/resources/group_role_management_policy#notification_level GroupRoleManagementPolicy#notification_level}
	NotificationLevel *string `field:"required" json:"notificationLevel" yaml:"notificationLevel"`
	// The additional recipients to notify.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.1.0/docs/resources/group_role_management_policy#additional_recipients GroupRoleManagementPolicy#additional_recipients}
	AdditionalRecipients *[]*string `field:"optional" json:"additionalRecipients" yaml:"additionalRecipients"`
}


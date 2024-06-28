// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grouprolemanagementpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GroupRoleManagementPolicyConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// ID of the group to which this policy is assigned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.53.0/docs/resources/group_role_management_policy#group_id GroupRoleManagementPolicy#group_id}
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
	// The ID of the role of this policy to the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.53.0/docs/resources/group_role_management_policy#role_id GroupRoleManagementPolicy#role_id}
	RoleId *string `field:"required" json:"roleId" yaml:"roleId"`
	// activation_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.53.0/docs/resources/group_role_management_policy#activation_rules GroupRoleManagementPolicy#activation_rules}
	ActivationRules *GroupRoleManagementPolicyActivationRules `field:"optional" json:"activationRules" yaml:"activationRules"`
	// active_assignment_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.53.0/docs/resources/group_role_management_policy#active_assignment_rules GroupRoleManagementPolicy#active_assignment_rules}
	ActiveAssignmentRules *GroupRoleManagementPolicyActiveAssignmentRules `field:"optional" json:"activeAssignmentRules" yaml:"activeAssignmentRules"`
	// eligible_assignment_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.53.0/docs/resources/group_role_management_policy#eligible_assignment_rules GroupRoleManagementPolicy#eligible_assignment_rules}
	EligibleAssignmentRules *GroupRoleManagementPolicyEligibleAssignmentRules `field:"optional" json:"eligibleAssignmentRules" yaml:"eligibleAssignmentRules"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.53.0/docs/resources/group_role_management_policy#id GroupRoleManagementPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// notification_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.53.0/docs/resources/group_role_management_policy#notification_rules GroupRoleManagementPolicy#notification_rules}
	NotificationRules *GroupRoleManagementPolicyNotificationRules `field:"optional" json:"notificationRules" yaml:"notificationRules"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.53.0/docs/resources/group_role_management_policy#timeouts GroupRoleManagementPolicy#timeouts}
	Timeouts *GroupRoleManagementPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grouprolemanagementpolicy


type GroupRoleManagementPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.50.0/docs/resources/group_role_management_policy#create GroupRoleManagementPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.50.0/docs/resources/group_role_management_policy#delete GroupRoleManagementPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.50.0/docs/resources/group_role_management_policy#read GroupRoleManagementPolicy#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.50.0/docs/resources/group_role_management_policy#update GroupRoleManagementPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


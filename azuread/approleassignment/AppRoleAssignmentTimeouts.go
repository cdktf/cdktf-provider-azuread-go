// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package approleassignment


type AppRoleAssignmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.53.1/docs/resources/app_role_assignment#create AppRoleAssignment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.53.1/docs/resources/app_role_assignment#delete AppRoleAssignment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.53.1/docs/resources/app_role_assignment#read AppRoleAssignment#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package directoryroleassignment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DirectoryRoleAssignmentConfig struct {
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
	// The object ID of the member principal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.49.1/docs/resources/directory_role_assignment#principal_object_id DirectoryRoleAssignment#principal_object_id}
	PrincipalObjectId *string `field:"required" json:"principalObjectId" yaml:"principalObjectId"`
	// The object ID of the directory role for this assignment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.49.1/docs/resources/directory_role_assignment#role_id DirectoryRoleAssignment#role_id}
	RoleId *string `field:"required" json:"roleId" yaml:"roleId"`
	// Identifier of the app-specific scope when the assignment scope is app-specific.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.49.1/docs/resources/directory_role_assignment#app_scope_id DirectoryRoleAssignment#app_scope_id}
	AppScopeId *string `field:"optional" json:"appScopeId" yaml:"appScopeId"`
	// Identifier of the app-specific scope when the assignment scope is app-specific.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.49.1/docs/resources/directory_role_assignment#app_scope_object_id DirectoryRoleAssignment#app_scope_object_id}
	AppScopeObjectId *string `field:"optional" json:"appScopeObjectId" yaml:"appScopeObjectId"`
	// Identifier of the directory object representing the scope of the assignment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.49.1/docs/resources/directory_role_assignment#directory_scope_id DirectoryRoleAssignment#directory_scope_id}
	DirectoryScopeId *string `field:"optional" json:"directoryScopeId" yaml:"directoryScopeId"`
	// Identifier of the directory object representing the scope of the assignment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.49.1/docs/resources/directory_role_assignment#directory_scope_object_id DirectoryRoleAssignment#directory_scope_object_id}
	DirectoryScopeObjectId *string `field:"optional" json:"directoryScopeObjectId" yaml:"directoryScopeObjectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.49.1/docs/resources/directory_role_assignment#id DirectoryRoleAssignment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.49.1/docs/resources/directory_role_assignment#timeouts DirectoryRoleAssignment#timeouts}
	Timeouts *DirectoryRoleAssignmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


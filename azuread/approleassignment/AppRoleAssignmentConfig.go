// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package approleassignment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppRoleAssignmentConfig struct {
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
	// The ID of the app role to be assigned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.0/docs/resources/app_role_assignment#app_role_id AppRoleAssignment#app_role_id}
	AppRoleId *string `field:"required" json:"appRoleId" yaml:"appRoleId"`
	// The object ID of the user, group or service principal to be assigned this app role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.0/docs/resources/app_role_assignment#principal_object_id AppRoleAssignment#principal_object_id}
	PrincipalObjectId *string `field:"required" json:"principalObjectId" yaml:"principalObjectId"`
	// The object ID of the service principal representing the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.0/docs/resources/app_role_assignment#resource_object_id AppRoleAssignment#resource_object_id}
	ResourceObjectId *string `field:"required" json:"resourceObjectId" yaml:"resourceObjectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.0/docs/resources/app_role_assignment#id AppRoleAssignment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.0/docs/resources/app_role_assignment#timeouts AppRoleAssignment#timeouts}
	Timeouts *AppRoleAssignmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


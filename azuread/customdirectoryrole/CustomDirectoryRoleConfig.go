// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package customdirectoryrole

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CustomDirectoryRoleConfig struct {
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
	// The display name of the custom directory role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/custom_directory_role#display_name CustomDirectoryRole#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Indicates whether the role is enabled for assignment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/custom_directory_role#enabled CustomDirectoryRole#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// permissions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/custom_directory_role#permissions CustomDirectoryRole#permissions}
	Permissions interface{} `field:"required" json:"permissions" yaml:"permissions"`
	// The version of the role definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/custom_directory_role#version CustomDirectoryRole#version}
	Version *string `field:"required" json:"version" yaml:"version"`
	// The description of the custom directory role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/custom_directory_role#description CustomDirectoryRole#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/custom_directory_role#id CustomDirectoryRole#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Custom template identifier that is typically used if one needs an identifier to be the same across different directories.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/custom_directory_role#template_id CustomDirectoryRole#template_id}
	TemplateId *string `field:"optional" json:"templateId" yaml:"templateId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/custom_directory_role#timeouts CustomDirectoryRole#timeouts}
	Timeouts *CustomDirectoryRoleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


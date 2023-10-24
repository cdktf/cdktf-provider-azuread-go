// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package customdirectoryrole


type CustomDirectoryRoleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/custom_directory_role#create CustomDirectoryRole#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/custom_directory_role#delete CustomDirectoryRole#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/custom_directory_role#read CustomDirectoryRole#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/custom_directory_role#update CustomDirectoryRole#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


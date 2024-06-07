// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package groupmember


type GroupMemberTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/group_member#create GroupMember#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/group_member#delete GroupMember#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/group_member#read GroupMember#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/group_member#update GroupMember#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


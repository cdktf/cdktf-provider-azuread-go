// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package invitation


type InvitationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/invitation#create Invitation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/invitation#delete Invitation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/invitation#read Invitation#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}


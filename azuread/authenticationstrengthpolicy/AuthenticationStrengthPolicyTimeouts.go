// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package authenticationstrengthpolicy


type AuthenticationStrengthPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.45.0/docs/resources/authentication_strength_policy#create AuthenticationStrengthPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.45.0/docs/resources/authentication_strength_policy#delete AuthenticationStrengthPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.45.0/docs/resources/authentication_strength_policy#read AuthenticationStrengthPolicy#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.45.0/docs/resources/authentication_strength_policy#update AuthenticationStrengthPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


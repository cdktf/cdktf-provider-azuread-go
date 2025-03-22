// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package claimsmappingpolicy


type ClaimsMappingPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.2.0/docs/resources/claims_mapping_policy#create ClaimsMappingPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.2.0/docs/resources/claims_mapping_policy#delete ClaimsMappingPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.2.0/docs/resources/claims_mapping_policy#read ClaimsMappingPolicy#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.2.0/docs/resources/claims_mapping_policy#update ClaimsMappingPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


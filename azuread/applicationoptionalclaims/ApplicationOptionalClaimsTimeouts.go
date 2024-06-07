// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationoptionalclaims


type ApplicationOptionalClaimsTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/application_optional_claims#create ApplicationOptionalClaimsA#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/application_optional_claims#delete ApplicationOptionalClaimsA#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/application_optional_claims#read ApplicationOptionalClaimsA#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/application_optional_claims#update ApplicationOptionalClaimsA#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationoptionalclaims


type ApplicationOptionalClaimsSaml2TokenA struct {
	// The name of the optional claim.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.2.0/docs/resources/application_optional_claims#name ApplicationOptionalClaimsA#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// List of additional properties of the claim.
	//
	// If a property exists in this list, it modifies the behaviour of the optional claim
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.2.0/docs/resources/application_optional_claims#additional_properties ApplicationOptionalClaimsA#additional_properties}
	AdditionalProperties *[]*string `field:"optional" json:"additionalProperties" yaml:"additionalProperties"`
	// Whether the claim specified by the client is necessary to ensure a smooth authorization experience.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.2.0/docs/resources/application_optional_claims#essential ApplicationOptionalClaimsA#essential}
	Essential interface{} `field:"optional" json:"essential" yaml:"essential"`
	// The source of the claim.
	//
	// If `source` is absent, the claim is a predefined optional claim. If `source` is `user`, the value of `name` is the extension property from the user object
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.2.0/docs/resources/application_optional_claims#source ApplicationOptionalClaimsA#source}
	Source *string `field:"optional" json:"source" yaml:"source"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package application


type ApplicationOptionalClaims struct {
	// access_token block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.53.1/docs/resources/application#access_token Application#access_token}
	AccessToken interface{} `field:"optional" json:"accessToken" yaml:"accessToken"`
	// id_token block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.53.1/docs/resources/application#id_token Application#id_token}
	IdToken interface{} `field:"optional" json:"idToken" yaml:"idToken"`
	// saml2_token block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.53.1/docs/resources/application#saml2_token Application#saml2_token}
	Saml2Token interface{} `field:"optional" json:"saml2Token" yaml:"saml2Token"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package application


type ApplicationSinglePageApplication struct {
	// The URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.6.0/docs/resources/application#redirect_uris Application#redirect_uris}
	RedirectUris *[]*string `field:"optional" json:"redirectUris" yaml:"redirectUris"`
}


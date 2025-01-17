// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationredirecturis


type ApplicationRedirectUrisTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.1.0/docs/resources/application_redirect_uris#create ApplicationRedirectUris#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.1.0/docs/resources/application_redirect_uris#delete ApplicationRedirectUris#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.1.0/docs/resources/application_redirect_uris#read ApplicationRedirectUris#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.1.0/docs/resources/application_redirect_uris#update ApplicationRedirectUris#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


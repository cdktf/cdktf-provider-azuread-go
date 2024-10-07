// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package application


type ApplicationFeatureTags struct {
	// Whether this application represents a custom SAML application for linked service principals.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.0.2/docs/resources/application#custom_single_sign_on Application#custom_single_sign_on}
	CustomSingleSignOn interface{} `field:"optional" json:"customSingleSignOn" yaml:"customSingleSignOn"`
	// Whether this application represents an Enterprise Application for linked service principals.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.0.2/docs/resources/application#enterprise Application#enterprise}
	Enterprise interface{} `field:"optional" json:"enterprise" yaml:"enterprise"`
	// Whether this application represents a gallery application for linked service principals.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.0.2/docs/resources/application#gallery Application#gallery}
	Gallery interface{} `field:"optional" json:"gallery" yaml:"gallery"`
	// Whether this application is invisible to users in My Apps and Office 365 Launcher.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.0.2/docs/resources/application#hide Application#hide}
	Hide interface{} `field:"optional" json:"hide" yaml:"hide"`
}


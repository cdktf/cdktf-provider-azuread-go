// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceprincipal


type ServicePrincipalFeatures struct {
	// Whether this service principal represents a custom SAML application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.53.0/docs/resources/service_principal#custom_single_sign_on_app ServicePrincipal#custom_single_sign_on_app}
	CustomSingleSignOnApp interface{} `field:"optional" json:"customSingleSignOnApp" yaml:"customSingleSignOnApp"`
	// Whether this service principal represents an Enterprise Application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.53.0/docs/resources/service_principal#enterprise_application ServicePrincipal#enterprise_application}
	EnterpriseApplication interface{} `field:"optional" json:"enterpriseApplication" yaml:"enterpriseApplication"`
	// Whether this service principal represents a gallery application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.53.0/docs/resources/service_principal#gallery_application ServicePrincipal#gallery_application}
	GalleryApplication interface{} `field:"optional" json:"galleryApplication" yaml:"galleryApplication"`
	// Whether this app is visible to users in My Apps and Office 365 Launcher.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.53.0/docs/resources/service_principal#visible_to_users ServicePrincipal#visible_to_users}
	VisibleToUsers interface{} `field:"optional" json:"visibleToUsers" yaml:"visibleToUsers"`
}


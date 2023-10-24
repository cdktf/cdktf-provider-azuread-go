// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package application


type ApplicationApiOauth2PermissionScope struct {
	// The unique identifier of the delegated permission.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/application#id Application#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/application#admin_consent_description Application#admin_consent_description}
	AdminConsentDescription *string `field:"optional" json:"adminConsentDescription" yaml:"adminConsentDescription"`
	// Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/application#admin_consent_display_name Application#admin_consent_display_name}
	AdminConsentDisplayName *string `field:"optional" json:"adminConsentDisplayName" yaml:"adminConsentDisplayName"`
	// Determines if the permission scope is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/application#enabled Application#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/application#type Application#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/application#user_consent_description Application#user_consent_description}
	UserConsentDescription *string `field:"optional" json:"userConsentDescription" yaml:"userConsentDescription"`
	// Display name for the delegated permission that appears in the end user consent experience.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/application#user_consent_display_name Application#user_consent_display_name}
	UserConsentDisplayName *string `field:"optional" json:"userConsentDisplayName" yaml:"userConsentDisplayName"`
	// The value that is used for the `scp` claim in OAuth 2.0 access tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/application#value Application#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}


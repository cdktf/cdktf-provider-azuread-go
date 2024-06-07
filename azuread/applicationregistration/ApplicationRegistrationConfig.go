// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationregistration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApplicationRegistrationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The display name for the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/application_registration#display_name ApplicationRegistration#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Description of the application as shown to end users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/application_registration#description ApplicationRegistration#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Configures the `groups` claim that the app expects issued in a user or OAuth access token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/application_registration#group_membership_claims ApplicationRegistration#group_membership_claims}
	GroupMembershipClaims *[]*string `field:"optional" json:"groupMembershipClaims" yaml:"groupMembershipClaims"`
	// URL of the home page for the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/application_registration#homepage_url ApplicationRegistration#homepage_url}
	HomepageUrl *string `field:"optional" json:"homepageUrl" yaml:"homepageUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/application_registration#id ApplicationRegistration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether this application can request an access token using OAuth implicit flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/application_registration#implicit_access_token_issuance_enabled ApplicationRegistration#implicit_access_token_issuance_enabled}
	ImplicitAccessTokenIssuanceEnabled interface{} `field:"optional" json:"implicitAccessTokenIssuanceEnabled" yaml:"implicitAccessTokenIssuanceEnabled"`
	// Whether this application can request an ID token using OAuth implicit flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/application_registration#implicit_id_token_issuance_enabled ApplicationRegistration#implicit_id_token_issuance_enabled}
	ImplicitIdTokenIssuanceEnabled interface{} `field:"optional" json:"implicitIdTokenIssuanceEnabled" yaml:"implicitIdTokenIssuanceEnabled"`
	// URL of the logout page for the application, where the session is cleared for single sign-out.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/application_registration#logout_url ApplicationRegistration#logout_url}
	LogoutUrl *string `field:"optional" json:"logoutUrl" yaml:"logoutUrl"`
	// URL of the marketing page for the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/application_registration#marketing_url ApplicationRegistration#marketing_url}
	MarketingUrl *string `field:"optional" json:"marketingUrl" yaml:"marketingUrl"`
	// User-specified notes relevant for the management of the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/application_registration#notes ApplicationRegistration#notes}
	Notes *string `field:"optional" json:"notes" yaml:"notes"`
	// URL of the privacy statement for the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/application_registration#privacy_statement_url ApplicationRegistration#privacy_statement_url}
	PrivacyStatementUrl *string `field:"optional" json:"privacyStatementUrl" yaml:"privacyStatementUrl"`
	// The access token version expected by this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/application_registration#requested_access_token_version ApplicationRegistration#requested_access_token_version}
	RequestedAccessTokenVersion *float64 `field:"optional" json:"requestedAccessTokenVersion" yaml:"requestedAccessTokenVersion"`
	// References application or contact information from a service or asset management database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/application_registration#service_management_reference ApplicationRegistration#service_management_reference}
	ServiceManagementReference *string `field:"optional" json:"serviceManagementReference" yaml:"serviceManagementReference"`
	// The Microsoft account types that are supported for the current application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/application_registration#sign_in_audience ApplicationRegistration#sign_in_audience}
	SignInAudience *string `field:"optional" json:"signInAudience" yaml:"signInAudience"`
	// URL of the support page for the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/application_registration#support_url ApplicationRegistration#support_url}
	SupportUrl *string `field:"optional" json:"supportUrl" yaml:"supportUrl"`
	// URL of the terms of service statement for the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/application_registration#terms_of_service_url ApplicationRegistration#terms_of_service_url}
	TermsOfServiceUrl *string `field:"optional" json:"termsOfServiceUrl" yaml:"termsOfServiceUrl"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/application_registration#timeouts ApplicationRegistration#timeouts}
	Timeouts *ApplicationRegistrationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


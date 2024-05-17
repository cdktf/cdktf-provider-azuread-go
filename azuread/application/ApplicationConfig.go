// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package application

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApplicationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.50.0/docs/resources/application#display_name Application#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// api block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.50.0/docs/resources/application#api Application#api}
	Api *ApplicationApi `field:"optional" json:"api" yaml:"api"`
	// app_role block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.50.0/docs/resources/application#app_role Application#app_role}
	AppRole interface{} `field:"optional" json:"appRole" yaml:"appRole"`
	// Description of the application as shown to end users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.50.0/docs/resources/application#description Application#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies whether this application supports device authentication without a user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.50.0/docs/resources/application#device_only_auth_enabled Application#device_only_auth_enabled}
	DeviceOnlyAuthEnabled interface{} `field:"optional" json:"deviceOnlyAuthEnabled" yaml:"deviceOnlyAuthEnabled"`
	// Specifies whether the application is a public client.
	//
	// Appropriate for apps using token grant flows that don't use a redirect URI
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.50.0/docs/resources/application#fallback_public_client_enabled Application#fallback_public_client_enabled}
	FallbackPublicClientEnabled interface{} `field:"optional" json:"fallbackPublicClientEnabled" yaml:"fallbackPublicClientEnabled"`
	// feature_tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.50.0/docs/resources/application#feature_tags Application#feature_tags}
	FeatureTags interface{} `field:"optional" json:"featureTags" yaml:"featureTags"`
	// Configures the `groups` claim issued in a user or OAuth 2.0 access token that the app expects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.50.0/docs/resources/application#group_membership_claims Application#group_membership_claims}
	GroupMembershipClaims *[]*string `field:"optional" json:"groupMembershipClaims" yaml:"groupMembershipClaims"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.50.0/docs/resources/application#id Application#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The user-defined URI(s) that uniquely identify an application within its Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.50.0/docs/resources/application#identifier_uris Application#identifier_uris}
	IdentifierUris *[]*string `field:"optional" json:"identifierUris" yaml:"identifierUris"`
	// Base64 encoded logo image in gif, png or jpeg format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.50.0/docs/resources/application#logo_image Application#logo_image}
	LogoImage *string `field:"optional" json:"logoImage" yaml:"logoImage"`
	// URL of the application's marketing page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.50.0/docs/resources/application#marketing_url Application#marketing_url}
	MarketingUrl *string `field:"optional" json:"marketingUrl" yaml:"marketingUrl"`
	// User-specified notes relevant for the management of the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.50.0/docs/resources/application#notes Application#notes}
	Notes *string `field:"optional" json:"notes" yaml:"notes"`
	// Specifies whether, as part of OAuth 2.0 token requests, Azure AD allows POST requests, as opposed to GET requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.50.0/docs/resources/application#oauth2_post_response_required Application#oauth2_post_response_required}
	Oauth2PostResponseRequired interface{} `field:"optional" json:"oauth2PostResponseRequired" yaml:"oauth2PostResponseRequired"`
	// optional_claims block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.50.0/docs/resources/application#optional_claims Application#optional_claims}
	OptionalClaims *ApplicationOptionalClaims `field:"optional" json:"optionalClaims" yaml:"optionalClaims"`
	// A list of object IDs of principals that will be granted ownership of the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.50.0/docs/resources/application#owners Application#owners}
	Owners *[]*string `field:"optional" json:"owners" yaml:"owners"`
	// If `true`, will return an error if an existing application is found with the same name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.50.0/docs/resources/application#prevent_duplicate_names Application#prevent_duplicate_names}
	PreventDuplicateNames interface{} `field:"optional" json:"preventDuplicateNames" yaml:"preventDuplicateNames"`
	// URL of the application's privacy statement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.50.0/docs/resources/application#privacy_statement_url Application#privacy_statement_url}
	PrivacyStatementUrl *string `field:"optional" json:"privacyStatementUrl" yaml:"privacyStatementUrl"`
	// public_client block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.50.0/docs/resources/application#public_client Application#public_client}
	PublicClient *ApplicationPublicClient `field:"optional" json:"publicClient" yaml:"publicClient"`
	// required_resource_access block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.50.0/docs/resources/application#required_resource_access Application#required_resource_access}
	RequiredResourceAccess interface{} `field:"optional" json:"requiredResourceAccess" yaml:"requiredResourceAccess"`
	// References application or service contact information from a Service or Asset Management database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.50.0/docs/resources/application#service_management_reference Application#service_management_reference}
	ServiceManagementReference *string `field:"optional" json:"serviceManagementReference" yaml:"serviceManagementReference"`
	// The Microsoft account types that are supported for the current application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.50.0/docs/resources/application#sign_in_audience Application#sign_in_audience}
	SignInAudience *string `field:"optional" json:"signInAudience" yaml:"signInAudience"`
	// single_page_application block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.50.0/docs/resources/application#single_page_application Application#single_page_application}
	SinglePageApplication *ApplicationSinglePageApplication `field:"optional" json:"singlePageApplication" yaml:"singlePageApplication"`
	// URL of the application's support page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.50.0/docs/resources/application#support_url Application#support_url}
	SupportUrl *string `field:"optional" json:"supportUrl" yaml:"supportUrl"`
	// A set of tags to apply to the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.50.0/docs/resources/application#tags Application#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// Unique ID of the application template from which this application is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.50.0/docs/resources/application#template_id Application#template_id}
	TemplateId *string `field:"optional" json:"templateId" yaml:"templateId"`
	// URL of the application's terms of service statement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.50.0/docs/resources/application#terms_of_service_url Application#terms_of_service_url}
	TermsOfServiceUrl *string `field:"optional" json:"termsOfServiceUrl" yaml:"termsOfServiceUrl"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.50.0/docs/resources/application#timeouts Application#timeouts}
	Timeouts *ApplicationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// web block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.50.0/docs/resources/application#web Application#web}
	Web *ApplicationWeb `field:"optional" json:"web" yaml:"web"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceprincipal

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServicePrincipalConfig struct {
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
	// Whether or not the service principal account is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.46.0/docs/resources/service_principal#account_enabled ServicePrincipal#account_enabled}
	AccountEnabled interface{} `field:"optional" json:"accountEnabled" yaml:"accountEnabled"`
	// A list of alternative names, used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.46.0/docs/resources/service_principal#alternative_names ServicePrincipal#alternative_names}
	AlternativeNames *[]*string `field:"optional" json:"alternativeNames" yaml:"alternativeNames"`
	// The application ID (client ID) of the application for which to create a service principal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.46.0/docs/resources/service_principal#application_id ServicePrincipal#application_id}
	ApplicationId *string `field:"optional" json:"applicationId" yaml:"applicationId"`
	// Whether this service principal requires an app role assignment to a user or group before Azure AD will issue a user or access token to the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.46.0/docs/resources/service_principal#app_role_assignment_required ServicePrincipal#app_role_assignment_required}
	AppRoleAssignmentRequired interface{} `field:"optional" json:"appRoleAssignmentRequired" yaml:"appRoleAssignmentRequired"`
	// The client ID of the application for which to create a service principal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.46.0/docs/resources/service_principal#client_id ServicePrincipal#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Description of the service principal provided for internal end-users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.46.0/docs/resources/service_principal#description ServicePrincipal#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// features block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.46.0/docs/resources/service_principal#features ServicePrincipal#features}
	Features interface{} `field:"optional" json:"features" yaml:"features"`
	// feature_tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.46.0/docs/resources/service_principal#feature_tags ServicePrincipal#feature_tags}
	FeatureTags interface{} `field:"optional" json:"featureTags" yaml:"featureTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.46.0/docs/resources/service_principal#id ServicePrincipal#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The URL where the service provider redirects the user to Azure AD to authenticate.
	//
	// Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps. When blank, Azure AD performs IdP-initiated sign-on for applications configured with SAML-based single sign-on
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.46.0/docs/resources/service_principal#login_url ServicePrincipal#login_url}
	LoginUrl *string `field:"optional" json:"loginUrl" yaml:"loginUrl"`
	// Free text field to capture information about the service principal, typically used for operational purposes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.46.0/docs/resources/service_principal#notes ServicePrincipal#notes}
	Notes *string `field:"optional" json:"notes" yaml:"notes"`
	// List of email addresses where Azure AD sends a notification when the active certificate is near the expiration date.
	//
	// This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.46.0/docs/resources/service_principal#notification_email_addresses ServicePrincipal#notification_email_addresses}
	NotificationEmailAddresses *[]*string `field:"optional" json:"notificationEmailAddresses" yaml:"notificationEmailAddresses"`
	// A list of object IDs of principals that will be granted ownership of the service principal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.46.0/docs/resources/service_principal#owners ServicePrincipal#owners}
	Owners *[]*string `field:"optional" json:"owners" yaml:"owners"`
	// The single sign-on mode configured for this application.
	//
	// Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.46.0/docs/resources/service_principal#preferred_single_sign_on_mode ServicePrincipal#preferred_single_sign_on_mode}
	PreferredSingleSignOnMode *string `field:"optional" json:"preferredSingleSignOnMode" yaml:"preferredSingleSignOnMode"`
	// saml_single_sign_on block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.46.0/docs/resources/service_principal#saml_single_sign_on ServicePrincipal#saml_single_sign_on}
	SamlSingleSignOn *ServicePrincipalSamlSingleSignOn `field:"optional" json:"samlSingleSignOn" yaml:"samlSingleSignOn"`
	// A set of tags to apply to the service principal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.46.0/docs/resources/service_principal#tags ServicePrincipal#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.46.0/docs/resources/service_principal#timeouts ServicePrincipal#timeouts}
	Timeouts *ServicePrincipalTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// When true, the resource will return an existing service principal instead of failing with an error.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.46.0/docs/resources/service_principal#use_existing ServicePrincipal#use_existing}
	UseExisting interface{} `field:"optional" json:"useExisting" yaml:"useExisting"`
}


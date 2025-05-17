// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationfederatedidentitycredential

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApplicationFederatedIdentityCredentialConfig struct {
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
	// The resource ID of the application for which this federated identity credential should be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/application_federated_identity_credential#application_id ApplicationFederatedIdentityCredential#application_id}
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// List of audiences that can appear in the external token.
	//
	// This specifies what should be accepted in the `aud` claim of incoming tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/application_federated_identity_credential#audiences ApplicationFederatedIdentityCredential#audiences}
	Audiences *[]*string `field:"required" json:"audiences" yaml:"audiences"`
	// A unique display name for the federated identity credential.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/application_federated_identity_credential#display_name ApplicationFederatedIdentityCredential#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The URL of the external identity provider, which must match the issuer claim of the external token being exchanged.
	//
	// The combination of the values of issuer and subject must be unique on the app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/application_federated_identity_credential#issuer ApplicationFederatedIdentityCredential#issuer}
	Issuer *string `field:"required" json:"issuer" yaml:"issuer"`
	// The identifier of the external software workload within the external identity provider.
	//
	// The combination of issuer and subject must be unique on the app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/application_federated_identity_credential#subject ApplicationFederatedIdentityCredential#subject}
	Subject *string `field:"required" json:"subject" yaml:"subject"`
	// A description for the federated identity credential.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/application_federated_identity_credential#description ApplicationFederatedIdentityCredential#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/application_federated_identity_credential#id ApplicationFederatedIdentityCredential#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/application_federated_identity_credential#timeouts ApplicationFederatedIdentityCredential#timeouts}
	Timeouts *ApplicationFederatedIdentityCredentialTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


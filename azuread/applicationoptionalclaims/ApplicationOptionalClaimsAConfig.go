// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationoptionalclaims

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApplicationOptionalClaimsAConfig struct {
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
	// The resource ID of the application to which these optional claims belong.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.49.0/docs/resources/application_optional_claims#application_id ApplicationOptionalClaimsA#application_id}
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// access_token block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.49.0/docs/resources/application_optional_claims#access_token ApplicationOptionalClaimsA#access_token}
	AccessToken interface{} `field:"optional" json:"accessToken" yaml:"accessToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.49.0/docs/resources/application_optional_claims#id ApplicationOptionalClaimsA#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// id_token block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.49.0/docs/resources/application_optional_claims#id_token ApplicationOptionalClaimsA#id_token}
	IdToken interface{} `field:"optional" json:"idToken" yaml:"idToken"`
	// saml2_token block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.49.0/docs/resources/application_optional_claims#saml2_token ApplicationOptionalClaimsA#saml2_token}
	Saml2Token interface{} `field:"optional" json:"saml2Token" yaml:"saml2Token"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.49.0/docs/resources/application_optional_claims#timeouts ApplicationOptionalClaimsA#timeouts}
	Timeouts *ApplicationOptionalClaimsTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


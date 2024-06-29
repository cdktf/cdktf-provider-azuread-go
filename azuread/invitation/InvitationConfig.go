// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package invitation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type InvitationConfig struct {
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
	// The URL that the user should be redirected to once the invitation is redeemed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.53.1/docs/resources/invitation#redirect_url Invitation#redirect_url}
	RedirectUrl *string `field:"required" json:"redirectUrl" yaml:"redirectUrl"`
	// The email address of the user being invited.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.53.1/docs/resources/invitation#user_email_address Invitation#user_email_address}
	UserEmailAddress *string `field:"required" json:"userEmailAddress" yaml:"userEmailAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.53.1/docs/resources/invitation#id Invitation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// message block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.53.1/docs/resources/invitation#message Invitation#message}
	Message *InvitationMessage `field:"optional" json:"message" yaml:"message"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.53.1/docs/resources/invitation#timeouts Invitation#timeouts}
	Timeouts *InvitationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The display name of the user being invited.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.53.1/docs/resources/invitation#user_display_name Invitation#user_display_name}
	UserDisplayName *string `field:"optional" json:"userDisplayName" yaml:"userDisplayName"`
	// The user type of the user being invited.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.53.1/docs/resources/invitation#user_type Invitation#user_type}
	UserType *string `field:"optional" json:"userType" yaml:"userType"`
}


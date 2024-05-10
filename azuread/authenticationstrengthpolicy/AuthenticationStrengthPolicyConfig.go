// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package authenticationstrengthpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AuthenticationStrengthPolicyConfig struct {
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
	// The allowed MFA methods for this policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.49.0/docs/resources/authentication_strength_policy#allowed_combinations AuthenticationStrengthPolicy#allowed_combinations}
	AllowedCombinations *[]*string `field:"required" json:"allowedCombinations" yaml:"allowedCombinations"`
	// The display name for the authentication strength policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.49.0/docs/resources/authentication_strength_policy#display_name AuthenticationStrengthPolicy#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The description for the authentication strength policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.49.0/docs/resources/authentication_strength_policy#description AuthenticationStrengthPolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.49.0/docs/resources/authentication_strength_policy#id AuthenticationStrengthPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.49.0/docs/resources/authentication_strength_policy#timeouts AuthenticationStrengthPolicy#timeouts}
	Timeouts *AuthenticationStrengthPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


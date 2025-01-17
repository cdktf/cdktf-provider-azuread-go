// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package conditionalaccesspolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConditionalAccessPolicyConfig struct {
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
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.1.0/docs/resources/conditional_access_policy#conditions ConditionalAccessPolicy#conditions}
	Conditions *ConditionalAccessPolicyConditions `field:"required" json:"conditions" yaml:"conditions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.1.0/docs/resources/conditional_access_policy#display_name ConditionalAccessPolicy#display_name}.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.1.0/docs/resources/conditional_access_policy#state ConditionalAccessPolicy#state}.
	State *string `field:"required" json:"state" yaml:"state"`
	// grant_controls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.1.0/docs/resources/conditional_access_policy#grant_controls ConditionalAccessPolicy#grant_controls}
	GrantControls *ConditionalAccessPolicyGrantControls `field:"optional" json:"grantControls" yaml:"grantControls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.1.0/docs/resources/conditional_access_policy#id ConditionalAccessPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// session_controls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.1.0/docs/resources/conditional_access_policy#session_controls ConditionalAccessPolicy#session_controls}
	SessionControls *ConditionalAccessPolicySessionControls `field:"optional" json:"sessionControls" yaml:"sessionControls"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.1.0/docs/resources/conditional_access_policy#timeouts ConditionalAccessPolicy#timeouts}
	Timeouts *ConditionalAccessPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


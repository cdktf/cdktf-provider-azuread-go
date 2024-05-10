// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationapiaccess

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApplicationApiAccessConfig struct {
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
	// The client ID of the API to which access is being granted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.49.0/docs/resources/application_api_access#api_client_id ApplicationApiAccess#api_client_id}
	ApiClientId *string `field:"required" json:"apiClientId" yaml:"apiClientId"`
	// The resource ID of the application to which this API access is granted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.49.0/docs/resources/application_api_access#application_id ApplicationApiAccess#application_id}
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.49.0/docs/resources/application_api_access#id ApplicationApiAccess#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A set of role IDs to be granted to the application, as published by the API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.49.0/docs/resources/application_api_access#role_ids ApplicationApiAccess#role_ids}
	RoleIds *[]*string `field:"optional" json:"roleIds" yaml:"roleIds"`
	// A set of scope IDs to be granted to the application, as published by the API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.49.0/docs/resources/application_api_access#scope_ids ApplicationApiAccess#scope_ids}
	ScopeIds *[]*string `field:"optional" json:"scopeIds" yaml:"scopeIds"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.49.0/docs/resources/application_api_access#timeouts ApplicationApiAccess#timeouts}
	Timeouts *ApplicationApiAccessTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


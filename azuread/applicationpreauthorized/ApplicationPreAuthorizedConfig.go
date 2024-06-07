// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationpreauthorized

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApplicationPreAuthorizedConfig struct {
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
	// The IDs of the permission scopes required by the pre-authorized application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/application_pre_authorized#permission_ids ApplicationPreAuthorized#permission_ids}
	PermissionIds *[]*string `field:"required" json:"permissionIds" yaml:"permissionIds"`
	// The resource ID of the application to which this pre-authorized application should be added.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/application_pre_authorized#application_id ApplicationPreAuthorized#application_id}
	ApplicationId *string `field:"optional" json:"applicationId" yaml:"applicationId"`
	// The object ID of the application to which this pre-authorized application should be added.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/application_pre_authorized#application_object_id ApplicationPreAuthorized#application_object_id}
	ApplicationObjectId *string `field:"optional" json:"applicationObjectId" yaml:"applicationObjectId"`
	// The application ID of the pre-authorized application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/application_pre_authorized#authorized_app_id ApplicationPreAuthorized#authorized_app_id}
	AuthorizedAppId *string `field:"optional" json:"authorizedAppId" yaml:"authorizedAppId"`
	// The client ID of the pre-authorized application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/application_pre_authorized#authorized_client_id ApplicationPreAuthorized#authorized_client_id}
	AuthorizedClientId *string `field:"optional" json:"authorizedClientId" yaml:"authorizedClientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/application_pre_authorized#id ApplicationPreAuthorized#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/application_pre_authorized#timeouts ApplicationPreAuthorized#timeouts}
	Timeouts *ApplicationPreAuthorizedTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


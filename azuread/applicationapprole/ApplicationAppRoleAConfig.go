// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationapprole

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApplicationAppRoleAConfig struct {
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
	// Specifies whether this app role definition can be assigned to users and groups by setting to `User`, or to other applications (that are accessing this application in a standalone scenario) by setting to `Application`, or to both.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/application_app_role#allowed_member_types ApplicationAppRoleA#allowed_member_types}
	AllowedMemberTypes *[]*string `field:"required" json:"allowedMemberTypes" yaml:"allowedMemberTypes"`
	// The resource ID of the application to which this app role should be applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/application_app_role#application_id ApplicationAppRoleA#application_id}
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// Description of the app role that appears when the role is being assigned and, if the role functions as an application permissions, during the consent experiences.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/application_app_role#description ApplicationAppRoleA#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// Display name for the app role that appears during app role assignment and in consent experiences.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/application_app_role#display_name ApplicationAppRoleA#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The unique identifier of the app role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/application_app_role#role_id ApplicationAppRoleA#role_id}
	RoleId *string `field:"required" json:"roleId" yaml:"roleId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/application_app_role#id ApplicationAppRoleA#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/application_app_role#timeouts ApplicationAppRoleA#timeouts}
	Timeouts *ApplicationAppRoleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The value that is used for the `roles` claim in ID tokens and OAuth access tokens that are authenticating an assigned service or user principal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/application_app_role#value ApplicationAppRoleA#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}


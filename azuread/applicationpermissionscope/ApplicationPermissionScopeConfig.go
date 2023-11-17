// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationpermissionscope

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApplicationPermissionScopeConfig struct {
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
	// Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.46.0/docs/resources/application_permission_scope#admin_consent_description ApplicationPermissionScope#admin_consent_description}
	AdminConsentDescription *string `field:"required" json:"adminConsentDescription" yaml:"adminConsentDescription"`
	// Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.46.0/docs/resources/application_permission_scope#admin_consent_display_name ApplicationPermissionScope#admin_consent_display_name}
	AdminConsentDisplayName *string `field:"required" json:"adminConsentDisplayName" yaml:"adminConsentDisplayName"`
	// The resource ID of the application to which this permission scope should be applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.46.0/docs/resources/application_permission_scope#application_id ApplicationPermissionScope#application_id}
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The unique identifier of the permission scope.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.46.0/docs/resources/application_permission_scope#scope_id ApplicationPermissionScope#scope_id}
	ScopeId *string `field:"required" json:"scopeId" yaml:"scopeId"`
	// The value that is used for the `scp` claim in OAuth access tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.46.0/docs/resources/application_permission_scope#value ApplicationPermissionScope#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.46.0/docs/resources/application_permission_scope#id ApplicationPermissionScope#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.46.0/docs/resources/application_permission_scope#timeouts ApplicationPermissionScope#timeouts}
	Timeouts *ApplicationPermissionScopeTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.46.0/docs/resources/application_permission_scope#type ApplicationPermissionScope#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.46.0/docs/resources/application_permission_scope#user_consent_description ApplicationPermissionScope#user_consent_description}
	UserConsentDescription *string `field:"optional" json:"userConsentDescription" yaml:"userConsentDescription"`
	// Display name for the delegated permission that appears in the end user consent experience.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.46.0/docs/resources/application_permission_scope#user_consent_display_name ApplicationPermissionScope#user_consent_display_name}
	UserConsentDisplayName *string `field:"optional" json:"userConsentDisplayName" yaml:"userConsentDisplayName"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package directoryroleeligibilityschedulerequest

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DirectoryRoleEligibilityScheduleRequestConfig struct {
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
	// Identifier of the directory object representing the scope of the role eligibility schedule request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.53.0/docs/resources/directory_role_eligibility_schedule_request#directory_scope_id DirectoryRoleEligibilityScheduleRequest#directory_scope_id}
	DirectoryScopeId *string `field:"required" json:"directoryScopeId" yaml:"directoryScopeId"`
	// Justification for why the role is assigned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.53.0/docs/resources/directory_role_eligibility_schedule_request#justification DirectoryRoleEligibilityScheduleRequest#justification}
	Justification *string `field:"required" json:"justification" yaml:"justification"`
	// The object ID of the member principal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.53.0/docs/resources/directory_role_eligibility_schedule_request#principal_id DirectoryRoleEligibilityScheduleRequest#principal_id}
	PrincipalId *string `field:"required" json:"principalId" yaml:"principalId"`
	// The object ID of the directory role for this role eligibility schedule request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.53.0/docs/resources/directory_role_eligibility_schedule_request#role_definition_id DirectoryRoleEligibilityScheduleRequest#role_definition_id}
	RoleDefinitionId *string `field:"required" json:"roleDefinitionId" yaml:"roleDefinitionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.53.0/docs/resources/directory_role_eligibility_schedule_request#id DirectoryRoleEligibilityScheduleRequest#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.53.0/docs/resources/directory_role_eligibility_schedule_request#timeouts DirectoryRoleEligibilityScheduleRequest#timeouts}
	Timeouts *DirectoryRoleEligibilityScheduleRequestTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


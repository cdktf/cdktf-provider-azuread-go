// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesspackageassignmentpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccessPackageAssignmentPolicyConfig struct {
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
	// The ID of the access package that will contain the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.6.0/docs/resources/access_package_assignment_policy#access_package_id AccessPackageAssignmentPolicy#access_package_id}
	AccessPackageId *string `field:"required" json:"accessPackageId" yaml:"accessPackageId"`
	// The description of the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.6.0/docs/resources/access_package_assignment_policy#description AccessPackageAssignmentPolicy#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// The display name of the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.6.0/docs/resources/access_package_assignment_policy#display_name AccessPackageAssignmentPolicy#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// approval_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.6.0/docs/resources/access_package_assignment_policy#approval_settings AccessPackageAssignmentPolicy#approval_settings}
	ApprovalSettings *AccessPackageAssignmentPolicyApprovalSettings `field:"optional" json:"approvalSettings" yaml:"approvalSettings"`
	// assignment_review_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.6.0/docs/resources/access_package_assignment_policy#assignment_review_settings AccessPackageAssignmentPolicy#assignment_review_settings}
	AssignmentReviewSettings *AccessPackageAssignmentPolicyAssignmentReviewSettings `field:"optional" json:"assignmentReviewSettings" yaml:"assignmentReviewSettings"`
	// How many days this assignment is valid for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.6.0/docs/resources/access_package_assignment_policy#duration_in_days AccessPackageAssignmentPolicy#duration_in_days}
	DurationInDays *float64 `field:"optional" json:"durationInDays" yaml:"durationInDays"`
	// The date that this assignment expires, formatted as an RFC3339 date string in UTC (e.g. 2018-01-01T01:02:03Z).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.6.0/docs/resources/access_package_assignment_policy#expiration_date AccessPackageAssignmentPolicy#expiration_date}
	ExpirationDate *string `field:"optional" json:"expirationDate" yaml:"expirationDate"`
	// When enabled, users will be able to request extension of their access to this package before their access expires.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.6.0/docs/resources/access_package_assignment_policy#extension_enabled AccessPackageAssignmentPolicy#extension_enabled}
	ExtensionEnabled interface{} `field:"optional" json:"extensionEnabled" yaml:"extensionEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.6.0/docs/resources/access_package_assignment_policy#id AccessPackageAssignmentPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// question block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.6.0/docs/resources/access_package_assignment_policy#question AccessPackageAssignmentPolicy#question}
	Question interface{} `field:"optional" json:"question" yaml:"question"`
	// requestor_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.6.0/docs/resources/access_package_assignment_policy#requestor_settings AccessPackageAssignmentPolicy#requestor_settings}
	RequestorSettings *AccessPackageAssignmentPolicyRequestorSettings `field:"optional" json:"requestorSettings" yaml:"requestorSettings"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.6.0/docs/resources/access_package_assignment_policy#timeouts AccessPackageAssignmentPolicy#timeouts}
	Timeouts *AccessPackageAssignmentPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


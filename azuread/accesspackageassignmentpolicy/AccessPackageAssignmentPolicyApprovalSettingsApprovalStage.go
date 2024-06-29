// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesspackageassignmentpolicy


type AccessPackageAssignmentPolicyApprovalSettingsApprovalStage struct {
	// Decision must be made in how many days?
	//
	// If a request is not approved within this time period after it is made, it will be automatically rejected
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.53.1/docs/resources/access_package_assignment_policy#approval_timeout_in_days AccessPackageAssignmentPolicy#approval_timeout_in_days}
	ApprovalTimeoutInDays *float64 `field:"required" json:"approvalTimeoutInDays" yaml:"approvalTimeoutInDays"`
	// If no action taken, forward to alternate approvers?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.53.1/docs/resources/access_package_assignment_policy#alternative_approval_enabled AccessPackageAssignmentPolicy#alternative_approval_enabled}
	AlternativeApprovalEnabled interface{} `field:"optional" json:"alternativeApprovalEnabled" yaml:"alternativeApprovalEnabled"`
	// alternative_approver block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.53.1/docs/resources/access_package_assignment_policy#alternative_approver AccessPackageAssignmentPolicy#alternative_approver}
	AlternativeApprover interface{} `field:"optional" json:"alternativeApprover" yaml:"alternativeApprover"`
	// Whether an approver must provide a justification for their decision. Justification is visible to other approvers and the requestor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.53.1/docs/resources/access_package_assignment_policy#approver_justification_required AccessPackageAssignmentPolicy#approver_justification_required}
	ApproverJustificationRequired interface{} `field:"optional" json:"approverJustificationRequired" yaml:"approverJustificationRequired"`
	// Forward to alternate approver(s) after how many days?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.53.1/docs/resources/access_package_assignment_policy#enable_alternative_approval_in_days AccessPackageAssignmentPolicy#enable_alternative_approval_in_days}
	EnableAlternativeApprovalInDays *float64 `field:"optional" json:"enableAlternativeApprovalInDays" yaml:"enableAlternativeApprovalInDays"`
	// primary_approver block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.53.1/docs/resources/access_package_assignment_policy#primary_approver AccessPackageAssignmentPolicy#primary_approver}
	PrimaryApprover interface{} `field:"optional" json:"primaryApprover" yaml:"primaryApprover"`
}


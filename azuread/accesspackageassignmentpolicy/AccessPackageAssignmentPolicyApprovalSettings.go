// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesspackageassignmentpolicy


type AccessPackageAssignmentPolicyApprovalSettings struct {
	// Whether an approval is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.48.0/docs/resources/access_package_assignment_policy#approval_required AccessPackageAssignmentPolicy#approval_required}
	ApprovalRequired interface{} `field:"optional" json:"approvalRequired" yaml:"approvalRequired"`
	// Whether an approval is required to grant extension. Same approval settings used to approve initial access will apply.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.48.0/docs/resources/access_package_assignment_policy#approval_required_for_extension AccessPackageAssignmentPolicy#approval_required_for_extension}
	ApprovalRequiredForExtension interface{} `field:"optional" json:"approvalRequiredForExtension" yaml:"approvalRequiredForExtension"`
	// approval_stage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.48.0/docs/resources/access_package_assignment_policy#approval_stage AccessPackageAssignmentPolicy#approval_stage}
	ApprovalStage interface{} `field:"optional" json:"approvalStage" yaml:"approvalStage"`
	// Whether requestor are required to provide a justification to request an access package.
	//
	// Justification is visible to other approvers and the requestor
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.48.0/docs/resources/access_package_assignment_policy#requestor_justification_required AccessPackageAssignmentPolicy#requestor_justification_required}
	RequestorJustificationRequired interface{} `field:"optional" json:"requestorJustificationRequired" yaml:"requestorJustificationRequired"`
}


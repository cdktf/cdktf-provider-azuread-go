// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesspackageassignmentpolicy


type AccessPackageAssignmentPolicyQuestion struct {
	// text block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.53.1/docs/resources/access_package_assignment_policy#text AccessPackageAssignmentPolicy#text}
	Text *AccessPackageAssignmentPolicyQuestionText `field:"required" json:"text" yaml:"text"`
	// choice block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.53.1/docs/resources/access_package_assignment_policy#choice AccessPackageAssignmentPolicy#choice}
	Choice interface{} `field:"optional" json:"choice" yaml:"choice"`
	// Whether this question is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.53.1/docs/resources/access_package_assignment_policy#required AccessPackageAssignmentPolicy#required}
	Required interface{} `field:"optional" json:"required" yaml:"required"`
	// The sequence number of this question.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.53.1/docs/resources/access_package_assignment_policy#sequence AccessPackageAssignmentPolicy#sequence}
	Sequence *float64 `field:"optional" json:"sequence" yaml:"sequence"`
}


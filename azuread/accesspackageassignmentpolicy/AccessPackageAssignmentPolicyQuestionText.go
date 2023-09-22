// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesspackageassignmentpolicy


type AccessPackageAssignmentPolicyQuestionText struct {
	// The default text of this question.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.43.0/docs/resources/access_package_assignment_policy#default_text AccessPackageAssignmentPolicy#default_text}
	DefaultText *string `field:"required" json:"defaultText" yaml:"defaultText"`
	// localized_text block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.43.0/docs/resources/access_package_assignment_policy#localized_text AccessPackageAssignmentPolicy#localized_text}
	LocalizedText interface{} `field:"optional" json:"localizedText" yaml:"localizedText"`
}


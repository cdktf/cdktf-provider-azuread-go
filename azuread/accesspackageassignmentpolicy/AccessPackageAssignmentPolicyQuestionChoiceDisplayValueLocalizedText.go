// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesspackageassignmentpolicy


type AccessPackageAssignmentPolicyQuestionChoiceDisplayValueLocalizedText struct {
	// The localized content of this question.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.3.0/docs/resources/access_package_assignment_policy#content AccessPackageAssignmentPolicy#content}
	Content *string `field:"required" json:"content" yaml:"content"`
	// The language code of this question content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.3.0/docs/resources/access_package_assignment_policy#language_code AccessPackageAssignmentPolicy#language_code}
	LanguageCode *string `field:"required" json:"languageCode" yaml:"languageCode"`
}


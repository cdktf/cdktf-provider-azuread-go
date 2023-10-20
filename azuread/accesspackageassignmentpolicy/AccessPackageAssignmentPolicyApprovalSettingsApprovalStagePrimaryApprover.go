// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesspackageassignmentpolicy


type AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApprover struct {
	// Type of users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.0/docs/resources/access_package_assignment_policy#subject_type AccessPackageAssignmentPolicy#subject_type}
	SubjectType *string `field:"required" json:"subjectType" yaml:"subjectType"`
	// For a user in an approval stage, this property indicates whether the user is a backup fallback approver.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.0/docs/resources/access_package_assignment_policy#backup AccessPackageAssignmentPolicy#backup}
	Backup interface{} `field:"optional" json:"backup" yaml:"backup"`
	// The object ID of the subject.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.0/docs/resources/access_package_assignment_policy#object_id AccessPackageAssignmentPolicy#object_id}
	ObjectId *string `field:"optional" json:"objectId" yaml:"objectId"`
}


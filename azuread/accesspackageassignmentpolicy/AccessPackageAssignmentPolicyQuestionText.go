package accesspackageassignmentpolicy


type AccessPackageAssignmentPolicyQuestionText struct {
	// The default text of this question.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/access_package_assignment_policy#default_text AccessPackageAssignmentPolicy#default_text}
	DefaultText *string `field:"required" json:"defaultText" yaml:"defaultText"`
	// localized_text block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/access_package_assignment_policy#localized_text AccessPackageAssignmentPolicy#localized_text}
	LocalizedText interface{} `field:"optional" json:"localizedText" yaml:"localizedText"`
}


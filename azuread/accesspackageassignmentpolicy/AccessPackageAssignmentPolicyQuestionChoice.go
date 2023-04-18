package accesspackageassignmentpolicy


type AccessPackageAssignmentPolicyQuestionChoice struct {
	// The actual value of this choice.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.37.1/docs/resources/access_package_assignment_policy#actual_value AccessPackageAssignmentPolicy#actual_value}
	ActualValue *string `field:"required" json:"actualValue" yaml:"actualValue"`
	// display_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.37.1/docs/resources/access_package_assignment_policy#display_value AccessPackageAssignmentPolicy#display_value}
	DisplayValue *AccessPackageAssignmentPolicyQuestionChoiceDisplayValue `field:"required" json:"displayValue" yaml:"displayValue"`
}


package accesspackageassignmentpolicy


type AccessPackageAssignmentPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.37.1/docs/resources/access_package_assignment_policy#create AccessPackageAssignmentPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.37.1/docs/resources/access_package_assignment_policy#delete AccessPackageAssignmentPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.37.1/docs/resources/access_package_assignment_policy#read AccessPackageAssignmentPolicy#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.37.1/docs/resources/access_package_assignment_policy#update AccessPackageAssignmentPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


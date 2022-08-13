// Prebuilt azuread Provider for Terraform CDK (cdktf)
package azuread


type DirectoryRoleAssignmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/directory_role_assignment#create DirectoryRoleAssignment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/directory_role_assignment#delete DirectoryRoleAssignment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/directory_role_assignment#read DirectoryRoleAssignment#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/directory_role_assignment#update DirectoryRoleAssignment#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


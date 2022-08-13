// Prebuilt azuread Provider for Terraform CDK (cdktf)
package azuread


type AppRoleAssignmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/app_role_assignment#create AppRoleAssignment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/app_role_assignment#delete AppRoleAssignment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/app_role_assignment#read AppRoleAssignment#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}


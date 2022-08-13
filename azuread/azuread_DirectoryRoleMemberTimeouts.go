// Prebuilt azuread Provider for Terraform CDK (cdktf)
package azuread


type DirectoryRoleMemberTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/directory_role_member#create DirectoryRoleMember#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/directory_role_member#delete DirectoryRoleMember#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/directory_role_member#read DirectoryRoleMember#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/directory_role_member#update DirectoryRoleMember#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


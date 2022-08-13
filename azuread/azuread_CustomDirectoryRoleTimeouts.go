// Prebuilt azuread Provider for Terraform CDK (cdktf)
package azuread


type CustomDirectoryRoleTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/custom_directory_role#create CustomDirectoryRole#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/custom_directory_role#delete CustomDirectoryRole#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/custom_directory_role#read CustomDirectoryRole#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/custom_directory_role#update CustomDirectoryRole#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


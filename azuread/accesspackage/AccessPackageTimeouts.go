package accesspackage


type AccessPackageTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/access_package#create AccessPackage#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/access_package#delete AccessPackage#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/access_package#read AccessPackage#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/access_package#update AccessPackage#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


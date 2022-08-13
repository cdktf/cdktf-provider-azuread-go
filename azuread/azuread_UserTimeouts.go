// Prebuilt azuread Provider for Terraform CDK (cdktf)
package azuread


type UserTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/user#create User#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/user#delete User#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/user#read User#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/user#update User#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


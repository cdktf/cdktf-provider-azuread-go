// Prebuilt azuread Provider for Terraform CDK (cdktf)
package azuread


type AdministrativeUnitTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/administrative_unit#create AdministrativeUnit#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/administrative_unit#delete AdministrativeUnit#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/administrative_unit#read AdministrativeUnit#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/administrative_unit#update AdministrativeUnit#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


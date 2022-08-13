// Prebuilt azuread Provider for Terraform CDK (cdktf)
package azuread


type ApplicationPasswordTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/application_password#create ApplicationPassword#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/application_password#delete ApplicationPassword#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/application_password#read ApplicationPassword#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/application_password#update ApplicationPassword#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


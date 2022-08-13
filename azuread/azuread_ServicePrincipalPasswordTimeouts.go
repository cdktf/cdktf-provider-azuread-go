// Prebuilt azuread Provider for Terraform CDK (cdktf)
package azuread


type ServicePrincipalPasswordTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/service_principal_password#create ServicePrincipalPassword#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/service_principal_password#delete ServicePrincipalPassword#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/service_principal_password#read ServicePrincipalPassword#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/service_principal_password#update ServicePrincipalPassword#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


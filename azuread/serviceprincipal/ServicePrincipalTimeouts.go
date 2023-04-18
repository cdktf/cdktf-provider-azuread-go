package serviceprincipal


type ServicePrincipalTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.37.1/docs/resources/service_principal#create ServicePrincipal#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.37.1/docs/resources/service_principal#delete ServicePrincipal#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.37.1/docs/resources/service_principal#read ServicePrincipal#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.37.1/docs/resources/service_principal#update ServicePrincipal#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


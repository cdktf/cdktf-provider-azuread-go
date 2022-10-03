package applicationpreauthorized


type ApplicationPreAuthorizedTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/application_pre_authorized#create ApplicationPreAuthorized#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/application_pre_authorized#delete ApplicationPreAuthorized#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/application_pre_authorized#read ApplicationPreAuthorized#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/application_pre_authorized#update ApplicationPreAuthorized#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


package applicationcertificate


type ApplicationCertificateTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.40.0/docs/resources/application_certificate#create ApplicationCertificate#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.40.0/docs/resources/application_certificate#delete ApplicationCertificate#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.40.0/docs/resources/application_certificate#read ApplicationCertificate#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.40.0/docs/resources/application_certificate#update ApplicationCertificate#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


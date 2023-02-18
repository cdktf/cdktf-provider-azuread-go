package serviceprincipalcertificate


type ServicePrincipalCertificateTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/service_principal_certificate#create ServicePrincipalCertificate#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/service_principal_certificate#delete ServicePrincipalCertificate#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/service_principal_certificate#read ServicePrincipalCertificate#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/service_principal_certificate#update ServicePrincipalCertificate#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

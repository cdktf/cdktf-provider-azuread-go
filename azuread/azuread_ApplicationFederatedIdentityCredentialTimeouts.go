// Prebuilt azuread Provider for Terraform CDK (cdktf)
package azuread


type ApplicationFederatedIdentityCredentialTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/application_federated_identity_credential#create ApplicationFederatedIdentityCredential#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/application_federated_identity_credential#delete ApplicationFederatedIdentityCredential#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/application_federated_identity_credential#read ApplicationFederatedIdentityCredential#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/application_federated_identity_credential#update ApplicationFederatedIdentityCredential#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


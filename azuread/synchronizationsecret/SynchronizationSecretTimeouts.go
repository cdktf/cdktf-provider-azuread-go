package synchronizationsecret


type SynchronizationSecretTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.38.0/docs/resources/synchronization_secret#create SynchronizationSecret#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.38.0/docs/resources/synchronization_secret#delete SynchronizationSecret#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.38.0/docs/resources/synchronization_secret#read SynchronizationSecret#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.38.0/docs/resources/synchronization_secret#update SynchronizationSecret#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


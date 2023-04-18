package synchronizationsecret


type SynchronizationSecretCredential struct {
	// Name for this key-value pair.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.37.1/docs/resources/synchronization_secret#key SynchronizationSecret#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Value for this key-value pair.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.37.1/docs/resources/synchronization_secret#value SynchronizationSecret#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}


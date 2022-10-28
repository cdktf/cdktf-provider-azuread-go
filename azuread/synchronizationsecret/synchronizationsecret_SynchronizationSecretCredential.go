package synchronizationsecret


type SynchronizationSecretCredential struct {
	// Name for this key-value pair.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/synchronization_secret#key SynchronizationSecret#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Value for this key-value pair.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/synchronization_secret#value SynchronizationSecret#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}


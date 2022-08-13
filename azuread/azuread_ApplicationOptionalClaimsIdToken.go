// Prebuilt azuread Provider for Terraform CDK (cdktf)
package azuread


type ApplicationOptionalClaimsIdToken struct {
	// The name of the optional claim.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/application#name Application#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// List of additional properties of the claim.
	//
	// If a property exists in this list, it modifies the behaviour of the optional claim
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/application#additional_properties Application#additional_properties}
	AdditionalProperties *[]*string `field:"optional" json:"additionalProperties" yaml:"additionalProperties"`
	// Whether the claim specified by the client is necessary to ensure a smooth authorization experience.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/application#essential Application#essential}
	Essential interface{} `field:"optional" json:"essential" yaml:"essential"`
	// The source of the claim.
	//
	// If `source` is absent, the claim is a predefined optional claim. If `source` is `user`, the value of `name` is the extension property from the user object
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/application#source Application#source}
	Source *string `field:"optional" json:"source" yaml:"source"`
}


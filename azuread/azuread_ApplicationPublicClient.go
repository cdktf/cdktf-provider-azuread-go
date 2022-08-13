// Prebuilt azuread Provider for Terraform CDK (cdktf)
package azuread


type ApplicationPublicClient struct {
	// The URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/application#redirect_uris Application#redirect_uris}
	RedirectUris *[]*string `field:"optional" json:"redirectUris" yaml:"redirectUris"`
}


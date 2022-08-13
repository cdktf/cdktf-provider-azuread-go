// Prebuilt azuread Provider for Terraform CDK (cdktf)
package azuread


type ApplicationWeb struct {
	// Home page or landing page of the application.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/application#homepage_url Application#homepage_url}
	HomepageUrl *string `field:"optional" json:"homepageUrl" yaml:"homepageUrl"`
	// implicit_grant block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/application#implicit_grant Application#implicit_grant}
	ImplicitGrant *ApplicationWebImplicitGrant `field:"optional" json:"implicitGrant" yaml:"implicitGrant"`
	// The URL that will be used by Microsoft's authorization service to sign out a user using front-channel, back-channel or SAML logout protocols.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/application#logout_url Application#logout_url}
	LogoutUrl *string `field:"optional" json:"logoutUrl" yaml:"logoutUrl"`
	// The URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/application#redirect_uris Application#redirect_uris}
	RedirectUris *[]*string `field:"optional" json:"redirectUris" yaml:"redirectUris"`
}


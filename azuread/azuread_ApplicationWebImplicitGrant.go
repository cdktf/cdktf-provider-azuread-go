// Prebuilt azuread Provider for Terraform CDK (cdktf)
package azuread


type ApplicationWebImplicitGrant struct {
	// Whether this web application can request an access token using OAuth 2.0 implicit flow.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/application#access_token_issuance_enabled Application#access_token_issuance_enabled}
	AccessTokenIssuanceEnabled interface{} `field:"optional" json:"accessTokenIssuanceEnabled" yaml:"accessTokenIssuanceEnabled"`
	// Whether this web application can request an ID token using OAuth 2.0 implicit flow.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/application#id_token_issuance_enabled Application#id_token_issuance_enabled}
	IdTokenIssuanceEnabled interface{} `field:"optional" json:"idTokenIssuanceEnabled" yaml:"idTokenIssuanceEnabled"`
}


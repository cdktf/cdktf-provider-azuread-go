// Prebuilt azuread Provider for Terraform CDK (cdktf)
package azuread


type ApplicationApi struct {
	// Used for bundling consent if you have a solution that contains two parts: a client app and a custom web API app.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/application#known_client_applications Application#known_client_applications}
	KnownClientApplications *[]*string `field:"optional" json:"knownClientApplications" yaml:"knownClientApplications"`
	// Allows an application to use claims mapping without specifying a custom signing key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/application#mapped_claims_enabled Application#mapped_claims_enabled}
	MappedClaimsEnabled interface{} `field:"optional" json:"mappedClaimsEnabled" yaml:"mappedClaimsEnabled"`
	// oauth2_permission_scope block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/application#oauth2_permission_scope Application#oauth2_permission_scope}
	Oauth2PermissionScope interface{} `field:"optional" json:"oauth2PermissionScope" yaml:"oauth2PermissionScope"`
	// The access token version expected by this resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/application#requested_access_token_version Application#requested_access_token_version}
	RequestedAccessTokenVersion *float64 `field:"optional" json:"requestedAccessTokenVersion" yaml:"requestedAccessTokenVersion"`
}


// Prebuilt azuread Provider for Terraform CDK (cdktf)
package azuread


type AzureadProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread#alias AzureadProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Base64 encoded PKCS#12 certificate bundle to use when authenticating as a Service Principal using a Client Certificate.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread#client_certificate AzureadProvider#client_certificate}
	ClientCertificate *string `field:"optional" json:"clientCertificate" yaml:"clientCertificate"`
	// The password to decrypt the Client Certificate. For use when authenticating as a Service Principal using a Client Certificate.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread#client_certificate_password AzureadProvider#client_certificate_password}
	ClientCertificatePassword *string `field:"optional" json:"clientCertificatePassword" yaml:"clientCertificatePassword"`
	// The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service Principal using a Client Certificate.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread#client_certificate_path AzureadProvider#client_certificate_path}
	ClientCertificatePath *string `field:"optional" json:"clientCertificatePath" yaml:"clientCertificatePath"`
	// The Client ID which should be used for service principal authentication.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread#client_id AzureadProvider#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The application password to use when authenticating as a Service Principal using a Client Secret.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread#client_secret AzureadProvider#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// Disable the Terraform Partner ID, which is used if a custom `partner_id` isn't specified.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread#disable_terraform_partner_id AzureadProvider#disable_terraform_partner_id}
	DisableTerraformPartnerId interface{} `field:"optional" json:"disableTerraformPartnerId" yaml:"disableTerraformPartnerId"`
	// The cloud environment which should be used.
	//
	// Possible values are: `global` (also `public`), `usgovernmentl4` (also `usgovernment`), `usgovernmentl5` (also `dod`), and `china`. Defaults to `global`
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread#environment AzureadProvider#environment}
	Environment *string `field:"optional" json:"environment" yaml:"environment"`
	// The path to a custom endpoint for Managed Identity - in most circumstances this should be detected automatically.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread#msi_endpoint AzureadProvider#msi_endpoint}
	MsiEndpoint *string `field:"optional" json:"msiEndpoint" yaml:"msiEndpoint"`
	// The bearer token for the request to the OIDC provider.
	//
	// For use when authenticating as a Service Principal using OpenID Connect.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread#oidc_request_token AzureadProvider#oidc_request_token}
	OidcRequestToken *string `field:"optional" json:"oidcRequestToken" yaml:"oidcRequestToken"`
	// The URL for the OIDC provider from which to request an ID token.
	//
	// For use when authenticating as a Service Principal using OpenID Connect.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread#oidc_request_url AzureadProvider#oidc_request_url}
	OidcRequestUrl *string `field:"optional" json:"oidcRequestUrl" yaml:"oidcRequestUrl"`
	// The ID token for use when authenticating as a Service Principal using OpenID Connect.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread#oidc_token AzureadProvider#oidc_token}
	OidcToken *string `field:"optional" json:"oidcToken" yaml:"oidcToken"`
	// A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread#partner_id AzureadProvider#partner_id}
	PartnerId *string `field:"optional" json:"partnerId" yaml:"partnerId"`
	// The Tenant ID which should be used. Works with all authentication methods except Managed Identity.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread#tenant_id AzureadProvider#tenant_id}
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
	// Allow Azure CLI to be used for Authentication.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread#use_cli AzureadProvider#use_cli}
	UseCli interface{} `field:"optional" json:"useCli" yaml:"useCli"`
	// Allow Managed Identity to be used for Authentication.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread#use_msi AzureadProvider#use_msi}
	UseMsi interface{} `field:"optional" json:"useMsi" yaml:"useMsi"`
	// Allow OpenID Connect to be used for authentication.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread#use_oidc AzureadProvider#use_oidc}
	UseOidc interface{} `field:"optional" json:"useOidc" yaml:"useOidc"`
}


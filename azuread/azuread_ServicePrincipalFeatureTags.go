// Prebuilt azuread Provider for Terraform CDK (cdktf)
package azuread


type ServicePrincipalFeatureTags struct {
	// Whether this service principal represents a custom SAML application.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/service_principal#custom_single_sign_on ServicePrincipal#custom_single_sign_on}
	CustomSingleSignOn interface{} `field:"optional" json:"customSingleSignOn" yaml:"customSingleSignOn"`
	// Whether this service principal represents an Enterprise Application.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/service_principal#enterprise ServicePrincipal#enterprise}
	Enterprise interface{} `field:"optional" json:"enterprise" yaml:"enterprise"`
	// Whether this service principal represents a gallery application.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/service_principal#gallery ServicePrincipal#gallery}
	Gallery interface{} `field:"optional" json:"gallery" yaml:"gallery"`
	// Whether this app is invisible to users in My Apps and Office 365 Launcher.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/service_principal#hide ServicePrincipal#hide}
	Hide interface{} `field:"optional" json:"hide" yaml:"hide"`
}

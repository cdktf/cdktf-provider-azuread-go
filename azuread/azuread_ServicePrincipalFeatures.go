// Prebuilt azuread Provider for Terraform CDK (cdktf)
package azuread


type ServicePrincipalFeatures struct {
	// Whether this service principal represents a custom SAML application.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/service_principal#custom_single_sign_on_app ServicePrincipal#custom_single_sign_on_app}
	CustomSingleSignOnApp interface{} `field:"optional" json:"customSingleSignOnApp" yaml:"customSingleSignOnApp"`
	// Whether this service principal represents an Enterprise Application.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/service_principal#enterprise_application ServicePrincipal#enterprise_application}
	EnterpriseApplication interface{} `field:"optional" json:"enterpriseApplication" yaml:"enterpriseApplication"`
	// Whether this service principal represents a gallery application.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/service_principal#gallery_application ServicePrincipal#gallery_application}
	GalleryApplication interface{} `field:"optional" json:"galleryApplication" yaml:"galleryApplication"`
	// Whether this app is visible to users in My Apps and Office 365 Launcher.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/service_principal#visible_to_users ServicePrincipal#visible_to_users}
	VisibleToUsers interface{} `field:"optional" json:"visibleToUsers" yaml:"visibleToUsers"`
}


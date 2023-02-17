package application


type ApplicationAppRole struct {
	// Specifies whether this app role definition can be assigned to users and groups by setting to `User`, or to other applications (that are accessing this application in a standalone scenario) by setting to `Application`, or to both.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/application#allowed_member_types Application#allowed_member_types}
	AllowedMemberTypes *[]*string `field:"required" json:"allowedMemberTypes" yaml:"allowedMemberTypes"`
	// Description of the app role that appears when the role is being assigned and, if the role functions as an application permissions, during the consent experiences.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/application#description Application#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// Display name for the app role that appears during app role assignment and in consent experiences.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/application#display_name Application#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The unique identifier of the app role.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/application#id Application#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Determines if the app role is enabled.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/application#enabled Application#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The value that is used for the `roles` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/application#value Application#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}


package conditionalaccesspolicy


type ConditionalAccessPolicyGrantControls struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.39.0/docs/resources/conditional_access_policy#built_in_controls ConditionalAccessPolicy#built_in_controls}.
	BuiltInControls *[]*string `field:"required" json:"builtInControls" yaml:"builtInControls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.39.0/docs/resources/conditional_access_policy#operator ConditionalAccessPolicy#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.39.0/docs/resources/conditional_access_policy#custom_authentication_factors ConditionalAccessPolicy#custom_authentication_factors}.
	CustomAuthenticationFactors *[]*string `field:"optional" json:"customAuthenticationFactors" yaml:"customAuthenticationFactors"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.39.0/docs/resources/conditional_access_policy#terms_of_use ConditionalAccessPolicy#terms_of_use}.
	TermsOfUse *[]*string `field:"optional" json:"termsOfUse" yaml:"termsOfUse"`
}


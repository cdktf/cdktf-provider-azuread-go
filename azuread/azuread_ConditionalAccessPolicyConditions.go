// Prebuilt azuread Provider for Terraform CDK (cdktf)
package azuread


type ConditionalAccessPolicyConditions struct {
	// applications block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/conditional_access_policy#applications ConditionalAccessPolicy#applications}
	Applications *ConditionalAccessPolicyConditionsApplications `field:"required" json:"applications" yaml:"applications"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/conditional_access_policy#client_app_types ConditionalAccessPolicy#client_app_types}.
	ClientAppTypes *[]*string `field:"required" json:"clientAppTypes" yaml:"clientAppTypes"`
	// users block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/conditional_access_policy#users ConditionalAccessPolicy#users}
	Users *ConditionalAccessPolicyConditionsUsers `field:"required" json:"users" yaml:"users"`
	// devices block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/conditional_access_policy#devices ConditionalAccessPolicy#devices}
	Devices *ConditionalAccessPolicyConditionsDevices `field:"optional" json:"devices" yaml:"devices"`
	// locations block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/conditional_access_policy#locations ConditionalAccessPolicy#locations}
	Locations *ConditionalAccessPolicyConditionsLocations `field:"optional" json:"locations" yaml:"locations"`
	// platforms block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/conditional_access_policy#platforms ConditionalAccessPolicy#platforms}
	Platforms *ConditionalAccessPolicyConditionsPlatforms `field:"optional" json:"platforms" yaml:"platforms"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/conditional_access_policy#sign_in_risk_levels ConditionalAccessPolicy#sign_in_risk_levels}.
	SignInRiskLevels *[]*string `field:"optional" json:"signInRiskLevels" yaml:"signInRiskLevels"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/conditional_access_policy#user_risk_levels ConditionalAccessPolicy#user_risk_levels}.
	UserRiskLevels *[]*string `field:"optional" json:"userRiskLevels" yaml:"userRiskLevels"`
}


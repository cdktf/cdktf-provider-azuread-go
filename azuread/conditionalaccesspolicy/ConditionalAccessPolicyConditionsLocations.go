package conditionalaccesspolicy


type ConditionalAccessPolicyConditionsLocations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.39.0/docs/resources/conditional_access_policy#included_locations ConditionalAccessPolicy#included_locations}.
	IncludedLocations *[]*string `field:"required" json:"includedLocations" yaml:"includedLocations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.39.0/docs/resources/conditional_access_policy#excluded_locations ConditionalAccessPolicy#excluded_locations}.
	ExcludedLocations *[]*string `field:"optional" json:"excludedLocations" yaml:"excludedLocations"`
}


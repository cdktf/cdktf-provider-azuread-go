// Prebuilt azuread Provider for Terraform CDK (cdktf)
package azuread


type ConditionalAccessPolicyConditionsPlatforms struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/conditional_access_policy#included_platforms ConditionalAccessPolicy#included_platforms}.
	IncludedPlatforms *[]*string `field:"required" json:"includedPlatforms" yaml:"includedPlatforms"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/conditional_access_policy#excluded_platforms ConditionalAccessPolicy#excluded_platforms}.
	ExcludedPlatforms *[]*string `field:"optional" json:"excludedPlatforms" yaml:"excludedPlatforms"`
}


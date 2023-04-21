package conditionalaccesspolicy


type ConditionalAccessPolicyConditionsUsers struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.37.2/docs/resources/conditional_access_policy#excluded_groups ConditionalAccessPolicy#excluded_groups}.
	ExcludedGroups *[]*string `field:"optional" json:"excludedGroups" yaml:"excludedGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.37.2/docs/resources/conditional_access_policy#excluded_roles ConditionalAccessPolicy#excluded_roles}.
	ExcludedRoles *[]*string `field:"optional" json:"excludedRoles" yaml:"excludedRoles"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.37.2/docs/resources/conditional_access_policy#excluded_users ConditionalAccessPolicy#excluded_users}.
	ExcludedUsers *[]*string `field:"optional" json:"excludedUsers" yaml:"excludedUsers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.37.2/docs/resources/conditional_access_policy#included_groups ConditionalAccessPolicy#included_groups}.
	IncludedGroups *[]*string `field:"optional" json:"includedGroups" yaml:"includedGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.37.2/docs/resources/conditional_access_policy#included_roles ConditionalAccessPolicy#included_roles}.
	IncludedRoles *[]*string `field:"optional" json:"includedRoles" yaml:"includedRoles"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.37.2/docs/resources/conditional_access_policy#included_users ConditionalAccessPolicy#included_users}.
	IncludedUsers *[]*string `field:"optional" json:"includedUsers" yaml:"includedUsers"`
}


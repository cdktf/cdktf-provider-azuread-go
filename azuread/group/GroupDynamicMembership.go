package group


type GroupDynamicMembership struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/group#enabled Group#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Rule to determine members for a dynamic group. Required when `group_types` contains 'DynamicMembership'.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/group#rule Group#rule}
	Rule *string `field:"required" json:"rule" yaml:"rule"`
}


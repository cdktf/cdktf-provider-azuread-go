package conditionalaccesspolicy


type ConditionalAccessPolicyConditionsDevicesFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.39.0/docs/resources/conditional_access_policy#mode ConditionalAccessPolicy#mode}.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.39.0/docs/resources/conditional_access_policy#rule ConditionalAccessPolicy#rule}.
	Rule *string `field:"required" json:"rule" yaml:"rule"`
}


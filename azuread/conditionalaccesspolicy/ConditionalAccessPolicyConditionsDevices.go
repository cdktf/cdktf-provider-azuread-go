package conditionalaccesspolicy


type ConditionalAccessPolicyConditionsDevices struct {
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.40.0/docs/resources/conditional_access_policy#filter ConditionalAccessPolicy#filter}
	Filter *ConditionalAccessPolicyConditionsDevicesFilter `field:"optional" json:"filter" yaml:"filter"`
}


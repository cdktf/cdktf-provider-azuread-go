// Prebuilt azuread Provider for Terraform CDK (cdktf)
package azuread


type ConditionalAccessPolicyConditionsDevices struct {
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/conditional_access_policy#filter ConditionalAccessPolicy#filter}
	Filter *ConditionalAccessPolicyConditionsDevicesFilter `field:"optional" json:"filter" yaml:"filter"`
}


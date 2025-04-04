// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package conditionalaccesspolicy


type ConditionalAccessPolicyConditionsDevices struct {
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.3.0/docs/resources/conditional_access_policy#filter ConditionalAccessPolicy#filter}
	Filter *ConditionalAccessPolicyConditionsDevicesFilter `field:"optional" json:"filter" yaml:"filter"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package conditionalaccesspolicy


type ConditionalAccessPolicyConditionsClientApplications struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.5.0/docs/resources/conditional_access_policy#excluded_service_principals ConditionalAccessPolicy#excluded_service_principals}.
	ExcludedServicePrincipals *[]*string `field:"optional" json:"excludedServicePrincipals" yaml:"excludedServicePrincipals"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.5.0/docs/resources/conditional_access_policy#filter ConditionalAccessPolicy#filter}
	Filter *ConditionalAccessPolicyConditionsClientApplicationsFilter `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.5.0/docs/resources/conditional_access_policy#included_service_principals ConditionalAccessPolicy#included_service_principals}.
	IncludedServicePrincipals *[]*string `field:"optional" json:"includedServicePrincipals" yaml:"includedServicePrincipals"`
}


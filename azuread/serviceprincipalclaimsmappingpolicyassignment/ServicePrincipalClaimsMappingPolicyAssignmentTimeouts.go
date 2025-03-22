// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceprincipalclaimsmappingpolicyassignment


type ServicePrincipalClaimsMappingPolicyAssignmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.2.0/docs/resources/service_principal_claims_mapping_policy_assignment#create ServicePrincipalClaimsMappingPolicyAssignment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.2.0/docs/resources/service_principal_claims_mapping_policy_assignment#delete ServicePrincipalClaimsMappingPolicyAssignment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.2.0/docs/resources/service_principal_claims_mapping_policy_assignment#read ServicePrincipalClaimsMappingPolicyAssignment#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}


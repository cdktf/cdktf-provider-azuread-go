// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationfallbackpublicclient


type ApplicationFallbackPublicClientTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.53.1/docs/resources/application_fallback_public_client#create ApplicationFallbackPublicClient#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.53.1/docs/resources/application_fallback_public_client#delete ApplicationFallbackPublicClient#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.53.1/docs/resources/application_fallback_public_client#read ApplicationFallbackPublicClient#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}


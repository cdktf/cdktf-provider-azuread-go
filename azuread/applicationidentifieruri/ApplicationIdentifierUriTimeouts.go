// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationidentifieruri


type ApplicationIdentifierUriTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.46.0/docs/resources/application_identifier_uri#create ApplicationIdentifierUri#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.46.0/docs/resources/application_identifier_uri#delete ApplicationIdentifierUri#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.46.0/docs/resources/application_identifier_uri#read ApplicationIdentifierUri#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataazureadadministrativeunit


type DataAzureadAdministrativeUnitTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.5.0/docs/data-sources/administrative_unit#create DataAzureadAdministrativeUnit#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.5.0/docs/data-sources/administrative_unit#delete DataAzureadAdministrativeUnit#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.5.0/docs/data-sources/administrative_unit#read DataAzureadAdministrativeUnit#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.5.0/docs/data-sources/administrative_unit#update DataAzureadAdministrativeUnit#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


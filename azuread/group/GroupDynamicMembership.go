// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package group


type GroupDynamicMembership struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.42.0/docs/resources/group#enabled Group#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Rule to determine members for a dynamic group. Required when `group_types` contains 'DynamicMembership'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.42.0/docs/resources/group#rule Group#rule}
	Rule *string `field:"required" json:"rule" yaml:"rule"`
}


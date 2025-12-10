// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package groupwithoutmembers


type GroupWithoutMembersDynamicMembership struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.7.0/docs/resources/group_without_members#enabled GroupWithoutMembers#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Rule to determine members for a dynamic group. Required when `group_types` contains 'DynamicMembership'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.7.0/docs/resources/group_without_members#rule GroupWithoutMembers#rule}
	Rule *string `field:"required" json:"rule" yaml:"rule"`
}


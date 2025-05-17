// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package namedlocation


type NamedLocationIp struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/named_location#ip_ranges NamedLocation#ip_ranges}.
	IpRanges *[]*string `field:"required" json:"ipRanges" yaml:"ipRanges"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/named_location#trusted NamedLocation#trusted}.
	Trusted interface{} `field:"optional" json:"trusted" yaml:"trusted"`
}


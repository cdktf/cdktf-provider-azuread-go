// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package namedlocation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NamedLocationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/named_location#display_name NamedLocation#display_name}.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// country block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/named_location#country NamedLocation#country}
	Country *NamedLocationCountry `field:"optional" json:"country" yaml:"country"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/named_location#id NamedLocation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// ip block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/named_location#ip NamedLocation#ip}
	Ip *NamedLocationIp `field:"optional" json:"ip" yaml:"ip"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/named_location#timeouts NamedLocation#timeouts}
	Timeouts *NamedLocationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


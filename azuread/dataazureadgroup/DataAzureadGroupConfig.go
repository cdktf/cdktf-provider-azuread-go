// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataazureadgroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAzureadGroupConfig struct {
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
	// The display name for the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.48.0/docs/data-sources/group#display_name DataAzureadGroup#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.48.0/docs/data-sources/group#id DataAzureadGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether the group is mail-enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.48.0/docs/data-sources/group#mail_enabled DataAzureadGroup#mail_enabled}
	MailEnabled interface{} `field:"optional" json:"mailEnabled" yaml:"mailEnabled"`
	// The mail alias for the group, unique in the organisation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.48.0/docs/data-sources/group#mail_nickname DataAzureadGroup#mail_nickname}
	MailNickname *string `field:"optional" json:"mailNickname" yaml:"mailNickname"`
	// The object ID of the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.48.0/docs/data-sources/group#object_id DataAzureadGroup#object_id}
	ObjectId *string `field:"optional" json:"objectId" yaml:"objectId"`
	// Whether the group is a security group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.48.0/docs/data-sources/group#security_enabled DataAzureadGroup#security_enabled}
	SecurityEnabled interface{} `field:"optional" json:"securityEnabled" yaml:"securityEnabled"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.48.0/docs/data-sources/group#timeouts DataAzureadGroup#timeouts}
	Timeouts *DataAzureadGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


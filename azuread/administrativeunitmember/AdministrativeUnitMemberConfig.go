// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package administrativeunitmember

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AdministrativeUnitMemberConfig struct {
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
	// The object ID of the administrative unit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.1.0/docs/resources/administrative_unit_member#administrative_unit_object_id AdministrativeUnitMember#administrative_unit_object_id}
	AdministrativeUnitObjectId *string `field:"optional" json:"administrativeUnitObjectId" yaml:"administrativeUnitObjectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.1.0/docs/resources/administrative_unit_member#id AdministrativeUnitMember#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The object ID of the member.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.1.0/docs/resources/administrative_unit_member#member_object_id AdministrativeUnitMember#member_object_id}
	MemberObjectId *string `field:"optional" json:"memberObjectId" yaml:"memberObjectId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.1.0/docs/resources/administrative_unit_member#timeouts AdministrativeUnitMember#timeouts}
	Timeouts *AdministrativeUnitMemberTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package administrativeunit

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AdministrativeUnitConfig struct {
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
	// The display name for the administrative unit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.49.0/docs/resources/administrative_unit#display_name AdministrativeUnit#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The description for the administrative unit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.49.0/docs/resources/administrative_unit#description AdministrativeUnit#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether the administrative unit and its members are hidden or publicly viewable in the directory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.49.0/docs/resources/administrative_unit#hidden_membership_enabled AdministrativeUnit#hidden_membership_enabled}
	HiddenMembershipEnabled interface{} `field:"optional" json:"hiddenMembershipEnabled" yaml:"hiddenMembershipEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.49.0/docs/resources/administrative_unit#id AdministrativeUnit#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A set of object IDs of members who should be present in this administrative unit.
	//
	// Supported object types are Users or Groups
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.49.0/docs/resources/administrative_unit#members AdministrativeUnit#members}
	Members *[]*string `field:"optional" json:"members" yaml:"members"`
	// If `true`, will return an error if an existing administrative unit is found with the same name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.49.0/docs/resources/administrative_unit#prevent_duplicate_names AdministrativeUnit#prevent_duplicate_names}
	PreventDuplicateNames interface{} `field:"optional" json:"preventDuplicateNames" yaml:"preventDuplicateNames"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.49.0/docs/resources/administrative_unit#timeouts AdministrativeUnit#timeouts}
	Timeouts *AdministrativeUnitTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


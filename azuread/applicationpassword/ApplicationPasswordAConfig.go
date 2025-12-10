// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationpassword

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApplicationPasswordAConfig struct {
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
	// The resource ID of the application for which this password should be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.7.0/docs/resources/application_password#application_id ApplicationPasswordA#application_id}
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// A display name for the password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.7.0/docs/resources/application_password#display_name ApplicationPasswordA#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The end date until which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.7.0/docs/resources/application_password#end_date ApplicationPasswordA#end_date}
	EndDate *string `field:"optional" json:"endDate" yaml:"endDate"`
	// A relative duration for which the password is valid until, for example `240h` (10 days) or `2400h30m`.
	//
	// Changing this field forces a new resource to be created
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.7.0/docs/resources/application_password#end_date_relative ApplicationPasswordA#end_date_relative}
	EndDateRelative *string `field:"optional" json:"endDateRelative" yaml:"endDateRelative"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.7.0/docs/resources/application_password#id ApplicationPasswordA#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Arbitrary map of values that, when changed, will trigger rotation of the password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.7.0/docs/resources/application_password#rotate_when_changed ApplicationPasswordA#rotate_when_changed}
	RotateWhenChanged *map[string]*string `field:"optional" json:"rotateWhenChanged" yaml:"rotateWhenChanged"`
	// The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.7.0/docs/resources/application_password#start_date ApplicationPasswordA#start_date}
	StartDate *string `field:"optional" json:"startDate" yaml:"startDate"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.7.0/docs/resources/application_password#timeouts ApplicationPasswordA#timeouts}
	Timeouts *ApplicationPasswordTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


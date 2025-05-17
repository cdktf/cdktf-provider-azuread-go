// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package application


type ApplicationPassword struct {
	// A display name for the password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/application#display_name Application#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The end date until which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/application#end_date Application#end_date}
	EndDate *string `field:"optional" json:"endDate" yaml:"endDate"`
	// The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/application#start_date Application#start_date}
	StartDate *string `field:"optional" json:"startDate" yaml:"startDate"`
}


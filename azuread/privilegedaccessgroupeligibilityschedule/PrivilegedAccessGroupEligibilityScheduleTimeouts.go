// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privilegedaccessgroupeligibilityschedule


type PrivilegedAccessGroupEligibilityScheduleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.6.0/docs/resources/privileged_access_group_eligibility_schedule#create PrivilegedAccessGroupEligibilitySchedule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.6.0/docs/resources/privileged_access_group_eligibility_schedule#delete PrivilegedAccessGroupEligibilitySchedule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.6.0/docs/resources/privileged_access_group_eligibility_schedule#read PrivilegedAccessGroupEligibilitySchedule#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.6.0/docs/resources/privileged_access_group_eligibility_schedule#update PrivilegedAccessGroupEligibilitySchedule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


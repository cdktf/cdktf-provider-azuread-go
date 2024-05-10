// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privilegedaccessgroupassignmentschedule


type PrivilegedAccessGroupAssignmentScheduleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.49.0/docs/resources/privileged_access_group_assignment_schedule#create PrivilegedAccessGroupAssignmentSchedule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.49.0/docs/resources/privileged_access_group_assignment_schedule#delete PrivilegedAccessGroupAssignmentSchedule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.49.0/docs/resources/privileged_access_group_assignment_schedule#read PrivilegedAccessGroupAssignmentSchedule#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.49.0/docs/resources/privileged_access_group_assignment_schedule#update PrivilegedAccessGroupAssignmentSchedule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


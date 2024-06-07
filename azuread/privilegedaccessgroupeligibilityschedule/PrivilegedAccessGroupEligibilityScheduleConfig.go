// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privilegedaccessgroupeligibilityschedule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PrivilegedAccessGroupEligibilityScheduleConfig struct {
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
	// The ID of the assignment to the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/privileged_access_group_eligibility_schedule#assignment_type PrivilegedAccessGroupEligibilitySchedule#assignment_type}
	AssignmentType *string `field:"required" json:"assignmentType" yaml:"assignmentType"`
	// The ID of the Group representing the scope of the assignment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/privileged_access_group_eligibility_schedule#group_id PrivilegedAccessGroupEligibilitySchedule#group_id}
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
	// The ID of the Principal assigned to the schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/privileged_access_group_eligibility_schedule#principal_id PrivilegedAccessGroupEligibilitySchedule#principal_id}
	PrincipalId *string `field:"required" json:"principalId" yaml:"principalId"`
	// The duration of the assignment, formatted as an ISO8601 duration string (e.g. P3D for 3 days).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/privileged_access_group_eligibility_schedule#duration PrivilegedAccessGroupEligibilitySchedule#duration}
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
	// The date that this assignment expires, formatted as an RFC3339 date string in UTC (e.g. 2018-01-01T01:02:03Z).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/privileged_access_group_eligibility_schedule#expiration_date PrivilegedAccessGroupEligibilitySchedule#expiration_date}
	ExpirationDate *string `field:"optional" json:"expirationDate" yaml:"expirationDate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/privileged_access_group_eligibility_schedule#id PrivilegedAccessGroupEligibilitySchedule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The justification for the assignment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/privileged_access_group_eligibility_schedule#justification PrivilegedAccessGroupEligibilitySchedule#justification}
	Justification *string `field:"optional" json:"justification" yaml:"justification"`
	// Is the assignment permanent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/privileged_access_group_eligibility_schedule#permanent_assignment PrivilegedAccessGroupEligibilitySchedule#permanent_assignment}
	PermanentAssignment interface{} `field:"optional" json:"permanentAssignment" yaml:"permanentAssignment"`
	// The date that this assignment starts, formatted as an RFC3339 date string in UTC (e.g. 2018-01-01T01:02:03Z).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/privileged_access_group_eligibility_schedule#start_date PrivilegedAccessGroupEligibilitySchedule#start_date}
	StartDate *string `field:"optional" json:"startDate" yaml:"startDate"`
	// The ticket number authorising the assignment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/privileged_access_group_eligibility_schedule#ticket_number PrivilegedAccessGroupEligibilitySchedule#ticket_number}
	TicketNumber *string `field:"optional" json:"ticketNumber" yaml:"ticketNumber"`
	// The ticket system authorising the assignment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/privileged_access_group_eligibility_schedule#ticket_system PrivilegedAccessGroupEligibilitySchedule#ticket_system}
	TicketSystem *string `field:"optional" json:"ticketSystem" yaml:"ticketSystem"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/privileged_access_group_eligibility_schedule#timeouts PrivilegedAccessGroupEligibilitySchedule#timeouts}
	Timeouts *PrivilegedAccessGroupEligibilityScheduleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


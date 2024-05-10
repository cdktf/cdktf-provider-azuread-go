// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package synchronizationjobprovisionondemand


type SynchronizationJobProvisionOnDemandParameter struct {
	// The identifier of the synchronization rule to be applied.
	//
	// This rule ID is defined in the schema for a given synchronization job or template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.49.0/docs/resources/synchronization_job_provision_on_demand#rule_id SynchronizationJobProvisionOnDemand#rule_id}
	RuleId *string `field:"required" json:"ruleId" yaml:"ruleId"`
	// subject block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.49.0/docs/resources/synchronization_job_provision_on_demand#subject SynchronizationJobProvisionOnDemand#subject}
	Subject interface{} `field:"required" json:"subject" yaml:"subject"`
}


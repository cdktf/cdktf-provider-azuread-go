// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package synchronizationjobprovisionondemand


type SynchronizationJobProvisionOnDemandTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.50.0/docs/resources/synchronization_job_provision_on_demand#create SynchronizationJobProvisionOnDemand#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.50.0/docs/resources/synchronization_job_provision_on_demand#delete SynchronizationJobProvisionOnDemand#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.50.0/docs/resources/synchronization_job_provision_on_demand#read SynchronizationJobProvisionOnDemand#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}


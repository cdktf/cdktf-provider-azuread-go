// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package synchronizationjob


type SynchronizationJobTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.1.0/docs/resources/synchronization_job#create SynchronizationJob#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.1.0/docs/resources/synchronization_job#delete SynchronizationJob#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.1.0/docs/resources/synchronization_job#read SynchronizationJob#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.1.0/docs/resources/synchronization_job#update SynchronizationJob#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


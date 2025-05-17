// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationknownclients

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApplicationKnownClientsConfig struct {
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
	// The resource ID of the application to which this API access is granted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/application_known_clients#application_id ApplicationKnownClients#application_id}
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// A list of known client IDs, used for bundling consent if you have a solution that includes an API and a client application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/application_known_clients#known_client_ids ApplicationKnownClients#known_client_ids}
	KnownClientIds *[]*string `field:"required" json:"knownClientIds" yaml:"knownClientIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/application_known_clients#id ApplicationKnownClients#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/application_known_clients#timeouts ApplicationKnownClients#timeouts}
	Timeouts *ApplicationKnownClientsTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


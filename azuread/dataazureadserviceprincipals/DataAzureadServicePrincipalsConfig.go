// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataazureadserviceprincipals

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAzureadServicePrincipalsConfig struct {
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
	// The client IDs of the applications associated with the service principals.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/data-sources/service_principals#client_ids DataAzureadServicePrincipals#client_ids}
	ClientIds *[]*string `field:"optional" json:"clientIds" yaml:"clientIds"`
	// The display names of the applications associated with the service principals.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/data-sources/service_principals#display_names DataAzureadServicePrincipals#display_names}
	DisplayNames *[]*string `field:"optional" json:"displayNames" yaml:"displayNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/data-sources/service_principals#id DataAzureadServicePrincipals#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Ignore missing service principals and return the service principals that were found.
	//
	// The data source will still fail if no service principals are found
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/data-sources/service_principals#ignore_missing DataAzureadServicePrincipals#ignore_missing}
	IgnoreMissing interface{} `field:"optional" json:"ignoreMissing" yaml:"ignoreMissing"`
	// The object IDs of the service principals.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/data-sources/service_principals#object_ids DataAzureadServicePrincipals#object_ids}
	ObjectIds *[]*string `field:"optional" json:"objectIds" yaml:"objectIds"`
	// Fetch all service principals with no filter and return all that were found.
	//
	// The data source will still fail if no service principals are found.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/data-sources/service_principals#return_all DataAzureadServicePrincipals#return_all}
	ReturnAll interface{} `field:"optional" json:"returnAll" yaml:"returnAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/data-sources/service_principals#timeouts DataAzureadServicePrincipals#timeouts}
	Timeouts *DataAzureadServicePrincipalsTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


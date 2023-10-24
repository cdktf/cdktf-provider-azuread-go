// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesspackagecatalog

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccessPackageCatalogConfig struct {
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
	// The description of the access package catalog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/access_package_catalog#description AccessPackageCatalog#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// The display name of the access package catalog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/access_package_catalog#display_name AccessPackageCatalog#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Whether the access packages in this catalog can be requested by users outside the tenant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/access_package_catalog#externally_visible AccessPackageCatalog#externally_visible}
	ExternallyVisible interface{} `field:"optional" json:"externallyVisible" yaml:"externallyVisible"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/access_package_catalog#id AccessPackageCatalog#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether the access packages in this catalog are available for management.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/access_package_catalog#published AccessPackageCatalog#published}
	Published interface{} `field:"optional" json:"published" yaml:"published"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/access_package_catalog#timeouts AccessPackageCatalog#timeouts}
	Timeouts *AccessPackageCatalogTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


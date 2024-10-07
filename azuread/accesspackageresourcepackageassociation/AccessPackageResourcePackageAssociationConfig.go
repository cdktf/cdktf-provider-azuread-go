// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesspackageresourcepackageassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccessPackageResourcePackageAssociationConfig struct {
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
	// The ID of access package this resource association is configured to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.0.2/docs/resources/access_package_resource_package_association#access_package_id AccessPackageResourcePackageAssociation#access_package_id}
	AccessPackageId *string `field:"required" json:"accessPackageId" yaml:"accessPackageId"`
	// The ID of the access package catalog association.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.0.2/docs/resources/access_package_resource_package_association#catalog_resource_association_id AccessPackageResourcePackageAssociation#catalog_resource_association_id}
	CatalogResourceAssociationId *string `field:"required" json:"catalogResourceAssociationId" yaml:"catalogResourceAssociationId"`
	// The role of access type to the specified resource, valid values are `Member` and `Owner`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.0.2/docs/resources/access_package_resource_package_association#access_type AccessPackageResourcePackageAssociation#access_type}
	AccessType *string `field:"optional" json:"accessType" yaml:"accessType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.0.2/docs/resources/access_package_resource_package_association#id AccessPackageResourcePackageAssociation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.0.2/docs/resources/access_package_resource_package_association#timeouts AccessPackageResourcePackageAssociation#timeouts}
	Timeouts *AccessPackageResourcePackageAssociationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


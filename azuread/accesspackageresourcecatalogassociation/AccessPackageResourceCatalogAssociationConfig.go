// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesspackageresourcecatalogassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccessPackageResourceCatalogAssociationConfig struct {
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
	// The unique ID of the access package catalog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/access_package_resource_catalog_association#catalog_id AccessPackageResourceCatalogAssociation#catalog_id}
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// The unique identifier of the resource in the origin system.
	//
	// In the case of an Azure AD group, this is the identifier of the group
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/access_package_resource_catalog_association#resource_origin_id AccessPackageResourceCatalogAssociation#resource_origin_id}
	ResourceOriginId *string `field:"required" json:"resourceOriginId" yaml:"resourceOriginId"`
	// The type of the resource in the origin system, such as SharePointOnline, AadApplication or AadGroup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/access_package_resource_catalog_association#resource_origin_system AccessPackageResourceCatalogAssociation#resource_origin_system}
	ResourceOriginSystem *string `field:"required" json:"resourceOriginSystem" yaml:"resourceOriginSystem"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/access_package_resource_catalog_association#id AccessPackageResourceCatalogAssociation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/access_package_resource_catalog_association#timeouts AccessPackageResourceCatalogAssociation#timeouts}
	Timeouts *AccessPackageResourceCatalogAssociationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


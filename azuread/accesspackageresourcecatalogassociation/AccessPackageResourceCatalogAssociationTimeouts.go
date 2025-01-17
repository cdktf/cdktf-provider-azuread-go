// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesspackageresourcecatalogassociation


type AccessPackageResourceCatalogAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.1.0/docs/resources/access_package_resource_catalog_association#create AccessPackageResourceCatalogAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.1.0/docs/resources/access_package_resource_catalog_association#delete AccessPackageResourceCatalogAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.1.0/docs/resources/access_package_resource_catalog_association#read AccessPackageResourceCatalogAssociation#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesspackageresourcepackageassociation


type AccessPackageResourcePackageAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/access_package_resource_package_association#create AccessPackageResourcePackageAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/access_package_resource_package_association#delete AccessPackageResourcePackageAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/access_package_resource_package_association#read AccessPackageResourcePackageAssociation#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}


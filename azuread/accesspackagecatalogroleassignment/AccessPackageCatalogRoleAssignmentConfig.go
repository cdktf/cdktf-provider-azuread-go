// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesspackagecatalogroleassignment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccessPackageCatalogRoleAssignmentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.50.0/docs/resources/access_package_catalog_role_assignment#catalog_id AccessPackageCatalogRoleAssignment#catalog_id}
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// The object ID of the member principal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.50.0/docs/resources/access_package_catalog_role_assignment#principal_object_id AccessPackageCatalogRoleAssignment#principal_object_id}
	PrincipalObjectId *string `field:"required" json:"principalObjectId" yaml:"principalObjectId"`
	// The object ID of the catalog role for this assignment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.50.0/docs/resources/access_package_catalog_role_assignment#role_id AccessPackageCatalogRoleAssignment#role_id}
	RoleId *string `field:"required" json:"roleId" yaml:"roleId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.50.0/docs/resources/access_package_catalog_role_assignment#id AccessPackageCatalogRoleAssignment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.50.0/docs/resources/access_package_catalog_role_assignment#timeouts AccessPackageCatalogRoleAssignment#timeouts}
	Timeouts *AccessPackageCatalogRoleAssignmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


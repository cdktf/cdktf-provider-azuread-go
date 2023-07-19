package accesspackagecatalogroleassignment


type AccessPackageCatalogRoleAssignmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.40.0/docs/resources/access_package_catalog_role_assignment#create AccessPackageCatalogRoleAssignment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.40.0/docs/resources/access_package_catalog_role_assignment#delete AccessPackageCatalogRoleAssignment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.40.0/docs/resources/access_package_catalog_role_assignment#read AccessPackageCatalogRoleAssignment#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.40.0/docs/resources/access_package_catalog_role_assignment#update AccessPackageCatalogRoleAssignment#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


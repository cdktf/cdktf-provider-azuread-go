package accesspackageresourcepackageassociation


type AccessPackageResourcePackageAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/access_package_resource_package_association#create AccessPackageResourcePackageAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/access_package_resource_package_association#delete AccessPackageResourcePackageAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/access_package_resource_package_association#read AccessPackageResourcePackageAssociation#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}


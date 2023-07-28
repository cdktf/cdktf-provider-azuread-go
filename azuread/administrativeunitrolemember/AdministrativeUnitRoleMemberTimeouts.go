package administrativeunitrolemember


type AdministrativeUnitRoleMemberTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.41.0/docs/resources/administrative_unit_role_member#create AdministrativeUnitRoleMember#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.41.0/docs/resources/administrative_unit_role_member#delete AdministrativeUnitRoleMember#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.41.0/docs/resources/administrative_unit_role_member#read AdministrativeUnitRoleMember#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.41.0/docs/resources/administrative_unit_role_member#update AdministrativeUnitRoleMember#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


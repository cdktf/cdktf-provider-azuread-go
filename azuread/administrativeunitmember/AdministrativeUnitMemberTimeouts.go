package administrativeunitmember


type AdministrativeUnitMemberTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.37.2/docs/resources/administrative_unit_member#create AdministrativeUnitMember#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.37.2/docs/resources/administrative_unit_member#delete AdministrativeUnitMember#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.37.2/docs/resources/administrative_unit_member#read AdministrativeUnitMember#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.37.2/docs/resources/administrative_unit_member#update AdministrativeUnitMember#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


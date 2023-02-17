package invitation


type InvitationTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/invitation#create Invitation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/invitation#delete Invitation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/invitation#read Invitation#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/invitation#update Invitation#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


package invitation


type InvitationMessage struct {
	// Email addresses of additional recipients the invitation message should be sent to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.37.1/docs/resources/invitation#additional_recipients Invitation#additional_recipients}
	AdditionalRecipients *[]*string `field:"optional" json:"additionalRecipients" yaml:"additionalRecipients"`
	// Customized message body you want to send if you don't want to send the default message.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.37.1/docs/resources/invitation#body Invitation#body}
	Body *string `field:"optional" json:"body" yaml:"body"`
	// The language you want to send the default message in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.37.1/docs/resources/invitation#language Invitation#language}
	Language *string `field:"optional" json:"language" yaml:"language"`
}


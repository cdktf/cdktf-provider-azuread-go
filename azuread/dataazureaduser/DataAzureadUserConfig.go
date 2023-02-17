package dataazureaduser

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAzureadUserConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/d/user#id DataAzureadUser#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The email alias of the user.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/d/user#mail_nickname DataAzureadUser#mail_nickname}
	MailNickname *string `field:"optional" json:"mailNickname" yaml:"mailNickname"`
	// The object ID of the user.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/d/user#object_id DataAzureadUser#object_id}
	ObjectId *string `field:"optional" json:"objectId" yaml:"objectId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/d/user#timeouts DataAzureadUser#timeouts}
	Timeouts *DataAzureadUserTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The user principal name (UPN) of the user.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/d/user#user_principal_name DataAzureadUser#user_principal_name}
	UserPrincipalName *string `field:"optional" json:"userPrincipalName" yaml:"userPrincipalName"`
}


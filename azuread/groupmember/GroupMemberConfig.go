package groupmember

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GroupMemberConfig struct {
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
	// The object ID of the group you want to add the member to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.38.0/docs/resources/group_member#group_object_id GroupMember#group_object_id}
	GroupObjectId *string `field:"required" json:"groupObjectId" yaml:"groupObjectId"`
	// The object ID of the principal you want to add as a member to the group.
	//
	// Supported object types are Users, Groups or Service Principals
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.38.0/docs/resources/group_member#member_object_id GroupMember#member_object_id}
	MemberObjectId *string `field:"required" json:"memberObjectId" yaml:"memberObjectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.38.0/docs/resources/group_member#id GroupMember#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.38.0/docs/resources/group_member#timeouts GroupMember#timeouts}
	Timeouts *GroupMemberTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


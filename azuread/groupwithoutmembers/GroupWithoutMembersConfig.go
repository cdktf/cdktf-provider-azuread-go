// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package groupwithoutmembers

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GroupWithoutMembersConfig struct {
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
	// The display name for the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/group_without_members#display_name GroupWithoutMembers#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The administrative unit IDs in which the group should be.
	//
	// If empty, the group will be created at the tenant level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/group_without_members#administrative_unit_ids GroupWithoutMembers#administrative_unit_ids}
	AdministrativeUnitIds *[]*string `field:"optional" json:"administrativeUnitIds" yaml:"administrativeUnitIds"`
	// Indicates whether this group can be assigned to an Azure Active Directory role.
	//
	// This property can only be `true` for security-enabled groups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/group_without_members#assignable_to_role GroupWithoutMembers#assignable_to_role}
	AssignableToRole interface{} `field:"optional" json:"assignableToRole" yaml:"assignableToRole"`
	// Indicates whether new members added to the group will be auto-subscribed to receive email notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/group_without_members#auto_subscribe_new_members GroupWithoutMembers#auto_subscribe_new_members}
	AutoSubscribeNewMembers interface{} `field:"optional" json:"autoSubscribeNewMembers" yaml:"autoSubscribeNewMembers"`
	// The group behaviours for a Microsoft 365 group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/group_without_members#behaviors GroupWithoutMembers#behaviors}
	Behaviors *[]*string `field:"optional" json:"behaviors" yaml:"behaviors"`
	// The description for the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/group_without_members#description GroupWithoutMembers#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// dynamic_membership block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/group_without_members#dynamic_membership GroupWithoutMembers#dynamic_membership}
	DynamicMembership *GroupWithoutMembersDynamicMembership `field:"optional" json:"dynamicMembership" yaml:"dynamicMembership"`
	// Indicates whether people external to the organization can send messages to the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/group_without_members#external_senders_allowed GroupWithoutMembers#external_senders_allowed}
	ExternalSendersAllowed interface{} `field:"optional" json:"externalSendersAllowed" yaml:"externalSendersAllowed"`
	// Indicates whether the group is displayed in certain parts of the Outlook user interface: in the Address Book, in address lists for selecting message recipients, and in the Browse Groups dialog for searching groups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/group_without_members#hide_from_address_lists GroupWithoutMembers#hide_from_address_lists}
	HideFromAddressLists interface{} `field:"optional" json:"hideFromAddressLists" yaml:"hideFromAddressLists"`
	// Indicates whether the group is displayed in Outlook clients, such as Outlook for Windows and Outlook on the web.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/group_without_members#hide_from_outlook_clients GroupWithoutMembers#hide_from_outlook_clients}
	HideFromOutlookClients interface{} `field:"optional" json:"hideFromOutlookClients" yaml:"hideFromOutlookClients"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/group_without_members#id GroupWithoutMembers#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether the group is a mail enabled, with a shared group mailbox.
	//
	// At least one of `mail_enabled` or `security_enabled` must be specified. A group can be mail enabled _and_ security enabled
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/group_without_members#mail_enabled GroupWithoutMembers#mail_enabled}
	MailEnabled interface{} `field:"optional" json:"mailEnabled" yaml:"mailEnabled"`
	// The mail alias for the group, unique in the organisation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/group_without_members#mail_nickname GroupWithoutMembers#mail_nickname}
	MailNickname *string `field:"optional" json:"mailNickname" yaml:"mailNickname"`
	// Indicates the target on-premise group type the group will be written back as.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/group_without_members#onpremises_group_type GroupWithoutMembers#onpremises_group_type}
	OnpremisesGroupType *string `field:"optional" json:"onpremisesGroupType" yaml:"onpremisesGroupType"`
	// A set of owners who own this group. Supported object types are Users or Service Principals.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/group_without_members#owners GroupWithoutMembers#owners}
	Owners *[]*string `field:"optional" json:"owners" yaml:"owners"`
	// If `true`, will return an error if an existing group is found with the same name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/group_without_members#prevent_duplicate_names GroupWithoutMembers#prevent_duplicate_names}
	PreventDuplicateNames interface{} `field:"optional" json:"preventDuplicateNames" yaml:"preventDuplicateNames"`
	// The group provisioning options for a Microsoft 365 group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/group_without_members#provisioning_options GroupWithoutMembers#provisioning_options}
	ProvisioningOptions *[]*string `field:"optional" json:"provisioningOptions" yaml:"provisioningOptions"`
	// Whether the group is a security group for controlling access to in-app resources.
	//
	// At least one of `security_enabled` or `mail_enabled` must be specified. A group can be security enabled _and_ mail enabled
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/group_without_members#security_enabled GroupWithoutMembers#security_enabled}
	SecurityEnabled interface{} `field:"optional" json:"securityEnabled" yaml:"securityEnabled"`
	// The colour theme for a Microsoft 365 group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/group_without_members#theme GroupWithoutMembers#theme}
	Theme *string `field:"optional" json:"theme" yaml:"theme"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/group_without_members#timeouts GroupWithoutMembers#timeouts}
	Timeouts *GroupWithoutMembersTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// A set of group types to configure for the group.
	//
	// `Unified` specifies a Microsoft 365 group. Required when `mail_enabled` is true
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/group_without_members#types GroupWithoutMembers#types}
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
	// Specifies the group join policy and group content visibility.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/group_without_members#visibility GroupWithoutMembers#visibility}
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
	// Whether this group should be synced from Azure AD to the on-premises directory when Azure AD Connect is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/group_without_members#writeback_enabled GroupWithoutMembers#writeback_enabled}
	WritebackEnabled interface{} `field:"optional" json:"writebackEnabled" yaml:"writebackEnabled"`
}


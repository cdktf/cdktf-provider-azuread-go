// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package group

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-azuread.group.Group",
		reflect.TypeOf((*Group)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "administrativeUnitIds", GoGetter: "AdministrativeUnitIds"},
			_jsii_.MemberProperty{JsiiProperty: "administrativeUnitIdsInput", GoGetter: "AdministrativeUnitIdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "assignableToRole", GoGetter: "AssignableToRole"},
			_jsii_.MemberProperty{JsiiProperty: "assignableToRoleInput", GoGetter: "AssignableToRoleInput"},
			_jsii_.MemberProperty{JsiiProperty: "autoSubscribeNewMembers", GoGetter: "AutoSubscribeNewMembers"},
			_jsii_.MemberProperty{JsiiProperty: "autoSubscribeNewMembersInput", GoGetter: "AutoSubscribeNewMembersInput"},
			_jsii_.MemberProperty{JsiiProperty: "behaviors", GoGetter: "Behaviors"},
			_jsii_.MemberProperty{JsiiProperty: "behaviorsInput", GoGetter: "BehaviorsInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "displayName", GoGetter: "DisplayName"},
			_jsii_.MemberProperty{JsiiProperty: "displayNameInput", GoGetter: "DisplayNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "dynamicMembership", GoGetter: "DynamicMembership"},
			_jsii_.MemberProperty{JsiiProperty: "dynamicMembershipInput", GoGetter: "DynamicMembershipInput"},
			_jsii_.MemberProperty{JsiiProperty: "externalSendersAllowed", GoGetter: "ExternalSendersAllowed"},
			_jsii_.MemberProperty{JsiiProperty: "externalSendersAllowedInput", GoGetter: "ExternalSendersAllowedInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "hideFromAddressLists", GoGetter: "HideFromAddressLists"},
			_jsii_.MemberProperty{JsiiProperty: "hideFromAddressListsInput", GoGetter: "HideFromAddressListsInput"},
			_jsii_.MemberProperty{JsiiProperty: "hideFromOutlookClients", GoGetter: "HideFromOutlookClients"},
			_jsii_.MemberProperty{JsiiProperty: "hideFromOutlookClientsInput", GoGetter: "HideFromOutlookClientsInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "mail", GoGetter: "Mail"},
			_jsii_.MemberProperty{JsiiProperty: "mailEnabled", GoGetter: "MailEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "mailEnabledInput", GoGetter: "MailEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "mailNickname", GoGetter: "MailNickname"},
			_jsii_.MemberProperty{JsiiProperty: "mailNicknameInput", GoGetter: "MailNicknameInput"},
			_jsii_.MemberProperty{JsiiProperty: "members", GoGetter: "Members"},
			_jsii_.MemberProperty{JsiiProperty: "membersInput", GoGetter: "MembersInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "objectId", GoGetter: "ObjectId"},
			_jsii_.MemberProperty{JsiiProperty: "onpremisesDomainName", GoGetter: "OnpremisesDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "onpremisesGroupType", GoGetter: "OnpremisesGroupType"},
			_jsii_.MemberProperty{JsiiProperty: "onpremisesGroupTypeInput", GoGetter: "OnpremisesGroupTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "onpremisesNetbiosName", GoGetter: "OnpremisesNetbiosName"},
			_jsii_.MemberProperty{JsiiProperty: "onpremisesSamAccountName", GoGetter: "OnpremisesSamAccountName"},
			_jsii_.MemberProperty{JsiiProperty: "onpremisesSecurityIdentifier", GoGetter: "OnpremisesSecurityIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "onpremisesSyncEnabled", GoGetter: "OnpremisesSyncEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "owners", GoGetter: "Owners"},
			_jsii_.MemberProperty{JsiiProperty: "ownersInput", GoGetter: "OwnersInput"},
			_jsii_.MemberProperty{JsiiProperty: "preferredLanguage", GoGetter: "PreferredLanguage"},
			_jsii_.MemberProperty{JsiiProperty: "preventDuplicateNames", GoGetter: "PreventDuplicateNames"},
			_jsii_.MemberProperty{JsiiProperty: "preventDuplicateNamesInput", GoGetter: "PreventDuplicateNamesInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "provisioningOptions", GoGetter: "ProvisioningOptions"},
			_jsii_.MemberProperty{JsiiProperty: "provisioningOptionsInput", GoGetter: "ProvisioningOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "proxyAddresses", GoGetter: "ProxyAddresses"},
			_jsii_.MemberMethod{JsiiMethod: "putDynamicMembership", GoMethod: "PutDynamicMembership"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAdministrativeUnitIds", GoMethod: "ResetAdministrativeUnitIds"},
			_jsii_.MemberMethod{JsiiMethod: "resetAssignableToRole", GoMethod: "ResetAssignableToRole"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoSubscribeNewMembers", GoMethod: "ResetAutoSubscribeNewMembers"},
			_jsii_.MemberMethod{JsiiMethod: "resetBehaviors", GoMethod: "ResetBehaviors"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetDynamicMembership", GoMethod: "ResetDynamicMembership"},
			_jsii_.MemberMethod{JsiiMethod: "resetExternalSendersAllowed", GoMethod: "ResetExternalSendersAllowed"},
			_jsii_.MemberMethod{JsiiMethod: "resetHideFromAddressLists", GoMethod: "ResetHideFromAddressLists"},
			_jsii_.MemberMethod{JsiiMethod: "resetHideFromOutlookClients", GoMethod: "ResetHideFromOutlookClients"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetMailEnabled", GoMethod: "ResetMailEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetMailNickname", GoMethod: "ResetMailNickname"},
			_jsii_.MemberMethod{JsiiMethod: "resetMembers", GoMethod: "ResetMembers"},
			_jsii_.MemberMethod{JsiiMethod: "resetOnpremisesGroupType", GoMethod: "ResetOnpremisesGroupType"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOwners", GoMethod: "ResetOwners"},
			_jsii_.MemberMethod{JsiiMethod: "resetPreventDuplicateNames", GoMethod: "ResetPreventDuplicateNames"},
			_jsii_.MemberMethod{JsiiMethod: "resetProvisioningOptions", GoMethod: "ResetProvisioningOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecurityEnabled", GoMethod: "ResetSecurityEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetTheme", GoMethod: "ResetTheme"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberMethod{JsiiMethod: "resetTypes", GoMethod: "ResetTypes"},
			_jsii_.MemberMethod{JsiiMethod: "resetVisibility", GoMethod: "ResetVisibility"},
			_jsii_.MemberMethod{JsiiMethod: "resetWritebackEnabled", GoMethod: "ResetWritebackEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "securityEnabled", GoGetter: "SecurityEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "securityEnabledInput", GoGetter: "SecurityEnabledInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "theme", GoGetter: "Theme"},
			_jsii_.MemberProperty{JsiiProperty: "themeInput", GoGetter: "ThemeInput"},
			_jsii_.MemberProperty{JsiiProperty: "timeouts", GoGetter: "Timeouts"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutsInput", GoGetter: "TimeoutsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "types", GoGetter: "Types"},
			_jsii_.MemberProperty{JsiiProperty: "typesInput", GoGetter: "TypesInput"},
			_jsii_.MemberProperty{JsiiProperty: "visibility", GoGetter: "Visibility"},
			_jsii_.MemberProperty{JsiiProperty: "visibilityInput", GoGetter: "VisibilityInput"},
			_jsii_.MemberProperty{JsiiProperty: "writebackEnabled", GoGetter: "WritebackEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "writebackEnabledInput", GoGetter: "WritebackEnabledInput"},
		},
		func() interface{} {
			j := jsiiProxy_Group{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azuread.group.GroupConfig",
		reflect.TypeOf((*GroupConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azuread.group.GroupDynamicMembership",
		reflect.TypeOf((*GroupDynamicMembership)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azuread.group.GroupDynamicMembershipOutputReference",
		reflect.TypeOf((*GroupDynamicMembershipOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "rule", GoGetter: "Rule"},
			_jsii_.MemberProperty{JsiiProperty: "ruleInput", GoGetter: "RuleInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GroupDynamicMembershipOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azuread.group.GroupTimeouts",
		reflect.TypeOf((*GroupTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azuread.group.GroupTimeoutsOutputReference",
		reflect.TypeOf((*GroupTimeoutsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "create", GoGetter: "Create"},
			_jsii_.MemberProperty{JsiiProperty: "createInput", GoGetter: "CreateInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "delete", GoGetter: "Delete"},
			_jsii_.MemberProperty{JsiiProperty: "deleteInput", GoGetter: "DeleteInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "read", GoGetter: "Read"},
			_jsii_.MemberProperty{JsiiProperty: "readInput", GoGetter: "ReadInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCreate", GoMethod: "ResetCreate"},
			_jsii_.MemberMethod{JsiiMethod: "resetDelete", GoMethod: "ResetDelete"},
			_jsii_.MemberMethod{JsiiMethod: "resetRead", GoMethod: "ResetRead"},
			_jsii_.MemberMethod{JsiiMethod: "resetUpdate", GoMethod: "ResetUpdate"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "update", GoGetter: "Update"},
			_jsii_.MemberProperty{JsiiProperty: "updateInput", GoGetter: "UpdateInput"},
		},
		func() interface{} {
			j := jsiiProxy_GroupTimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}

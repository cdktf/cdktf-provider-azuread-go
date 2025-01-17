// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package conditionalaccesspolicy

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-azuread.conditionalAccessPolicy.ConditionalAccessPolicy",
		reflect.TypeOf((*ConditionalAccessPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "conditions", GoGetter: "Conditions"},
			_jsii_.MemberProperty{JsiiProperty: "conditionsInput", GoGetter: "ConditionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "displayName", GoGetter: "DisplayName"},
			_jsii_.MemberProperty{JsiiProperty: "displayNameInput", GoGetter: "DisplayNameInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "grantControls", GoGetter: "GrantControls"},
			_jsii_.MemberProperty{JsiiProperty: "grantControlsInput", GoGetter: "GrantControlsInput"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "objectId", GoGetter: "ObjectId"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putConditions", GoMethod: "PutConditions"},
			_jsii_.MemberMethod{JsiiMethod: "putGrantControls", GoMethod: "PutGrantControls"},
			_jsii_.MemberMethod{JsiiMethod: "putSessionControls", GoMethod: "PutSessionControls"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetGrantControls", GoMethod: "ResetGrantControls"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSessionControls", GoMethod: "ResetSessionControls"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "sessionControls", GoGetter: "SessionControls"},
			_jsii_.MemberProperty{JsiiProperty: "sessionControlsInput", GoGetter: "SessionControlsInput"},
			_jsii_.MemberProperty{JsiiProperty: "state", GoGetter: "State"},
			_jsii_.MemberProperty{JsiiProperty: "stateInput", GoGetter: "StateInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "timeouts", GoGetter: "Timeouts"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutsInput", GoGetter: "TimeoutsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_ConditionalAccessPolicy{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azuread.conditionalAccessPolicy.ConditionalAccessPolicyConditions",
		reflect.TypeOf((*ConditionalAccessPolicyConditions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azuread.conditionalAccessPolicy.ConditionalAccessPolicyConditionsApplications",
		reflect.TypeOf((*ConditionalAccessPolicyConditionsApplications)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azuread.conditionalAccessPolicy.ConditionalAccessPolicyConditionsApplicationsOutputReference",
		reflect.TypeOf((*ConditionalAccessPolicyConditionsApplicationsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "excludedApplications", GoGetter: "ExcludedApplications"},
			_jsii_.MemberProperty{JsiiProperty: "excludedApplicationsInput", GoGetter: "ExcludedApplicationsInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "includedApplications", GoGetter: "IncludedApplications"},
			_jsii_.MemberProperty{JsiiProperty: "includedApplicationsInput", GoGetter: "IncludedApplicationsInput"},
			_jsii_.MemberProperty{JsiiProperty: "includedUserActions", GoGetter: "IncludedUserActions"},
			_jsii_.MemberProperty{JsiiProperty: "includedUserActionsInput", GoGetter: "IncludedUserActionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetExcludedApplications", GoMethod: "ResetExcludedApplications"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludedApplications", GoMethod: "ResetIncludedApplications"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludedUserActions", GoMethod: "ResetIncludedUserActions"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ConditionalAccessPolicyConditionsApplicationsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azuread.conditionalAccessPolicy.ConditionalAccessPolicyConditionsClientApplications",
		reflect.TypeOf((*ConditionalAccessPolicyConditionsClientApplications)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azuread.conditionalAccessPolicy.ConditionalAccessPolicyConditionsClientApplicationsOutputReference",
		reflect.TypeOf((*ConditionalAccessPolicyConditionsClientApplicationsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "excludedServicePrincipals", GoGetter: "ExcludedServicePrincipals"},
			_jsii_.MemberProperty{JsiiProperty: "excludedServicePrincipalsInput", GoGetter: "ExcludedServicePrincipalsInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "includedServicePrincipals", GoGetter: "IncludedServicePrincipals"},
			_jsii_.MemberProperty{JsiiProperty: "includedServicePrincipalsInput", GoGetter: "IncludedServicePrincipalsInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetExcludedServicePrincipals", GoMethod: "ResetExcludedServicePrincipals"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludedServicePrincipals", GoMethod: "ResetIncludedServicePrincipals"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ConditionalAccessPolicyConditionsClientApplicationsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azuread.conditionalAccessPolicy.ConditionalAccessPolicyConditionsDevices",
		reflect.TypeOf((*ConditionalAccessPolicyConditionsDevices)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azuread.conditionalAccessPolicy.ConditionalAccessPolicyConditionsDevicesFilter",
		reflect.TypeOf((*ConditionalAccessPolicyConditionsDevicesFilter)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azuread.conditionalAccessPolicy.ConditionalAccessPolicyConditionsDevicesFilterOutputReference",
		reflect.TypeOf((*ConditionalAccessPolicyConditionsDevicesFilterOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "mode", GoGetter: "Mode"},
			_jsii_.MemberProperty{JsiiProperty: "modeInput", GoGetter: "ModeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "rule", GoGetter: "Rule"},
			_jsii_.MemberProperty{JsiiProperty: "ruleInput", GoGetter: "RuleInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ConditionalAccessPolicyConditionsDevicesFilterOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azuread.conditionalAccessPolicy.ConditionalAccessPolicyConditionsDevicesOutputReference",
		reflect.TypeOf((*ConditionalAccessPolicyConditionsDevicesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "filter", GoGetter: "Filter"},
			_jsii_.MemberProperty{JsiiProperty: "filterInput", GoGetter: "FilterInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putFilter", GoMethod: "PutFilter"},
			_jsii_.MemberMethod{JsiiMethod: "resetFilter", GoMethod: "ResetFilter"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ConditionalAccessPolicyConditionsDevicesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azuread.conditionalAccessPolicy.ConditionalAccessPolicyConditionsLocations",
		reflect.TypeOf((*ConditionalAccessPolicyConditionsLocations)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azuread.conditionalAccessPolicy.ConditionalAccessPolicyConditionsLocationsOutputReference",
		reflect.TypeOf((*ConditionalAccessPolicyConditionsLocationsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "excludedLocations", GoGetter: "ExcludedLocations"},
			_jsii_.MemberProperty{JsiiProperty: "excludedLocationsInput", GoGetter: "ExcludedLocationsInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "includedLocations", GoGetter: "IncludedLocations"},
			_jsii_.MemberProperty{JsiiProperty: "includedLocationsInput", GoGetter: "IncludedLocationsInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetExcludedLocations", GoMethod: "ResetExcludedLocations"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ConditionalAccessPolicyConditionsLocationsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azuread.conditionalAccessPolicy.ConditionalAccessPolicyConditionsOutputReference",
		reflect.TypeOf((*ConditionalAccessPolicyConditionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "applications", GoGetter: "Applications"},
			_jsii_.MemberProperty{JsiiProperty: "applicationsInput", GoGetter: "ApplicationsInput"},
			_jsii_.MemberProperty{JsiiProperty: "clientApplications", GoGetter: "ClientApplications"},
			_jsii_.MemberProperty{JsiiProperty: "clientApplicationsInput", GoGetter: "ClientApplicationsInput"},
			_jsii_.MemberProperty{JsiiProperty: "clientAppTypes", GoGetter: "ClientAppTypes"},
			_jsii_.MemberProperty{JsiiProperty: "clientAppTypesInput", GoGetter: "ClientAppTypesInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "devices", GoGetter: "Devices"},
			_jsii_.MemberProperty{JsiiProperty: "devicesInput", GoGetter: "DevicesInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "insiderRiskLevels", GoGetter: "InsiderRiskLevels"},
			_jsii_.MemberProperty{JsiiProperty: "insiderRiskLevelsInput", GoGetter: "InsiderRiskLevelsInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "locations", GoGetter: "Locations"},
			_jsii_.MemberProperty{JsiiProperty: "locationsInput", GoGetter: "LocationsInput"},
			_jsii_.MemberProperty{JsiiProperty: "platforms", GoGetter: "Platforms"},
			_jsii_.MemberProperty{JsiiProperty: "platformsInput", GoGetter: "PlatformsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putApplications", GoMethod: "PutApplications"},
			_jsii_.MemberMethod{JsiiMethod: "putClientApplications", GoMethod: "PutClientApplications"},
			_jsii_.MemberMethod{JsiiMethod: "putDevices", GoMethod: "PutDevices"},
			_jsii_.MemberMethod{JsiiMethod: "putLocations", GoMethod: "PutLocations"},
			_jsii_.MemberMethod{JsiiMethod: "putPlatforms", GoMethod: "PutPlatforms"},
			_jsii_.MemberMethod{JsiiMethod: "putUsers", GoMethod: "PutUsers"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientApplications", GoMethod: "ResetClientApplications"},
			_jsii_.MemberMethod{JsiiMethod: "resetDevices", GoMethod: "ResetDevices"},
			_jsii_.MemberMethod{JsiiMethod: "resetInsiderRiskLevels", GoMethod: "ResetInsiderRiskLevels"},
			_jsii_.MemberMethod{JsiiMethod: "resetLocations", GoMethod: "ResetLocations"},
			_jsii_.MemberMethod{JsiiMethod: "resetPlatforms", GoMethod: "ResetPlatforms"},
			_jsii_.MemberMethod{JsiiMethod: "resetServicePrincipalRiskLevels", GoMethod: "ResetServicePrincipalRiskLevels"},
			_jsii_.MemberMethod{JsiiMethod: "resetSignInRiskLevels", GoMethod: "ResetSignInRiskLevels"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserRiskLevels", GoMethod: "ResetUserRiskLevels"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "servicePrincipalRiskLevels", GoGetter: "ServicePrincipalRiskLevels"},
			_jsii_.MemberProperty{JsiiProperty: "servicePrincipalRiskLevelsInput", GoGetter: "ServicePrincipalRiskLevelsInput"},
			_jsii_.MemberProperty{JsiiProperty: "signInRiskLevels", GoGetter: "SignInRiskLevels"},
			_jsii_.MemberProperty{JsiiProperty: "signInRiskLevelsInput", GoGetter: "SignInRiskLevelsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "userRiskLevels", GoGetter: "UserRiskLevels"},
			_jsii_.MemberProperty{JsiiProperty: "userRiskLevelsInput", GoGetter: "UserRiskLevelsInput"},
			_jsii_.MemberProperty{JsiiProperty: "users", GoGetter: "Users"},
			_jsii_.MemberProperty{JsiiProperty: "usersInput", GoGetter: "UsersInput"},
		},
		func() interface{} {
			j := jsiiProxy_ConditionalAccessPolicyConditionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azuread.conditionalAccessPolicy.ConditionalAccessPolicyConditionsPlatforms",
		reflect.TypeOf((*ConditionalAccessPolicyConditionsPlatforms)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azuread.conditionalAccessPolicy.ConditionalAccessPolicyConditionsPlatformsOutputReference",
		reflect.TypeOf((*ConditionalAccessPolicyConditionsPlatformsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "excludedPlatforms", GoGetter: "ExcludedPlatforms"},
			_jsii_.MemberProperty{JsiiProperty: "excludedPlatformsInput", GoGetter: "ExcludedPlatformsInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "includedPlatforms", GoGetter: "IncludedPlatforms"},
			_jsii_.MemberProperty{JsiiProperty: "includedPlatformsInput", GoGetter: "IncludedPlatformsInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetExcludedPlatforms", GoMethod: "ResetExcludedPlatforms"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ConditionalAccessPolicyConditionsPlatformsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azuread.conditionalAccessPolicy.ConditionalAccessPolicyConditionsUsers",
		reflect.TypeOf((*ConditionalAccessPolicyConditionsUsers)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azuread.conditionalAccessPolicy.ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUsers",
		reflect.TypeOf((*ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUsers)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azuread.conditionalAccessPolicy.ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUsersExternalTenants",
		reflect.TypeOf((*ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUsersExternalTenants)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azuread.conditionalAccessPolicy.ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUsersExternalTenantsList",
		reflect.TypeOf((*ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUsersExternalTenantsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUsersExternalTenantsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azuread.conditionalAccessPolicy.ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUsersExternalTenantsOutputReference",
		reflect.TypeOf((*ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUsersExternalTenantsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "members", GoGetter: "Members"},
			_jsii_.MemberProperty{JsiiProperty: "membershipKind", GoGetter: "MembershipKind"},
			_jsii_.MemberProperty{JsiiProperty: "membershipKindInput", GoGetter: "MembershipKindInput"},
			_jsii_.MemberProperty{JsiiProperty: "membersInput", GoGetter: "MembersInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetMembers", GoMethod: "ResetMembers"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUsersExternalTenantsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azuread.conditionalAccessPolicy.ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUsersList",
		reflect.TypeOf((*ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUsersList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUsersList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azuread.conditionalAccessPolicy.ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUsersOutputReference",
		reflect.TypeOf((*ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUsersOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "externalTenants", GoGetter: "ExternalTenants"},
			_jsii_.MemberProperty{JsiiProperty: "externalTenantsInput", GoGetter: "ExternalTenantsInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "guestOrExternalUserTypes", GoGetter: "GuestOrExternalUserTypes"},
			_jsii_.MemberProperty{JsiiProperty: "guestOrExternalUserTypesInput", GoGetter: "GuestOrExternalUserTypesInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putExternalTenants", GoMethod: "PutExternalTenants"},
			_jsii_.MemberMethod{JsiiMethod: "resetExternalTenants", GoMethod: "ResetExternalTenants"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUsersOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azuread.conditionalAccessPolicy.ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsers",
		reflect.TypeOf((*ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsers)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azuread.conditionalAccessPolicy.ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenants",
		reflect.TypeOf((*ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenants)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azuread.conditionalAccessPolicy.ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsList",
		reflect.TypeOf((*ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azuread.conditionalAccessPolicy.ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsOutputReference",
		reflect.TypeOf((*ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "members", GoGetter: "Members"},
			_jsii_.MemberProperty{JsiiProperty: "membershipKind", GoGetter: "MembershipKind"},
			_jsii_.MemberProperty{JsiiProperty: "membershipKindInput", GoGetter: "MembershipKindInput"},
			_jsii_.MemberProperty{JsiiProperty: "membersInput", GoGetter: "MembersInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetMembers", GoMethod: "ResetMembers"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azuread.conditionalAccessPolicy.ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersList",
		reflect.TypeOf((*ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azuread.conditionalAccessPolicy.ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersOutputReference",
		reflect.TypeOf((*ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "externalTenants", GoGetter: "ExternalTenants"},
			_jsii_.MemberProperty{JsiiProperty: "externalTenantsInput", GoGetter: "ExternalTenantsInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "guestOrExternalUserTypes", GoGetter: "GuestOrExternalUserTypes"},
			_jsii_.MemberProperty{JsiiProperty: "guestOrExternalUserTypesInput", GoGetter: "GuestOrExternalUserTypesInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putExternalTenants", GoMethod: "PutExternalTenants"},
			_jsii_.MemberMethod{JsiiMethod: "resetExternalTenants", GoMethod: "ResetExternalTenants"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azuread.conditionalAccessPolicy.ConditionalAccessPolicyConditionsUsersOutputReference",
		reflect.TypeOf((*ConditionalAccessPolicyConditionsUsersOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "excludedGroups", GoGetter: "ExcludedGroups"},
			_jsii_.MemberProperty{JsiiProperty: "excludedGroupsInput", GoGetter: "ExcludedGroupsInput"},
			_jsii_.MemberProperty{JsiiProperty: "excludedGuestsOrExternalUsers", GoGetter: "ExcludedGuestsOrExternalUsers"},
			_jsii_.MemberProperty{JsiiProperty: "excludedGuestsOrExternalUsersInput", GoGetter: "ExcludedGuestsOrExternalUsersInput"},
			_jsii_.MemberProperty{JsiiProperty: "excludedRoles", GoGetter: "ExcludedRoles"},
			_jsii_.MemberProperty{JsiiProperty: "excludedRolesInput", GoGetter: "ExcludedRolesInput"},
			_jsii_.MemberProperty{JsiiProperty: "excludedUsers", GoGetter: "ExcludedUsers"},
			_jsii_.MemberProperty{JsiiProperty: "excludedUsersInput", GoGetter: "ExcludedUsersInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "includedGroups", GoGetter: "IncludedGroups"},
			_jsii_.MemberProperty{JsiiProperty: "includedGroupsInput", GoGetter: "IncludedGroupsInput"},
			_jsii_.MemberProperty{JsiiProperty: "includedGuestsOrExternalUsers", GoGetter: "IncludedGuestsOrExternalUsers"},
			_jsii_.MemberProperty{JsiiProperty: "includedGuestsOrExternalUsersInput", GoGetter: "IncludedGuestsOrExternalUsersInput"},
			_jsii_.MemberProperty{JsiiProperty: "includedRoles", GoGetter: "IncludedRoles"},
			_jsii_.MemberProperty{JsiiProperty: "includedRolesInput", GoGetter: "IncludedRolesInput"},
			_jsii_.MemberProperty{JsiiProperty: "includedUsers", GoGetter: "IncludedUsers"},
			_jsii_.MemberProperty{JsiiProperty: "includedUsersInput", GoGetter: "IncludedUsersInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putExcludedGuestsOrExternalUsers", GoMethod: "PutExcludedGuestsOrExternalUsers"},
			_jsii_.MemberMethod{JsiiMethod: "putIncludedGuestsOrExternalUsers", GoMethod: "PutIncludedGuestsOrExternalUsers"},
			_jsii_.MemberMethod{JsiiMethod: "resetExcludedGroups", GoMethod: "ResetExcludedGroups"},
			_jsii_.MemberMethod{JsiiMethod: "resetExcludedGuestsOrExternalUsers", GoMethod: "ResetExcludedGuestsOrExternalUsers"},
			_jsii_.MemberMethod{JsiiMethod: "resetExcludedRoles", GoMethod: "ResetExcludedRoles"},
			_jsii_.MemberMethod{JsiiMethod: "resetExcludedUsers", GoMethod: "ResetExcludedUsers"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludedGroups", GoMethod: "ResetIncludedGroups"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludedGuestsOrExternalUsers", GoMethod: "ResetIncludedGuestsOrExternalUsers"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludedRoles", GoMethod: "ResetIncludedRoles"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludedUsers", GoMethod: "ResetIncludedUsers"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azuread.conditionalAccessPolicy.ConditionalAccessPolicyConfig",
		reflect.TypeOf((*ConditionalAccessPolicyConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azuread.conditionalAccessPolicy.ConditionalAccessPolicyGrantControls",
		reflect.TypeOf((*ConditionalAccessPolicyGrantControls)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azuread.conditionalAccessPolicy.ConditionalAccessPolicyGrantControlsOutputReference",
		reflect.TypeOf((*ConditionalAccessPolicyGrantControlsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "authenticationStrengthPolicyId", GoGetter: "AuthenticationStrengthPolicyId"},
			_jsii_.MemberProperty{JsiiProperty: "authenticationStrengthPolicyIdInput", GoGetter: "AuthenticationStrengthPolicyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "builtInControls", GoGetter: "BuiltInControls"},
			_jsii_.MemberProperty{JsiiProperty: "builtInControlsInput", GoGetter: "BuiltInControlsInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "customAuthenticationFactors", GoGetter: "CustomAuthenticationFactors"},
			_jsii_.MemberProperty{JsiiProperty: "customAuthenticationFactorsInput", GoGetter: "CustomAuthenticationFactorsInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "operator", GoGetter: "Operator"},
			_jsii_.MemberProperty{JsiiProperty: "operatorInput", GoGetter: "OperatorInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthenticationStrengthPolicyId", GoMethod: "ResetAuthenticationStrengthPolicyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetBuiltInControls", GoMethod: "ResetBuiltInControls"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomAuthenticationFactors", GoMethod: "ResetCustomAuthenticationFactors"},
			_jsii_.MemberMethod{JsiiMethod: "resetTermsOfUse", GoMethod: "ResetTermsOfUse"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "termsOfUse", GoGetter: "TermsOfUse"},
			_jsii_.MemberProperty{JsiiProperty: "termsOfUseInput", GoGetter: "TermsOfUseInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ConditionalAccessPolicyGrantControlsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azuread.conditionalAccessPolicy.ConditionalAccessPolicySessionControls",
		reflect.TypeOf((*ConditionalAccessPolicySessionControls)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azuread.conditionalAccessPolicy.ConditionalAccessPolicySessionControlsOutputReference",
		reflect.TypeOf((*ConditionalAccessPolicySessionControlsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "applicationEnforcedRestrictionsEnabled", GoGetter: "ApplicationEnforcedRestrictionsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "applicationEnforcedRestrictionsEnabledInput", GoGetter: "ApplicationEnforcedRestrictionsEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "cloudAppSecurityPolicy", GoGetter: "CloudAppSecurityPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cloudAppSecurityPolicyInput", GoGetter: "CloudAppSecurityPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "disableResilienceDefaults", GoGetter: "DisableResilienceDefaults"},
			_jsii_.MemberProperty{JsiiProperty: "disableResilienceDefaultsInput", GoGetter: "DisableResilienceDefaultsInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "persistentBrowserMode", GoGetter: "PersistentBrowserMode"},
			_jsii_.MemberProperty{JsiiProperty: "persistentBrowserModeInput", GoGetter: "PersistentBrowserModeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetApplicationEnforcedRestrictionsEnabled", GoMethod: "ResetApplicationEnforcedRestrictionsEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloudAppSecurityPolicy", GoMethod: "ResetCloudAppSecurityPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisableResilienceDefaults", GoMethod: "ResetDisableResilienceDefaults"},
			_jsii_.MemberMethod{JsiiMethod: "resetPersistentBrowserMode", GoMethod: "ResetPersistentBrowserMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetSignInFrequency", GoMethod: "ResetSignInFrequency"},
			_jsii_.MemberMethod{JsiiMethod: "resetSignInFrequencyAuthenticationType", GoMethod: "ResetSignInFrequencyAuthenticationType"},
			_jsii_.MemberMethod{JsiiMethod: "resetSignInFrequencyInterval", GoMethod: "ResetSignInFrequencyInterval"},
			_jsii_.MemberMethod{JsiiMethod: "resetSignInFrequencyPeriod", GoMethod: "ResetSignInFrequencyPeriod"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "signInFrequency", GoGetter: "SignInFrequency"},
			_jsii_.MemberProperty{JsiiProperty: "signInFrequencyAuthenticationType", GoGetter: "SignInFrequencyAuthenticationType"},
			_jsii_.MemberProperty{JsiiProperty: "signInFrequencyAuthenticationTypeInput", GoGetter: "SignInFrequencyAuthenticationTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "signInFrequencyInput", GoGetter: "SignInFrequencyInput"},
			_jsii_.MemberProperty{JsiiProperty: "signInFrequencyInterval", GoGetter: "SignInFrequencyInterval"},
			_jsii_.MemberProperty{JsiiProperty: "signInFrequencyIntervalInput", GoGetter: "SignInFrequencyIntervalInput"},
			_jsii_.MemberProperty{JsiiProperty: "signInFrequencyPeriod", GoGetter: "SignInFrequencyPeriod"},
			_jsii_.MemberProperty{JsiiProperty: "signInFrequencyPeriodInput", GoGetter: "SignInFrequencyPeriodInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azuread.conditionalAccessPolicy.ConditionalAccessPolicyTimeouts",
		reflect.TypeOf((*ConditionalAccessPolicyTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azuread.conditionalAccessPolicy.ConditionalAccessPolicyTimeoutsOutputReference",
		reflect.TypeOf((*ConditionalAccessPolicyTimeoutsOutputReference)(nil)).Elem(),
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
			j := jsiiProxy_ConditionalAccessPolicyTimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}

// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataazureaduser

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-azuread.dataAzureadUser.DataAzureadUser",
		reflect.TypeOf((*DataAzureadUser)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accountEnabled", GoGetter: "AccountEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "ageGroup", GoGetter: "AgeGroup"},
			_jsii_.MemberProperty{JsiiProperty: "businessPhones", GoGetter: "BusinessPhones"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "city", GoGetter: "City"},
			_jsii_.MemberProperty{JsiiProperty: "companyName", GoGetter: "CompanyName"},
			_jsii_.MemberProperty{JsiiProperty: "consentProvidedForMinor", GoGetter: "ConsentProvidedForMinor"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "costCenter", GoGetter: "CostCenter"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "country", GoGetter: "Country"},
			_jsii_.MemberProperty{JsiiProperty: "creationType", GoGetter: "CreationType"},
			_jsii_.MemberProperty{JsiiProperty: "department", GoGetter: "Department"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "displayName", GoGetter: "DisplayName"},
			_jsii_.MemberProperty{JsiiProperty: "division", GoGetter: "Division"},
			_jsii_.MemberProperty{JsiiProperty: "employeeId", GoGetter: "EmployeeId"},
			_jsii_.MemberProperty{JsiiProperty: "employeeIdInput", GoGetter: "EmployeeIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "employeeType", GoGetter: "EmployeeType"},
			_jsii_.MemberProperty{JsiiProperty: "externalUserState", GoGetter: "ExternalUserState"},
			_jsii_.MemberProperty{JsiiProperty: "faxNumber", GoGetter: "FaxNumber"},
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
			_jsii_.MemberProperty{JsiiProperty: "givenName", GoGetter: "GivenName"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "imAddresses", GoGetter: "ImAddresses"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "jobTitle", GoGetter: "JobTitle"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "mail", GoGetter: "Mail"},
			_jsii_.MemberProperty{JsiiProperty: "mailInput", GoGetter: "MailInput"},
			_jsii_.MemberProperty{JsiiProperty: "mailNickname", GoGetter: "MailNickname"},
			_jsii_.MemberProperty{JsiiProperty: "mailNicknameInput", GoGetter: "MailNicknameInput"},
			_jsii_.MemberProperty{JsiiProperty: "managerId", GoGetter: "ManagerId"},
			_jsii_.MemberProperty{JsiiProperty: "mobilePhone", GoGetter: "MobilePhone"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "objectId", GoGetter: "ObjectId"},
			_jsii_.MemberProperty{JsiiProperty: "objectIdInput", GoGetter: "ObjectIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "officeLocation", GoGetter: "OfficeLocation"},
			_jsii_.MemberProperty{JsiiProperty: "onpremisesDistinguishedName", GoGetter: "OnpremisesDistinguishedName"},
			_jsii_.MemberProperty{JsiiProperty: "onpremisesDomainName", GoGetter: "OnpremisesDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "onpremisesImmutableId", GoGetter: "OnpremisesImmutableId"},
			_jsii_.MemberProperty{JsiiProperty: "onpremisesSamAccountName", GoGetter: "OnpremisesSamAccountName"},
			_jsii_.MemberProperty{JsiiProperty: "onpremisesSecurityIdentifier", GoGetter: "OnpremisesSecurityIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "onpremisesSyncEnabled", GoGetter: "OnpremisesSyncEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "onpremisesUserPrincipalName", GoGetter: "OnpremisesUserPrincipalName"},
			_jsii_.MemberProperty{JsiiProperty: "otherMails", GoGetter: "OtherMails"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "postalCode", GoGetter: "PostalCode"},
			_jsii_.MemberProperty{JsiiProperty: "preferredLanguage", GoGetter: "PreferredLanguage"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "proxyAddresses", GoGetter: "ProxyAddresses"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetEmployeeId", GoMethod: "ResetEmployeeId"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetMail", GoMethod: "ResetMail"},
			_jsii_.MemberMethod{JsiiMethod: "resetMailNickname", GoMethod: "ResetMailNickname"},
			_jsii_.MemberMethod{JsiiMethod: "resetObjectId", GoMethod: "ResetObjectId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserPrincipalName", GoMethod: "ResetUserPrincipalName"},
			_jsii_.MemberProperty{JsiiProperty: "showInAddressList", GoGetter: "ShowInAddressList"},
			_jsii_.MemberProperty{JsiiProperty: "state", GoGetter: "State"},
			_jsii_.MemberProperty{JsiiProperty: "streetAddress", GoGetter: "StreetAddress"},
			_jsii_.MemberProperty{JsiiProperty: "surname", GoGetter: "Surname"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "timeouts", GoGetter: "Timeouts"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutsInput", GoGetter: "TimeoutsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "usageLocation", GoGetter: "UsageLocation"},
			_jsii_.MemberProperty{JsiiProperty: "userPrincipalName", GoGetter: "UserPrincipalName"},
			_jsii_.MemberProperty{JsiiProperty: "userPrincipalNameInput", GoGetter: "UserPrincipalNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "userType", GoGetter: "UserType"},
		},
		func() interface{} {
			j := jsiiProxy_DataAzureadUser{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azuread.dataAzureadUser.DataAzureadUserConfig",
		reflect.TypeOf((*DataAzureadUserConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azuread.dataAzureadUser.DataAzureadUserTimeouts",
		reflect.TypeOf((*DataAzureadUserTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azuread.dataAzureadUser.DataAzureadUserTimeoutsOutputReference",
		reflect.TypeOf((*DataAzureadUserTimeoutsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "read", GoGetter: "Read"},
			_jsii_.MemberProperty{JsiiProperty: "readInput", GoGetter: "ReadInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetRead", GoMethod: "ResetRead"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataAzureadUserTimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}

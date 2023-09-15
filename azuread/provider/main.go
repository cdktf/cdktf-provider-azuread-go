// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-azuread.provider.AzureadProvider",
		reflect.TypeOf((*AzureadProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "alias", GoGetter: "Alias"},
			_jsii_.MemberProperty{JsiiProperty: "aliasInput", GoGetter: "AliasInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "clientCertificate", GoGetter: "ClientCertificate"},
			_jsii_.MemberProperty{JsiiProperty: "clientCertificateInput", GoGetter: "ClientCertificateInput"},
			_jsii_.MemberProperty{JsiiProperty: "clientCertificatePassword", GoGetter: "ClientCertificatePassword"},
			_jsii_.MemberProperty{JsiiProperty: "clientCertificatePasswordInput", GoGetter: "ClientCertificatePasswordInput"},
			_jsii_.MemberProperty{JsiiProperty: "clientCertificatePath", GoGetter: "ClientCertificatePath"},
			_jsii_.MemberProperty{JsiiProperty: "clientCertificatePathInput", GoGetter: "ClientCertificatePathInput"},
			_jsii_.MemberProperty{JsiiProperty: "clientId", GoGetter: "ClientId"},
			_jsii_.MemberProperty{JsiiProperty: "clientIdFilePath", GoGetter: "ClientIdFilePath"},
			_jsii_.MemberProperty{JsiiProperty: "clientIdFilePathInput", GoGetter: "ClientIdFilePathInput"},
			_jsii_.MemberProperty{JsiiProperty: "clientIdInput", GoGetter: "ClientIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "clientSecret", GoGetter: "ClientSecret"},
			_jsii_.MemberProperty{JsiiProperty: "clientSecretFilePath", GoGetter: "ClientSecretFilePath"},
			_jsii_.MemberProperty{JsiiProperty: "clientSecretFilePathInput", GoGetter: "ClientSecretFilePathInput"},
			_jsii_.MemberProperty{JsiiProperty: "clientSecretInput", GoGetter: "ClientSecretInput"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "disableTerraformPartnerId", GoGetter: "DisableTerraformPartnerId"},
			_jsii_.MemberProperty{JsiiProperty: "disableTerraformPartnerIdInput", GoGetter: "DisableTerraformPartnerIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "environment", GoGetter: "Environment"},
			_jsii_.MemberProperty{JsiiProperty: "environmentInput", GoGetter: "EnvironmentInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "metaAttributes", GoGetter: "MetaAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "metadataHost", GoGetter: "MetadataHost"},
			_jsii_.MemberProperty{JsiiProperty: "metadataHostInput", GoGetter: "MetadataHostInput"},
			_jsii_.MemberProperty{JsiiProperty: "msiEndpoint", GoGetter: "MsiEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "msiEndpointInput", GoGetter: "MsiEndpointInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "oidcRequestToken", GoGetter: "OidcRequestToken"},
			_jsii_.MemberProperty{JsiiProperty: "oidcRequestTokenInput", GoGetter: "OidcRequestTokenInput"},
			_jsii_.MemberProperty{JsiiProperty: "oidcRequestUrl", GoGetter: "OidcRequestUrl"},
			_jsii_.MemberProperty{JsiiProperty: "oidcRequestUrlInput", GoGetter: "OidcRequestUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "oidcToken", GoGetter: "OidcToken"},
			_jsii_.MemberProperty{JsiiProperty: "oidcTokenFilePath", GoGetter: "OidcTokenFilePath"},
			_jsii_.MemberProperty{JsiiProperty: "oidcTokenFilePathInput", GoGetter: "OidcTokenFilePathInput"},
			_jsii_.MemberProperty{JsiiProperty: "oidcTokenInput", GoGetter: "OidcTokenInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "partnerId", GoGetter: "PartnerId"},
			_jsii_.MemberProperty{JsiiProperty: "partnerIdInput", GoGetter: "PartnerIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAlias", GoMethod: "ResetAlias"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientCertificate", GoMethod: "ResetClientCertificate"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientCertificatePassword", GoMethod: "ResetClientCertificatePassword"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientCertificatePath", GoMethod: "ResetClientCertificatePath"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientId", GoMethod: "ResetClientId"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientIdFilePath", GoMethod: "ResetClientIdFilePath"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientSecret", GoMethod: "ResetClientSecret"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientSecretFilePath", GoMethod: "ResetClientSecretFilePath"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisableTerraformPartnerId", GoMethod: "ResetDisableTerraformPartnerId"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnvironment", GoMethod: "ResetEnvironment"},
			_jsii_.MemberMethod{JsiiMethod: "resetMetadataHost", GoMethod: "ResetMetadataHost"},
			_jsii_.MemberMethod{JsiiMethod: "resetMsiEndpoint", GoMethod: "ResetMsiEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "resetOidcRequestToken", GoMethod: "ResetOidcRequestToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetOidcRequestUrl", GoMethod: "ResetOidcRequestUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetOidcToken", GoMethod: "ResetOidcToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetOidcTokenFilePath", GoMethod: "ResetOidcTokenFilePath"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPartnerId", GoMethod: "ResetPartnerId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTenantId", GoMethod: "ResetTenantId"},
			_jsii_.MemberMethod{JsiiMethod: "resetUseCli", GoMethod: "ResetUseCli"},
			_jsii_.MemberMethod{JsiiMethod: "resetUseMsi", GoMethod: "ResetUseMsi"},
			_jsii_.MemberMethod{JsiiMethod: "resetUseOidc", GoMethod: "ResetUseOidc"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tenantId", GoGetter: "TenantId"},
			_jsii_.MemberProperty{JsiiProperty: "tenantIdInput", GoGetter: "TenantIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformProviderSource", GoGetter: "TerraformProviderSource"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "useCli", GoGetter: "UseCli"},
			_jsii_.MemberProperty{JsiiProperty: "useCliInput", GoGetter: "UseCliInput"},
			_jsii_.MemberProperty{JsiiProperty: "useMsi", GoGetter: "UseMsi"},
			_jsii_.MemberProperty{JsiiProperty: "useMsiInput", GoGetter: "UseMsiInput"},
			_jsii_.MemberProperty{JsiiProperty: "useOidc", GoGetter: "UseOidc"},
			_jsii_.MemberProperty{JsiiProperty: "useOidcInput", GoGetter: "UseOidcInput"},
		},
		func() interface{} {
			j := jsiiProxy_AzureadProvider{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azuread.provider.AzureadProviderConfig",
		reflect.TypeOf((*AzureadProviderConfig)(nil)).Elem(),
	)
}

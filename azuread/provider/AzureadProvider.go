// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azuread-go/azuread/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azuread-go/azuread/v13/provider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs azuread}.
type AzureadProvider interface {
	cdktf.TerraformProvider
	AdoPipelineServiceConnectionId() *string
	SetAdoPipelineServiceConnectionId(val *string)
	AdoPipelineServiceConnectionIdInput() *string
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientCertificate() *string
	SetClientCertificate(val *string)
	ClientCertificateInput() *string
	ClientCertificatePassword() *string
	SetClientCertificatePassword(val *string)
	ClientCertificatePasswordInput() *string
	ClientCertificatePath() *string
	SetClientCertificatePath(val *string)
	ClientCertificatePathInput() *string
	ClientId() *string
	SetClientId(val *string)
	ClientIdFilePath() *string
	SetClientIdFilePath(val *string)
	ClientIdFilePathInput() *string
	ClientIdInput() *string
	ClientSecret() *string
	SetClientSecret(val *string)
	ClientSecretFilePath() *string
	SetClientSecretFilePath(val *string)
	ClientSecretFilePathInput() *string
	ClientSecretInput() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	DisableTerraformPartnerId() interface{}
	SetDisableTerraformPartnerId(val interface{})
	DisableTerraformPartnerIdInput() interface{}
	Environment() *string
	SetEnvironment(val *string)
	EnvironmentInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	MetaAttributes() *map[string]interface{}
	MetadataHost() *string
	SetMetadataHost(val *string)
	MetadataHostInput() *string
	MsiEndpoint() *string
	SetMsiEndpoint(val *string)
	MsiEndpointInput() *string
	// The tree node.
	Node() constructs.Node
	OidcRequestToken() *string
	SetOidcRequestToken(val *string)
	OidcRequestTokenInput() *string
	OidcRequestUrl() *string
	SetOidcRequestUrl(val *string)
	OidcRequestUrlInput() *string
	OidcToken() *string
	SetOidcToken(val *string)
	OidcTokenFilePath() *string
	SetOidcTokenFilePath(val *string)
	OidcTokenFilePathInput() *string
	OidcTokenInput() *string
	PartnerId() *string
	SetPartnerId(val *string)
	PartnerIdInput() *string
	// Experimental.
	RawOverrides() interface{}
	TenantId() *string
	SetTenantId(val *string)
	TenantIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	UseAksWorkloadIdentity() interface{}
	SetUseAksWorkloadIdentity(val interface{})
	UseAksWorkloadIdentityInput() interface{}
	UseCli() interface{}
	SetUseCli(val interface{})
	UseCliInput() interface{}
	UseMsi() interface{}
	SetUseMsi(val interface{})
	UseMsiInput() interface{}
	UseOidc() interface{}
	SetUseOidc(val interface{})
	UseOidcInput() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAdoPipelineServiceConnectionId()
	ResetAlias()
	ResetClientCertificate()
	ResetClientCertificatePassword()
	ResetClientCertificatePath()
	ResetClientId()
	ResetClientIdFilePath()
	ResetClientSecret()
	ResetClientSecretFilePath()
	ResetDisableTerraformPartnerId()
	ResetEnvironment()
	ResetMetadataHost()
	ResetMsiEndpoint()
	ResetOidcRequestToken()
	ResetOidcRequestUrl()
	ResetOidcToken()
	ResetOidcTokenFilePath()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPartnerId()
	ResetTenantId()
	ResetUseAksWorkloadIdentity()
	ResetUseCli()
	ResetUseMsi()
	ResetUseOidc()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for AzureadProvider
type jsiiProxy_AzureadProvider struct {
	internal.Type__cdktfTerraformProvider
}

func (j *jsiiProxy_AzureadProvider) AdoPipelineServiceConnectionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adoPipelineServiceConnectionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) AdoPipelineServiceConnectionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adoPipelineServiceConnectionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) ClientCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) ClientCertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) ClientCertificatePassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificatePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) ClientCertificatePasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificatePasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) ClientCertificatePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificatePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) ClientCertificatePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificatePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) ClientIdFilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdFilePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) ClientIdFilePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdFilePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) ClientSecretFilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretFilePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) ClientSecretFilePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretFilePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) ClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) DisableTerraformPartnerId() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableTerraformPartnerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) DisableTerraformPartnerIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableTerraformPartnerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) Environment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) EnvironmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) MetadataHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) MetadataHostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataHostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) MsiEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"msiEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) MsiEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"msiEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) OidcRequestToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oidcRequestToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) OidcRequestTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oidcRequestTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) OidcRequestUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oidcRequestUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) OidcRequestUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oidcRequestUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) OidcToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oidcToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) OidcTokenFilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oidcTokenFilePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) OidcTokenFilePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oidcTokenFilePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) OidcTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oidcTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) PartnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partnerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) PartnerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partnerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) TenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) UseAksWorkloadIdentity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useAksWorkloadIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) UseAksWorkloadIdentityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useAksWorkloadIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) UseCli() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCli",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) UseCliInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCliInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) UseMsi() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useMsi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) UseMsiInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useMsiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) UseOidc() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useOidc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureadProvider) UseOidcInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useOidcInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs azuread} Resource.
func NewAzureadProvider(scope constructs.Construct, id *string, config *AzureadProviderConfig) AzureadProvider {
	_init_.Initialize()

	if err := validateNewAzureadProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AzureadProvider{}

	_jsii_.Create(
		"@cdktf/provider-azuread.provider.AzureadProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs azuread} Resource.
func NewAzureadProvider_Override(a AzureadProvider, scope constructs.Construct, id *string, config *AzureadProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azuread.provider.AzureadProvider",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AzureadProvider)SetAdoPipelineServiceConnectionId(val *string) {
	_jsii_.Set(
		j,
		"adoPipelineServiceConnectionId",
		val,
	)
}

func (j *jsiiProxy_AzureadProvider)SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_AzureadProvider)SetClientCertificate(val *string) {
	_jsii_.Set(
		j,
		"clientCertificate",
		val,
	)
}

func (j *jsiiProxy_AzureadProvider)SetClientCertificatePassword(val *string) {
	_jsii_.Set(
		j,
		"clientCertificatePassword",
		val,
	)
}

func (j *jsiiProxy_AzureadProvider)SetClientCertificatePath(val *string) {
	_jsii_.Set(
		j,
		"clientCertificatePath",
		val,
	)
}

func (j *jsiiProxy_AzureadProvider)SetClientId(val *string) {
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_AzureadProvider)SetClientIdFilePath(val *string) {
	_jsii_.Set(
		j,
		"clientIdFilePath",
		val,
	)
}

func (j *jsiiProxy_AzureadProvider)SetClientSecret(val *string) {
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_AzureadProvider)SetClientSecretFilePath(val *string) {
	_jsii_.Set(
		j,
		"clientSecretFilePath",
		val,
	)
}

func (j *jsiiProxy_AzureadProvider)SetDisableTerraformPartnerId(val interface{}) {
	if err := j.validateSetDisableTerraformPartnerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableTerraformPartnerId",
		val,
	)
}

func (j *jsiiProxy_AzureadProvider)SetEnvironment(val *string) {
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_AzureadProvider)SetMetadataHost(val *string) {
	_jsii_.Set(
		j,
		"metadataHost",
		val,
	)
}

func (j *jsiiProxy_AzureadProvider)SetMsiEndpoint(val *string) {
	_jsii_.Set(
		j,
		"msiEndpoint",
		val,
	)
}

func (j *jsiiProxy_AzureadProvider)SetOidcRequestToken(val *string) {
	_jsii_.Set(
		j,
		"oidcRequestToken",
		val,
	)
}

func (j *jsiiProxy_AzureadProvider)SetOidcRequestUrl(val *string) {
	_jsii_.Set(
		j,
		"oidcRequestUrl",
		val,
	)
}

func (j *jsiiProxy_AzureadProvider)SetOidcToken(val *string) {
	_jsii_.Set(
		j,
		"oidcToken",
		val,
	)
}

func (j *jsiiProxy_AzureadProvider)SetOidcTokenFilePath(val *string) {
	_jsii_.Set(
		j,
		"oidcTokenFilePath",
		val,
	)
}

func (j *jsiiProxy_AzureadProvider)SetPartnerId(val *string) {
	_jsii_.Set(
		j,
		"partnerId",
		val,
	)
}

func (j *jsiiProxy_AzureadProvider)SetTenantId(val *string) {
	_jsii_.Set(
		j,
		"tenantId",
		val,
	)
}

func (j *jsiiProxy_AzureadProvider)SetUseAksWorkloadIdentity(val interface{}) {
	if err := j.validateSetUseAksWorkloadIdentityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useAksWorkloadIdentity",
		val,
	)
}

func (j *jsiiProxy_AzureadProvider)SetUseCli(val interface{}) {
	if err := j.validateSetUseCliParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useCli",
		val,
	)
}

func (j *jsiiProxy_AzureadProvider)SetUseMsi(val interface{}) {
	if err := j.validateSetUseMsiParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useMsi",
		val,
	)
}

func (j *jsiiProxy_AzureadProvider)SetUseOidc(val interface{}) {
	if err := j.validateSetUseOidcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useOidc",
		val,
	)
}

// Generates CDKTF code for importing a AzureadProvider resource upon running "cdktf plan <stack-name>".
func AzureadProvider_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateAzureadProvider_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azuread.provider.AzureadProvider",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func AzureadProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAzureadProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azuread.provider.AzureadProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AzureadProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAzureadProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azuread.provider.AzureadProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AzureadProvider_IsTerraformProvider(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAzureadProvider_IsTerraformProviderParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azuread.provider.AzureadProvider",
		"isTerraformProvider",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AzureadProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azuread.provider.AzureadProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AzureadProvider) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AzureadProvider) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AzureadProvider) ResetAdoPipelineServiceConnectionId() {
	_jsii_.InvokeVoid(
		a,
		"resetAdoPipelineServiceConnectionId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureadProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		a,
		"resetAlias",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureadProvider) ResetClientCertificate() {
	_jsii_.InvokeVoid(
		a,
		"resetClientCertificate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureadProvider) ResetClientCertificatePassword() {
	_jsii_.InvokeVoid(
		a,
		"resetClientCertificatePassword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureadProvider) ResetClientCertificatePath() {
	_jsii_.InvokeVoid(
		a,
		"resetClientCertificatePath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureadProvider) ResetClientId() {
	_jsii_.InvokeVoid(
		a,
		"resetClientId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureadProvider) ResetClientIdFilePath() {
	_jsii_.InvokeVoid(
		a,
		"resetClientIdFilePath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureadProvider) ResetClientSecret() {
	_jsii_.InvokeVoid(
		a,
		"resetClientSecret",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureadProvider) ResetClientSecretFilePath() {
	_jsii_.InvokeVoid(
		a,
		"resetClientSecretFilePath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureadProvider) ResetDisableTerraformPartnerId() {
	_jsii_.InvokeVoid(
		a,
		"resetDisableTerraformPartnerId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureadProvider) ResetEnvironment() {
	_jsii_.InvokeVoid(
		a,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureadProvider) ResetMetadataHost() {
	_jsii_.InvokeVoid(
		a,
		"resetMetadataHost",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureadProvider) ResetMsiEndpoint() {
	_jsii_.InvokeVoid(
		a,
		"resetMsiEndpoint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureadProvider) ResetOidcRequestToken() {
	_jsii_.InvokeVoid(
		a,
		"resetOidcRequestToken",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureadProvider) ResetOidcRequestUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetOidcRequestUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureadProvider) ResetOidcToken() {
	_jsii_.InvokeVoid(
		a,
		"resetOidcToken",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureadProvider) ResetOidcTokenFilePath() {
	_jsii_.InvokeVoid(
		a,
		"resetOidcTokenFilePath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureadProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureadProvider) ResetPartnerId() {
	_jsii_.InvokeVoid(
		a,
		"resetPartnerId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureadProvider) ResetTenantId() {
	_jsii_.InvokeVoid(
		a,
		"resetTenantId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureadProvider) ResetUseAksWorkloadIdentity() {
	_jsii_.InvokeVoid(
		a,
		"resetUseAksWorkloadIdentity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureadProvider) ResetUseCli() {
	_jsii_.InvokeVoid(
		a,
		"resetUseCli",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureadProvider) ResetUseMsi() {
	_jsii_.InvokeVoid(
		a,
		"resetUseMsi",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureadProvider) ResetUseOidc() {
	_jsii_.InvokeVoid(
		a,
		"resetUseOidc",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureadProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureadProvider) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureadProvider) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureadProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureadProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureadProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


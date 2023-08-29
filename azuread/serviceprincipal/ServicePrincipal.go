// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceprincipal

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azuread-go/azuread/v10/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azuread-go/azuread/v10/serviceprincipal/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azuread/2.41.0/docs/resources/service_principal azuread_service_principal}.
type ServicePrincipal interface {
	cdktf.TerraformResource
	AccountEnabled() interface{}
	SetAccountEnabled(val interface{})
	AccountEnabledInput() interface{}
	AlternativeNames() *[]*string
	SetAlternativeNames(val *[]*string)
	AlternativeNamesInput() *[]*string
	ApplicationId() *string
	SetApplicationId(val *string)
	ApplicationIdInput() *string
	ApplicationTenantId() *string
	AppRoleAssignmentRequired() interface{}
	SetAppRoleAssignmentRequired(val interface{})
	AppRoleAssignmentRequiredInput() interface{}
	AppRoleIds() cdktf.StringMap
	AppRoles() ServicePrincipalAppRolesList
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisplayName() *string
	Features() ServicePrincipalFeaturesList
	FeaturesInput() interface{}
	FeatureTags() ServicePrincipalFeatureTagsList
	FeatureTagsInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HomepageUrl() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoginUrl() *string
	SetLoginUrl(val *string)
	LoginUrlInput() *string
	LogoutUrl() *string
	// The tree node.
	Node() constructs.Node
	Notes() *string
	SetNotes(val *string)
	NotesInput() *string
	NotificationEmailAddresses() *[]*string
	SetNotificationEmailAddresses(val *[]*string)
	NotificationEmailAddressesInput() *[]*string
	Oauth2PermissionScopeIds() cdktf.StringMap
	Oauth2PermissionScopes() ServicePrincipalOauth2PermissionScopesList
	ObjectId() *string
	Owners() *[]*string
	SetOwners(val *[]*string)
	OwnersInput() *[]*string
	PreferredSingleSignOnMode() *string
	SetPreferredSingleSignOnMode(val *string)
	PreferredSingleSignOnModeInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	RedirectUris() *[]*string
	SamlMetadataUrl() *string
	SamlSingleSignOn() ServicePrincipalSamlSingleSignOnOutputReference
	SamlSingleSignOnInput() *ServicePrincipalSamlSingleSignOn
	ServicePrincipalNames() *[]*string
	SignInAudience() *string
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ServicePrincipalTimeoutsOutputReference
	TimeoutsInput() interface{}
	Type() *string
	UseExisting() interface{}
	SetUseExisting(val interface{})
	UseExistingInput() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutFeatures(value interface{})
	PutFeatureTags(value interface{})
	PutSamlSingleSignOn(value *ServicePrincipalSamlSingleSignOn)
	PutTimeouts(value *ServicePrincipalTimeouts)
	ResetAccountEnabled()
	ResetAlternativeNames()
	ResetAppRoleAssignmentRequired()
	ResetDescription()
	ResetFeatures()
	ResetFeatureTags()
	ResetId()
	ResetLoginUrl()
	ResetNotes()
	ResetNotificationEmailAddresses()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetOwners()
	ResetPreferredSingleSignOnMode()
	ResetSamlSingleSignOn()
	ResetTags()
	ResetTimeouts()
	ResetUseExisting()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ServicePrincipal
type jsiiProxy_ServicePrincipal struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ServicePrincipal) AccountEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accountEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) AccountEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accountEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) AlternativeNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"alternativeNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) AlternativeNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"alternativeNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) ApplicationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) ApplicationTenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationTenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) AppRoleAssignmentRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appRoleAssignmentRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) AppRoleAssignmentRequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appRoleAssignmentRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) AppRoleIds() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"appRoleIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) AppRoles() ServicePrincipalAppRolesList {
	var returns ServicePrincipalAppRolesList
	_jsii_.Get(
		j,
		"appRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) Features() ServicePrincipalFeaturesList {
	var returns ServicePrincipalFeaturesList
	_jsii_.Get(
		j,
		"features",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) FeaturesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"featuresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) FeatureTags() ServicePrincipalFeatureTagsList {
	var returns ServicePrincipalFeatureTagsList
	_jsii_.Get(
		j,
		"featureTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) FeatureTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"featureTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) HomepageUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homepageUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) LoginUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) LoginUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) LogoutUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoutUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) Notes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) NotesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) NotificationEmailAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notificationEmailAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) NotificationEmailAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notificationEmailAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) Oauth2PermissionScopeIds() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"oauth2PermissionScopeIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) Oauth2PermissionScopes() ServicePrincipalOauth2PermissionScopesList {
	var returns ServicePrincipalOauth2PermissionScopesList
	_jsii_.Get(
		j,
		"oauth2PermissionScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) ObjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) Owners() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"owners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) OwnersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ownersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) PreferredSingleSignOnMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredSingleSignOnMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) PreferredSingleSignOnModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredSingleSignOnModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) RedirectUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"redirectUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) SamlMetadataUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samlMetadataUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) SamlSingleSignOn() ServicePrincipalSamlSingleSignOnOutputReference {
	var returns ServicePrincipalSamlSingleSignOnOutputReference
	_jsii_.Get(
		j,
		"samlSingleSignOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) SamlSingleSignOnInput() *ServicePrincipalSamlSingleSignOn {
	var returns *ServicePrincipalSamlSingleSignOn
	_jsii_.Get(
		j,
		"samlSingleSignOnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) ServicePrincipalNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"servicePrincipalNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) SignInAudience() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signInAudience",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) Timeouts() ServicePrincipalTimeoutsOutputReference {
	var returns ServicePrincipalTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) UseExisting() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useExisting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) UseExistingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useExistingInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azuread/2.41.0/docs/resources/service_principal azuread_service_principal} Resource.
func NewServicePrincipal(scope constructs.Construct, id *string, config *ServicePrincipalConfig) ServicePrincipal {
	_init_.Initialize()

	if err := validateNewServicePrincipalParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServicePrincipal{}

	_jsii_.Create(
		"@cdktf/provider-azuread.servicePrincipal.ServicePrincipal",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azuread/2.41.0/docs/resources/service_principal azuread_service_principal} Resource.
func NewServicePrincipal_Override(s ServicePrincipal, scope constructs.Construct, id *string, config *ServicePrincipalConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azuread.servicePrincipal.ServicePrincipal",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_ServicePrincipal)SetAccountEnabled(val interface{}) {
	if err := j.validateSetAccountEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountEnabled",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipal)SetAlternativeNames(val *[]*string) {
	if err := j.validateSetAlternativeNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alternativeNames",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipal)SetApplicationId(val *string) {
	if err := j.validateSetApplicationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationId",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipal)SetAppRoleAssignmentRequired(val interface{}) {
	if err := j.validateSetAppRoleAssignmentRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appRoleAssignmentRequired",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipal)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipal)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipal)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipal)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipal)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipal)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipal)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipal)SetLoginUrl(val *string) {
	if err := j.validateSetLoginUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loginUrl",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipal)SetNotes(val *string) {
	if err := j.validateSetNotesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notes",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipal)SetNotificationEmailAddresses(val *[]*string) {
	if err := j.validateSetNotificationEmailAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notificationEmailAddresses",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipal)SetOwners(val *[]*string) {
	if err := j.validateSetOwnersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"owners",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipal)SetPreferredSingleSignOnMode(val *string) {
	if err := j.validateSetPreferredSingleSignOnModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredSingleSignOnMode",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipal)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipal)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipal)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipal)SetUseExisting(val interface{}) {
	if err := j.validateSetUseExistingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useExisting",
		val,
	)
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
func ServicePrincipal_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServicePrincipal_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azuread.servicePrincipal.ServicePrincipal",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ServicePrincipal_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServicePrincipal_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azuread.servicePrincipal.ServicePrincipal",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ServicePrincipal_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServicePrincipal_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azuread.servicePrincipal.ServicePrincipal",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ServicePrincipal_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azuread.servicePrincipal.ServicePrincipal",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_ServicePrincipal) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_ServicePrincipal) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipal) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipal) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipal) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipal) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipal) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipal) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipal) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipal) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipal) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipal) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_ServicePrincipal) PutFeatures(value interface{}) {
	if err := s.validatePutFeaturesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putFeatures",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServicePrincipal) PutFeatureTags(value interface{}) {
	if err := s.validatePutFeatureTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putFeatureTags",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServicePrincipal) PutSamlSingleSignOn(value *ServicePrincipalSamlSingleSignOn) {
	if err := s.validatePutSamlSingleSignOnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSamlSingleSignOn",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServicePrincipal) PutTimeouts(value *ServicePrincipalTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServicePrincipal) ResetAccountEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetAccountEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicePrincipal) ResetAlternativeNames() {
	_jsii_.InvokeVoid(
		s,
		"resetAlternativeNames",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicePrincipal) ResetAppRoleAssignmentRequired() {
	_jsii_.InvokeVoid(
		s,
		"resetAppRoleAssignmentRequired",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicePrincipal) ResetDescription() {
	_jsii_.InvokeVoid(
		s,
		"resetDescription",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicePrincipal) ResetFeatures() {
	_jsii_.InvokeVoid(
		s,
		"resetFeatures",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicePrincipal) ResetFeatureTags() {
	_jsii_.InvokeVoid(
		s,
		"resetFeatureTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicePrincipal) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicePrincipal) ResetLoginUrl() {
	_jsii_.InvokeVoid(
		s,
		"resetLoginUrl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicePrincipal) ResetNotes() {
	_jsii_.InvokeVoid(
		s,
		"resetNotes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicePrincipal) ResetNotificationEmailAddresses() {
	_jsii_.InvokeVoid(
		s,
		"resetNotificationEmailAddresses",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicePrincipal) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicePrincipal) ResetOwners() {
	_jsii_.InvokeVoid(
		s,
		"resetOwners",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicePrincipal) ResetPreferredSingleSignOnMode() {
	_jsii_.InvokeVoid(
		s,
		"resetPreferredSingleSignOnMode",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicePrincipal) ResetSamlSingleSignOn() {
	_jsii_.InvokeVoid(
		s,
		"resetSamlSingleSignOn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicePrincipal) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicePrincipal) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicePrincipal) ResetUseExisting() {
	_jsii_.InvokeVoid(
		s,
		"resetUseExisting",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicePrincipal) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipal) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipal) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipal) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


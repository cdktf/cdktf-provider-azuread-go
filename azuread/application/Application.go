// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package application

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azuread-go/azuread/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azuread-go/azuread/v14/application/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azuread/3.7.0/docs/resources/application azuread_application}.
type Application interface {
	cdktf.TerraformResource
	Api() ApplicationApiOutputReference
	ApiInput() *ApplicationApi
	AppRole() ApplicationAppRoleList
	AppRoleIds() cdktf.StringMap
	AppRoleInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientId() *string
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
	DeviceOnlyAuthEnabled() interface{}
	SetDeviceOnlyAuthEnabled(val interface{})
	DeviceOnlyAuthEnabledInput() interface{}
	DisabledByMicrosoft() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	FallbackPublicClientEnabled() interface{}
	SetFallbackPublicClientEnabled(val interface{})
	FallbackPublicClientEnabledInput() interface{}
	FeatureTags() ApplicationFeatureTagsList
	FeatureTagsInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GroupMembershipClaims() *[]*string
	SetGroupMembershipClaims(val *[]*string)
	GroupMembershipClaimsInput() *[]*string
	Id() *string
	SetId(val *string)
	IdentifierUris() *[]*string
	SetIdentifierUris(val *[]*string)
	IdentifierUrisInput() *[]*string
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogoImage() *string
	SetLogoImage(val *string)
	LogoImageInput() *string
	LogoUrl() *string
	MarketingUrl() *string
	SetMarketingUrl(val *string)
	MarketingUrlInput() *string
	// The tree node.
	Node() constructs.Node
	Notes() *string
	SetNotes(val *string)
	NotesInput() *string
	Oauth2PermissionScopeIds() cdktf.StringMap
	Oauth2PostResponseRequired() interface{}
	SetOauth2PostResponseRequired(val interface{})
	Oauth2PostResponseRequiredInput() interface{}
	ObjectId() *string
	OptionalClaims() ApplicationOptionalClaimsOutputReference
	OptionalClaimsInput() *ApplicationOptionalClaims
	Owners() *[]*string
	SetOwners(val *[]*string)
	OwnersInput() *[]*string
	Password() ApplicationPasswordOutputReference
	PasswordInput() *ApplicationPassword
	PreventDuplicateNames() interface{}
	SetPreventDuplicateNames(val interface{})
	PreventDuplicateNamesInput() interface{}
	PrivacyStatementUrl() *string
	SetPrivacyStatementUrl(val *string)
	PrivacyStatementUrlInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PublicClient() ApplicationPublicClientOutputReference
	PublicClientInput() *ApplicationPublicClient
	PublisherDomain() *string
	// Experimental.
	RawOverrides() interface{}
	RequiredResourceAccess() ApplicationRequiredResourceAccessList
	RequiredResourceAccessInput() interface{}
	ServiceManagementReference() *string
	SetServiceManagementReference(val *string)
	ServiceManagementReferenceInput() *string
	SignInAudience() *string
	SetSignInAudience(val *string)
	SignInAudienceInput() *string
	SinglePageApplication() ApplicationSinglePageApplicationOutputReference
	SinglePageApplicationInput() *ApplicationSinglePageApplication
	SupportUrl() *string
	SetSupportUrl(val *string)
	SupportUrlInput() *string
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
	TemplateId() *string
	SetTemplateId(val *string)
	TemplateIdInput() *string
	TermsOfServiceUrl() *string
	SetTermsOfServiceUrl(val *string)
	TermsOfServiceUrlInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ApplicationTimeoutsOutputReference
	TimeoutsInput() interface{}
	Web() ApplicationWebOutputReference
	WebInput() *ApplicationWeb
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutApi(value *ApplicationApi)
	PutAppRole(value interface{})
	PutFeatureTags(value interface{})
	PutOptionalClaims(value *ApplicationOptionalClaims)
	PutPassword(value *ApplicationPassword)
	PutPublicClient(value *ApplicationPublicClient)
	PutRequiredResourceAccess(value interface{})
	PutSinglePageApplication(value *ApplicationSinglePageApplication)
	PutTimeouts(value *ApplicationTimeouts)
	PutWeb(value *ApplicationWeb)
	ResetApi()
	ResetAppRole()
	ResetDescription()
	ResetDeviceOnlyAuthEnabled()
	ResetFallbackPublicClientEnabled()
	ResetFeatureTags()
	ResetGroupMembershipClaims()
	ResetId()
	ResetIdentifierUris()
	ResetLogoImage()
	ResetMarketingUrl()
	ResetNotes()
	ResetOauth2PostResponseRequired()
	ResetOptionalClaims()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetOwners()
	ResetPassword()
	ResetPreventDuplicateNames()
	ResetPrivacyStatementUrl()
	ResetPublicClient()
	ResetRequiredResourceAccess()
	ResetServiceManagementReference()
	ResetSignInAudience()
	ResetSinglePageApplication()
	ResetSupportUrl()
	ResetTags()
	ResetTemplateId()
	ResetTermsOfServiceUrl()
	ResetTimeouts()
	ResetWeb()
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

// The jsii proxy struct for Application
type jsiiProxy_Application struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Application) Api() ApplicationApiOutputReference {
	var returns ApplicationApiOutputReference
	_jsii_.Get(
		j,
		"api",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) ApiInput() *ApplicationApi {
	var returns *ApplicationApi
	_jsii_.Get(
		j,
		"apiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) AppRole() ApplicationAppRoleList {
	var returns ApplicationAppRoleList
	_jsii_.Get(
		j,
		"appRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) AppRoleIds() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"appRoleIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) AppRoleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) DeviceOnlyAuthEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deviceOnlyAuthEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) DeviceOnlyAuthEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deviceOnlyAuthEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) DisabledByMicrosoft() *string {
	var returns *string
	_jsii_.Get(
		j,
		"disabledByMicrosoft",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) FallbackPublicClientEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fallbackPublicClientEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) FallbackPublicClientEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fallbackPublicClientEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) FeatureTags() ApplicationFeatureTagsList {
	var returns ApplicationFeatureTagsList
	_jsii_.Get(
		j,
		"featureTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) FeatureTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"featureTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) GroupMembershipClaims() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupMembershipClaims",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) GroupMembershipClaimsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupMembershipClaimsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) IdentifierUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identifierUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) IdentifierUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identifierUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) LogoImage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) LogoImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) LogoUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) MarketingUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"marketingUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) MarketingUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"marketingUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) Notes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) NotesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) Oauth2PermissionScopeIds() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"oauth2PermissionScopeIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) Oauth2PostResponseRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oauth2PostResponseRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) Oauth2PostResponseRequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oauth2PostResponseRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) ObjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) OptionalClaims() ApplicationOptionalClaimsOutputReference {
	var returns ApplicationOptionalClaimsOutputReference
	_jsii_.Get(
		j,
		"optionalClaims",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) OptionalClaimsInput() *ApplicationOptionalClaims {
	var returns *ApplicationOptionalClaims
	_jsii_.Get(
		j,
		"optionalClaimsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) Owners() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"owners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) OwnersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ownersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) Password() ApplicationPasswordOutputReference {
	var returns ApplicationPasswordOutputReference
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) PasswordInput() *ApplicationPassword {
	var returns *ApplicationPassword
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) PreventDuplicateNames() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preventDuplicateNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) PreventDuplicateNamesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preventDuplicateNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) PrivacyStatementUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privacyStatementUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) PrivacyStatementUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privacyStatementUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) PublicClient() ApplicationPublicClientOutputReference {
	var returns ApplicationPublicClientOutputReference
	_jsii_.Get(
		j,
		"publicClient",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) PublicClientInput() *ApplicationPublicClient {
	var returns *ApplicationPublicClient
	_jsii_.Get(
		j,
		"publicClientInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) PublisherDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publisherDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) RequiredResourceAccess() ApplicationRequiredResourceAccessList {
	var returns ApplicationRequiredResourceAccessList
	_jsii_.Get(
		j,
		"requiredResourceAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) RequiredResourceAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requiredResourceAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) ServiceManagementReference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceManagementReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) ServiceManagementReferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceManagementReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) SignInAudience() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signInAudience",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) SignInAudienceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signInAudienceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) SinglePageApplication() ApplicationSinglePageApplicationOutputReference {
	var returns ApplicationSinglePageApplicationOutputReference
	_jsii_.Get(
		j,
		"singlePageApplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) SinglePageApplicationInput() *ApplicationSinglePageApplication {
	var returns *ApplicationSinglePageApplication
	_jsii_.Get(
		j,
		"singlePageApplicationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) SupportUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) SupportUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) TemplateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) TemplateIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) TermsOfServiceUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"termsOfServiceUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) TermsOfServiceUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"termsOfServiceUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) Timeouts() ApplicationTimeoutsOutputReference {
	var returns ApplicationTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) Web() ApplicationWebOutputReference {
	var returns ApplicationWebOutputReference
	_jsii_.Get(
		j,
		"web",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) WebInput() *ApplicationWeb {
	var returns *ApplicationWeb
	_jsii_.Get(
		j,
		"webInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azuread/3.7.0/docs/resources/application azuread_application} Resource.
func NewApplication(scope constructs.Construct, id *string, config *ApplicationConfig) Application {
	_init_.Initialize()

	if err := validateNewApplicationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Application{}

	_jsii_.Create(
		"@cdktf/provider-azuread.application.Application",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azuread/3.7.0/docs/resources/application azuread_application} Resource.
func NewApplication_Override(a Application, scope constructs.Construct, id *string, config *ApplicationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azuread.application.Application",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_Application)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Application)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Application)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Application)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Application)SetDeviceOnlyAuthEnabled(val interface{}) {
	if err := j.validateSetDeviceOnlyAuthEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceOnlyAuthEnabled",
		val,
	)
}

func (j *jsiiProxy_Application)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_Application)SetFallbackPublicClientEnabled(val interface{}) {
	if err := j.validateSetFallbackPublicClientEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fallbackPublicClientEnabled",
		val,
	)
}

func (j *jsiiProxy_Application)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Application)SetGroupMembershipClaims(val *[]*string) {
	if err := j.validateSetGroupMembershipClaimsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupMembershipClaims",
		val,
	)
}

func (j *jsiiProxy_Application)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Application)SetIdentifierUris(val *[]*string) {
	if err := j.validateSetIdentifierUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identifierUris",
		val,
	)
}

func (j *jsiiProxy_Application)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Application)SetLogoImage(val *string) {
	if err := j.validateSetLogoImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logoImage",
		val,
	)
}

func (j *jsiiProxy_Application)SetMarketingUrl(val *string) {
	if err := j.validateSetMarketingUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"marketingUrl",
		val,
	)
}

func (j *jsiiProxy_Application)SetNotes(val *string) {
	if err := j.validateSetNotesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notes",
		val,
	)
}

func (j *jsiiProxy_Application)SetOauth2PostResponseRequired(val interface{}) {
	if err := j.validateSetOauth2PostResponseRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauth2PostResponseRequired",
		val,
	)
}

func (j *jsiiProxy_Application)SetOwners(val *[]*string) {
	if err := j.validateSetOwnersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"owners",
		val,
	)
}

func (j *jsiiProxy_Application)SetPreventDuplicateNames(val interface{}) {
	if err := j.validateSetPreventDuplicateNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preventDuplicateNames",
		val,
	)
}

func (j *jsiiProxy_Application)SetPrivacyStatementUrl(val *string) {
	if err := j.validateSetPrivacyStatementUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privacyStatementUrl",
		val,
	)
}

func (j *jsiiProxy_Application)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Application)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Application)SetServiceManagementReference(val *string) {
	if err := j.validateSetServiceManagementReferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceManagementReference",
		val,
	)
}

func (j *jsiiProxy_Application)SetSignInAudience(val *string) {
	if err := j.validateSetSignInAudienceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signInAudience",
		val,
	)
}

func (j *jsiiProxy_Application)SetSupportUrl(val *string) {
	if err := j.validateSetSupportUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportUrl",
		val,
	)
}

func (j *jsiiProxy_Application)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_Application)SetTemplateId(val *string) {
	if err := j.validateSetTemplateIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"templateId",
		val,
	)
}

func (j *jsiiProxy_Application)SetTermsOfServiceUrl(val *string) {
	if err := j.validateSetTermsOfServiceUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"termsOfServiceUrl",
		val,
	)
}

// Generates CDKTF code for importing a Application resource upon running "cdktf plan <stack-name>".
func Application_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateApplication_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azuread.application.Application",
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
func Application_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApplication_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azuread.application.Application",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Application_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApplication_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azuread.application.Application",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Application_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApplication_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azuread.application.Application",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Application_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azuread.application.Application",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_Application) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_Application) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_Application) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Application) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Application) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Application) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Application) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Application) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Application) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Application) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Application) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Application) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Application) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_Application) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Application) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_Application) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_Application) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_Application) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_Application) PutApi(value *ApplicationApi) {
	if err := a.validatePutApiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putApi",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Application) PutAppRole(value interface{}) {
	if err := a.validatePutAppRoleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAppRole",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Application) PutFeatureTags(value interface{}) {
	if err := a.validatePutFeatureTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFeatureTags",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Application) PutOptionalClaims(value *ApplicationOptionalClaims) {
	if err := a.validatePutOptionalClaimsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOptionalClaims",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Application) PutPassword(value *ApplicationPassword) {
	if err := a.validatePutPasswordParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPassword",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Application) PutPublicClient(value *ApplicationPublicClient) {
	if err := a.validatePutPublicClientParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPublicClient",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Application) PutRequiredResourceAccess(value interface{}) {
	if err := a.validatePutRequiredResourceAccessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRequiredResourceAccess",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Application) PutSinglePageApplication(value *ApplicationSinglePageApplication) {
	if err := a.validatePutSinglePageApplicationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSinglePageApplication",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Application) PutTimeouts(value *ApplicationTimeouts) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Application) PutWeb(value *ApplicationWeb) {
	if err := a.validatePutWebParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putWeb",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_Application) ResetApi() {
	_jsii_.InvokeVoid(
		a,
		"resetApi",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Application) ResetAppRole() {
	_jsii_.InvokeVoid(
		a,
		"resetAppRole",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Application) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Application) ResetDeviceOnlyAuthEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetDeviceOnlyAuthEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Application) ResetFallbackPublicClientEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetFallbackPublicClientEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Application) ResetFeatureTags() {
	_jsii_.InvokeVoid(
		a,
		"resetFeatureTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Application) ResetGroupMembershipClaims() {
	_jsii_.InvokeVoid(
		a,
		"resetGroupMembershipClaims",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Application) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Application) ResetIdentifierUris() {
	_jsii_.InvokeVoid(
		a,
		"resetIdentifierUris",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Application) ResetLogoImage() {
	_jsii_.InvokeVoid(
		a,
		"resetLogoImage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Application) ResetMarketingUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetMarketingUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Application) ResetNotes() {
	_jsii_.InvokeVoid(
		a,
		"resetNotes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Application) ResetOauth2PostResponseRequired() {
	_jsii_.InvokeVoid(
		a,
		"resetOauth2PostResponseRequired",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Application) ResetOptionalClaims() {
	_jsii_.InvokeVoid(
		a,
		"resetOptionalClaims",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Application) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Application) ResetOwners() {
	_jsii_.InvokeVoid(
		a,
		"resetOwners",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Application) ResetPassword() {
	_jsii_.InvokeVoid(
		a,
		"resetPassword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Application) ResetPreventDuplicateNames() {
	_jsii_.InvokeVoid(
		a,
		"resetPreventDuplicateNames",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Application) ResetPrivacyStatementUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetPrivacyStatementUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Application) ResetPublicClient() {
	_jsii_.InvokeVoid(
		a,
		"resetPublicClient",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Application) ResetRequiredResourceAccess() {
	_jsii_.InvokeVoid(
		a,
		"resetRequiredResourceAccess",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Application) ResetServiceManagementReference() {
	_jsii_.InvokeVoid(
		a,
		"resetServiceManagementReference",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Application) ResetSignInAudience() {
	_jsii_.InvokeVoid(
		a,
		"resetSignInAudience",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Application) ResetSinglePageApplication() {
	_jsii_.InvokeVoid(
		a,
		"resetSinglePageApplication",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Application) ResetSupportUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetSupportUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Application) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Application) ResetTemplateId() {
	_jsii_.InvokeVoid(
		a,
		"resetTemplateId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Application) ResetTermsOfServiceUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetTermsOfServiceUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Application) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Application) ResetWeb() {
	_jsii_.InvokeVoid(
		a,
		"resetWeb",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Application) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Application) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Application) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Application) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Application) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Application) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


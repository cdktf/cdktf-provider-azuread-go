// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationregistration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azuread-go/azuread/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azuread-go/azuread/v13/applicationregistration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azuread/3.3.0/docs/resources/application_registration azuread_application_registration}.
type ApplicationRegistration interface {
	cdktf.TerraformResource
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
	DisabledByMicrosoft() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
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
	HomepageUrl() *string
	SetHomepageUrl(val *string)
	HomepageUrlInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	ImplicitAccessTokenIssuanceEnabled() interface{}
	SetImplicitAccessTokenIssuanceEnabled(val interface{})
	ImplicitAccessTokenIssuanceEnabledInput() interface{}
	ImplicitIdTokenIssuanceEnabled() interface{}
	SetImplicitIdTokenIssuanceEnabled(val interface{})
	ImplicitIdTokenIssuanceEnabledInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogoutUrl() *string
	SetLogoutUrl(val *string)
	LogoutUrlInput() *string
	MarketingUrl() *string
	SetMarketingUrl(val *string)
	MarketingUrlInput() *string
	// The tree node.
	Node() constructs.Node
	Notes() *string
	SetNotes(val *string)
	NotesInput() *string
	ObjectId() *string
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
	PublisherDomain() *string
	// Experimental.
	RawOverrides() interface{}
	RequestedAccessTokenVersion() *float64
	SetRequestedAccessTokenVersion(val *float64)
	RequestedAccessTokenVersionInput() *float64
	ServiceManagementReference() *string
	SetServiceManagementReference(val *string)
	ServiceManagementReferenceInput() *string
	SignInAudience() *string
	SetSignInAudience(val *string)
	SignInAudienceInput() *string
	SupportUrl() *string
	SetSupportUrl(val *string)
	SupportUrlInput() *string
	TermsOfServiceUrl() *string
	SetTermsOfServiceUrl(val *string)
	TermsOfServiceUrlInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ApplicationRegistrationTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutTimeouts(value *ApplicationRegistrationTimeouts)
	ResetDescription()
	ResetGroupMembershipClaims()
	ResetHomepageUrl()
	ResetId()
	ResetImplicitAccessTokenIssuanceEnabled()
	ResetImplicitIdTokenIssuanceEnabled()
	ResetLogoutUrl()
	ResetMarketingUrl()
	ResetNotes()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrivacyStatementUrl()
	ResetRequestedAccessTokenVersion()
	ResetServiceManagementReference()
	ResetSignInAudience()
	ResetSupportUrl()
	ResetTermsOfServiceUrl()
	ResetTimeouts()
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

// The jsii proxy struct for ApplicationRegistration
type jsiiProxy_ApplicationRegistration struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApplicationRegistration) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) DisabledByMicrosoft() *string {
	var returns *string
	_jsii_.Get(
		j,
		"disabledByMicrosoft",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) GroupMembershipClaims() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupMembershipClaims",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) GroupMembershipClaimsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupMembershipClaimsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) HomepageUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homepageUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) HomepageUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homepageUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) ImplicitAccessTokenIssuanceEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"implicitAccessTokenIssuanceEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) ImplicitAccessTokenIssuanceEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"implicitAccessTokenIssuanceEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) ImplicitIdTokenIssuanceEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"implicitIdTokenIssuanceEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) ImplicitIdTokenIssuanceEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"implicitIdTokenIssuanceEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) LogoutUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoutUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) LogoutUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoutUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) MarketingUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"marketingUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) MarketingUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"marketingUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) Notes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) NotesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) ObjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) PrivacyStatementUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privacyStatementUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) PrivacyStatementUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privacyStatementUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) PublisherDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publisherDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) RequestedAccessTokenVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestedAccessTokenVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) RequestedAccessTokenVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestedAccessTokenVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) ServiceManagementReference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceManagementReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) ServiceManagementReferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceManagementReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) SignInAudience() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signInAudience",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) SignInAudienceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signInAudienceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) SupportUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) SupportUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) TermsOfServiceUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"termsOfServiceUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) TermsOfServiceUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"termsOfServiceUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) Timeouts() ApplicationRegistrationTimeoutsOutputReference {
	var returns ApplicationRegistrationTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationRegistration) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azuread/3.3.0/docs/resources/application_registration azuread_application_registration} Resource.
func NewApplicationRegistration(scope constructs.Construct, id *string, config *ApplicationRegistrationConfig) ApplicationRegistration {
	_init_.Initialize()

	if err := validateNewApplicationRegistrationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationRegistration{}

	_jsii_.Create(
		"@cdktf/provider-azuread.applicationRegistration.ApplicationRegistration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azuread/3.3.0/docs/resources/application_registration azuread_application_registration} Resource.
func NewApplicationRegistration_Override(a ApplicationRegistration, scope constructs.Construct, id *string, config *ApplicationRegistrationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azuread.applicationRegistration.ApplicationRegistration",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApplicationRegistration)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ApplicationRegistration)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApplicationRegistration)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApplicationRegistration)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ApplicationRegistration)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_ApplicationRegistration)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ApplicationRegistration)SetGroupMembershipClaims(val *[]*string) {
	if err := j.validateSetGroupMembershipClaimsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupMembershipClaims",
		val,
	)
}

func (j *jsiiProxy_ApplicationRegistration)SetHomepageUrl(val *string) {
	if err := j.validateSetHomepageUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"homepageUrl",
		val,
	)
}

func (j *jsiiProxy_ApplicationRegistration)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ApplicationRegistration)SetImplicitAccessTokenIssuanceEnabled(val interface{}) {
	if err := j.validateSetImplicitAccessTokenIssuanceEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"implicitAccessTokenIssuanceEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationRegistration)SetImplicitIdTokenIssuanceEnabled(val interface{}) {
	if err := j.validateSetImplicitIdTokenIssuanceEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"implicitIdTokenIssuanceEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationRegistration)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApplicationRegistration)SetLogoutUrl(val *string) {
	if err := j.validateSetLogoutUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logoutUrl",
		val,
	)
}

func (j *jsiiProxy_ApplicationRegistration)SetMarketingUrl(val *string) {
	if err := j.validateSetMarketingUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"marketingUrl",
		val,
	)
}

func (j *jsiiProxy_ApplicationRegistration)SetNotes(val *string) {
	if err := j.validateSetNotesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notes",
		val,
	)
}

func (j *jsiiProxy_ApplicationRegistration)SetPrivacyStatementUrl(val *string) {
	if err := j.validateSetPrivacyStatementUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privacyStatementUrl",
		val,
	)
}

func (j *jsiiProxy_ApplicationRegistration)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApplicationRegistration)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ApplicationRegistration)SetRequestedAccessTokenVersion(val *float64) {
	if err := j.validateSetRequestedAccessTokenVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestedAccessTokenVersion",
		val,
	)
}

func (j *jsiiProxy_ApplicationRegistration)SetServiceManagementReference(val *string) {
	if err := j.validateSetServiceManagementReferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceManagementReference",
		val,
	)
}

func (j *jsiiProxy_ApplicationRegistration)SetSignInAudience(val *string) {
	if err := j.validateSetSignInAudienceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signInAudience",
		val,
	)
}

func (j *jsiiProxy_ApplicationRegistration)SetSupportUrl(val *string) {
	if err := j.validateSetSupportUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportUrl",
		val,
	)
}

func (j *jsiiProxy_ApplicationRegistration)SetTermsOfServiceUrl(val *string) {
	if err := j.validateSetTermsOfServiceUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"termsOfServiceUrl",
		val,
	)
}

// Generates CDKTF code for importing a ApplicationRegistration resource upon running "cdktf plan <stack-name>".
func ApplicationRegistration_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateApplicationRegistration_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azuread.applicationRegistration.ApplicationRegistration",
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
func ApplicationRegistration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApplicationRegistration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azuread.applicationRegistration.ApplicationRegistration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApplicationRegistration_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApplicationRegistration_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azuread.applicationRegistration.ApplicationRegistration",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApplicationRegistration_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApplicationRegistration_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azuread.applicationRegistration.ApplicationRegistration",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApplicationRegistration_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azuread.applicationRegistration.ApplicationRegistration",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_ApplicationRegistration) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_ApplicationRegistration) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_ApplicationRegistration) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApplicationRegistration) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApplicationRegistration) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApplicationRegistration) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApplicationRegistration) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApplicationRegistration) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApplicationRegistration) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApplicationRegistration) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApplicationRegistration) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApplicationRegistration) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationRegistration) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_ApplicationRegistration) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApplicationRegistration) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ApplicationRegistration) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_ApplicationRegistration) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ApplicationRegistration) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApplicationRegistration) PutTimeouts(value *ApplicationRegistrationTimeouts) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationRegistration) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationRegistration) ResetGroupMembershipClaims() {
	_jsii_.InvokeVoid(
		a,
		"resetGroupMembershipClaims",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationRegistration) ResetHomepageUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetHomepageUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationRegistration) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationRegistration) ResetImplicitAccessTokenIssuanceEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetImplicitAccessTokenIssuanceEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationRegistration) ResetImplicitIdTokenIssuanceEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetImplicitIdTokenIssuanceEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationRegistration) ResetLogoutUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetLogoutUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationRegistration) ResetMarketingUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetMarketingUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationRegistration) ResetNotes() {
	_jsii_.InvokeVoid(
		a,
		"resetNotes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationRegistration) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationRegistration) ResetPrivacyStatementUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetPrivacyStatementUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationRegistration) ResetRequestedAccessTokenVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestedAccessTokenVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationRegistration) ResetServiceManagementReference() {
	_jsii_.InvokeVoid(
		a,
		"resetServiceManagementReference",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationRegistration) ResetSignInAudience() {
	_jsii_.InvokeVoid(
		a,
		"resetSignInAudience",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationRegistration) ResetSupportUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetSupportUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationRegistration) ResetTermsOfServiceUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetTermsOfServiceUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationRegistration) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationRegistration) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationRegistration) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationRegistration) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationRegistration) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationRegistration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationRegistration) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


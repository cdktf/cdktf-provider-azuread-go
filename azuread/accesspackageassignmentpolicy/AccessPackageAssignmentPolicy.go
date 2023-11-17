// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesspackageassignmentpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azuread-go/azuread/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azuread-go/azuread/v11/accesspackageassignmentpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azuread/2.46.0/docs/resources/access_package_assignment_policy azuread_access_package_assignment_policy}.
type AccessPackageAssignmentPolicy interface {
	cdktf.TerraformResource
	AccessPackageId() *string
	SetAccessPackageId(val *string)
	AccessPackageIdInput() *string
	ApprovalSettings() AccessPackageAssignmentPolicyApprovalSettingsOutputReference
	ApprovalSettingsInput() *AccessPackageAssignmentPolicyApprovalSettings
	AssignmentReviewSettings() AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference
	AssignmentReviewSettingsInput() *AccessPackageAssignmentPolicyAssignmentReviewSettings
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
	SetDisplayName(val *string)
	DisplayNameInput() *string
	DurationInDays() *float64
	SetDurationInDays(val *float64)
	DurationInDaysInput() *float64
	ExpirationDate() *string
	SetExpirationDate(val *string)
	ExpirationDateInput() *string
	ExtensionEnabled() interface{}
	SetExtensionEnabled(val interface{})
	ExtensionEnabledInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	Question() AccessPackageAssignmentPolicyQuestionList
	QuestionInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	RequestorSettings() AccessPackageAssignmentPolicyRequestorSettingsOutputReference
	RequestorSettingsInput() *AccessPackageAssignmentPolicyRequestorSettings
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() AccessPackageAssignmentPolicyTimeoutsOutputReference
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
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutApprovalSettings(value *AccessPackageAssignmentPolicyApprovalSettings)
	PutAssignmentReviewSettings(value *AccessPackageAssignmentPolicyAssignmentReviewSettings)
	PutQuestion(value interface{})
	PutRequestorSettings(value *AccessPackageAssignmentPolicyRequestorSettings)
	PutTimeouts(value *AccessPackageAssignmentPolicyTimeouts)
	ResetApprovalSettings()
	ResetAssignmentReviewSettings()
	ResetDurationInDays()
	ResetExpirationDate()
	ResetExtensionEnabled()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetQuestion()
	ResetRequestorSettings()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for AccessPackageAssignmentPolicy
type jsiiProxy_AccessPackageAssignmentPolicy struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy) AccessPackageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessPackageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy) AccessPackageIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessPackageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy) ApprovalSettings() AccessPackageAssignmentPolicyApprovalSettingsOutputReference {
	var returns AccessPackageAssignmentPolicyApprovalSettingsOutputReference
	_jsii_.Get(
		j,
		"approvalSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy) ApprovalSettingsInput() *AccessPackageAssignmentPolicyApprovalSettings {
	var returns *AccessPackageAssignmentPolicyApprovalSettings
	_jsii_.Get(
		j,
		"approvalSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy) AssignmentReviewSettings() AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference {
	var returns AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference
	_jsii_.Get(
		j,
		"assignmentReviewSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy) AssignmentReviewSettingsInput() *AccessPackageAssignmentPolicyAssignmentReviewSettings {
	var returns *AccessPackageAssignmentPolicyAssignmentReviewSettings
	_jsii_.Get(
		j,
		"assignmentReviewSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy) DurationInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"durationInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy) DurationInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"durationInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy) ExpirationDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expirationDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy) ExpirationDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expirationDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy) ExtensionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extensionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy) ExtensionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extensionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy) Question() AccessPackageAssignmentPolicyQuestionList {
	var returns AccessPackageAssignmentPolicyQuestionList
	_jsii_.Get(
		j,
		"question",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy) QuestionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"questionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy) RequestorSettings() AccessPackageAssignmentPolicyRequestorSettingsOutputReference {
	var returns AccessPackageAssignmentPolicyRequestorSettingsOutputReference
	_jsii_.Get(
		j,
		"requestorSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy) RequestorSettingsInput() *AccessPackageAssignmentPolicyRequestorSettings {
	var returns *AccessPackageAssignmentPolicyRequestorSettings
	_jsii_.Get(
		j,
		"requestorSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy) Timeouts() AccessPackageAssignmentPolicyTimeoutsOutputReference {
	var returns AccessPackageAssignmentPolicyTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azuread/2.46.0/docs/resources/access_package_assignment_policy azuread_access_package_assignment_policy} Resource.
func NewAccessPackageAssignmentPolicy(scope constructs.Construct, id *string, config *AccessPackageAssignmentPolicyConfig) AccessPackageAssignmentPolicy {
	_init_.Initialize()

	if err := validateNewAccessPackageAssignmentPolicyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccessPackageAssignmentPolicy{}

	_jsii_.Create(
		"@cdktf/provider-azuread.accessPackageAssignmentPolicy.AccessPackageAssignmentPolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azuread/2.46.0/docs/resources/access_package_assignment_policy azuread_access_package_assignment_policy} Resource.
func NewAccessPackageAssignmentPolicy_Override(a AccessPackageAssignmentPolicy, scope constructs.Construct, id *string, config *AccessPackageAssignmentPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azuread.accessPackageAssignmentPolicy.AccessPackageAssignmentPolicy",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy)SetAccessPackageId(val *string) {
	if err := j.validateSetAccessPackageIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessPackageId",
		val,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy)SetDurationInDays(val *float64) {
	if err := j.validateSetDurationInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"durationInDays",
		val,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy)SetExpirationDate(val *string) {
	if err := j.validateSetExpirationDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expirationDate",
		val,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy)SetExtensionEnabled(val interface{}) {
	if err := j.validateSetExtensionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extensionEnabled",
		val,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicy)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a AccessPackageAssignmentPolicy resource upon running "cdktf plan <stack-name>".
func AccessPackageAssignmentPolicy_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateAccessPackageAssignmentPolicy_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azuread.accessPackageAssignmentPolicy.AccessPackageAssignmentPolicy",
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
func AccessPackageAssignmentPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAccessPackageAssignmentPolicy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azuread.accessPackageAssignmentPolicy.AccessPackageAssignmentPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AccessPackageAssignmentPolicy_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAccessPackageAssignmentPolicy_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azuread.accessPackageAssignmentPolicy.AccessPackageAssignmentPolicy",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AccessPackageAssignmentPolicy_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAccessPackageAssignmentPolicy_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azuread.accessPackageAssignmentPolicy.AccessPackageAssignmentPolicy",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AccessPackageAssignmentPolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azuread.accessPackageAssignmentPolicy.AccessPackageAssignmentPolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AccessPackageAssignmentPolicy) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AccessPackageAssignmentPolicy) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AccessPackageAssignmentPolicy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicy) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicy) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AccessPackageAssignmentPolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicy) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AccessPackageAssignmentPolicy) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AccessPackageAssignmentPolicy) PutApprovalSettings(value *AccessPackageAssignmentPolicyApprovalSettings) {
	if err := a.validatePutApprovalSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putApprovalSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccessPackageAssignmentPolicy) PutAssignmentReviewSettings(value *AccessPackageAssignmentPolicyAssignmentReviewSettings) {
	if err := a.validatePutAssignmentReviewSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAssignmentReviewSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccessPackageAssignmentPolicy) PutQuestion(value interface{}) {
	if err := a.validatePutQuestionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putQuestion",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccessPackageAssignmentPolicy) PutRequestorSettings(value *AccessPackageAssignmentPolicyRequestorSettings) {
	if err := a.validatePutRequestorSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRequestorSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccessPackageAssignmentPolicy) PutTimeouts(value *AccessPackageAssignmentPolicyTimeouts) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccessPackageAssignmentPolicy) ResetApprovalSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetApprovalSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPackageAssignmentPolicy) ResetAssignmentReviewSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetAssignmentReviewSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPackageAssignmentPolicy) ResetDurationInDays() {
	_jsii_.InvokeVoid(
		a,
		"resetDurationInDays",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPackageAssignmentPolicy) ResetExpirationDate() {
	_jsii_.InvokeVoid(
		a,
		"resetExpirationDate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPackageAssignmentPolicy) ResetExtensionEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetExtensionEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPackageAssignmentPolicy) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPackageAssignmentPolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPackageAssignmentPolicy) ResetQuestion() {
	_jsii_.InvokeVoid(
		a,
		"resetQuestion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPackageAssignmentPolicy) ResetRequestorSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestorSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPackageAssignmentPolicy) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPackageAssignmentPolicy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessPackageAssignmentPolicy) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessPackageAssignmentPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessPackageAssignmentPolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


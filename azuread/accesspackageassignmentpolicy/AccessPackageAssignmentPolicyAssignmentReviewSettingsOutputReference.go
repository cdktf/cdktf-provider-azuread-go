// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesspackageassignmentpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azuread-go/azuread/v11/jsii"

	"github.com/cdktf/cdktf-provider-azuread-go/azuread/v11/accesspackageassignmentpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference interface {
	cdktf.ComplexObject
	AccessRecommendationEnabled() interface{}
	SetAccessRecommendationEnabled(val interface{})
	AccessRecommendationEnabledInput() interface{}
	AccessReviewTimeoutBehavior() *string
	SetAccessReviewTimeoutBehavior(val *string)
	AccessReviewTimeoutBehaviorInput() *string
	ApproverJustificationRequired() interface{}
	SetApproverJustificationRequired(val interface{})
	ApproverJustificationRequiredInput() interface{}
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DurationInDays() *float64
	SetDurationInDays(val *float64)
	DurationInDaysInput() *float64
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *AccessPackageAssignmentPolicyAssignmentReviewSettings
	SetInternalValue(val *AccessPackageAssignmentPolicyAssignmentReviewSettings)
	Reviewer() AccessPackageAssignmentPolicyAssignmentReviewSettingsReviewerList
	ReviewerInput() interface{}
	ReviewFrequency() *string
	SetReviewFrequency(val *string)
	ReviewFrequencyInput() *string
	ReviewType() *string
	SetReviewType(val *string)
	ReviewTypeInput() *string
	StartingOn() *string
	SetStartingOn(val *string)
	StartingOnInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutReviewer(value interface{})
	ResetAccessRecommendationEnabled()
	ResetAccessReviewTimeoutBehavior()
	ResetApproverJustificationRequired()
	ResetDurationInDays()
	ResetEnabled()
	ResetReviewer()
	ResetReviewFrequency()
	ResetReviewType()
	ResetStartingOn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference
type jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference) AccessRecommendationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessRecommendationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference) AccessRecommendationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessRecommendationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference) AccessReviewTimeoutBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessReviewTimeoutBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference) AccessReviewTimeoutBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessReviewTimeoutBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference) ApproverJustificationRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"approverJustificationRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference) ApproverJustificationRequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"approverJustificationRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference) DurationInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"durationInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference) DurationInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"durationInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference) InternalValue() *AccessPackageAssignmentPolicyAssignmentReviewSettings {
	var returns *AccessPackageAssignmentPolicyAssignmentReviewSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference) Reviewer() AccessPackageAssignmentPolicyAssignmentReviewSettingsReviewerList {
	var returns AccessPackageAssignmentPolicyAssignmentReviewSettingsReviewerList
	_jsii_.Get(
		j,
		"reviewer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference) ReviewerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reviewerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference) ReviewFrequency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reviewFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference) ReviewFrequencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reviewFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference) ReviewType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reviewType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference) ReviewTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reviewTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference) StartingOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startingOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference) StartingOnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startingOnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewAccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azuread.accessPackageAssignmentPolicy.AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference_Override(a AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azuread.accessPackageAssignmentPolicy.AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference)SetAccessRecommendationEnabled(val interface{}) {
	if err := j.validateSetAccessRecommendationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessRecommendationEnabled",
		val,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference)SetAccessReviewTimeoutBehavior(val *string) {
	if err := j.validateSetAccessReviewTimeoutBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessReviewTimeoutBehavior",
		val,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference)SetApproverJustificationRequired(val interface{}) {
	if err := j.validateSetApproverJustificationRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"approverJustificationRequired",
		val,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference)SetDurationInDays(val *float64) {
	if err := j.validateSetDurationInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"durationInDays",
		val,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference)SetInternalValue(val *AccessPackageAssignmentPolicyAssignmentReviewSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference)SetReviewFrequency(val *string) {
	if err := j.validateSetReviewFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reviewFrequency",
		val,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference)SetReviewType(val *string) {
	if err := j.validateSetReviewTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reviewType",
		val,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference)SetStartingOn(val *string) {
	if err := j.validateSetStartingOnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startingOn",
		val,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference) PutReviewer(value interface{}) {
	if err := a.validatePutReviewerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putReviewer",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference) ResetAccessRecommendationEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessRecommendationEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference) ResetAccessReviewTimeoutBehavior() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessReviewTimeoutBehavior",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference) ResetApproverJustificationRequired() {
	_jsii_.InvokeVoid(
		a,
		"resetApproverJustificationRequired",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference) ResetDurationInDays() {
	_jsii_.InvokeVoid(
		a,
		"resetDurationInDays",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference) ResetReviewer() {
	_jsii_.InvokeVoid(
		a,
		"resetReviewer",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference) ResetReviewFrequency() {
	_jsii_.InvokeVoid(
		a,
		"resetReviewFrequency",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference) ResetReviewType() {
	_jsii_.InvokeVoid(
		a,
		"resetReviewType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference) ResetStartingOn() {
	_jsii_.InvokeVoid(
		a,
		"resetStartingOn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessPackageAssignmentPolicyAssignmentReviewSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


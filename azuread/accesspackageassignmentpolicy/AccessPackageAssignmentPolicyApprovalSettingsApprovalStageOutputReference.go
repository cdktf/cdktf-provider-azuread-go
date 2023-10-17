// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesspackageassignmentpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azuread-go/azuread/v11/jsii"

	"github.com/cdktf/cdktf-provider-azuread-go/azuread/v11/accesspackageassignmentpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference interface {
	cdktf.ComplexObject
	AlternativeApprovalEnabled() interface{}
	SetAlternativeApprovalEnabled(val interface{})
	AlternativeApprovalEnabledInput() interface{}
	AlternativeApprover() AccessPackageAssignmentPolicyApprovalSettingsApprovalStageAlternativeApproverList
	AlternativeApproverInput() interface{}
	ApprovalTimeoutInDays() *float64
	SetApprovalTimeoutInDays(val *float64)
	ApprovalTimeoutInDaysInput() *float64
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
	EnableAlternativeApprovalInDays() *float64
	SetEnableAlternativeApprovalInDays(val *float64)
	EnableAlternativeApprovalInDaysInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PrimaryApprover() AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverList
	PrimaryApproverInput() interface{}
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
	PutAlternativeApprover(value interface{})
	PutPrimaryApprover(value interface{})
	ResetAlternativeApprovalEnabled()
	ResetAlternativeApprover()
	ResetApproverJustificationRequired()
	ResetEnableAlternativeApprovalInDays()
	ResetPrimaryApprover()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference
type jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference) AlternativeApprovalEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alternativeApprovalEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference) AlternativeApprovalEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alternativeApprovalEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference) AlternativeApprover() AccessPackageAssignmentPolicyApprovalSettingsApprovalStageAlternativeApproverList {
	var returns AccessPackageAssignmentPolicyApprovalSettingsApprovalStageAlternativeApproverList
	_jsii_.Get(
		j,
		"alternativeApprover",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference) AlternativeApproverInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alternativeApproverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference) ApprovalTimeoutInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"approvalTimeoutInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference) ApprovalTimeoutInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"approvalTimeoutInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference) ApproverJustificationRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"approverJustificationRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference) ApproverJustificationRequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"approverJustificationRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference) EnableAlternativeApprovalInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"enableAlternativeApprovalInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference) EnableAlternativeApprovalInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"enableAlternativeApprovalInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference) PrimaryApprover() AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverList {
	var returns AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverList
	_jsii_.Get(
		j,
		"primaryApprover",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference) PrimaryApproverInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"primaryApproverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference {
	_init_.Initialize()

	if err := validateNewAccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azuread.accessPackageAssignmentPolicy.AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference_Override(a AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azuread.accessPackageAssignmentPolicy.AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference)SetAlternativeApprovalEnabled(val interface{}) {
	if err := j.validateSetAlternativeApprovalEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alternativeApprovalEnabled",
		val,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference)SetApprovalTimeoutInDays(val *float64) {
	if err := j.validateSetApprovalTimeoutInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"approvalTimeoutInDays",
		val,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference)SetApproverJustificationRequired(val interface{}) {
	if err := j.validateSetApproverJustificationRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"approverJustificationRequired",
		val,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference)SetEnableAlternativeApprovalInDays(val *float64) {
	if err := j.validateSetEnableAlternativeApprovalInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableAlternativeApprovalInDays",
		val,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference) PutAlternativeApprover(value interface{}) {
	if err := a.validatePutAlternativeApproverParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAlternativeApprover",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference) PutPrimaryApprover(value interface{}) {
	if err := a.validatePutPrimaryApproverParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPrimaryApprover",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference) ResetAlternativeApprovalEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetAlternativeApprovalEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference) ResetAlternativeApprover() {
	_jsii_.InvokeVoid(
		a,
		"resetAlternativeApprover",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference) ResetApproverJustificationRequired() {
	_jsii_.InvokeVoid(
		a,
		"resetApproverJustificationRequired",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference) ResetEnableAlternativeApprovalInDays() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableAlternativeApprovalInDays",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference) ResetPrimaryApprover() {
	_jsii_.InvokeVoid(
		a,
		"resetPrimaryApprover",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStageOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


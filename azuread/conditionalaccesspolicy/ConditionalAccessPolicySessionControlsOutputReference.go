// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package conditionalaccesspolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azuread-go/azuread/v13/jsii"

	"github.com/cdktf/cdktf-provider-azuread-go/azuread/v13/conditionalaccesspolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConditionalAccessPolicySessionControlsOutputReference interface {
	cdktf.ComplexObject
	ApplicationEnforcedRestrictionsEnabled() interface{}
	SetApplicationEnforcedRestrictionsEnabled(val interface{})
	ApplicationEnforcedRestrictionsEnabledInput() interface{}
	CloudAppSecurityPolicy() *string
	SetCloudAppSecurityPolicy(val *string)
	CloudAppSecurityPolicyInput() *string
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
	DisableResilienceDefaults() interface{}
	SetDisableResilienceDefaults(val interface{})
	DisableResilienceDefaultsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *ConditionalAccessPolicySessionControls
	SetInternalValue(val *ConditionalAccessPolicySessionControls)
	PersistentBrowserMode() *string
	SetPersistentBrowserMode(val *string)
	PersistentBrowserModeInput() *string
	SignInFrequency() *float64
	SetSignInFrequency(val *float64)
	SignInFrequencyAuthenticationType() *string
	SetSignInFrequencyAuthenticationType(val *string)
	SignInFrequencyAuthenticationTypeInput() *string
	SignInFrequencyInput() *float64
	SignInFrequencyInterval() *string
	SetSignInFrequencyInterval(val *string)
	SignInFrequencyIntervalInput() *string
	SignInFrequencyPeriod() *string
	SetSignInFrequencyPeriod(val *string)
	SignInFrequencyPeriodInput() *string
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
	ResetApplicationEnforcedRestrictionsEnabled()
	ResetCloudAppSecurityPolicy()
	ResetDisableResilienceDefaults()
	ResetPersistentBrowserMode()
	ResetSignInFrequency()
	ResetSignInFrequencyAuthenticationType()
	ResetSignInFrequencyInterval()
	ResetSignInFrequencyPeriod()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConditionalAccessPolicySessionControlsOutputReference
type jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference) ApplicationEnforcedRestrictionsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applicationEnforcedRestrictionsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference) ApplicationEnforcedRestrictionsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applicationEnforcedRestrictionsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference) CloudAppSecurityPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudAppSecurityPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference) CloudAppSecurityPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudAppSecurityPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference) DisableResilienceDefaults() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableResilienceDefaults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference) DisableResilienceDefaultsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableResilienceDefaultsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference) InternalValue() *ConditionalAccessPolicySessionControls {
	var returns *ConditionalAccessPolicySessionControls
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference) PersistentBrowserMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"persistentBrowserMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference) PersistentBrowserModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"persistentBrowserModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference) SignInFrequency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"signInFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference) SignInFrequencyAuthenticationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signInFrequencyAuthenticationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference) SignInFrequencyAuthenticationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signInFrequencyAuthenticationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference) SignInFrequencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"signInFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference) SignInFrequencyInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signInFrequencyInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference) SignInFrequencyIntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signInFrequencyIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference) SignInFrequencyPeriod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signInFrequencyPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference) SignInFrequencyPeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signInFrequencyPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewConditionalAccessPolicySessionControlsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ConditionalAccessPolicySessionControlsOutputReference {
	_init_.Initialize()

	if err := validateNewConditionalAccessPolicySessionControlsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azuread.conditionalAccessPolicy.ConditionalAccessPolicySessionControlsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewConditionalAccessPolicySessionControlsOutputReference_Override(c ConditionalAccessPolicySessionControlsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azuread.conditionalAccessPolicy.ConditionalAccessPolicySessionControlsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference)SetApplicationEnforcedRestrictionsEnabled(val interface{}) {
	if err := j.validateSetApplicationEnforcedRestrictionsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationEnforcedRestrictionsEnabled",
		val,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference)SetCloudAppSecurityPolicy(val *string) {
	if err := j.validateSetCloudAppSecurityPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudAppSecurityPolicy",
		val,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference)SetDisableResilienceDefaults(val interface{}) {
	if err := j.validateSetDisableResilienceDefaultsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableResilienceDefaults",
		val,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference)SetInternalValue(val *ConditionalAccessPolicySessionControls) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference)SetPersistentBrowserMode(val *string) {
	if err := j.validateSetPersistentBrowserModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"persistentBrowserMode",
		val,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference)SetSignInFrequency(val *float64) {
	if err := j.validateSetSignInFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signInFrequency",
		val,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference)SetSignInFrequencyAuthenticationType(val *string) {
	if err := j.validateSetSignInFrequencyAuthenticationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signInFrequencyAuthenticationType",
		val,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference)SetSignInFrequencyInterval(val *string) {
	if err := j.validateSetSignInFrequencyIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signInFrequencyInterval",
		val,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference)SetSignInFrequencyPeriod(val *string) {
	if err := j.validateSetSignInFrequencyPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signInFrequencyPeriod",
		val,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference) ResetApplicationEnforcedRestrictionsEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetApplicationEnforcedRestrictionsEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference) ResetCloudAppSecurityPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetCloudAppSecurityPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference) ResetDisableResilienceDefaults() {
	_jsii_.InvokeVoid(
		c,
		"resetDisableResilienceDefaults",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference) ResetPersistentBrowserMode() {
	_jsii_.InvokeVoid(
		c,
		"resetPersistentBrowserMode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference) ResetSignInFrequency() {
	_jsii_.InvokeVoid(
		c,
		"resetSignInFrequency",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference) ResetSignInFrequencyAuthenticationType() {
	_jsii_.InvokeVoid(
		c,
		"resetSignInFrequencyAuthenticationType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference) ResetSignInFrequencyInterval() {
	_jsii_.InvokeVoid(
		c,
		"resetSignInFrequencyInterval",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference) ResetSignInFrequencyPeriod() {
	_jsii_.InvokeVoid(
		c,
		"resetSignInFrequencyPeriod",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConditionalAccessPolicySessionControlsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


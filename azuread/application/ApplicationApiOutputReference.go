// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package application

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azuread-go/azuread/v10/jsii"

	"github.com/cdktf/cdktf-provider-azuread-go/azuread/v10/application/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApplicationApiOutputReference interface {
	cdktf.ComplexObject
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
	// Experimental.
	Fqn() *string
	InternalValue() *ApplicationApi
	SetInternalValue(val *ApplicationApi)
	KnownClientApplications() *[]*string
	SetKnownClientApplications(val *[]*string)
	KnownClientApplicationsInput() *[]*string
	MappedClaimsEnabled() interface{}
	SetMappedClaimsEnabled(val interface{})
	MappedClaimsEnabledInput() interface{}
	Oauth2PermissionScope() ApplicationApiOauth2PermissionScopeList
	Oauth2PermissionScopeInput() interface{}
	RequestedAccessTokenVersion() *float64
	SetRequestedAccessTokenVersion(val *float64)
	RequestedAccessTokenVersionInput() *float64
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
	PutOauth2PermissionScope(value interface{})
	ResetKnownClientApplications()
	ResetMappedClaimsEnabled()
	ResetOauth2PermissionScope()
	ResetRequestedAccessTokenVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApplicationApiOutputReference
type jsiiProxy_ApplicationApiOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApplicationApiOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationApiOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationApiOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationApiOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationApiOutputReference) InternalValue() *ApplicationApi {
	var returns *ApplicationApi
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationApiOutputReference) KnownClientApplications() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"knownClientApplications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationApiOutputReference) KnownClientApplicationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"knownClientApplicationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationApiOutputReference) MappedClaimsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mappedClaimsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationApiOutputReference) MappedClaimsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mappedClaimsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationApiOutputReference) Oauth2PermissionScope() ApplicationApiOauth2PermissionScopeList {
	var returns ApplicationApiOauth2PermissionScopeList
	_jsii_.Get(
		j,
		"oauth2PermissionScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationApiOutputReference) Oauth2PermissionScopeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oauth2PermissionScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationApiOutputReference) RequestedAccessTokenVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestedAccessTokenVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationApiOutputReference) RequestedAccessTokenVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestedAccessTokenVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationApiOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationApiOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApplicationApiOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApplicationApiOutputReference {
	_init_.Initialize()

	if err := validateNewApplicationApiOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationApiOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azuread.application.ApplicationApiOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApplicationApiOutputReference_Override(a ApplicationApiOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azuread.application.ApplicationApiOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApplicationApiOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApplicationApiOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApplicationApiOutputReference)SetInternalValue(val *ApplicationApi) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApplicationApiOutputReference)SetKnownClientApplications(val *[]*string) {
	if err := j.validateSetKnownClientApplicationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"knownClientApplications",
		val,
	)
}

func (j *jsiiProxy_ApplicationApiOutputReference)SetMappedClaimsEnabled(val interface{}) {
	if err := j.validateSetMappedClaimsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mappedClaimsEnabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationApiOutputReference)SetRequestedAccessTokenVersion(val *float64) {
	if err := j.validateSetRequestedAccessTokenVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestedAccessTokenVersion",
		val,
	)
}

func (j *jsiiProxy_ApplicationApiOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApplicationApiOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApplicationApiOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationApiOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApplicationApiOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApplicationApiOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApplicationApiOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApplicationApiOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApplicationApiOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApplicationApiOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApplicationApiOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApplicationApiOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApplicationApiOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationApiOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApplicationApiOutputReference) PutOauth2PermissionScope(value interface{}) {
	if err := a.validatePutOauth2PermissionScopeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOauth2PermissionScope",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationApiOutputReference) ResetKnownClientApplications() {
	_jsii_.InvokeVoid(
		a,
		"resetKnownClientApplications",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationApiOutputReference) ResetMappedClaimsEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetMappedClaimsEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationApiOutputReference) ResetOauth2PermissionScope() {
	_jsii_.InvokeVoid(
		a,
		"resetOauth2PermissionScope",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationApiOutputReference) ResetRequestedAccessTokenVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestedAccessTokenVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationApiOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ApplicationApiOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


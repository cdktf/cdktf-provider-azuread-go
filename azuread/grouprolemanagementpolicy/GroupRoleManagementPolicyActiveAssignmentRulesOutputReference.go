// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grouprolemanagementpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azuread-go/azuread/v13/jsii"

	"github.com/cdktf/cdktf-provider-azuread-go/azuread/v13/grouprolemanagementpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GroupRoleManagementPolicyActiveAssignmentRulesOutputReference interface {
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
	ExpirationRequired() interface{}
	SetExpirationRequired(val interface{})
	ExpirationRequiredInput() interface{}
	ExpireAfter() *string
	SetExpireAfter(val *string)
	ExpireAfterInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GroupRoleManagementPolicyActiveAssignmentRules
	SetInternalValue(val *GroupRoleManagementPolicyActiveAssignmentRules)
	RequireJustification() interface{}
	SetRequireJustification(val interface{})
	RequireJustificationInput() interface{}
	RequireMultifactorAuthentication() interface{}
	SetRequireMultifactorAuthentication(val interface{})
	RequireMultifactorAuthenticationInput() interface{}
	RequireTicketInfo() interface{}
	SetRequireTicketInfo(val interface{})
	RequireTicketInfoInput() interface{}
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
	ResetExpirationRequired()
	ResetExpireAfter()
	ResetRequireJustification()
	ResetRequireMultifactorAuthentication()
	ResetRequireTicketInfo()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GroupRoleManagementPolicyActiveAssignmentRulesOutputReference
type jsiiProxy_GroupRoleManagementPolicyActiveAssignmentRulesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GroupRoleManagementPolicyActiveAssignmentRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyActiveAssignmentRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyActiveAssignmentRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyActiveAssignmentRulesOutputReference) ExpirationRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"expirationRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyActiveAssignmentRulesOutputReference) ExpirationRequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"expirationRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyActiveAssignmentRulesOutputReference) ExpireAfter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expireAfter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyActiveAssignmentRulesOutputReference) ExpireAfterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expireAfterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyActiveAssignmentRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyActiveAssignmentRulesOutputReference) InternalValue() *GroupRoleManagementPolicyActiveAssignmentRules {
	var returns *GroupRoleManagementPolicyActiveAssignmentRules
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyActiveAssignmentRulesOutputReference) RequireJustification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireJustification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyActiveAssignmentRulesOutputReference) RequireJustificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireJustificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyActiveAssignmentRulesOutputReference) RequireMultifactorAuthentication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireMultifactorAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyActiveAssignmentRulesOutputReference) RequireMultifactorAuthenticationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireMultifactorAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyActiveAssignmentRulesOutputReference) RequireTicketInfo() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireTicketInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyActiveAssignmentRulesOutputReference) RequireTicketInfoInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireTicketInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyActiveAssignmentRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyActiveAssignmentRulesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGroupRoleManagementPolicyActiveAssignmentRulesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GroupRoleManagementPolicyActiveAssignmentRulesOutputReference {
	_init_.Initialize()

	if err := validateNewGroupRoleManagementPolicyActiveAssignmentRulesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GroupRoleManagementPolicyActiveAssignmentRulesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azuread.groupRoleManagementPolicy.GroupRoleManagementPolicyActiveAssignmentRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGroupRoleManagementPolicyActiveAssignmentRulesOutputReference_Override(g GroupRoleManagementPolicyActiveAssignmentRulesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azuread.groupRoleManagementPolicy.GroupRoleManagementPolicyActiveAssignmentRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GroupRoleManagementPolicyActiveAssignmentRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GroupRoleManagementPolicyActiveAssignmentRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GroupRoleManagementPolicyActiveAssignmentRulesOutputReference)SetExpirationRequired(val interface{}) {
	if err := j.validateSetExpirationRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expirationRequired",
		val,
	)
}

func (j *jsiiProxy_GroupRoleManagementPolicyActiveAssignmentRulesOutputReference)SetExpireAfter(val *string) {
	if err := j.validateSetExpireAfterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expireAfter",
		val,
	)
}

func (j *jsiiProxy_GroupRoleManagementPolicyActiveAssignmentRulesOutputReference)SetInternalValue(val *GroupRoleManagementPolicyActiveAssignmentRules) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GroupRoleManagementPolicyActiveAssignmentRulesOutputReference)SetRequireJustification(val interface{}) {
	if err := j.validateSetRequireJustificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireJustification",
		val,
	)
}

func (j *jsiiProxy_GroupRoleManagementPolicyActiveAssignmentRulesOutputReference)SetRequireMultifactorAuthentication(val interface{}) {
	if err := j.validateSetRequireMultifactorAuthenticationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireMultifactorAuthentication",
		val,
	)
}

func (j *jsiiProxy_GroupRoleManagementPolicyActiveAssignmentRulesOutputReference)SetRequireTicketInfo(val interface{}) {
	if err := j.validateSetRequireTicketInfoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireTicketInfo",
		val,
	)
}

func (j *jsiiProxy_GroupRoleManagementPolicyActiveAssignmentRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GroupRoleManagementPolicyActiveAssignmentRulesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GroupRoleManagementPolicyActiveAssignmentRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupRoleManagementPolicyActiveAssignmentRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupRoleManagementPolicyActiveAssignmentRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupRoleManagementPolicyActiveAssignmentRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupRoleManagementPolicyActiveAssignmentRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupRoleManagementPolicyActiveAssignmentRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupRoleManagementPolicyActiveAssignmentRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupRoleManagementPolicyActiveAssignmentRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupRoleManagementPolicyActiveAssignmentRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupRoleManagementPolicyActiveAssignmentRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupRoleManagementPolicyActiveAssignmentRulesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupRoleManagementPolicyActiveAssignmentRulesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupRoleManagementPolicyActiveAssignmentRulesOutputReference) ResetExpirationRequired() {
	_jsii_.InvokeVoid(
		g,
		"resetExpirationRequired",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupRoleManagementPolicyActiveAssignmentRulesOutputReference) ResetExpireAfter() {
	_jsii_.InvokeVoid(
		g,
		"resetExpireAfter",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupRoleManagementPolicyActiveAssignmentRulesOutputReference) ResetRequireJustification() {
	_jsii_.InvokeVoid(
		g,
		"resetRequireJustification",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupRoleManagementPolicyActiveAssignmentRulesOutputReference) ResetRequireMultifactorAuthentication() {
	_jsii_.InvokeVoid(
		g,
		"resetRequireMultifactorAuthentication",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupRoleManagementPolicyActiveAssignmentRulesOutputReference) ResetRequireTicketInfo() {
	_jsii_.InvokeVoid(
		g,
		"resetRequireTicketInfo",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupRoleManagementPolicyActiveAssignmentRulesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupRoleManagementPolicyActiveAssignmentRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


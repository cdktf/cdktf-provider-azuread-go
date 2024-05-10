// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grouprolemanagementpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azuread-go/azuread/v12/jsii"

	"github.com/cdktf/cdktf-provider-azuread-go/azuread/v12/grouprolemanagementpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GroupRoleManagementPolicyActivationRulesOutputReference interface {
	cdktf.ComplexObject
	ApprovalStage() GroupRoleManagementPolicyActivationRulesApprovalStageOutputReference
	ApprovalStageInput() *GroupRoleManagementPolicyActivationRulesApprovalStage
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
	InternalValue() *GroupRoleManagementPolicyActivationRules
	SetInternalValue(val *GroupRoleManagementPolicyActivationRules)
	MaximumDuration() *string
	SetMaximumDuration(val *string)
	MaximumDurationInput() *string
	RequireApproval() interface{}
	SetRequireApproval(val interface{})
	RequireApprovalInput() interface{}
	RequiredConditionalAccessAuthenticationContext() *string
	SetRequiredConditionalAccessAuthenticationContext(val *string)
	RequiredConditionalAccessAuthenticationContextInput() *string
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
	PutApprovalStage(value *GroupRoleManagementPolicyActivationRulesApprovalStage)
	ResetApprovalStage()
	ResetMaximumDuration()
	ResetRequireApproval()
	ResetRequiredConditionalAccessAuthenticationContext()
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

// The jsii proxy struct for GroupRoleManagementPolicyActivationRulesOutputReference
type jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference) ApprovalStage() GroupRoleManagementPolicyActivationRulesApprovalStageOutputReference {
	var returns GroupRoleManagementPolicyActivationRulesApprovalStageOutputReference
	_jsii_.Get(
		j,
		"approvalStage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference) ApprovalStageInput() *GroupRoleManagementPolicyActivationRulesApprovalStage {
	var returns *GroupRoleManagementPolicyActivationRulesApprovalStage
	_jsii_.Get(
		j,
		"approvalStageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference) InternalValue() *GroupRoleManagementPolicyActivationRules {
	var returns *GroupRoleManagementPolicyActivationRules
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference) MaximumDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maximumDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference) MaximumDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maximumDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference) RequireApproval() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireApproval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference) RequireApprovalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireApprovalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference) RequiredConditionalAccessAuthenticationContext() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requiredConditionalAccessAuthenticationContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference) RequiredConditionalAccessAuthenticationContextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requiredConditionalAccessAuthenticationContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference) RequireJustification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireJustification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference) RequireJustificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireJustificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference) RequireMultifactorAuthentication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireMultifactorAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference) RequireMultifactorAuthenticationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireMultifactorAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference) RequireTicketInfo() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireTicketInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference) RequireTicketInfoInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireTicketInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGroupRoleManagementPolicyActivationRulesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GroupRoleManagementPolicyActivationRulesOutputReference {
	_init_.Initialize()

	if err := validateNewGroupRoleManagementPolicyActivationRulesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azuread.groupRoleManagementPolicy.GroupRoleManagementPolicyActivationRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGroupRoleManagementPolicyActivationRulesOutputReference_Override(g GroupRoleManagementPolicyActivationRulesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azuread.groupRoleManagementPolicy.GroupRoleManagementPolicyActivationRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference)SetInternalValue(val *GroupRoleManagementPolicyActivationRules) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference)SetMaximumDuration(val *string) {
	if err := j.validateSetMaximumDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumDuration",
		val,
	)
}

func (j *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference)SetRequireApproval(val interface{}) {
	if err := j.validateSetRequireApprovalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireApproval",
		val,
	)
}

func (j *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference)SetRequiredConditionalAccessAuthenticationContext(val *string) {
	if err := j.validateSetRequiredConditionalAccessAuthenticationContextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requiredConditionalAccessAuthenticationContext",
		val,
	)
}

func (j *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference)SetRequireJustification(val interface{}) {
	if err := j.validateSetRequireJustificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireJustification",
		val,
	)
}

func (j *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference)SetRequireMultifactorAuthentication(val interface{}) {
	if err := j.validateSetRequireMultifactorAuthenticationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireMultifactorAuthentication",
		val,
	)
}

func (j *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference)SetRequireTicketInfo(val interface{}) {
	if err := j.validateSetRequireTicketInfoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireTicketInfo",
		val,
	)
}

func (j *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference) PutApprovalStage(value *GroupRoleManagementPolicyActivationRulesApprovalStage) {
	if err := g.validatePutApprovalStageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putApprovalStage",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference) ResetApprovalStage() {
	_jsii_.InvokeVoid(
		g,
		"resetApprovalStage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference) ResetMaximumDuration() {
	_jsii_.InvokeVoid(
		g,
		"resetMaximumDuration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference) ResetRequireApproval() {
	_jsii_.InvokeVoid(
		g,
		"resetRequireApproval",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference) ResetRequiredConditionalAccessAuthenticationContext() {
	_jsii_.InvokeVoid(
		g,
		"resetRequiredConditionalAccessAuthenticationContext",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference) ResetRequireJustification() {
	_jsii_.InvokeVoid(
		g,
		"resetRequireJustification",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference) ResetRequireMultifactorAuthentication() {
	_jsii_.InvokeVoid(
		g,
		"resetRequireMultifactorAuthentication",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference) ResetRequireTicketInfo() {
	_jsii_.InvokeVoid(
		g,
		"resetRequireTicketInfo",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GroupRoleManagementPolicyActivationRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


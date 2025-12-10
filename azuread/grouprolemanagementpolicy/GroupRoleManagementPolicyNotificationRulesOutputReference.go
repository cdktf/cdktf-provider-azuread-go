// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grouprolemanagementpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azuread-go/azuread/v14/jsii"

	"github.com/cdktf/cdktf-provider-azuread-go/azuread/v14/grouprolemanagementpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GroupRoleManagementPolicyNotificationRulesOutputReference interface {
	cdktf.ComplexObject
	ActiveAssignments() GroupRoleManagementPolicyNotificationRulesActiveAssignmentsOutputReference
	ActiveAssignmentsInput() *GroupRoleManagementPolicyNotificationRulesActiveAssignments
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
	EligibleActivations() GroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference
	EligibleActivationsInput() *GroupRoleManagementPolicyNotificationRulesEligibleActivations
	EligibleAssignments() GroupRoleManagementPolicyNotificationRulesEligibleAssignmentsOutputReference
	EligibleAssignmentsInput() *GroupRoleManagementPolicyNotificationRulesEligibleAssignments
	// Experimental.
	Fqn() *string
	InternalValue() *GroupRoleManagementPolicyNotificationRules
	SetInternalValue(val *GroupRoleManagementPolicyNotificationRules)
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutActiveAssignments(value *GroupRoleManagementPolicyNotificationRulesActiveAssignments)
	PutEligibleActivations(value *GroupRoleManagementPolicyNotificationRulesEligibleActivations)
	PutEligibleAssignments(value *GroupRoleManagementPolicyNotificationRulesEligibleAssignments)
	ResetActiveAssignments()
	ResetEligibleActivations()
	ResetEligibleAssignments()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GroupRoleManagementPolicyNotificationRulesOutputReference
type jsiiProxy_GroupRoleManagementPolicyNotificationRulesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesOutputReference) ActiveAssignments() GroupRoleManagementPolicyNotificationRulesActiveAssignmentsOutputReference {
	var returns GroupRoleManagementPolicyNotificationRulesActiveAssignmentsOutputReference
	_jsii_.Get(
		j,
		"activeAssignments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesOutputReference) ActiveAssignmentsInput() *GroupRoleManagementPolicyNotificationRulesActiveAssignments {
	var returns *GroupRoleManagementPolicyNotificationRulesActiveAssignments
	_jsii_.Get(
		j,
		"activeAssignmentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesOutputReference) EligibleActivations() GroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference {
	var returns GroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference
	_jsii_.Get(
		j,
		"eligibleActivations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesOutputReference) EligibleActivationsInput() *GroupRoleManagementPolicyNotificationRulesEligibleActivations {
	var returns *GroupRoleManagementPolicyNotificationRulesEligibleActivations
	_jsii_.Get(
		j,
		"eligibleActivationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesOutputReference) EligibleAssignments() GroupRoleManagementPolicyNotificationRulesEligibleAssignmentsOutputReference {
	var returns GroupRoleManagementPolicyNotificationRulesEligibleAssignmentsOutputReference
	_jsii_.Get(
		j,
		"eligibleAssignments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesOutputReference) EligibleAssignmentsInput() *GroupRoleManagementPolicyNotificationRulesEligibleAssignments {
	var returns *GroupRoleManagementPolicyNotificationRulesEligibleAssignments
	_jsii_.Get(
		j,
		"eligibleAssignmentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesOutputReference) InternalValue() *GroupRoleManagementPolicyNotificationRules {
	var returns *GroupRoleManagementPolicyNotificationRules
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGroupRoleManagementPolicyNotificationRulesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GroupRoleManagementPolicyNotificationRulesOutputReference {
	_init_.Initialize()

	if err := validateNewGroupRoleManagementPolicyNotificationRulesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GroupRoleManagementPolicyNotificationRulesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azuread.groupRoleManagementPolicy.GroupRoleManagementPolicyNotificationRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGroupRoleManagementPolicyNotificationRulesOutputReference_Override(g GroupRoleManagementPolicyNotificationRulesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azuread.groupRoleManagementPolicy.GroupRoleManagementPolicyNotificationRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesOutputReference)SetInternalValue(val *GroupRoleManagementPolicyNotificationRules) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesOutputReference) PutActiveAssignments(value *GroupRoleManagementPolicyNotificationRulesActiveAssignments) {
	if err := g.validatePutActiveAssignmentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putActiveAssignments",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesOutputReference) PutEligibleActivations(value *GroupRoleManagementPolicyNotificationRulesEligibleActivations) {
	if err := g.validatePutEligibleActivationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEligibleActivations",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesOutputReference) PutEligibleAssignments(value *GroupRoleManagementPolicyNotificationRulesEligibleAssignments) {
	if err := g.validatePutEligibleAssignmentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEligibleAssignments",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesOutputReference) ResetActiveAssignments() {
	_jsii_.InvokeVoid(
		g,
		"resetActiveAssignments",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesOutputReference) ResetEligibleActivations() {
	_jsii_.InvokeVoid(
		g,
		"resetEligibleActivations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesOutputReference) ResetEligibleAssignments() {
	_jsii_.InvokeVoid(
		g,
		"resetEligibleAssignments",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


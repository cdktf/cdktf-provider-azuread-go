// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package conditionalaccesspolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azuread-go/azuread/v11/jsii"

	"github.com/cdktf/cdktf-provider-azuread-go/azuread/v11/conditionalaccesspolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUsersExternalTenantsList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUsersExternalTenantsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUsersExternalTenantsList
type jsiiProxy_ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUsersExternalTenantsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUsersExternalTenantsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUsersExternalTenantsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUsersExternalTenantsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUsersExternalTenantsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUsersExternalTenantsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUsersExternalTenantsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUsersExternalTenantsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUsersExternalTenantsList {
	_init_.Initialize()

	if err := validateNewConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUsersExternalTenantsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUsersExternalTenantsList{}

	_jsii_.Create(
		"@cdktf/provider-azuread.conditionalAccessPolicy.ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUsersExternalTenantsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUsersExternalTenantsList_Override(c ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUsersExternalTenantsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azuread.conditionalAccessPolicy.ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUsersExternalTenantsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		c,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUsersExternalTenantsList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUsersExternalTenantsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUsersExternalTenantsList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUsersExternalTenantsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (c *jsiiProxy_ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUsersExternalTenantsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUsersExternalTenantsList) Get(index *float64) ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUsersExternalTenantsOutputReference {
	if err := c.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUsersExternalTenantsOutputReference

	_jsii_.Invoke(
		c,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUsersExternalTenantsList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUsersExternalTenantsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

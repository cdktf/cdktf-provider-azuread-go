// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package conditionalaccesspolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azuread-go/azuread/v13/jsii"

	"github.com/cdktf/cdktf-provider-azuread-go/azuread/v13/conditionalaccesspolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Members() *[]*string
	SetMembers(val *[]*string)
	MembershipKind() *string
	SetMembershipKind(val *string)
	MembershipKindInput() *string
	MembersInput() *[]*string
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
	ResetMembers()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsOutputReference
type jsiiProxy_ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsOutputReference) Members() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"members",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsOutputReference) MembershipKind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"membershipKind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsOutputReference) MembershipKindInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"membershipKindInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsOutputReference) MembersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"membersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsOutputReference {
	_init_.Initialize()

	if err := validateNewConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azuread.conditionalAccessPolicy.ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsOutputReference_Override(c ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azuread.conditionalAccessPolicy.ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsOutputReference)SetMembers(val *[]*string) {
	if err := j.validateSetMembersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"members",
		val,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsOutputReference)SetMembershipKind(val *string) {
	if err := j.validateSetMembershipKindParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"membershipKind",
		val,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsOutputReference) ResetMembers() {
	_jsii_.InvokeVoid(
		c,
		"resetMembers",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ConditionalAccessPolicyConditionsUsersIncludedGuestsOrExternalUsersExternalTenantsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


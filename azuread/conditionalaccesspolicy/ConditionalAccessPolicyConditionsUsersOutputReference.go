package conditionalaccesspolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azuread-go/azuread/v6/jsii"

	"github.com/cdktf/cdktf-provider-azuread-go/azuread/v6/conditionalaccesspolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConditionalAccessPolicyConditionsUsersOutputReference interface {
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
	ExcludedGroups() *[]*string
	SetExcludedGroups(val *[]*string)
	ExcludedGroupsInput() *[]*string
	ExcludedRoles() *[]*string
	SetExcludedRoles(val *[]*string)
	ExcludedRolesInput() *[]*string
	ExcludedUsers() *[]*string
	SetExcludedUsers(val *[]*string)
	ExcludedUsersInput() *[]*string
	// Experimental.
	Fqn() *string
	IncludedGroups() *[]*string
	SetIncludedGroups(val *[]*string)
	IncludedGroupsInput() *[]*string
	IncludedRoles() *[]*string
	SetIncludedRoles(val *[]*string)
	IncludedRolesInput() *[]*string
	IncludedUsers() *[]*string
	SetIncludedUsers(val *[]*string)
	IncludedUsersInput() *[]*string
	InternalValue() *ConditionalAccessPolicyConditionsUsers
	SetInternalValue(val *ConditionalAccessPolicyConditionsUsers)
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
	ResetExcludedGroups()
	ResetExcludedRoles()
	ResetExcludedUsers()
	ResetIncludedGroups()
	ResetIncludedRoles()
	ResetIncludedUsers()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConditionalAccessPolicyConditionsUsersOutputReference
type jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference) ExcludedGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference) ExcludedGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference) ExcludedRoles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference) ExcludedRolesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedRolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference) ExcludedUsers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedUsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference) ExcludedUsersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedUsersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference) IncludedGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference) IncludedGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference) IncludedRoles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference) IncludedRolesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedRolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference) IncludedUsers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedUsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference) IncludedUsersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedUsersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference) InternalValue() *ConditionalAccessPolicyConditionsUsers {
	var returns *ConditionalAccessPolicyConditionsUsers
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewConditionalAccessPolicyConditionsUsersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ConditionalAccessPolicyConditionsUsersOutputReference {
	_init_.Initialize()

	if err := validateNewConditionalAccessPolicyConditionsUsersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azuread.conditionalAccessPolicy.ConditionalAccessPolicyConditionsUsersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewConditionalAccessPolicyConditionsUsersOutputReference_Override(c ConditionalAccessPolicyConditionsUsersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azuread.conditionalAccessPolicy.ConditionalAccessPolicyConditionsUsersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference)SetExcludedGroups(val *[]*string) {
	if err := j.validateSetExcludedGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedGroups",
		val,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference)SetExcludedRoles(val *[]*string) {
	if err := j.validateSetExcludedRolesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedRoles",
		val,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference)SetExcludedUsers(val *[]*string) {
	if err := j.validateSetExcludedUsersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedUsers",
		val,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference)SetIncludedGroups(val *[]*string) {
	if err := j.validateSetIncludedGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includedGroups",
		val,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference)SetIncludedRoles(val *[]*string) {
	if err := j.validateSetIncludedRolesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includedRoles",
		val,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference)SetIncludedUsers(val *[]*string) {
	if err := j.validateSetIncludedUsersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includedUsers",
		val,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference)SetInternalValue(val *ConditionalAccessPolicyConditionsUsers) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference) ResetExcludedGroups() {
	_jsii_.InvokeVoid(
		c,
		"resetExcludedGroups",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference) ResetExcludedRoles() {
	_jsii_.InvokeVoid(
		c,
		"resetExcludedRoles",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference) ResetExcludedUsers() {
	_jsii_.InvokeVoid(
		c,
		"resetExcludedUsers",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference) ResetIncludedGroups() {
	_jsii_.InvokeVoid(
		c,
		"resetIncludedGroups",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference) ResetIncludedRoles() {
	_jsii_.InvokeVoid(
		c,
		"resetIncludedRoles",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference) ResetIncludedUsers() {
	_jsii_.InvokeVoid(
		c,
		"resetIncludedUsers",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ConditionalAccessPolicyConditionsUsersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


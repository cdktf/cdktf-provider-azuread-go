package conditionalaccesspolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azuread-go/azuread/v3/jsii"

	"github.com/cdktf/cdktf-provider-azuread-go/azuread/v3/conditionalaccesspolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConditionalAccessPolicyConditionsLocationsOutputReference interface {
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
	ExcludedLocations() *[]*string
	SetExcludedLocations(val *[]*string)
	ExcludedLocationsInput() *[]*string
	// Experimental.
	Fqn() *string
	IncludedLocations() *[]*string
	SetIncludedLocations(val *[]*string)
	IncludedLocationsInput() *[]*string
	InternalValue() *ConditionalAccessPolicyConditionsLocations
	SetInternalValue(val *ConditionalAccessPolicyConditionsLocations)
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
	ResetExcludedLocations()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConditionalAccessPolicyConditionsLocationsOutputReference
type jsiiProxy_ConditionalAccessPolicyConditionsLocationsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsLocationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsLocationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsLocationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsLocationsOutputReference) ExcludedLocations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedLocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsLocationsOutputReference) ExcludedLocationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedLocationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsLocationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsLocationsOutputReference) IncludedLocations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedLocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsLocationsOutputReference) IncludedLocationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedLocationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsLocationsOutputReference) InternalValue() *ConditionalAccessPolicyConditionsLocations {
	var returns *ConditionalAccessPolicyConditionsLocations
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsLocationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsLocationsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewConditionalAccessPolicyConditionsLocationsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ConditionalAccessPolicyConditionsLocationsOutputReference {
	_init_.Initialize()

	if err := validateNewConditionalAccessPolicyConditionsLocationsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConditionalAccessPolicyConditionsLocationsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azuread.conditionalAccessPolicy.ConditionalAccessPolicyConditionsLocationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewConditionalAccessPolicyConditionsLocationsOutputReference_Override(c ConditionalAccessPolicyConditionsLocationsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azuread.conditionalAccessPolicy.ConditionalAccessPolicyConditionsLocationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsLocationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsLocationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsLocationsOutputReference)SetExcludedLocations(val *[]*string) {
	if err := j.validateSetExcludedLocationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedLocations",
		val,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsLocationsOutputReference)SetIncludedLocations(val *[]*string) {
	if err := j.validateSetIncludedLocationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includedLocations",
		val,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsLocationsOutputReference)SetInternalValue(val *ConditionalAccessPolicyConditionsLocations) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsLocationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsLocationsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ConditionalAccessPolicyConditionsLocationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConditionalAccessPolicyConditionsLocationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ConditionalAccessPolicyConditionsLocationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ConditionalAccessPolicyConditionsLocationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ConditionalAccessPolicyConditionsLocationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ConditionalAccessPolicyConditionsLocationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ConditionalAccessPolicyConditionsLocationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ConditionalAccessPolicyConditionsLocationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ConditionalAccessPolicyConditionsLocationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ConditionalAccessPolicyConditionsLocationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ConditionalAccessPolicyConditionsLocationsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConditionalAccessPolicyConditionsLocationsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ConditionalAccessPolicyConditionsLocationsOutputReference) ResetExcludedLocations() {
	_jsii_.InvokeVoid(
		c,
		"resetExcludedLocations",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConditionalAccessPolicyConditionsLocationsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ConditionalAccessPolicyConditionsLocationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


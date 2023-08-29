// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package conditionalaccesspolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azuread-go/azuread/v10/jsii"

	"github.com/cdktf/cdktf-provider-azuread-go/azuread/v10/conditionalaccesspolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConditionalAccessPolicyConditionsOutputReference interface {
	cdktf.ComplexObject
	Applications() ConditionalAccessPolicyConditionsApplicationsOutputReference
	ApplicationsInput() *ConditionalAccessPolicyConditionsApplications
	ClientApplications() ConditionalAccessPolicyConditionsClientApplicationsOutputReference
	ClientApplicationsInput() *ConditionalAccessPolicyConditionsClientApplications
	ClientAppTypes() *[]*string
	SetClientAppTypes(val *[]*string)
	ClientAppTypesInput() *[]*string
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
	Devices() ConditionalAccessPolicyConditionsDevicesOutputReference
	DevicesInput() *ConditionalAccessPolicyConditionsDevices
	// Experimental.
	Fqn() *string
	InternalValue() *ConditionalAccessPolicyConditions
	SetInternalValue(val *ConditionalAccessPolicyConditions)
	Locations() ConditionalAccessPolicyConditionsLocationsOutputReference
	LocationsInput() *ConditionalAccessPolicyConditionsLocations
	Platforms() ConditionalAccessPolicyConditionsPlatformsOutputReference
	PlatformsInput() *ConditionalAccessPolicyConditionsPlatforms
	ServicePrincipalRiskLevels() *[]*string
	SetServicePrincipalRiskLevels(val *[]*string)
	ServicePrincipalRiskLevelsInput() *[]*string
	SignInRiskLevels() *[]*string
	SetSignInRiskLevels(val *[]*string)
	SignInRiskLevelsInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserRiskLevels() *[]*string
	SetUserRiskLevels(val *[]*string)
	UserRiskLevelsInput() *[]*string
	Users() ConditionalAccessPolicyConditionsUsersOutputReference
	UsersInput() *ConditionalAccessPolicyConditionsUsers
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
	PutApplications(value *ConditionalAccessPolicyConditionsApplications)
	PutClientApplications(value *ConditionalAccessPolicyConditionsClientApplications)
	PutDevices(value *ConditionalAccessPolicyConditionsDevices)
	PutLocations(value *ConditionalAccessPolicyConditionsLocations)
	PutPlatforms(value *ConditionalAccessPolicyConditionsPlatforms)
	PutUsers(value *ConditionalAccessPolicyConditionsUsers)
	ResetClientApplications()
	ResetDevices()
	ResetLocations()
	ResetPlatforms()
	ResetServicePrincipalRiskLevels()
	ResetSignInRiskLevels()
	ResetUserRiskLevels()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConditionalAccessPolicyConditionsOutputReference
type jsiiProxy_ConditionalAccessPolicyConditionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) Applications() ConditionalAccessPolicyConditionsApplicationsOutputReference {
	var returns ConditionalAccessPolicyConditionsApplicationsOutputReference
	_jsii_.Get(
		j,
		"applications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) ApplicationsInput() *ConditionalAccessPolicyConditionsApplications {
	var returns *ConditionalAccessPolicyConditionsApplications
	_jsii_.Get(
		j,
		"applicationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) ClientApplications() ConditionalAccessPolicyConditionsClientApplicationsOutputReference {
	var returns ConditionalAccessPolicyConditionsClientApplicationsOutputReference
	_jsii_.Get(
		j,
		"clientApplications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) ClientApplicationsInput() *ConditionalAccessPolicyConditionsClientApplications {
	var returns *ConditionalAccessPolicyConditionsClientApplications
	_jsii_.Get(
		j,
		"clientApplicationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) ClientAppTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clientAppTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) ClientAppTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clientAppTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) Devices() ConditionalAccessPolicyConditionsDevicesOutputReference {
	var returns ConditionalAccessPolicyConditionsDevicesOutputReference
	_jsii_.Get(
		j,
		"devices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) DevicesInput() *ConditionalAccessPolicyConditionsDevices {
	var returns *ConditionalAccessPolicyConditionsDevices
	_jsii_.Get(
		j,
		"devicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) InternalValue() *ConditionalAccessPolicyConditions {
	var returns *ConditionalAccessPolicyConditions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) Locations() ConditionalAccessPolicyConditionsLocationsOutputReference {
	var returns ConditionalAccessPolicyConditionsLocationsOutputReference
	_jsii_.Get(
		j,
		"locations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) LocationsInput() *ConditionalAccessPolicyConditionsLocations {
	var returns *ConditionalAccessPolicyConditionsLocations
	_jsii_.Get(
		j,
		"locationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) Platforms() ConditionalAccessPolicyConditionsPlatformsOutputReference {
	var returns ConditionalAccessPolicyConditionsPlatformsOutputReference
	_jsii_.Get(
		j,
		"platforms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) PlatformsInput() *ConditionalAccessPolicyConditionsPlatforms {
	var returns *ConditionalAccessPolicyConditionsPlatforms
	_jsii_.Get(
		j,
		"platformsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) ServicePrincipalRiskLevels() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"servicePrincipalRiskLevels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) ServicePrincipalRiskLevelsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"servicePrincipalRiskLevelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) SignInRiskLevels() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"signInRiskLevels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) SignInRiskLevelsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"signInRiskLevelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) UserRiskLevels() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"userRiskLevels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) UserRiskLevelsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"userRiskLevelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) Users() ConditionalAccessPolicyConditionsUsersOutputReference {
	var returns ConditionalAccessPolicyConditionsUsersOutputReference
	_jsii_.Get(
		j,
		"users",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) UsersInput() *ConditionalAccessPolicyConditionsUsers {
	var returns *ConditionalAccessPolicyConditionsUsers
	_jsii_.Get(
		j,
		"usersInput",
		&returns,
	)
	return returns
}


func NewConditionalAccessPolicyConditionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ConditionalAccessPolicyConditionsOutputReference {
	_init_.Initialize()

	if err := validateNewConditionalAccessPolicyConditionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConditionalAccessPolicyConditionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azuread.conditionalAccessPolicy.ConditionalAccessPolicyConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewConditionalAccessPolicyConditionsOutputReference_Override(c ConditionalAccessPolicyConditionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azuread.conditionalAccessPolicy.ConditionalAccessPolicyConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference)SetClientAppTypes(val *[]*string) {
	if err := j.validateSetClientAppTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientAppTypes",
		val,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference)SetInternalValue(val *ConditionalAccessPolicyConditions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference)SetServicePrincipalRiskLevels(val *[]*string) {
	if err := j.validateSetServicePrincipalRiskLevelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"servicePrincipalRiskLevels",
		val,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference)SetSignInRiskLevels(val *[]*string) {
	if err := j.validateSetSignInRiskLevelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signInRiskLevels",
		val,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference)SetUserRiskLevels(val *[]*string) {
	if err := j.validateSetUserRiskLevelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userRiskLevels",
		val,
	)
}

func (c *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) PutApplications(value *ConditionalAccessPolicyConditionsApplications) {
	if err := c.validatePutApplicationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putApplications",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) PutClientApplications(value *ConditionalAccessPolicyConditionsClientApplications) {
	if err := c.validatePutClientApplicationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putClientApplications",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) PutDevices(value *ConditionalAccessPolicyConditionsDevices) {
	if err := c.validatePutDevicesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDevices",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) PutLocations(value *ConditionalAccessPolicyConditionsLocations) {
	if err := c.validatePutLocationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLocations",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) PutPlatforms(value *ConditionalAccessPolicyConditionsPlatforms) {
	if err := c.validatePutPlatformsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPlatforms",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) PutUsers(value *ConditionalAccessPolicyConditionsUsers) {
	if err := c.validatePutUsersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putUsers",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) ResetClientApplications() {
	_jsii_.InvokeVoid(
		c,
		"resetClientApplications",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) ResetDevices() {
	_jsii_.InvokeVoid(
		c,
		"resetDevices",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) ResetLocations() {
	_jsii_.InvokeVoid(
		c,
		"resetLocations",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) ResetPlatforms() {
	_jsii_.InvokeVoid(
		c,
		"resetPlatforms",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) ResetServicePrincipalRiskLevels() {
	_jsii_.InvokeVoid(
		c,
		"resetServicePrincipalRiskLevels",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) ResetSignInRiskLevels() {
	_jsii_.InvokeVoid(
		c,
		"resetSignInRiskLevels",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) ResetUserRiskLevels() {
	_jsii_.InvokeVoid(
		c,
		"resetUserRiskLevels",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ConditionalAccessPolicyConditionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


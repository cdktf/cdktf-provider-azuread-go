package application

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azuread-go/azuread/v7/jsii"

	"github.com/cdktf/cdktf-provider-azuread-go/azuread/v7/application/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApplicationApiOauth2PermissionScopeOutputReference interface {
	cdktf.ComplexObject
	AdminConsentDescription() *string
	SetAdminConsentDescription(val *string)
	AdminConsentDescriptionInput() *string
	AdminConsentDisplayName() *string
	SetAdminConsentDisplayName(val *string)
	AdminConsentDisplayNameInput() *string
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
	UserConsentDescription() *string
	SetUserConsentDescription(val *string)
	UserConsentDescriptionInput() *string
	UserConsentDisplayName() *string
	SetUserConsentDisplayName(val *string)
	UserConsentDisplayNameInput() *string
	Value() *string
	SetValue(val *string)
	ValueInput() *string
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
	ResetAdminConsentDescription()
	ResetAdminConsentDisplayName()
	ResetEnabled()
	ResetType()
	ResetUserConsentDescription()
	ResetUserConsentDisplayName()
	ResetValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApplicationApiOauth2PermissionScopeOutputReference
type jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference) AdminConsentDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminConsentDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference) AdminConsentDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminConsentDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference) AdminConsentDisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminConsentDisplayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference) AdminConsentDisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminConsentDisplayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference) UserConsentDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userConsentDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference) UserConsentDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userConsentDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference) UserConsentDisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userConsentDisplayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference) UserConsentDisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userConsentDisplayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference) ValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewApplicationApiOauth2PermissionScopeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ApplicationApiOauth2PermissionScopeOutputReference {
	_init_.Initialize()

	if err := validateNewApplicationApiOauth2PermissionScopeOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azuread.application.ApplicationApiOauth2PermissionScopeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewApplicationApiOauth2PermissionScopeOutputReference_Override(a ApplicationApiOauth2PermissionScopeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azuread.application.ApplicationApiOauth2PermissionScopeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference)SetAdminConsentDescription(val *string) {
	if err := j.validateSetAdminConsentDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminConsentDescription",
		val,
	)
}

func (j *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference)SetAdminConsentDisplayName(val *string) {
	if err := j.validateSetAdminConsentDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminConsentDisplayName",
		val,
	)
}

func (j *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference)SetUserConsentDescription(val *string) {
	if err := j.validateSetUserConsentDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userConsentDescription",
		val,
	)
}

func (j *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference)SetUserConsentDisplayName(val *string) {
	if err := j.validateSetUserConsentDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userConsentDisplayName",
		val,
	)
}

func (j *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference)SetValue(val *string) {
	if err := j.validateSetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (a *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference) ResetAdminConsentDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetAdminConsentDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference) ResetAdminConsentDisplayName() {
	_jsii_.InvokeVoid(
		a,
		"resetAdminConsentDisplayName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		a,
		"resetType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference) ResetUserConsentDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetUserConsentDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference) ResetUserConsentDisplayName() {
	_jsii_.InvokeVoid(
		a,
		"resetUserConsentDisplayName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference) ResetValue() {
	_jsii_.InvokeVoid(
		a,
		"resetValue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ApplicationApiOauth2PermissionScopeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


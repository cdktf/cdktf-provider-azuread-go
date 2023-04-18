package application

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azuread-go/azuread/v7/jsii"

	"github.com/cdktf/cdktf-provider-azuread-go/azuread/v7/application/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApplicationFeatureTagsOutputReference interface {
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
	CustomSingleSignOn() interface{}
	SetCustomSingleSignOn(val interface{})
	CustomSingleSignOnInput() interface{}
	Enterprise() interface{}
	SetEnterprise(val interface{})
	EnterpriseInput() interface{}
	// Experimental.
	Fqn() *string
	Gallery() interface{}
	SetGallery(val interface{})
	GalleryInput() interface{}
	Hide() interface{}
	SetHide(val interface{})
	HideInput() interface{}
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
	ResetCustomSingleSignOn()
	ResetEnterprise()
	ResetGallery()
	ResetHide()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApplicationFeatureTagsOutputReference
type jsiiProxy_ApplicationFeatureTagsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApplicationFeatureTagsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationFeatureTagsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationFeatureTagsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationFeatureTagsOutputReference) CustomSingleSignOn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customSingleSignOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationFeatureTagsOutputReference) CustomSingleSignOnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customSingleSignOnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationFeatureTagsOutputReference) Enterprise() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enterprise",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationFeatureTagsOutputReference) EnterpriseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enterpriseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationFeatureTagsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationFeatureTagsOutputReference) Gallery() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gallery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationFeatureTagsOutputReference) GalleryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"galleryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationFeatureTagsOutputReference) Hide() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hide",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationFeatureTagsOutputReference) HideInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationFeatureTagsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationFeatureTagsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationFeatureTagsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApplicationFeatureTagsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ApplicationFeatureTagsOutputReference {
	_init_.Initialize()

	if err := validateNewApplicationFeatureTagsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationFeatureTagsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azuread.application.ApplicationFeatureTagsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewApplicationFeatureTagsOutputReference_Override(a ApplicationFeatureTagsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azuread.application.ApplicationFeatureTagsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_ApplicationFeatureTagsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApplicationFeatureTagsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApplicationFeatureTagsOutputReference)SetCustomSingleSignOn(val interface{}) {
	if err := j.validateSetCustomSingleSignOnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customSingleSignOn",
		val,
	)
}

func (j *jsiiProxy_ApplicationFeatureTagsOutputReference)SetEnterprise(val interface{}) {
	if err := j.validateSetEnterpriseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enterprise",
		val,
	)
}

func (j *jsiiProxy_ApplicationFeatureTagsOutputReference)SetGallery(val interface{}) {
	if err := j.validateSetGalleryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gallery",
		val,
	)
}

func (j *jsiiProxy_ApplicationFeatureTagsOutputReference)SetHide(val interface{}) {
	if err := j.validateSetHideParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hide",
		val,
	)
}

func (j *jsiiProxy_ApplicationFeatureTagsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApplicationFeatureTagsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApplicationFeatureTagsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApplicationFeatureTagsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationFeatureTagsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApplicationFeatureTagsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApplicationFeatureTagsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApplicationFeatureTagsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApplicationFeatureTagsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApplicationFeatureTagsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApplicationFeatureTagsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApplicationFeatureTagsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApplicationFeatureTagsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApplicationFeatureTagsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationFeatureTagsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApplicationFeatureTagsOutputReference) ResetCustomSingleSignOn() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomSingleSignOn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationFeatureTagsOutputReference) ResetEnterprise() {
	_jsii_.InvokeVoid(
		a,
		"resetEnterprise",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationFeatureTagsOutputReference) ResetGallery() {
	_jsii_.InvokeVoid(
		a,
		"resetGallery",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationFeatureTagsOutputReference) ResetHide() {
	_jsii_.InvokeVoid(
		a,
		"resetHide",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationFeatureTagsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ApplicationFeatureTagsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


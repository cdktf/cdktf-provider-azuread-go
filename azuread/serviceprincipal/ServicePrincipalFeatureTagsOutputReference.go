package serviceprincipal

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azuread-go/azuread/v7/jsii"

	"github.com/cdktf/cdktf-provider-azuread-go/azuread/v7/serviceprincipal/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServicePrincipalFeatureTagsOutputReference interface {
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

// The jsii proxy struct for ServicePrincipalFeatureTagsOutputReference
type jsiiProxy_ServicePrincipalFeatureTagsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServicePrincipalFeatureTagsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipalFeatureTagsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipalFeatureTagsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipalFeatureTagsOutputReference) CustomSingleSignOn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customSingleSignOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipalFeatureTagsOutputReference) CustomSingleSignOnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customSingleSignOnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipalFeatureTagsOutputReference) Enterprise() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enterprise",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipalFeatureTagsOutputReference) EnterpriseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enterpriseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipalFeatureTagsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipalFeatureTagsOutputReference) Gallery() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gallery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipalFeatureTagsOutputReference) GalleryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"galleryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipalFeatureTagsOutputReference) Hide() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hide",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipalFeatureTagsOutputReference) HideInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipalFeatureTagsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipalFeatureTagsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipalFeatureTagsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewServicePrincipalFeatureTagsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ServicePrincipalFeatureTagsOutputReference {
	_init_.Initialize()

	if err := validateNewServicePrincipalFeatureTagsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServicePrincipalFeatureTagsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azuread.servicePrincipal.ServicePrincipalFeatureTagsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewServicePrincipalFeatureTagsOutputReference_Override(s ServicePrincipalFeatureTagsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azuread.servicePrincipal.ServicePrincipalFeatureTagsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_ServicePrincipalFeatureTagsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipalFeatureTagsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipalFeatureTagsOutputReference)SetCustomSingleSignOn(val interface{}) {
	if err := j.validateSetCustomSingleSignOnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customSingleSignOn",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipalFeatureTagsOutputReference)SetEnterprise(val interface{}) {
	if err := j.validateSetEnterpriseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enterprise",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipalFeatureTagsOutputReference)SetGallery(val interface{}) {
	if err := j.validateSetGalleryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gallery",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipalFeatureTagsOutputReference)SetHide(val interface{}) {
	if err := j.validateSetHideParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hide",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipalFeatureTagsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipalFeatureTagsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipalFeatureTagsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_ServicePrincipalFeatureTagsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipalFeatureTagsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipalFeatureTagsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipalFeatureTagsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipalFeatureTagsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipalFeatureTagsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipalFeatureTagsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipalFeatureTagsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipalFeatureTagsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipalFeatureTagsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipalFeatureTagsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipalFeatureTagsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipalFeatureTagsOutputReference) ResetCustomSingleSignOn() {
	_jsii_.InvokeVoid(
		s,
		"resetCustomSingleSignOn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicePrincipalFeatureTagsOutputReference) ResetEnterprise() {
	_jsii_.InvokeVoid(
		s,
		"resetEnterprise",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicePrincipalFeatureTagsOutputReference) ResetGallery() {
	_jsii_.InvokeVoid(
		s,
		"resetGallery",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicePrincipalFeatureTagsOutputReference) ResetHide() {
	_jsii_.InvokeVoid(
		s,
		"resetHide",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicePrincipalFeatureTagsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipalFeatureTagsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


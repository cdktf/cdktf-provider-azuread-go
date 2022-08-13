// Prebuilt azuread Provider for Terraform CDK (cdktf)
package azuread

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azuread-go/azuread/v2/jsii"

	"github.com/hashicorp/cdktf-provider-azuread-go/azuread/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApplicationWebOutputReference interface {
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
	HomepageUrl() *string
	SetHomepageUrl(val *string)
	HomepageUrlInput() *string
	ImplicitGrant() ApplicationWebImplicitGrantOutputReference
	ImplicitGrantInput() *ApplicationWebImplicitGrant
	InternalValue() *ApplicationWeb
	SetInternalValue(val *ApplicationWeb)
	LogoutUrl() *string
	SetLogoutUrl(val *string)
	LogoutUrlInput() *string
	RedirectUris() *[]*string
	SetRedirectUris(val *[]*string)
	RedirectUrisInput() *[]*string
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
	PutImplicitGrant(value *ApplicationWebImplicitGrant)
	ResetHomepageUrl()
	ResetImplicitGrant()
	ResetLogoutUrl()
	ResetRedirectUris()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApplicationWebOutputReference
type jsiiProxy_ApplicationWebOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApplicationWebOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationWebOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationWebOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationWebOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationWebOutputReference) HomepageUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homepageUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationWebOutputReference) HomepageUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homepageUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationWebOutputReference) ImplicitGrant() ApplicationWebImplicitGrantOutputReference {
	var returns ApplicationWebImplicitGrantOutputReference
	_jsii_.Get(
		j,
		"implicitGrant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationWebOutputReference) ImplicitGrantInput() *ApplicationWebImplicitGrant {
	var returns *ApplicationWebImplicitGrant
	_jsii_.Get(
		j,
		"implicitGrantInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationWebOutputReference) InternalValue() *ApplicationWeb {
	var returns *ApplicationWeb
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationWebOutputReference) LogoutUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoutUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationWebOutputReference) LogoutUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoutUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationWebOutputReference) RedirectUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"redirectUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationWebOutputReference) RedirectUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"redirectUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationWebOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationWebOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApplicationWebOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApplicationWebOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ApplicationWebOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azuread.ApplicationWebOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApplicationWebOutputReference_Override(a ApplicationWebOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azuread.ApplicationWebOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApplicationWebOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApplicationWebOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApplicationWebOutputReference) SetHomepageUrl(val *string) {
	_jsii_.Set(
		j,
		"homepageUrl",
		val,
	)
}

func (j *jsiiProxy_ApplicationWebOutputReference) SetInternalValue(val *ApplicationWeb) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApplicationWebOutputReference) SetLogoutUrl(val *string) {
	_jsii_.Set(
		j,
		"logoutUrl",
		val,
	)
}

func (j *jsiiProxy_ApplicationWebOutputReference) SetRedirectUris(val *[]*string) {
	_jsii_.Set(
		j,
		"redirectUris",
		val,
	)
}

func (j *jsiiProxy_ApplicationWebOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApplicationWebOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApplicationWebOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationWebOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationWebOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationWebOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationWebOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationWebOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationWebOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationWebOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationWebOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationWebOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationWebOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationWebOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationWebOutputReference) PutImplicitGrant(value *ApplicationWebImplicitGrant) {
	_jsii_.InvokeVoid(
		a,
		"putImplicitGrant",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationWebOutputReference) ResetHomepageUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetHomepageUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationWebOutputReference) ResetImplicitGrant() {
	_jsii_.InvokeVoid(
		a,
		"resetImplicitGrant",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationWebOutputReference) ResetLogoutUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetLogoutUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationWebOutputReference) ResetRedirectUris() {
	_jsii_.InvokeVoid(
		a,
		"resetRedirectUris",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationWebOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationWebOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


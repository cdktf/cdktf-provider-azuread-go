package application

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azuread-go/azuread/v9/jsii"

	"github.com/cdktf/cdktf-provider-azuread-go/azuread/v9/application/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApplicationOptionalClaimsOutputReference interface {
	cdktf.ComplexObject
	AccessToken() ApplicationOptionalClaimsAccessTokenList
	AccessTokenInput() interface{}
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
	IdToken() ApplicationOptionalClaimsIdTokenList
	IdTokenInput() interface{}
	InternalValue() *ApplicationOptionalClaims
	SetInternalValue(val *ApplicationOptionalClaims)
	Saml2Token() ApplicationOptionalClaimsSaml2TokenList
	Saml2TokenInput() interface{}
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
	PutAccessToken(value interface{})
	PutIdToken(value interface{})
	PutSaml2Token(value interface{})
	ResetAccessToken()
	ResetIdToken()
	ResetSaml2Token()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApplicationOptionalClaimsOutputReference
type jsiiProxy_ApplicationOptionalClaimsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApplicationOptionalClaimsOutputReference) AccessToken() ApplicationOptionalClaimsAccessTokenList {
	var returns ApplicationOptionalClaimsAccessTokenList
	_jsii_.Get(
		j,
		"accessToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationOptionalClaimsOutputReference) AccessTokenInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationOptionalClaimsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationOptionalClaimsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationOptionalClaimsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationOptionalClaimsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationOptionalClaimsOutputReference) IdToken() ApplicationOptionalClaimsIdTokenList {
	var returns ApplicationOptionalClaimsIdTokenList
	_jsii_.Get(
		j,
		"idToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationOptionalClaimsOutputReference) IdTokenInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"idTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationOptionalClaimsOutputReference) InternalValue() *ApplicationOptionalClaims {
	var returns *ApplicationOptionalClaims
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationOptionalClaimsOutputReference) Saml2Token() ApplicationOptionalClaimsSaml2TokenList {
	var returns ApplicationOptionalClaimsSaml2TokenList
	_jsii_.Get(
		j,
		"saml2Token",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationOptionalClaimsOutputReference) Saml2TokenInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"saml2TokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationOptionalClaimsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationOptionalClaimsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApplicationOptionalClaimsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApplicationOptionalClaimsOutputReference {
	_init_.Initialize()

	if err := validateNewApplicationOptionalClaimsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationOptionalClaimsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azuread.application.ApplicationOptionalClaimsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApplicationOptionalClaimsOutputReference_Override(a ApplicationOptionalClaimsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azuread.application.ApplicationOptionalClaimsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApplicationOptionalClaimsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApplicationOptionalClaimsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApplicationOptionalClaimsOutputReference)SetInternalValue(val *ApplicationOptionalClaims) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApplicationOptionalClaimsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApplicationOptionalClaimsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApplicationOptionalClaimsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationOptionalClaimsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApplicationOptionalClaimsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApplicationOptionalClaimsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApplicationOptionalClaimsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApplicationOptionalClaimsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApplicationOptionalClaimsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApplicationOptionalClaimsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApplicationOptionalClaimsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApplicationOptionalClaimsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApplicationOptionalClaimsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationOptionalClaimsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApplicationOptionalClaimsOutputReference) PutAccessToken(value interface{}) {
	if err := a.validatePutAccessTokenParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAccessToken",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationOptionalClaimsOutputReference) PutIdToken(value interface{}) {
	if err := a.validatePutIdTokenParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIdToken",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationOptionalClaimsOutputReference) PutSaml2Token(value interface{}) {
	if err := a.validatePutSaml2TokenParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSaml2Token",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationOptionalClaimsOutputReference) ResetAccessToken() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessToken",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationOptionalClaimsOutputReference) ResetIdToken() {
	_jsii_.InvokeVoid(
		a,
		"resetIdToken",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationOptionalClaimsOutputReference) ResetSaml2Token() {
	_jsii_.InvokeVoid(
		a,
		"resetSaml2Token",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationOptionalClaimsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ApplicationOptionalClaimsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


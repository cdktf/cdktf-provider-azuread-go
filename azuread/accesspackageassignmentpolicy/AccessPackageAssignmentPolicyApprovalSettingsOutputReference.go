package accesspackageassignmentpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azuread-go/azuread/v7/jsii"

	"github.com/cdktf/cdktf-provider-azuread-go/azuread/v7/accesspackageassignmentpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccessPackageAssignmentPolicyApprovalSettingsOutputReference interface {
	cdktf.ComplexObject
	ApprovalRequired() interface{}
	SetApprovalRequired(val interface{})
	ApprovalRequiredForExtension() interface{}
	SetApprovalRequiredForExtension(val interface{})
	ApprovalRequiredForExtensionInput() interface{}
	ApprovalRequiredInput() interface{}
	ApprovalStage() AccessPackageAssignmentPolicyApprovalSettingsApprovalStageList
	ApprovalStageInput() interface{}
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
	InternalValue() *AccessPackageAssignmentPolicyApprovalSettings
	SetInternalValue(val *AccessPackageAssignmentPolicyApprovalSettings)
	RequestorJustificationRequired() interface{}
	SetRequestorJustificationRequired(val interface{})
	RequestorJustificationRequiredInput() interface{}
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
	PutApprovalStage(value interface{})
	ResetApprovalRequired()
	ResetApprovalRequiredForExtension()
	ResetApprovalStage()
	ResetRequestorJustificationRequired()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AccessPackageAssignmentPolicyApprovalSettingsOutputReference
type jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsOutputReference) ApprovalRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"approvalRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsOutputReference) ApprovalRequiredForExtension() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"approvalRequiredForExtension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsOutputReference) ApprovalRequiredForExtensionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"approvalRequiredForExtensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsOutputReference) ApprovalRequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"approvalRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsOutputReference) ApprovalStage() AccessPackageAssignmentPolicyApprovalSettingsApprovalStageList {
	var returns AccessPackageAssignmentPolicyApprovalSettingsApprovalStageList
	_jsii_.Get(
		j,
		"approvalStage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsOutputReference) ApprovalStageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"approvalStageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsOutputReference) InternalValue() *AccessPackageAssignmentPolicyApprovalSettings {
	var returns *AccessPackageAssignmentPolicyApprovalSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsOutputReference) RequestorJustificationRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestorJustificationRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsOutputReference) RequestorJustificationRequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestorJustificationRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAccessPackageAssignmentPolicyApprovalSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AccessPackageAssignmentPolicyApprovalSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewAccessPackageAssignmentPolicyApprovalSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azuread.accessPackageAssignmentPolicy.AccessPackageAssignmentPolicyApprovalSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAccessPackageAssignmentPolicyApprovalSettingsOutputReference_Override(a AccessPackageAssignmentPolicyApprovalSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azuread.accessPackageAssignmentPolicy.AccessPackageAssignmentPolicyApprovalSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsOutputReference)SetApprovalRequired(val interface{}) {
	if err := j.validateSetApprovalRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"approvalRequired",
		val,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsOutputReference)SetApprovalRequiredForExtension(val interface{}) {
	if err := j.validateSetApprovalRequiredForExtensionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"approvalRequiredForExtension",
		val,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsOutputReference)SetInternalValue(val *AccessPackageAssignmentPolicyApprovalSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsOutputReference)SetRequestorJustificationRequired(val interface{}) {
	if err := j.validateSetRequestorJustificationRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestorJustificationRequired",
		val,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsOutputReference) PutApprovalStage(value interface{}) {
	if err := a.validatePutApprovalStageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putApprovalStage",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsOutputReference) ResetApprovalRequired() {
	_jsii_.InvokeVoid(
		a,
		"resetApprovalRequired",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsOutputReference) ResetApprovalRequiredForExtension() {
	_jsii_.InvokeVoid(
		a,
		"resetApprovalRequiredForExtension",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsOutputReference) ResetApprovalStage() {
	_jsii_.InvokeVoid(
		a,
		"resetApprovalStage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsOutputReference) ResetRequestorJustificationRequired() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestorJustificationRequired",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


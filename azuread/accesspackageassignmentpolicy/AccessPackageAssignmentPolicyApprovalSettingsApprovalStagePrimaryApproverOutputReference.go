package accesspackageassignmentpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azuread-go/azuread/v7/jsii"

	"github.com/cdktf/cdktf-provider-azuread-go/azuread/v7/accesspackageassignmentpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverOutputReference interface {
	cdktf.ComplexObject
	Backup() interface{}
	SetBackup(val interface{})
	BackupInput() interface{}
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
	ObjectId() *string
	SetObjectId(val *string)
	ObjectIdInput() *string
	SubjectType() *string
	SetSubjectType(val *string)
	SubjectTypeInput() *string
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
	ResetBackup()
	ResetObjectId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverOutputReference
type jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverOutputReference) Backup() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverOutputReference) BackupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverOutputReference) ObjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverOutputReference) ObjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverOutputReference) SubjectType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverOutputReference) SubjectTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverOutputReference {
	_init_.Initialize()

	if err := validateNewAccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azuread.accessPackageAssignmentPolicy.AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverOutputReference_Override(a AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azuread.accessPackageAssignmentPolicy.AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverOutputReference)SetBackup(val interface{}) {
	if err := j.validateSetBackupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backup",
		val,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverOutputReference)SetObjectId(val *string) {
	if err := j.validateSetObjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectId",
		val,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverOutputReference)SetSubjectType(val *string) {
	if err := j.validateSetSubjectTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subjectType",
		val,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverOutputReference) ResetBackup() {
	_jsii_.InvokeVoid(
		a,
		"resetBackup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverOutputReference) ResetObjectId() {
	_jsii_.InvokeVoid(
		a,
		"resetObjectId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


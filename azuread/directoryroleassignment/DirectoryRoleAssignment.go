package directoryroleassignment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azuread-go/azuread/v5/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azuread-go/azuread/v5/directoryroleassignment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azuread/r/directory_role_assignment azuread_directory_role_assignment}.
type DirectoryRoleAssignment interface {
	cdktf.TerraformResource
	AppScopeId() *string
	SetAppScopeId(val *string)
	AppScopeIdInput() *string
	AppScopeObjectId() *string
	SetAppScopeObjectId(val *string)
	AppScopeObjectIdInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DirectoryScopeId() *string
	SetDirectoryScopeId(val *string)
	DirectoryScopeIdInput() *string
	DirectoryScopeObjectId() *string
	SetDirectoryScopeObjectId(val *string)
	DirectoryScopeObjectIdInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	PrincipalObjectId() *string
	SetPrincipalObjectId(val *string)
	PrincipalObjectIdInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	RoleId() *string
	SetRoleId(val *string)
	RoleIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DirectoryRoleAssignmentTimeoutsOutputReference
	TimeoutsInput() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutTimeouts(value *DirectoryRoleAssignmentTimeouts)
	ResetAppScopeId()
	ResetAppScopeObjectId()
	ResetDirectoryScopeId()
	ResetDirectoryScopeObjectId()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DirectoryRoleAssignment
type jsiiProxy_DirectoryRoleAssignment struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DirectoryRoleAssignment) AppScopeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appScopeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryRoleAssignment) AppScopeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appScopeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryRoleAssignment) AppScopeObjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appScopeObjectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryRoleAssignment) AppScopeObjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appScopeObjectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryRoleAssignment) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryRoleAssignment) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryRoleAssignment) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryRoleAssignment) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryRoleAssignment) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryRoleAssignment) DirectoryScopeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryScopeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryRoleAssignment) DirectoryScopeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryScopeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryRoleAssignment) DirectoryScopeObjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryScopeObjectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryRoleAssignment) DirectoryScopeObjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryScopeObjectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryRoleAssignment) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryRoleAssignment) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryRoleAssignment) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryRoleAssignment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryRoleAssignment) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryRoleAssignment) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryRoleAssignment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryRoleAssignment) PrincipalObjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalObjectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryRoleAssignment) PrincipalObjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalObjectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryRoleAssignment) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryRoleAssignment) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryRoleAssignment) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryRoleAssignment) RoleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryRoleAssignment) RoleIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryRoleAssignment) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryRoleAssignment) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryRoleAssignment) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryRoleAssignment) Timeouts() DirectoryRoleAssignmentTimeoutsOutputReference {
	var returns DirectoryRoleAssignmentTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectoryRoleAssignment) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azuread/r/directory_role_assignment azuread_directory_role_assignment} Resource.
func NewDirectoryRoleAssignment(scope constructs.Construct, id *string, config *DirectoryRoleAssignmentConfig) DirectoryRoleAssignment {
	_init_.Initialize()

	if err := validateNewDirectoryRoleAssignmentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DirectoryRoleAssignment{}

	_jsii_.Create(
		"@cdktf/provider-azuread.directoryRoleAssignment.DirectoryRoleAssignment",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azuread/r/directory_role_assignment azuread_directory_role_assignment} Resource.
func NewDirectoryRoleAssignment_Override(d DirectoryRoleAssignment, scope constructs.Construct, id *string, config *DirectoryRoleAssignmentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azuread.directoryRoleAssignment.DirectoryRoleAssignment",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DirectoryRoleAssignment)SetAppScopeId(val *string) {
	if err := j.validateSetAppScopeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appScopeId",
		val,
	)
}

func (j *jsiiProxy_DirectoryRoleAssignment)SetAppScopeObjectId(val *string) {
	if err := j.validateSetAppScopeObjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appScopeObjectId",
		val,
	)
}

func (j *jsiiProxy_DirectoryRoleAssignment)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DirectoryRoleAssignment)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DirectoryRoleAssignment)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DirectoryRoleAssignment)SetDirectoryScopeId(val *string) {
	if err := j.validateSetDirectoryScopeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"directoryScopeId",
		val,
	)
}

func (j *jsiiProxy_DirectoryRoleAssignment)SetDirectoryScopeObjectId(val *string) {
	if err := j.validateSetDirectoryScopeObjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"directoryScopeObjectId",
		val,
	)
}

func (j *jsiiProxy_DirectoryRoleAssignment)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DirectoryRoleAssignment)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DirectoryRoleAssignment)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DirectoryRoleAssignment)SetPrincipalObjectId(val *string) {
	if err := j.validateSetPrincipalObjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"principalObjectId",
		val,
	)
}

func (j *jsiiProxy_DirectoryRoleAssignment)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DirectoryRoleAssignment)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DirectoryRoleAssignment)SetRoleId(val *string) {
	if err := j.validateSetRoleIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleId",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func DirectoryRoleAssignment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDirectoryRoleAssignment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azuread.directoryRoleAssignment.DirectoryRoleAssignment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DirectoryRoleAssignment_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDirectoryRoleAssignment_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azuread.directoryRoleAssignment.DirectoryRoleAssignment",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DirectoryRoleAssignment_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDirectoryRoleAssignment_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azuread.directoryRoleAssignment.DirectoryRoleAssignment",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DirectoryRoleAssignment_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azuread.directoryRoleAssignment.DirectoryRoleAssignment",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DirectoryRoleAssignment) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DirectoryRoleAssignment) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectoryRoleAssignment) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectoryRoleAssignment) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectoryRoleAssignment) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectoryRoleAssignment) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectoryRoleAssignment) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectoryRoleAssignment) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectoryRoleAssignment) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectoryRoleAssignment) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectoryRoleAssignment) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectoryRoleAssignment) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DirectoryRoleAssignment) PutTimeouts(value *DirectoryRoleAssignmentTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DirectoryRoleAssignment) ResetAppScopeId() {
	_jsii_.InvokeVoid(
		d,
		"resetAppScopeId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DirectoryRoleAssignment) ResetAppScopeObjectId() {
	_jsii_.InvokeVoid(
		d,
		"resetAppScopeObjectId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DirectoryRoleAssignment) ResetDirectoryScopeId() {
	_jsii_.InvokeVoid(
		d,
		"resetDirectoryScopeId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DirectoryRoleAssignment) ResetDirectoryScopeObjectId() {
	_jsii_.InvokeVoid(
		d,
		"resetDirectoryScopeObjectId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DirectoryRoleAssignment) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DirectoryRoleAssignment) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DirectoryRoleAssignment) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DirectoryRoleAssignment) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectoryRoleAssignment) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectoryRoleAssignment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectoryRoleAssignment) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


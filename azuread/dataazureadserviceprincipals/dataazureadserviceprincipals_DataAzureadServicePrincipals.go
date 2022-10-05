package dataazureadserviceprincipals

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azuread-go/azuread/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azuread-go/azuread/v3/dataazureadserviceprincipals/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azuread/d/service_principals azuread_service_principals}.
type DataAzureadServicePrincipals interface {
	cdktf.TerraformDataSource
	ApplicationIds() *[]*string
	SetApplicationIds(val *[]*string)
	ApplicationIdsInput() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	DisplayNames() *[]*string
	SetDisplayNames(val *[]*string)
	DisplayNamesInput() *[]*string
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
	IgnoreMissing() interface{}
	SetIgnoreMissing(val interface{})
	IgnoreMissingInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	ObjectIds() *[]*string
	SetObjectIds(val *[]*string)
	ObjectIdsInput() *[]*string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	ReturnAll() interface{}
	SetReturnAll(val interface{})
	ReturnAllInput() interface{}
	ServicePrincipals() DataAzureadServicePrincipalsServicePrincipalsList
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DataAzureadServicePrincipalsTimeoutsOutputReference
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
	PutTimeouts(value *DataAzureadServicePrincipalsTimeouts)
	ResetApplicationIds()
	ResetDisplayNames()
	ResetId()
	ResetIgnoreMissing()
	ResetObjectIds()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetReturnAll()
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

// The jsii proxy struct for DataAzureadServicePrincipals
type jsiiProxy_DataAzureadServicePrincipals struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAzureadServicePrincipals) ApplicationIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"applicationIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadServicePrincipals) ApplicationIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"applicationIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadServicePrincipals) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadServicePrincipals) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadServicePrincipals) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadServicePrincipals) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadServicePrincipals) DisplayNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"displayNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadServicePrincipals) DisplayNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"displayNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadServicePrincipals) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadServicePrincipals) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadServicePrincipals) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadServicePrincipals) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadServicePrincipals) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadServicePrincipals) IgnoreMissing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreMissing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadServicePrincipals) IgnoreMissingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreMissingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadServicePrincipals) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadServicePrincipals) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadServicePrincipals) ObjectIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"objectIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadServicePrincipals) ObjectIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"objectIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadServicePrincipals) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadServicePrincipals) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadServicePrincipals) ReturnAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"returnAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadServicePrincipals) ReturnAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"returnAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadServicePrincipals) ServicePrincipals() DataAzureadServicePrincipalsServicePrincipalsList {
	var returns DataAzureadServicePrincipalsServicePrincipalsList
	_jsii_.Get(
		j,
		"servicePrincipals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadServicePrincipals) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadServicePrincipals) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadServicePrincipals) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadServicePrincipals) Timeouts() DataAzureadServicePrincipalsTimeoutsOutputReference {
	var returns DataAzureadServicePrincipalsTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadServicePrincipals) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azuread/d/service_principals azuread_service_principals} Data Source.
func NewDataAzureadServicePrincipals(scope constructs.Construct, id *string, config *DataAzureadServicePrincipalsConfig) DataAzureadServicePrincipals {
	_init_.Initialize()

	if err := validateNewDataAzureadServicePrincipalsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAzureadServicePrincipals{}

	_jsii_.Create(
		"@cdktf/provider-azuread.dataAzureadServicePrincipals.DataAzureadServicePrincipals",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azuread/d/service_principals azuread_service_principals} Data Source.
func NewDataAzureadServicePrincipals_Override(d DataAzureadServicePrincipals, scope constructs.Construct, id *string, config *DataAzureadServicePrincipalsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azuread.dataAzureadServicePrincipals.DataAzureadServicePrincipals",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAzureadServicePrincipals)SetApplicationIds(val *[]*string) {
	if err := j.validateSetApplicationIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationIds",
		val,
	)
}

func (j *jsiiProxy_DataAzureadServicePrincipals)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAzureadServicePrincipals)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAzureadServicePrincipals)SetDisplayNames(val *[]*string) {
	if err := j.validateSetDisplayNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayNames",
		val,
	)
}

func (j *jsiiProxy_DataAzureadServicePrincipals)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAzureadServicePrincipals)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAzureadServicePrincipals)SetIgnoreMissing(val interface{}) {
	if err := j.validateSetIgnoreMissingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreMissing",
		val,
	)
}

func (j *jsiiProxy_DataAzureadServicePrincipals)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAzureadServicePrincipals)SetObjectIds(val *[]*string) {
	if err := j.validateSetObjectIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectIds",
		val,
	)
}

func (j *jsiiProxy_DataAzureadServicePrincipals)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAzureadServicePrincipals)SetReturnAll(val interface{}) {
	if err := j.validateSetReturnAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"returnAll",
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
func DataAzureadServicePrincipals_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzureadServicePrincipals_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azuread.dataAzureadServicePrincipals.DataAzureadServicePrincipals",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAzureadServicePrincipals_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azuread.dataAzureadServicePrincipals.DataAzureadServicePrincipals",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAzureadServicePrincipals) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAzureadServicePrincipals) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAzureadServicePrincipals) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAzureadServicePrincipals) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAzureadServicePrincipals) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAzureadServicePrincipals) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAzureadServicePrincipals) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAzureadServicePrincipals) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAzureadServicePrincipals) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAzureadServicePrincipals) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAzureadServicePrincipals) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAzureadServicePrincipals) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAzureadServicePrincipals) PutTimeouts(value *DataAzureadServicePrincipalsTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAzureadServicePrincipals) ResetApplicationIds() {
	_jsii_.InvokeVoid(
		d,
		"resetApplicationIds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzureadServicePrincipals) ResetDisplayNames() {
	_jsii_.InvokeVoid(
		d,
		"resetDisplayNames",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzureadServicePrincipals) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzureadServicePrincipals) ResetIgnoreMissing() {
	_jsii_.InvokeVoid(
		d,
		"resetIgnoreMissing",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzureadServicePrincipals) ResetObjectIds() {
	_jsii_.InvokeVoid(
		d,
		"resetObjectIds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzureadServicePrincipals) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzureadServicePrincipals) ResetReturnAll() {
	_jsii_.InvokeVoid(
		d,
		"resetReturnAll",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzureadServicePrincipals) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzureadServicePrincipals) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzureadServicePrincipals) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzureadServicePrincipals) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzureadServicePrincipals) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


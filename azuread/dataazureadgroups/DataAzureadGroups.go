// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataazureadgroups

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azuread-go/azuread/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azuread-go/azuread/v14/dataazureadgroups/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azuread/3.5.0/docs/data-sources/groups azuread_groups}.
type DataAzureadGroups interface {
	cdktf.TerraformDataSource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayNamePrefix() *string
	SetDisplayNamePrefix(val *string)
	DisplayNamePrefixInput() *string
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
	MailEnabled() interface{}
	SetMailEnabled(val interface{})
	MailEnabledInput() interface{}
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
	SecurityEnabled() interface{}
	SetSecurityEnabled(val interface{})
	SecurityEnabledInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DataAzureadGroupsTimeoutsOutputReference
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
	PutTimeouts(value *DataAzureadGroupsTimeouts)
	ResetDisplayNamePrefix()
	ResetDisplayNames()
	ResetId()
	ResetIgnoreMissing()
	ResetMailEnabled()
	ResetObjectIds()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetReturnAll()
	ResetSecurityEnabled()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAzureadGroups
type jsiiProxy_DataAzureadGroups struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAzureadGroups) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadGroups) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadGroups) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadGroups) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadGroups) DisplayNamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNamePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadGroups) DisplayNamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNamePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadGroups) DisplayNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"displayNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadGroups) DisplayNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"displayNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadGroups) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadGroups) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadGroups) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadGroups) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadGroups) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadGroups) IgnoreMissing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreMissing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadGroups) IgnoreMissingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreMissingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadGroups) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadGroups) MailEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mailEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadGroups) MailEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mailEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadGroups) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadGroups) ObjectIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"objectIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadGroups) ObjectIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"objectIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadGroups) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadGroups) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadGroups) ReturnAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"returnAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadGroups) ReturnAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"returnAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadGroups) SecurityEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"securityEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadGroups) SecurityEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"securityEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadGroups) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadGroups) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadGroups) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadGroups) Timeouts() DataAzureadGroupsTimeoutsOutputReference {
	var returns DataAzureadGroupsTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadGroups) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azuread/3.5.0/docs/data-sources/groups azuread_groups} Data Source.
func NewDataAzureadGroups(scope constructs.Construct, id *string, config *DataAzureadGroupsConfig) DataAzureadGroups {
	_init_.Initialize()

	if err := validateNewDataAzureadGroupsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAzureadGroups{}

	_jsii_.Create(
		"@cdktf/provider-azuread.dataAzureadGroups.DataAzureadGroups",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azuread/3.5.0/docs/data-sources/groups azuread_groups} Data Source.
func NewDataAzureadGroups_Override(d DataAzureadGroups, scope constructs.Construct, id *string, config *DataAzureadGroupsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azuread.dataAzureadGroups.DataAzureadGroups",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAzureadGroups)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAzureadGroups)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAzureadGroups)SetDisplayNamePrefix(val *string) {
	if err := j.validateSetDisplayNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayNamePrefix",
		val,
	)
}

func (j *jsiiProxy_DataAzureadGroups)SetDisplayNames(val *[]*string) {
	if err := j.validateSetDisplayNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayNames",
		val,
	)
}

func (j *jsiiProxy_DataAzureadGroups)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAzureadGroups)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAzureadGroups)SetIgnoreMissing(val interface{}) {
	if err := j.validateSetIgnoreMissingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreMissing",
		val,
	)
}

func (j *jsiiProxy_DataAzureadGroups)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAzureadGroups)SetMailEnabled(val interface{}) {
	if err := j.validateSetMailEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mailEnabled",
		val,
	)
}

func (j *jsiiProxy_DataAzureadGroups)SetObjectIds(val *[]*string) {
	if err := j.validateSetObjectIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectIds",
		val,
	)
}

func (j *jsiiProxy_DataAzureadGroups)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAzureadGroups)SetReturnAll(val interface{}) {
	if err := j.validateSetReturnAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"returnAll",
		val,
	)
}

func (j *jsiiProxy_DataAzureadGroups)SetSecurityEnabled(val interface{}) {
	if err := j.validateSetSecurityEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityEnabled",
		val,
	)
}

// Generates CDKTF code for importing a DataAzureadGroups resource upon running "cdktf plan <stack-name>".
func DataAzureadGroups_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataAzureadGroups_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azuread.dataAzureadGroups.DataAzureadGroups",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
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
func DataAzureadGroups_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzureadGroups_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azuread.dataAzureadGroups.DataAzureadGroups",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAzureadGroups_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzureadGroups_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azuread.dataAzureadGroups.DataAzureadGroups",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAzureadGroups_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzureadGroups_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azuread.dataAzureadGroups.DataAzureadGroups",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAzureadGroups_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azuread.dataAzureadGroups.DataAzureadGroups",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAzureadGroups) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAzureadGroups) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAzureadGroups) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAzureadGroups) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAzureadGroups) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAzureadGroups) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAzureadGroups) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAzureadGroups) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAzureadGroups) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAzureadGroups) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAzureadGroups) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAzureadGroups) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAzureadGroups) PutTimeouts(value *DataAzureadGroupsTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAzureadGroups) ResetDisplayNamePrefix() {
	_jsii_.InvokeVoid(
		d,
		"resetDisplayNamePrefix",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzureadGroups) ResetDisplayNames() {
	_jsii_.InvokeVoid(
		d,
		"resetDisplayNames",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzureadGroups) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzureadGroups) ResetIgnoreMissing() {
	_jsii_.InvokeVoid(
		d,
		"resetIgnoreMissing",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzureadGroups) ResetMailEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetMailEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzureadGroups) ResetObjectIds() {
	_jsii_.InvokeVoid(
		d,
		"resetObjectIds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzureadGroups) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzureadGroups) ResetReturnAll() {
	_jsii_.InvokeVoid(
		d,
		"resetReturnAll",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzureadGroups) ResetSecurityEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetSecurityEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzureadGroups) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzureadGroups) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzureadGroups) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzureadGroups) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzureadGroups) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzureadGroups) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzureadGroups) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


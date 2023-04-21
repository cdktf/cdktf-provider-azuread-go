package dataazureaddomains

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azuread-go/azuread/v7/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azuread-go/azuread/v7/dataazureaddomains/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azuread/2.37.2/docs/data-sources/domains azuread_domains}.
type DataAzureadDomains interface {
	cdktf.TerraformDataSource
	AdminManaged() interface{}
	SetAdminManaged(val interface{})
	AdminManagedInput() interface{}
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
	Domains() DataAzureadDomainsDomainsList
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
	IncludeUnverified() interface{}
	SetIncludeUnverified(val interface{})
	IncludeUnverifiedInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	OnlyDefault() interface{}
	SetOnlyDefault(val interface{})
	OnlyDefaultInput() interface{}
	OnlyInitial() interface{}
	SetOnlyInitial(val interface{})
	OnlyInitialInput() interface{}
	OnlyRoot() interface{}
	SetOnlyRoot(val interface{})
	OnlyRootInput() interface{}
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	SupportsServices() *[]*string
	SetSupportsServices(val *[]*string)
	SupportsServicesInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DataAzureadDomainsTimeoutsOutputReference
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
	PutTimeouts(value *DataAzureadDomainsTimeouts)
	ResetAdminManaged()
	ResetId()
	ResetIncludeUnverified()
	ResetOnlyDefault()
	ResetOnlyInitial()
	ResetOnlyRoot()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSupportsServices()
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

// The jsii proxy struct for DataAzureadDomains
type jsiiProxy_DataAzureadDomains struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAzureadDomains) AdminManaged() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adminManaged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadDomains) AdminManagedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adminManagedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadDomains) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadDomains) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadDomains) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadDomains) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadDomains) Domains() DataAzureadDomainsDomainsList {
	var returns DataAzureadDomainsDomainsList
	_jsii_.Get(
		j,
		"domains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadDomains) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadDomains) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadDomains) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadDomains) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadDomains) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadDomains) IncludeUnverified() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeUnverified",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadDomains) IncludeUnverifiedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeUnverifiedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadDomains) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadDomains) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadDomains) OnlyDefault() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onlyDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadDomains) OnlyDefaultInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onlyDefaultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadDomains) OnlyInitial() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onlyInitial",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadDomains) OnlyInitialInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onlyInitialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadDomains) OnlyRoot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onlyRoot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadDomains) OnlyRootInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onlyRootInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadDomains) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadDomains) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadDomains) SupportsServices() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportsServices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadDomains) SupportsServicesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportsServicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadDomains) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadDomains) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadDomains) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadDomains) Timeouts() DataAzureadDomainsTimeoutsOutputReference {
	var returns DataAzureadDomainsTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadDomains) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azuread/2.37.2/docs/data-sources/domains azuread_domains} Data Source.
func NewDataAzureadDomains(scope constructs.Construct, id *string, config *DataAzureadDomainsConfig) DataAzureadDomains {
	_init_.Initialize()

	if err := validateNewDataAzureadDomainsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAzureadDomains{}

	_jsii_.Create(
		"@cdktf/provider-azuread.dataAzureadDomains.DataAzureadDomains",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azuread/2.37.2/docs/data-sources/domains azuread_domains} Data Source.
func NewDataAzureadDomains_Override(d DataAzureadDomains, scope constructs.Construct, id *string, config *DataAzureadDomainsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azuread.dataAzureadDomains.DataAzureadDomains",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAzureadDomains)SetAdminManaged(val interface{}) {
	if err := j.validateSetAdminManagedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminManaged",
		val,
	)
}

func (j *jsiiProxy_DataAzureadDomains)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAzureadDomains)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAzureadDomains)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAzureadDomains)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAzureadDomains)SetIncludeUnverified(val interface{}) {
	if err := j.validateSetIncludeUnverifiedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeUnverified",
		val,
	)
}

func (j *jsiiProxy_DataAzureadDomains)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAzureadDomains)SetOnlyDefault(val interface{}) {
	if err := j.validateSetOnlyDefaultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onlyDefault",
		val,
	)
}

func (j *jsiiProxy_DataAzureadDomains)SetOnlyInitial(val interface{}) {
	if err := j.validateSetOnlyInitialParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onlyInitial",
		val,
	)
}

func (j *jsiiProxy_DataAzureadDomains)SetOnlyRoot(val interface{}) {
	if err := j.validateSetOnlyRootParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onlyRoot",
		val,
	)
}

func (j *jsiiProxy_DataAzureadDomains)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAzureadDomains)SetSupportsServices(val *[]*string) {
	if err := j.validateSetSupportsServicesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportsServices",
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
func DataAzureadDomains_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzureadDomains_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azuread.dataAzureadDomains.DataAzureadDomains",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAzureadDomains_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzureadDomains_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azuread.dataAzureadDomains.DataAzureadDomains",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAzureadDomains_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzureadDomains_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azuread.dataAzureadDomains.DataAzureadDomains",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAzureadDomains_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azuread.dataAzureadDomains.DataAzureadDomains",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAzureadDomains) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAzureadDomains) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAzureadDomains) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAzureadDomains) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAzureadDomains) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAzureadDomains) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAzureadDomains) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAzureadDomains) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAzureadDomains) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAzureadDomains) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAzureadDomains) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAzureadDomains) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAzureadDomains) PutTimeouts(value *DataAzureadDomainsTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAzureadDomains) ResetAdminManaged() {
	_jsii_.InvokeVoid(
		d,
		"resetAdminManaged",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzureadDomains) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzureadDomains) ResetIncludeUnverified() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludeUnverified",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzureadDomains) ResetOnlyDefault() {
	_jsii_.InvokeVoid(
		d,
		"resetOnlyDefault",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzureadDomains) ResetOnlyInitial() {
	_jsii_.InvokeVoid(
		d,
		"resetOnlyInitial",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzureadDomains) ResetOnlyRoot() {
	_jsii_.InvokeVoid(
		d,
		"resetOnlyRoot",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzureadDomains) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzureadDomains) ResetSupportsServices() {
	_jsii_.InvokeVoid(
		d,
		"resetSupportsServices",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzureadDomains) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzureadDomains) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzureadDomains) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzureadDomains) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzureadDomains) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


package dataazureaduser

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azuread-go/azuread/v9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azuread-go/azuread/v9/dataazureaduser/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azuread/2.40.0/docs/data-sources/user azuread_user}.
type DataAzureadUser interface {
	cdktf.TerraformDataSource
	AccountEnabled() cdktf.IResolvable
	AgeGroup() *string
	BusinessPhones() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	City() *string
	CompanyName() *string
	ConsentProvidedForMinor() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CostCenter() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	Country() *string
	CreationType() *string
	Department() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayName() *string
	Division() *string
	EmployeeId() *string
	SetEmployeeId(val *string)
	EmployeeIdInput() *string
	EmployeeType() *string
	ExternalUserState() *string
	FaxNumber() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GivenName() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	ImAddresses() *[]*string
	JobTitle() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Mail() *string
	SetMail(val *string)
	MailInput() *string
	MailNickname() *string
	SetMailNickname(val *string)
	MailNicknameInput() *string
	ManagerId() *string
	MobilePhone() *string
	// The tree node.
	Node() constructs.Node
	ObjectId() *string
	SetObjectId(val *string)
	ObjectIdInput() *string
	OfficeLocation() *string
	OnpremisesDistinguishedName() *string
	OnpremisesDomainName() *string
	OnpremisesImmutableId() *string
	OnpremisesSamAccountName() *string
	OnpremisesSecurityIdentifier() *string
	OnpremisesSyncEnabled() cdktf.IResolvable
	OnpremisesUserPrincipalName() *string
	OtherMails() *[]*string
	PostalCode() *string
	PreferredLanguage() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	ProxyAddresses() *[]*string
	// Experimental.
	RawOverrides() interface{}
	ShowInAddressList() cdktf.IResolvable
	State() *string
	StreetAddress() *string
	Surname() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DataAzureadUserTimeoutsOutputReference
	TimeoutsInput() interface{}
	UsageLocation() *string
	UserPrincipalName() *string
	SetUserPrincipalName(val *string)
	UserPrincipalNameInput() *string
	UserType() *string
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
	PutTimeouts(value *DataAzureadUserTimeouts)
	ResetEmployeeId()
	ResetId()
	ResetMail()
	ResetMailNickname()
	ResetObjectId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTimeouts()
	ResetUserPrincipalName()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAzureadUser
type jsiiProxy_DataAzureadUser struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAzureadUser) AccountEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"accountEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) AgeGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ageGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) BusinessPhones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"businessPhones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) City() *string {
	var returns *string
	_jsii_.Get(
		j,
		"city",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) CompanyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"companyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) ConsentProvidedForMinor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consentProvidedForMinor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) CostCenter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"costCenter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) Country() *string {
	var returns *string
	_jsii_.Get(
		j,
		"country",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) CreationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) Department() *string {
	var returns *string
	_jsii_.Get(
		j,
		"department",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) Division() *string {
	var returns *string
	_jsii_.Get(
		j,
		"division",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) EmployeeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"employeeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) EmployeeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"employeeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) EmployeeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"employeeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) ExternalUserState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalUserState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) FaxNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"faxNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) GivenName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"givenName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) ImAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"imAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) JobTitle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobTitle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) Mail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) MailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) MailNickname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailNickname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) MailNicknameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailNicknameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) ManagerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) MobilePhone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mobilePhone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) ObjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) ObjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) OfficeLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"officeLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) OnpremisesDistinguishedName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onpremisesDistinguishedName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) OnpremisesDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onpremisesDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) OnpremisesImmutableId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onpremisesImmutableId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) OnpremisesSamAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onpremisesSamAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) OnpremisesSecurityIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onpremisesSecurityIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) OnpremisesSyncEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"onpremisesSyncEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) OnpremisesUserPrincipalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onpremisesUserPrincipalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) OtherMails() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"otherMails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) PostalCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postalCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) PreferredLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) ProxyAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"proxyAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) ShowInAddressList() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"showInAddressList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) StreetAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streetAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) Surname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"surname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) Timeouts() DataAzureadUserTimeoutsOutputReference {
	var returns DataAzureadUserTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) UsageLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usageLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) UserPrincipalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPrincipalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) UserPrincipalNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPrincipalNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzureadUser) UserType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azuread/2.40.0/docs/data-sources/user azuread_user} Data Source.
func NewDataAzureadUser(scope constructs.Construct, id *string, config *DataAzureadUserConfig) DataAzureadUser {
	_init_.Initialize()

	if err := validateNewDataAzureadUserParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAzureadUser{}

	_jsii_.Create(
		"@cdktf/provider-azuread.dataAzureadUser.DataAzureadUser",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azuread/2.40.0/docs/data-sources/user azuread_user} Data Source.
func NewDataAzureadUser_Override(d DataAzureadUser, scope constructs.Construct, id *string, config *DataAzureadUserConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azuread.dataAzureadUser.DataAzureadUser",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAzureadUser)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAzureadUser)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAzureadUser)SetEmployeeId(val *string) {
	if err := j.validateSetEmployeeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"employeeId",
		val,
	)
}

func (j *jsiiProxy_DataAzureadUser)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAzureadUser)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAzureadUser)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAzureadUser)SetMail(val *string) {
	if err := j.validateSetMailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mail",
		val,
	)
}

func (j *jsiiProxy_DataAzureadUser)SetMailNickname(val *string) {
	if err := j.validateSetMailNicknameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mailNickname",
		val,
	)
}

func (j *jsiiProxy_DataAzureadUser)SetObjectId(val *string) {
	if err := j.validateSetObjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectId",
		val,
	)
}

func (j *jsiiProxy_DataAzureadUser)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAzureadUser)SetUserPrincipalName(val *string) {
	if err := j.validateSetUserPrincipalNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userPrincipalName",
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
func DataAzureadUser_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzureadUser_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azuread.dataAzureadUser.DataAzureadUser",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAzureadUser_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzureadUser_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azuread.dataAzureadUser.DataAzureadUser",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAzureadUser_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzureadUser_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azuread.dataAzureadUser.DataAzureadUser",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAzureadUser_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azuread.dataAzureadUser.DataAzureadUser",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAzureadUser) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAzureadUser) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAzureadUser) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAzureadUser) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAzureadUser) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAzureadUser) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAzureadUser) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAzureadUser) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAzureadUser) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAzureadUser) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAzureadUser) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAzureadUser) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAzureadUser) PutTimeouts(value *DataAzureadUserTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAzureadUser) ResetEmployeeId() {
	_jsii_.InvokeVoid(
		d,
		"resetEmployeeId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzureadUser) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzureadUser) ResetMail() {
	_jsii_.InvokeVoid(
		d,
		"resetMail",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzureadUser) ResetMailNickname() {
	_jsii_.InvokeVoid(
		d,
		"resetMailNickname",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzureadUser) ResetObjectId() {
	_jsii_.InvokeVoid(
		d,
		"resetObjectId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzureadUser) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzureadUser) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzureadUser) ResetUserPrincipalName() {
	_jsii_.InvokeVoid(
		d,
		"resetUserPrincipalName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzureadUser) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzureadUser) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzureadUser) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzureadUser) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


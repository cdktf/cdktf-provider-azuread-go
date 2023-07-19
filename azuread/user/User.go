package user

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azuread-go/azuread/v9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azuread-go/azuread/v9/user/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azuread/2.40.0/docs/resources/user azuread_user}.
type User interface {
	cdktf.TerraformResource
	AboutMe() *string
	AccountEnabled() interface{}
	SetAccountEnabled(val interface{})
	AccountEnabledInput() interface{}
	AgeGroup() *string
	SetAgeGroup(val *string)
	AgeGroupInput() *string
	BusinessPhones() *[]*string
	SetBusinessPhones(val *[]*string)
	BusinessPhonesInput() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	City() *string
	SetCity(val *string)
	CityInput() *string
	CompanyName() *string
	SetCompanyName(val *string)
	CompanyNameInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConsentProvidedForMinor() *string
	SetConsentProvidedForMinor(val *string)
	ConsentProvidedForMinorInput() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CostCenter() *string
	SetCostCenter(val *string)
	CostCenterInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	Country() *string
	SetCountry(val *string)
	CountryInput() *string
	CreationType() *string
	Department() *string
	SetDepartment(val *string)
	DepartmentInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisablePasswordExpiration() interface{}
	SetDisablePasswordExpiration(val interface{})
	DisablePasswordExpirationInput() interface{}
	DisableStrongPassword() interface{}
	SetDisableStrongPassword(val interface{})
	DisableStrongPasswordInput() interface{}
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	Division() *string
	SetDivision(val *string)
	DivisionInput() *string
	EmployeeId() *string
	SetEmployeeId(val *string)
	EmployeeIdInput() *string
	EmployeeType() *string
	SetEmployeeType(val *string)
	EmployeeTypeInput() *string
	ExternalUserState() *string
	FaxNumber() *string
	SetFaxNumber(val *string)
	FaxNumberInput() *string
	ForcePasswordChange() interface{}
	SetForcePasswordChange(val interface{})
	ForcePasswordChangeInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GivenName() *string
	SetGivenName(val *string)
	GivenNameInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	ImAddresses() *[]*string
	JobTitle() *string
	SetJobTitle(val *string)
	JobTitleInput() *string
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
	SetManagerId(val *string)
	ManagerIdInput() *string
	MobilePhone() *string
	SetMobilePhone(val *string)
	MobilePhoneInput() *string
	// The tree node.
	Node() constructs.Node
	ObjectId() *string
	OfficeLocation() *string
	SetOfficeLocation(val *string)
	OfficeLocationInput() *string
	OnpremisesDistinguishedName() *string
	OnpremisesDomainName() *string
	OnpremisesImmutableId() *string
	SetOnpremisesImmutableId(val *string)
	OnpremisesImmutableIdInput() *string
	OnpremisesSamAccountName() *string
	OnpremisesSecurityIdentifier() *string
	OnpremisesSyncEnabled() cdktf.IResolvable
	OnpremisesUserPrincipalName() *string
	OtherMails() *[]*string
	SetOtherMails(val *[]*string)
	OtherMailsInput() *[]*string
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	PostalCode() *string
	SetPostalCode(val *string)
	PostalCodeInput() *string
	PreferredLanguage() *string
	SetPreferredLanguage(val *string)
	PreferredLanguageInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	ProxyAddresses() *[]*string
	// Experimental.
	RawOverrides() interface{}
	ShowInAddressList() interface{}
	SetShowInAddressList(val interface{})
	ShowInAddressListInput() interface{}
	State() *string
	SetState(val *string)
	StateInput() *string
	StreetAddress() *string
	SetStreetAddress(val *string)
	StreetAddressInput() *string
	Surname() *string
	SetSurname(val *string)
	SurnameInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() UserTimeoutsOutputReference
	TimeoutsInput() interface{}
	UsageLocation() *string
	SetUsageLocation(val *string)
	UsageLocationInput() *string
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
	PutTimeouts(value *UserTimeouts)
	ResetAccountEnabled()
	ResetAgeGroup()
	ResetBusinessPhones()
	ResetCity()
	ResetCompanyName()
	ResetConsentProvidedForMinor()
	ResetCostCenter()
	ResetCountry()
	ResetDepartment()
	ResetDisablePasswordExpiration()
	ResetDisableStrongPassword()
	ResetDivision()
	ResetEmployeeId()
	ResetEmployeeType()
	ResetFaxNumber()
	ResetForcePasswordChange()
	ResetGivenName()
	ResetId()
	ResetJobTitle()
	ResetMail()
	ResetMailNickname()
	ResetManagerId()
	ResetMobilePhone()
	ResetOfficeLocation()
	ResetOnpremisesImmutableId()
	ResetOtherMails()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPassword()
	ResetPostalCode()
	ResetPreferredLanguage()
	ResetShowInAddressList()
	ResetState()
	ResetStreetAddress()
	ResetSurname()
	ResetTimeouts()
	ResetUsageLocation()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for User
type jsiiProxy_User struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_User) AboutMe() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aboutMe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) AccountEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accountEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) AccountEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accountEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) AgeGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ageGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) AgeGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ageGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) BusinessPhones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"businessPhones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) BusinessPhonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"businessPhonesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) City() *string {
	var returns *string
	_jsii_.Get(
		j,
		"city",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) CityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) CompanyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"companyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) CompanyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"companyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ConsentProvidedForMinor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consentProvidedForMinor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ConsentProvidedForMinorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consentProvidedForMinorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) CostCenter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"costCenter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) CostCenterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"costCenterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Country() *string {
	var returns *string
	_jsii_.Get(
		j,
		"country",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) CountryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) CreationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Department() *string {
	var returns *string
	_jsii_.Get(
		j,
		"department",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DepartmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"departmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DisablePasswordExpiration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disablePasswordExpiration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DisablePasswordExpirationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disablePasswordExpirationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DisableStrongPassword() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableStrongPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DisableStrongPasswordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableStrongPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Division() *string {
	var returns *string
	_jsii_.Get(
		j,
		"division",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DivisionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"divisionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) EmployeeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"employeeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) EmployeeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"employeeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) EmployeeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"employeeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) EmployeeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"employeeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ExternalUserState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalUserState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) FaxNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"faxNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) FaxNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"faxNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ForcePasswordChange() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forcePasswordChange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ForcePasswordChangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forcePasswordChangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) GivenName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"givenName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) GivenNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"givenNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ImAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"imAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) JobTitle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobTitle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) JobTitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobTitleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Mail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) MailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) MailNickname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailNickname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) MailNicknameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailNicknameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ManagerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ManagerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) MobilePhone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mobilePhone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) MobilePhoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mobilePhoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ObjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) OfficeLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"officeLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) OfficeLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"officeLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) OnpremisesDistinguishedName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onpremisesDistinguishedName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) OnpremisesDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onpremisesDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) OnpremisesImmutableId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onpremisesImmutableId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) OnpremisesImmutableIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onpremisesImmutableIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) OnpremisesSamAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onpremisesSamAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) OnpremisesSecurityIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onpremisesSecurityIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) OnpremisesSyncEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"onpremisesSyncEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) OnpremisesUserPrincipalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onpremisesUserPrincipalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) OtherMails() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"otherMails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) OtherMailsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"otherMailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) PostalCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postalCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) PostalCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postalCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) PreferredLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) PreferredLanguageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredLanguageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ProxyAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"proxyAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ShowInAddressList() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showInAddressList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ShowInAddressListInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showInAddressListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) StateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) StreetAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streetAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) StreetAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streetAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Surname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"surname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) SurnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"surnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Timeouts() UserTimeoutsOutputReference {
	var returns UserTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) UsageLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usageLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) UsageLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usageLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) UserPrincipalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPrincipalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) UserPrincipalNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPrincipalNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) UserType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azuread/2.40.0/docs/resources/user azuread_user} Resource.
func NewUser(scope constructs.Construct, id *string, config *UserConfig) User {
	_init_.Initialize()

	if err := validateNewUserParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_User{}

	_jsii_.Create(
		"@cdktf/provider-azuread.user.User",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azuread/2.40.0/docs/resources/user azuread_user} Resource.
func NewUser_Override(u User, scope constructs.Construct, id *string, config *UserConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azuread.user.User",
		[]interface{}{scope, id, config},
		u,
	)
}

func (j *jsiiProxy_User)SetAccountEnabled(val interface{}) {
	if err := j.validateSetAccountEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountEnabled",
		val,
	)
}

func (j *jsiiProxy_User)SetAgeGroup(val *string) {
	if err := j.validateSetAgeGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ageGroup",
		val,
	)
}

func (j *jsiiProxy_User)SetBusinessPhones(val *[]*string) {
	if err := j.validateSetBusinessPhonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"businessPhones",
		val,
	)
}

func (j *jsiiProxy_User)SetCity(val *string) {
	if err := j.validateSetCityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"city",
		val,
	)
}

func (j *jsiiProxy_User)SetCompanyName(val *string) {
	if err := j.validateSetCompanyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"companyName",
		val,
	)
}

func (j *jsiiProxy_User)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_User)SetConsentProvidedForMinor(val *string) {
	if err := j.validateSetConsentProvidedForMinorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consentProvidedForMinor",
		val,
	)
}

func (j *jsiiProxy_User)SetCostCenter(val *string) {
	if err := j.validateSetCostCenterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"costCenter",
		val,
	)
}

func (j *jsiiProxy_User)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_User)SetCountry(val *string) {
	if err := j.validateSetCountryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"country",
		val,
	)
}

func (j *jsiiProxy_User)SetDepartment(val *string) {
	if err := j.validateSetDepartmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"department",
		val,
	)
}

func (j *jsiiProxy_User)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_User)SetDisablePasswordExpiration(val interface{}) {
	if err := j.validateSetDisablePasswordExpirationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disablePasswordExpiration",
		val,
	)
}

func (j *jsiiProxy_User)SetDisableStrongPassword(val interface{}) {
	if err := j.validateSetDisableStrongPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableStrongPassword",
		val,
	)
}

func (j *jsiiProxy_User)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_User)SetDivision(val *string) {
	if err := j.validateSetDivisionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"division",
		val,
	)
}

func (j *jsiiProxy_User)SetEmployeeId(val *string) {
	if err := j.validateSetEmployeeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"employeeId",
		val,
	)
}

func (j *jsiiProxy_User)SetEmployeeType(val *string) {
	if err := j.validateSetEmployeeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"employeeType",
		val,
	)
}

func (j *jsiiProxy_User)SetFaxNumber(val *string) {
	if err := j.validateSetFaxNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"faxNumber",
		val,
	)
}

func (j *jsiiProxy_User)SetForcePasswordChange(val interface{}) {
	if err := j.validateSetForcePasswordChangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forcePasswordChange",
		val,
	)
}

func (j *jsiiProxy_User)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_User)SetGivenName(val *string) {
	if err := j.validateSetGivenNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"givenName",
		val,
	)
}

func (j *jsiiProxy_User)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_User)SetJobTitle(val *string) {
	if err := j.validateSetJobTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobTitle",
		val,
	)
}

func (j *jsiiProxy_User)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_User)SetMail(val *string) {
	if err := j.validateSetMailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mail",
		val,
	)
}

func (j *jsiiProxy_User)SetMailNickname(val *string) {
	if err := j.validateSetMailNicknameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mailNickname",
		val,
	)
}

func (j *jsiiProxy_User)SetManagerId(val *string) {
	if err := j.validateSetManagerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managerId",
		val,
	)
}

func (j *jsiiProxy_User)SetMobilePhone(val *string) {
	if err := j.validateSetMobilePhoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mobilePhone",
		val,
	)
}

func (j *jsiiProxy_User)SetOfficeLocation(val *string) {
	if err := j.validateSetOfficeLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"officeLocation",
		val,
	)
}

func (j *jsiiProxy_User)SetOnpremisesImmutableId(val *string) {
	if err := j.validateSetOnpremisesImmutableIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onpremisesImmutableId",
		val,
	)
}

func (j *jsiiProxy_User)SetOtherMails(val *[]*string) {
	if err := j.validateSetOtherMailsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"otherMails",
		val,
	)
}

func (j *jsiiProxy_User)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_User)SetPostalCode(val *string) {
	if err := j.validateSetPostalCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postalCode",
		val,
	)
}

func (j *jsiiProxy_User)SetPreferredLanguage(val *string) {
	if err := j.validateSetPreferredLanguageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredLanguage",
		val,
	)
}

func (j *jsiiProxy_User)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_User)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_User)SetShowInAddressList(val interface{}) {
	if err := j.validateSetShowInAddressListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showInAddressList",
		val,
	)
}

func (j *jsiiProxy_User)SetState(val *string) {
	if err := j.validateSetStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"state",
		val,
	)
}

func (j *jsiiProxy_User)SetStreetAddress(val *string) {
	if err := j.validateSetStreetAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streetAddress",
		val,
	)
}

func (j *jsiiProxy_User)SetSurname(val *string) {
	if err := j.validateSetSurnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"surname",
		val,
	)
}

func (j *jsiiProxy_User)SetUsageLocation(val *string) {
	if err := j.validateSetUsageLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usageLocation",
		val,
	)
}

func (j *jsiiProxy_User)SetUserPrincipalName(val *string) {
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
func User_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateUser_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azuread.user.User",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func User_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateUser_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azuread.user.User",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func User_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateUser_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azuread.user.User",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func User_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azuread.user.User",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (u *jsiiProxy_User) AddOverride(path *string, value interface{}) {
	if err := u.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (u *jsiiProxy_User) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := u.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		u,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := u.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		u,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := u.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		u,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := u.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		u,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := u.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		u,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := u.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		u,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := u.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		u,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetStringAttribute(terraformAttribute *string) *string {
	if err := u.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		u,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := u.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		u,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := u.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		u,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) OverrideLogicalId(newLogicalId *string) {
	if err := u.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (u *jsiiProxy_User) PutTimeouts(value *UserTimeouts) {
	if err := u.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (u *jsiiProxy_User) ResetAccountEnabled() {
	_jsii_.InvokeVoid(
		u,
		"resetAccountEnabled",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetAgeGroup() {
	_jsii_.InvokeVoid(
		u,
		"resetAgeGroup",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetBusinessPhones() {
	_jsii_.InvokeVoid(
		u,
		"resetBusinessPhones",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetCity() {
	_jsii_.InvokeVoid(
		u,
		"resetCity",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetCompanyName() {
	_jsii_.InvokeVoid(
		u,
		"resetCompanyName",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetConsentProvidedForMinor() {
	_jsii_.InvokeVoid(
		u,
		"resetConsentProvidedForMinor",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetCostCenter() {
	_jsii_.InvokeVoid(
		u,
		"resetCostCenter",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetCountry() {
	_jsii_.InvokeVoid(
		u,
		"resetCountry",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetDepartment() {
	_jsii_.InvokeVoid(
		u,
		"resetDepartment",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetDisablePasswordExpiration() {
	_jsii_.InvokeVoid(
		u,
		"resetDisablePasswordExpiration",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetDisableStrongPassword() {
	_jsii_.InvokeVoid(
		u,
		"resetDisableStrongPassword",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetDivision() {
	_jsii_.InvokeVoid(
		u,
		"resetDivision",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetEmployeeId() {
	_jsii_.InvokeVoid(
		u,
		"resetEmployeeId",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetEmployeeType() {
	_jsii_.InvokeVoid(
		u,
		"resetEmployeeType",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetFaxNumber() {
	_jsii_.InvokeVoid(
		u,
		"resetFaxNumber",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetForcePasswordChange() {
	_jsii_.InvokeVoid(
		u,
		"resetForcePasswordChange",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetGivenName() {
	_jsii_.InvokeVoid(
		u,
		"resetGivenName",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetId() {
	_jsii_.InvokeVoid(
		u,
		"resetId",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetJobTitle() {
	_jsii_.InvokeVoid(
		u,
		"resetJobTitle",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetMail() {
	_jsii_.InvokeVoid(
		u,
		"resetMail",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetMailNickname() {
	_jsii_.InvokeVoid(
		u,
		"resetMailNickname",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetManagerId() {
	_jsii_.InvokeVoid(
		u,
		"resetManagerId",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetMobilePhone() {
	_jsii_.InvokeVoid(
		u,
		"resetMobilePhone",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetOfficeLocation() {
	_jsii_.InvokeVoid(
		u,
		"resetOfficeLocation",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetOnpremisesImmutableId() {
	_jsii_.InvokeVoid(
		u,
		"resetOnpremisesImmutableId",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetOtherMails() {
	_jsii_.InvokeVoid(
		u,
		"resetOtherMails",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		u,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetPassword() {
	_jsii_.InvokeVoid(
		u,
		"resetPassword",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetPostalCode() {
	_jsii_.InvokeVoid(
		u,
		"resetPostalCode",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetPreferredLanguage() {
	_jsii_.InvokeVoid(
		u,
		"resetPreferredLanguage",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetShowInAddressList() {
	_jsii_.InvokeVoid(
		u,
		"resetShowInAddressList",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetState() {
	_jsii_.InvokeVoid(
		u,
		"resetState",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetStreetAddress() {
	_jsii_.InvokeVoid(
		u,
		"resetStreetAddress",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetSurname() {
	_jsii_.InvokeVoid(
		u,
		"resetSurname",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetTimeouts() {
	_jsii_.InvokeVoid(
		u,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetUsageLocation() {
	_jsii_.InvokeVoid(
		u,
		"resetUsageLocation",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		u,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


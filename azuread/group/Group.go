package group

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azuread-go/azuread/v9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azuread-go/azuread/v9/group/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azuread/2.41.0/docs/resources/group azuread_group}.
type Group interface {
	cdktf.TerraformResource
	AdministrativeUnitIds() *[]*string
	SetAdministrativeUnitIds(val *[]*string)
	AdministrativeUnitIdsInput() *[]*string
	AssignableToRole() interface{}
	SetAssignableToRole(val interface{})
	AssignableToRoleInput() interface{}
	AutoSubscribeNewMembers() interface{}
	SetAutoSubscribeNewMembers(val interface{})
	AutoSubscribeNewMembersInput() interface{}
	Behaviors() *[]*string
	SetBehaviors(val *[]*string)
	BehaviorsInput() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	DynamicMembership() GroupDynamicMembershipOutputReference
	DynamicMembershipInput() *GroupDynamicMembership
	ExternalSendersAllowed() interface{}
	SetExternalSendersAllowed(val interface{})
	ExternalSendersAllowedInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HideFromAddressLists() interface{}
	SetHideFromAddressLists(val interface{})
	HideFromAddressListsInput() interface{}
	HideFromOutlookClients() interface{}
	SetHideFromOutlookClients(val interface{})
	HideFromOutlookClientsInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Mail() *string
	MailEnabled() interface{}
	SetMailEnabled(val interface{})
	MailEnabledInput() interface{}
	MailNickname() *string
	SetMailNickname(val *string)
	MailNicknameInput() *string
	Members() *[]*string
	SetMembers(val *[]*string)
	MembersInput() *[]*string
	// The tree node.
	Node() constructs.Node
	ObjectId() *string
	OnpremisesDomainName() *string
	OnpremisesGroupType() *string
	SetOnpremisesGroupType(val *string)
	OnpremisesGroupTypeInput() *string
	OnpremisesNetbiosName() *string
	OnpremisesSamAccountName() *string
	OnpremisesSecurityIdentifier() *string
	OnpremisesSyncEnabled() cdktf.IResolvable
	Owners() *[]*string
	SetOwners(val *[]*string)
	OwnersInput() *[]*string
	PreferredLanguage() *string
	PreventDuplicateNames() interface{}
	SetPreventDuplicateNames(val interface{})
	PreventDuplicateNamesInput() interface{}
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	ProvisioningOptions() *[]*string
	SetProvisioningOptions(val *[]*string)
	ProvisioningOptionsInput() *[]*string
	ProxyAddresses() *[]*string
	// Experimental.
	RawOverrides() interface{}
	SecurityEnabled() interface{}
	SetSecurityEnabled(val interface{})
	SecurityEnabledInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Theme() *string
	SetTheme(val *string)
	ThemeInput() *string
	Timeouts() GroupTimeoutsOutputReference
	TimeoutsInput() interface{}
	Types() *[]*string
	SetTypes(val *[]*string)
	TypesInput() *[]*string
	Visibility() *string
	SetVisibility(val *string)
	VisibilityInput() *string
	WritebackEnabled() interface{}
	SetWritebackEnabled(val interface{})
	WritebackEnabledInput() interface{}
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
	PutDynamicMembership(value *GroupDynamicMembership)
	PutTimeouts(value *GroupTimeouts)
	ResetAdministrativeUnitIds()
	ResetAssignableToRole()
	ResetAutoSubscribeNewMembers()
	ResetBehaviors()
	ResetDescription()
	ResetDynamicMembership()
	ResetExternalSendersAllowed()
	ResetHideFromAddressLists()
	ResetHideFromOutlookClients()
	ResetId()
	ResetMailEnabled()
	ResetMailNickname()
	ResetMembers()
	ResetOnpremisesGroupType()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetOwners()
	ResetPreventDuplicateNames()
	ResetProvisioningOptions()
	ResetSecurityEnabled()
	ResetTheme()
	ResetTimeouts()
	ResetTypes()
	ResetVisibility()
	ResetWritebackEnabled()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Group
type jsiiProxy_Group struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Group) AdministrativeUnitIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"administrativeUnitIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) AdministrativeUnitIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"administrativeUnitIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) AssignableToRole() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assignableToRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) AssignableToRoleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assignableToRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) AutoSubscribeNewMembers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoSubscribeNewMembers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) AutoSubscribeNewMembersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoSubscribeNewMembersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) Behaviors() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"behaviors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) BehaviorsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"behaviorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) DynamicMembership() GroupDynamicMembershipOutputReference {
	var returns GroupDynamicMembershipOutputReference
	_jsii_.Get(
		j,
		"dynamicMembership",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) DynamicMembershipInput() *GroupDynamicMembership {
	var returns *GroupDynamicMembership
	_jsii_.Get(
		j,
		"dynamicMembershipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) ExternalSendersAllowed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalSendersAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) ExternalSendersAllowedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalSendersAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) HideFromAddressLists() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideFromAddressLists",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) HideFromAddressListsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideFromAddressListsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) HideFromOutlookClients() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideFromOutlookClients",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) HideFromOutlookClientsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideFromOutlookClientsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) Mail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) MailEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mailEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) MailEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mailEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) MailNickname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailNickname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) MailNicknameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailNicknameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) Members() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"members",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) MembersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"membersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) ObjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) OnpremisesDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onpremisesDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) OnpremisesGroupType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onpremisesGroupType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) OnpremisesGroupTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onpremisesGroupTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) OnpremisesNetbiosName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onpremisesNetbiosName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) OnpremisesSamAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onpremisesSamAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) OnpremisesSecurityIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onpremisesSecurityIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) OnpremisesSyncEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"onpremisesSyncEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) Owners() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"owners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) OwnersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ownersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) PreferredLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) PreventDuplicateNames() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preventDuplicateNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) PreventDuplicateNamesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preventDuplicateNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) ProvisioningOptions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"provisioningOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) ProvisioningOptionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"provisioningOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) ProxyAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"proxyAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) SecurityEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"securityEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) SecurityEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"securityEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) Theme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"theme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) ThemeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"themeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) Timeouts() GroupTimeoutsOutputReference {
	var returns GroupTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) Types() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"types",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) TypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"typesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) Visibility() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) VisibilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visibilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) WritebackEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"writebackEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Group) WritebackEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"writebackEnabledInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azuread/2.41.0/docs/resources/group azuread_group} Resource.
func NewGroup(scope constructs.Construct, id *string, config *GroupConfig) Group {
	_init_.Initialize()

	if err := validateNewGroupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Group{}

	_jsii_.Create(
		"@cdktf/provider-azuread.group.Group",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azuread/2.41.0/docs/resources/group azuread_group} Resource.
func NewGroup_Override(g Group, scope constructs.Construct, id *string, config *GroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azuread.group.Group",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_Group)SetAdministrativeUnitIds(val *[]*string) {
	if err := j.validateSetAdministrativeUnitIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"administrativeUnitIds",
		val,
	)
}

func (j *jsiiProxy_Group)SetAssignableToRole(val interface{}) {
	if err := j.validateSetAssignableToRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assignableToRole",
		val,
	)
}

func (j *jsiiProxy_Group)SetAutoSubscribeNewMembers(val interface{}) {
	if err := j.validateSetAutoSubscribeNewMembersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoSubscribeNewMembers",
		val,
	)
}

func (j *jsiiProxy_Group)SetBehaviors(val *[]*string) {
	if err := j.validateSetBehaviorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"behaviors",
		val,
	)
}

func (j *jsiiProxy_Group)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Group)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Group)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Group)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Group)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_Group)SetExternalSendersAllowed(val interface{}) {
	if err := j.validateSetExternalSendersAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalSendersAllowed",
		val,
	)
}

func (j *jsiiProxy_Group)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Group)SetHideFromAddressLists(val interface{}) {
	if err := j.validateSetHideFromAddressListsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hideFromAddressLists",
		val,
	)
}

func (j *jsiiProxy_Group)SetHideFromOutlookClients(val interface{}) {
	if err := j.validateSetHideFromOutlookClientsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hideFromOutlookClients",
		val,
	)
}

func (j *jsiiProxy_Group)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Group)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Group)SetMailEnabled(val interface{}) {
	if err := j.validateSetMailEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mailEnabled",
		val,
	)
}

func (j *jsiiProxy_Group)SetMailNickname(val *string) {
	if err := j.validateSetMailNicknameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mailNickname",
		val,
	)
}

func (j *jsiiProxy_Group)SetMembers(val *[]*string) {
	if err := j.validateSetMembersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"members",
		val,
	)
}

func (j *jsiiProxy_Group)SetOnpremisesGroupType(val *string) {
	if err := j.validateSetOnpremisesGroupTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onpremisesGroupType",
		val,
	)
}

func (j *jsiiProxy_Group)SetOwners(val *[]*string) {
	if err := j.validateSetOwnersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"owners",
		val,
	)
}

func (j *jsiiProxy_Group)SetPreventDuplicateNames(val interface{}) {
	if err := j.validateSetPreventDuplicateNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preventDuplicateNames",
		val,
	)
}

func (j *jsiiProxy_Group)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Group)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Group)SetProvisioningOptions(val *[]*string) {
	if err := j.validateSetProvisioningOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioningOptions",
		val,
	)
}

func (j *jsiiProxy_Group)SetSecurityEnabled(val interface{}) {
	if err := j.validateSetSecurityEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityEnabled",
		val,
	)
}

func (j *jsiiProxy_Group)SetTheme(val *string) {
	if err := j.validateSetThemeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"theme",
		val,
	)
}

func (j *jsiiProxy_Group)SetTypes(val *[]*string) {
	if err := j.validateSetTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"types",
		val,
	)
}

func (j *jsiiProxy_Group)SetVisibility(val *string) {
	if err := j.validateSetVisibilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visibility",
		val,
	)
}

func (j *jsiiProxy_Group)SetWritebackEnabled(val interface{}) {
	if err := j.validateSetWritebackEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"writebackEnabled",
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
func Group_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azuread.group.Group",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Group_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGroup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azuread.group.Group",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Group_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGroup_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azuread.group.Group",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Group_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azuread.group.Group",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_Group) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_Group) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Group) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Group) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Group) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Group) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Group) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Group) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Group) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Group) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Group) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Group) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_Group) PutDynamicMembership(value *GroupDynamicMembership) {
	if err := g.validatePutDynamicMembershipParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDynamicMembership",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_Group) PutTimeouts(value *GroupTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_Group) ResetAdministrativeUnitIds() {
	_jsii_.InvokeVoid(
		g,
		"resetAdministrativeUnitIds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetAssignableToRole() {
	_jsii_.InvokeVoid(
		g,
		"resetAssignableToRole",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetAutoSubscribeNewMembers() {
	_jsii_.InvokeVoid(
		g,
		"resetAutoSubscribeNewMembers",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetBehaviors() {
	_jsii_.InvokeVoid(
		g,
		"resetBehaviors",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetDynamicMembership() {
	_jsii_.InvokeVoid(
		g,
		"resetDynamicMembership",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetExternalSendersAllowed() {
	_jsii_.InvokeVoid(
		g,
		"resetExternalSendersAllowed",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetHideFromAddressLists() {
	_jsii_.InvokeVoid(
		g,
		"resetHideFromAddressLists",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetHideFromOutlookClients() {
	_jsii_.InvokeVoid(
		g,
		"resetHideFromOutlookClients",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetMailEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetMailEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetMailNickname() {
	_jsii_.InvokeVoid(
		g,
		"resetMailNickname",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetMembers() {
	_jsii_.InvokeVoid(
		g,
		"resetMembers",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetOnpremisesGroupType() {
	_jsii_.InvokeVoid(
		g,
		"resetOnpremisesGroupType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetOwners() {
	_jsii_.InvokeVoid(
		g,
		"resetOwners",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetPreventDuplicateNames() {
	_jsii_.InvokeVoid(
		g,
		"resetPreventDuplicateNames",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetProvisioningOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetProvisioningOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetSecurityEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetSecurityEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetTheme() {
	_jsii_.InvokeVoid(
		g,
		"resetTheme",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetTypes() {
	_jsii_.InvokeVoid(
		g,
		"resetTypes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetVisibility() {
	_jsii_.InvokeVoid(
		g,
		"resetVisibility",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) ResetWritebackEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetWritebackEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Group) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Group) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Group) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Group) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


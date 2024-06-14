// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privilegedaccessgroupassignmentschedule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azuread-go/azuread/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azuread-go/azuread/v12/privilegedaccessgroupassignmentschedule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azuread/2.52.0/docs/resources/privileged_access_group_assignment_schedule azuread_privileged_access_group_assignment_schedule}.
type PrivilegedAccessGroupAssignmentSchedule interface {
	cdktf.TerraformResource
	AssignmentType() *string
	SetAssignmentType(val *string)
	AssignmentTypeInput() *string
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
	Duration() *string
	SetDuration(val *string)
	DurationInput() *string
	ExpirationDate() *string
	SetExpirationDate(val *string)
	ExpirationDateInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GroupId() *string
	SetGroupId(val *string)
	GroupIdInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Justification() *string
	SetJustification(val *string)
	JustificationInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	PermanentAssignment() interface{}
	SetPermanentAssignment(val interface{})
	PermanentAssignmentInput() interface{}
	PrincipalId() *string
	SetPrincipalId(val *string)
	PrincipalIdInput() *string
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
	StartDate() *string
	SetStartDate(val *string)
	StartDateInput() *string
	Status() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TicketNumber() *string
	SetTicketNumber(val *string)
	TicketNumberInput() *string
	TicketSystem() *string
	SetTicketSystem(val *string)
	TicketSystemInput() *string
	Timeouts() PrivilegedAccessGroupAssignmentScheduleTimeoutsOutputReference
	TimeoutsInput() interface{}
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutTimeouts(value *PrivilegedAccessGroupAssignmentScheduleTimeouts)
	ResetDuration()
	ResetExpirationDate()
	ResetId()
	ResetJustification()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPermanentAssignment()
	ResetStartDate()
	ResetTicketNumber()
	ResetTicketSystem()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
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

// The jsii proxy struct for PrivilegedAccessGroupAssignmentSchedule
type jsiiProxy_PrivilegedAccessGroupAssignmentSchedule struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) AssignmentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assignmentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) AssignmentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assignmentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) Duration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"duration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) DurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"durationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) ExpirationDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expirationDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) ExpirationDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expirationDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) GroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) GroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) Justification() *string {
	var returns *string
	_jsii_.Get(
		j,
		"justification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) JustificationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"justificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) PermanentAssignment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"permanentAssignment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) PermanentAssignmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"permanentAssignmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) PrincipalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) PrincipalIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) StartDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) StartDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) TicketNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ticketNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) TicketNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ticketNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) TicketSystem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ticketSystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) TicketSystemInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ticketSystemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) Timeouts() PrivilegedAccessGroupAssignmentScheduleTimeoutsOutputReference {
	var returns PrivilegedAccessGroupAssignmentScheduleTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azuread/2.52.0/docs/resources/privileged_access_group_assignment_schedule azuread_privileged_access_group_assignment_schedule} Resource.
func NewPrivilegedAccessGroupAssignmentSchedule(scope constructs.Construct, id *string, config *PrivilegedAccessGroupAssignmentScheduleConfig) PrivilegedAccessGroupAssignmentSchedule {
	_init_.Initialize()

	if err := validateNewPrivilegedAccessGroupAssignmentScheduleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_PrivilegedAccessGroupAssignmentSchedule{}

	_jsii_.Create(
		"@cdktf/provider-azuread.privilegedAccessGroupAssignmentSchedule.PrivilegedAccessGroupAssignmentSchedule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azuread/2.52.0/docs/resources/privileged_access_group_assignment_schedule azuread_privileged_access_group_assignment_schedule} Resource.
func NewPrivilegedAccessGroupAssignmentSchedule_Override(p PrivilegedAccessGroupAssignmentSchedule, scope constructs.Construct, id *string, config *PrivilegedAccessGroupAssignmentScheduleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azuread.privilegedAccessGroupAssignmentSchedule.PrivilegedAccessGroupAssignmentSchedule",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule)SetAssignmentType(val *string) {
	if err := j.validateSetAssignmentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assignmentType",
		val,
	)
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule)SetDuration(val *string) {
	if err := j.validateSetDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"duration",
		val,
	)
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule)SetExpirationDate(val *string) {
	if err := j.validateSetExpirationDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expirationDate",
		val,
	)
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule)SetGroupId(val *string) {
	if err := j.validateSetGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupId",
		val,
	)
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule)SetJustification(val *string) {
	if err := j.validateSetJustificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"justification",
		val,
	)
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule)SetPermanentAssignment(val interface{}) {
	if err := j.validateSetPermanentAssignmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permanentAssignment",
		val,
	)
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule)SetPrincipalId(val *string) {
	if err := j.validateSetPrincipalIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"principalId",
		val,
	)
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule)SetStartDate(val *string) {
	if err := j.validateSetStartDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startDate",
		val,
	)
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule)SetTicketNumber(val *string) {
	if err := j.validateSetTicketNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ticketNumber",
		val,
	)
}

func (j *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule)SetTicketSystem(val *string) {
	if err := j.validateSetTicketSystemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ticketSystem",
		val,
	)
}

// Generates CDKTF code for importing a PrivilegedAccessGroupAssignmentSchedule resource upon running "cdktf plan <stack-name>".
func PrivilegedAccessGroupAssignmentSchedule_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validatePrivilegedAccessGroupAssignmentSchedule_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azuread.privilegedAccessGroupAssignmentSchedule.PrivilegedAccessGroupAssignmentSchedule",
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
func PrivilegedAccessGroupAssignmentSchedule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePrivilegedAccessGroupAssignmentSchedule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azuread.privilegedAccessGroupAssignmentSchedule.PrivilegedAccessGroupAssignmentSchedule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PrivilegedAccessGroupAssignmentSchedule_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePrivilegedAccessGroupAssignmentSchedule_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azuread.privilegedAccessGroupAssignmentSchedule.PrivilegedAccessGroupAssignmentSchedule",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PrivilegedAccessGroupAssignmentSchedule_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePrivilegedAccessGroupAssignmentSchedule_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azuread.privilegedAccessGroupAssignmentSchedule.PrivilegedAccessGroupAssignmentSchedule",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PrivilegedAccessGroupAssignmentSchedule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azuread.privilegedAccessGroupAssignmentSchedule.PrivilegedAccessGroupAssignmentSchedule",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) AddMoveTarget(moveTarget *string) {
	if err := p.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (p *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := p.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (p *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) MoveFromId(id *string) {
	if err := p.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveFromId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) MoveTo(moveTarget *string, index interface{}) {
	if err := p.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (p *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) MoveToId(id *string) {
	if err := p.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveToId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) PutTimeouts(value *PrivilegedAccessGroupAssignmentScheduleTimeouts) {
	if err := p.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) ResetDuration() {
	_jsii_.InvokeVoid(
		p,
		"resetDuration",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) ResetExpirationDate() {
	_jsii_.InvokeVoid(
		p,
		"resetExpirationDate",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) ResetId() {
	_jsii_.InvokeVoid(
		p,
		"resetId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) ResetJustification() {
	_jsii_.InvokeVoid(
		p,
		"resetJustification",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) ResetPermanentAssignment() {
	_jsii_.InvokeVoid(
		p,
		"resetPermanentAssignment",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) ResetStartDate() {
	_jsii_.InvokeVoid(
		p,
		"resetStartDate",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) ResetTicketNumber() {
	_jsii_.InvokeVoid(
		p,
		"resetTicketNumber",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) ResetTicketSystem() {
	_jsii_.InvokeVoid(
		p,
		"resetTicketSystem",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) ResetTimeouts() {
	_jsii_.InvokeVoid(
		p,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivilegedAccessGroupAssignmentSchedule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


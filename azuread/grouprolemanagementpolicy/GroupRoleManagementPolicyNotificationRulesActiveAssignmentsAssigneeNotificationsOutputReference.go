// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grouprolemanagementpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azuread-go/azuread/v13/jsii"

	"github.com/cdktf/cdktf-provider-azuread-go/azuread/v13/grouprolemanagementpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotificationsOutputReference interface {
	cdktf.ComplexObject
	AdditionalRecipients() *[]*string
	SetAdditionalRecipients(val *[]*string)
	AdditionalRecipientsInput() *[]*string
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
	DefaultRecipients() interface{}
	SetDefaultRecipients(val interface{})
	DefaultRecipientsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotifications
	SetInternalValue(val *GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotifications)
	NotificationLevel() *string
	SetNotificationLevel(val *string)
	NotificationLevelInput() *string
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
	ResetAdditionalRecipients()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotificationsOutputReference
type jsiiProxy_GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotificationsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotificationsOutputReference) AdditionalRecipients() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalRecipients",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotificationsOutputReference) AdditionalRecipientsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalRecipientsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotificationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotificationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotificationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotificationsOutputReference) DefaultRecipients() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultRecipients",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotificationsOutputReference) DefaultRecipientsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultRecipientsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotificationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotificationsOutputReference) InternalValue() *GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotifications {
	var returns *GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotifications
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotificationsOutputReference) NotificationLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotificationsOutputReference) NotificationLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotificationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotificationsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotificationsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotificationsOutputReference {
	_init_.Initialize()

	if err := validateNewGroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotificationsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotificationsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azuread.groupRoleManagementPolicy.GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotificationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotificationsOutputReference_Override(g GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotificationsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azuread.groupRoleManagementPolicy.GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotificationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotificationsOutputReference)SetAdditionalRecipients(val *[]*string) {
	if err := j.validateSetAdditionalRecipientsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalRecipients",
		val,
	)
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotificationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotificationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotificationsOutputReference)SetDefaultRecipients(val interface{}) {
	if err := j.validateSetDefaultRecipientsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultRecipients",
		val,
	)
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotificationsOutputReference)SetInternalValue(val *GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotifications) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotificationsOutputReference)SetNotificationLevel(val *string) {
	if err := j.validateSetNotificationLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notificationLevel",
		val,
	)
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotificationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotificationsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotificationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotificationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotificationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotificationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotificationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotificationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotificationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotificationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotificationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotificationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotificationsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotificationsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotificationsOutputReference) ResetAdditionalRecipients() {
	_jsii_.InvokeVoid(
		g,
		"resetAdditionalRecipients",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotificationsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesActiveAssignmentsAssigneeNotificationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


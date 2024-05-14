// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package grouprolemanagementpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azuread-go/azuread/v12/jsii"

	"github.com/cdktf/cdktf-provider-azuread-go/azuread/v12/grouprolemanagementpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference interface {
	cdktf.ComplexObject
	AdminNotifications() GroupRoleManagementPolicyNotificationRulesEligibleActivationsAdminNotificationsOutputReference
	AdminNotificationsInput() *GroupRoleManagementPolicyNotificationRulesEligibleActivationsAdminNotifications
	ApproverNotifications() GroupRoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference
	ApproverNotificationsInput() *GroupRoleManagementPolicyNotificationRulesEligibleActivationsApproverNotifications
	AssigneeNotifications() GroupRoleManagementPolicyNotificationRulesEligibleActivationsAssigneeNotificationsOutputReference
	AssigneeNotificationsInput() *GroupRoleManagementPolicyNotificationRulesEligibleActivationsAssigneeNotifications
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
	InternalValue() *GroupRoleManagementPolicyNotificationRulesEligibleActivations
	SetInternalValue(val *GroupRoleManagementPolicyNotificationRulesEligibleActivations)
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
	PutAdminNotifications(value *GroupRoleManagementPolicyNotificationRulesEligibleActivationsAdminNotifications)
	PutApproverNotifications(value *GroupRoleManagementPolicyNotificationRulesEligibleActivationsApproverNotifications)
	PutAssigneeNotifications(value *GroupRoleManagementPolicyNotificationRulesEligibleActivationsAssigneeNotifications)
	ResetAdminNotifications()
	ResetApproverNotifications()
	ResetAssigneeNotifications()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference
type jsiiProxy_GroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) AdminNotifications() GroupRoleManagementPolicyNotificationRulesEligibleActivationsAdminNotificationsOutputReference {
	var returns GroupRoleManagementPolicyNotificationRulesEligibleActivationsAdminNotificationsOutputReference
	_jsii_.Get(
		j,
		"adminNotifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) AdminNotificationsInput() *GroupRoleManagementPolicyNotificationRulesEligibleActivationsAdminNotifications {
	var returns *GroupRoleManagementPolicyNotificationRulesEligibleActivationsAdminNotifications
	_jsii_.Get(
		j,
		"adminNotificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) ApproverNotifications() GroupRoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference {
	var returns GroupRoleManagementPolicyNotificationRulesEligibleActivationsApproverNotificationsOutputReference
	_jsii_.Get(
		j,
		"approverNotifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) ApproverNotificationsInput() *GroupRoleManagementPolicyNotificationRulesEligibleActivationsApproverNotifications {
	var returns *GroupRoleManagementPolicyNotificationRulesEligibleActivationsApproverNotifications
	_jsii_.Get(
		j,
		"approverNotificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) AssigneeNotifications() GroupRoleManagementPolicyNotificationRulesEligibleActivationsAssigneeNotificationsOutputReference {
	var returns GroupRoleManagementPolicyNotificationRulesEligibleActivationsAssigneeNotificationsOutputReference
	_jsii_.Get(
		j,
		"assigneeNotifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) AssigneeNotificationsInput() *GroupRoleManagementPolicyNotificationRulesEligibleActivationsAssigneeNotifications {
	var returns *GroupRoleManagementPolicyNotificationRulesEligibleActivationsAssigneeNotifications
	_jsii_.Get(
		j,
		"assigneeNotificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) InternalValue() *GroupRoleManagementPolicyNotificationRulesEligibleActivations {
	var returns *GroupRoleManagementPolicyNotificationRulesEligibleActivations
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference {
	_init_.Initialize()

	if err := validateNewGroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azuread.groupRoleManagementPolicy.GroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference_Override(g GroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azuread.groupRoleManagementPolicy.GroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference)SetInternalValue(val *GroupRoleManagementPolicyNotificationRulesEligibleActivations) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) PutAdminNotifications(value *GroupRoleManagementPolicyNotificationRulesEligibleActivationsAdminNotifications) {
	if err := g.validatePutAdminNotificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAdminNotifications",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) PutApproverNotifications(value *GroupRoleManagementPolicyNotificationRulesEligibleActivationsApproverNotifications) {
	if err := g.validatePutApproverNotificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putApproverNotifications",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) PutAssigneeNotifications(value *GroupRoleManagementPolicyNotificationRulesEligibleActivationsAssigneeNotifications) {
	if err := g.validatePutAssigneeNotificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAssigneeNotifications",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) ResetAdminNotifications() {
	_jsii_.InvokeVoid(
		g,
		"resetAdminNotifications",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) ResetApproverNotifications() {
	_jsii_.InvokeVoid(
		g,
		"resetApproverNotifications",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) ResetAssigneeNotifications() {
	_jsii_.InvokeVoid(
		g,
		"resetAssigneeNotifications",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GroupRoleManagementPolicyNotificationRulesEligibleActivationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

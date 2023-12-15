// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package groupmember

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GroupMember) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (g *jsiiProxy_GroupMember) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (g *jsiiProxy_GroupMember) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GroupMember) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GroupMember) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GroupMember) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GroupMember) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GroupMember) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GroupMember) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GroupMember) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GroupMember) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GroupMember) validateImportFromParameters(id *string) error {
	return nil
}

func (g *jsiiProxy_GroupMember) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GroupMember) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (g *jsiiProxy_GroupMember) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (g *jsiiProxy_GroupMember) validateMoveToIdParameters(id *string) error {
	return nil
}

func (g *jsiiProxy_GroupMember) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (g *jsiiProxy_GroupMember) validatePutTimeoutsParameters(value *GroupMemberTimeouts) error {
	return nil
}

func validateGroupMember_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateGroupMember_IsConstructParameters(x interface{}) error {
	return nil
}

func validateGroupMember_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateGroupMember_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_GroupMember) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GroupMember) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GroupMember) validateSetGroupObjectIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GroupMember) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GroupMember) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_GroupMember) validateSetMemberObjectIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GroupMember) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func validateNewGroupMemberParameters(scope constructs.Construct, id *string, config *GroupMemberConfig) error {
	return nil
}


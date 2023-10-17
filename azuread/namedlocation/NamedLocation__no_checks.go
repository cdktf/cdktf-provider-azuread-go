// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package namedlocation

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NamedLocation) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (n *jsiiProxy_NamedLocation) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (n *jsiiProxy_NamedLocation) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_NamedLocation) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_NamedLocation) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_NamedLocation) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_NamedLocation) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_NamedLocation) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_NamedLocation) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_NamedLocation) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_NamedLocation) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_NamedLocation) validateImportFromParameters(id *string) error {
	return nil
}

func (n *jsiiProxy_NamedLocation) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_NamedLocation) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (n *jsiiProxy_NamedLocation) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (n *jsiiProxy_NamedLocation) validatePutCountryParameters(value *NamedLocationCountry) error {
	return nil
}

func (n *jsiiProxy_NamedLocation) validatePutIpParameters(value *NamedLocationIp) error {
	return nil
}

func (n *jsiiProxy_NamedLocation) validatePutTimeoutsParameters(value *NamedLocationTimeouts) error {
	return nil
}

func validateNamedLocation_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateNamedLocation_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNamedLocation_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateNamedLocation_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_NamedLocation) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NamedLocation) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NamedLocation) validateSetDisplayNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NamedLocation) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NamedLocation) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_NamedLocation) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func validateNewNamedLocationParameters(scope constructs.Construct, id *string, config *NamedLocationConfig) error {
	return nil
}


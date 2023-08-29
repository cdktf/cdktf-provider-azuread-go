// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataazureadapplication

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAzureadApplicationAppRolesList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAzureadApplicationAppRolesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAzureadApplicationAppRolesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAzureadApplicationAppRolesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAzureadApplicationAppRolesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAzureadApplicationAppRolesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


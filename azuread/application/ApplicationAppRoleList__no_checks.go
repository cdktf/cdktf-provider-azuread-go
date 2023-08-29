// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package application

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApplicationAppRoleList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_ApplicationAppRoleList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ApplicationAppRoleList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationAppRoleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationAppRoleList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ApplicationAppRoleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewApplicationAppRoleListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


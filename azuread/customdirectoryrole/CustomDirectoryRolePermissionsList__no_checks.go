// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package customdirectoryrole

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CustomDirectoryRolePermissionsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CustomDirectoryRolePermissionsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CustomDirectoryRolePermissionsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CustomDirectoryRolePermissionsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CustomDirectoryRolePermissionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CustomDirectoryRolePermissionsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CustomDirectoryRolePermissionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCustomDirectoryRolePermissionsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


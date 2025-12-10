// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataazureadusers

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAzureadUsersUsersList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataAzureadUsersUsersList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAzureadUsersUsersList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAzureadUsersUsersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAzureadUsersUsersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAzureadUsersUsersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAzureadUsersUsersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


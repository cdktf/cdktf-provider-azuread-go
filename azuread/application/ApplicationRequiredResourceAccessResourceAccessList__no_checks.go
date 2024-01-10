// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package application

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApplicationRequiredResourceAccessResourceAccessList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_ApplicationRequiredResourceAccessResourceAccessList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_ApplicationRequiredResourceAccessResourceAccessList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ApplicationRequiredResourceAccessResourceAccessList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationRequiredResourceAccessResourceAccessList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationRequiredResourceAccessResourceAccessList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ApplicationRequiredResourceAccessResourceAccessList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewApplicationRequiredResourceAccessResourceAccessListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


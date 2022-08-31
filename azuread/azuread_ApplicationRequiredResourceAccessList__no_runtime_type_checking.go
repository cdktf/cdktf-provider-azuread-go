//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt azuread Provider for Terraform CDK (cdktf)
package azuread

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApplicationRequiredResourceAccessList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_ApplicationRequiredResourceAccessList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ApplicationRequiredResourceAccessList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationRequiredResourceAccessList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationRequiredResourceAccessList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ApplicationRequiredResourceAccessList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewApplicationRequiredResourceAccessListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt azuread Provider for Terraform CDK (cdktf)
package azuread

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAzureadApplicationRequiredResourceAccessList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAzureadApplicationRequiredResourceAccessList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAzureadApplicationRequiredResourceAccessList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAzureadApplicationRequiredResourceAccessList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAzureadApplicationRequiredResourceAccessList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAzureadApplicationRequiredResourceAccessListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


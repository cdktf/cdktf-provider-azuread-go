//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt azuread Provider for Terraform CDK (cdktf)
package azuread

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAzureadServicePrincipalsServicePrincipalsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAzureadServicePrincipalsServicePrincipalsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAzureadServicePrincipalsServicePrincipalsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAzureadServicePrincipalsServicePrincipalsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAzureadServicePrincipalsServicePrincipalsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAzureadServicePrincipalsServicePrincipalsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

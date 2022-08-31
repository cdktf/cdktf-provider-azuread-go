//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt azuread Provider for Terraform CDK (cdktf)
package azuread

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAzureadApplicationOptionalClaimsIdTokenList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAzureadApplicationOptionalClaimsIdTokenList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAzureadApplicationOptionalClaimsIdTokenList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAzureadApplicationOptionalClaimsIdTokenList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAzureadApplicationOptionalClaimsIdTokenList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAzureadApplicationOptionalClaimsIdTokenListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


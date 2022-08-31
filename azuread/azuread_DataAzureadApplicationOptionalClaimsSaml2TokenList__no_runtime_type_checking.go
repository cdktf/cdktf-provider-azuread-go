//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt azuread Provider for Terraform CDK (cdktf)
package azuread

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAzureadApplicationOptionalClaimsSaml2TokenList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAzureadApplicationOptionalClaimsSaml2TokenList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAzureadApplicationOptionalClaimsSaml2TokenList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAzureadApplicationOptionalClaimsSaml2TokenList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAzureadApplicationOptionalClaimsSaml2TokenList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAzureadApplicationOptionalClaimsSaml2TokenListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


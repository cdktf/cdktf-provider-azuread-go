//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package application

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApplicationOptionalClaimsSaml2TokenList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_ApplicationOptionalClaimsSaml2TokenList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ApplicationOptionalClaimsSaml2TokenList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationOptionalClaimsSaml2TokenList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationOptionalClaimsSaml2TokenList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ApplicationOptionalClaimsSaml2TokenList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewApplicationOptionalClaimsSaml2TokenListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

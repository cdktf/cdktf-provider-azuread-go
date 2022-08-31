//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt azuread Provider for Terraform CDK (cdktf)
package azuread

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAzureadServicePrincipalOauth2PermissionScopesList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAzureadServicePrincipalOauth2PermissionScopesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAzureadServicePrincipalOauth2PermissionScopesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAzureadServicePrincipalOauth2PermissionScopesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAzureadServicePrincipalOauth2PermissionScopesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAzureadServicePrincipalOauth2PermissionScopesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


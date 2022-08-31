//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt azuread Provider for Terraform CDK (cdktf)
package azuread

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServicePrincipalAppRolesList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ServicePrincipalAppRolesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ServicePrincipalAppRolesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ServicePrincipalAppRolesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ServicePrincipalAppRolesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewServicePrincipalAppRolesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


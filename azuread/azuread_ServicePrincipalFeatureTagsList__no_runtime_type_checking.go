//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt azuread Provider for Terraform CDK (cdktf)
package azuread

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServicePrincipalFeatureTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ServicePrincipalFeatureTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ServicePrincipalFeatureTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ServicePrincipalFeatureTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ServicePrincipalFeatureTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ServicePrincipalFeatureTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewServicePrincipalFeatureTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


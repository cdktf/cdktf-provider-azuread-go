// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package application

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApplicationOptionalClaimsAccessTokenList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_ApplicationOptionalClaimsAccessTokenList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_ApplicationOptionalClaimsAccessTokenList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ApplicationOptionalClaimsAccessTokenList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationOptionalClaimsAccessTokenList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationOptionalClaimsAccessTokenList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ApplicationOptionalClaimsAccessTokenList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewApplicationOptionalClaimsAccessTokenListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


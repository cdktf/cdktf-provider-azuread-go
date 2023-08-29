// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AzureadProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (a *jsiiProxy_AzureadProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateAzureadProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAzureadProvider_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateAzureadProvider_IsTerraformProviderParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_AzureadProvider) validateSetDisableTerraformPartnerIdParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AzureadProvider) validateSetUseCliParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AzureadProvider) validateSetUseMsiParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AzureadProvider) validateSetUseOidcParameters(val interface{}) error {
	return nil
}

func validateNewAzureadProviderParameters(scope constructs.Construct, id *string, config *AzureadProviderConfig) error {
	return nil
}


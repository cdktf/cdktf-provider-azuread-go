// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package synchronizationsecret

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SynchronizationSecretCredentialList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SynchronizationSecretCredentialList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SynchronizationSecretCredentialList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SynchronizationSecretCredentialList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SynchronizationSecretCredentialList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SynchronizationSecretCredentialList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SynchronizationSecretCredentialList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSynchronizationSecretCredentialListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


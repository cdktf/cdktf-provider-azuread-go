// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationfederatedidentitycredential


type ApplicationFederatedIdentityCredentialTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.3.0/docs/resources/application_federated_identity_credential#create ApplicationFederatedIdentityCredential#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.3.0/docs/resources/application_federated_identity_credential#delete ApplicationFederatedIdentityCredential#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.3.0/docs/resources/application_federated_identity_credential#read ApplicationFederatedIdentityCredential#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.3.0/docs/resources/application_federated_identity_credential#update ApplicationFederatedIdentityCredential#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceprincipaltokensigningcertificate


type ServicePrincipalTokenSigningCertificateTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/service_principal_token_signing_certificate#create ServicePrincipalTokenSigningCertificate#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/service_principal_token_signing_certificate#delete ServicePrincipalTokenSigningCertificate#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/service_principal_token_signing_certificate#read ServicePrincipalTokenSigningCertificate#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.51.0/docs/resources/service_principal_token_signing_certificate#update ServicePrincipalTokenSigningCertificate#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


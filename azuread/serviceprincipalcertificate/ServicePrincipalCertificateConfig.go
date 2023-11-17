// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceprincipalcertificate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServicePrincipalCertificateConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The object ID of the service principal for which this certificate should be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.46.0/docs/resources/service_principal_certificate#service_principal_id ServicePrincipalCertificate#service_principal_id}
	ServicePrincipalId *string `field:"required" json:"servicePrincipalId" yaml:"servicePrincipalId"`
	// The certificate data, which can be PEM encoded, base64 encoded DER or hexadecimal encoded DER.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.46.0/docs/resources/service_principal_certificate#value ServicePrincipalCertificate#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// Specifies the encoding used for the supplied certificate data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.46.0/docs/resources/service_principal_certificate#encoding ServicePrincipalCertificate#encoding}
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
	// The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.46.0/docs/resources/service_principal_certificate#end_date ServicePrincipalCertificate#end_date}
	EndDate *string `field:"optional" json:"endDate" yaml:"endDate"`
	// A relative duration for which the certificate is valid until, for example `240h` (10 days) or `2400h30m`.
	//
	// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h"
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.46.0/docs/resources/service_principal_certificate#end_date_relative ServicePrincipalCertificate#end_date_relative}
	EndDateRelative *string `field:"optional" json:"endDateRelative" yaml:"endDateRelative"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.46.0/docs/resources/service_principal_certificate#id ServicePrincipalCertificate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A UUID used to uniquely identify this certificate. If not specified a UUID will be automatically generated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.46.0/docs/resources/service_principal_certificate#key_id ServicePrincipalCertificate#key_id}
	KeyId *string `field:"optional" json:"keyId" yaml:"keyId"`
	// The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.46.0/docs/resources/service_principal_certificate#start_date ServicePrincipalCertificate#start_date}
	StartDate *string `field:"optional" json:"startDate" yaml:"startDate"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.46.0/docs/resources/service_principal_certificate#timeouts ServicePrincipalCertificate#timeouts}
	Timeouts *ServicePrincipalCertificateTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The type of key/certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.46.0/docs/resources/service_principal_certificate#type ServicePrincipalCertificate#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}


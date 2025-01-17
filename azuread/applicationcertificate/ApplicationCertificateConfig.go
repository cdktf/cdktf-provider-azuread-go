// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationcertificate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApplicationCertificateConfig struct {
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
	// The resource ID of the application for which this certificate should be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.1.0/docs/resources/application_certificate#application_id ApplicationCertificate#application_id}
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The certificate data, which can be PEM encoded, base64 encoded DER or hexadecimal encoded DER.
	//
	// See also the `encoding` argument
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.1.0/docs/resources/application_certificate#value ApplicationCertificate#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// Specifies the encoding used for the supplied certificate data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.1.0/docs/resources/application_certificate#encoding ApplicationCertificate#encoding}
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
	// The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If omitted, the API will decide a suitable expiry date, which is typically around 2 years from the start date.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.1.0/docs/resources/application_certificate#end_date ApplicationCertificate#end_date}
	EndDate *string `field:"optional" json:"endDate" yaml:"endDate"`
	// A relative duration for which the certificate is valid until, for example `240h` (10 days) or `2400h30m`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.1.0/docs/resources/application_certificate#end_date_relative ApplicationCertificate#end_date_relative}
	EndDateRelative *string `field:"optional" json:"endDateRelative" yaml:"endDateRelative"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.1.0/docs/resources/application_certificate#id ApplicationCertificate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A UUID used to uniquely identify this certificate. If omitted, a random UUID will be automatically generated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.1.0/docs/resources/application_certificate#key_id ApplicationCertificate#key_id}
	KeyId *string `field:"optional" json:"keyId" yaml:"keyId"`
	// The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date and time are use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.1.0/docs/resources/application_certificate#start_date ApplicationCertificate#start_date}
	StartDate *string `field:"optional" json:"startDate" yaml:"startDate"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.1.0/docs/resources/application_certificate#timeouts ApplicationCertificate#timeouts}
	Timeouts *ApplicationCertificateTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The type of key/certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.1.0/docs/resources/application_certificate#type ApplicationCertificate#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}


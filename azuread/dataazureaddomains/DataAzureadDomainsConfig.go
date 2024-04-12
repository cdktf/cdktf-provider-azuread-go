// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataazureaddomains

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAzureadDomainsConfig struct {
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
	// Set to `true` to only return domains whose DNS is managed by Microsoft 365.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.48.0/docs/data-sources/domains#admin_managed DataAzureadDomains#admin_managed}
	AdminManaged interface{} `field:"optional" json:"adminManaged" yaml:"adminManaged"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.48.0/docs/data-sources/domains#id DataAzureadDomains#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Set to `true` if unverified Azure AD domains should be included.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.48.0/docs/data-sources/domains#include_unverified DataAzureadDomains#include_unverified}
	IncludeUnverified interface{} `field:"optional" json:"includeUnverified" yaml:"includeUnverified"`
	// Set to `true` to only return the default domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.48.0/docs/data-sources/domains#only_default DataAzureadDomains#only_default}
	OnlyDefault interface{} `field:"optional" json:"onlyDefault" yaml:"onlyDefault"`
	// Set to `true` to only return the initial domain, which is your primary Azure Active Directory tenant domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.48.0/docs/data-sources/domains#only_initial DataAzureadDomains#only_initial}
	OnlyInitial interface{} `field:"optional" json:"onlyInitial" yaml:"onlyInitial"`
	// Set to `true` to only return verified root domains. Excludes subdomains and unverified domains.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.48.0/docs/data-sources/domains#only_root DataAzureadDomains#only_root}
	OnlyRoot interface{} `field:"optional" json:"onlyRoot" yaml:"onlyRoot"`
	// A list of supported services that must be supported by a domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.48.0/docs/data-sources/domains#supports_services DataAzureadDomains#supports_services}
	SupportsServices *[]*string `field:"optional" json:"supportsServices" yaml:"supportsServices"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.48.0/docs/data-sources/domains#timeouts DataAzureadDomains#timeouts}
	Timeouts *DataAzureadDomainsTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


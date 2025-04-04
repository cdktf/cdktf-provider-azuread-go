// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package namedlocation


type NamedLocationCountry struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.3.0/docs/resources/named_location#countries_and_regions NamedLocation#countries_and_regions}.
	CountriesAndRegions *[]*string `field:"required" json:"countriesAndRegions" yaml:"countriesAndRegions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.3.0/docs/resources/named_location#country_lookup_method NamedLocation#country_lookup_method}.
	CountryLookupMethod *string `field:"optional" json:"countryLookupMethod" yaml:"countryLookupMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.3.0/docs/resources/named_location#include_unknown_countries_and_regions NamedLocation#include_unknown_countries_and_regions}.
	IncludeUnknownCountriesAndRegions interface{} `field:"optional" json:"includeUnknownCountriesAndRegions" yaml:"includeUnknownCountriesAndRegions"`
}


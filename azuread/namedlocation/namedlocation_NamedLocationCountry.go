package namedlocation


type NamedLocationCountry struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/named_location#countries_and_regions NamedLocation#countries_and_regions}.
	CountriesAndRegions *[]*string `field:"required" json:"countriesAndRegions" yaml:"countriesAndRegions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/named_location#include_unknown_countries_and_regions NamedLocation#include_unknown_countries_and_regions}.
	IncludeUnknownCountriesAndRegions interface{} `field:"optional" json:"includeUnknownCountriesAndRegions" yaml:"includeUnknownCountriesAndRegions"`
}


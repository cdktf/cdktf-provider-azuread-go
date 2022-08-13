// Prebuilt azuread Provider for Terraform CDK (cdktf)
package azuread


type NamedLocationIp struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/named_location#ip_ranges NamedLocation#ip_ranges}.
	IpRanges *[]*string `field:"required" json:"ipRanges" yaml:"ipRanges"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/named_location#trusted NamedLocation#trusted}.
	Trusted interface{} `field:"optional" json:"trusted" yaml:"trusted"`
}


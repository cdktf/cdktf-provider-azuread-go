package accesspackageassignmentpolicy


type AccessPackageAssignmentPolicyRequestorSettings struct {
	// requestor block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.38.0/docs/resources/access_package_assignment_policy#requestor AccessPackageAssignmentPolicy#requestor}
	Requestor interface{} `field:"optional" json:"requestor" yaml:"requestor"`
	// Whether to accept requests now, when disabled, no new requests can be made using this policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.38.0/docs/resources/access_package_assignment_policy#requests_accepted AccessPackageAssignmentPolicy#requests_accepted}
	RequestsAccepted interface{} `field:"optional" json:"requestsAccepted" yaml:"requestsAccepted"`
	// Specify the scopes of the requestors.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.38.0/docs/resources/access_package_assignment_policy#scope_type AccessPackageAssignmentPolicy#scope_type}
	ScopeType *string `field:"optional" json:"scopeType" yaml:"scopeType"`
}


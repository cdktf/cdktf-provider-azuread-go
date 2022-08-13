// Prebuilt azuread Provider for Terraform CDK (cdktf)
package azuread


type ServicePrincipalDelegatedPermissionGrantTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/service_principal_delegated_permission_grant#create ServicePrincipalDelegatedPermissionGrant#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/service_principal_delegated_permission_grant#delete ServicePrincipalDelegatedPermissionGrant#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/service_principal_delegated_permission_grant#read ServicePrincipalDelegatedPermissionGrant#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/service_principal_delegated_permission_grant#update ServicePrincipalDelegatedPermissionGrant#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


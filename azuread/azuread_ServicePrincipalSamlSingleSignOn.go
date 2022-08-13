// Prebuilt azuread Provider for Terraform CDK (cdktf)
package azuread


type ServicePrincipalSamlSingleSignOn struct {
	// The relative URI the service provider would redirect to after completion of the single sign-on flow.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/service_principal#relay_state ServicePrincipal#relay_state}
	RelayState *string `field:"optional" json:"relayState" yaml:"relayState"`
}


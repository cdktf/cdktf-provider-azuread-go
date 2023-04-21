package serviceprincipal


type ServicePrincipalSamlSingleSignOn struct {
	// The relative URI the service provider would redirect to after completion of the single sign-on flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.37.2/docs/resources/service_principal#relay_state ServicePrincipal#relay_state}
	RelayState *string `field:"optional" json:"relayState" yaml:"relayState"`
}


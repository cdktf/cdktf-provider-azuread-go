package application


type ApplicationRequiredResourceAccess struct {
	// resource_access block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/application#resource_access Application#resource_access}
	ResourceAccess interface{} `field:"required" json:"resourceAccess" yaml:"resourceAccess"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/application#resource_app_id Application#resource_app_id}.
	ResourceAppId *string `field:"required" json:"resourceAppId" yaml:"resourceAppId"`
}


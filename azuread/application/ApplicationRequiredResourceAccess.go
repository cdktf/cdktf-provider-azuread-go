// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package application


type ApplicationRequiredResourceAccess struct {
	// resource_access block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.2.0/docs/resources/application#resource_access Application#resource_access}
	ResourceAccess interface{} `field:"required" json:"resourceAccess" yaml:"resourceAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.2.0/docs/resources/application#resource_app_id Application#resource_app_id}.
	ResourceAppId *string `field:"required" json:"resourceAppId" yaml:"resourceAppId"`
}


package customdirectoryrole


type CustomDirectoryRolePermissions struct {
	// Set of tasks that can be performed on a resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/r/custom_directory_role#allowed_resource_actions CustomDirectoryRole#allowed_resource_actions}
	AllowedResourceActions *[]*string `field:"required" json:"allowedResourceActions" yaml:"allowedResourceActions"`
}


package dataazureadserviceprincipal

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAzureadServicePrincipalConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The application ID (client ID) of the application associated with this service principal.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/d/service_principal#application_id DataAzureadServicePrincipal#application_id}
	ApplicationId *string `field:"optional" json:"applicationId" yaml:"applicationId"`
	// The display name of the application associated with this service principal.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/d/service_principal#display_name DataAzureadServicePrincipal#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/d/service_principal#id DataAzureadServicePrincipal#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The object ID of the service principal.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/d/service_principal#object_id DataAzureadServicePrincipal#object_id}
	ObjectId *string `field:"optional" json:"objectId" yaml:"objectId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/d/service_principal#timeouts DataAzureadServicePrincipal#timeouts}
	Timeouts *DataAzureadServicePrincipalTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

package dataazureadapplicationtemplate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAzureadApplicationTemplateConfig struct {
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
	// The display name for the application template.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/d/application_template#display_name DataAzureadApplicationTemplate#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/d/application_template#id DataAzureadApplicationTemplate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The application template's ID.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/d/application_template#template_id DataAzureadApplicationTemplate#template_id}
	TemplateId *string `field:"optional" json:"templateId" yaml:"templateId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azuread/d/application_template#timeouts DataAzureadApplicationTemplate#timeouts}
	Timeouts *DataAzureadApplicationTemplateTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationfromtemplate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApplicationFromTemplateConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
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
	// The display name for the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/application_from_template#display_name ApplicationFromTemplate#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The UUID of the template to instantiate for this application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/application_from_template#template_id ApplicationFromTemplate#template_id}
	TemplateId *string `field:"required" json:"templateId" yaml:"templateId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/application_from_template#id ApplicationFromTemplate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/3.4.0/docs/resources/application_from_template#timeouts ApplicationFromTemplate#timeouts}
	Timeouts *ApplicationFromTemplateTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


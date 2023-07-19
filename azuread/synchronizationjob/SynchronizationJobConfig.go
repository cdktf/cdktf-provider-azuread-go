package synchronizationjob

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SynchronizationJobConfig struct {
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
	// The object ID of the service principal for which this synchronization job should be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.40.0/docs/resources/synchronization_job#service_principal_id SynchronizationJob#service_principal_id}
	ServicePrincipalId *string `field:"required" json:"servicePrincipalId" yaml:"servicePrincipalId"`
	// Identifier of the synchronization template this job is based on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.40.0/docs/resources/synchronization_job#template_id SynchronizationJob#template_id}
	TemplateId *string `field:"required" json:"templateId" yaml:"templateId"`
	// Whether or not the synchronization job is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.40.0/docs/resources/synchronization_job#enabled SynchronizationJob#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.40.0/docs/resources/synchronization_job#id SynchronizationJob#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.40.0/docs/resources/synchronization_job#timeouts SynchronizationJob#timeouts}
	Timeouts *SynchronizationJobTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


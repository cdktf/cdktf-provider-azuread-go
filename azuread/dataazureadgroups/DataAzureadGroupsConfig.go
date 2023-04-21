package dataazureadgroups

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAzureadGroupsConfig struct {
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
	// Common display name prefix of the groups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.37.2/docs/data-sources/groups#display_name_prefix DataAzureadGroups#display_name_prefix}
	DisplayNamePrefix *string `field:"optional" json:"displayNamePrefix" yaml:"displayNamePrefix"`
	// The display names of the groups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.37.2/docs/data-sources/groups#display_names DataAzureadGroups#display_names}
	DisplayNames *[]*string `field:"optional" json:"displayNames" yaml:"displayNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.37.2/docs/data-sources/groups#id DataAzureadGroups#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Ignore missing groups and return groups that were found.
	//
	// The data source will still fail if no groups are found
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.37.2/docs/data-sources/groups#ignore_missing DataAzureadGroups#ignore_missing}
	IgnoreMissing interface{} `field:"optional" json:"ignoreMissing" yaml:"ignoreMissing"`
	// Whether the groups are mail-enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.37.2/docs/data-sources/groups#mail_enabled DataAzureadGroups#mail_enabled}
	MailEnabled interface{} `field:"optional" json:"mailEnabled" yaml:"mailEnabled"`
	// The object IDs of the groups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.37.2/docs/data-sources/groups#object_ids DataAzureadGroups#object_ids}
	ObjectIds *[]*string `field:"optional" json:"objectIds" yaml:"objectIds"`
	// Retrieve all groups with no filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.37.2/docs/data-sources/groups#return_all DataAzureadGroups#return_all}
	ReturnAll interface{} `field:"optional" json:"returnAll" yaml:"returnAll"`
	// Whether the groups are security-enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.37.2/docs/data-sources/groups#security_enabled DataAzureadGroups#security_enabled}
	SecurityEnabled interface{} `field:"optional" json:"securityEnabled" yaml:"securityEnabled"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.37.2/docs/data-sources/groups#timeouts DataAzureadGroups#timeouts}
	Timeouts *DataAzureadGroupsTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


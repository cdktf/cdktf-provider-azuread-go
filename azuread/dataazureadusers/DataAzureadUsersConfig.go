package dataazureadusers

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAzureadUsersConfig struct {
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
	// The employee identifier assigned to the user by the organisation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.40.0/docs/data-sources/users#employee_ids DataAzureadUsers#employee_ids}
	EmployeeIds *[]*string `field:"optional" json:"employeeIds" yaml:"employeeIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.40.0/docs/data-sources/users#id DataAzureadUsers#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Ignore missing users and return users that were found.
	//
	// The data source will still fail if no users are found
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.40.0/docs/data-sources/users#ignore_missing DataAzureadUsers#ignore_missing}
	IgnoreMissing interface{} `field:"optional" json:"ignoreMissing" yaml:"ignoreMissing"`
	// The email aliases of the users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.40.0/docs/data-sources/users#mail_nicknames DataAzureadUsers#mail_nicknames}
	MailNicknames *[]*string `field:"optional" json:"mailNicknames" yaml:"mailNicknames"`
	// The object IDs of the users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.40.0/docs/data-sources/users#object_ids DataAzureadUsers#object_ids}
	ObjectIds *[]*string `field:"optional" json:"objectIds" yaml:"objectIds"`
	// Fetch all users with no filter and return all that were found.
	//
	// The data source will still fail if no users are found.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.40.0/docs/data-sources/users#return_all DataAzureadUsers#return_all}
	ReturnAll interface{} `field:"optional" json:"returnAll" yaml:"returnAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.40.0/docs/data-sources/users#timeouts DataAzureadUsers#timeouts}
	Timeouts *DataAzureadUsersTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The user principal names (UPNs) of the users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.40.0/docs/data-sources/users#user_principal_names DataAzureadUsers#user_principal_names}
	UserPrincipalNames *[]*string `field:"optional" json:"userPrincipalNames" yaml:"userPrincipalNames"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package user

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type UserConfig struct {
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
	// The name to display in the address book for the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/user#display_name User#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The user principal name (UPN) of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/user#user_principal_name User#user_principal_name}
	UserPrincipalName *string `field:"required" json:"userPrincipalName" yaml:"userPrincipalName"`
	// Whether or not the account should be enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/user#account_enabled User#account_enabled}
	AccountEnabled interface{} `field:"optional" json:"accountEnabled" yaml:"accountEnabled"`
	// The age group of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/user#age_group User#age_group}
	AgeGroup *string `field:"optional" json:"ageGroup" yaml:"ageGroup"`
	// The telephone numbers for the user.
	//
	// Only one number can be set for this property. Read-only for users synced with Azure AD Connect
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/user#business_phones User#business_phones}
	BusinessPhones *[]*string `field:"optional" json:"businessPhones" yaml:"businessPhones"`
	// The city in which the user is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/user#city User#city}
	City *string `field:"optional" json:"city" yaml:"city"`
	// The company name which the user is associated.
	//
	// This property can be useful for describing the company that an external user comes from
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/user#company_name User#company_name}
	CompanyName *string `field:"optional" json:"companyName" yaml:"companyName"`
	// Whether consent has been obtained for minors.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/user#consent_provided_for_minor User#consent_provided_for_minor}
	ConsentProvidedForMinor *string `field:"optional" json:"consentProvidedForMinor" yaml:"consentProvidedForMinor"`
	// The cost center associated with the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/user#cost_center User#cost_center}
	CostCenter *string `field:"optional" json:"costCenter" yaml:"costCenter"`
	// The country/region in which the user is located, e.g. `US` or `UK`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/user#country User#country}
	Country *string `field:"optional" json:"country" yaml:"country"`
	// The name for the department in which the user works.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/user#department User#department}
	Department *string `field:"optional" json:"department" yaml:"department"`
	// Whether the users password is exempt from expiring.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/user#disable_password_expiration User#disable_password_expiration}
	DisablePasswordExpiration interface{} `field:"optional" json:"disablePasswordExpiration" yaml:"disablePasswordExpiration"`
	// Whether the user is allowed weaker passwords than the default policy to be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/user#disable_strong_password User#disable_strong_password}
	DisableStrongPassword interface{} `field:"optional" json:"disableStrongPassword" yaml:"disableStrongPassword"`
	// The name of the division in which the user works.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/user#division User#division}
	Division *string `field:"optional" json:"division" yaml:"division"`
	// The employee identifier assigned to the user by the organisation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/user#employee_id User#employee_id}
	EmployeeId *string `field:"optional" json:"employeeId" yaml:"employeeId"`
	// Captures enterprise worker type. For example, Employee, Contractor, Consultant, or Vendor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/user#employee_type User#employee_type}
	EmployeeType *string `field:"optional" json:"employeeType" yaml:"employeeType"`
	// The fax number of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/user#fax_number User#fax_number}
	FaxNumber *string `field:"optional" json:"faxNumber" yaml:"faxNumber"`
	// Whether the user is forced to change the password during the next sign-in.
	//
	// Only takes effect when also changing the password
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/user#force_password_change User#force_password_change}
	ForcePasswordChange interface{} `field:"optional" json:"forcePasswordChange" yaml:"forcePasswordChange"`
	// The given name (first name) of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/user#given_name User#given_name}
	GivenName *string `field:"optional" json:"givenName" yaml:"givenName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/user#id User#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The user’s job title.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/user#job_title User#job_title}
	JobTitle *string `field:"optional" json:"jobTitle" yaml:"jobTitle"`
	// The SMTP address for the user. Cannot be unset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/user#mail User#mail}
	Mail *string `field:"optional" json:"mail" yaml:"mail"`
	// The mail alias for the user. Defaults to the user name part of the user principal name (UPN).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/user#mail_nickname User#mail_nickname}
	MailNickname *string `field:"optional" json:"mailNickname" yaml:"mailNickname"`
	// The object ID of the user's manager.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/user#manager_id User#manager_id}
	ManagerId *string `field:"optional" json:"managerId" yaml:"managerId"`
	// The primary cellular telephone number for the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/user#mobile_phone User#mobile_phone}
	MobilePhone *string `field:"optional" json:"mobilePhone" yaml:"mobilePhone"`
	// The office location in the user's place of business.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/user#office_location User#office_location}
	OfficeLocation *string `field:"optional" json:"officeLocation" yaml:"officeLocation"`
	// The value used to associate an on-premise Active Directory user account with their Azure AD user object.
	//
	// This must be specified if you are using a federated domain for the user's `user_principal_name` property when creating a new user account
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/user#onpremises_immutable_id User#onpremises_immutable_id}
	OnpremisesImmutableId *string `field:"optional" json:"onpremisesImmutableId" yaml:"onpremisesImmutableId"`
	// Additional email addresses for the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/user#other_mails User#other_mails}
	OtherMails *[]*string `field:"optional" json:"otherMails" yaml:"otherMails"`
	// The password for the user.
	//
	// The password must satisfy minimum requirements as specified by the password policy. The maximum length is 256 characters. This property is required when creating a new user
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/user#password User#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The postal code for the user's postal address.
	//
	// The postal code is specific to the user's country/region. In the United States of America, this attribute contains the ZIP code
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/user#postal_code User#postal_code}
	PostalCode *string `field:"optional" json:"postalCode" yaml:"postalCode"`
	// The user's preferred language, in ISO 639-1 notation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/user#preferred_language User#preferred_language}
	PreferredLanguage *string `field:"optional" json:"preferredLanguage" yaml:"preferredLanguage"`
	// Whether or not the Outlook global address list should include this user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/user#show_in_address_list User#show_in_address_list}
	ShowInAddressList interface{} `field:"optional" json:"showInAddressList" yaml:"showInAddressList"`
	// The state or province in the user's address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/user#state User#state}
	State *string `field:"optional" json:"state" yaml:"state"`
	// The street address of the user's place of business.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/user#street_address User#street_address}
	StreetAddress *string `field:"optional" json:"streetAddress" yaml:"streetAddress"`
	// The user's surname (family name or last name).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/user#surname User#surname}
	Surname *string `field:"optional" json:"surname" yaml:"surname"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/user#timeouts User#timeouts}
	Timeouts *UserTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The usage location of the user.
	//
	// Required for users that will be assigned licenses due to legal requirement to check for availability of services in countries. The usage location is a two letter country code (ISO standard 3166). Examples include: `NO`, `JP`, and `GB`. Cannot be reset to null once set
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.44.1/docs/resources/user#usage_location User#usage_location}
	UsageLocation *string `field:"optional" json:"usageLocation" yaml:"usageLocation"`
}


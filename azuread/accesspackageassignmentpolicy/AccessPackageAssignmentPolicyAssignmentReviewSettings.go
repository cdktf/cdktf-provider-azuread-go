package accesspackageassignmentpolicy


type AccessPackageAssignmentPolicyAssignmentReviewSettings struct {
	// Whether to show Show reviewer decision helpers.
	//
	// If enabled, system recommendations based on users' access information will be shown to the reviewers. The reviewer will be recommended to approve the review if the user has signed-in at least once during the last 30 days. The reviewer will be recommended to deny the review if the user has not signed-in during the last 30 days
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.37.2/docs/resources/access_package_assignment_policy#access_recommendation_enabled AccessPackageAssignmentPolicy#access_recommendation_enabled}
	AccessRecommendationEnabled interface{} `field:"optional" json:"accessRecommendationEnabled" yaml:"accessRecommendationEnabled"`
	// What actions the system takes if reviewers don't respond in time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.37.2/docs/resources/access_package_assignment_policy#access_review_timeout_behavior AccessPackageAssignmentPolicy#access_review_timeout_behavior}
	AccessReviewTimeoutBehavior *string `field:"optional" json:"accessReviewTimeoutBehavior" yaml:"accessReviewTimeoutBehavior"`
	// Whether a reviewer need provide a justification for their decision. Justification is visible to other reviewers and the requestor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.37.2/docs/resources/access_package_assignment_policy#approver_justification_required AccessPackageAssignmentPolicy#approver_justification_required}
	ApproverJustificationRequired interface{} `field:"optional" json:"approverJustificationRequired" yaml:"approverJustificationRequired"`
	// How many days each occurrence of the access review series will run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.37.2/docs/resources/access_package_assignment_policy#duration_in_days AccessPackageAssignmentPolicy#duration_in_days}
	DurationInDays *float64 `field:"optional" json:"durationInDays" yaml:"durationInDays"`
	// Whether to enable assignment review.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.37.2/docs/resources/access_package_assignment_policy#enabled AccessPackageAssignmentPolicy#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// reviewer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.37.2/docs/resources/access_package_assignment_policy#reviewer AccessPackageAssignmentPolicy#reviewer}
	Reviewer interface{} `field:"optional" json:"reviewer" yaml:"reviewer"`
	// This will determine how often the access review campaign runs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.37.2/docs/resources/access_package_assignment_policy#review_frequency AccessPackageAssignmentPolicy#review_frequency}
	ReviewFrequency *string `field:"optional" json:"reviewFrequency" yaml:"reviewFrequency"`
	// Self review or specific reviewers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.37.2/docs/resources/access_package_assignment_policy#review_type AccessPackageAssignmentPolicy#review_type}
	ReviewType *string `field:"optional" json:"reviewType" yaml:"reviewType"`
	// This is the date the access review campaign will start on, formatted as an RFC3339 date string in UTC(e.g. 2018-01-01T01:02:03Z), default is now. Once an access review has been created, you cannot update its start date.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azuread/2.37.2/docs/resources/access_package_assignment_policy#starting_on AccessPackageAssignmentPolicy#starting_on}
	StartingOn *string `field:"optional" json:"startingOn" yaml:"startingOn"`
}


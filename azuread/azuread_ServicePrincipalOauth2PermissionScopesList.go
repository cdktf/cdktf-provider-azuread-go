// Prebuilt azuread Provider for Terraform CDK (cdktf)
package azuread

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azuread-go/azuread/v2/jsii"

	"github.com/hashicorp/cdktf-provider-azuread-go/azuread/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServicePrincipalOauth2PermissionScopesList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) ServicePrincipalOauth2PermissionScopesOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServicePrincipalOauth2PermissionScopesList
type jsiiProxy_ServicePrincipalOauth2PermissionScopesList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_ServicePrincipalOauth2PermissionScopesList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipalOauth2PermissionScopesList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipalOauth2PermissionScopesList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipalOauth2PermissionScopesList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipalOauth2PermissionScopesList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewServicePrincipalOauth2PermissionScopesList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ServicePrincipalOauth2PermissionScopesList {
	_init_.Initialize()

	j := jsiiProxy_ServicePrincipalOauth2PermissionScopesList{}

	_jsii_.Create(
		"@cdktf/provider-azuread.ServicePrincipalOauth2PermissionScopesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewServicePrincipalOauth2PermissionScopesList_Override(s ServicePrincipalOauth2PermissionScopesList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azuread.ServicePrincipalOauth2PermissionScopesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		s,
	)
}

func (j *jsiiProxy_ServicePrincipalOauth2PermissionScopesList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipalOauth2PermissionScopesList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ServicePrincipalOauth2PermissionScopesList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (s *jsiiProxy_ServicePrincipalOauth2PermissionScopesList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipalOauth2PermissionScopesList) Get(index *float64) ServicePrincipalOauth2PermissionScopesOutputReference {
	var returns ServicePrincipalOauth2PermissionScopesOutputReference

	_jsii_.Invoke(
		s,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipalOauth2PermissionScopesList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipalOauth2PermissionScopesList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

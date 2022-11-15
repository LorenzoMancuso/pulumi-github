// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetDependabotOrganizationSecrets(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetDependabotOrganizationSecretsResult, error) {
	var rv GetDependabotOrganizationSecretsResult
	err := ctx.Invoke("github:index/getDependabotOrganizationSecrets:getDependabotOrganizationSecrets", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getDependabotOrganizationSecrets.
type GetDependabotOrganizationSecretsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id      string                                   `pulumi:"id"`
	Secrets []GetDependabotOrganizationSecretsSecret `pulumi:"secrets"`
}

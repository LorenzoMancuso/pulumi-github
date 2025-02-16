// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetOrganizationWebhooks(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetOrganizationWebhooksResult, error) {
	var rv GetOrganizationWebhooksResult
	err := ctx.Invoke("github:index/getOrganizationWebhooks:getOrganizationWebhooks", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getOrganizationWebhooks.
type GetOrganizationWebhooksResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id       string                           `pulumi:"id"`
	Webhooks []GetOrganizationWebhooksWebhook `pulumi:"webhooks"`
}

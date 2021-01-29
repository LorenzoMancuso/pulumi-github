// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to retrieve basic information about a GitHub Organization.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-github/sdk/v3/go/github"
// 	"github.com/pulumi/pulumi-github/sdk/v3/go/github/"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := github.GetOrganization(ctx, &github.GetOrganizationArgs{
// 			Name: "github",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetOrganization(ctx *pulumi.Context, args *GetOrganizationArgs, opts ...pulumi.InvokeOption) (*GetOrganizationResult, error) {
	var rv GetOrganizationResult
	err := ctx.Invoke("github:index/getOrganization:getOrganization", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOrganization.
type GetOrganizationArgs struct {
	Name string `pulumi:"name"`
}

// A collection of values returned by getOrganization.
type GetOrganizationResult struct {
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id     string `pulumi:"id"`
	Login  string `pulumi:"login"`
	Name   string `pulumi:"name"`
	NodeId string `pulumi:"nodeId"`
	// The plan name for the organization account
	Plan string `pulumi:"plan"`
}

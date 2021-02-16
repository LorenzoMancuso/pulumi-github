// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to retrieve information about a specific GitHub milestone in a repository.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-github/sdk/v3/go/github"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := github.LookupRepositoryMilestone(ctx, &github.LookupRepositoryMilestoneArgs{
// 			Number:     1,
// 			Owner:      "example-owner",
// 			Repository: "example-repository",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupRepositoryMilestone(ctx *pulumi.Context, args *LookupRepositoryMilestoneArgs, opts ...pulumi.InvokeOption) (*LookupRepositoryMilestoneResult, error) {
	var rv LookupRepositoryMilestoneResult
	err := ctx.Invoke("github:index/getRepositoryMilestone:getRepositoryMilestone", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRepositoryMilestone.
type LookupRepositoryMilestoneArgs struct {
	// The number of the milestone.
	Number int `pulumi:"number"`
	// Owner of the repository.
	Owner string `pulumi:"owner"`
	// Name of the repository to retrieve the milestone from.
	Repository string `pulumi:"repository"`
}

// A collection of values returned by getRepositoryMilestone.
type LookupRepositoryMilestoneResult struct {
	// Description of the milestone.
	Description string `pulumi:"description"`
	// The milestone due date (in ISO-8601 `yyyy-mm-dd` format).
	DueDate string `pulumi:"dueDate"`
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	Number     int    `pulumi:"number"`
	Owner      string `pulumi:"owner"`
	Repository string `pulumi:"repository"`
	// State of the milestone.
	State string `pulumi:"state"`
	// Title of the milestone.
	Title string `pulumi:"title"`
}

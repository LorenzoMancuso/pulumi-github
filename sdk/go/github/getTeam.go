// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to retrieve information about a GitHub team.
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
// 		_, err := github.LookupTeam(ctx, &github.LookupTeamArgs{
// 			Slug: "example",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupTeam(ctx *pulumi.Context, args *LookupTeamArgs, opts ...pulumi.InvokeOption) (*LookupTeamResult, error) {
	var rv LookupTeamResult
	err := ctx.Invoke("github:index/getTeam:getTeam", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTeam.
type LookupTeamArgs struct {
	// The team slug.
	Slug string `pulumi:"slug"`
}

// A collection of values returned by getTeam.
type LookupTeamResult struct {
	// the team's description.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List of team members
	Members []string `pulumi:"members"`
	// the team's full name.
	Name string `pulumi:"name"`
	// the Node ID of the team.
	NodeId string `pulumi:"nodeId"`
	// the team's permission level.
	Permission string `pulumi:"permission"`
	// the team's privacy type.
	Privacy string `pulumi:"privacy"`
	Slug    string `pulumi:"slug"`
}

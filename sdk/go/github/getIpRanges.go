// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to retrieve information about GitHub's IP addresses.
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
// 		_, err := github.GetIpRanges(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetIpRanges(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetIpRangesResult, error) {
	var rv GetIpRangesResult
	err := ctx.Invoke("github:index/getIpRanges:getIpRanges", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getIpRanges.
type GetIpRangesResult struct {
	// An Array of IP addresses in CIDR format specifying the Git servers.
	Gits []string `pulumi:"gits"`
	// An Array of IP addresses in CIDR format specifying the addresses that incoming service hooks will originate from.
	Hooks []string `pulumi:"hooks"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// An Array of IP addresses in CIDR format specifying the A records for GitHub Importer.
	Importers []string `pulumi:"importers"`
	// An Array of IP addresses in CIDR format specifying the A records for GitHub Pages.
	Pages []string `pulumi:"pages"`
}

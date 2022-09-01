// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about GitHub's IP addresses.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-github/sdk/v4/go/github"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := github.GetIpRanges(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
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
	// An array of IP addresses in CIDR format specifying the addresses that incoming requests from GitHub actions will originate from.
	Actions []string `pulumi:"actions"`
	// A subset of the `actions` array that contains IP addresses in IPv4 CIDR format.
	ActionsIpv4s []string `pulumi:"actionsIpv4s"`
	// A subset of the `actions` array that contains IP addresses in IPv6 CIDR format.
	ActionsIpv6s []string `pulumi:"actionsIpv6s"`
	// A subset of the `api` array that contains IP addresses in IPv4 CIDR format.
	ApiIpv4s []string `pulumi:"apiIpv4s"`
	// A subset of the `api` array that contains IP addresses in IPv6 CIDR format.
	ApiIpv6s []string `pulumi:"apiIpv6s"`
	// An Array of IP addresses in CIDR format for the GitHub API.
	Apis []string `pulumi:"apis"`
	// A subset of the `dependabot` array that contains IP addresses in IPv4 CIDR format.
	DependabotIpv4s []string `pulumi:"dependabotIpv4s"`
	// A subset of the `dependabot` array that contains IP addresses in IPv6 CIDR format.
	DependabotIpv6s []string `pulumi:"dependabotIpv6s"`
	// An array of IP addresses in CIDR format specifying the A records for dependabot.
	Dependabots []string `pulumi:"dependabots"`
	// A subset of the `git` array that contains IP addresses in IPv4 CIDR format.
	GitIpv4s []string `pulumi:"gitIpv4s"`
	// A subset of the `git` array that contains IP addresses in IPv6 CIDR format.
	GitIpv6s []string `pulumi:"gitIpv6s"`
	// An Array of IP addresses in CIDR format specifying the Git servers.
	Gits []string `pulumi:"gits"`
	// An Array of IP addresses in CIDR format specifying the addresses that incoming service hooks will originate from.
	Hooks []string `pulumi:"hooks"`
	// A subset of the `hooks` array that contains IP addresses in IPv4 CIDR format.
	HooksIpv4s []string `pulumi:"hooksIpv4s"`
	// A subset of the `hooks` array that contains IP addresses in IPv6 CIDR format.
	HooksIpv6s []string `pulumi:"hooksIpv6s"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A subset of the `importer` array that contains IP addresses in IPv4 CIDR format.
	ImporterIpv4s []string `pulumi:"importerIpv4s"`
	// A subset of the `importer` array that contains IP addresses in IPv6 CIDR format.
	ImporterIpv6s []string `pulumi:"importerIpv6s"`
	// An Array of IP addresses in CIDR format specifying the A records for GitHub Importer.
	Importers []string `pulumi:"importers"`
	// An Array of IP addresses in CIDR format specifying the A records for GitHub Pages.
	Pages []string `pulumi:"pages"`
	// A subset of the `pages` array that contains IP addresses in IPv4 CIDR format.
	PagesIpv4s []string `pulumi:"pagesIpv4s"`
	// A subset of the `pages` array that contains IP addresses in IPv6 CIDR format.
	PagesIpv6s []string `pulumi:"pagesIpv6s"`
	// A subset of the `web` array that contains IP addresses in IPv4 CIDR format.
	WebIpv4s []string `pulumi:"webIpv4s"`
	// A subset of the `web` array that contains IP addresses in IPv6 CIDR format.
	WebIpv6s []string `pulumi:"webIpv6s"`
	// An Array of IP addresses in CIDR format for GitHub Web.
	Webs []string `pulumi:"webs"`
}

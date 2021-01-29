// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to retrieve information about a GitHub release in a specific repository.
//
// ## Example Usage
//
// To retrieve the latest release that is present in a repository:
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
// 		_, err := github.GetRelease(ctx, &github.GetReleaseArgs{
// 			Owner:      "example-owner",
// 			Repository: "example-repository",
// 			RetrieveBy: "latest",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// To retrieve a specific release from a repository based on it's ID:
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
// 		_, err := github.GetRelease(ctx, &github.GetReleaseArgs{
// 			Id:         12345,
// 			Owner:      "example-owner",
// 			Repository: "example-repository",
// 			RetrieveBy: "id",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// Finally, to retrieve a release based on it's tag:
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
// 		opt0 := "v1.0.0"
// 		_, err := github.GetRelease(ctx, &github.GetReleaseArgs{
// 			Owner:      "example-owner",
// 			ReleaseTag: &opt0,
// 			Repository: "example-repository",
// 			RetrieveBy: "tag",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetRelease(ctx *pulumi.Context, args *GetReleaseArgs, opts ...pulumi.InvokeOption) (*GetReleaseResult, error) {
	var rv GetReleaseResult
	err := ctx.Invoke("github:index/getRelease:getRelease", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRelease.
type GetReleaseArgs struct {
	// Owner of the repository.
	Owner string `pulumi:"owner"`
	// ID of the release to retrieve. Must be specified when `retrieveBy` = `id`.
	ReleaseId *int `pulumi:"releaseId"`
	// Tag of the release to retrieve. Must be specified when `retrieveBy` = `tag`.
	ReleaseTag *string `pulumi:"releaseTag"`
	// Name of the repository to retrieve the release from.
	Repository string `pulumi:"repository"`
	// Describes how to fetch the release. Valid values are `id`, `tag`, `latest`.
	RetrieveBy string `pulumi:"retrieveBy"`
}

// A collection of values returned by getRelease.
type GetReleaseResult struct {
	// URL of any associated assets with the release
	AssertsUrl string `pulumi:"assertsUrl"`
	// Contents of the description (body) of a release
	Body string `pulumi:"body"`
	// Date of release creation
	CreatedAt string `pulumi:"createdAt"`
	// (`Boolean`) indicates whether the release is a draft
	Draft bool `pulumi:"draft"`
	// URL directing to detailed information on the release
	HtmlUrl string `pulumi:"htmlUrl"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Name of release
	Name  string `pulumi:"name"`
	Owner string `pulumi:"owner"`
	// (`Boolean`) indicates whether the release is a prerelease
	Prerelease bool `pulumi:"prerelease"`
	// Date of release publishing
	PublishedAt string `pulumi:"publishedAt"`
	// ID of release
	ReleaseId *int `pulumi:"releaseId"`
	// Tag of release
	ReleaseTag *string `pulumi:"releaseTag"`
	Repository string  `pulumi:"repository"`
	RetrieveBy string  `pulumi:"retrieveBy"`
	// Download URL of a specific release in `tar.gz` format
	TarballUrl string `pulumi:"tarballUrl"`
	// Commitish value that determines where the Git release is created from
	TargetCommitish string `pulumi:"targetCommitish"`
	// URL that can be used to upload Assets to the release
	UploadUrl string `pulumi:"uploadUrl"`
	// Base URL of the release
	Url string `pulumi:"url"`
	// Download URL of a specific release in `zip` format
	ZipballUrl string `pulumi:"zipballUrl"`
}

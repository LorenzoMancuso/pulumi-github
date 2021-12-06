// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about a GitHub repository.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-github/sdk/v4/go/github"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "hashicorp/terraform"
// 		_, err := github.LookupRepository(ctx, &GetRepositoryArgs{
// 			FullName: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupRepository(ctx *pulumi.Context, args *LookupRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupRepositoryResult, error) {
	var rv LookupRepositoryResult
	err := ctx.Invoke("github:index/getRepository:getRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRepository.
type LookupRepositoryArgs struct {
	// A description of the repository.
	Description *string `pulumi:"description"`
	// Full name of the repository (in `org/name` format).
	FullName *string `pulumi:"fullName"`
	// URL of a page describing the project.
	HomepageUrl *string `pulumi:"homepageUrl"`
	// The name of the repository.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getRepository.
type LookupRepositoryResult struct {
	// Whether the repository allows auto-merging pull requests.
	AllowAutoMerge bool `pulumi:"allowAutoMerge"`
	// Whether the repository allows merge commits.
	AllowMergeCommit bool `pulumi:"allowMergeCommit"`
	// Whether the repository allows rebase merges.
	AllowRebaseMerge bool `pulumi:"allowRebaseMerge"`
	// Whether the repository allows squash merges.
	AllowSquashMerge bool `pulumi:"allowSquashMerge"`
	// Whether the repository is archived.
	Archived bool `pulumi:"archived"`
	// The list of this repository's branches. Each element of `branches` has the following attributes:
	Branches []GetRepositoryBranch `pulumi:"branches"`
	// The name of the default branch of the repository.
	DefaultBranch string `pulumi:"defaultBranch"`
	// A description of the repository.
	Description *string `pulumi:"description"`
	FullName    string  `pulumi:"fullName"`
	// URL that can be provided to `git clone` to clone the repository anonymously via the git protocol.
	GitCloneUrl string `pulumi:"gitCloneUrl"`
	// Whether the repository has Downloads feature enabled.
	HasDownloads bool `pulumi:"hasDownloads"`
	// Whether the repository has GitHub Issues enabled.
	HasIssues bool `pulumi:"hasIssues"`
	// Whether the repository has the GitHub Projects enabled.
	HasProjects bool `pulumi:"hasProjects"`
	// Whether the repository has the GitHub Wiki enabled.
	HasWiki bool `pulumi:"hasWiki"`
	// URL of a page describing the project.
	HomepageUrl *string `pulumi:"homepageUrl"`
	// URL to the repository on the web.
	HtmlUrl string `pulumi:"htmlUrl"`
	// URL that can be provided to `git clone` to clone the repository via HTTPS.
	HttpCloneUrl string `pulumi:"httpCloneUrl"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Name of the branch.
	Name string `pulumi:"name"`
	// GraphQL global node id for use with v4 API
	NodeId string `pulumi:"nodeId"`
	// The repository's GitHub Pages configuration.
	Pages []GetRepositoryPage `pulumi:"pages"`
	// Whether the repository is private.
	Private bool `pulumi:"private"`
	// GitHub ID for the repository
	RepoId int `pulumi:"repoId"`
	// URL that can be provided to `git clone` to clone the repository via SSH.
	SshCloneUrl string `pulumi:"sshCloneUrl"`
	// URL that can be provided to `svn checkout` to check out the repository via GitHub's Subversion protocol emulation.
	SvnUrl string `pulumi:"svnUrl"`
	// The list of topics of the repository.
	Topics []string `pulumi:"topics"`
	// Whether the repository is public, private or internal.
	Visibility string `pulumi:"visibility"`
}

func LookupRepositoryOutput(ctx *pulumi.Context, args LookupRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRepositoryResult, error) {
			args := v.(LookupRepositoryArgs)
			r, err := LookupRepository(ctx, &args, opts...)
			return *r, err
		}).(LookupRepositoryResultOutput)
}

// A collection of arguments for invoking getRepository.
type LookupRepositoryOutputArgs struct {
	// A description of the repository.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// Full name of the repository (in `org/name` format).
	FullName pulumi.StringPtrInput `pulumi:"fullName"`
	// URL of a page describing the project.
	HomepageUrl pulumi.StringPtrInput `pulumi:"homepageUrl"`
	// The name of the repository.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getRepository.
type LookupRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRepositoryResult)(nil)).Elem()
}

func (o LookupRepositoryResultOutput) ToLookupRepositoryResultOutput() LookupRepositoryResultOutput {
	return o
}

func (o LookupRepositoryResultOutput) ToLookupRepositoryResultOutputWithContext(ctx context.Context) LookupRepositoryResultOutput {
	return o
}

// Whether the repository allows auto-merging pull requests.
func (o LookupRepositoryResultOutput) AllowAutoMerge() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRepositoryResult) bool { return v.AllowAutoMerge }).(pulumi.BoolOutput)
}

// Whether the repository allows merge commits.
func (o LookupRepositoryResultOutput) AllowMergeCommit() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRepositoryResult) bool { return v.AllowMergeCommit }).(pulumi.BoolOutput)
}

// Whether the repository allows rebase merges.
func (o LookupRepositoryResultOutput) AllowRebaseMerge() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRepositoryResult) bool { return v.AllowRebaseMerge }).(pulumi.BoolOutput)
}

// Whether the repository allows squash merges.
func (o LookupRepositoryResultOutput) AllowSquashMerge() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRepositoryResult) bool { return v.AllowSquashMerge }).(pulumi.BoolOutput)
}

// Whether the repository is archived.
func (o LookupRepositoryResultOutput) Archived() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRepositoryResult) bool { return v.Archived }).(pulumi.BoolOutput)
}

// The list of this repository's branches. Each element of `branches` has the following attributes:
func (o LookupRepositoryResultOutput) Branches() GetRepositoryBranchArrayOutput {
	return o.ApplyT(func(v LookupRepositoryResult) []GetRepositoryBranch { return v.Branches }).(GetRepositoryBranchArrayOutput)
}

// The name of the default branch of the repository.
func (o LookupRepositoryResultOutput) DefaultBranch() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryResult) string { return v.DefaultBranch }).(pulumi.StringOutput)
}

// A description of the repository.
func (o LookupRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupRepositoryResultOutput) FullName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryResult) string { return v.FullName }).(pulumi.StringOutput)
}

// URL that can be provided to `git clone` to clone the repository anonymously via the git protocol.
func (o LookupRepositoryResultOutput) GitCloneUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryResult) string { return v.GitCloneUrl }).(pulumi.StringOutput)
}

// Whether the repository has Downloads feature enabled.
func (o LookupRepositoryResultOutput) HasDownloads() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRepositoryResult) bool { return v.HasDownloads }).(pulumi.BoolOutput)
}

// Whether the repository has GitHub Issues enabled.
func (o LookupRepositoryResultOutput) HasIssues() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRepositoryResult) bool { return v.HasIssues }).(pulumi.BoolOutput)
}

// Whether the repository has the GitHub Projects enabled.
func (o LookupRepositoryResultOutput) HasProjects() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRepositoryResult) bool { return v.HasProjects }).(pulumi.BoolOutput)
}

// Whether the repository has the GitHub Wiki enabled.
func (o LookupRepositoryResultOutput) HasWiki() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRepositoryResult) bool { return v.HasWiki }).(pulumi.BoolOutput)
}

// URL of a page describing the project.
func (o LookupRepositoryResultOutput) HomepageUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRepositoryResult) *string { return v.HomepageUrl }).(pulumi.StringPtrOutput)
}

// URL to the repository on the web.
func (o LookupRepositoryResultOutput) HtmlUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryResult) string { return v.HtmlUrl }).(pulumi.StringOutput)
}

// URL that can be provided to `git clone` to clone the repository via HTTPS.
func (o LookupRepositoryResultOutput) HttpCloneUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryResult) string { return v.HttpCloneUrl }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of the branch.
func (o LookupRepositoryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryResult) string { return v.Name }).(pulumi.StringOutput)
}

// GraphQL global node id for use with v4 API
func (o LookupRepositoryResultOutput) NodeId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryResult) string { return v.NodeId }).(pulumi.StringOutput)
}

// The repository's GitHub Pages configuration.
func (o LookupRepositoryResultOutput) Pages() GetRepositoryPageArrayOutput {
	return o.ApplyT(func(v LookupRepositoryResult) []GetRepositoryPage { return v.Pages }).(GetRepositoryPageArrayOutput)
}

// Whether the repository is private.
func (o LookupRepositoryResultOutput) Private() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRepositoryResult) bool { return v.Private }).(pulumi.BoolOutput)
}

// GitHub ID for the repository
func (o LookupRepositoryResultOutput) RepoId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRepositoryResult) int { return v.RepoId }).(pulumi.IntOutput)
}

// URL that can be provided to `git clone` to clone the repository via SSH.
func (o LookupRepositoryResultOutput) SshCloneUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryResult) string { return v.SshCloneUrl }).(pulumi.StringOutput)
}

// URL that can be provided to `svn checkout` to check out the repository via GitHub's Subversion protocol emulation.
func (o LookupRepositoryResultOutput) SvnUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryResult) string { return v.SvnUrl }).(pulumi.StringOutput)
}

// The list of topics of the repository.
func (o LookupRepositoryResultOutput) Topics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRepositoryResult) []string { return v.Topics }).(pulumi.StringArrayOutput)
}

// Whether the repository is public, private or internal.
func (o LookupRepositoryResultOutput) Visibility() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryResult) string { return v.Visibility }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRepositoryResultOutput{})
}

// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRelease(ctx *pulumi.Context, args *LookupReleaseArgs, opts ...pulumi.InvokeOption) (*LookupReleaseResult, error) {
	var rv LookupReleaseResult
	err := ctx.Invoke("github:index/getRelease:getRelease", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRelease.
type LookupReleaseArgs struct {
	Owner      string  `pulumi:"owner"`
	ReleaseId  *int    `pulumi:"releaseId"`
	ReleaseTag *string `pulumi:"releaseTag"`
	Repository string  `pulumi:"repository"`
	RetrieveBy string  `pulumi:"retrieveBy"`
}

// A collection of values returned by getRelease.
type LookupReleaseResult struct {
	// Deprecated: use assets_url instead
	AssertsUrl string            `pulumi:"assertsUrl"`
	Assets     []GetReleaseAsset `pulumi:"assets"`
	AssetsUrl  string            `pulumi:"assetsUrl"`
	Body       string            `pulumi:"body"`
	CreatedAt  string            `pulumi:"createdAt"`
	Draft      bool              `pulumi:"draft"`
	HtmlUrl    string            `pulumi:"htmlUrl"`
	// The provider-assigned unique ID for this managed resource.
	Id              string  `pulumi:"id"`
	Name            string  `pulumi:"name"`
	Owner           string  `pulumi:"owner"`
	Prerelease      bool    `pulumi:"prerelease"`
	PublishedAt     string  `pulumi:"publishedAt"`
	ReleaseId       *int    `pulumi:"releaseId"`
	ReleaseTag      *string `pulumi:"releaseTag"`
	Repository      string  `pulumi:"repository"`
	RetrieveBy      string  `pulumi:"retrieveBy"`
	TarballUrl      string  `pulumi:"tarballUrl"`
	TargetCommitish string  `pulumi:"targetCommitish"`
	UploadUrl       string  `pulumi:"uploadUrl"`
	Url             string  `pulumi:"url"`
	ZipballUrl      string  `pulumi:"zipballUrl"`
}

func LookupReleaseOutput(ctx *pulumi.Context, args LookupReleaseOutputArgs, opts ...pulumi.InvokeOption) LookupReleaseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReleaseResult, error) {
			args := v.(LookupReleaseArgs)
			r, err := LookupRelease(ctx, &args, opts...)
			var s LookupReleaseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReleaseResultOutput)
}

// A collection of arguments for invoking getRelease.
type LookupReleaseOutputArgs struct {
	Owner      pulumi.StringInput    `pulumi:"owner"`
	ReleaseId  pulumi.IntPtrInput    `pulumi:"releaseId"`
	ReleaseTag pulumi.StringPtrInput `pulumi:"releaseTag"`
	Repository pulumi.StringInput    `pulumi:"repository"`
	RetrieveBy pulumi.StringInput    `pulumi:"retrieveBy"`
}

func (LookupReleaseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReleaseArgs)(nil)).Elem()
}

// A collection of values returned by getRelease.
type LookupReleaseResultOutput struct{ *pulumi.OutputState }

func (LookupReleaseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReleaseResult)(nil)).Elem()
}

func (o LookupReleaseResultOutput) ToLookupReleaseResultOutput() LookupReleaseResultOutput {
	return o
}

func (o LookupReleaseResultOutput) ToLookupReleaseResultOutputWithContext(ctx context.Context) LookupReleaseResultOutput {
	return o
}

// Deprecated: use assets_url instead
func (o LookupReleaseResultOutput) AssertsUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReleaseResult) string { return v.AssertsUrl }).(pulumi.StringOutput)
}

func (o LookupReleaseResultOutput) Assets() GetReleaseAssetArrayOutput {
	return o.ApplyT(func(v LookupReleaseResult) []GetReleaseAsset { return v.Assets }).(GetReleaseAssetArrayOutput)
}

func (o LookupReleaseResultOutput) AssetsUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReleaseResult) string { return v.AssetsUrl }).(pulumi.StringOutput)
}

func (o LookupReleaseResultOutput) Body() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReleaseResult) string { return v.Body }).(pulumi.StringOutput)
}

func (o LookupReleaseResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReleaseResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupReleaseResultOutput) Draft() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupReleaseResult) bool { return v.Draft }).(pulumi.BoolOutput)
}

func (o LookupReleaseResultOutput) HtmlUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReleaseResult) string { return v.HtmlUrl }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupReleaseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReleaseResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupReleaseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReleaseResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupReleaseResultOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReleaseResult) string { return v.Owner }).(pulumi.StringOutput)
}

func (o LookupReleaseResultOutput) Prerelease() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupReleaseResult) bool { return v.Prerelease }).(pulumi.BoolOutput)
}

func (o LookupReleaseResultOutput) PublishedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReleaseResult) string { return v.PublishedAt }).(pulumi.StringOutput)
}

func (o LookupReleaseResultOutput) ReleaseId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupReleaseResult) *int { return v.ReleaseId }).(pulumi.IntPtrOutput)
}

func (o LookupReleaseResultOutput) ReleaseTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReleaseResult) *string { return v.ReleaseTag }).(pulumi.StringPtrOutput)
}

func (o LookupReleaseResultOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReleaseResult) string { return v.Repository }).(pulumi.StringOutput)
}

func (o LookupReleaseResultOutput) RetrieveBy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReleaseResult) string { return v.RetrieveBy }).(pulumi.StringOutput)
}

func (o LookupReleaseResultOutput) TarballUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReleaseResult) string { return v.TarballUrl }).(pulumi.StringOutput)
}

func (o LookupReleaseResultOutput) TargetCommitish() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReleaseResult) string { return v.TargetCommitish }).(pulumi.StringOutput)
}

func (o LookupReleaseResultOutput) UploadUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReleaseResult) string { return v.UploadUrl }).(pulumi.StringOutput)
}

func (o LookupReleaseResultOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReleaseResult) string { return v.Url }).(pulumi.StringOutput)
}

func (o LookupReleaseResultOutput) ZipballUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReleaseResult) string { return v.ZipballUrl }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReleaseResultOutput{})
}

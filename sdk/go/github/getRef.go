// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetRef(ctx *pulumi.Context, args *GetRefArgs, opts ...pulumi.InvokeOption) (*GetRefResult, error) {
	var rv GetRefResult
	err := ctx.Invoke("github:index/getRef:getRef", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRef.
type GetRefArgs struct {
	Ref        string `pulumi:"ref"`
	Repository string `pulumi:"repository"`
}

// A collection of values returned by getRef.
type GetRefResult struct {
	Etag string `pulumi:"etag"`
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	Ref        string `pulumi:"ref"`
	Repository string `pulumi:"repository"`
	Sha        string `pulumi:"sha"`
}

func GetRefOutput(ctx *pulumi.Context, args GetRefOutputArgs, opts ...pulumi.InvokeOption) GetRefResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRefResult, error) {
			args := v.(GetRefArgs)
			r, err := GetRef(ctx, &args, opts...)
			var s GetRefResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRefResultOutput)
}

// A collection of arguments for invoking getRef.
type GetRefOutputArgs struct {
	Ref        pulumi.StringInput `pulumi:"ref"`
	Repository pulumi.StringInput `pulumi:"repository"`
}

func (GetRefOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRefArgs)(nil)).Elem()
}

// A collection of values returned by getRef.
type GetRefResultOutput struct{ *pulumi.OutputState }

func (GetRefResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRefResult)(nil)).Elem()
}

func (o GetRefResultOutput) ToGetRefResultOutput() GetRefResultOutput {
	return o
}

func (o GetRefResultOutput) ToGetRefResultOutputWithContext(ctx context.Context) GetRefResultOutput {
	return o
}

func (o GetRefResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v GetRefResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRefResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRefResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetRefResultOutput) Ref() pulumi.StringOutput {
	return o.ApplyT(func(v GetRefResult) string { return v.Ref }).(pulumi.StringOutput)
}

func (o GetRefResultOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v GetRefResult) string { return v.Repository }).(pulumi.StringOutput)
}

func (o GetRefResultOutput) Sha() pulumi.StringOutput {
	return o.ApplyT(func(v GetRefResult) string { return v.Sha }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRefResultOutput{})
}

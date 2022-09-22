// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BranchProtectionV3 struct {
	pulumi.CustomResourceState

	Branch                        pulumi.StringOutput                                   `pulumi:"branch"`
	EnforceAdmins                 pulumi.BoolPtrOutput                                  `pulumi:"enforceAdmins"`
	Etag                          pulumi.StringOutput                                   `pulumi:"etag"`
	Repository                    pulumi.StringOutput                                   `pulumi:"repository"`
	RequireConversationResolution pulumi.BoolPtrOutput                                  `pulumi:"requireConversationResolution"`
	RequireSignedCommits          pulumi.BoolPtrOutput                                  `pulumi:"requireSignedCommits"`
	RequiredPullRequestReviews    BranchProtectionV3RequiredPullRequestReviewsPtrOutput `pulumi:"requiredPullRequestReviews"`
	RequiredStatusChecks          BranchProtectionV3RequiredStatusChecksPtrOutput       `pulumi:"requiredStatusChecks"`
	Restrictions                  BranchProtectionV3RestrictionsPtrOutput               `pulumi:"restrictions"`
}

// NewBranchProtectionV3 registers a new resource with the given unique name, arguments, and options.
func NewBranchProtectionV3(ctx *pulumi.Context,
	name string, args *BranchProtectionV3Args, opts ...pulumi.ResourceOption) (*BranchProtectionV3, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Branch == nil {
		return nil, errors.New("invalid value for required argument 'Branch'")
	}
	if args.Repository == nil {
		return nil, errors.New("invalid value for required argument 'Repository'")
	}
	var resource BranchProtectionV3
	err := ctx.RegisterResource("github:index/branchProtectionV3:BranchProtectionV3", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBranchProtectionV3 gets an existing BranchProtectionV3 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBranchProtectionV3(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BranchProtectionV3State, opts ...pulumi.ResourceOption) (*BranchProtectionV3, error) {
	var resource BranchProtectionV3
	err := ctx.ReadResource("github:index/branchProtectionV3:BranchProtectionV3", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BranchProtectionV3 resources.
type branchProtectionV3State struct {
	Branch                        *string                                       `pulumi:"branch"`
	EnforceAdmins                 *bool                                         `pulumi:"enforceAdmins"`
	Etag                          *string                                       `pulumi:"etag"`
	Repository                    *string                                       `pulumi:"repository"`
	RequireConversationResolution *bool                                         `pulumi:"requireConversationResolution"`
	RequireSignedCommits          *bool                                         `pulumi:"requireSignedCommits"`
	RequiredPullRequestReviews    *BranchProtectionV3RequiredPullRequestReviews `pulumi:"requiredPullRequestReviews"`
	RequiredStatusChecks          *BranchProtectionV3RequiredStatusChecks       `pulumi:"requiredStatusChecks"`
	Restrictions                  *BranchProtectionV3Restrictions               `pulumi:"restrictions"`
}

type BranchProtectionV3State struct {
	Branch                        pulumi.StringPtrInput
	EnforceAdmins                 pulumi.BoolPtrInput
	Etag                          pulumi.StringPtrInput
	Repository                    pulumi.StringPtrInput
	RequireConversationResolution pulumi.BoolPtrInput
	RequireSignedCommits          pulumi.BoolPtrInput
	RequiredPullRequestReviews    BranchProtectionV3RequiredPullRequestReviewsPtrInput
	RequiredStatusChecks          BranchProtectionV3RequiredStatusChecksPtrInput
	Restrictions                  BranchProtectionV3RestrictionsPtrInput
}

func (BranchProtectionV3State) ElementType() reflect.Type {
	return reflect.TypeOf((*branchProtectionV3State)(nil)).Elem()
}

type branchProtectionV3Args struct {
	Branch                        string                                        `pulumi:"branch"`
	EnforceAdmins                 *bool                                         `pulumi:"enforceAdmins"`
	Repository                    string                                        `pulumi:"repository"`
	RequireConversationResolution *bool                                         `pulumi:"requireConversationResolution"`
	RequireSignedCommits          *bool                                         `pulumi:"requireSignedCommits"`
	RequiredPullRequestReviews    *BranchProtectionV3RequiredPullRequestReviews `pulumi:"requiredPullRequestReviews"`
	RequiredStatusChecks          *BranchProtectionV3RequiredStatusChecks       `pulumi:"requiredStatusChecks"`
	Restrictions                  *BranchProtectionV3Restrictions               `pulumi:"restrictions"`
}

// The set of arguments for constructing a BranchProtectionV3 resource.
type BranchProtectionV3Args struct {
	Branch                        pulumi.StringInput
	EnforceAdmins                 pulumi.BoolPtrInput
	Repository                    pulumi.StringInput
	RequireConversationResolution pulumi.BoolPtrInput
	RequireSignedCommits          pulumi.BoolPtrInput
	RequiredPullRequestReviews    BranchProtectionV3RequiredPullRequestReviewsPtrInput
	RequiredStatusChecks          BranchProtectionV3RequiredStatusChecksPtrInput
	Restrictions                  BranchProtectionV3RestrictionsPtrInput
}

func (BranchProtectionV3Args) ElementType() reflect.Type {
	return reflect.TypeOf((*branchProtectionV3Args)(nil)).Elem()
}

type BranchProtectionV3Input interface {
	pulumi.Input

	ToBranchProtectionV3Output() BranchProtectionV3Output
	ToBranchProtectionV3OutputWithContext(ctx context.Context) BranchProtectionV3Output
}

func (*BranchProtectionV3) ElementType() reflect.Type {
	return reflect.TypeOf((**BranchProtectionV3)(nil)).Elem()
}

func (i *BranchProtectionV3) ToBranchProtectionV3Output() BranchProtectionV3Output {
	return i.ToBranchProtectionV3OutputWithContext(context.Background())
}

func (i *BranchProtectionV3) ToBranchProtectionV3OutputWithContext(ctx context.Context) BranchProtectionV3Output {
	return pulumi.ToOutputWithContext(ctx, i).(BranchProtectionV3Output)
}

// BranchProtectionV3ArrayInput is an input type that accepts BranchProtectionV3Array and BranchProtectionV3ArrayOutput values.
// You can construct a concrete instance of `BranchProtectionV3ArrayInput` via:
//
//	BranchProtectionV3Array{ BranchProtectionV3Args{...} }
type BranchProtectionV3ArrayInput interface {
	pulumi.Input

	ToBranchProtectionV3ArrayOutput() BranchProtectionV3ArrayOutput
	ToBranchProtectionV3ArrayOutputWithContext(context.Context) BranchProtectionV3ArrayOutput
}

type BranchProtectionV3Array []BranchProtectionV3Input

func (BranchProtectionV3Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BranchProtectionV3)(nil)).Elem()
}

func (i BranchProtectionV3Array) ToBranchProtectionV3ArrayOutput() BranchProtectionV3ArrayOutput {
	return i.ToBranchProtectionV3ArrayOutputWithContext(context.Background())
}

func (i BranchProtectionV3Array) ToBranchProtectionV3ArrayOutputWithContext(ctx context.Context) BranchProtectionV3ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchProtectionV3ArrayOutput)
}

// BranchProtectionV3MapInput is an input type that accepts BranchProtectionV3Map and BranchProtectionV3MapOutput values.
// You can construct a concrete instance of `BranchProtectionV3MapInput` via:
//
//	BranchProtectionV3Map{ "key": BranchProtectionV3Args{...} }
type BranchProtectionV3MapInput interface {
	pulumi.Input

	ToBranchProtectionV3MapOutput() BranchProtectionV3MapOutput
	ToBranchProtectionV3MapOutputWithContext(context.Context) BranchProtectionV3MapOutput
}

type BranchProtectionV3Map map[string]BranchProtectionV3Input

func (BranchProtectionV3Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BranchProtectionV3)(nil)).Elem()
}

func (i BranchProtectionV3Map) ToBranchProtectionV3MapOutput() BranchProtectionV3MapOutput {
	return i.ToBranchProtectionV3MapOutputWithContext(context.Background())
}

func (i BranchProtectionV3Map) ToBranchProtectionV3MapOutputWithContext(ctx context.Context) BranchProtectionV3MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchProtectionV3MapOutput)
}

type BranchProtectionV3Output struct{ *pulumi.OutputState }

func (BranchProtectionV3Output) ElementType() reflect.Type {
	return reflect.TypeOf((**BranchProtectionV3)(nil)).Elem()
}

func (o BranchProtectionV3Output) ToBranchProtectionV3Output() BranchProtectionV3Output {
	return o
}

func (o BranchProtectionV3Output) ToBranchProtectionV3OutputWithContext(ctx context.Context) BranchProtectionV3Output {
	return o
}

func (o BranchProtectionV3Output) Branch() pulumi.StringOutput {
	return o.ApplyT(func(v *BranchProtectionV3) pulumi.StringOutput { return v.Branch }).(pulumi.StringOutput)
}

func (o BranchProtectionV3Output) EnforceAdmins() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BranchProtectionV3) pulumi.BoolPtrOutput { return v.EnforceAdmins }).(pulumi.BoolPtrOutput)
}

func (o BranchProtectionV3Output) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *BranchProtectionV3) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o BranchProtectionV3Output) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v *BranchProtectionV3) pulumi.StringOutput { return v.Repository }).(pulumi.StringOutput)
}

func (o BranchProtectionV3Output) RequireConversationResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BranchProtectionV3) pulumi.BoolPtrOutput { return v.RequireConversationResolution }).(pulumi.BoolPtrOutput)
}

func (o BranchProtectionV3Output) RequireSignedCommits() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BranchProtectionV3) pulumi.BoolPtrOutput { return v.RequireSignedCommits }).(pulumi.BoolPtrOutput)
}

func (o BranchProtectionV3Output) RequiredPullRequestReviews() BranchProtectionV3RequiredPullRequestReviewsPtrOutput {
	return o.ApplyT(func(v *BranchProtectionV3) BranchProtectionV3RequiredPullRequestReviewsPtrOutput {
		return v.RequiredPullRequestReviews
	}).(BranchProtectionV3RequiredPullRequestReviewsPtrOutput)
}

func (o BranchProtectionV3Output) RequiredStatusChecks() BranchProtectionV3RequiredStatusChecksPtrOutput {
	return o.ApplyT(func(v *BranchProtectionV3) BranchProtectionV3RequiredStatusChecksPtrOutput {
		return v.RequiredStatusChecks
	}).(BranchProtectionV3RequiredStatusChecksPtrOutput)
}

func (o BranchProtectionV3Output) Restrictions() BranchProtectionV3RestrictionsPtrOutput {
	return o.ApplyT(func(v *BranchProtectionV3) BranchProtectionV3RestrictionsPtrOutput { return v.Restrictions }).(BranchProtectionV3RestrictionsPtrOutput)
}

type BranchProtectionV3ArrayOutput struct{ *pulumi.OutputState }

func (BranchProtectionV3ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BranchProtectionV3)(nil)).Elem()
}

func (o BranchProtectionV3ArrayOutput) ToBranchProtectionV3ArrayOutput() BranchProtectionV3ArrayOutput {
	return o
}

func (o BranchProtectionV3ArrayOutput) ToBranchProtectionV3ArrayOutputWithContext(ctx context.Context) BranchProtectionV3ArrayOutput {
	return o
}

func (o BranchProtectionV3ArrayOutput) Index(i pulumi.IntInput) BranchProtectionV3Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BranchProtectionV3 {
		return vs[0].([]*BranchProtectionV3)[vs[1].(int)]
	}).(BranchProtectionV3Output)
}

type BranchProtectionV3MapOutput struct{ *pulumi.OutputState }

func (BranchProtectionV3MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BranchProtectionV3)(nil)).Elem()
}

func (o BranchProtectionV3MapOutput) ToBranchProtectionV3MapOutput() BranchProtectionV3MapOutput {
	return o
}

func (o BranchProtectionV3MapOutput) ToBranchProtectionV3MapOutputWithContext(ctx context.Context) BranchProtectionV3MapOutput {
	return o
}

func (o BranchProtectionV3MapOutput) MapIndex(k pulumi.StringInput) BranchProtectionV3Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BranchProtectionV3 {
		return vs[0].(map[string]*BranchProtectionV3)[vs[1].(string)]
	}).(BranchProtectionV3Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BranchProtectionV3Input)(nil)).Elem(), &BranchProtectionV3{})
	pulumi.RegisterInputType(reflect.TypeOf((*BranchProtectionV3ArrayInput)(nil)).Elem(), BranchProtectionV3Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*BranchProtectionV3MapInput)(nil)).Elem(), BranchProtectionV3Map{})
	pulumi.RegisterOutputType(BranchProtectionV3Output{})
	pulumi.RegisterOutputType(BranchProtectionV3ArrayOutput{})
	pulumi.RegisterOutputType(BranchProtectionV3MapOutput{})
}

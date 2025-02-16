// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BranchProtection struct {
	pulumi.CustomResourceState

	AllowsDeletions   pulumi.BoolPtrOutput     `pulumi:"allowsDeletions"`
	AllowsForcePushes pulumi.BoolPtrOutput     `pulumi:"allowsForcePushes"`
	BlocksCreations   pulumi.BoolPtrOutput     `pulumi:"blocksCreations"`
	EnforceAdmins     pulumi.BoolPtrOutput     `pulumi:"enforceAdmins"`
	LockBranch        pulumi.BoolPtrOutput     `pulumi:"lockBranch"`
	Pattern           pulumi.StringOutput      `pulumi:"pattern"`
	PushRestrictions  pulumi.StringArrayOutput `pulumi:"pushRestrictions"`
	// Node ID or name of repository
	RepositoryId                  pulumi.StringOutput                                  `pulumi:"repositoryId"`
	RequireConversationResolution pulumi.BoolPtrOutput                                 `pulumi:"requireConversationResolution"`
	RequireSignedCommits          pulumi.BoolPtrOutput                                 `pulumi:"requireSignedCommits"`
	RequiredLinearHistory         pulumi.BoolPtrOutput                                 `pulumi:"requiredLinearHistory"`
	RequiredPullRequestReviews    BranchProtectionRequiredPullRequestReviewArrayOutput `pulumi:"requiredPullRequestReviews"`
	RequiredStatusChecks          BranchProtectionRequiredStatusCheckArrayOutput       `pulumi:"requiredStatusChecks"`
}

// NewBranchProtection registers a new resource with the given unique name, arguments, and options.
func NewBranchProtection(ctx *pulumi.Context,
	name string, args *BranchProtectionArgs, opts ...pulumi.ResourceOption) (*BranchProtection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Pattern == nil {
		return nil, errors.New("invalid value for required argument 'Pattern'")
	}
	if args.RepositoryId == nil {
		return nil, errors.New("invalid value for required argument 'RepositoryId'")
	}
	var resource BranchProtection
	err := ctx.RegisterResource("github:index/branchProtection:BranchProtection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBranchProtection gets an existing BranchProtection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBranchProtection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BranchProtectionState, opts ...pulumi.ResourceOption) (*BranchProtection, error) {
	var resource BranchProtection
	err := ctx.ReadResource("github:index/branchProtection:BranchProtection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BranchProtection resources.
type branchProtectionState struct {
	AllowsDeletions   *bool    `pulumi:"allowsDeletions"`
	AllowsForcePushes *bool    `pulumi:"allowsForcePushes"`
	BlocksCreations   *bool    `pulumi:"blocksCreations"`
	EnforceAdmins     *bool    `pulumi:"enforceAdmins"`
	LockBranch        *bool    `pulumi:"lockBranch"`
	Pattern           *string  `pulumi:"pattern"`
	PushRestrictions  []string `pulumi:"pushRestrictions"`
	// Node ID or name of repository
	RepositoryId                  *string                                     `pulumi:"repositoryId"`
	RequireConversationResolution *bool                                       `pulumi:"requireConversationResolution"`
	RequireSignedCommits          *bool                                       `pulumi:"requireSignedCommits"`
	RequiredLinearHistory         *bool                                       `pulumi:"requiredLinearHistory"`
	RequiredPullRequestReviews    []BranchProtectionRequiredPullRequestReview `pulumi:"requiredPullRequestReviews"`
	RequiredStatusChecks          []BranchProtectionRequiredStatusCheck       `pulumi:"requiredStatusChecks"`
}

type BranchProtectionState struct {
	AllowsDeletions   pulumi.BoolPtrInput
	AllowsForcePushes pulumi.BoolPtrInput
	BlocksCreations   pulumi.BoolPtrInput
	EnforceAdmins     pulumi.BoolPtrInput
	LockBranch        pulumi.BoolPtrInput
	Pattern           pulumi.StringPtrInput
	PushRestrictions  pulumi.StringArrayInput
	// Node ID or name of repository
	RepositoryId                  pulumi.StringPtrInput
	RequireConversationResolution pulumi.BoolPtrInput
	RequireSignedCommits          pulumi.BoolPtrInput
	RequiredLinearHistory         pulumi.BoolPtrInput
	RequiredPullRequestReviews    BranchProtectionRequiredPullRequestReviewArrayInput
	RequiredStatusChecks          BranchProtectionRequiredStatusCheckArrayInput
}

func (BranchProtectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*branchProtectionState)(nil)).Elem()
}

type branchProtectionArgs struct {
	AllowsDeletions   *bool    `pulumi:"allowsDeletions"`
	AllowsForcePushes *bool    `pulumi:"allowsForcePushes"`
	BlocksCreations   *bool    `pulumi:"blocksCreations"`
	EnforceAdmins     *bool    `pulumi:"enforceAdmins"`
	LockBranch        *bool    `pulumi:"lockBranch"`
	Pattern           string   `pulumi:"pattern"`
	PushRestrictions  []string `pulumi:"pushRestrictions"`
	// Node ID or name of repository
	RepositoryId                  string                                      `pulumi:"repositoryId"`
	RequireConversationResolution *bool                                       `pulumi:"requireConversationResolution"`
	RequireSignedCommits          *bool                                       `pulumi:"requireSignedCommits"`
	RequiredLinearHistory         *bool                                       `pulumi:"requiredLinearHistory"`
	RequiredPullRequestReviews    []BranchProtectionRequiredPullRequestReview `pulumi:"requiredPullRequestReviews"`
	RequiredStatusChecks          []BranchProtectionRequiredStatusCheck       `pulumi:"requiredStatusChecks"`
}

// The set of arguments for constructing a BranchProtection resource.
type BranchProtectionArgs struct {
	AllowsDeletions   pulumi.BoolPtrInput
	AllowsForcePushes pulumi.BoolPtrInput
	BlocksCreations   pulumi.BoolPtrInput
	EnforceAdmins     pulumi.BoolPtrInput
	LockBranch        pulumi.BoolPtrInput
	Pattern           pulumi.StringInput
	PushRestrictions  pulumi.StringArrayInput
	// Node ID or name of repository
	RepositoryId                  pulumi.StringInput
	RequireConversationResolution pulumi.BoolPtrInput
	RequireSignedCommits          pulumi.BoolPtrInput
	RequiredLinearHistory         pulumi.BoolPtrInput
	RequiredPullRequestReviews    BranchProtectionRequiredPullRequestReviewArrayInput
	RequiredStatusChecks          BranchProtectionRequiredStatusCheckArrayInput
}

func (BranchProtectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*branchProtectionArgs)(nil)).Elem()
}

type BranchProtectionInput interface {
	pulumi.Input

	ToBranchProtectionOutput() BranchProtectionOutput
	ToBranchProtectionOutputWithContext(ctx context.Context) BranchProtectionOutput
}

func (*BranchProtection) ElementType() reflect.Type {
	return reflect.TypeOf((**BranchProtection)(nil)).Elem()
}

func (i *BranchProtection) ToBranchProtectionOutput() BranchProtectionOutput {
	return i.ToBranchProtectionOutputWithContext(context.Background())
}

func (i *BranchProtection) ToBranchProtectionOutputWithContext(ctx context.Context) BranchProtectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchProtectionOutput)
}

// BranchProtectionArrayInput is an input type that accepts BranchProtectionArray and BranchProtectionArrayOutput values.
// You can construct a concrete instance of `BranchProtectionArrayInput` via:
//
//          BranchProtectionArray{ BranchProtectionArgs{...} }
type BranchProtectionArrayInput interface {
	pulumi.Input

	ToBranchProtectionArrayOutput() BranchProtectionArrayOutput
	ToBranchProtectionArrayOutputWithContext(context.Context) BranchProtectionArrayOutput
}

type BranchProtectionArray []BranchProtectionInput

func (BranchProtectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BranchProtection)(nil)).Elem()
}

func (i BranchProtectionArray) ToBranchProtectionArrayOutput() BranchProtectionArrayOutput {
	return i.ToBranchProtectionArrayOutputWithContext(context.Background())
}

func (i BranchProtectionArray) ToBranchProtectionArrayOutputWithContext(ctx context.Context) BranchProtectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchProtectionArrayOutput)
}

// BranchProtectionMapInput is an input type that accepts BranchProtectionMap and BranchProtectionMapOutput values.
// You can construct a concrete instance of `BranchProtectionMapInput` via:
//
//          BranchProtectionMap{ "key": BranchProtectionArgs{...} }
type BranchProtectionMapInput interface {
	pulumi.Input

	ToBranchProtectionMapOutput() BranchProtectionMapOutput
	ToBranchProtectionMapOutputWithContext(context.Context) BranchProtectionMapOutput
}

type BranchProtectionMap map[string]BranchProtectionInput

func (BranchProtectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BranchProtection)(nil)).Elem()
}

func (i BranchProtectionMap) ToBranchProtectionMapOutput() BranchProtectionMapOutput {
	return i.ToBranchProtectionMapOutputWithContext(context.Background())
}

func (i BranchProtectionMap) ToBranchProtectionMapOutputWithContext(ctx context.Context) BranchProtectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchProtectionMapOutput)
}

type BranchProtectionOutput struct{ *pulumi.OutputState }

func (BranchProtectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BranchProtection)(nil)).Elem()
}

func (o BranchProtectionOutput) ToBranchProtectionOutput() BranchProtectionOutput {
	return o
}

func (o BranchProtectionOutput) ToBranchProtectionOutputWithContext(ctx context.Context) BranchProtectionOutput {
	return o
}

func (o BranchProtectionOutput) AllowsDeletions() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BranchProtection) pulumi.BoolPtrOutput { return v.AllowsDeletions }).(pulumi.BoolPtrOutput)
}

func (o BranchProtectionOutput) AllowsForcePushes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BranchProtection) pulumi.BoolPtrOutput { return v.AllowsForcePushes }).(pulumi.BoolPtrOutput)
}

func (o BranchProtectionOutput) BlocksCreations() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BranchProtection) pulumi.BoolPtrOutput { return v.BlocksCreations }).(pulumi.BoolPtrOutput)
}

func (o BranchProtectionOutput) EnforceAdmins() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BranchProtection) pulumi.BoolPtrOutput { return v.EnforceAdmins }).(pulumi.BoolPtrOutput)
}

func (o BranchProtectionOutput) LockBranch() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BranchProtection) pulumi.BoolPtrOutput { return v.LockBranch }).(pulumi.BoolPtrOutput)
}

func (o BranchProtectionOutput) Pattern() pulumi.StringOutput {
	return o.ApplyT(func(v *BranchProtection) pulumi.StringOutput { return v.Pattern }).(pulumi.StringOutput)
}

func (o BranchProtectionOutput) PushRestrictions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BranchProtection) pulumi.StringArrayOutput { return v.PushRestrictions }).(pulumi.StringArrayOutput)
}

// Node ID or name of repository
func (o BranchProtectionOutput) RepositoryId() pulumi.StringOutput {
	return o.ApplyT(func(v *BranchProtection) pulumi.StringOutput { return v.RepositoryId }).(pulumi.StringOutput)
}

func (o BranchProtectionOutput) RequireConversationResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BranchProtection) pulumi.BoolPtrOutput { return v.RequireConversationResolution }).(pulumi.BoolPtrOutput)
}

func (o BranchProtectionOutput) RequireSignedCommits() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BranchProtection) pulumi.BoolPtrOutput { return v.RequireSignedCommits }).(pulumi.BoolPtrOutput)
}

func (o BranchProtectionOutput) RequiredLinearHistory() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BranchProtection) pulumi.BoolPtrOutput { return v.RequiredLinearHistory }).(pulumi.BoolPtrOutput)
}

func (o BranchProtectionOutput) RequiredPullRequestReviews() BranchProtectionRequiredPullRequestReviewArrayOutput {
	return o.ApplyT(func(v *BranchProtection) BranchProtectionRequiredPullRequestReviewArrayOutput {
		return v.RequiredPullRequestReviews
	}).(BranchProtectionRequiredPullRequestReviewArrayOutput)
}

func (o BranchProtectionOutput) RequiredStatusChecks() BranchProtectionRequiredStatusCheckArrayOutput {
	return o.ApplyT(func(v *BranchProtection) BranchProtectionRequiredStatusCheckArrayOutput {
		return v.RequiredStatusChecks
	}).(BranchProtectionRequiredStatusCheckArrayOutput)
}

type BranchProtectionArrayOutput struct{ *pulumi.OutputState }

func (BranchProtectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BranchProtection)(nil)).Elem()
}

func (o BranchProtectionArrayOutput) ToBranchProtectionArrayOutput() BranchProtectionArrayOutput {
	return o
}

func (o BranchProtectionArrayOutput) ToBranchProtectionArrayOutputWithContext(ctx context.Context) BranchProtectionArrayOutput {
	return o
}

func (o BranchProtectionArrayOutput) Index(i pulumi.IntInput) BranchProtectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BranchProtection {
		return vs[0].([]*BranchProtection)[vs[1].(int)]
	}).(BranchProtectionOutput)
}

type BranchProtectionMapOutput struct{ *pulumi.OutputState }

func (BranchProtectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BranchProtection)(nil)).Elem()
}

func (o BranchProtectionMapOutput) ToBranchProtectionMapOutput() BranchProtectionMapOutput {
	return o
}

func (o BranchProtectionMapOutput) ToBranchProtectionMapOutputWithContext(ctx context.Context) BranchProtectionMapOutput {
	return o
}

func (o BranchProtectionMapOutput) MapIndex(k pulumi.StringInput) BranchProtectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BranchProtection {
		return vs[0].(map[string]*BranchProtection)[vs[1].(string)]
	}).(BranchProtectionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BranchProtectionInput)(nil)).Elem(), &BranchProtection{})
	pulumi.RegisterInputType(reflect.TypeOf((*BranchProtectionArrayInput)(nil)).Elem(), BranchProtectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BranchProtectionMapInput)(nil)).Elem(), BranchProtectionMap{})
	pulumi.RegisterOutputType(BranchProtectionOutput{})
	pulumi.RegisterOutputType(BranchProtectionArrayOutput{})
	pulumi.RegisterOutputType(BranchProtectionMapOutput{})
}

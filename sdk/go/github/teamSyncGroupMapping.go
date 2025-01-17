// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TeamSyncGroupMapping struct {
	pulumi.CustomResourceState

	Etag     pulumi.StringOutput                  `pulumi:"etag"`
	Groups   TeamSyncGroupMappingGroupArrayOutput `pulumi:"groups"`
	TeamSlug pulumi.StringOutput                  `pulumi:"teamSlug"`
}

// NewTeamSyncGroupMapping registers a new resource with the given unique name, arguments, and options.
func NewTeamSyncGroupMapping(ctx *pulumi.Context,
	name string, args *TeamSyncGroupMappingArgs, opts ...pulumi.ResourceOption) (*TeamSyncGroupMapping, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.TeamSlug == nil {
		return nil, errors.New("invalid value for required argument 'TeamSlug'")
	}
	var resource TeamSyncGroupMapping
	err := ctx.RegisterResource("github:index/teamSyncGroupMapping:TeamSyncGroupMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTeamSyncGroupMapping gets an existing TeamSyncGroupMapping resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTeamSyncGroupMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TeamSyncGroupMappingState, opts ...pulumi.ResourceOption) (*TeamSyncGroupMapping, error) {
	var resource TeamSyncGroupMapping
	err := ctx.ReadResource("github:index/teamSyncGroupMapping:TeamSyncGroupMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TeamSyncGroupMapping resources.
type teamSyncGroupMappingState struct {
	Etag     *string                     `pulumi:"etag"`
	Groups   []TeamSyncGroupMappingGroup `pulumi:"groups"`
	TeamSlug *string                     `pulumi:"teamSlug"`
}

type TeamSyncGroupMappingState struct {
	Etag     pulumi.StringPtrInput
	Groups   TeamSyncGroupMappingGroupArrayInput
	TeamSlug pulumi.StringPtrInput
}

func (TeamSyncGroupMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*teamSyncGroupMappingState)(nil)).Elem()
}

type teamSyncGroupMappingArgs struct {
	Groups   []TeamSyncGroupMappingGroup `pulumi:"groups"`
	TeamSlug string                      `pulumi:"teamSlug"`
}

// The set of arguments for constructing a TeamSyncGroupMapping resource.
type TeamSyncGroupMappingArgs struct {
	Groups   TeamSyncGroupMappingGroupArrayInput
	TeamSlug pulumi.StringInput
}

func (TeamSyncGroupMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*teamSyncGroupMappingArgs)(nil)).Elem()
}

type TeamSyncGroupMappingInput interface {
	pulumi.Input

	ToTeamSyncGroupMappingOutput() TeamSyncGroupMappingOutput
	ToTeamSyncGroupMappingOutputWithContext(ctx context.Context) TeamSyncGroupMappingOutput
}

func (*TeamSyncGroupMapping) ElementType() reflect.Type {
	return reflect.TypeOf((**TeamSyncGroupMapping)(nil)).Elem()
}

func (i *TeamSyncGroupMapping) ToTeamSyncGroupMappingOutput() TeamSyncGroupMappingOutput {
	return i.ToTeamSyncGroupMappingOutputWithContext(context.Background())
}

func (i *TeamSyncGroupMapping) ToTeamSyncGroupMappingOutputWithContext(ctx context.Context) TeamSyncGroupMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamSyncGroupMappingOutput)
}

// TeamSyncGroupMappingArrayInput is an input type that accepts TeamSyncGroupMappingArray and TeamSyncGroupMappingArrayOutput values.
// You can construct a concrete instance of `TeamSyncGroupMappingArrayInput` via:
//
//          TeamSyncGroupMappingArray{ TeamSyncGroupMappingArgs{...} }
type TeamSyncGroupMappingArrayInput interface {
	pulumi.Input

	ToTeamSyncGroupMappingArrayOutput() TeamSyncGroupMappingArrayOutput
	ToTeamSyncGroupMappingArrayOutputWithContext(context.Context) TeamSyncGroupMappingArrayOutput
}

type TeamSyncGroupMappingArray []TeamSyncGroupMappingInput

func (TeamSyncGroupMappingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TeamSyncGroupMapping)(nil)).Elem()
}

func (i TeamSyncGroupMappingArray) ToTeamSyncGroupMappingArrayOutput() TeamSyncGroupMappingArrayOutput {
	return i.ToTeamSyncGroupMappingArrayOutputWithContext(context.Background())
}

func (i TeamSyncGroupMappingArray) ToTeamSyncGroupMappingArrayOutputWithContext(ctx context.Context) TeamSyncGroupMappingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamSyncGroupMappingArrayOutput)
}

// TeamSyncGroupMappingMapInput is an input type that accepts TeamSyncGroupMappingMap and TeamSyncGroupMappingMapOutput values.
// You can construct a concrete instance of `TeamSyncGroupMappingMapInput` via:
//
//          TeamSyncGroupMappingMap{ "key": TeamSyncGroupMappingArgs{...} }
type TeamSyncGroupMappingMapInput interface {
	pulumi.Input

	ToTeamSyncGroupMappingMapOutput() TeamSyncGroupMappingMapOutput
	ToTeamSyncGroupMappingMapOutputWithContext(context.Context) TeamSyncGroupMappingMapOutput
}

type TeamSyncGroupMappingMap map[string]TeamSyncGroupMappingInput

func (TeamSyncGroupMappingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TeamSyncGroupMapping)(nil)).Elem()
}

func (i TeamSyncGroupMappingMap) ToTeamSyncGroupMappingMapOutput() TeamSyncGroupMappingMapOutput {
	return i.ToTeamSyncGroupMappingMapOutputWithContext(context.Background())
}

func (i TeamSyncGroupMappingMap) ToTeamSyncGroupMappingMapOutputWithContext(ctx context.Context) TeamSyncGroupMappingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamSyncGroupMappingMapOutput)
}

type TeamSyncGroupMappingOutput struct{ *pulumi.OutputState }

func (TeamSyncGroupMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TeamSyncGroupMapping)(nil)).Elem()
}

func (o TeamSyncGroupMappingOutput) ToTeamSyncGroupMappingOutput() TeamSyncGroupMappingOutput {
	return o
}

func (o TeamSyncGroupMappingOutput) ToTeamSyncGroupMappingOutputWithContext(ctx context.Context) TeamSyncGroupMappingOutput {
	return o
}

func (o TeamSyncGroupMappingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *TeamSyncGroupMapping) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o TeamSyncGroupMappingOutput) Groups() TeamSyncGroupMappingGroupArrayOutput {
	return o.ApplyT(func(v *TeamSyncGroupMapping) TeamSyncGroupMappingGroupArrayOutput { return v.Groups }).(TeamSyncGroupMappingGroupArrayOutput)
}

func (o TeamSyncGroupMappingOutput) TeamSlug() pulumi.StringOutput {
	return o.ApplyT(func(v *TeamSyncGroupMapping) pulumi.StringOutput { return v.TeamSlug }).(pulumi.StringOutput)
}

type TeamSyncGroupMappingArrayOutput struct{ *pulumi.OutputState }

func (TeamSyncGroupMappingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TeamSyncGroupMapping)(nil)).Elem()
}

func (o TeamSyncGroupMappingArrayOutput) ToTeamSyncGroupMappingArrayOutput() TeamSyncGroupMappingArrayOutput {
	return o
}

func (o TeamSyncGroupMappingArrayOutput) ToTeamSyncGroupMappingArrayOutputWithContext(ctx context.Context) TeamSyncGroupMappingArrayOutput {
	return o
}

func (o TeamSyncGroupMappingArrayOutput) Index(i pulumi.IntInput) TeamSyncGroupMappingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TeamSyncGroupMapping {
		return vs[0].([]*TeamSyncGroupMapping)[vs[1].(int)]
	}).(TeamSyncGroupMappingOutput)
}

type TeamSyncGroupMappingMapOutput struct{ *pulumi.OutputState }

func (TeamSyncGroupMappingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TeamSyncGroupMapping)(nil)).Elem()
}

func (o TeamSyncGroupMappingMapOutput) ToTeamSyncGroupMappingMapOutput() TeamSyncGroupMappingMapOutput {
	return o
}

func (o TeamSyncGroupMappingMapOutput) ToTeamSyncGroupMappingMapOutputWithContext(ctx context.Context) TeamSyncGroupMappingMapOutput {
	return o
}

func (o TeamSyncGroupMappingMapOutput) MapIndex(k pulumi.StringInput) TeamSyncGroupMappingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TeamSyncGroupMapping {
		return vs[0].(map[string]*TeamSyncGroupMapping)[vs[1].(string)]
	}).(TeamSyncGroupMappingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TeamSyncGroupMappingInput)(nil)).Elem(), &TeamSyncGroupMapping{})
	pulumi.RegisterInputType(reflect.TypeOf((*TeamSyncGroupMappingArrayInput)(nil)).Elem(), TeamSyncGroupMappingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TeamSyncGroupMappingMapInput)(nil)).Elem(), TeamSyncGroupMappingMap{})
	pulumi.RegisterOutputType(TeamSyncGroupMappingOutput{})
	pulumi.RegisterOutputType(TeamSyncGroupMappingArrayOutput{})
	pulumi.RegisterOutputType(TeamSyncGroupMappingMapOutput{})
}

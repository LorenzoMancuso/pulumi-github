// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EmuGroupMapping struct {
	pulumi.CustomResourceState

	Etag     pulumi.StringOutput `pulumi:"etag"`
	GroupId  pulumi.IntOutput    `pulumi:"groupId"`
	TeamSlug pulumi.StringOutput `pulumi:"teamSlug"`
}

// NewEmuGroupMapping registers a new resource with the given unique name, arguments, and options.
func NewEmuGroupMapping(ctx *pulumi.Context,
	name string, args *EmuGroupMappingArgs, opts ...pulumi.ResourceOption) (*EmuGroupMapping, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	if args.TeamSlug == nil {
		return nil, errors.New("invalid value for required argument 'TeamSlug'")
	}
	var resource EmuGroupMapping
	err := ctx.RegisterResource("github:index/emuGroupMapping:EmuGroupMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEmuGroupMapping gets an existing EmuGroupMapping resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEmuGroupMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EmuGroupMappingState, opts ...pulumi.ResourceOption) (*EmuGroupMapping, error) {
	var resource EmuGroupMapping
	err := ctx.ReadResource("github:index/emuGroupMapping:EmuGroupMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EmuGroupMapping resources.
type emuGroupMappingState struct {
	Etag     *string `pulumi:"etag"`
	GroupId  *int    `pulumi:"groupId"`
	TeamSlug *string `pulumi:"teamSlug"`
}

type EmuGroupMappingState struct {
	Etag     pulumi.StringPtrInput
	GroupId  pulumi.IntPtrInput
	TeamSlug pulumi.StringPtrInput
}

func (EmuGroupMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*emuGroupMappingState)(nil)).Elem()
}

type emuGroupMappingArgs struct {
	GroupId  int    `pulumi:"groupId"`
	TeamSlug string `pulumi:"teamSlug"`
}

// The set of arguments for constructing a EmuGroupMapping resource.
type EmuGroupMappingArgs struct {
	GroupId  pulumi.IntInput
	TeamSlug pulumi.StringInput
}

func (EmuGroupMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*emuGroupMappingArgs)(nil)).Elem()
}

type EmuGroupMappingInput interface {
	pulumi.Input

	ToEmuGroupMappingOutput() EmuGroupMappingOutput
	ToEmuGroupMappingOutputWithContext(ctx context.Context) EmuGroupMappingOutput
}

func (*EmuGroupMapping) ElementType() reflect.Type {
	return reflect.TypeOf((**EmuGroupMapping)(nil)).Elem()
}

func (i *EmuGroupMapping) ToEmuGroupMappingOutput() EmuGroupMappingOutput {
	return i.ToEmuGroupMappingOutputWithContext(context.Background())
}

func (i *EmuGroupMapping) ToEmuGroupMappingOutputWithContext(ctx context.Context) EmuGroupMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmuGroupMappingOutput)
}

// EmuGroupMappingArrayInput is an input type that accepts EmuGroupMappingArray and EmuGroupMappingArrayOutput values.
// You can construct a concrete instance of `EmuGroupMappingArrayInput` via:
//
//          EmuGroupMappingArray{ EmuGroupMappingArgs{...} }
type EmuGroupMappingArrayInput interface {
	pulumi.Input

	ToEmuGroupMappingArrayOutput() EmuGroupMappingArrayOutput
	ToEmuGroupMappingArrayOutputWithContext(context.Context) EmuGroupMappingArrayOutput
}

type EmuGroupMappingArray []EmuGroupMappingInput

func (EmuGroupMappingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EmuGroupMapping)(nil)).Elem()
}

func (i EmuGroupMappingArray) ToEmuGroupMappingArrayOutput() EmuGroupMappingArrayOutput {
	return i.ToEmuGroupMappingArrayOutputWithContext(context.Background())
}

func (i EmuGroupMappingArray) ToEmuGroupMappingArrayOutputWithContext(ctx context.Context) EmuGroupMappingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmuGroupMappingArrayOutput)
}

// EmuGroupMappingMapInput is an input type that accepts EmuGroupMappingMap and EmuGroupMappingMapOutput values.
// You can construct a concrete instance of `EmuGroupMappingMapInput` via:
//
//          EmuGroupMappingMap{ "key": EmuGroupMappingArgs{...} }
type EmuGroupMappingMapInput interface {
	pulumi.Input

	ToEmuGroupMappingMapOutput() EmuGroupMappingMapOutput
	ToEmuGroupMappingMapOutputWithContext(context.Context) EmuGroupMappingMapOutput
}

type EmuGroupMappingMap map[string]EmuGroupMappingInput

func (EmuGroupMappingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EmuGroupMapping)(nil)).Elem()
}

func (i EmuGroupMappingMap) ToEmuGroupMappingMapOutput() EmuGroupMappingMapOutput {
	return i.ToEmuGroupMappingMapOutputWithContext(context.Background())
}

func (i EmuGroupMappingMap) ToEmuGroupMappingMapOutputWithContext(ctx context.Context) EmuGroupMappingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmuGroupMappingMapOutput)
}

type EmuGroupMappingOutput struct{ *pulumi.OutputState }

func (EmuGroupMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EmuGroupMapping)(nil)).Elem()
}

func (o EmuGroupMappingOutput) ToEmuGroupMappingOutput() EmuGroupMappingOutput {
	return o
}

func (o EmuGroupMappingOutput) ToEmuGroupMappingOutputWithContext(ctx context.Context) EmuGroupMappingOutput {
	return o
}

func (o EmuGroupMappingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *EmuGroupMapping) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o EmuGroupMappingOutput) GroupId() pulumi.IntOutput {
	return o.ApplyT(func(v *EmuGroupMapping) pulumi.IntOutput { return v.GroupId }).(pulumi.IntOutput)
}

func (o EmuGroupMappingOutput) TeamSlug() pulumi.StringOutput {
	return o.ApplyT(func(v *EmuGroupMapping) pulumi.StringOutput { return v.TeamSlug }).(pulumi.StringOutput)
}

type EmuGroupMappingArrayOutput struct{ *pulumi.OutputState }

func (EmuGroupMappingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EmuGroupMapping)(nil)).Elem()
}

func (o EmuGroupMappingArrayOutput) ToEmuGroupMappingArrayOutput() EmuGroupMappingArrayOutput {
	return o
}

func (o EmuGroupMappingArrayOutput) ToEmuGroupMappingArrayOutputWithContext(ctx context.Context) EmuGroupMappingArrayOutput {
	return o
}

func (o EmuGroupMappingArrayOutput) Index(i pulumi.IntInput) EmuGroupMappingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EmuGroupMapping {
		return vs[0].([]*EmuGroupMapping)[vs[1].(int)]
	}).(EmuGroupMappingOutput)
}

type EmuGroupMappingMapOutput struct{ *pulumi.OutputState }

func (EmuGroupMappingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EmuGroupMapping)(nil)).Elem()
}

func (o EmuGroupMappingMapOutput) ToEmuGroupMappingMapOutput() EmuGroupMappingMapOutput {
	return o
}

func (o EmuGroupMappingMapOutput) ToEmuGroupMappingMapOutputWithContext(ctx context.Context) EmuGroupMappingMapOutput {
	return o
}

func (o EmuGroupMappingMapOutput) MapIndex(k pulumi.StringInput) EmuGroupMappingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EmuGroupMapping {
		return vs[0].(map[string]*EmuGroupMapping)[vs[1].(string)]
	}).(EmuGroupMappingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EmuGroupMappingInput)(nil)).Elem(), &EmuGroupMapping{})
	pulumi.RegisterInputType(reflect.TypeOf((*EmuGroupMappingArrayInput)(nil)).Elem(), EmuGroupMappingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EmuGroupMappingMapInput)(nil)).Elem(), EmuGroupMappingMap{})
	pulumi.RegisterOutputType(EmuGroupMappingOutput{})
	pulumi.RegisterOutputType(EmuGroupMappingArrayOutput{})
	pulumi.RegisterOutputType(EmuGroupMappingMapOutput{})
}

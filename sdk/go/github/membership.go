// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Membership struct {
	pulumi.CustomResourceState

	Etag     pulumi.StringOutput    `pulumi:"etag"`
	Role     pulumi.StringPtrOutput `pulumi:"role"`
	Username pulumi.StringOutput    `pulumi:"username"`
}

// NewMembership registers a new resource with the given unique name, arguments, and options.
func NewMembership(ctx *pulumi.Context,
	name string, args *MembershipArgs, opts ...pulumi.ResourceOption) (*Membership, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	var resource Membership
	err := ctx.RegisterResource("github:index/membership:Membership", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMembership gets an existing Membership resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMembership(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MembershipState, opts ...pulumi.ResourceOption) (*Membership, error) {
	var resource Membership
	err := ctx.ReadResource("github:index/membership:Membership", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Membership resources.
type membershipState struct {
	Etag     *string `pulumi:"etag"`
	Role     *string `pulumi:"role"`
	Username *string `pulumi:"username"`
}

type MembershipState struct {
	Etag     pulumi.StringPtrInput
	Role     pulumi.StringPtrInput
	Username pulumi.StringPtrInput
}

func (MembershipState) ElementType() reflect.Type {
	return reflect.TypeOf((*membershipState)(nil)).Elem()
}

type membershipArgs struct {
	Role     *string `pulumi:"role"`
	Username string  `pulumi:"username"`
}

// The set of arguments for constructing a Membership resource.
type MembershipArgs struct {
	Role     pulumi.StringPtrInput
	Username pulumi.StringInput
}

func (MembershipArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*membershipArgs)(nil)).Elem()
}

type MembershipInput interface {
	pulumi.Input

	ToMembershipOutput() MembershipOutput
	ToMembershipOutputWithContext(ctx context.Context) MembershipOutput
}

func (*Membership) ElementType() reflect.Type {
	return reflect.TypeOf((**Membership)(nil)).Elem()
}

func (i *Membership) ToMembershipOutput() MembershipOutput {
	return i.ToMembershipOutputWithContext(context.Background())
}

func (i *Membership) ToMembershipOutputWithContext(ctx context.Context) MembershipOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MembershipOutput)
}

// MembershipArrayInput is an input type that accepts MembershipArray and MembershipArrayOutput values.
// You can construct a concrete instance of `MembershipArrayInput` via:
//
//	MembershipArray{ MembershipArgs{...} }
type MembershipArrayInput interface {
	pulumi.Input

	ToMembershipArrayOutput() MembershipArrayOutput
	ToMembershipArrayOutputWithContext(context.Context) MembershipArrayOutput
}

type MembershipArray []MembershipInput

func (MembershipArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Membership)(nil)).Elem()
}

func (i MembershipArray) ToMembershipArrayOutput() MembershipArrayOutput {
	return i.ToMembershipArrayOutputWithContext(context.Background())
}

func (i MembershipArray) ToMembershipArrayOutputWithContext(ctx context.Context) MembershipArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MembershipArrayOutput)
}

// MembershipMapInput is an input type that accepts MembershipMap and MembershipMapOutput values.
// You can construct a concrete instance of `MembershipMapInput` via:
//
//	MembershipMap{ "key": MembershipArgs{...} }
type MembershipMapInput interface {
	pulumi.Input

	ToMembershipMapOutput() MembershipMapOutput
	ToMembershipMapOutputWithContext(context.Context) MembershipMapOutput
}

type MembershipMap map[string]MembershipInput

func (MembershipMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Membership)(nil)).Elem()
}

func (i MembershipMap) ToMembershipMapOutput() MembershipMapOutput {
	return i.ToMembershipMapOutputWithContext(context.Background())
}

func (i MembershipMap) ToMembershipMapOutputWithContext(ctx context.Context) MembershipMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MembershipMapOutput)
}

type MembershipOutput struct{ *pulumi.OutputState }

func (MembershipOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Membership)(nil)).Elem()
}

func (o MembershipOutput) ToMembershipOutput() MembershipOutput {
	return o
}

func (o MembershipOutput) ToMembershipOutputWithContext(ctx context.Context) MembershipOutput {
	return o
}

func (o MembershipOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Membership) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o MembershipOutput) Role() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Membership) pulumi.StringPtrOutput { return v.Role }).(pulumi.StringPtrOutput)
}

func (o MembershipOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *Membership) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

type MembershipArrayOutput struct{ *pulumi.OutputState }

func (MembershipArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Membership)(nil)).Elem()
}

func (o MembershipArrayOutput) ToMembershipArrayOutput() MembershipArrayOutput {
	return o
}

func (o MembershipArrayOutput) ToMembershipArrayOutputWithContext(ctx context.Context) MembershipArrayOutput {
	return o
}

func (o MembershipArrayOutput) Index(i pulumi.IntInput) MembershipOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Membership {
		return vs[0].([]*Membership)[vs[1].(int)]
	}).(MembershipOutput)
}

type MembershipMapOutput struct{ *pulumi.OutputState }

func (MembershipMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Membership)(nil)).Elem()
}

func (o MembershipMapOutput) ToMembershipMapOutput() MembershipMapOutput {
	return o
}

func (o MembershipMapOutput) ToMembershipMapOutputWithContext(ctx context.Context) MembershipMapOutput {
	return o
}

func (o MembershipMapOutput) MapIndex(k pulumi.StringInput) MembershipOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Membership {
		return vs[0].(map[string]*Membership)[vs[1].(string)]
	}).(MembershipOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MembershipInput)(nil)).Elem(), &Membership{})
	pulumi.RegisterInputType(reflect.TypeOf((*MembershipArrayInput)(nil)).Elem(), MembershipArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MembershipMapInput)(nil)).Elem(), MembershipMap{})
	pulumi.RegisterOutputType(MembershipOutput{})
	pulumi.RegisterOutputType(MembershipArrayOutput{})
	pulumi.RegisterOutputType(MembershipMapOutput{})
}

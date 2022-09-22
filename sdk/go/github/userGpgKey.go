// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type UserGpgKey struct {
	pulumi.CustomResourceState

	ArmoredPublicKey pulumi.StringOutput `pulumi:"armoredPublicKey"`
	Etag             pulumi.StringOutput `pulumi:"etag"`
	KeyId            pulumi.StringOutput `pulumi:"keyId"`
}

// NewUserGpgKey registers a new resource with the given unique name, arguments, and options.
func NewUserGpgKey(ctx *pulumi.Context,
	name string, args *UserGpgKeyArgs, opts ...pulumi.ResourceOption) (*UserGpgKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ArmoredPublicKey == nil {
		return nil, errors.New("invalid value for required argument 'ArmoredPublicKey'")
	}
	var resource UserGpgKey
	err := ctx.RegisterResource("github:index/userGpgKey:UserGpgKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserGpgKey gets an existing UserGpgKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserGpgKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserGpgKeyState, opts ...pulumi.ResourceOption) (*UserGpgKey, error) {
	var resource UserGpgKey
	err := ctx.ReadResource("github:index/userGpgKey:UserGpgKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserGpgKey resources.
type userGpgKeyState struct {
	ArmoredPublicKey *string `pulumi:"armoredPublicKey"`
	Etag             *string `pulumi:"etag"`
	KeyId            *string `pulumi:"keyId"`
}

type UserGpgKeyState struct {
	ArmoredPublicKey pulumi.StringPtrInput
	Etag             pulumi.StringPtrInput
	KeyId            pulumi.StringPtrInput
}

func (UserGpgKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*userGpgKeyState)(nil)).Elem()
}

type userGpgKeyArgs struct {
	ArmoredPublicKey string `pulumi:"armoredPublicKey"`
}

// The set of arguments for constructing a UserGpgKey resource.
type UserGpgKeyArgs struct {
	ArmoredPublicKey pulumi.StringInput
}

func (UserGpgKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userGpgKeyArgs)(nil)).Elem()
}

type UserGpgKeyInput interface {
	pulumi.Input

	ToUserGpgKeyOutput() UserGpgKeyOutput
	ToUserGpgKeyOutputWithContext(ctx context.Context) UserGpgKeyOutput
}

func (*UserGpgKey) ElementType() reflect.Type {
	return reflect.TypeOf((**UserGpgKey)(nil)).Elem()
}

func (i *UserGpgKey) ToUserGpgKeyOutput() UserGpgKeyOutput {
	return i.ToUserGpgKeyOutputWithContext(context.Background())
}

func (i *UserGpgKey) ToUserGpgKeyOutputWithContext(ctx context.Context) UserGpgKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserGpgKeyOutput)
}

// UserGpgKeyArrayInput is an input type that accepts UserGpgKeyArray and UserGpgKeyArrayOutput values.
// You can construct a concrete instance of `UserGpgKeyArrayInput` via:
//
//	UserGpgKeyArray{ UserGpgKeyArgs{...} }
type UserGpgKeyArrayInput interface {
	pulumi.Input

	ToUserGpgKeyArrayOutput() UserGpgKeyArrayOutput
	ToUserGpgKeyArrayOutputWithContext(context.Context) UserGpgKeyArrayOutput
}

type UserGpgKeyArray []UserGpgKeyInput

func (UserGpgKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserGpgKey)(nil)).Elem()
}

func (i UserGpgKeyArray) ToUserGpgKeyArrayOutput() UserGpgKeyArrayOutput {
	return i.ToUserGpgKeyArrayOutputWithContext(context.Background())
}

func (i UserGpgKeyArray) ToUserGpgKeyArrayOutputWithContext(ctx context.Context) UserGpgKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserGpgKeyArrayOutput)
}

// UserGpgKeyMapInput is an input type that accepts UserGpgKeyMap and UserGpgKeyMapOutput values.
// You can construct a concrete instance of `UserGpgKeyMapInput` via:
//
//	UserGpgKeyMap{ "key": UserGpgKeyArgs{...} }
type UserGpgKeyMapInput interface {
	pulumi.Input

	ToUserGpgKeyMapOutput() UserGpgKeyMapOutput
	ToUserGpgKeyMapOutputWithContext(context.Context) UserGpgKeyMapOutput
}

type UserGpgKeyMap map[string]UserGpgKeyInput

func (UserGpgKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserGpgKey)(nil)).Elem()
}

func (i UserGpgKeyMap) ToUserGpgKeyMapOutput() UserGpgKeyMapOutput {
	return i.ToUserGpgKeyMapOutputWithContext(context.Background())
}

func (i UserGpgKeyMap) ToUserGpgKeyMapOutputWithContext(ctx context.Context) UserGpgKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserGpgKeyMapOutput)
}

type UserGpgKeyOutput struct{ *pulumi.OutputState }

func (UserGpgKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserGpgKey)(nil)).Elem()
}

func (o UserGpgKeyOutput) ToUserGpgKeyOutput() UserGpgKeyOutput {
	return o
}

func (o UserGpgKeyOutput) ToUserGpgKeyOutputWithContext(ctx context.Context) UserGpgKeyOutput {
	return o
}

func (o UserGpgKeyOutput) ArmoredPublicKey() pulumi.StringOutput {
	return o.ApplyT(func(v *UserGpgKey) pulumi.StringOutput { return v.ArmoredPublicKey }).(pulumi.StringOutput)
}

func (o UserGpgKeyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *UserGpgKey) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o UserGpgKeyOutput) KeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *UserGpgKey) pulumi.StringOutput { return v.KeyId }).(pulumi.StringOutput)
}

type UserGpgKeyArrayOutput struct{ *pulumi.OutputState }

func (UserGpgKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserGpgKey)(nil)).Elem()
}

func (o UserGpgKeyArrayOutput) ToUserGpgKeyArrayOutput() UserGpgKeyArrayOutput {
	return o
}

func (o UserGpgKeyArrayOutput) ToUserGpgKeyArrayOutputWithContext(ctx context.Context) UserGpgKeyArrayOutput {
	return o
}

func (o UserGpgKeyArrayOutput) Index(i pulumi.IntInput) UserGpgKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserGpgKey {
		return vs[0].([]*UserGpgKey)[vs[1].(int)]
	}).(UserGpgKeyOutput)
}

type UserGpgKeyMapOutput struct{ *pulumi.OutputState }

func (UserGpgKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserGpgKey)(nil)).Elem()
}

func (o UserGpgKeyMapOutput) ToUserGpgKeyMapOutput() UserGpgKeyMapOutput {
	return o
}

func (o UserGpgKeyMapOutput) ToUserGpgKeyMapOutputWithContext(ctx context.Context) UserGpgKeyMapOutput {
	return o
}

func (o UserGpgKeyMapOutput) MapIndex(k pulumi.StringInput) UserGpgKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserGpgKey {
		return vs[0].(map[string]*UserGpgKey)[vs[1].(string)]
	}).(UserGpgKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserGpgKeyInput)(nil)).Elem(), &UserGpgKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserGpgKeyArrayInput)(nil)).Elem(), UserGpgKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserGpgKeyMapInput)(nil)).Elem(), UserGpgKeyMap{})
	pulumi.RegisterOutputType(UserGpgKeyOutput{})
	pulumi.RegisterOutputType(UserGpgKeyArrayOutput{})
	pulumi.RegisterOutputType(UserGpgKeyMapOutput{})
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a GitHub user's SSH key resource.
//
// This resource allows you to add/remove SSH keys from your user account.
//
// ## Import
//
// SSH keys can be imported using their ID e.g.
//
// ```sh
//  $ pulumi import github:index/userSshKey:UserSshKey example 1234567
// ```
type UserSshKey struct {
	pulumi.CustomResourceState

	Etag pulumi.StringOutput `pulumi:"etag"`
	// The public SSH key to add to your GitHub account.
	Key pulumi.StringOutput `pulumi:"key"`
	// A descriptive name for the new key. e.g. `Personal MacBook Air`
	Title pulumi.StringOutput `pulumi:"title"`
	// The URL of the SSH key
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewUserSshKey registers a new resource with the given unique name, arguments, and options.
func NewUserSshKey(ctx *pulumi.Context,
	name string, args *UserSshKeyArgs, opts ...pulumi.ResourceOption) (*UserSshKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Title == nil {
		return nil, errors.New("invalid value for required argument 'Title'")
	}
	var resource UserSshKey
	err := ctx.RegisterResource("github:index/userSshKey:UserSshKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserSshKey gets an existing UserSshKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserSshKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserSshKeyState, opts ...pulumi.ResourceOption) (*UserSshKey, error) {
	var resource UserSshKey
	err := ctx.ReadResource("github:index/userSshKey:UserSshKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserSshKey resources.
type userSshKeyState struct {
	Etag *string `pulumi:"etag"`
	// The public SSH key to add to your GitHub account.
	Key *string `pulumi:"key"`
	// A descriptive name for the new key. e.g. `Personal MacBook Air`
	Title *string `pulumi:"title"`
	// The URL of the SSH key
	Url *string `pulumi:"url"`
}

type UserSshKeyState struct {
	Etag pulumi.StringPtrInput
	// The public SSH key to add to your GitHub account.
	Key pulumi.StringPtrInput
	// A descriptive name for the new key. e.g. `Personal MacBook Air`
	Title pulumi.StringPtrInput
	// The URL of the SSH key
	Url pulumi.StringPtrInput
}

func (UserSshKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*userSshKeyState)(nil)).Elem()
}

type userSshKeyArgs struct {
	// The public SSH key to add to your GitHub account.
	Key string `pulumi:"key"`
	// A descriptive name for the new key. e.g. `Personal MacBook Air`
	Title string `pulumi:"title"`
}

// The set of arguments for constructing a UserSshKey resource.
type UserSshKeyArgs struct {
	// The public SSH key to add to your GitHub account.
	Key pulumi.StringInput
	// A descriptive name for the new key. e.g. `Personal MacBook Air`
	Title pulumi.StringInput
}

func (UserSshKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userSshKeyArgs)(nil)).Elem()
}

type UserSshKeyInput interface {
	pulumi.Input

	ToUserSshKeyOutput() UserSshKeyOutput
	ToUserSshKeyOutputWithContext(ctx context.Context) UserSshKeyOutput
}

func (*UserSshKey) ElementType() reflect.Type {
	return reflect.TypeOf((*UserSshKey)(nil))
}

func (i *UserSshKey) ToUserSshKeyOutput() UserSshKeyOutput {
	return i.ToUserSshKeyOutputWithContext(context.Background())
}

func (i *UserSshKey) ToUserSshKeyOutputWithContext(ctx context.Context) UserSshKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserSshKeyOutput)
}

func (i *UserSshKey) ToUserSshKeyPtrOutput() UserSshKeyPtrOutput {
	return i.ToUserSshKeyPtrOutputWithContext(context.Background())
}

func (i *UserSshKey) ToUserSshKeyPtrOutputWithContext(ctx context.Context) UserSshKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserSshKeyPtrOutput)
}

type UserSshKeyPtrInput interface {
	pulumi.Input

	ToUserSshKeyPtrOutput() UserSshKeyPtrOutput
	ToUserSshKeyPtrOutputWithContext(ctx context.Context) UserSshKeyPtrOutput
}

type userSshKeyPtrType UserSshKeyArgs

func (*userSshKeyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserSshKey)(nil))
}

func (i *userSshKeyPtrType) ToUserSshKeyPtrOutput() UserSshKeyPtrOutput {
	return i.ToUserSshKeyPtrOutputWithContext(context.Background())
}

func (i *userSshKeyPtrType) ToUserSshKeyPtrOutputWithContext(ctx context.Context) UserSshKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserSshKeyPtrOutput)
}

// UserSshKeyArrayInput is an input type that accepts UserSshKeyArray and UserSshKeyArrayOutput values.
// You can construct a concrete instance of `UserSshKeyArrayInput` via:
//
//          UserSshKeyArray{ UserSshKeyArgs{...} }
type UserSshKeyArrayInput interface {
	pulumi.Input

	ToUserSshKeyArrayOutput() UserSshKeyArrayOutput
	ToUserSshKeyArrayOutputWithContext(context.Context) UserSshKeyArrayOutput
}

type UserSshKeyArray []UserSshKeyInput

func (UserSshKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*UserSshKey)(nil))
}

func (i UserSshKeyArray) ToUserSshKeyArrayOutput() UserSshKeyArrayOutput {
	return i.ToUserSshKeyArrayOutputWithContext(context.Background())
}

func (i UserSshKeyArray) ToUserSshKeyArrayOutputWithContext(ctx context.Context) UserSshKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserSshKeyArrayOutput)
}

// UserSshKeyMapInput is an input type that accepts UserSshKeyMap and UserSshKeyMapOutput values.
// You can construct a concrete instance of `UserSshKeyMapInput` via:
//
//          UserSshKeyMap{ "key": UserSshKeyArgs{...} }
type UserSshKeyMapInput interface {
	pulumi.Input

	ToUserSshKeyMapOutput() UserSshKeyMapOutput
	ToUserSshKeyMapOutputWithContext(context.Context) UserSshKeyMapOutput
}

type UserSshKeyMap map[string]UserSshKeyInput

func (UserSshKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*UserSshKey)(nil))
}

func (i UserSshKeyMap) ToUserSshKeyMapOutput() UserSshKeyMapOutput {
	return i.ToUserSshKeyMapOutputWithContext(context.Background())
}

func (i UserSshKeyMap) ToUserSshKeyMapOutputWithContext(ctx context.Context) UserSshKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserSshKeyMapOutput)
}

type UserSshKeyOutput struct {
	*pulumi.OutputState
}

func (UserSshKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserSshKey)(nil))
}

func (o UserSshKeyOutput) ToUserSshKeyOutput() UserSshKeyOutput {
	return o
}

func (o UserSshKeyOutput) ToUserSshKeyOutputWithContext(ctx context.Context) UserSshKeyOutput {
	return o
}

func (o UserSshKeyOutput) ToUserSshKeyPtrOutput() UserSshKeyPtrOutput {
	return o.ToUserSshKeyPtrOutputWithContext(context.Background())
}

func (o UserSshKeyOutput) ToUserSshKeyPtrOutputWithContext(ctx context.Context) UserSshKeyPtrOutput {
	return o.ApplyT(func(v UserSshKey) *UserSshKey {
		return &v
	}).(UserSshKeyPtrOutput)
}

type UserSshKeyPtrOutput struct {
	*pulumi.OutputState
}

func (UserSshKeyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserSshKey)(nil))
}

func (o UserSshKeyPtrOutput) ToUserSshKeyPtrOutput() UserSshKeyPtrOutput {
	return o
}

func (o UserSshKeyPtrOutput) ToUserSshKeyPtrOutputWithContext(ctx context.Context) UserSshKeyPtrOutput {
	return o
}

type UserSshKeyArrayOutput struct{ *pulumi.OutputState }

func (UserSshKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UserSshKey)(nil))
}

func (o UserSshKeyArrayOutput) ToUserSshKeyArrayOutput() UserSshKeyArrayOutput {
	return o
}

func (o UserSshKeyArrayOutput) ToUserSshKeyArrayOutputWithContext(ctx context.Context) UserSshKeyArrayOutput {
	return o
}

func (o UserSshKeyArrayOutput) Index(i pulumi.IntInput) UserSshKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) UserSshKey {
		return vs[0].([]UserSshKey)[vs[1].(int)]
	}).(UserSshKeyOutput)
}

type UserSshKeyMapOutput struct{ *pulumi.OutputState }

func (UserSshKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserSshKey)(nil))
}

func (o UserSshKeyMapOutput) ToUserSshKeyMapOutput() UserSshKeyMapOutput {
	return o
}

func (o UserSshKeyMapOutput) ToUserSshKeyMapOutputWithContext(ctx context.Context) UserSshKeyMapOutput {
	return o
}

func (o UserSshKeyMapOutput) MapIndex(k pulumi.StringInput) UserSshKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserSshKey {
		return vs[0].(map[string]UserSshKey)[vs[1].(string)]
	}).(UserSshKeyOutput)
}

func init() {
	pulumi.RegisterOutputType(UserSshKeyOutput{})
	pulumi.RegisterOutputType(UserSshKeyPtrOutput{})
	pulumi.RegisterOutputType(UserSshKeyArrayOutput{})
	pulumi.RegisterOutputType(UserSshKeyMapOutput{})
}

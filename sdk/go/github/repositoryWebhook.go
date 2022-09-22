// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RepositoryWebhook struct {
	pulumi.CustomResourceState

	Active        pulumi.BoolPtrOutput                    `pulumi:"active"`
	Configuration RepositoryWebhookConfigurationPtrOutput `pulumi:"configuration"`
	Etag          pulumi.StringOutput                     `pulumi:"etag"`
	Events        pulumi.StringArrayOutput                `pulumi:"events"`
	Repository    pulumi.StringOutput                     `pulumi:"repository"`
	Url           pulumi.StringOutput                     `pulumi:"url"`
}

// NewRepositoryWebhook registers a new resource with the given unique name, arguments, and options.
func NewRepositoryWebhook(ctx *pulumi.Context,
	name string, args *RepositoryWebhookArgs, opts ...pulumi.ResourceOption) (*RepositoryWebhook, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Events == nil {
		return nil, errors.New("invalid value for required argument 'Events'")
	}
	if args.Repository == nil {
		return nil, errors.New("invalid value for required argument 'Repository'")
	}
	var resource RepositoryWebhook
	err := ctx.RegisterResource("github:index/repositoryWebhook:RepositoryWebhook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepositoryWebhook gets an existing RepositoryWebhook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepositoryWebhook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryWebhookState, opts ...pulumi.ResourceOption) (*RepositoryWebhook, error) {
	var resource RepositoryWebhook
	err := ctx.ReadResource("github:index/repositoryWebhook:RepositoryWebhook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RepositoryWebhook resources.
type repositoryWebhookState struct {
	Active        *bool                           `pulumi:"active"`
	Configuration *RepositoryWebhookConfiguration `pulumi:"configuration"`
	Etag          *string                         `pulumi:"etag"`
	Events        []string                        `pulumi:"events"`
	Repository    *string                         `pulumi:"repository"`
	Url           *string                         `pulumi:"url"`
}

type RepositoryWebhookState struct {
	Active        pulumi.BoolPtrInput
	Configuration RepositoryWebhookConfigurationPtrInput
	Etag          pulumi.StringPtrInput
	Events        pulumi.StringArrayInput
	Repository    pulumi.StringPtrInput
	Url           pulumi.StringPtrInput
}

func (RepositoryWebhookState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryWebhookState)(nil)).Elem()
}

type repositoryWebhookArgs struct {
	Active        *bool                           `pulumi:"active"`
	Configuration *RepositoryWebhookConfiguration `pulumi:"configuration"`
	Events        []string                        `pulumi:"events"`
	Repository    string                          `pulumi:"repository"`
}

// The set of arguments for constructing a RepositoryWebhook resource.
type RepositoryWebhookArgs struct {
	Active        pulumi.BoolPtrInput
	Configuration RepositoryWebhookConfigurationPtrInput
	Events        pulumi.StringArrayInput
	Repository    pulumi.StringInput
}

func (RepositoryWebhookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryWebhookArgs)(nil)).Elem()
}

type RepositoryWebhookInput interface {
	pulumi.Input

	ToRepositoryWebhookOutput() RepositoryWebhookOutput
	ToRepositoryWebhookOutputWithContext(ctx context.Context) RepositoryWebhookOutput
}

func (*RepositoryWebhook) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryWebhook)(nil)).Elem()
}

func (i *RepositoryWebhook) ToRepositoryWebhookOutput() RepositoryWebhookOutput {
	return i.ToRepositoryWebhookOutputWithContext(context.Background())
}

func (i *RepositoryWebhook) ToRepositoryWebhookOutputWithContext(ctx context.Context) RepositoryWebhookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryWebhookOutput)
}

// RepositoryWebhookArrayInput is an input type that accepts RepositoryWebhookArray and RepositoryWebhookArrayOutput values.
// You can construct a concrete instance of `RepositoryWebhookArrayInput` via:
//
//	RepositoryWebhookArray{ RepositoryWebhookArgs{...} }
type RepositoryWebhookArrayInput interface {
	pulumi.Input

	ToRepositoryWebhookArrayOutput() RepositoryWebhookArrayOutput
	ToRepositoryWebhookArrayOutputWithContext(context.Context) RepositoryWebhookArrayOutput
}

type RepositoryWebhookArray []RepositoryWebhookInput

func (RepositoryWebhookArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryWebhook)(nil)).Elem()
}

func (i RepositoryWebhookArray) ToRepositoryWebhookArrayOutput() RepositoryWebhookArrayOutput {
	return i.ToRepositoryWebhookArrayOutputWithContext(context.Background())
}

func (i RepositoryWebhookArray) ToRepositoryWebhookArrayOutputWithContext(ctx context.Context) RepositoryWebhookArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryWebhookArrayOutput)
}

// RepositoryWebhookMapInput is an input type that accepts RepositoryWebhookMap and RepositoryWebhookMapOutput values.
// You can construct a concrete instance of `RepositoryWebhookMapInput` via:
//
//	RepositoryWebhookMap{ "key": RepositoryWebhookArgs{...} }
type RepositoryWebhookMapInput interface {
	pulumi.Input

	ToRepositoryWebhookMapOutput() RepositoryWebhookMapOutput
	ToRepositoryWebhookMapOutputWithContext(context.Context) RepositoryWebhookMapOutput
}

type RepositoryWebhookMap map[string]RepositoryWebhookInput

func (RepositoryWebhookMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryWebhook)(nil)).Elem()
}

func (i RepositoryWebhookMap) ToRepositoryWebhookMapOutput() RepositoryWebhookMapOutput {
	return i.ToRepositoryWebhookMapOutputWithContext(context.Background())
}

func (i RepositoryWebhookMap) ToRepositoryWebhookMapOutputWithContext(ctx context.Context) RepositoryWebhookMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryWebhookMapOutput)
}

type RepositoryWebhookOutput struct{ *pulumi.OutputState }

func (RepositoryWebhookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryWebhook)(nil)).Elem()
}

func (o RepositoryWebhookOutput) ToRepositoryWebhookOutput() RepositoryWebhookOutput {
	return o
}

func (o RepositoryWebhookOutput) ToRepositoryWebhookOutputWithContext(ctx context.Context) RepositoryWebhookOutput {
	return o
}

func (o RepositoryWebhookOutput) Active() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RepositoryWebhook) pulumi.BoolPtrOutput { return v.Active }).(pulumi.BoolPtrOutput)
}

func (o RepositoryWebhookOutput) Configuration() RepositoryWebhookConfigurationPtrOutput {
	return o.ApplyT(func(v *RepositoryWebhook) RepositoryWebhookConfigurationPtrOutput { return v.Configuration }).(RepositoryWebhookConfigurationPtrOutput)
}

func (o RepositoryWebhookOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryWebhook) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o RepositoryWebhookOutput) Events() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RepositoryWebhook) pulumi.StringArrayOutput { return v.Events }).(pulumi.StringArrayOutput)
}

func (o RepositoryWebhookOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryWebhook) pulumi.StringOutput { return v.Repository }).(pulumi.StringOutput)
}

func (o RepositoryWebhookOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryWebhook) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

type RepositoryWebhookArrayOutput struct{ *pulumi.OutputState }

func (RepositoryWebhookArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryWebhook)(nil)).Elem()
}

func (o RepositoryWebhookArrayOutput) ToRepositoryWebhookArrayOutput() RepositoryWebhookArrayOutput {
	return o
}

func (o RepositoryWebhookArrayOutput) ToRepositoryWebhookArrayOutputWithContext(ctx context.Context) RepositoryWebhookArrayOutput {
	return o
}

func (o RepositoryWebhookArrayOutput) Index(i pulumi.IntInput) RepositoryWebhookOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RepositoryWebhook {
		return vs[0].([]*RepositoryWebhook)[vs[1].(int)]
	}).(RepositoryWebhookOutput)
}

type RepositoryWebhookMapOutput struct{ *pulumi.OutputState }

func (RepositoryWebhookMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryWebhook)(nil)).Elem()
}

func (o RepositoryWebhookMapOutput) ToRepositoryWebhookMapOutput() RepositoryWebhookMapOutput {
	return o
}

func (o RepositoryWebhookMapOutput) ToRepositoryWebhookMapOutputWithContext(ctx context.Context) RepositoryWebhookMapOutput {
	return o
}

func (o RepositoryWebhookMapOutput) MapIndex(k pulumi.StringInput) RepositoryWebhookOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RepositoryWebhook {
		return vs[0].(map[string]*RepositoryWebhook)[vs[1].(string)]
	}).(RepositoryWebhookOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryWebhookInput)(nil)).Elem(), &RepositoryWebhook{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryWebhookArrayInput)(nil)).Elem(), RepositoryWebhookArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryWebhookMapInput)(nil)).Elem(), RepositoryWebhookMap{})
	pulumi.RegisterOutputType(RepositoryWebhookOutput{})
	pulumi.RegisterOutputType(RepositoryWebhookArrayOutput{})
	pulumi.RegisterOutputType(RepositoryWebhookMapOutput{})
}

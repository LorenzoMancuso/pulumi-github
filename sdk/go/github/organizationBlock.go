// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This resource allows you to create and manage blocks for GitHub organizations.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-github/sdk/v3/go/github/"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := github.NewOrganizationBlock(ctx, "example", &github.OrganizationBlockArgs{
// 			Username: pulumi.String("paultyng"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type OrganizationBlock struct {
	pulumi.CustomResourceState

	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the user to block.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewOrganizationBlock registers a new resource with the given unique name, arguments, and options.
func NewOrganizationBlock(ctx *pulumi.Context,
	name string, args *OrganizationBlockArgs, opts ...pulumi.ResourceOption) (*OrganizationBlock, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	var resource OrganizationBlock
	err := ctx.RegisterResource("github:index/organizationBlock:OrganizationBlock", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationBlock gets an existing OrganizationBlock resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationBlock(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationBlockState, opts ...pulumi.ResourceOption) (*OrganizationBlock, error) {
	var resource OrganizationBlock
	err := ctx.ReadResource("github:index/organizationBlock:OrganizationBlock", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationBlock resources.
type organizationBlockState struct {
	Etag *string `pulumi:"etag"`
	// The name of the user to block.
	Username *string `pulumi:"username"`
}

type OrganizationBlockState struct {
	Etag pulumi.StringPtrInput
	// The name of the user to block.
	Username pulumi.StringPtrInput
}

func (OrganizationBlockState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationBlockState)(nil)).Elem()
}

type organizationBlockArgs struct {
	// The name of the user to block.
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a OrganizationBlock resource.
type OrganizationBlockArgs struct {
	// The name of the user to block.
	Username pulumi.StringInput
}

func (OrganizationBlockArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationBlockArgs)(nil)).Elem()
}

type OrganizationBlockInput interface {
	pulumi.Input

	ToOrganizationBlockOutput() OrganizationBlockOutput
	ToOrganizationBlockOutputWithContext(ctx context.Context) OrganizationBlockOutput
}

func (*OrganizationBlock) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationBlock)(nil))
}

func (i *OrganizationBlock) ToOrganizationBlockOutput() OrganizationBlockOutput {
	return i.ToOrganizationBlockOutputWithContext(context.Background())
}

func (i *OrganizationBlock) ToOrganizationBlockOutputWithContext(ctx context.Context) OrganizationBlockOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationBlockOutput)
}

type OrganizationBlockOutput struct {
	*pulumi.OutputState
}

func (OrganizationBlockOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationBlock)(nil))
}

func (o OrganizationBlockOutput) ToOrganizationBlockOutput() OrganizationBlockOutput {
	return o
}

func (o OrganizationBlockOutput) ToOrganizationBlockOutputWithContext(ctx context.Context) OrganizationBlockOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(OrganizationBlockOutput{})
}

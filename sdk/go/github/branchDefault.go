// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a GitHub branch default resource.
//
// This resource allows you to set the default branch for a given repository.
//
// Note that use of this resource is incompatible with the `defaultBranch` option of the `Repository` resource.  Using both will result in plans always showing a diff.
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
// 		example, err := github.NewRepository(ctx, "example", &github.RepositoryArgs{
// 			Description: pulumi.String("My awesome codebase"),
// 			AutoInit:    pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		development, err := github.NewBranch(ctx, "development", &github.BranchArgs{
// 			Repository: example.Name,
// 			Branch:     pulumi.String("development"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = github.NewBranchDefault(ctx, "_default", &github.BranchDefaultArgs{
// 			Repository: example.Name,
// 			Branch:     development.Branch,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// GitHub Branch Defaults can be imported using an ID made up of `repository`, e.g.
//
// ```sh
//  $ pulumi import github:index/branchDefault:BranchDefault branch_default my-repo
// ```
type BranchDefault struct {
	pulumi.CustomResourceState

	// The branch (e.g. `main`)
	Branch pulumi.StringOutput `pulumi:"branch"`
	// The GitHub repository
	Repository pulumi.StringOutput `pulumi:"repository"`
}

// NewBranchDefault registers a new resource with the given unique name, arguments, and options.
func NewBranchDefault(ctx *pulumi.Context,
	name string, args *BranchDefaultArgs, opts ...pulumi.ResourceOption) (*BranchDefault, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Branch == nil {
		return nil, errors.New("invalid value for required argument 'Branch'")
	}
	if args.Repository == nil {
		return nil, errors.New("invalid value for required argument 'Repository'")
	}
	var resource BranchDefault
	err := ctx.RegisterResource("github:index/branchDefault:BranchDefault", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBranchDefault gets an existing BranchDefault resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBranchDefault(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BranchDefaultState, opts ...pulumi.ResourceOption) (*BranchDefault, error) {
	var resource BranchDefault
	err := ctx.ReadResource("github:index/branchDefault:BranchDefault", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BranchDefault resources.
type branchDefaultState struct {
	// The branch (e.g. `main`)
	Branch *string `pulumi:"branch"`
	// The GitHub repository
	Repository *string `pulumi:"repository"`
}

type BranchDefaultState struct {
	// The branch (e.g. `main`)
	Branch pulumi.StringPtrInput
	// The GitHub repository
	Repository pulumi.StringPtrInput
}

func (BranchDefaultState) ElementType() reflect.Type {
	return reflect.TypeOf((*branchDefaultState)(nil)).Elem()
}

type branchDefaultArgs struct {
	// The branch (e.g. `main`)
	Branch string `pulumi:"branch"`
	// The GitHub repository
	Repository string `pulumi:"repository"`
}

// The set of arguments for constructing a BranchDefault resource.
type BranchDefaultArgs struct {
	// The branch (e.g. `main`)
	Branch pulumi.StringInput
	// The GitHub repository
	Repository pulumi.StringInput
}

func (BranchDefaultArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*branchDefaultArgs)(nil)).Elem()
}

type BranchDefaultInput interface {
	pulumi.Input

	ToBranchDefaultOutput() BranchDefaultOutput
	ToBranchDefaultOutputWithContext(ctx context.Context) BranchDefaultOutput
}

func (*BranchDefault) ElementType() reflect.Type {
	return reflect.TypeOf((*BranchDefault)(nil))
}

func (i *BranchDefault) ToBranchDefaultOutput() BranchDefaultOutput {
	return i.ToBranchDefaultOutputWithContext(context.Background())
}

func (i *BranchDefault) ToBranchDefaultOutputWithContext(ctx context.Context) BranchDefaultOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchDefaultOutput)
}

type BranchDefaultOutput struct {
	*pulumi.OutputState
}

func (BranchDefaultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BranchDefault)(nil))
}

func (o BranchDefaultOutput) ToBranchDefaultOutput() BranchDefaultOutput {
	return o
}

func (o BranchDefaultOutput) ToBranchDefaultOutputWithContext(ctx context.Context) BranchDefaultOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BranchDefaultOutput{})
}

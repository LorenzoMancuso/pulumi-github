// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a GitHub repository milestone resource.
//
// This resource allows you to create and manage milestones for a Github Repository within an organization or user account.
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
// 		_, err := github.NewRepositoryMilestone(ctx, "example", &github.RepositoryMilestoneArgs{
// 			Owner:      pulumi.String("example-owner"),
// 			Repository: pulumi.String("example-repository"),
// 			Title:      pulumi.String("v1.1.0"),
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
// A GitHub Repository Milestone can be imported using an ID made up of `owner/repository/number`, e.g.
//
// ```sh
//  $ pulumi import github:index/repositoryMilestone:RepositoryMilestone example example-owner/example-repository/1
// ```
type RepositoryMilestone struct {
	pulumi.CustomResourceState

	// A description of the milestone.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The milestone due date. In `yyyy-mm-dd` format.
	DueDate pulumi.StringPtrOutput `pulumi:"dueDate"`
	// The number of the milestone.
	Number pulumi.IntOutput `pulumi:"number"`
	// The owner of the Github Repository.
	Owner pulumi.StringOutput `pulumi:"owner"`
	// The name of the Github Repository.
	Repository pulumi.StringOutput `pulumi:"repository"`
	// The state of the milestone. Either `open` or `closed`. Default: `open`
	State pulumi.StringPtrOutput `pulumi:"state"`
	// The title of the milestone.
	Title pulumi.StringOutput `pulumi:"title"`
}

// NewRepositoryMilestone registers a new resource with the given unique name, arguments, and options.
func NewRepositoryMilestone(ctx *pulumi.Context,
	name string, args *RepositoryMilestoneArgs, opts ...pulumi.ResourceOption) (*RepositoryMilestone, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Owner == nil {
		return nil, errors.New("invalid value for required argument 'Owner'")
	}
	if args.Repository == nil {
		return nil, errors.New("invalid value for required argument 'Repository'")
	}
	if args.Title == nil {
		return nil, errors.New("invalid value for required argument 'Title'")
	}
	var resource RepositoryMilestone
	err := ctx.RegisterResource("github:index/repositoryMilestone:RepositoryMilestone", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepositoryMilestone gets an existing RepositoryMilestone resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepositoryMilestone(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryMilestoneState, opts ...pulumi.ResourceOption) (*RepositoryMilestone, error) {
	var resource RepositoryMilestone
	err := ctx.ReadResource("github:index/repositoryMilestone:RepositoryMilestone", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RepositoryMilestone resources.
type repositoryMilestoneState struct {
	// A description of the milestone.
	Description *string `pulumi:"description"`
	// The milestone due date. In `yyyy-mm-dd` format.
	DueDate *string `pulumi:"dueDate"`
	// The number of the milestone.
	Number *int `pulumi:"number"`
	// The owner of the Github Repository.
	Owner *string `pulumi:"owner"`
	// The name of the Github Repository.
	Repository *string `pulumi:"repository"`
	// The state of the milestone. Either `open` or `closed`. Default: `open`
	State *string `pulumi:"state"`
	// The title of the milestone.
	Title *string `pulumi:"title"`
}

type RepositoryMilestoneState struct {
	// A description of the milestone.
	Description pulumi.StringPtrInput
	// The milestone due date. In `yyyy-mm-dd` format.
	DueDate pulumi.StringPtrInput
	// The number of the milestone.
	Number pulumi.IntPtrInput
	// The owner of the Github Repository.
	Owner pulumi.StringPtrInput
	// The name of the Github Repository.
	Repository pulumi.StringPtrInput
	// The state of the milestone. Either `open` or `closed`. Default: `open`
	State pulumi.StringPtrInput
	// The title of the milestone.
	Title pulumi.StringPtrInput
}

func (RepositoryMilestoneState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryMilestoneState)(nil)).Elem()
}

type repositoryMilestoneArgs struct {
	// A description of the milestone.
	Description *string `pulumi:"description"`
	// The milestone due date. In `yyyy-mm-dd` format.
	DueDate *string `pulumi:"dueDate"`
	// The owner of the Github Repository.
	Owner string `pulumi:"owner"`
	// The name of the Github Repository.
	Repository string `pulumi:"repository"`
	// The state of the milestone. Either `open` or `closed`. Default: `open`
	State *string `pulumi:"state"`
	// The title of the milestone.
	Title string `pulumi:"title"`
}

// The set of arguments for constructing a RepositoryMilestone resource.
type RepositoryMilestoneArgs struct {
	// A description of the milestone.
	Description pulumi.StringPtrInput
	// The milestone due date. In `yyyy-mm-dd` format.
	DueDate pulumi.StringPtrInput
	// The owner of the Github Repository.
	Owner pulumi.StringInput
	// The name of the Github Repository.
	Repository pulumi.StringInput
	// The state of the milestone. Either `open` or `closed`. Default: `open`
	State pulumi.StringPtrInput
	// The title of the milestone.
	Title pulumi.StringInput
}

func (RepositoryMilestoneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryMilestoneArgs)(nil)).Elem()
}

type RepositoryMilestoneInput interface {
	pulumi.Input

	ToRepositoryMilestoneOutput() RepositoryMilestoneOutput
	ToRepositoryMilestoneOutputWithContext(ctx context.Context) RepositoryMilestoneOutput
}

func (*RepositoryMilestone) ElementType() reflect.Type {
	return reflect.TypeOf((*RepositoryMilestone)(nil))
}

func (i *RepositoryMilestone) ToRepositoryMilestoneOutput() RepositoryMilestoneOutput {
	return i.ToRepositoryMilestoneOutputWithContext(context.Background())
}

func (i *RepositoryMilestone) ToRepositoryMilestoneOutputWithContext(ctx context.Context) RepositoryMilestoneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryMilestoneOutput)
}

type RepositoryMilestoneOutput struct {
	*pulumi.OutputState
}

func (RepositoryMilestoneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RepositoryMilestone)(nil))
}

func (o RepositoryMilestoneOutput) ToRepositoryMilestoneOutput() RepositoryMilestoneOutput {
	return o
}

func (o RepositoryMilestoneOutput) ToRepositoryMilestoneOutputWithContext(ctx context.Context) RepositoryMilestoneOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RepositoryMilestoneOutput{})
}

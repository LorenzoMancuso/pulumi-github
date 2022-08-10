// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-github/sdk/v4/go/github"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := github.NewMembership(ctx, "membershipForSomeUser", &github.MembershipArgs{
//				Username: pulumi.String("SomeUser"),
//				Role:     pulumi.String("member"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = github.NewMembership(ctx, "membershipForAnotherUser", &github.MembershipArgs{
//				Username: pulumi.String("AnotherUser"),
//				Role:     pulumi.String("member"),
//			})
//			if err != nil {
//				return err
//			}
//			someTeam, err := github.NewTeam(ctx, "someTeam", &github.TeamArgs{
//				Description: pulumi.String("Some cool team"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = github.NewTeamMembers(ctx, "someTeamMembers", &github.TeamMembersArgs{
//				TeamId: someTeam.ID(),
//				Members: TeamMembersMemberArray{
//					&TeamMembersMemberArgs{
//						Username: pulumi.String("SomeUser"),
//						Role:     pulumi.String("maintainer"),
//					},
//					&TeamMembersMemberArgs{
//						Username: pulumi.String("AnotherUser"),
//						Role:     pulumi.String("member"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// GitHub Team Membership can be imported using the team ID `teamid`, e.g.
//
// ```sh
//
//	$ pulumi import github:index/teamMembers:TeamMembers some_team 1234567
//
// ```
type TeamMembers struct {
	pulumi.CustomResourceState

	Etag pulumi.StringOutput `pulumi:"etag"`
	// List of team members. See Members below for details.
	Members TeamMembersMemberArrayOutput `pulumi:"members"`
	// The GitHub team id
	TeamId pulumi.StringOutput `pulumi:"teamId"`
}

// NewTeamMembers registers a new resource with the given unique name, arguments, and options.
func NewTeamMembers(ctx *pulumi.Context,
	name string, args *TeamMembersArgs, opts ...pulumi.ResourceOption) (*TeamMembers, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.TeamId == nil {
		return nil, errors.New("invalid value for required argument 'TeamId'")
	}
	var resource TeamMembers
	err := ctx.RegisterResource("github:index/teamMembers:TeamMembers", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTeamMembers gets an existing TeamMembers resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTeamMembers(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TeamMembersState, opts ...pulumi.ResourceOption) (*TeamMembers, error) {
	var resource TeamMembers
	err := ctx.ReadResource("github:index/teamMembers:TeamMembers", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TeamMembers resources.
type teamMembersState struct {
	Etag *string `pulumi:"etag"`
	// List of team members. See Members below for details.
	Members []TeamMembersMember `pulumi:"members"`
	// The GitHub team id
	TeamId *string `pulumi:"teamId"`
}

type TeamMembersState struct {
	Etag pulumi.StringPtrInput
	// List of team members. See Members below for details.
	Members TeamMembersMemberArrayInput
	// The GitHub team id
	TeamId pulumi.StringPtrInput
}

func (TeamMembersState) ElementType() reflect.Type {
	return reflect.TypeOf((*teamMembersState)(nil)).Elem()
}

type teamMembersArgs struct {
	// List of team members. See Members below for details.
	Members []TeamMembersMember `pulumi:"members"`
	// The GitHub team id
	TeamId string `pulumi:"teamId"`
}

// The set of arguments for constructing a TeamMembers resource.
type TeamMembersArgs struct {
	// List of team members. See Members below for details.
	Members TeamMembersMemberArrayInput
	// The GitHub team id
	TeamId pulumi.StringInput
}

func (TeamMembersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*teamMembersArgs)(nil)).Elem()
}

type TeamMembersInput interface {
	pulumi.Input

	ToTeamMembersOutput() TeamMembersOutput
	ToTeamMembersOutputWithContext(ctx context.Context) TeamMembersOutput
}

func (*TeamMembers) ElementType() reflect.Type {
	return reflect.TypeOf((**TeamMembers)(nil)).Elem()
}

func (i *TeamMembers) ToTeamMembersOutput() TeamMembersOutput {
	return i.ToTeamMembersOutputWithContext(context.Background())
}

func (i *TeamMembers) ToTeamMembersOutputWithContext(ctx context.Context) TeamMembersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamMembersOutput)
}

// TeamMembersArrayInput is an input type that accepts TeamMembersArray and TeamMembersArrayOutput values.
// You can construct a concrete instance of `TeamMembersArrayInput` via:
//
//	TeamMembersArray{ TeamMembersArgs{...} }
type TeamMembersArrayInput interface {
	pulumi.Input

	ToTeamMembersArrayOutput() TeamMembersArrayOutput
	ToTeamMembersArrayOutputWithContext(context.Context) TeamMembersArrayOutput
}

type TeamMembersArray []TeamMembersInput

func (TeamMembersArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TeamMembers)(nil)).Elem()
}

func (i TeamMembersArray) ToTeamMembersArrayOutput() TeamMembersArrayOutput {
	return i.ToTeamMembersArrayOutputWithContext(context.Background())
}

func (i TeamMembersArray) ToTeamMembersArrayOutputWithContext(ctx context.Context) TeamMembersArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamMembersArrayOutput)
}

// TeamMembersMapInput is an input type that accepts TeamMembersMap and TeamMembersMapOutput values.
// You can construct a concrete instance of `TeamMembersMapInput` via:
//
//	TeamMembersMap{ "key": TeamMembersArgs{...} }
type TeamMembersMapInput interface {
	pulumi.Input

	ToTeamMembersMapOutput() TeamMembersMapOutput
	ToTeamMembersMapOutputWithContext(context.Context) TeamMembersMapOutput
}

type TeamMembersMap map[string]TeamMembersInput

func (TeamMembersMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TeamMembers)(nil)).Elem()
}

func (i TeamMembersMap) ToTeamMembersMapOutput() TeamMembersMapOutput {
	return i.ToTeamMembersMapOutputWithContext(context.Background())
}

func (i TeamMembersMap) ToTeamMembersMapOutputWithContext(ctx context.Context) TeamMembersMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamMembersMapOutput)
}

type TeamMembersOutput struct{ *pulumi.OutputState }

func (TeamMembersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TeamMembers)(nil)).Elem()
}

func (o TeamMembersOutput) ToTeamMembersOutput() TeamMembersOutput {
	return o
}

func (o TeamMembersOutput) ToTeamMembersOutputWithContext(ctx context.Context) TeamMembersOutput {
	return o
}

func (o TeamMembersOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *TeamMembers) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// List of team members. See Members below for details.
func (o TeamMembersOutput) Members() TeamMembersMemberArrayOutput {
	return o.ApplyT(func(v *TeamMembers) TeamMembersMemberArrayOutput { return v.Members }).(TeamMembersMemberArrayOutput)
}

// The GitHub team id
func (o TeamMembersOutput) TeamId() pulumi.StringOutput {
	return o.ApplyT(func(v *TeamMembers) pulumi.StringOutput { return v.TeamId }).(pulumi.StringOutput)
}

type TeamMembersArrayOutput struct{ *pulumi.OutputState }

func (TeamMembersArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TeamMembers)(nil)).Elem()
}

func (o TeamMembersArrayOutput) ToTeamMembersArrayOutput() TeamMembersArrayOutput {
	return o
}

func (o TeamMembersArrayOutput) ToTeamMembersArrayOutputWithContext(ctx context.Context) TeamMembersArrayOutput {
	return o
}

func (o TeamMembersArrayOutput) Index(i pulumi.IntInput) TeamMembersOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TeamMembers {
		return vs[0].([]*TeamMembers)[vs[1].(int)]
	}).(TeamMembersOutput)
}

type TeamMembersMapOutput struct{ *pulumi.OutputState }

func (TeamMembersMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TeamMembers)(nil)).Elem()
}

func (o TeamMembersMapOutput) ToTeamMembersMapOutput() TeamMembersMapOutput {
	return o
}

func (o TeamMembersMapOutput) ToTeamMembersMapOutputWithContext(ctx context.Context) TeamMembersMapOutput {
	return o
}

func (o TeamMembersMapOutput) MapIndex(k pulumi.StringInput) TeamMembersOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TeamMembers {
		return vs[0].(map[string]*TeamMembers)[vs[1].(string)]
	}).(TeamMembersOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TeamMembersInput)(nil)).Elem(), &TeamMembers{})
	pulumi.RegisterInputType(reflect.TypeOf((*TeamMembersArrayInput)(nil)).Elem(), TeamMembersArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TeamMembersMapInput)(nil)).Elem(), TeamMembersMap{})
	pulumi.RegisterOutputType(TeamMembersOutput{})
	pulumi.RegisterOutputType(TeamMembersArrayOutput{})
	pulumi.RegisterOutputType(TeamMembersMapOutput{})
}

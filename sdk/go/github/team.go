// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Team struct {
	pulumi.CustomResourceState

	CreateDefaultMaintainer pulumi.BoolPtrOutput   `pulumi:"createDefaultMaintainer"`
	Description             pulumi.StringPtrOutput `pulumi:"description"`
	Etag                    pulumi.StringOutput    `pulumi:"etag"`
	LdapDn                  pulumi.StringPtrOutput `pulumi:"ldapDn"`
	MembersCount            pulumi.IntOutput       `pulumi:"membersCount"`
	Name                    pulumi.StringOutput    `pulumi:"name"`
	NodeId                  pulumi.StringOutput    `pulumi:"nodeId"`
	ParentTeamId            pulumi.IntPtrOutput    `pulumi:"parentTeamId"`
	Privacy                 pulumi.StringPtrOutput `pulumi:"privacy"`
	Slug                    pulumi.StringOutput    `pulumi:"slug"`
}

// NewTeam registers a new resource with the given unique name, arguments, and options.
func NewTeam(ctx *pulumi.Context,
	name string, args *TeamArgs, opts ...pulumi.ResourceOption) (*Team, error) {
	if args == nil {
		args = &TeamArgs{}
	}

	var resource Team
	err := ctx.RegisterResource("github:index/team:Team", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTeam gets an existing Team resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTeam(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TeamState, opts ...pulumi.ResourceOption) (*Team, error) {
	var resource Team
	err := ctx.ReadResource("github:index/team:Team", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Team resources.
type teamState struct {
	CreateDefaultMaintainer *bool   `pulumi:"createDefaultMaintainer"`
	Description             *string `pulumi:"description"`
	Etag                    *string `pulumi:"etag"`
	LdapDn                  *string `pulumi:"ldapDn"`
	MembersCount            *int    `pulumi:"membersCount"`
	Name                    *string `pulumi:"name"`
	NodeId                  *string `pulumi:"nodeId"`
	ParentTeamId            *int    `pulumi:"parentTeamId"`
	Privacy                 *string `pulumi:"privacy"`
	Slug                    *string `pulumi:"slug"`
}

type TeamState struct {
	CreateDefaultMaintainer pulumi.BoolPtrInput
	Description             pulumi.StringPtrInput
	Etag                    pulumi.StringPtrInput
	LdapDn                  pulumi.StringPtrInput
	MembersCount            pulumi.IntPtrInput
	Name                    pulumi.StringPtrInput
	NodeId                  pulumi.StringPtrInput
	ParentTeamId            pulumi.IntPtrInput
	Privacy                 pulumi.StringPtrInput
	Slug                    pulumi.StringPtrInput
}

func (TeamState) ElementType() reflect.Type {
	return reflect.TypeOf((*teamState)(nil)).Elem()
}

type teamArgs struct {
	CreateDefaultMaintainer *bool   `pulumi:"createDefaultMaintainer"`
	Description             *string `pulumi:"description"`
	LdapDn                  *string `pulumi:"ldapDn"`
	Name                    *string `pulumi:"name"`
	ParentTeamId            *int    `pulumi:"parentTeamId"`
	Privacy                 *string `pulumi:"privacy"`
}

// The set of arguments for constructing a Team resource.
type TeamArgs struct {
	CreateDefaultMaintainer pulumi.BoolPtrInput
	Description             pulumi.StringPtrInput
	LdapDn                  pulumi.StringPtrInput
	Name                    pulumi.StringPtrInput
	ParentTeamId            pulumi.IntPtrInput
	Privacy                 pulumi.StringPtrInput
}

func (TeamArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*teamArgs)(nil)).Elem()
}

type TeamInput interface {
	pulumi.Input

	ToTeamOutput() TeamOutput
	ToTeamOutputWithContext(ctx context.Context) TeamOutput
}

func (*Team) ElementType() reflect.Type {
	return reflect.TypeOf((**Team)(nil)).Elem()
}

func (i *Team) ToTeamOutput() TeamOutput {
	return i.ToTeamOutputWithContext(context.Background())
}

func (i *Team) ToTeamOutputWithContext(ctx context.Context) TeamOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamOutput)
}

// TeamArrayInput is an input type that accepts TeamArray and TeamArrayOutput values.
// You can construct a concrete instance of `TeamArrayInput` via:
//
//          TeamArray{ TeamArgs{...} }
type TeamArrayInput interface {
	pulumi.Input

	ToTeamArrayOutput() TeamArrayOutput
	ToTeamArrayOutputWithContext(context.Context) TeamArrayOutput
}

type TeamArray []TeamInput

func (TeamArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Team)(nil)).Elem()
}

func (i TeamArray) ToTeamArrayOutput() TeamArrayOutput {
	return i.ToTeamArrayOutputWithContext(context.Background())
}

func (i TeamArray) ToTeamArrayOutputWithContext(ctx context.Context) TeamArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamArrayOutput)
}

// TeamMapInput is an input type that accepts TeamMap and TeamMapOutput values.
// You can construct a concrete instance of `TeamMapInput` via:
//
//          TeamMap{ "key": TeamArgs{...} }
type TeamMapInput interface {
	pulumi.Input

	ToTeamMapOutput() TeamMapOutput
	ToTeamMapOutputWithContext(context.Context) TeamMapOutput
}

type TeamMap map[string]TeamInput

func (TeamMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Team)(nil)).Elem()
}

func (i TeamMap) ToTeamMapOutput() TeamMapOutput {
	return i.ToTeamMapOutputWithContext(context.Background())
}

func (i TeamMap) ToTeamMapOutputWithContext(ctx context.Context) TeamMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamMapOutput)
}

type TeamOutput struct{ *pulumi.OutputState }

func (TeamOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Team)(nil)).Elem()
}

func (o TeamOutput) ToTeamOutput() TeamOutput {
	return o
}

func (o TeamOutput) ToTeamOutputWithContext(ctx context.Context) TeamOutput {
	return o
}

func (o TeamOutput) CreateDefaultMaintainer() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Team) pulumi.BoolPtrOutput { return v.CreateDefaultMaintainer }).(pulumi.BoolPtrOutput)
}

func (o TeamOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Team) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o TeamOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Team) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o TeamOutput) LdapDn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Team) pulumi.StringPtrOutput { return v.LdapDn }).(pulumi.StringPtrOutput)
}

func (o TeamOutput) MembersCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Team) pulumi.IntOutput { return v.MembersCount }).(pulumi.IntOutput)
}

func (o TeamOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Team) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o TeamOutput) NodeId() pulumi.StringOutput {
	return o.ApplyT(func(v *Team) pulumi.StringOutput { return v.NodeId }).(pulumi.StringOutput)
}

func (o TeamOutput) ParentTeamId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Team) pulumi.IntPtrOutput { return v.ParentTeamId }).(pulumi.IntPtrOutput)
}

func (o TeamOutput) Privacy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Team) pulumi.StringPtrOutput { return v.Privacy }).(pulumi.StringPtrOutput)
}

func (o TeamOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v *Team) pulumi.StringOutput { return v.Slug }).(pulumi.StringOutput)
}

type TeamArrayOutput struct{ *pulumi.OutputState }

func (TeamArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Team)(nil)).Elem()
}

func (o TeamArrayOutput) ToTeamArrayOutput() TeamArrayOutput {
	return o
}

func (o TeamArrayOutput) ToTeamArrayOutputWithContext(ctx context.Context) TeamArrayOutput {
	return o
}

func (o TeamArrayOutput) Index(i pulumi.IntInput) TeamOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Team {
		return vs[0].([]*Team)[vs[1].(int)]
	}).(TeamOutput)
}

type TeamMapOutput struct{ *pulumi.OutputState }

func (TeamMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Team)(nil)).Elem()
}

func (o TeamMapOutput) ToTeamMapOutput() TeamMapOutput {
	return o
}

func (o TeamMapOutput) ToTeamMapOutputWithContext(ctx context.Context) TeamMapOutput {
	return o
}

func (o TeamMapOutput) MapIndex(k pulumi.StringInput) TeamOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Team {
		return vs[0].(map[string]*Team)[vs[1].(string)]
	}).(TeamOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TeamInput)(nil)).Elem(), &Team{})
	pulumi.RegisterInputType(reflect.TypeOf((*TeamArrayInput)(nil)).Elem(), TeamArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TeamMapInput)(nil)).Elem(), TeamMap{})
	pulumi.RegisterOutputType(TeamOutput{})
	pulumi.RegisterOutputType(TeamArrayOutput{})
	pulumi.RegisterOutputType(TeamMapOutput{})
}

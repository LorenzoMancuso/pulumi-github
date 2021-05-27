// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AppAuth struct {
	Id             string `pulumi:"id"`
	InstallationId string `pulumi:"installationId"`
	PemFile        string `pulumi:"pemFile"`
}

// AppAuthInput is an input type that accepts AppAuthArgs and AppAuthOutput values.
// You can construct a concrete instance of `AppAuthInput` via:
//
//          AppAuthArgs{...}
type AppAuthInput interface {
	pulumi.Input

	ToAppAuthOutput() AppAuthOutput
	ToAppAuthOutputWithContext(context.Context) AppAuthOutput
}

type AppAuthArgs struct {
	Id             pulumi.StringInput `pulumi:"id"`
	InstallationId pulumi.StringInput `pulumi:"installationId"`
	PemFile        pulumi.StringInput `pulumi:"pemFile"`
}

func (AppAuthArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AppAuth)(nil)).Elem()
}

func (i AppAuthArgs) ToAppAuthOutput() AppAuthOutput {
	return i.ToAppAuthOutputWithContext(context.Background())
}

func (i AppAuthArgs) ToAppAuthOutputWithContext(ctx context.Context) AppAuthOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppAuthOutput)
}

type AppAuthOutput struct{ *pulumi.OutputState }

func (AppAuthOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppAuth)(nil)).Elem()
}

func (o AppAuthOutput) ToAppAuthOutput() AppAuthOutput {
	return o
}

func (o AppAuthOutput) ToAppAuthOutputWithContext(ctx context.Context) AppAuthOutput {
	return o
}

func (o AppAuthOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v AppAuth) string { return v.Id }).(pulumi.StringOutput)
}

func (o AppAuthOutput) InstallationId() pulumi.StringOutput {
	return o.ApplyT(func(v AppAuth) string { return v.InstallationId }).(pulumi.StringOutput)
}

func (o AppAuthOutput) PemFile() pulumi.StringOutput {
	return o.ApplyT(func(v AppAuth) string { return v.PemFile }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AppAuthOutput{})
}

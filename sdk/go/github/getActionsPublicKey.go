// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to retrieve information about a GitHub Actions public key. This data source is required to be used with other GitHub secrets interactions.
// Note that the provider `token` must have admin rights to a repository to retrieve it's action public key.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-github/sdk/v3/go/github"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := github.GetActionsPublicKey(ctx, &github.GetActionsPublicKeyArgs{
// 			Repository: "example_repo",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetActionsPublicKey(ctx *pulumi.Context, args *GetActionsPublicKeyArgs, opts ...pulumi.InvokeOption) (*GetActionsPublicKeyResult, error) {
	var rv GetActionsPublicKeyResult
	err := ctx.Invoke("github:index/getActionsPublicKey:getActionsPublicKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getActionsPublicKey.
type GetActionsPublicKeyArgs struct {
	// Name of the repository to get public key from.
	Repository string `pulumi:"repository"`
}

// A collection of values returned by getActionsPublicKey.
type GetActionsPublicKeyResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Actual key retrieved.
	Key string `pulumi:"key"`
	// ID of the key that has been retrieved.
	KeyId      string `pulumi:"keyId"`
	Repository string `pulumi:"repository"`
}

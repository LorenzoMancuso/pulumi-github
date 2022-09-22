// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    public static class GetActionsPublicKey
    {
        public static Task<GetActionsPublicKeyResult> InvokeAsync(GetActionsPublicKeyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetActionsPublicKeyResult>("github:index/getActionsPublicKey:getActionsPublicKey", args ?? new GetActionsPublicKeyArgs(), options.WithDefaults());

        public static Output<GetActionsPublicKeyResult> Invoke(GetActionsPublicKeyInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetActionsPublicKeyResult>("github:index/getActionsPublicKey:getActionsPublicKey", args ?? new GetActionsPublicKeyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetActionsPublicKeyArgs : global::Pulumi.InvokeArgs
    {
        [Input("repository", required: true)]
        public string Repository { get; set; } = null!;

        public GetActionsPublicKeyArgs()
        {
        }
        public static new GetActionsPublicKeyArgs Empty => new GetActionsPublicKeyArgs();
    }

    public sealed class GetActionsPublicKeyInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("repository", required: true)]
        public Input<string> Repository { get; set; } = null!;

        public GetActionsPublicKeyInvokeArgs()
        {
        }
        public static new GetActionsPublicKeyInvokeArgs Empty => new GetActionsPublicKeyInvokeArgs();
    }


    [OutputType]
    public sealed class GetActionsPublicKeyResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Key;
        public readonly string KeyId;
        public readonly string Repository;

        [OutputConstructor]
        private GetActionsPublicKeyResult(
            string id,

            string key,

            string keyId,

            string repository)
        {
            Id = id;
            Key = key;
            KeyId = keyId;
            Repository = repository;
        }
    }
}

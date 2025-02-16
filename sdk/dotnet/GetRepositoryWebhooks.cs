// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    public static class GetRepositoryWebhooks
    {
        public static Task<GetRepositoryWebhooksResult> InvokeAsync(GetRepositoryWebhooksArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRepositoryWebhooksResult>("github:index/getRepositoryWebhooks:getRepositoryWebhooks", args ?? new GetRepositoryWebhooksArgs(), options.WithDefaults());

        public static Output<GetRepositoryWebhooksResult> Invoke(GetRepositoryWebhooksInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRepositoryWebhooksResult>("github:index/getRepositoryWebhooks:getRepositoryWebhooks", args ?? new GetRepositoryWebhooksInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRepositoryWebhooksArgs : global::Pulumi.InvokeArgs
    {
        [Input("repository", required: true)]
        public string Repository { get; set; } = null!;

        public GetRepositoryWebhooksArgs()
        {
        }
        public static new GetRepositoryWebhooksArgs Empty => new GetRepositoryWebhooksArgs();
    }

    public sealed class GetRepositoryWebhooksInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("repository", required: true)]
        public Input<string> Repository { get; set; } = null!;

        public GetRepositoryWebhooksInvokeArgs()
        {
        }
        public static new GetRepositoryWebhooksInvokeArgs Empty => new GetRepositoryWebhooksInvokeArgs();
    }


    [OutputType]
    public sealed class GetRepositoryWebhooksResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Repository;
        public readonly ImmutableArray<Outputs.GetRepositoryWebhooksWebhookResult> Webhooks;

        [OutputConstructor]
        private GetRepositoryWebhooksResult(
            string id,

            string repository,

            ImmutableArray<Outputs.GetRepositoryWebhooksWebhookResult> webhooks)
        {
            Id = id;
            Repository = repository;
            Webhooks = webhooks;
        }
    }
}

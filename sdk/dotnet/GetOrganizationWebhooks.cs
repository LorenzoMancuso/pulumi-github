// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    public static class GetOrganizationWebhooks
    {
        public static Task<GetOrganizationWebhooksResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOrganizationWebhooksResult>("github:index/getOrganizationWebhooks:getOrganizationWebhooks", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetOrganizationWebhooksResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetOrganizationWebhooksWebhookResult> Webhooks;

        [OutputConstructor]
        private GetOrganizationWebhooksResult(
            string id,

            ImmutableArray<Outputs.GetOrganizationWebhooksWebhookResult> webhooks)
        {
            Id = id;
            Webhooks = webhooks;
        }
    }
}

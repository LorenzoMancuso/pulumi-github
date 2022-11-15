// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github.Outputs
{

    [OutputType]
    public sealed class GetRepositoryWebhooksWebhookResult
    {
        public readonly bool Active;
        public readonly int Id;
        public readonly string Name;
        public readonly string Type;
        public readonly string Url;

        [OutputConstructor]
        private GetRepositoryWebhooksWebhookResult(
            bool active,

            int id,

            string name,

            string type,

            string url)
        {
            Active = active;
            Id = id;
            Name = name;
            Type = type;
            Url = url;
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    public static class GetGithubApp
    {
        public static Task<GetGithubAppResult> InvokeAsync(GetGithubAppArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGithubAppResult>("github:index/getGithubApp:getGithubApp", args ?? new GetGithubAppArgs(), options.WithDefaults());

        public static Output<GetGithubAppResult> Invoke(GetGithubAppInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGithubAppResult>("github:index/getGithubApp:getGithubApp", args ?? new GetGithubAppInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGithubAppArgs : global::Pulumi.InvokeArgs
    {
        [Input("slug", required: true)]
        public string Slug { get; set; } = null!;

        public GetGithubAppArgs()
        {
        }
        public static new GetGithubAppArgs Empty => new GetGithubAppArgs();
    }

    public sealed class GetGithubAppInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("slug", required: true)]
        public Input<string> Slug { get; set; } = null!;

        public GetGithubAppInvokeArgs()
        {
        }
        public static new GetGithubAppInvokeArgs Empty => new GetGithubAppInvokeArgs();
    }


    [OutputType]
    public sealed class GetGithubAppResult
    {
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string NodeId;
        public readonly string Slug;

        [OutputConstructor]
        private GetGithubAppResult(
            string description,

            string id,

            string name,

            string nodeId,

            string slug)
        {
            Description = description;
            Id = id;
            Name = name;
            NodeId = nodeId;
            Slug = slug;
        }
    }
}
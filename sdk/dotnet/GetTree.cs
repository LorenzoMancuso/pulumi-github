// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    public static class GetTree
    {
        public static Task<GetTreeResult> InvokeAsync(GetTreeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTreeResult>("github:index/getTree:getTree", args ?? new GetTreeArgs(), options.WithDefaults());

        public static Output<GetTreeResult> Invoke(GetTreeInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetTreeResult>("github:index/getTree:getTree", args ?? new GetTreeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTreeArgs : global::Pulumi.InvokeArgs
    {
        [Input("recursive")]
        public bool? Recursive { get; set; }

        [Input("repository", required: true)]
        public string Repository { get; set; } = null!;

        [Input("treeSha", required: true)]
        public string TreeSha { get; set; } = null!;

        public GetTreeArgs()
        {
        }
        public static new GetTreeArgs Empty => new GetTreeArgs();
    }

    public sealed class GetTreeInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("recursive")]
        public Input<bool>? Recursive { get; set; }

        [Input("repository", required: true)]
        public Input<string> Repository { get; set; } = null!;

        [Input("treeSha", required: true)]
        public Input<string> TreeSha { get; set; } = null!;

        public GetTreeInvokeArgs()
        {
        }
        public static new GetTreeInvokeArgs Empty => new GetTreeInvokeArgs();
    }


    [OutputType]
    public sealed class GetTreeResult
    {
        public readonly ImmutableArray<Outputs.GetTreeEntryResult> Entries;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool? Recursive;
        public readonly string Repository;
        public readonly string TreeSha;

        [OutputConstructor]
        private GetTreeResult(
            ImmutableArray<Outputs.GetTreeEntryResult> entries,

            string id,

            bool? recursive,

            string repository,

            string treeSha)
        {
            Entries = entries;
            Id = id;
            Recursive = recursive;
            Repository = repository;
            TreeSha = treeSha;
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    public static class GetRepositories
    {
        public static Task<GetRepositoriesResult> InvokeAsync(GetRepositoriesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRepositoriesResult>("github:index/getRepositories:getRepositories", args ?? new GetRepositoriesArgs(), options.WithDefaults());

        public static Output<GetRepositoriesResult> Invoke(GetRepositoriesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRepositoriesResult>("github:index/getRepositories:getRepositories", args ?? new GetRepositoriesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRepositoriesArgs : global::Pulumi.InvokeArgs
    {
        [Input("includeRepoId")]
        public bool? IncludeRepoId { get; set; }

        [Input("query", required: true)]
        public string Query { get; set; } = null!;

        [Input("resultsPerPage")]
        public int? ResultsPerPage { get; set; }

        [Input("sort")]
        public string? Sort { get; set; }

        public GetRepositoriesArgs()
        {
        }
        public static new GetRepositoriesArgs Empty => new GetRepositoriesArgs();
    }

    public sealed class GetRepositoriesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("includeRepoId")]
        public Input<bool>? IncludeRepoId { get; set; }

        [Input("query", required: true)]
        public Input<string> Query { get; set; } = null!;

        [Input("resultsPerPage")]
        public Input<int>? ResultsPerPage { get; set; }

        [Input("sort")]
        public Input<string>? Sort { get; set; }

        public GetRepositoriesInvokeArgs()
        {
        }
        public static new GetRepositoriesInvokeArgs Empty => new GetRepositoriesInvokeArgs();
    }


    [OutputType]
    public sealed class GetRepositoriesResult
    {
        public readonly ImmutableArray<string> FullNames;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool? IncludeRepoId;
        public readonly ImmutableArray<string> Names;
        public readonly string Query;
        public readonly ImmutableArray<int> RepoIds;
        public readonly int? ResultsPerPage;
        public readonly string? Sort;

        [OutputConstructor]
        private GetRepositoriesResult(
            ImmutableArray<string> fullNames,

            string id,

            bool? includeRepoId,

            ImmutableArray<string> names,

            string query,

            ImmutableArray<int> repoIds,

            int? resultsPerPage,

            string? sort)
        {
            FullNames = fullNames;
            Id = id;
            IncludeRepoId = includeRepoId;
            Names = names;
            Query = query;
            RepoIds = repoIds;
            ResultsPerPage = resultsPerPage;
            Sort = sort;
        }
    }
}

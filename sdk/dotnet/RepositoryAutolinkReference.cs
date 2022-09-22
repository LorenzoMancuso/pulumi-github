// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    [GithubResourceType("github:index/repositoryAutolinkReference:RepositoryAutolinkReference")]
    public partial class RepositoryAutolinkReference : global::Pulumi.CustomResource
    {
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// This prefix appended by a number will generate a link any time it is found in an issue, pull request, or commit
        /// </summary>
        [Output("keyPrefix")]
        public Output<string> KeyPrefix { get; private set; } = null!;

        /// <summary>
        /// The repository name
        /// </summary>
        [Output("repository")]
        public Output<string> Repository { get; private set; } = null!;

        /// <summary>
        /// The template of the target URL used for the links; must be a valid URL and contain `&lt;num&gt;` for the reference number
        /// </summary>
        [Output("targetUrlTemplate")]
        public Output<string> TargetUrlTemplate { get; private set; } = null!;


        /// <summary>
        /// Create a RepositoryAutolinkReference resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RepositoryAutolinkReference(string name, RepositoryAutolinkReferenceArgs args, CustomResourceOptions? options = null)
            : base("github:index/repositoryAutolinkReference:RepositoryAutolinkReference", name, args ?? new RepositoryAutolinkReferenceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RepositoryAutolinkReference(string name, Input<string> id, RepositoryAutolinkReferenceState? state = null, CustomResourceOptions? options = null)
            : base("github:index/repositoryAutolinkReference:RepositoryAutolinkReference", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing RepositoryAutolinkReference resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RepositoryAutolinkReference Get(string name, Input<string> id, RepositoryAutolinkReferenceState? state = null, CustomResourceOptions? options = null)
        {
            return new RepositoryAutolinkReference(name, id, state, options);
        }
    }

    public sealed class RepositoryAutolinkReferenceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This prefix appended by a number will generate a link any time it is found in an issue, pull request, or commit
        /// </summary>
        [Input("keyPrefix", required: true)]
        public Input<string> KeyPrefix { get; set; } = null!;

        /// <summary>
        /// The repository name
        /// </summary>
        [Input("repository", required: true)]
        public Input<string> Repository { get; set; } = null!;

        /// <summary>
        /// The template of the target URL used for the links; must be a valid URL and contain `&lt;num&gt;` for the reference number
        /// </summary>
        [Input("targetUrlTemplate", required: true)]
        public Input<string> TargetUrlTemplate { get; set; } = null!;

        public RepositoryAutolinkReferenceArgs()
        {
        }
        public static new RepositoryAutolinkReferenceArgs Empty => new RepositoryAutolinkReferenceArgs();
    }

    public sealed class RepositoryAutolinkReferenceState : global::Pulumi.ResourceArgs
    {
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// This prefix appended by a number will generate a link any time it is found in an issue, pull request, or commit
        /// </summary>
        [Input("keyPrefix")]
        public Input<string>? KeyPrefix { get; set; }

        /// <summary>
        /// The repository name
        /// </summary>
        [Input("repository")]
        public Input<string>? Repository { get; set; }

        /// <summary>
        /// The template of the target URL used for the links; must be a valid URL and contain `&lt;num&gt;` for the reference number
        /// </summary>
        [Input("targetUrlTemplate")]
        public Input<string>? TargetUrlTemplate { get; set; }

        public RepositoryAutolinkReferenceState()
        {
        }
        public static new RepositoryAutolinkReferenceState Empty => new RepositoryAutolinkReferenceState();
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    [GithubResourceType("github:index/repositoryTagProtection:RepositoryTagProtection")]
    public partial class RepositoryTagProtection : global::Pulumi.CustomResource
    {
        [Output("pattern")]
        public Output<string> Pattern { get; private set; } = null!;

        [Output("repository")]
        public Output<string> Repository { get; private set; } = null!;

        [Output("tagProtectionId")]
        public Output<int> TagProtectionId { get; private set; } = null!;


        /// <summary>
        /// Create a RepositoryTagProtection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RepositoryTagProtection(string name, RepositoryTagProtectionArgs args, CustomResourceOptions? options = null)
            : base("github:index/repositoryTagProtection:RepositoryTagProtection", name, args ?? new RepositoryTagProtectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RepositoryTagProtection(string name, Input<string> id, RepositoryTagProtectionState? state = null, CustomResourceOptions? options = null)
            : base("github:index/repositoryTagProtection:RepositoryTagProtection", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RepositoryTagProtection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RepositoryTagProtection Get(string name, Input<string> id, RepositoryTagProtectionState? state = null, CustomResourceOptions? options = null)
        {
            return new RepositoryTagProtection(name, id, state, options);
        }
    }

    public sealed class RepositoryTagProtectionArgs : global::Pulumi.ResourceArgs
    {
        [Input("pattern", required: true)]
        public Input<string> Pattern { get; set; } = null!;

        [Input("repository", required: true)]
        public Input<string> Repository { get; set; } = null!;

        public RepositoryTagProtectionArgs()
        {
        }
        public static new RepositoryTagProtectionArgs Empty => new RepositoryTagProtectionArgs();
    }

    public sealed class RepositoryTagProtectionState : global::Pulumi.ResourceArgs
    {
        [Input("pattern")]
        public Input<string>? Pattern { get; set; }

        [Input("repository")]
        public Input<string>? Repository { get; set; }

        [Input("tagProtectionId")]
        public Input<int>? TagProtectionId { get; set; }

        public RepositoryTagProtectionState()
        {
        }
        public static new RepositoryTagProtectionState Empty => new RepositoryTagProtectionState();
    }
}

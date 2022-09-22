// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    [GithubResourceType("github:index/organizationBlock:OrganizationBlock")]
    public partial class OrganizationBlock : global::Pulumi.CustomResource
    {
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        [Output("username")]
        public Output<string> Username { get; private set; } = null!;


        /// <summary>
        /// Create a OrganizationBlock resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OrganizationBlock(string name, OrganizationBlockArgs args, CustomResourceOptions? options = null)
            : base("github:index/organizationBlock:OrganizationBlock", name, args ?? new OrganizationBlockArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OrganizationBlock(string name, Input<string> id, OrganizationBlockState? state = null, CustomResourceOptions? options = null)
            : base("github:index/organizationBlock:OrganizationBlock", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing OrganizationBlock resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OrganizationBlock Get(string name, Input<string> id, OrganizationBlockState? state = null, CustomResourceOptions? options = null)
        {
            return new OrganizationBlock(name, id, state, options);
        }
    }

    public sealed class OrganizationBlockArgs : global::Pulumi.ResourceArgs
    {
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public OrganizationBlockArgs()
        {
        }
        public static new OrganizationBlockArgs Empty => new OrganizationBlockArgs();
    }

    public sealed class OrganizationBlockState : global::Pulumi.ResourceArgs
    {
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        [Input("username")]
        public Input<string>? Username { get; set; }

        public OrganizationBlockState()
        {
        }
        public static new OrganizationBlockState Empty => new OrganizationBlockState();
    }
}

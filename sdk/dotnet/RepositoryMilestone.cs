// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    /// <summary>
    /// Provides a GitHub repository milestone resource.
    /// 
    /// This resource allows you to create and manage milestones for a Github Repository within an organization or user account.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Github = Pulumi.Github;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         // Create a milestone for a repository
    ///         var example = new Github.RepositoryMilestone("example", new Github.RepositoryMilestoneArgs
    ///         {
    ///             Owner = "example-owner",
    ///             Repository = "example-repository",
    ///             Title = "v1.1.0",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// A GitHub Repository Milestone can be imported using an ID made up of `owner/repository/number`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import github:index/repositoryMilestone:RepositoryMilestone example example-owner/example-repository/1
    /// ```
    /// </summary>
    public partial class RepositoryMilestone : Pulumi.CustomResource
    {
        /// <summary>
        /// A description of the milestone.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The milestone due date. In `yyyy-mm-dd` format.
        /// </summary>
        [Output("dueDate")]
        public Output<string?> DueDate { get; private set; } = null!;

        /// <summary>
        /// The number of the milestone.
        /// </summary>
        [Output("number")]
        public Output<int> Number { get; private set; } = null!;

        /// <summary>
        /// The owner of the Github Repository.
        /// </summary>
        [Output("owner")]
        public Output<string> Owner { get; private set; } = null!;

        /// <summary>
        /// The name of the Github Repository.
        /// </summary>
        [Output("repository")]
        public Output<string> Repository { get; private set; } = null!;

        /// <summary>
        /// The state of the milestone. Either `open` or `closed`. Default: `open`
        /// </summary>
        [Output("state")]
        public Output<string?> State { get; private set; } = null!;

        /// <summary>
        /// The title of the milestone.
        /// </summary>
        [Output("title")]
        public Output<string> Title { get; private set; } = null!;


        /// <summary>
        /// Create a RepositoryMilestone resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RepositoryMilestone(string name, RepositoryMilestoneArgs args, CustomResourceOptions? options = null)
            : base("github:index/repositoryMilestone:RepositoryMilestone", name, args ?? new RepositoryMilestoneArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RepositoryMilestone(string name, Input<string> id, RepositoryMilestoneState? state = null, CustomResourceOptions? options = null)
            : base("github:index/repositoryMilestone:RepositoryMilestone", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RepositoryMilestone resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RepositoryMilestone Get(string name, Input<string> id, RepositoryMilestoneState? state = null, CustomResourceOptions? options = null)
        {
            return new RepositoryMilestone(name, id, state, options);
        }
    }

    public sealed class RepositoryMilestoneArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A description of the milestone.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The milestone due date. In `yyyy-mm-dd` format.
        /// </summary>
        [Input("dueDate")]
        public Input<string>? DueDate { get; set; }

        /// <summary>
        /// The owner of the Github Repository.
        /// </summary>
        [Input("owner", required: true)]
        public Input<string> Owner { get; set; } = null!;

        /// <summary>
        /// The name of the Github Repository.
        /// </summary>
        [Input("repository", required: true)]
        public Input<string> Repository { get; set; } = null!;

        /// <summary>
        /// The state of the milestone. Either `open` or `closed`. Default: `open`
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// The title of the milestone.
        /// </summary>
        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        public RepositoryMilestoneArgs()
        {
        }
    }

    public sealed class RepositoryMilestoneState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A description of the milestone.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The milestone due date. In `yyyy-mm-dd` format.
        /// </summary>
        [Input("dueDate")]
        public Input<string>? DueDate { get; set; }

        /// <summary>
        /// The number of the milestone.
        /// </summary>
        [Input("number")]
        public Input<int>? Number { get; set; }

        /// <summary>
        /// The owner of the Github Repository.
        /// </summary>
        [Input("owner")]
        public Input<string>? Owner { get; set; }

        /// <summary>
        /// The name of the Github Repository.
        /// </summary>
        [Input("repository")]
        public Input<string>? Repository { get; set; }

        /// <summary>
        /// The state of the milestone. Either `open` or `closed`. Default: `open`
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// The title of the milestone.
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        public RepositoryMilestoneState()
        {
        }
    }
}

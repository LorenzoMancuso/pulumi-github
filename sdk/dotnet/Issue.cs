// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    [GithubResourceType("github:index/issue:Issue")]
    public partial class Issue : global::Pulumi.CustomResource
    {
        /// <summary>
        /// List of Logins for Users to assign to this issue
        /// </summary>
        [Output("assignees")]
        public Output<ImmutableArray<string>> Assignees { get; private set; } = null!;

        [Output("body")]
        public Output<string?> Body { get; private set; } = null!;

        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        [Output("issueId")]
        public Output<int> IssueId { get; private set; } = null!;

        /// <summary>
        /// List of names of labels on the issue
        /// </summary>
        [Output("labels")]
        public Output<ImmutableArray<string>> Labels { get; private set; } = null!;

        [Output("milestoneNumber")]
        public Output<int?> MilestoneNumber { get; private set; } = null!;

        [Output("number")]
        public Output<int> Number { get; private set; } = null!;

        [Output("repository")]
        public Output<string> Repository { get; private set; } = null!;

        [Output("title")]
        public Output<string> Title { get; private set; } = null!;


        /// <summary>
        /// Create a Issue resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Issue(string name, IssueArgs args, CustomResourceOptions? options = null)
            : base("github:index/issue:Issue", name, args ?? new IssueArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Issue(string name, Input<string> id, IssueState? state = null, CustomResourceOptions? options = null)
            : base("github:index/issue:Issue", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Issue resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Issue Get(string name, Input<string> id, IssueState? state = null, CustomResourceOptions? options = null)
        {
            return new Issue(name, id, state, options);
        }
    }

    public sealed class IssueArgs : global::Pulumi.ResourceArgs
    {
        [Input("assignees")]
        private InputList<string>? _assignees;

        /// <summary>
        /// List of Logins for Users to assign to this issue
        /// </summary>
        public InputList<string> Assignees
        {
            get => _assignees ?? (_assignees = new InputList<string>());
            set => _assignees = value;
        }

        [Input("body")]
        public Input<string>? Body { get; set; }

        [Input("labels")]
        private InputList<string>? _labels;

        /// <summary>
        /// List of names of labels on the issue
        /// </summary>
        public InputList<string> Labels
        {
            get => _labels ?? (_labels = new InputList<string>());
            set => _labels = value;
        }

        [Input("milestoneNumber")]
        public Input<int>? MilestoneNumber { get; set; }

        [Input("repository", required: true)]
        public Input<string> Repository { get; set; } = null!;

        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        public IssueArgs()
        {
        }
        public static new IssueArgs Empty => new IssueArgs();
    }

    public sealed class IssueState : global::Pulumi.ResourceArgs
    {
        [Input("assignees")]
        private InputList<string>? _assignees;

        /// <summary>
        /// List of Logins for Users to assign to this issue
        /// </summary>
        public InputList<string> Assignees
        {
            get => _assignees ?? (_assignees = new InputList<string>());
            set => _assignees = value;
        }

        [Input("body")]
        public Input<string>? Body { get; set; }

        [Input("etag")]
        public Input<string>? Etag { get; set; }

        [Input("issueId")]
        public Input<int>? IssueId { get; set; }

        [Input("labels")]
        private InputList<string>? _labels;

        /// <summary>
        /// List of names of labels on the issue
        /// </summary>
        public InputList<string> Labels
        {
            get => _labels ?? (_labels = new InputList<string>());
            set => _labels = value;
        }

        [Input("milestoneNumber")]
        public Input<int>? MilestoneNumber { get; set; }

        [Input("number")]
        public Input<int>? Number { get; set; }

        [Input("repository")]
        public Input<string>? Repository { get; set; }

        [Input("title")]
        public Input<string>? Title { get; set; }

        public IssueState()
        {
        }
        public static new IssueState Empty => new IssueState();
    }
}

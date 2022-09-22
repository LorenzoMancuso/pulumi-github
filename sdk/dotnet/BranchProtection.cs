// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    [GithubResourceType("github:index/branchProtection:BranchProtection")]
    public partial class BranchProtection : global::Pulumi.CustomResource
    {
        [Output("allowsDeletions")]
        public Output<bool?> AllowsDeletions { get; private set; } = null!;

        [Output("allowsForcePushes")]
        public Output<bool?> AllowsForcePushes { get; private set; } = null!;

        [Output("blocksCreations")]
        public Output<bool?> BlocksCreations { get; private set; } = null!;

        [Output("enforceAdmins")]
        public Output<bool?> EnforceAdmins { get; private set; } = null!;

        [Output("pattern")]
        public Output<string> Pattern { get; private set; } = null!;

        [Output("pushRestrictions")]
        public Output<ImmutableArray<string>> PushRestrictions { get; private set; } = null!;

        /// <summary>
        /// Node ID or name of repository
        /// </summary>
        [Output("repositoryId")]
        public Output<string> RepositoryId { get; private set; } = null!;

        [Output("requireConversationResolution")]
        public Output<bool?> RequireConversationResolution { get; private set; } = null!;

        [Output("requireSignedCommits")]
        public Output<bool?> RequireSignedCommits { get; private set; } = null!;

        [Output("requiredLinearHistory")]
        public Output<bool?> RequiredLinearHistory { get; private set; } = null!;

        [Output("requiredPullRequestReviews")]
        public Output<ImmutableArray<Outputs.BranchProtectionRequiredPullRequestReview>> RequiredPullRequestReviews { get; private set; } = null!;

        [Output("requiredStatusChecks")]
        public Output<ImmutableArray<Outputs.BranchProtectionRequiredStatusCheck>> RequiredStatusChecks { get; private set; } = null!;


        /// <summary>
        /// Create a BranchProtection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BranchProtection(string name, BranchProtectionArgs args, CustomResourceOptions? options = null)
            : base("github:index/branchProtection:BranchProtection", name, args ?? new BranchProtectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BranchProtection(string name, Input<string> id, BranchProtectionState? state = null, CustomResourceOptions? options = null)
            : base("github:index/branchProtection:BranchProtection", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BranchProtection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BranchProtection Get(string name, Input<string> id, BranchProtectionState? state = null, CustomResourceOptions? options = null)
        {
            return new BranchProtection(name, id, state, options);
        }
    }

    public sealed class BranchProtectionArgs : global::Pulumi.ResourceArgs
    {
        [Input("allowsDeletions")]
        public Input<bool>? AllowsDeletions { get; set; }

        [Input("allowsForcePushes")]
        public Input<bool>? AllowsForcePushes { get; set; }

        [Input("blocksCreations")]
        public Input<bool>? BlocksCreations { get; set; }

        [Input("enforceAdmins")]
        public Input<bool>? EnforceAdmins { get; set; }

        [Input("pattern", required: true)]
        public Input<string> Pattern { get; set; } = null!;

        [Input("pushRestrictions")]
        private InputList<string>? _pushRestrictions;
        public InputList<string> PushRestrictions
        {
            get => _pushRestrictions ?? (_pushRestrictions = new InputList<string>());
            set => _pushRestrictions = value;
        }

        /// <summary>
        /// Node ID or name of repository
        /// </summary>
        [Input("repositoryId", required: true)]
        public Input<string> RepositoryId { get; set; } = null!;

        [Input("requireConversationResolution")]
        public Input<bool>? RequireConversationResolution { get; set; }

        [Input("requireSignedCommits")]
        public Input<bool>? RequireSignedCommits { get; set; }

        [Input("requiredLinearHistory")]
        public Input<bool>? RequiredLinearHistory { get; set; }

        [Input("requiredPullRequestReviews")]
        private InputList<Inputs.BranchProtectionRequiredPullRequestReviewArgs>? _requiredPullRequestReviews;
        public InputList<Inputs.BranchProtectionRequiredPullRequestReviewArgs> RequiredPullRequestReviews
        {
            get => _requiredPullRequestReviews ?? (_requiredPullRequestReviews = new InputList<Inputs.BranchProtectionRequiredPullRequestReviewArgs>());
            set => _requiredPullRequestReviews = value;
        }

        [Input("requiredStatusChecks")]
        private InputList<Inputs.BranchProtectionRequiredStatusCheckArgs>? _requiredStatusChecks;
        public InputList<Inputs.BranchProtectionRequiredStatusCheckArgs> RequiredStatusChecks
        {
            get => _requiredStatusChecks ?? (_requiredStatusChecks = new InputList<Inputs.BranchProtectionRequiredStatusCheckArgs>());
            set => _requiredStatusChecks = value;
        }

        public BranchProtectionArgs()
        {
        }
        public static new BranchProtectionArgs Empty => new BranchProtectionArgs();
    }

    public sealed class BranchProtectionState : global::Pulumi.ResourceArgs
    {
        [Input("allowsDeletions")]
        public Input<bool>? AllowsDeletions { get; set; }

        [Input("allowsForcePushes")]
        public Input<bool>? AllowsForcePushes { get; set; }

        [Input("blocksCreations")]
        public Input<bool>? BlocksCreations { get; set; }

        [Input("enforceAdmins")]
        public Input<bool>? EnforceAdmins { get; set; }

        [Input("pattern")]
        public Input<string>? Pattern { get; set; }

        [Input("pushRestrictions")]
        private InputList<string>? _pushRestrictions;
        public InputList<string> PushRestrictions
        {
            get => _pushRestrictions ?? (_pushRestrictions = new InputList<string>());
            set => _pushRestrictions = value;
        }

        /// <summary>
        /// Node ID or name of repository
        /// </summary>
        [Input("repositoryId")]
        public Input<string>? RepositoryId { get; set; }

        [Input("requireConversationResolution")]
        public Input<bool>? RequireConversationResolution { get; set; }

        [Input("requireSignedCommits")]
        public Input<bool>? RequireSignedCommits { get; set; }

        [Input("requiredLinearHistory")]
        public Input<bool>? RequiredLinearHistory { get; set; }

        [Input("requiredPullRequestReviews")]
        private InputList<Inputs.BranchProtectionRequiredPullRequestReviewGetArgs>? _requiredPullRequestReviews;
        public InputList<Inputs.BranchProtectionRequiredPullRequestReviewGetArgs> RequiredPullRequestReviews
        {
            get => _requiredPullRequestReviews ?? (_requiredPullRequestReviews = new InputList<Inputs.BranchProtectionRequiredPullRequestReviewGetArgs>());
            set => _requiredPullRequestReviews = value;
        }

        [Input("requiredStatusChecks")]
        private InputList<Inputs.BranchProtectionRequiredStatusCheckGetArgs>? _requiredStatusChecks;
        public InputList<Inputs.BranchProtectionRequiredStatusCheckGetArgs> RequiredStatusChecks
        {
            get => _requiredStatusChecks ?? (_requiredStatusChecks = new InputList<Inputs.BranchProtectionRequiredStatusCheckGetArgs>());
            set => _requiredStatusChecks = value;
        }

        public BranchProtectionState()
        {
        }
        public static new BranchProtectionState Empty => new BranchProtectionState();
    }
}

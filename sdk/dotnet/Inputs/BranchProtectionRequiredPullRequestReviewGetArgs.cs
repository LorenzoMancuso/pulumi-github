// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github.Inputs
{

    public sealed class BranchProtectionRequiredPullRequestReviewGetArgs : Pulumi.ResourceArgs
    {
        [Input("dismissStaleReviews")]
        public Input<bool>? DismissStaleReviews { get; set; }

        [Input("dismissalRestrictions")]
        private InputList<string>? _dismissalRestrictions;
        public InputList<string> DismissalRestrictions
        {
            get => _dismissalRestrictions ?? (_dismissalRestrictions = new InputList<string>());
            set => _dismissalRestrictions = value;
        }

        [Input("requireCodeOwnerReviews")]
        public Input<bool>? RequireCodeOwnerReviews { get; set; }

        [Input("requiredApprovingReviewCount")]
        public Input<int>? RequiredApprovingReviewCount { get; set; }

        public BranchProtectionRequiredPullRequestReviewGetArgs()
        {
        }
    }
}
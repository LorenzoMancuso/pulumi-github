// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github.Inputs
{

    public sealed class ActionsOrganizationPermissionsAllowedActionsConfigGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("githubOwnedAllowed", required: true)]
        public Input<bool> GithubOwnedAllowed { get; set; } = null!;

        [Input("patternsAlloweds")]
        private InputList<string>? _patternsAlloweds;
        public InputList<string> PatternsAlloweds
        {
            get => _patternsAlloweds ?? (_patternsAlloweds = new InputList<string>());
            set => _patternsAlloweds = value;
        }

        [Input("verifiedAllowed")]
        public Input<bool>? VerifiedAllowed { get; set; }

        public ActionsOrganizationPermissionsAllowedActionsConfigGetArgs()
        {
        }
        public static new ActionsOrganizationPermissionsAllowedActionsConfigGetArgs Empty => new ActionsOrganizationPermissionsAllowedActionsConfigGetArgs();
    }
}

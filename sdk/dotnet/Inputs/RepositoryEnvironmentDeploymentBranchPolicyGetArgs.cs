// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github.Inputs
{

    public sealed class RepositoryEnvironmentDeploymentBranchPolicyGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether only branches that match the specified name patterns can deploy to this environment.
        /// </summary>
        [Input("customBranchPolicies", required: true)]
        public Input<bool> CustomBranchPolicies { get; set; } = null!;

        /// <summary>
        /// Whether only branches with branch protection rules can deploy to this environment.
        /// </summary>
        [Input("protectedBranches", required: true)]
        public Input<bool> ProtectedBranches { get; set; } = null!;

        public RepositoryEnvironmentDeploymentBranchPolicyGetArgs()
        {
        }
        public static new RepositoryEnvironmentDeploymentBranchPolicyGetArgs Empty => new RepositoryEnvironmentDeploymentBranchPolicyGetArgs();
    }
}

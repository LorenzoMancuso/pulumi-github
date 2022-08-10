// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github.Inputs
{

    public sealed class ActionsOrganizationPermissionsEnabledRepositoriesConfigGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("repositoryIds", required: true)]
        private InputList<int>? _repositoryIds;

        /// <summary>
        /// List of repository IDs to enable for GitHub Actions.
        /// </summary>
        public InputList<int> RepositoryIds
        {
            get => _repositoryIds ?? (_repositoryIds = new InputList<int>());
            set => _repositoryIds = value;
        }

        public ActionsOrganizationPermissionsEnabledRepositoriesConfigGetArgs()
        {
        }
        public static new ActionsOrganizationPermissionsEnabledRepositoriesConfigGetArgs Empty => new ActionsOrganizationPermissionsEnabledRepositoriesConfigGetArgs();
    }
}

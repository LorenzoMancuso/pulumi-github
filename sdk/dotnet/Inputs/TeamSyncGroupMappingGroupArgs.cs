// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github.Inputs
{

    public sealed class TeamSyncGroupMappingGroupArgs : global::Pulumi.ResourceArgs
    {
        [Input("groupDescription", required: true)]
        public Input<string> GroupDescription { get; set; } = null!;

        [Input("groupId", required: true)]
        public Input<string> GroupId { get; set; } = null!;

        [Input("groupName", required: true)]
        public Input<string> GroupName { get; set; } = null!;

        public TeamSyncGroupMappingGroupArgs()
        {
        }
        public static new TeamSyncGroupMappingGroupArgs Empty => new TeamSyncGroupMappingGroupArgs();
    }
}

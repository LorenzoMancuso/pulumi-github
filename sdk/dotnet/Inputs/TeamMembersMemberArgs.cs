// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github.Inputs
{

    public sealed class TeamMembersMemberArgs : global::Pulumi.ResourceArgs
    {
        [Input("role")]
        public Input<string>? Role { get; set; }

        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public TeamMembersMemberArgs()
        {
        }
        public static new TeamMembersMemberArgs Empty => new TeamMembersMemberArgs();
    }
}

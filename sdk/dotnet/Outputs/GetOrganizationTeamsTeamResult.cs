// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github.Outputs
{

    [OutputType]
    public sealed class GetOrganizationTeamsTeamResult
    {
        public readonly string Description;
        public readonly int Id;
        public readonly ImmutableArray<string> Members;
        public readonly string Name;
        public readonly string NodeId;
        public readonly string Privacy;
        public readonly ImmutableArray<string> Repositories;
        public readonly string Slug;

        [OutputConstructor]
        private GetOrganizationTeamsTeamResult(
            string description,

            int id,

            ImmutableArray<string> members,

            string name,

            string nodeId,

            string privacy,

            ImmutableArray<string> repositories,

            string slug)
        {
            Description = description;
            Id = id;
            Members = members;
            Name = name;
            NodeId = nodeId;
            Privacy = privacy;
            Repositories = repositories;
            Slug = slug;
        }
    }
}

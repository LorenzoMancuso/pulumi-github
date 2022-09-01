// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    public static class GetTeam
    {
        /// <summary>
        /// Use this data source to retrieve information about a GitHub team.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Github = Pulumi.Github;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Github.GetTeam.Invoke(new()
        ///     {
        ///         Slug = "example",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetTeamResult> InvokeAsync(GetTeamArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTeamResult>("github:index/getTeam:getTeam", args ?? new GetTeamArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about a GitHub team.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Github = Pulumi.Github;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Github.GetTeam.Invoke(new()
        ///     {
        ///         Slug = "example",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetTeamResult> Invoke(GetTeamInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetTeamResult>("github:index/getTeam:getTeam", args ?? new GetTeamInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTeamArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Type of membershp to be requested to fill the list of members. Can be either "all" or "immediate". Default: "all"
        /// </summary>
        [Input("membershipType")]
        public string? MembershipType { get; set; }

        /// <summary>
        /// The team slug.
        /// </summary>
        [Input("slug", required: true)]
        public string Slug { get; set; } = null!;

        public GetTeamArgs()
        {
        }
        public static new GetTeamArgs Empty => new GetTeamArgs();
    }

    public sealed class GetTeamInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Type of membershp to be requested to fill the list of members. Can be either "all" or "immediate". Default: "all"
        /// </summary>
        [Input("membershipType")]
        public Input<string>? MembershipType { get; set; }

        /// <summary>
        /// The team slug.
        /// </summary>
        [Input("slug", required: true)]
        public Input<string> Slug { get; set; } = null!;

        public GetTeamInvokeArgs()
        {
        }
        public static new GetTeamInvokeArgs Empty => new GetTeamInvokeArgs();
    }


    [OutputType]
    public sealed class GetTeamResult
    {
        /// <summary>
        /// the team's description.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// List of team members (list of GitHub usernames)
        /// </summary>
        public readonly ImmutableArray<string> Members;
        public readonly string? MembershipType;
        /// <summary>
        /// the team's full name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// the Node ID of the team.
        /// </summary>
        public readonly string NodeId;
        /// <summary>
        /// the team's permission level.
        /// </summary>
        public readonly string Permission;
        /// <summary>
        /// the team's privacy type.
        /// </summary>
        public readonly string Privacy;
        /// <summary>
        /// List of team repositories (list of repo names)
        /// </summary>
        public readonly ImmutableArray<string> Repositories;
        public readonly string Slug;

        [OutputConstructor]
        private GetTeamResult(
            string description,

            string id,

            ImmutableArray<string> members,

            string? membershipType,

            string name,

            string nodeId,

            string permission,

            string privacy,

            ImmutableArray<string> repositories,

            string slug)
        {
            Description = description;
            Id = id;
            Members = members;
            MembershipType = membershipType;
            Name = name;
            NodeId = nodeId;
            Permission = permission;
            Privacy = privacy;
            Repositories = repositories;
            Slug = slug;
        }
    }
}

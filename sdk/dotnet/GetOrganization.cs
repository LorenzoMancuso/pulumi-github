// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    public static class GetOrganization
    {
        /// <summary>
        /// Use this data source to retrieve basic information about a GitHub Organization.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Github = Pulumi.Github;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var test = Output.Create(Github.GetOrganization.InvokeAsync(new Github.GetOrganizationArgs
        ///         {
        ///             Name = "github",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetOrganizationResult> InvokeAsync(GetOrganizationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetOrganizationResult>("github:index/getOrganization:getOrganization", args ?? new GetOrganizationArgs(), options.WithVersion());
    }


    public sealed class GetOrganizationArgs : Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetOrganizationArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetOrganizationResult
    {
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Login;
        public readonly string Name;
        public readonly string NodeId;
        /// <summary>
        /// The plan name for the organization account
        /// </summary>
        public readonly string Plan;
        /// <summary>
        /// (`list`) A list with the repositories on the organization
        /// </summary>
        public readonly ImmutableArray<string> Repositories;

        [OutputConstructor]
        private GetOrganizationResult(
            string description,

            string id,

            string login,

            string name,

            string nodeId,

            string plan,

            ImmutableArray<string> repositories)
        {
            Description = description;
            Id = id;
            Login = login;
            Name = name;
            NodeId = nodeId;
            Plan = plan;
            Repositories = repositories;
        }
    }
}

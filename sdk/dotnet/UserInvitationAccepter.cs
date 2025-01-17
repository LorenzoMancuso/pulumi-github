// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    [GithubResourceType("github:index/userInvitationAccepter:UserInvitationAccepter")]
    public partial class UserInvitationAccepter : global::Pulumi.CustomResource
    {
        [Output("invitationId")]
        public Output<string> InvitationId { get; private set; } = null!;


        /// <summary>
        /// Create a UserInvitationAccepter resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UserInvitationAccepter(string name, UserInvitationAccepterArgs args, CustomResourceOptions? options = null)
            : base("github:index/userInvitationAccepter:UserInvitationAccepter", name, args ?? new UserInvitationAccepterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UserInvitationAccepter(string name, Input<string> id, UserInvitationAccepterState? state = null, CustomResourceOptions? options = null)
            : base("github:index/userInvitationAccepter:UserInvitationAccepter", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UserInvitationAccepter resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UserInvitationAccepter Get(string name, Input<string> id, UserInvitationAccepterState? state = null, CustomResourceOptions? options = null)
        {
            return new UserInvitationAccepter(name, id, state, options);
        }
    }

    public sealed class UserInvitationAccepterArgs : global::Pulumi.ResourceArgs
    {
        [Input("invitationId", required: true)]
        public Input<string> InvitationId { get; set; } = null!;

        public UserInvitationAccepterArgs()
        {
        }
        public static new UserInvitationAccepterArgs Empty => new UserInvitationAccepterArgs();
    }

    public sealed class UserInvitationAccepterState : global::Pulumi.ResourceArgs
    {
        [Input("invitationId")]
        public Input<string>? InvitationId { get; set; }

        public UserInvitationAccepterState()
        {
        }
        public static new UserInvitationAccepterState Empty => new UserInvitationAccepterState();
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a GitHub repository collaborator resource.
 *
 * This resource allows you to add/remove collaborators from repositories in your
 * organization or personal account. For organization repositories, collaborators can
 * have explicit (and differing levels of) read, write, or administrator access to
 * specific repositories, without giving the user full organization membership.
 * For personal repositories, collaborators can only be granted write
 * (implictly includes read) permission.
 *
 * When applied, an invitation will be sent to the user to become a collaborator
 * on a repository. When destroyed, either the invitation will be cancelled or the
 * collaborator will be removed from the repository.
 *
 * Further documentation on GitHub collaborators:
 *
 * - [Adding outside collaborators to your personal repositories](https://help.github.com/en/github/setting-up-and-managing-your-github-user-account/managing-access-to-your-personal-repositories)
 * - [Adding outside collaborators to repositories in your organization](https://help.github.com/articles/adding-outside-collaborators-to-repositories-in-your-organization/)
 * - [Converting an organization member to an outside collaborator](https://help.github.com/articles/converting-an-organization-member-to-an-outside-collaborator/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * // Add a collaborator to a repository
 * const aRepoCollaborator = new github.RepositoryCollaborator("a_repo_collaborator", {
 *     permission: "admin",
 *     repository: "our-cool-repo",
 *     username: "SomeUser",
 * });
 * ```
 *
 * ## Import
 *
 * GitHub Repository Collaborators can be imported using an ID made up of `repository:username`, e.g.
 *
 * ```sh
 *  $ pulumi import github:index/repositoryCollaborator:RepositoryCollaborator collaborator terraform:someuser
 * ```
 */
export class RepositoryCollaborator extends pulumi.CustomResource {
    /**
     * Get an existing RepositoryCollaborator resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RepositoryCollaboratorState, opts?: pulumi.CustomResourceOptions): RepositoryCollaborator {
        return new RepositoryCollaborator(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:index/repositoryCollaborator:RepositoryCollaborator';

    /**
     * Returns true if the given object is an instance of RepositoryCollaborator.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RepositoryCollaborator {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RepositoryCollaborator.__pulumiType;
    }

    /**
     * ID of the invitation to be used in `github.UserInvitationAccepter`
     */
    public /*out*/ readonly invitationId!: pulumi.Output<string>;
    /**
     * The permission of the outside collaborator for the repository.
     * Must be one of `pull`, `push`, `maintain`, `triage` or `admin` for organization-owned repositories.
     * Must be `push` for personal repositories. Defaults to `push`.
     */
    public readonly permission!: pulumi.Output<string | undefined>;
    /**
     * Suppress plan diffs for `triage` and `maintain`.  Defaults to `false`.
     */
    public readonly permissionDiffSuppression!: pulumi.Output<boolean | undefined>;
    /**
     * The GitHub repository
     */
    public readonly repository!: pulumi.Output<string>;
    /**
     * The user to add to the repository as a collaborator.
     */
    public readonly username!: pulumi.Output<string>;

    /**
     * Create a RepositoryCollaborator resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RepositoryCollaboratorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RepositoryCollaboratorArgs | RepositoryCollaboratorState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as RepositoryCollaboratorState | undefined;
            inputs["invitationId"] = state ? state.invitationId : undefined;
            inputs["permission"] = state ? state.permission : undefined;
            inputs["permissionDiffSuppression"] = state ? state.permissionDiffSuppression : undefined;
            inputs["repository"] = state ? state.repository : undefined;
            inputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as RepositoryCollaboratorArgs | undefined;
            if ((!args || args.repository === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'repository'");
            }
            if ((!args || args.username === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'username'");
            }
            inputs["permission"] = args ? args.permission : undefined;
            inputs["permissionDiffSuppression"] = args ? args.permissionDiffSuppression : undefined;
            inputs["repository"] = args ? args.repository : undefined;
            inputs["username"] = args ? args.username : undefined;
            inputs["invitationId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(RepositoryCollaborator.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RepositoryCollaborator resources.
 */
export interface RepositoryCollaboratorState {
    /**
     * ID of the invitation to be used in `github.UserInvitationAccepter`
     */
    readonly invitationId?: pulumi.Input<string>;
    /**
     * The permission of the outside collaborator for the repository.
     * Must be one of `pull`, `push`, `maintain`, `triage` or `admin` for organization-owned repositories.
     * Must be `push` for personal repositories. Defaults to `push`.
     */
    readonly permission?: pulumi.Input<string>;
    /**
     * Suppress plan diffs for `triage` and `maintain`.  Defaults to `false`.
     */
    readonly permissionDiffSuppression?: pulumi.Input<boolean>;
    /**
     * The GitHub repository
     */
    readonly repository?: pulumi.Input<string>;
    /**
     * The user to add to the repository as a collaborator.
     */
    readonly username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RepositoryCollaborator resource.
 */
export interface RepositoryCollaboratorArgs {
    /**
     * The permission of the outside collaborator for the repository.
     * Must be one of `pull`, `push`, `maintain`, `triage` or `admin` for organization-owned repositories.
     * Must be `push` for personal repositories. Defaults to `push`.
     */
    readonly permission?: pulumi.Input<string>;
    /**
     * Suppress plan diffs for `triage` and `maintain`.  Defaults to `false`.
     */
    readonly permissionDiffSuppression?: pulumi.Input<boolean>;
    /**
     * The GitHub repository
     */
    readonly repository: pulumi.Input<string>;
    /**
     * The user to add to the repository as a collaborator.
     */
    readonly username: pulumi.Input<string>;
}

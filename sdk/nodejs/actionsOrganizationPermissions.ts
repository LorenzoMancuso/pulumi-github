// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * This resource allows you to create and manage GitHub Actions permissions within your GitHub enterprise organizations.
 * You must have admin access to an organization to use this resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const example = new github.Repository("example", {});
 * const test = new github.ActionsOrganizationPermissions("test", {
 *     allowedActions: "selected",
 *     enabledRepositories: "selected",
 *     allowedActionsConfig: {
 *         githubOwnedAllowed: true,
 *         patternsAlloweds: [
 *             "actions/cache@*",
 *             "actions/checkout@*",
 *         ],
 *         verifiedAllowed: true,
 *     },
 *     enabledRepositoriesConfig: {
 *         repositoryIds: [example.repoId],
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * This resource can be imported using the ID of the GitHub organization
 *
 * ```sh
 *  $ pulumi import github:index/actionsOrganizationPermissions:ActionsOrganizationPermissions test <github_organization_name>
 * ```
 */
export class ActionsOrganizationPermissions extends pulumi.CustomResource {
    /**
     * Get an existing ActionsOrganizationPermissions resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ActionsOrganizationPermissionsState, opts?: pulumi.CustomResourceOptions): ActionsOrganizationPermissions {
        return new ActionsOrganizationPermissions(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:index/actionsOrganizationPermissions:ActionsOrganizationPermissions';

    /**
     * Returns true if the given object is an instance of ActionsOrganizationPermissions.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ActionsOrganizationPermissions {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ActionsOrganizationPermissions.__pulumiType;
    }

    /**
     * The permissions policy that controls the actions that are allowed to run. Can be one of: `all`, `localOnly`, or `selected`.
     */
    public readonly allowedActions!: pulumi.Output<string | undefined>;
    /**
     * Sets the actions that are allowed in an organization. Only available when `allowedActions` = `selected`. See Allowed Actions Config below for details.
     */
    public readonly allowedActionsConfig!: pulumi.Output<outputs.ActionsOrganizationPermissionsAllowedActionsConfig | undefined>;
    /**
     * The policy that controls the repositories in the organization that are allowed to run GitHub Actions. Can be one of: `all`, `none`, or `selected`.
     */
    public readonly enabledRepositories!: pulumi.Output<string>;
    /**
     * Sets the list of selected repositories that are enabled for GitHub Actions in an organization. Only available when `enabledRepositories` = `selected`. See Enabled Repositories Config below for details.
     */
    public readonly enabledRepositoriesConfig!: pulumi.Output<outputs.ActionsOrganizationPermissionsEnabledRepositoriesConfig | undefined>;

    /**
     * Create a ActionsOrganizationPermissions resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ActionsOrganizationPermissionsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ActionsOrganizationPermissionsArgs | ActionsOrganizationPermissionsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ActionsOrganizationPermissionsState | undefined;
            resourceInputs["allowedActions"] = state ? state.allowedActions : undefined;
            resourceInputs["allowedActionsConfig"] = state ? state.allowedActionsConfig : undefined;
            resourceInputs["enabledRepositories"] = state ? state.enabledRepositories : undefined;
            resourceInputs["enabledRepositoriesConfig"] = state ? state.enabledRepositoriesConfig : undefined;
        } else {
            const args = argsOrState as ActionsOrganizationPermissionsArgs | undefined;
            if ((!args || args.enabledRepositories === undefined) && !opts.urn) {
                throw new Error("Missing required property 'enabledRepositories'");
            }
            resourceInputs["allowedActions"] = args ? args.allowedActions : undefined;
            resourceInputs["allowedActionsConfig"] = args ? args.allowedActionsConfig : undefined;
            resourceInputs["enabledRepositories"] = args ? args.enabledRepositories : undefined;
            resourceInputs["enabledRepositoriesConfig"] = args ? args.enabledRepositoriesConfig : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ActionsOrganizationPermissions.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ActionsOrganizationPermissions resources.
 */
export interface ActionsOrganizationPermissionsState {
    /**
     * The permissions policy that controls the actions that are allowed to run. Can be one of: `all`, `localOnly`, or `selected`.
     */
    allowedActions?: pulumi.Input<string>;
    /**
     * Sets the actions that are allowed in an organization. Only available when `allowedActions` = `selected`. See Allowed Actions Config below for details.
     */
    allowedActionsConfig?: pulumi.Input<inputs.ActionsOrganizationPermissionsAllowedActionsConfig>;
    /**
     * The policy that controls the repositories in the organization that are allowed to run GitHub Actions. Can be one of: `all`, `none`, or `selected`.
     */
    enabledRepositories?: pulumi.Input<string>;
    /**
     * Sets the list of selected repositories that are enabled for GitHub Actions in an organization. Only available when `enabledRepositories` = `selected`. See Enabled Repositories Config below for details.
     */
    enabledRepositoriesConfig?: pulumi.Input<inputs.ActionsOrganizationPermissionsEnabledRepositoriesConfig>;
}

/**
 * The set of arguments for constructing a ActionsOrganizationPermissions resource.
 */
export interface ActionsOrganizationPermissionsArgs {
    /**
     * The permissions policy that controls the actions that are allowed to run. Can be one of: `all`, `localOnly`, or `selected`.
     */
    allowedActions?: pulumi.Input<string>;
    /**
     * Sets the actions that are allowed in an organization. Only available when `allowedActions` = `selected`. See Allowed Actions Config below for details.
     */
    allowedActionsConfig?: pulumi.Input<inputs.ActionsOrganizationPermissionsAllowedActionsConfig>;
    /**
     * The policy that controls the repositories in the organization that are allowed to run GitHub Actions. Can be one of: `all`, `none`, or `selected`.
     */
    enabledRepositories: pulumi.Input<string>;
    /**
     * Sets the list of selected repositories that are enabled for GitHub Actions in an organization. Only available when `enabledRepositories` = `selected`. See Enabled Repositories Config below for details.
     */
    enabledRepositoriesConfig?: pulumi.Input<inputs.ActionsOrganizationPermissionsEnabledRepositoriesConfig>;
}

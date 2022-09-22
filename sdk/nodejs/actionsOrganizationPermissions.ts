// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

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

    public readonly allowedActions!: pulumi.Output<string | undefined>;
    public readonly allowedActionsConfig!: pulumi.Output<outputs.ActionsOrganizationPermissionsAllowedActionsConfig | undefined>;
    public readonly enabledRepositories!: pulumi.Output<string>;
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
    allowedActions?: pulumi.Input<string>;
    allowedActionsConfig?: pulumi.Input<inputs.ActionsOrganizationPermissionsAllowedActionsConfig>;
    enabledRepositories?: pulumi.Input<string>;
    enabledRepositoriesConfig?: pulumi.Input<inputs.ActionsOrganizationPermissionsEnabledRepositoriesConfig>;
}

/**
 * The set of arguments for constructing a ActionsOrganizationPermissions resource.
 */
export interface ActionsOrganizationPermissionsArgs {
    allowedActions?: pulumi.Input<string>;
    allowedActionsConfig?: pulumi.Input<inputs.ActionsOrganizationPermissionsAllowedActionsConfig>;
    enabledRepositories: pulumi.Input<string>;
    enabledRepositoriesConfig?: pulumi.Input<inputs.ActionsOrganizationPermissionsEnabledRepositoriesConfig>;
}

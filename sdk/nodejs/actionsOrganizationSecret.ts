// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class ActionsOrganizationSecret extends pulumi.CustomResource {
    /**
     * Get an existing ActionsOrganizationSecret resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ActionsOrganizationSecretState, opts?: pulumi.CustomResourceOptions): ActionsOrganizationSecret {
        return new ActionsOrganizationSecret(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:index/actionsOrganizationSecret:ActionsOrganizationSecret';

    /**
     * Returns true if the given object is an instance of ActionsOrganizationSecret.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ActionsOrganizationSecret {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ActionsOrganizationSecret.__pulumiType;
    }

    /**
     * Date of actionsSecret creation.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Plaintext value of the secret to be encrypted
     */
    public readonly plaintextValue!: pulumi.Output<string>;
    /**
     * Name of the secret
     */
    public readonly secretName!: pulumi.Output<string>;
    /**
     * An array of repository ids that can access the organization secret.
     */
    public readonly selectedRepositoryIds!: pulumi.Output<number[] | undefined>;
    /**
     * Date of actionsSecret update.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * Configures the access that repositories have to the organization secret.
     * Must be one of `all`, `private`, `selected`. `selectedRepositoryIds` is required if set to `selected`.
     */
    public readonly visibility!: pulumi.Output<string>;

    /**
     * Create a ActionsOrganizationSecret resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ActionsOrganizationSecretArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ActionsOrganizationSecretArgs | ActionsOrganizationSecretState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ActionsOrganizationSecretState | undefined;
            inputs["createdAt"] = state ? state.createdAt : undefined;
            inputs["plaintextValue"] = state ? state.plaintextValue : undefined;
            inputs["secretName"] = state ? state.secretName : undefined;
            inputs["selectedRepositoryIds"] = state ? state.selectedRepositoryIds : undefined;
            inputs["updatedAt"] = state ? state.updatedAt : undefined;
            inputs["visibility"] = state ? state.visibility : undefined;
        } else {
            const args = argsOrState as ActionsOrganizationSecretArgs | undefined;
            if ((!args || args.plaintextValue === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'plaintextValue'");
            }
            if ((!args || args.secretName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'secretName'");
            }
            if ((!args || args.visibility === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'visibility'");
            }
            inputs["plaintextValue"] = args ? args.plaintextValue : undefined;
            inputs["secretName"] = args ? args.secretName : undefined;
            inputs["selectedRepositoryIds"] = args ? args.selectedRepositoryIds : undefined;
            inputs["visibility"] = args ? args.visibility : undefined;
            inputs["createdAt"] = undefined /*out*/;
            inputs["updatedAt"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ActionsOrganizationSecret.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ActionsOrganizationSecret resources.
 */
export interface ActionsOrganizationSecretState {
    /**
     * Date of actionsSecret creation.
     */
    readonly createdAt?: pulumi.Input<string>;
    /**
     * Plaintext value of the secret to be encrypted
     */
    readonly plaintextValue?: pulumi.Input<string>;
    /**
     * Name of the secret
     */
    readonly secretName?: pulumi.Input<string>;
    /**
     * An array of repository ids that can access the organization secret.
     */
    readonly selectedRepositoryIds?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * Date of actionsSecret update.
     */
    readonly updatedAt?: pulumi.Input<string>;
    /**
     * Configures the access that repositories have to the organization secret.
     * Must be one of `all`, `private`, `selected`. `selectedRepositoryIds` is required if set to `selected`.
     */
    readonly visibility?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ActionsOrganizationSecret resource.
 */
export interface ActionsOrganizationSecretArgs {
    /**
     * Plaintext value of the secret to be encrypted
     */
    readonly plaintextValue: pulumi.Input<string>;
    /**
     * Name of the secret
     */
    readonly secretName: pulumi.Input<string>;
    /**
     * An array of repository ids that can access the organization secret.
     */
    readonly selectedRepositoryIds?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * Configures the access that repositories have to the organization secret.
     * Must be one of `all`, `private`, `selected`. `selectedRepositoryIds` is required if set to `selected`.
     */
    readonly visibility: pulumi.Input<string>;
}
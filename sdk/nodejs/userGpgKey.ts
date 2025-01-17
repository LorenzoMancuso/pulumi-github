// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class UserGpgKey extends pulumi.CustomResource {
    /**
     * Get an existing UserGpgKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserGpgKeyState, opts?: pulumi.CustomResourceOptions): UserGpgKey {
        return new UserGpgKey(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:index/userGpgKey:UserGpgKey';

    /**
     * Returns true if the given object is an instance of UserGpgKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserGpgKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserGpgKey.__pulumiType;
    }

    public readonly armoredPublicKey!: pulumi.Output<string>;
    public /*out*/ readonly etag!: pulumi.Output<string>;
    public /*out*/ readonly keyId!: pulumi.Output<string>;

    /**
     * Create a UserGpgKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserGpgKeyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserGpgKeyArgs | UserGpgKeyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserGpgKeyState | undefined;
            resourceInputs["armoredPublicKey"] = state ? state.armoredPublicKey : undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["keyId"] = state ? state.keyId : undefined;
        } else {
            const args = argsOrState as UserGpgKeyArgs | undefined;
            if ((!args || args.armoredPublicKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'armoredPublicKey'");
            }
            resourceInputs["armoredPublicKey"] = args ? args.armoredPublicKey : undefined;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["keyId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(UserGpgKey.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserGpgKey resources.
 */
export interface UserGpgKeyState {
    armoredPublicKey?: pulumi.Input<string>;
    etag?: pulumi.Input<string>;
    keyId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UserGpgKey resource.
 */
export interface UserGpgKeyArgs {
    armoredPublicKey: pulumi.Input<string>;
}

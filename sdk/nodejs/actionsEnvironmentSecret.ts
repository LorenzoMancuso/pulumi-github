// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Import
 *
 * This resource can be imported using an ID made up of the secret name
 *
 * ```sh
 *  $ pulumi import github:index/actionsEnvironmentSecret:ActionsEnvironmentSecret test_secret test_secret_name
 * ```
 *
 *  NOTEthe implementation is limited in that it won't fetch the value of the `plaintext_value` field when importing. You may need to ignore changes for the `plaintext_value` as a workaround.
 */
export class ActionsEnvironmentSecret extends pulumi.CustomResource {
    /**
     * Get an existing ActionsEnvironmentSecret resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ActionsEnvironmentSecretState, opts?: pulumi.CustomResourceOptions): ActionsEnvironmentSecret {
        return new ActionsEnvironmentSecret(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:index/actionsEnvironmentSecret:ActionsEnvironmentSecret';

    /**
     * Returns true if the given object is an instance of ActionsEnvironmentSecret.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ActionsEnvironmentSecret {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ActionsEnvironmentSecret.__pulumiType;
    }

    /**
     * Date of actionsEnvironmentSecret creation.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Encrypted value of the secret using the Github public key in Base64 format.
     */
    public readonly encryptedValue!: pulumi.Output<string | undefined>;
    /**
     * Name of the environment.
     */
    public readonly environment!: pulumi.Output<string>;
    /**
     * Plaintext value of the secret to be encrypted.
     */
    public readonly plaintextValue!: pulumi.Output<string | undefined>;
    /**
     * Name of the repository.
     */
    public readonly repository!: pulumi.Output<string>;
    /**
     * Name of the secret.
     */
    public readonly secretName!: pulumi.Output<string>;
    /**
     * Date of actionsEnvironmentSecret update.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a ActionsEnvironmentSecret resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ActionsEnvironmentSecretArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ActionsEnvironmentSecretArgs | ActionsEnvironmentSecretState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ActionsEnvironmentSecretState | undefined;
            inputs["createdAt"] = state ? state.createdAt : undefined;
            inputs["encryptedValue"] = state ? state.encryptedValue : undefined;
            inputs["environment"] = state ? state.environment : undefined;
            inputs["plaintextValue"] = state ? state.plaintextValue : undefined;
            inputs["repository"] = state ? state.repository : undefined;
            inputs["secretName"] = state ? state.secretName : undefined;
            inputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as ActionsEnvironmentSecretArgs | undefined;
            if ((!args || args.environment === undefined) && !opts.urn) {
                throw new Error("Missing required property 'environment'");
            }
            if ((!args || args.repository === undefined) && !opts.urn) {
                throw new Error("Missing required property 'repository'");
            }
            if ((!args || args.secretName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'secretName'");
            }
            inputs["encryptedValue"] = args ? args.encryptedValue : undefined;
            inputs["environment"] = args ? args.environment : undefined;
            inputs["plaintextValue"] = args ? args.plaintextValue : undefined;
            inputs["repository"] = args ? args.repository : undefined;
            inputs["secretName"] = args ? args.secretName : undefined;
            inputs["createdAt"] = undefined /*out*/;
            inputs["updatedAt"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ActionsEnvironmentSecret.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ActionsEnvironmentSecret resources.
 */
export interface ActionsEnvironmentSecretState {
    /**
     * Date of actionsEnvironmentSecret creation.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Encrypted value of the secret using the Github public key in Base64 format.
     */
    encryptedValue?: pulumi.Input<string>;
    /**
     * Name of the environment.
     */
    environment?: pulumi.Input<string>;
    /**
     * Plaintext value of the secret to be encrypted.
     */
    plaintextValue?: pulumi.Input<string>;
    /**
     * Name of the repository.
     */
    repository?: pulumi.Input<string>;
    /**
     * Name of the secret.
     */
    secretName?: pulumi.Input<string>;
    /**
     * Date of actionsEnvironmentSecret update.
     */
    updatedAt?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ActionsEnvironmentSecret resource.
 */
export interface ActionsEnvironmentSecretArgs {
    /**
     * Encrypted value of the secret using the Github public key in Base64 format.
     */
    encryptedValue?: pulumi.Input<string>;
    /**
     * Name of the environment.
     */
    environment: pulumi.Input<string>;
    /**
     * Plaintext value of the secret to be encrypted.
     */
    plaintextValue?: pulumi.Input<string>;
    /**
     * Name of the repository.
     */
    repository: pulumi.Input<string>;
    /**
     * Name of the secret.
     */
    secretName: pulumi.Input<string>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use this data source to retrieve information about a GitHub Actions public key. This data source is required to be used with other GitHub secrets interactions.
 * Note that the provider `token` must have admin rights to a repository to retrieve it's action public key.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 * 
 * const example = pulumi.output(github.SecretsPublicKey({
 *     repository: "exampleRepo",
 * }, { async: true }));
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-github/blob/master/website/docs/d/actions_public_key.html.markdown.
 */
export function getActionsPublicKey(args: GetActionsPublicKeyArgs, opts?: pulumi.InvokeOptions): Promise<GetActionsPublicKeyResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("github:index/getActionsPublicKey:getActionsPublicKey", {
        "repository": args.repository,
    }, opts);
}

/**
 * A collection of arguments for invoking getActionsPublicKey.
 */
export interface GetActionsPublicKeyArgs {
    /**
     * Name of the repository to get public key from.
     */
    readonly repository: string;
}

/**
 * A collection of values returned by getActionsPublicKey.
 */
export interface GetActionsPublicKeyResult {
    /**
     * Actual key retrieved.
     */
    readonly key: string;
    /**
     * ID of the key that has been retrieved.
     */
    readonly keyId: string;
    readonly repository: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The provider type for the github package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 */
export class Provider extends pulumi.ProviderResource {
    /** @internal */
    public static readonly __pulumiType = 'github';

    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Provider.__pulumiType;
    }


    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let inputs: pulumi.Inputs = {};
        inputs["baseUrl"] = (args ? args.baseUrl : undefined) || (utilities.getEnv("GITHUB_BASE_URL") || "https://api.github.com/");
        inputs["organization"] = (args ? args.organization : undefined) || utilities.getEnv("GITHUB_ORGANIZATION");
        inputs["owner"] = args ? args.owner : undefined;
        inputs["token"] = (args ? args.token : undefined) || utilities.getEnv("GITHUB_TOKEN");
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Provider.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    /**
     * The GitHub Base API URL
     */
    readonly baseUrl?: pulumi.Input<string>;
    /**
     * (Deprecated) The GitHub organization name to manage.
     *
     * @deprecated Use owner field (or GITHUB_OWNER ENV variable)
     */
    readonly organization?: pulumi.Input<string>;
    /**
     * The GitHub owner name to manage.
     */
    readonly owner?: pulumi.Input<string>;
    /**
     * The OAuth token used to connect to GitHub. If `anonymous` is false, `token` is required.
     */
    readonly token?: pulumi.Input<string>;
}

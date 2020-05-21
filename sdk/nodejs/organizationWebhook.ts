// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * This resource allows you to create and manage webhooks for GitHub organization.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const foo = new github.OrganizationWebhook("foo", {
 *     active: false,
 *     configuration: {
 *         contentType: "form",
 *         insecureSsl: false,
 *         url: "https://google.de/",
 *     },
 *     events: ["issues"],
 * });
 * ```
 */
export class OrganizationWebhook extends pulumi.CustomResource {
    /**
     * Get an existing OrganizationWebhook resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OrganizationWebhookState, opts?: pulumi.CustomResourceOptions): OrganizationWebhook {
        return new OrganizationWebhook(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:index/organizationWebhook:OrganizationWebhook';

    /**
     * Returns true if the given object is an instance of OrganizationWebhook.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OrganizationWebhook {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OrganizationWebhook.__pulumiType;
    }

    /**
     * Indicate of the webhook should receive events. Defaults to `true`.
     */
    public readonly active!: pulumi.Output<boolean | undefined>;
    /**
     * key/value pair of configuration for this webhook. Available keys are `url`, `contentType`, `secret` and `insecureSsl`.
     */
    public readonly configuration!: pulumi.Output<outputs.OrganizationWebhookConfiguration | undefined>;
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * A list of events which should trigger the webhook. See a list of [available events](https://developer.github.com/v3/activity/events/types/)
     */
    public readonly events!: pulumi.Output<string[]>;
    /**
     * URL of the webhook
     */
    public /*out*/ readonly url!: pulumi.Output<string>;

    /**
     * Create a OrganizationWebhook resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OrganizationWebhookArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OrganizationWebhookArgs | OrganizationWebhookState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as OrganizationWebhookState | undefined;
            inputs["active"] = state ? state.active : undefined;
            inputs["configuration"] = state ? state.configuration : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["events"] = state ? state.events : undefined;
            inputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as OrganizationWebhookArgs | undefined;
            if (!args || args.events === undefined) {
                throw new Error("Missing required property 'events'");
            }
            inputs["active"] = args ? args.active : undefined;
            inputs["configuration"] = args ? args.configuration : undefined;
            inputs["events"] = args ? args.events : undefined;
            inputs["etag"] = undefined /*out*/;
            inputs["url"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(OrganizationWebhook.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OrganizationWebhook resources.
 */
export interface OrganizationWebhookState {
    /**
     * Indicate of the webhook should receive events. Defaults to `true`.
     */
    readonly active?: pulumi.Input<boolean>;
    /**
     * key/value pair of configuration for this webhook. Available keys are `url`, `contentType`, `secret` and `insecureSsl`.
     */
    readonly configuration?: pulumi.Input<inputs.OrganizationWebhookConfiguration>;
    readonly etag?: pulumi.Input<string>;
    /**
     * A list of events which should trigger the webhook. See a list of [available events](https://developer.github.com/v3/activity/events/types/)
     */
    readonly events?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * URL of the webhook
     */
    readonly url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a OrganizationWebhook resource.
 */
export interface OrganizationWebhookArgs {
    /**
     * Indicate of the webhook should receive events. Defaults to `true`.
     */
    readonly active?: pulumi.Input<boolean>;
    /**
     * key/value pair of configuration for this webhook. Available keys are `url`, `contentType`, `secret` and `insecureSsl`.
     */
    readonly configuration?: pulumi.Input<inputs.OrganizationWebhookConfiguration>;
    /**
     * A list of events which should trigger the webhook. See a list of [available events](https://developer.github.com/v3/activity/events/types/)
     */
    readonly events: pulumi.Input<pulumi.Input<string>[]>;
}

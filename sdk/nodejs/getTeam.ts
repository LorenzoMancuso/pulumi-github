// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to retrieve information about a GitHub team.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const example = pulumi.output(github.getTeam({
 *     slug: "example",
 * }));
 * ```
 */
export function getTeam(args: GetTeamArgs, opts?: pulumi.InvokeOptions): Promise<GetTeamResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("github:index/getTeam:getTeam", {
        "slug": args.slug,
    }, opts);
}

/**
 * A collection of arguments for invoking getTeam.
 */
export interface GetTeamArgs {
    /**
     * The team slug.
     */
    slug: string;
}

/**
 * A collection of values returned by getTeam.
 */
export interface GetTeamResult {
    /**
     * the team's description.
     */
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * List of team members
     */
    readonly members: string[];
    /**
     * the team's full name.
     */
    readonly name: string;
    /**
     * the Node ID of the team.
     */
    readonly nodeId: string;
    /**
     * the team's permission level.
     */
    readonly permission: string;
    /**
     * the team's privacy type.
     */
    readonly privacy: string;
    /**
     * List of team repositories
     */
    readonly repositories: string[];
    readonly slug: string;
}

export function getTeamOutput(args: GetTeamOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTeamResult> {
    return pulumi.output(args).apply(a => getTeam(a, opts))
}

/**
 * A collection of arguments for invoking getTeam.
 */
export interface GetTeamOutputArgs {
    /**
     * The team slug.
     */
    slug: pulumi.Input<string>;
}

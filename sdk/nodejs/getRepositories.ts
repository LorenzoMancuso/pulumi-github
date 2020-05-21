// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * > **Note:** The data source will return a maximum of `1000` repositories
 * 	[as documented in official API docs](https://developer.github.com/v3/search/#about-the-search-api).
 *
 * Use this data source to retrieve a list of GitHub repositories using a search query.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const example = pulumi.output(github.getRepositories({
 *     query: "org:hashicorp language:Go",
 * }, { async: true }));
 * ```
 */
export function getRepositories(args: GetRepositoriesArgs, opts?: pulumi.InvokeOptions): Promise<GetRepositoriesResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("github:index/getRepositories:getRepositories", {
        "query": args.query,
        "sort": args.sort,
    }, opts);
}

/**
 * A collection of arguments for invoking getRepositories.
 */
export interface GetRepositoriesArgs {
    /**
     * Search query. See [documentation for the search syntax](https://help.github.com/articles/understanding-the-search-syntax/).
     */
    readonly query: string;
    /**
     * Sorts the repositories returned by the specified attribute. Valid values include `stars`, `fork`, and `updated`. Defaults to `updated`.
     */
    readonly sort?: string;
}

/**
 * A collection of values returned by getRepositories.
 */
export interface GetRepositoriesResult {
    readonly fullNames: string[];
    readonly names: string[];
    readonly query: string;
    readonly sort?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}

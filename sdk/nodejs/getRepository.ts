// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getRepository(args?: GetRepositoryArgs, opts?: pulumi.InvokeOptions): Promise<GetRepositoryResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("github:index/getRepository:getRepository", {
        "description": args.description,
        "fullName": args.fullName,
        "homepageUrl": args.homepageUrl,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getRepository.
 */
export interface GetRepositoryArgs {
    description?: string;
    fullName?: string;
    homepageUrl?: string;
    name?: string;
}

/**
 * A collection of values returned by getRepository.
 */
export interface GetRepositoryResult {
    readonly allowAutoMerge: boolean;
    readonly allowMergeCommit: boolean;
    readonly allowRebaseMerge: boolean;
    readonly allowSquashMerge: boolean;
    readonly archived: boolean;
    readonly defaultBranch: string;
    readonly description?: string;
    readonly fullName: string;
    readonly gitCloneUrl: string;
    readonly hasDownloads: boolean;
    readonly hasIssues: boolean;
    readonly hasProjects: boolean;
    readonly hasWiki: boolean;
    readonly homepageUrl?: string;
    readonly htmlUrl: string;
    readonly httpCloneUrl: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly isTemplate: boolean;
    readonly mergeCommitMessage: string;
    readonly mergeCommitTitle: string;
    readonly name: string;
    readonly nodeId: string;
    readonly pages: outputs.GetRepositoryPage[];
    readonly private: boolean;
    readonly repoId: number;
    readonly squashMergeCommitMessage: string;
    readonly squashMergeCommitTitle: string;
    readonly sshCloneUrl: string;
    readonly svnUrl: string;
    readonly template: outputs.GetRepositoryTemplate;
    readonly topics: string[];
    readonly visibility: string;
}

export function getRepositoryOutput(args?: GetRepositoryOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRepositoryResult> {
    return pulumi.output(args).apply(a => getRepository(a, opts))
}

/**
 * A collection of arguments for invoking getRepository.
 */
export interface GetRepositoryOutputArgs {
    description?: pulumi.Input<string>;
    fullName?: pulumi.Input<string>;
    homepageUrl?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
}

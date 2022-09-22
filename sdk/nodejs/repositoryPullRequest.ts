// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class RepositoryPullRequest extends pulumi.CustomResource {
    /**
     * Get an existing RepositoryPullRequest resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RepositoryPullRequestState, opts?: pulumi.CustomResourceOptions): RepositoryPullRequest {
        return new RepositoryPullRequest(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:index/repositoryPullRequest:RepositoryPullRequest';

    /**
     * Returns true if the given object is an instance of RepositoryPullRequest.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RepositoryPullRequest {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RepositoryPullRequest.__pulumiType;
    }

    public readonly baseRef!: pulumi.Output<string>;
    public readonly baseRepository!: pulumi.Output<string>;
    public /*out*/ readonly baseSha!: pulumi.Output<string>;
    public readonly body!: pulumi.Output<string | undefined>;
    public /*out*/ readonly draft!: pulumi.Output<boolean>;
    public readonly headRef!: pulumi.Output<string>;
    public /*out*/ readonly headSha!: pulumi.Output<string>;
    /**
     * List of names of labels on the PR
     */
    public /*out*/ readonly labels!: pulumi.Output<string[]>;
    public readonly maintainerCanModify!: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly number!: pulumi.Output<number>;
    public /*out*/ readonly openedAt!: pulumi.Output<number>;
    /**
     * Username of the PR creator
     */
    public /*out*/ readonly openedBy!: pulumi.Output<string>;
    public readonly owner!: pulumi.Output<string | undefined>;
    public /*out*/ readonly state!: pulumi.Output<string>;
    public readonly title!: pulumi.Output<string>;
    public /*out*/ readonly updatedAt!: pulumi.Output<number>;

    /**
     * Create a RepositoryPullRequest resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RepositoryPullRequestArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RepositoryPullRequestArgs | RepositoryPullRequestState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RepositoryPullRequestState | undefined;
            resourceInputs["baseRef"] = state ? state.baseRef : undefined;
            resourceInputs["baseRepository"] = state ? state.baseRepository : undefined;
            resourceInputs["baseSha"] = state ? state.baseSha : undefined;
            resourceInputs["body"] = state ? state.body : undefined;
            resourceInputs["draft"] = state ? state.draft : undefined;
            resourceInputs["headRef"] = state ? state.headRef : undefined;
            resourceInputs["headSha"] = state ? state.headSha : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["maintainerCanModify"] = state ? state.maintainerCanModify : undefined;
            resourceInputs["number"] = state ? state.number : undefined;
            resourceInputs["openedAt"] = state ? state.openedAt : undefined;
            resourceInputs["openedBy"] = state ? state.openedBy : undefined;
            resourceInputs["owner"] = state ? state.owner : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["title"] = state ? state.title : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as RepositoryPullRequestArgs | undefined;
            if ((!args || args.baseRef === undefined) && !opts.urn) {
                throw new Error("Missing required property 'baseRef'");
            }
            if ((!args || args.baseRepository === undefined) && !opts.urn) {
                throw new Error("Missing required property 'baseRepository'");
            }
            if ((!args || args.headRef === undefined) && !opts.urn) {
                throw new Error("Missing required property 'headRef'");
            }
            if ((!args || args.title === undefined) && !opts.urn) {
                throw new Error("Missing required property 'title'");
            }
            resourceInputs["baseRef"] = args ? args.baseRef : undefined;
            resourceInputs["baseRepository"] = args ? args.baseRepository : undefined;
            resourceInputs["body"] = args ? args.body : undefined;
            resourceInputs["headRef"] = args ? args.headRef : undefined;
            resourceInputs["maintainerCanModify"] = args ? args.maintainerCanModify : undefined;
            resourceInputs["owner"] = args ? args.owner : undefined;
            resourceInputs["title"] = args ? args.title : undefined;
            resourceInputs["baseSha"] = undefined /*out*/;
            resourceInputs["draft"] = undefined /*out*/;
            resourceInputs["headSha"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["number"] = undefined /*out*/;
            resourceInputs["openedAt"] = undefined /*out*/;
            resourceInputs["openedBy"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RepositoryPullRequest.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RepositoryPullRequest resources.
 */
export interface RepositoryPullRequestState {
    baseRef?: pulumi.Input<string>;
    baseRepository?: pulumi.Input<string>;
    baseSha?: pulumi.Input<string>;
    body?: pulumi.Input<string>;
    draft?: pulumi.Input<boolean>;
    headRef?: pulumi.Input<string>;
    headSha?: pulumi.Input<string>;
    /**
     * List of names of labels on the PR
     */
    labels?: pulumi.Input<pulumi.Input<string>[]>;
    maintainerCanModify?: pulumi.Input<boolean>;
    number?: pulumi.Input<number>;
    openedAt?: pulumi.Input<number>;
    /**
     * Username of the PR creator
     */
    openedBy?: pulumi.Input<string>;
    owner?: pulumi.Input<string>;
    state?: pulumi.Input<string>;
    title?: pulumi.Input<string>;
    updatedAt?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a RepositoryPullRequest resource.
 */
export interface RepositoryPullRequestArgs {
    baseRef: pulumi.Input<string>;
    baseRepository: pulumi.Input<string>;
    body?: pulumi.Input<string>;
    headRef: pulumi.Input<string>;
    maintainerCanModify?: pulumi.Input<boolean>;
    owner?: pulumi.Input<string>;
    title: pulumi.Input<string>;
}

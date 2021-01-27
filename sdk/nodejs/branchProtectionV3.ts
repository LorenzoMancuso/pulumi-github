// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Protects a GitHub branch.
 *
 * The `github.BranchProtection` resource has moved to the GraphQL API, while this resource will continue to leverage the REST API.
 *
 * This resource allows you to configure branch protection for repositories in your organization. When applied, the branch will be protected from forced pushes and deletion. Additional constraints, such as required status checks or restrictions on users, teams, and apps, can also be configured.
 *
 * ## Import
 *
 * GitHub Branch Protection can be imported using an ID made up of `repository:branch`, e.g.
 *
 * ```sh
 *  $ pulumi import github:index/branchProtectionV3:BranchProtectionV3 github_branch_protection_v3.terraform terraform:main
 * ```
 */
export class BranchProtectionV3 extends pulumi.CustomResource {
    /**
     * Get an existing BranchProtectionV3 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BranchProtectionV3State, opts?: pulumi.CustomResourceOptions): BranchProtectionV3 {
        return new BranchProtectionV3(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:index/branchProtectionV3:BranchProtectionV3';

    /**
     * Returns true if the given object is an instance of BranchProtectionV3.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BranchProtectionV3 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BranchProtectionV3.__pulumiType;
    }

    /**
     * The Git branch to protect.
     */
    public readonly branch!: pulumi.Output<string>;
    /**
     * Boolean, setting this to `true` enforces status checks for repository administrators.
     */
    public readonly enforceAdmins!: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The GitHub repository name.
     */
    public readonly repository!: pulumi.Output<string>;
    /**
     * Boolean, setting this to `true` requires all commits to be signed with GPG.
     */
    public readonly requireSignedCommits!: pulumi.Output<boolean | undefined>;
    /**
     * Enforce restrictions for pull request reviews. See Required Pull Request Reviews below for details.
     */
    public readonly requiredPullRequestReviews!: pulumi.Output<outputs.BranchProtectionV3RequiredPullRequestReviews | undefined>;
    /**
     * Enforce restrictions for required status checks. See Required Status Checks below for details.
     */
    public readonly requiredStatusChecks!: pulumi.Output<outputs.BranchProtectionV3RequiredStatusChecks | undefined>;
    /**
     * Enforce restrictions for the users and teams that may push to the branch. See Restrictions below for details.
     */
    public readonly restrictions!: pulumi.Output<outputs.BranchProtectionV3Restrictions | undefined>;

    /**
     * Create a BranchProtectionV3 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BranchProtectionV3Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BranchProtectionV3Args | BranchProtectionV3State, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as BranchProtectionV3State | undefined;
            inputs["branch"] = state ? state.branch : undefined;
            inputs["enforceAdmins"] = state ? state.enforceAdmins : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["repository"] = state ? state.repository : undefined;
            inputs["requireSignedCommits"] = state ? state.requireSignedCommits : undefined;
            inputs["requiredPullRequestReviews"] = state ? state.requiredPullRequestReviews : undefined;
            inputs["requiredStatusChecks"] = state ? state.requiredStatusChecks : undefined;
            inputs["restrictions"] = state ? state.restrictions : undefined;
        } else {
            const args = argsOrState as BranchProtectionV3Args | undefined;
            if ((!args || args.branch === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'branch'");
            }
            if ((!args || args.repository === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'repository'");
            }
            inputs["branch"] = args ? args.branch : undefined;
            inputs["enforceAdmins"] = args ? args.enforceAdmins : undefined;
            inputs["repository"] = args ? args.repository : undefined;
            inputs["requireSignedCommits"] = args ? args.requireSignedCommits : undefined;
            inputs["requiredPullRequestReviews"] = args ? args.requiredPullRequestReviews : undefined;
            inputs["requiredStatusChecks"] = args ? args.requiredStatusChecks : undefined;
            inputs["restrictions"] = args ? args.restrictions : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(BranchProtectionV3.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BranchProtectionV3 resources.
 */
export interface BranchProtectionV3State {
    /**
     * The Git branch to protect.
     */
    readonly branch?: pulumi.Input<string>;
    /**
     * Boolean, setting this to `true` enforces status checks for repository administrators.
     */
    readonly enforceAdmins?: pulumi.Input<boolean>;
    readonly etag?: pulumi.Input<string>;
    /**
     * The GitHub repository name.
     */
    readonly repository?: pulumi.Input<string>;
    /**
     * Boolean, setting this to `true` requires all commits to be signed with GPG.
     */
    readonly requireSignedCommits?: pulumi.Input<boolean>;
    /**
     * Enforce restrictions for pull request reviews. See Required Pull Request Reviews below for details.
     */
    readonly requiredPullRequestReviews?: pulumi.Input<inputs.BranchProtectionV3RequiredPullRequestReviews>;
    /**
     * Enforce restrictions for required status checks. See Required Status Checks below for details.
     */
    readonly requiredStatusChecks?: pulumi.Input<inputs.BranchProtectionV3RequiredStatusChecks>;
    /**
     * Enforce restrictions for the users and teams that may push to the branch. See Restrictions below for details.
     */
    readonly restrictions?: pulumi.Input<inputs.BranchProtectionV3Restrictions>;
}

/**
 * The set of arguments for constructing a BranchProtectionV3 resource.
 */
export interface BranchProtectionV3Args {
    /**
     * The Git branch to protect.
     */
    readonly branch: pulumi.Input<string>;
    /**
     * Boolean, setting this to `true` enforces status checks for repository administrators.
     */
    readonly enforceAdmins?: pulumi.Input<boolean>;
    /**
     * The GitHub repository name.
     */
    readonly repository: pulumi.Input<string>;
    /**
     * Boolean, setting this to `true` requires all commits to be signed with GPG.
     */
    readonly requireSignedCommits?: pulumi.Input<boolean>;
    /**
     * Enforce restrictions for pull request reviews. See Required Pull Request Reviews below for details.
     */
    readonly requiredPullRequestReviews?: pulumi.Input<inputs.BranchProtectionV3RequiredPullRequestReviews>;
    /**
     * Enforce restrictions for required status checks. See Required Status Checks below for details.
     */
    readonly requiredStatusChecks?: pulumi.Input<inputs.BranchProtectionV3RequiredStatusChecks>;
    /**
     * Enforce restrictions for the users and teams that may push to the branch. See Restrictions below for details.
     */
    readonly restrictions?: pulumi.Input<inputs.BranchProtectionV3Restrictions>;
}
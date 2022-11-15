// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface ActionsOrganizationPermissionsAllowedActionsConfig {
    githubOwnedAllowed: pulumi.Input<boolean>;
    patternsAlloweds?: pulumi.Input<pulumi.Input<string>[]>;
    verifiedAllowed?: pulumi.Input<boolean>;
}

export interface ActionsOrganizationPermissionsEnabledRepositoriesConfig {
    repositoryIds: pulumi.Input<pulumi.Input<number>[]>;
}

export interface BranchProtectionRequiredPullRequestReview {
    dismissStaleReviews?: pulumi.Input<boolean>;
    dismissalRestrictions?: pulumi.Input<pulumi.Input<string>[]>;
    pullRequestBypassers?: pulumi.Input<pulumi.Input<string>[]>;
    requireCodeOwnerReviews?: pulumi.Input<boolean>;
    requiredApprovingReviewCount?: pulumi.Input<number>;
    restrictDismissals?: pulumi.Input<boolean>;
}

export interface BranchProtectionRequiredStatusCheck {
    contexts?: pulumi.Input<pulumi.Input<string>[]>;
    strict?: pulumi.Input<boolean>;
}

export interface BranchProtectionV3RequiredPullRequestReviews {
    dismissStaleReviews?: pulumi.Input<boolean>;
    dismissalTeams?: pulumi.Input<pulumi.Input<string>[]>;
    dismissalUsers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * @deprecated Use enforce_admins instead
     */
    includeAdmins?: pulumi.Input<boolean>;
    requireCodeOwnerReviews?: pulumi.Input<boolean>;
    requiredApprovingReviewCount?: pulumi.Input<number>;
}

export interface BranchProtectionV3RequiredStatusChecks {
    contexts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * @deprecated Use enforce_admins instead
     */
    includeAdmins?: pulumi.Input<boolean>;
    strict?: pulumi.Input<boolean>;
}

export interface BranchProtectionV3Restrictions {
    apps?: pulumi.Input<pulumi.Input<string>[]>;
    teams?: pulumi.Input<pulumi.Input<string>[]>;
    users?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface OrganizationWebhookConfiguration {
    contentType?: pulumi.Input<string>;
    insecureSsl?: pulumi.Input<boolean>;
    secret?: pulumi.Input<string>;
    url: pulumi.Input<string>;
}

export interface ProviderAppAuth {
    id: pulumi.Input<string>;
    installationId: pulumi.Input<string>;
    pemFile: pulumi.Input<string>;
}

export interface RepositoryEnvironmentDeploymentBranchPolicy {
    customBranchPolicies: pulumi.Input<boolean>;
    protectedBranches: pulumi.Input<boolean>;
}

export interface RepositoryEnvironmentReviewer {
    teams?: pulumi.Input<pulumi.Input<number>[]>;
    users?: pulumi.Input<pulumi.Input<number>[]>;
}

export interface RepositoryPages {
    cname?: pulumi.Input<string>;
    custom404?: pulumi.Input<boolean>;
    htmlUrl?: pulumi.Input<string>;
    source: pulumi.Input<inputs.RepositoryPagesSource>;
    status?: pulumi.Input<string>;
    url?: pulumi.Input<string>;
}

export interface RepositoryPagesSource {
    branch: pulumi.Input<string>;
    path?: pulumi.Input<string>;
}

export interface RepositoryTemplate {
    includeAllBranches?: pulumi.Input<boolean>;
    owner: pulumi.Input<string>;
    repository: pulumi.Input<string>;
}

export interface RepositoryWebhookConfiguration {
    contentType?: pulumi.Input<string>;
    insecureSsl?: pulumi.Input<boolean>;
    secret?: pulumi.Input<string>;
    url: pulumi.Input<string>;
}

export interface TeamMembersMember {
    role?: pulumi.Input<string>;
    username: pulumi.Input<string>;
}

export interface TeamSyncGroupMappingGroup {
    groupDescription: pulumi.Input<string>;
    groupId: pulumi.Input<string>;
    groupName: pulumi.Input<string>;
}
export namespace config {
}

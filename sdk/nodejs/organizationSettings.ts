// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class OrganizationSettings extends pulumi.CustomResource {
    /**
     * Get an existing OrganizationSettings resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OrganizationSettingsState, opts?: pulumi.CustomResourceOptions): OrganizationSettings {
        return new OrganizationSettings(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:index/organizationSettings:OrganizationSettings';

    /**
     * Returns true if the given object is an instance of OrganizationSettings.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OrganizationSettings {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OrganizationSettings.__pulumiType;
    }

    public readonly advancedSecurityEnabledForNewRepositories!: pulumi.Output<boolean | undefined>;
    public readonly billingEmail!: pulumi.Output<string>;
    public readonly blog!: pulumi.Output<string | undefined>;
    public readonly company!: pulumi.Output<string | undefined>;
    public readonly defaultRepositoryPermission!: pulumi.Output<string | undefined>;
    public readonly dependabotAlertsEnabledForNewRepositories!: pulumi.Output<boolean | undefined>;
    public readonly dependabotSecurityUpdatesEnabledForNewRepositories!: pulumi.Output<boolean | undefined>;
    public readonly dependencyGraphEnabledForNewRepositories!: pulumi.Output<boolean | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly email!: pulumi.Output<string | undefined>;
    public readonly hasOrganizationProjects!: pulumi.Output<boolean | undefined>;
    public readonly hasRepositoryProjects!: pulumi.Output<boolean | undefined>;
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Setting to true allows organization members to create internal repositories. Only available to Enterprise Organizations.
     */
    public readonly membersCanCreateInternalRepositories!: pulumi.Output<boolean | undefined>;
    public readonly membersCanCreatePages!: pulumi.Output<boolean | undefined>;
    public readonly membersCanCreatePrivatePages!: pulumi.Output<boolean | undefined>;
    public readonly membersCanCreatePrivateRepositories!: pulumi.Output<boolean | undefined>;
    public readonly membersCanCreatePublicPages!: pulumi.Output<boolean | undefined>;
    public readonly membersCanCreatePublicRepositories!: pulumi.Output<boolean | undefined>;
    public readonly membersCanCreateRepositories!: pulumi.Output<boolean | undefined>;
    public readonly membersCanForkPrivateRepositories!: pulumi.Output<boolean | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly secretScanningEnabledForNewRepositories!: pulumi.Output<boolean | undefined>;
    public readonly secretScanningPushProtectionEnabledForNewRepositories!: pulumi.Output<boolean | undefined>;
    public readonly twitterUsername!: pulumi.Output<string | undefined>;
    public readonly webCommitSignoffRequired!: pulumi.Output<boolean | undefined>;

    /**
     * Create a OrganizationSettings resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OrganizationSettingsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OrganizationSettingsArgs | OrganizationSettingsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OrganizationSettingsState | undefined;
            resourceInputs["advancedSecurityEnabledForNewRepositories"] = state ? state.advancedSecurityEnabledForNewRepositories : undefined;
            resourceInputs["billingEmail"] = state ? state.billingEmail : undefined;
            resourceInputs["blog"] = state ? state.blog : undefined;
            resourceInputs["company"] = state ? state.company : undefined;
            resourceInputs["defaultRepositoryPermission"] = state ? state.defaultRepositoryPermission : undefined;
            resourceInputs["dependabotAlertsEnabledForNewRepositories"] = state ? state.dependabotAlertsEnabledForNewRepositories : undefined;
            resourceInputs["dependabotSecurityUpdatesEnabledForNewRepositories"] = state ? state.dependabotSecurityUpdatesEnabledForNewRepositories : undefined;
            resourceInputs["dependencyGraphEnabledForNewRepositories"] = state ? state.dependencyGraphEnabledForNewRepositories : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["email"] = state ? state.email : undefined;
            resourceInputs["hasOrganizationProjects"] = state ? state.hasOrganizationProjects : undefined;
            resourceInputs["hasRepositoryProjects"] = state ? state.hasRepositoryProjects : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["membersCanCreateInternalRepositories"] = state ? state.membersCanCreateInternalRepositories : undefined;
            resourceInputs["membersCanCreatePages"] = state ? state.membersCanCreatePages : undefined;
            resourceInputs["membersCanCreatePrivatePages"] = state ? state.membersCanCreatePrivatePages : undefined;
            resourceInputs["membersCanCreatePrivateRepositories"] = state ? state.membersCanCreatePrivateRepositories : undefined;
            resourceInputs["membersCanCreatePublicPages"] = state ? state.membersCanCreatePublicPages : undefined;
            resourceInputs["membersCanCreatePublicRepositories"] = state ? state.membersCanCreatePublicRepositories : undefined;
            resourceInputs["membersCanCreateRepositories"] = state ? state.membersCanCreateRepositories : undefined;
            resourceInputs["membersCanForkPrivateRepositories"] = state ? state.membersCanForkPrivateRepositories : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["secretScanningEnabledForNewRepositories"] = state ? state.secretScanningEnabledForNewRepositories : undefined;
            resourceInputs["secretScanningPushProtectionEnabledForNewRepositories"] = state ? state.secretScanningPushProtectionEnabledForNewRepositories : undefined;
            resourceInputs["twitterUsername"] = state ? state.twitterUsername : undefined;
            resourceInputs["webCommitSignoffRequired"] = state ? state.webCommitSignoffRequired : undefined;
        } else {
            const args = argsOrState as OrganizationSettingsArgs | undefined;
            if ((!args || args.billingEmail === undefined) && !opts.urn) {
                throw new Error("Missing required property 'billingEmail'");
            }
            resourceInputs["advancedSecurityEnabledForNewRepositories"] = args ? args.advancedSecurityEnabledForNewRepositories : undefined;
            resourceInputs["billingEmail"] = args ? args.billingEmail : undefined;
            resourceInputs["blog"] = args ? args.blog : undefined;
            resourceInputs["company"] = args ? args.company : undefined;
            resourceInputs["defaultRepositoryPermission"] = args ? args.defaultRepositoryPermission : undefined;
            resourceInputs["dependabotAlertsEnabledForNewRepositories"] = args ? args.dependabotAlertsEnabledForNewRepositories : undefined;
            resourceInputs["dependabotSecurityUpdatesEnabledForNewRepositories"] = args ? args.dependabotSecurityUpdatesEnabledForNewRepositories : undefined;
            resourceInputs["dependencyGraphEnabledForNewRepositories"] = args ? args.dependencyGraphEnabledForNewRepositories : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["email"] = args ? args.email : undefined;
            resourceInputs["hasOrganizationProjects"] = args ? args.hasOrganizationProjects : undefined;
            resourceInputs["hasRepositoryProjects"] = args ? args.hasRepositoryProjects : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["membersCanCreateInternalRepositories"] = args ? args.membersCanCreateInternalRepositories : undefined;
            resourceInputs["membersCanCreatePages"] = args ? args.membersCanCreatePages : undefined;
            resourceInputs["membersCanCreatePrivatePages"] = args ? args.membersCanCreatePrivatePages : undefined;
            resourceInputs["membersCanCreatePrivateRepositories"] = args ? args.membersCanCreatePrivateRepositories : undefined;
            resourceInputs["membersCanCreatePublicPages"] = args ? args.membersCanCreatePublicPages : undefined;
            resourceInputs["membersCanCreatePublicRepositories"] = args ? args.membersCanCreatePublicRepositories : undefined;
            resourceInputs["membersCanCreateRepositories"] = args ? args.membersCanCreateRepositories : undefined;
            resourceInputs["membersCanForkPrivateRepositories"] = args ? args.membersCanForkPrivateRepositories : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["secretScanningEnabledForNewRepositories"] = args ? args.secretScanningEnabledForNewRepositories : undefined;
            resourceInputs["secretScanningPushProtectionEnabledForNewRepositories"] = args ? args.secretScanningPushProtectionEnabledForNewRepositories : undefined;
            resourceInputs["twitterUsername"] = args ? args.twitterUsername : undefined;
            resourceInputs["webCommitSignoffRequired"] = args ? args.webCommitSignoffRequired : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(OrganizationSettings.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OrganizationSettings resources.
 */
export interface OrganizationSettingsState {
    advancedSecurityEnabledForNewRepositories?: pulumi.Input<boolean>;
    billingEmail?: pulumi.Input<string>;
    blog?: pulumi.Input<string>;
    company?: pulumi.Input<string>;
    defaultRepositoryPermission?: pulumi.Input<string>;
    dependabotAlertsEnabledForNewRepositories?: pulumi.Input<boolean>;
    dependabotSecurityUpdatesEnabledForNewRepositories?: pulumi.Input<boolean>;
    dependencyGraphEnabledForNewRepositories?: pulumi.Input<boolean>;
    description?: pulumi.Input<string>;
    email?: pulumi.Input<string>;
    hasOrganizationProjects?: pulumi.Input<boolean>;
    hasRepositoryProjects?: pulumi.Input<boolean>;
    location?: pulumi.Input<string>;
    /**
     * Setting to true allows organization members to create internal repositories. Only available to Enterprise Organizations.
     */
    membersCanCreateInternalRepositories?: pulumi.Input<boolean>;
    membersCanCreatePages?: pulumi.Input<boolean>;
    membersCanCreatePrivatePages?: pulumi.Input<boolean>;
    membersCanCreatePrivateRepositories?: pulumi.Input<boolean>;
    membersCanCreatePublicPages?: pulumi.Input<boolean>;
    membersCanCreatePublicRepositories?: pulumi.Input<boolean>;
    membersCanCreateRepositories?: pulumi.Input<boolean>;
    membersCanForkPrivateRepositories?: pulumi.Input<boolean>;
    name?: pulumi.Input<string>;
    secretScanningEnabledForNewRepositories?: pulumi.Input<boolean>;
    secretScanningPushProtectionEnabledForNewRepositories?: pulumi.Input<boolean>;
    twitterUsername?: pulumi.Input<string>;
    webCommitSignoffRequired?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a OrganizationSettings resource.
 */
export interface OrganizationSettingsArgs {
    advancedSecurityEnabledForNewRepositories?: pulumi.Input<boolean>;
    billingEmail: pulumi.Input<string>;
    blog?: pulumi.Input<string>;
    company?: pulumi.Input<string>;
    defaultRepositoryPermission?: pulumi.Input<string>;
    dependabotAlertsEnabledForNewRepositories?: pulumi.Input<boolean>;
    dependabotSecurityUpdatesEnabledForNewRepositories?: pulumi.Input<boolean>;
    dependencyGraphEnabledForNewRepositories?: pulumi.Input<boolean>;
    description?: pulumi.Input<string>;
    email?: pulumi.Input<string>;
    hasOrganizationProjects?: pulumi.Input<boolean>;
    hasRepositoryProjects?: pulumi.Input<boolean>;
    location?: pulumi.Input<string>;
    /**
     * Setting to true allows organization members to create internal repositories. Only available to Enterprise Organizations.
     */
    membersCanCreateInternalRepositories?: pulumi.Input<boolean>;
    membersCanCreatePages?: pulumi.Input<boolean>;
    membersCanCreatePrivatePages?: pulumi.Input<boolean>;
    membersCanCreatePrivateRepositories?: pulumi.Input<boolean>;
    membersCanCreatePublicPages?: pulumi.Input<boolean>;
    membersCanCreatePublicRepositories?: pulumi.Input<boolean>;
    membersCanCreateRepositories?: pulumi.Input<boolean>;
    membersCanForkPrivateRepositories?: pulumi.Input<boolean>;
    name?: pulumi.Input<string>;
    secretScanningEnabledForNewRepositories?: pulumi.Input<boolean>;
    secretScanningPushProtectionEnabledForNewRepositories?: pulumi.Input<boolean>;
    twitterUsername?: pulumi.Input<string>;
    webCommitSignoffRequired?: pulumi.Input<boolean>;
}
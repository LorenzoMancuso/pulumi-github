// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class OrganizationSettingsState extends com.pulumi.resources.ResourceArgs {

    public static final OrganizationSettingsState Empty = new OrganizationSettingsState();

    @Import(name="advancedSecurityEnabledForNewRepositories")
    private @Nullable Output<Boolean> advancedSecurityEnabledForNewRepositories;

    public Optional<Output<Boolean>> advancedSecurityEnabledForNewRepositories() {
        return Optional.ofNullable(this.advancedSecurityEnabledForNewRepositories);
    }

    @Import(name="billingEmail")
    private @Nullable Output<String> billingEmail;

    public Optional<Output<String>> billingEmail() {
        return Optional.ofNullable(this.billingEmail);
    }

    @Import(name="blog")
    private @Nullable Output<String> blog;

    public Optional<Output<String>> blog() {
        return Optional.ofNullable(this.blog);
    }

    @Import(name="company")
    private @Nullable Output<String> company;

    public Optional<Output<String>> company() {
        return Optional.ofNullable(this.company);
    }

    @Import(name="defaultRepositoryPermission")
    private @Nullable Output<String> defaultRepositoryPermission;

    public Optional<Output<String>> defaultRepositoryPermission() {
        return Optional.ofNullable(this.defaultRepositoryPermission);
    }

    @Import(name="dependabotAlertsEnabledForNewRepositories")
    private @Nullable Output<Boolean> dependabotAlertsEnabledForNewRepositories;

    public Optional<Output<Boolean>> dependabotAlertsEnabledForNewRepositories() {
        return Optional.ofNullable(this.dependabotAlertsEnabledForNewRepositories);
    }

    @Import(name="dependabotSecurityUpdatesEnabledForNewRepositories")
    private @Nullable Output<Boolean> dependabotSecurityUpdatesEnabledForNewRepositories;

    public Optional<Output<Boolean>> dependabotSecurityUpdatesEnabledForNewRepositories() {
        return Optional.ofNullable(this.dependabotSecurityUpdatesEnabledForNewRepositories);
    }

    @Import(name="dependencyGraphEnabledForNewRepositories")
    private @Nullable Output<Boolean> dependencyGraphEnabledForNewRepositories;

    public Optional<Output<Boolean>> dependencyGraphEnabledForNewRepositories() {
        return Optional.ofNullable(this.dependencyGraphEnabledForNewRepositories);
    }

    @Import(name="description")
    private @Nullable Output<String> description;

    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    @Import(name="email")
    private @Nullable Output<String> email;

    public Optional<Output<String>> email() {
        return Optional.ofNullable(this.email);
    }

    @Import(name="hasOrganizationProjects")
    private @Nullable Output<Boolean> hasOrganizationProjects;

    public Optional<Output<Boolean>> hasOrganizationProjects() {
        return Optional.ofNullable(this.hasOrganizationProjects);
    }

    @Import(name="hasRepositoryProjects")
    private @Nullable Output<Boolean> hasRepositoryProjects;

    public Optional<Output<Boolean>> hasRepositoryProjects() {
        return Optional.ofNullable(this.hasRepositoryProjects);
    }

    @Import(name="location")
    private @Nullable Output<String> location;

    public Optional<Output<String>> location() {
        return Optional.ofNullable(this.location);
    }

    /**
     * Setting to true allows organization members to create internal repositories. Only available to Enterprise Organizations.
     * 
     */
    @Import(name="membersCanCreateInternalRepositories")
    private @Nullable Output<Boolean> membersCanCreateInternalRepositories;

    /**
     * @return Setting to true allows organization members to create internal repositories. Only available to Enterprise Organizations.
     * 
     */
    public Optional<Output<Boolean>> membersCanCreateInternalRepositories() {
        return Optional.ofNullable(this.membersCanCreateInternalRepositories);
    }

    @Import(name="membersCanCreatePages")
    private @Nullable Output<Boolean> membersCanCreatePages;

    public Optional<Output<Boolean>> membersCanCreatePages() {
        return Optional.ofNullable(this.membersCanCreatePages);
    }

    @Import(name="membersCanCreatePrivatePages")
    private @Nullable Output<Boolean> membersCanCreatePrivatePages;

    public Optional<Output<Boolean>> membersCanCreatePrivatePages() {
        return Optional.ofNullable(this.membersCanCreatePrivatePages);
    }

    @Import(name="membersCanCreatePrivateRepositories")
    private @Nullable Output<Boolean> membersCanCreatePrivateRepositories;

    public Optional<Output<Boolean>> membersCanCreatePrivateRepositories() {
        return Optional.ofNullable(this.membersCanCreatePrivateRepositories);
    }

    @Import(name="membersCanCreatePublicPages")
    private @Nullable Output<Boolean> membersCanCreatePublicPages;

    public Optional<Output<Boolean>> membersCanCreatePublicPages() {
        return Optional.ofNullable(this.membersCanCreatePublicPages);
    }

    @Import(name="membersCanCreatePublicRepositories")
    private @Nullable Output<Boolean> membersCanCreatePublicRepositories;

    public Optional<Output<Boolean>> membersCanCreatePublicRepositories() {
        return Optional.ofNullable(this.membersCanCreatePublicRepositories);
    }

    @Import(name="membersCanCreateRepositories")
    private @Nullable Output<Boolean> membersCanCreateRepositories;

    public Optional<Output<Boolean>> membersCanCreateRepositories() {
        return Optional.ofNullable(this.membersCanCreateRepositories);
    }

    @Import(name="membersCanForkPrivateRepositories")
    private @Nullable Output<Boolean> membersCanForkPrivateRepositories;

    public Optional<Output<Boolean>> membersCanForkPrivateRepositories() {
        return Optional.ofNullable(this.membersCanForkPrivateRepositories);
    }

    @Import(name="name")
    private @Nullable Output<String> name;

    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    @Import(name="secretScanningEnabledForNewRepositories")
    private @Nullable Output<Boolean> secretScanningEnabledForNewRepositories;

    public Optional<Output<Boolean>> secretScanningEnabledForNewRepositories() {
        return Optional.ofNullable(this.secretScanningEnabledForNewRepositories);
    }

    @Import(name="secretScanningPushProtectionEnabledForNewRepositories")
    private @Nullable Output<Boolean> secretScanningPushProtectionEnabledForNewRepositories;

    public Optional<Output<Boolean>> secretScanningPushProtectionEnabledForNewRepositories() {
        return Optional.ofNullable(this.secretScanningPushProtectionEnabledForNewRepositories);
    }

    @Import(name="twitterUsername")
    private @Nullable Output<String> twitterUsername;

    public Optional<Output<String>> twitterUsername() {
        return Optional.ofNullable(this.twitterUsername);
    }

    @Import(name="webCommitSignoffRequired")
    private @Nullable Output<Boolean> webCommitSignoffRequired;

    public Optional<Output<Boolean>> webCommitSignoffRequired() {
        return Optional.ofNullable(this.webCommitSignoffRequired);
    }

    private OrganizationSettingsState() {}

    private OrganizationSettingsState(OrganizationSettingsState $) {
        this.advancedSecurityEnabledForNewRepositories = $.advancedSecurityEnabledForNewRepositories;
        this.billingEmail = $.billingEmail;
        this.blog = $.blog;
        this.company = $.company;
        this.defaultRepositoryPermission = $.defaultRepositoryPermission;
        this.dependabotAlertsEnabledForNewRepositories = $.dependabotAlertsEnabledForNewRepositories;
        this.dependabotSecurityUpdatesEnabledForNewRepositories = $.dependabotSecurityUpdatesEnabledForNewRepositories;
        this.dependencyGraphEnabledForNewRepositories = $.dependencyGraphEnabledForNewRepositories;
        this.description = $.description;
        this.email = $.email;
        this.hasOrganizationProjects = $.hasOrganizationProjects;
        this.hasRepositoryProjects = $.hasRepositoryProjects;
        this.location = $.location;
        this.membersCanCreateInternalRepositories = $.membersCanCreateInternalRepositories;
        this.membersCanCreatePages = $.membersCanCreatePages;
        this.membersCanCreatePrivatePages = $.membersCanCreatePrivatePages;
        this.membersCanCreatePrivateRepositories = $.membersCanCreatePrivateRepositories;
        this.membersCanCreatePublicPages = $.membersCanCreatePublicPages;
        this.membersCanCreatePublicRepositories = $.membersCanCreatePublicRepositories;
        this.membersCanCreateRepositories = $.membersCanCreateRepositories;
        this.membersCanForkPrivateRepositories = $.membersCanForkPrivateRepositories;
        this.name = $.name;
        this.secretScanningEnabledForNewRepositories = $.secretScanningEnabledForNewRepositories;
        this.secretScanningPushProtectionEnabledForNewRepositories = $.secretScanningPushProtectionEnabledForNewRepositories;
        this.twitterUsername = $.twitterUsername;
        this.webCommitSignoffRequired = $.webCommitSignoffRequired;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(OrganizationSettingsState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private OrganizationSettingsState $;

        public Builder() {
            $ = new OrganizationSettingsState();
        }

        public Builder(OrganizationSettingsState defaults) {
            $ = new OrganizationSettingsState(Objects.requireNonNull(defaults));
        }

        public Builder advancedSecurityEnabledForNewRepositories(@Nullable Output<Boolean> advancedSecurityEnabledForNewRepositories) {
            $.advancedSecurityEnabledForNewRepositories = advancedSecurityEnabledForNewRepositories;
            return this;
        }

        public Builder advancedSecurityEnabledForNewRepositories(Boolean advancedSecurityEnabledForNewRepositories) {
            return advancedSecurityEnabledForNewRepositories(Output.of(advancedSecurityEnabledForNewRepositories));
        }

        public Builder billingEmail(@Nullable Output<String> billingEmail) {
            $.billingEmail = billingEmail;
            return this;
        }

        public Builder billingEmail(String billingEmail) {
            return billingEmail(Output.of(billingEmail));
        }

        public Builder blog(@Nullable Output<String> blog) {
            $.blog = blog;
            return this;
        }

        public Builder blog(String blog) {
            return blog(Output.of(blog));
        }

        public Builder company(@Nullable Output<String> company) {
            $.company = company;
            return this;
        }

        public Builder company(String company) {
            return company(Output.of(company));
        }

        public Builder defaultRepositoryPermission(@Nullable Output<String> defaultRepositoryPermission) {
            $.defaultRepositoryPermission = defaultRepositoryPermission;
            return this;
        }

        public Builder defaultRepositoryPermission(String defaultRepositoryPermission) {
            return defaultRepositoryPermission(Output.of(defaultRepositoryPermission));
        }

        public Builder dependabotAlertsEnabledForNewRepositories(@Nullable Output<Boolean> dependabotAlertsEnabledForNewRepositories) {
            $.dependabotAlertsEnabledForNewRepositories = dependabotAlertsEnabledForNewRepositories;
            return this;
        }

        public Builder dependabotAlertsEnabledForNewRepositories(Boolean dependabotAlertsEnabledForNewRepositories) {
            return dependabotAlertsEnabledForNewRepositories(Output.of(dependabotAlertsEnabledForNewRepositories));
        }

        public Builder dependabotSecurityUpdatesEnabledForNewRepositories(@Nullable Output<Boolean> dependabotSecurityUpdatesEnabledForNewRepositories) {
            $.dependabotSecurityUpdatesEnabledForNewRepositories = dependabotSecurityUpdatesEnabledForNewRepositories;
            return this;
        }

        public Builder dependabotSecurityUpdatesEnabledForNewRepositories(Boolean dependabotSecurityUpdatesEnabledForNewRepositories) {
            return dependabotSecurityUpdatesEnabledForNewRepositories(Output.of(dependabotSecurityUpdatesEnabledForNewRepositories));
        }

        public Builder dependencyGraphEnabledForNewRepositories(@Nullable Output<Boolean> dependencyGraphEnabledForNewRepositories) {
            $.dependencyGraphEnabledForNewRepositories = dependencyGraphEnabledForNewRepositories;
            return this;
        }

        public Builder dependencyGraphEnabledForNewRepositories(Boolean dependencyGraphEnabledForNewRepositories) {
            return dependencyGraphEnabledForNewRepositories(Output.of(dependencyGraphEnabledForNewRepositories));
        }

        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        public Builder description(String description) {
            return description(Output.of(description));
        }

        public Builder email(@Nullable Output<String> email) {
            $.email = email;
            return this;
        }

        public Builder email(String email) {
            return email(Output.of(email));
        }

        public Builder hasOrganizationProjects(@Nullable Output<Boolean> hasOrganizationProjects) {
            $.hasOrganizationProjects = hasOrganizationProjects;
            return this;
        }

        public Builder hasOrganizationProjects(Boolean hasOrganizationProjects) {
            return hasOrganizationProjects(Output.of(hasOrganizationProjects));
        }

        public Builder hasRepositoryProjects(@Nullable Output<Boolean> hasRepositoryProjects) {
            $.hasRepositoryProjects = hasRepositoryProjects;
            return this;
        }

        public Builder hasRepositoryProjects(Boolean hasRepositoryProjects) {
            return hasRepositoryProjects(Output.of(hasRepositoryProjects));
        }

        public Builder location(@Nullable Output<String> location) {
            $.location = location;
            return this;
        }

        public Builder location(String location) {
            return location(Output.of(location));
        }

        /**
         * @param membersCanCreateInternalRepositories Setting to true allows organization members to create internal repositories. Only available to Enterprise Organizations.
         * 
         * @return builder
         * 
         */
        public Builder membersCanCreateInternalRepositories(@Nullable Output<Boolean> membersCanCreateInternalRepositories) {
            $.membersCanCreateInternalRepositories = membersCanCreateInternalRepositories;
            return this;
        }

        /**
         * @param membersCanCreateInternalRepositories Setting to true allows organization members to create internal repositories. Only available to Enterprise Organizations.
         * 
         * @return builder
         * 
         */
        public Builder membersCanCreateInternalRepositories(Boolean membersCanCreateInternalRepositories) {
            return membersCanCreateInternalRepositories(Output.of(membersCanCreateInternalRepositories));
        }

        public Builder membersCanCreatePages(@Nullable Output<Boolean> membersCanCreatePages) {
            $.membersCanCreatePages = membersCanCreatePages;
            return this;
        }

        public Builder membersCanCreatePages(Boolean membersCanCreatePages) {
            return membersCanCreatePages(Output.of(membersCanCreatePages));
        }

        public Builder membersCanCreatePrivatePages(@Nullable Output<Boolean> membersCanCreatePrivatePages) {
            $.membersCanCreatePrivatePages = membersCanCreatePrivatePages;
            return this;
        }

        public Builder membersCanCreatePrivatePages(Boolean membersCanCreatePrivatePages) {
            return membersCanCreatePrivatePages(Output.of(membersCanCreatePrivatePages));
        }

        public Builder membersCanCreatePrivateRepositories(@Nullable Output<Boolean> membersCanCreatePrivateRepositories) {
            $.membersCanCreatePrivateRepositories = membersCanCreatePrivateRepositories;
            return this;
        }

        public Builder membersCanCreatePrivateRepositories(Boolean membersCanCreatePrivateRepositories) {
            return membersCanCreatePrivateRepositories(Output.of(membersCanCreatePrivateRepositories));
        }

        public Builder membersCanCreatePublicPages(@Nullable Output<Boolean> membersCanCreatePublicPages) {
            $.membersCanCreatePublicPages = membersCanCreatePublicPages;
            return this;
        }

        public Builder membersCanCreatePublicPages(Boolean membersCanCreatePublicPages) {
            return membersCanCreatePublicPages(Output.of(membersCanCreatePublicPages));
        }

        public Builder membersCanCreatePublicRepositories(@Nullable Output<Boolean> membersCanCreatePublicRepositories) {
            $.membersCanCreatePublicRepositories = membersCanCreatePublicRepositories;
            return this;
        }

        public Builder membersCanCreatePublicRepositories(Boolean membersCanCreatePublicRepositories) {
            return membersCanCreatePublicRepositories(Output.of(membersCanCreatePublicRepositories));
        }

        public Builder membersCanCreateRepositories(@Nullable Output<Boolean> membersCanCreateRepositories) {
            $.membersCanCreateRepositories = membersCanCreateRepositories;
            return this;
        }

        public Builder membersCanCreateRepositories(Boolean membersCanCreateRepositories) {
            return membersCanCreateRepositories(Output.of(membersCanCreateRepositories));
        }

        public Builder membersCanForkPrivateRepositories(@Nullable Output<Boolean> membersCanForkPrivateRepositories) {
            $.membersCanForkPrivateRepositories = membersCanForkPrivateRepositories;
            return this;
        }

        public Builder membersCanForkPrivateRepositories(Boolean membersCanForkPrivateRepositories) {
            return membersCanForkPrivateRepositories(Output.of(membersCanForkPrivateRepositories));
        }

        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        public Builder name(String name) {
            return name(Output.of(name));
        }

        public Builder secretScanningEnabledForNewRepositories(@Nullable Output<Boolean> secretScanningEnabledForNewRepositories) {
            $.secretScanningEnabledForNewRepositories = secretScanningEnabledForNewRepositories;
            return this;
        }

        public Builder secretScanningEnabledForNewRepositories(Boolean secretScanningEnabledForNewRepositories) {
            return secretScanningEnabledForNewRepositories(Output.of(secretScanningEnabledForNewRepositories));
        }

        public Builder secretScanningPushProtectionEnabledForNewRepositories(@Nullable Output<Boolean> secretScanningPushProtectionEnabledForNewRepositories) {
            $.secretScanningPushProtectionEnabledForNewRepositories = secretScanningPushProtectionEnabledForNewRepositories;
            return this;
        }

        public Builder secretScanningPushProtectionEnabledForNewRepositories(Boolean secretScanningPushProtectionEnabledForNewRepositories) {
            return secretScanningPushProtectionEnabledForNewRepositories(Output.of(secretScanningPushProtectionEnabledForNewRepositories));
        }

        public Builder twitterUsername(@Nullable Output<String> twitterUsername) {
            $.twitterUsername = twitterUsername;
            return this;
        }

        public Builder twitterUsername(String twitterUsername) {
            return twitterUsername(Output.of(twitterUsername));
        }

        public Builder webCommitSignoffRequired(@Nullable Output<Boolean> webCommitSignoffRequired) {
            $.webCommitSignoffRequired = webCommitSignoffRequired;
            return this;
        }

        public Builder webCommitSignoffRequired(Boolean webCommitSignoffRequired) {
            return webCommitSignoffRequired(Output.of(webCommitSignoffRequired));
        }

        public OrganizationSettingsState build() {
            return $;
        }
    }

}

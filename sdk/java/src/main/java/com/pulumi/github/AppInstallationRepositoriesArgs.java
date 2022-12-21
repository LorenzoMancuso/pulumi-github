// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class AppInstallationRepositoriesArgs extends com.pulumi.resources.ResourceArgs {

    public static final AppInstallationRepositoriesArgs Empty = new AppInstallationRepositoriesArgs();

    @Import(name="installationId", required=true)
    private Output<String> installationId;

    public Output<String> installationId() {
        return this.installationId;
    }

    @Import(name="selectedRepositories", required=true)
    private Output<List<String>> selectedRepositories;

    public Output<List<String>> selectedRepositories() {
        return this.selectedRepositories;
    }

    private AppInstallationRepositoriesArgs() {}

    private AppInstallationRepositoriesArgs(AppInstallationRepositoriesArgs $) {
        this.installationId = $.installationId;
        this.selectedRepositories = $.selectedRepositories;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AppInstallationRepositoriesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AppInstallationRepositoriesArgs $;

        public Builder() {
            $ = new AppInstallationRepositoriesArgs();
        }

        public Builder(AppInstallationRepositoriesArgs defaults) {
            $ = new AppInstallationRepositoriesArgs(Objects.requireNonNull(defaults));
        }

        public Builder installationId(Output<String> installationId) {
            $.installationId = installationId;
            return this;
        }

        public Builder installationId(String installationId) {
            return installationId(Output.of(installationId));
        }

        public Builder selectedRepositories(Output<List<String>> selectedRepositories) {
            $.selectedRepositories = selectedRepositories;
            return this;
        }

        public Builder selectedRepositories(List<String> selectedRepositories) {
            return selectedRepositories(Output.of(selectedRepositories));
        }

        public Builder selectedRepositories(String... selectedRepositories) {
            return selectedRepositories(List.of(selectedRepositories));
        }

        public AppInstallationRepositoriesArgs build() {
            $.installationId = Objects.requireNonNull($.installationId, "expected parameter 'installationId' to be non-null");
            $.selectedRepositories = Objects.requireNonNull($.selectedRepositories, "expected parameter 'selectedRepositories' to be non-null");
            return $;
        }
    }

}
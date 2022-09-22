// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.util.List;
import java.util.Objects;


public final class ActionsOrganizationPermissionsEnabledRepositoriesConfigArgs extends com.pulumi.resources.ResourceArgs {

    public static final ActionsOrganizationPermissionsEnabledRepositoriesConfigArgs Empty = new ActionsOrganizationPermissionsEnabledRepositoriesConfigArgs();

    @Import(name="repositoryIds", required=true)
    private Output<List<Integer>> repositoryIds;

    public Output<List<Integer>> repositoryIds() {
        return this.repositoryIds;
    }

    private ActionsOrganizationPermissionsEnabledRepositoriesConfigArgs() {}

    private ActionsOrganizationPermissionsEnabledRepositoriesConfigArgs(ActionsOrganizationPermissionsEnabledRepositoriesConfigArgs $) {
        this.repositoryIds = $.repositoryIds;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ActionsOrganizationPermissionsEnabledRepositoriesConfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ActionsOrganizationPermissionsEnabledRepositoriesConfigArgs $;

        public Builder() {
            $ = new ActionsOrganizationPermissionsEnabledRepositoriesConfigArgs();
        }

        public Builder(ActionsOrganizationPermissionsEnabledRepositoriesConfigArgs defaults) {
            $ = new ActionsOrganizationPermissionsEnabledRepositoriesConfigArgs(Objects.requireNonNull(defaults));
        }

        public Builder repositoryIds(Output<List<Integer>> repositoryIds) {
            $.repositoryIds = repositoryIds;
            return this;
        }

        public Builder repositoryIds(List<Integer> repositoryIds) {
            return repositoryIds(Output.of(repositoryIds));
        }

        public Builder repositoryIds(Integer... repositoryIds) {
            return repositoryIds(List.of(repositoryIds));
        }

        public ActionsOrganizationPermissionsEnabledRepositoriesConfigArgs build() {
            $.repositoryIds = Objects.requireNonNull($.repositoryIds, "expected parameter 'repositoryIds' to be non-null");
            return $;
        }
    }

}

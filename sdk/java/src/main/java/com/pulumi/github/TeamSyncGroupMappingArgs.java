// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.github.inputs.TeamSyncGroupMappingGroupArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TeamSyncGroupMappingArgs extends com.pulumi.resources.ResourceArgs {

    public static final TeamSyncGroupMappingArgs Empty = new TeamSyncGroupMappingArgs();

    @Import(name="groups")
    private @Nullable Output<List<TeamSyncGroupMappingGroupArgs>> groups;

    public Optional<Output<List<TeamSyncGroupMappingGroupArgs>>> groups() {
        return Optional.ofNullable(this.groups);
    }

    @Import(name="teamSlug", required=true)
    private Output<String> teamSlug;

    public Output<String> teamSlug() {
        return this.teamSlug;
    }

    private TeamSyncGroupMappingArgs() {}

    private TeamSyncGroupMappingArgs(TeamSyncGroupMappingArgs $) {
        this.groups = $.groups;
        this.teamSlug = $.teamSlug;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TeamSyncGroupMappingArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TeamSyncGroupMappingArgs $;

        public Builder() {
            $ = new TeamSyncGroupMappingArgs();
        }

        public Builder(TeamSyncGroupMappingArgs defaults) {
            $ = new TeamSyncGroupMappingArgs(Objects.requireNonNull(defaults));
        }

        public Builder groups(@Nullable Output<List<TeamSyncGroupMappingGroupArgs>> groups) {
            $.groups = groups;
            return this;
        }

        public Builder groups(List<TeamSyncGroupMappingGroupArgs> groups) {
            return groups(Output.of(groups));
        }

        public Builder groups(TeamSyncGroupMappingGroupArgs... groups) {
            return groups(List.of(groups));
        }

        public Builder teamSlug(Output<String> teamSlug) {
            $.teamSlug = teamSlug;
            return this;
        }

        public Builder teamSlug(String teamSlug) {
            return teamSlug(Output.of(teamSlug));
        }

        public TeamSyncGroupMappingArgs build() {
            $.teamSlug = Objects.requireNonNull($.teamSlug, "expected parameter 'teamSlug' to be non-null");
            return $;
        }
    }

}

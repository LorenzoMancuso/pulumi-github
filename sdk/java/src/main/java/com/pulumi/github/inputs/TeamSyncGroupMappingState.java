// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.github.inputs.TeamSyncGroupMappingGroupArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TeamSyncGroupMappingState extends com.pulumi.resources.ResourceArgs {

    public static final TeamSyncGroupMappingState Empty = new TeamSyncGroupMappingState();

    @Import(name="etag")
    private @Nullable Output<String> etag;

    public Optional<Output<String>> etag() {
        return Optional.ofNullable(this.etag);
    }

    @Import(name="groups")
    private @Nullable Output<List<TeamSyncGroupMappingGroupArgs>> groups;

    public Optional<Output<List<TeamSyncGroupMappingGroupArgs>>> groups() {
        return Optional.ofNullable(this.groups);
    }

    @Import(name="teamSlug")
    private @Nullable Output<String> teamSlug;

    public Optional<Output<String>> teamSlug() {
        return Optional.ofNullable(this.teamSlug);
    }

    private TeamSyncGroupMappingState() {}

    private TeamSyncGroupMappingState(TeamSyncGroupMappingState $) {
        this.etag = $.etag;
        this.groups = $.groups;
        this.teamSlug = $.teamSlug;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TeamSyncGroupMappingState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TeamSyncGroupMappingState $;

        public Builder() {
            $ = new TeamSyncGroupMappingState();
        }

        public Builder(TeamSyncGroupMappingState defaults) {
            $ = new TeamSyncGroupMappingState(Objects.requireNonNull(defaults));
        }

        public Builder etag(@Nullable Output<String> etag) {
            $.etag = etag;
            return this;
        }

        public Builder etag(String etag) {
            return etag(Output.of(etag));
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

        public Builder teamSlug(@Nullable Output<String> teamSlug) {
            $.teamSlug = teamSlug;
            return this;
        }

        public Builder teamSlug(String teamSlug) {
            return teamSlug(Output.of(teamSlug));
        }

        public TeamSyncGroupMappingState build() {
            return $;
        }
    }

}
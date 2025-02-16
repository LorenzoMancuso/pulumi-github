// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TeamRepositoryState extends com.pulumi.resources.ResourceArgs {

    public static final TeamRepositoryState Empty = new TeamRepositoryState();

    @Import(name="etag")
    private @Nullable Output<String> etag;

    public Optional<Output<String>> etag() {
        return Optional.ofNullable(this.etag);
    }

    @Import(name="permission")
    private @Nullable Output<String> permission;

    public Optional<Output<String>> permission() {
        return Optional.ofNullable(this.permission);
    }

    @Import(name="repository")
    private @Nullable Output<String> repository;

    public Optional<Output<String>> repository() {
        return Optional.ofNullable(this.repository);
    }

    /**
     * ID or slug of team
     * 
     */
    @Import(name="teamId")
    private @Nullable Output<String> teamId;

    /**
     * @return ID or slug of team
     * 
     */
    public Optional<Output<String>> teamId() {
        return Optional.ofNullable(this.teamId);
    }

    private TeamRepositoryState() {}

    private TeamRepositoryState(TeamRepositoryState $) {
        this.etag = $.etag;
        this.permission = $.permission;
        this.repository = $.repository;
        this.teamId = $.teamId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TeamRepositoryState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TeamRepositoryState $;

        public Builder() {
            $ = new TeamRepositoryState();
        }

        public Builder(TeamRepositoryState defaults) {
            $ = new TeamRepositoryState(Objects.requireNonNull(defaults));
        }

        public Builder etag(@Nullable Output<String> etag) {
            $.etag = etag;
            return this;
        }

        public Builder etag(String etag) {
            return etag(Output.of(etag));
        }

        public Builder permission(@Nullable Output<String> permission) {
            $.permission = permission;
            return this;
        }

        public Builder permission(String permission) {
            return permission(Output.of(permission));
        }

        public Builder repository(@Nullable Output<String> repository) {
            $.repository = repository;
            return this;
        }

        public Builder repository(String repository) {
            return repository(Output.of(repository));
        }

        /**
         * @param teamId ID or slug of team
         * 
         * @return builder
         * 
         */
        public Builder teamId(@Nullable Output<String> teamId) {
            $.teamId = teamId;
            return this;
        }

        /**
         * @param teamId ID or slug of team
         * 
         * @return builder
         * 
         */
        public Builder teamId(String teamId) {
            return teamId(Output.of(teamId));
        }

        public TeamRepositoryState build() {
            return $;
        }
    }

}

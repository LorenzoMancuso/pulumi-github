// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RepositoryEnvironmentReviewerArgs extends com.pulumi.resources.ResourceArgs {

    public static final RepositoryEnvironmentReviewerArgs Empty = new RepositoryEnvironmentReviewerArgs();

    @Import(name="teams")
    private @Nullable Output<List<Integer>> teams;

    public Optional<Output<List<Integer>>> teams() {
        return Optional.ofNullable(this.teams);
    }

    @Import(name="users")
    private @Nullable Output<List<Integer>> users;

    public Optional<Output<List<Integer>>> users() {
        return Optional.ofNullable(this.users);
    }

    private RepositoryEnvironmentReviewerArgs() {}

    private RepositoryEnvironmentReviewerArgs(RepositoryEnvironmentReviewerArgs $) {
        this.teams = $.teams;
        this.users = $.users;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RepositoryEnvironmentReviewerArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RepositoryEnvironmentReviewerArgs $;

        public Builder() {
            $ = new RepositoryEnvironmentReviewerArgs();
        }

        public Builder(RepositoryEnvironmentReviewerArgs defaults) {
            $ = new RepositoryEnvironmentReviewerArgs(Objects.requireNonNull(defaults));
        }

        public Builder teams(@Nullable Output<List<Integer>> teams) {
            $.teams = teams;
            return this;
        }

        public Builder teams(List<Integer> teams) {
            return teams(Output.of(teams));
        }

        public Builder teams(Integer... teams) {
            return teams(List.of(teams));
        }

        public Builder users(@Nullable Output<List<Integer>> users) {
            $.users = users;
            return this;
        }

        public Builder users(List<Integer> users) {
            return users(Output.of(users));
        }

        public Builder users(Integer... users) {
            return users(List.of(users));
        }

        public RepositoryEnvironmentReviewerArgs build() {
            return $;
        }
    }

}

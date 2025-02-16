// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TeamMembershipArgs extends com.pulumi.resources.ResourceArgs {

    public static final TeamMembershipArgs Empty = new TeamMembershipArgs();

    @Import(name="role")
    private @Nullable Output<String> role;

    public Optional<Output<String>> role() {
        return Optional.ofNullable(this.role);
    }

    @Import(name="teamId", required=true)
    private Output<String> teamId;

    public Output<String> teamId() {
        return this.teamId;
    }

    @Import(name="username", required=true)
    private Output<String> username;

    public Output<String> username() {
        return this.username;
    }

    private TeamMembershipArgs() {}

    private TeamMembershipArgs(TeamMembershipArgs $) {
        this.role = $.role;
        this.teamId = $.teamId;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TeamMembershipArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TeamMembershipArgs $;

        public Builder() {
            $ = new TeamMembershipArgs();
        }

        public Builder(TeamMembershipArgs defaults) {
            $ = new TeamMembershipArgs(Objects.requireNonNull(defaults));
        }

        public Builder role(@Nullable Output<String> role) {
            $.role = role;
            return this;
        }

        public Builder role(String role) {
            return role(Output.of(role));
        }

        public Builder teamId(Output<String> teamId) {
            $.teamId = teamId;
            return this;
        }

        public Builder teamId(String teamId) {
            return teamId(Output.of(teamId));
        }

        public Builder username(Output<String> username) {
            $.username = username;
            return this;
        }

        public Builder username(String username) {
            return username(Output.of(username));
        }

        public TeamMembershipArgs build() {
            $.teamId = Objects.requireNonNull($.teamId, "expected parameter 'teamId' to be non-null");
            $.username = Objects.requireNonNull($.username, "expected parameter 'username' to be non-null");
            return $;
        }
    }

}

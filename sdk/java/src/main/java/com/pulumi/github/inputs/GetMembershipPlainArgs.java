// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetMembershipPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetMembershipPlainArgs Empty = new GetMembershipPlainArgs();

    @Import(name="organization")
    private @Nullable String organization;

    public Optional<String> organization() {
        return Optional.ofNullable(this.organization);
    }

    @Import(name="username", required=true)
    private String username;

    public String username() {
        return this.username;
    }

    private GetMembershipPlainArgs() {}

    private GetMembershipPlainArgs(GetMembershipPlainArgs $) {
        this.organization = $.organization;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetMembershipPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetMembershipPlainArgs $;

        public Builder() {
            $ = new GetMembershipPlainArgs();
        }

        public Builder(GetMembershipPlainArgs defaults) {
            $ = new GetMembershipPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder organization(@Nullable String organization) {
            $.organization = organization;
            return this;
        }

        public Builder username(String username) {
            $.username = username;
            return this;
        }

        public GetMembershipPlainArgs build() {
            $.username = Objects.requireNonNull($.username, "expected parameter 'username' to be non-null");
            return $;
        }
    }

}
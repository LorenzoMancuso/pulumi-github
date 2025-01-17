// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetActionsSecretsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetActionsSecretsPlainArgs Empty = new GetActionsSecretsPlainArgs();

    @Import(name="fullName")
    private @Nullable String fullName;

    public Optional<String> fullName() {
        return Optional.ofNullable(this.fullName);
    }

    @Import(name="name")
    private @Nullable String name;

    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    private GetActionsSecretsPlainArgs() {}

    private GetActionsSecretsPlainArgs(GetActionsSecretsPlainArgs $) {
        this.fullName = $.fullName;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetActionsSecretsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetActionsSecretsPlainArgs $;

        public Builder() {
            $ = new GetActionsSecretsPlainArgs();
        }

        public Builder(GetActionsSecretsPlainArgs defaults) {
            $ = new GetActionsSecretsPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder fullName(@Nullable String fullName) {
            $.fullName = fullName;
            return this;
        }

        public Builder name(@Nullable String name) {
            $.name = name;
            return this;
        }

        public GetActionsSecretsPlainArgs build() {
            return $;
        }
    }

}

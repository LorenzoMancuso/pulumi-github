// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetRepositoryArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetRepositoryArgs Empty = new GetRepositoryArgs();

    @Import(name="description")
    private @Nullable Output<String> description;

    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    @Import(name="fullName")
    private @Nullable Output<String> fullName;

    public Optional<Output<String>> fullName() {
        return Optional.ofNullable(this.fullName);
    }

    @Import(name="homepageUrl")
    private @Nullable Output<String> homepageUrl;

    public Optional<Output<String>> homepageUrl() {
        return Optional.ofNullable(this.homepageUrl);
    }

    @Import(name="name")
    private @Nullable Output<String> name;

    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    private GetRepositoryArgs() {}

    private GetRepositoryArgs(GetRepositoryArgs $) {
        this.description = $.description;
        this.fullName = $.fullName;
        this.homepageUrl = $.homepageUrl;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetRepositoryArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetRepositoryArgs $;

        public Builder() {
            $ = new GetRepositoryArgs();
        }

        public Builder(GetRepositoryArgs defaults) {
            $ = new GetRepositoryArgs(Objects.requireNonNull(defaults));
        }

        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        public Builder description(String description) {
            return description(Output.of(description));
        }

        public Builder fullName(@Nullable Output<String> fullName) {
            $.fullName = fullName;
            return this;
        }

        public Builder fullName(String fullName) {
            return fullName(Output.of(fullName));
        }

        public Builder homepageUrl(@Nullable Output<String> homepageUrl) {
            $.homepageUrl = homepageUrl;
            return this;
        }

        public Builder homepageUrl(String homepageUrl) {
            return homepageUrl(Output.of(homepageUrl));
        }

        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        public Builder name(String name) {
            return name(Output.of(name));
        }

        public GetRepositoryArgs build() {
            return $;
        }
    }

}

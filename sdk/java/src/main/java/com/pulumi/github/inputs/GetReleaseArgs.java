// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetReleaseArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetReleaseArgs Empty = new GetReleaseArgs();

    @Import(name="owner", required=true)
    private Output<String> owner;

    public Output<String> owner() {
        return this.owner;
    }

    @Import(name="releaseId")
    private @Nullable Output<Integer> releaseId;

    public Optional<Output<Integer>> releaseId() {
        return Optional.ofNullable(this.releaseId);
    }

    @Import(name="releaseTag")
    private @Nullable Output<String> releaseTag;

    public Optional<Output<String>> releaseTag() {
        return Optional.ofNullable(this.releaseTag);
    }

    @Import(name="repository", required=true)
    private Output<String> repository;

    public Output<String> repository() {
        return this.repository;
    }

    @Import(name="retrieveBy", required=true)
    private Output<String> retrieveBy;

    public Output<String> retrieveBy() {
        return this.retrieveBy;
    }

    private GetReleaseArgs() {}

    private GetReleaseArgs(GetReleaseArgs $) {
        this.owner = $.owner;
        this.releaseId = $.releaseId;
        this.releaseTag = $.releaseTag;
        this.repository = $.repository;
        this.retrieveBy = $.retrieveBy;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetReleaseArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetReleaseArgs $;

        public Builder() {
            $ = new GetReleaseArgs();
        }

        public Builder(GetReleaseArgs defaults) {
            $ = new GetReleaseArgs(Objects.requireNonNull(defaults));
        }

        public Builder owner(Output<String> owner) {
            $.owner = owner;
            return this;
        }

        public Builder owner(String owner) {
            return owner(Output.of(owner));
        }

        public Builder releaseId(@Nullable Output<Integer> releaseId) {
            $.releaseId = releaseId;
            return this;
        }

        public Builder releaseId(Integer releaseId) {
            return releaseId(Output.of(releaseId));
        }

        public Builder releaseTag(@Nullable Output<String> releaseTag) {
            $.releaseTag = releaseTag;
            return this;
        }

        public Builder releaseTag(String releaseTag) {
            return releaseTag(Output.of(releaseTag));
        }

        public Builder repository(Output<String> repository) {
            $.repository = repository;
            return this;
        }

        public Builder repository(String repository) {
            return repository(Output.of(repository));
        }

        public Builder retrieveBy(Output<String> retrieveBy) {
            $.retrieveBy = retrieveBy;
            return this;
        }

        public Builder retrieveBy(String retrieveBy) {
            return retrieveBy(Output.of(retrieveBy));
        }

        public GetReleaseArgs build() {
            $.owner = Objects.requireNonNull($.owner, "expected parameter 'owner' to be non-null");
            $.repository = Objects.requireNonNull($.repository, "expected parameter 'repository' to be non-null");
            $.retrieveBy = Objects.requireNonNull($.retrieveBy, "expected parameter 'retrieveBy' to be non-null");
            return $;
        }
    }

}

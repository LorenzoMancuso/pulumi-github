// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.github.inputs.BranchProtectionV3RequiredPullRequestReviewsArgs;
import com.pulumi.github.inputs.BranchProtectionV3RequiredStatusChecksArgs;
import com.pulumi.github.inputs.BranchProtectionV3RestrictionsArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BranchProtectionV3State extends com.pulumi.resources.ResourceArgs {

    public static final BranchProtectionV3State Empty = new BranchProtectionV3State();

    @Import(name="branch")
    private @Nullable Output<String> branch;

    public Optional<Output<String>> branch() {
        return Optional.ofNullable(this.branch);
    }

    @Import(name="enforceAdmins")
    private @Nullable Output<Boolean> enforceAdmins;

    public Optional<Output<Boolean>> enforceAdmins() {
        return Optional.ofNullable(this.enforceAdmins);
    }

    @Import(name="etag")
    private @Nullable Output<String> etag;

    public Optional<Output<String>> etag() {
        return Optional.ofNullable(this.etag);
    }

    @Import(name="repository")
    private @Nullable Output<String> repository;

    public Optional<Output<String>> repository() {
        return Optional.ofNullable(this.repository);
    }

    @Import(name="requireConversationResolution")
    private @Nullable Output<Boolean> requireConversationResolution;

    public Optional<Output<Boolean>> requireConversationResolution() {
        return Optional.ofNullable(this.requireConversationResolution);
    }

    @Import(name="requireSignedCommits")
    private @Nullable Output<Boolean> requireSignedCommits;

    public Optional<Output<Boolean>> requireSignedCommits() {
        return Optional.ofNullable(this.requireSignedCommits);
    }

    @Import(name="requiredPullRequestReviews")
    private @Nullable Output<BranchProtectionV3RequiredPullRequestReviewsArgs> requiredPullRequestReviews;

    public Optional<Output<BranchProtectionV3RequiredPullRequestReviewsArgs>> requiredPullRequestReviews() {
        return Optional.ofNullable(this.requiredPullRequestReviews);
    }

    @Import(name="requiredStatusChecks")
    private @Nullable Output<BranchProtectionV3RequiredStatusChecksArgs> requiredStatusChecks;

    public Optional<Output<BranchProtectionV3RequiredStatusChecksArgs>> requiredStatusChecks() {
        return Optional.ofNullable(this.requiredStatusChecks);
    }

    @Import(name="restrictions")
    private @Nullable Output<BranchProtectionV3RestrictionsArgs> restrictions;

    public Optional<Output<BranchProtectionV3RestrictionsArgs>> restrictions() {
        return Optional.ofNullable(this.restrictions);
    }

    private BranchProtectionV3State() {}

    private BranchProtectionV3State(BranchProtectionV3State $) {
        this.branch = $.branch;
        this.enforceAdmins = $.enforceAdmins;
        this.etag = $.etag;
        this.repository = $.repository;
        this.requireConversationResolution = $.requireConversationResolution;
        this.requireSignedCommits = $.requireSignedCommits;
        this.requiredPullRequestReviews = $.requiredPullRequestReviews;
        this.requiredStatusChecks = $.requiredStatusChecks;
        this.restrictions = $.restrictions;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BranchProtectionV3State defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BranchProtectionV3State $;

        public Builder() {
            $ = new BranchProtectionV3State();
        }

        public Builder(BranchProtectionV3State defaults) {
            $ = new BranchProtectionV3State(Objects.requireNonNull(defaults));
        }

        public Builder branch(@Nullable Output<String> branch) {
            $.branch = branch;
            return this;
        }

        public Builder branch(String branch) {
            return branch(Output.of(branch));
        }

        public Builder enforceAdmins(@Nullable Output<Boolean> enforceAdmins) {
            $.enforceAdmins = enforceAdmins;
            return this;
        }

        public Builder enforceAdmins(Boolean enforceAdmins) {
            return enforceAdmins(Output.of(enforceAdmins));
        }

        public Builder etag(@Nullable Output<String> etag) {
            $.etag = etag;
            return this;
        }

        public Builder etag(String etag) {
            return etag(Output.of(etag));
        }

        public Builder repository(@Nullable Output<String> repository) {
            $.repository = repository;
            return this;
        }

        public Builder repository(String repository) {
            return repository(Output.of(repository));
        }

        public Builder requireConversationResolution(@Nullable Output<Boolean> requireConversationResolution) {
            $.requireConversationResolution = requireConversationResolution;
            return this;
        }

        public Builder requireConversationResolution(Boolean requireConversationResolution) {
            return requireConversationResolution(Output.of(requireConversationResolution));
        }

        public Builder requireSignedCommits(@Nullable Output<Boolean> requireSignedCommits) {
            $.requireSignedCommits = requireSignedCommits;
            return this;
        }

        public Builder requireSignedCommits(Boolean requireSignedCommits) {
            return requireSignedCommits(Output.of(requireSignedCommits));
        }

        public Builder requiredPullRequestReviews(@Nullable Output<BranchProtectionV3RequiredPullRequestReviewsArgs> requiredPullRequestReviews) {
            $.requiredPullRequestReviews = requiredPullRequestReviews;
            return this;
        }

        public Builder requiredPullRequestReviews(BranchProtectionV3RequiredPullRequestReviewsArgs requiredPullRequestReviews) {
            return requiredPullRequestReviews(Output.of(requiredPullRequestReviews));
        }

        public Builder requiredStatusChecks(@Nullable Output<BranchProtectionV3RequiredStatusChecksArgs> requiredStatusChecks) {
            $.requiredStatusChecks = requiredStatusChecks;
            return this;
        }

        public Builder requiredStatusChecks(BranchProtectionV3RequiredStatusChecksArgs requiredStatusChecks) {
            return requiredStatusChecks(Output.of(requiredStatusChecks));
        }

        public Builder restrictions(@Nullable Output<BranchProtectionV3RestrictionsArgs> restrictions) {
            $.restrictions = restrictions;
            return this;
        }

        public Builder restrictions(BranchProtectionV3RestrictionsArgs restrictions) {
            return restrictions(Output.of(restrictions));
        }

        public BranchProtectionV3State build() {
            return $;
        }
    }

}

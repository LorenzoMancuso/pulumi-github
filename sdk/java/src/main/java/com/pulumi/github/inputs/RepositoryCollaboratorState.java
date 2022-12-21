// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RepositoryCollaboratorState extends com.pulumi.resources.ResourceArgs {

    public static final RepositoryCollaboratorState Empty = new RepositoryCollaboratorState();

    @Import(name="invitationId")
    private @Nullable Output<String> invitationId;

    public Optional<Output<String>> invitationId() {
        return Optional.ofNullable(this.invitationId);
    }

    @Import(name="permission")
    private @Nullable Output<String> permission;

    public Optional<Output<String>> permission() {
        return Optional.ofNullable(this.permission);
    }

    @Import(name="permissionDiffSuppression")
    private @Nullable Output<Boolean> permissionDiffSuppression;

    public Optional<Output<Boolean>> permissionDiffSuppression() {
        return Optional.ofNullable(this.permissionDiffSuppression);
    }

    @Import(name="repository")
    private @Nullable Output<String> repository;

    public Optional<Output<String>> repository() {
        return Optional.ofNullable(this.repository);
    }

    @Import(name="username")
    private @Nullable Output<String> username;

    public Optional<Output<String>> username() {
        return Optional.ofNullable(this.username);
    }

    private RepositoryCollaboratorState() {}

    private RepositoryCollaboratorState(RepositoryCollaboratorState $) {
        this.invitationId = $.invitationId;
        this.permission = $.permission;
        this.permissionDiffSuppression = $.permissionDiffSuppression;
        this.repository = $.repository;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RepositoryCollaboratorState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RepositoryCollaboratorState $;

        public Builder() {
            $ = new RepositoryCollaboratorState();
        }

        public Builder(RepositoryCollaboratorState defaults) {
            $ = new RepositoryCollaboratorState(Objects.requireNonNull(defaults));
        }

        public Builder invitationId(@Nullable Output<String> invitationId) {
            $.invitationId = invitationId;
            return this;
        }

        public Builder invitationId(String invitationId) {
            return invitationId(Output.of(invitationId));
        }

        public Builder permission(@Nullable Output<String> permission) {
            $.permission = permission;
            return this;
        }

        public Builder permission(String permission) {
            return permission(Output.of(permission));
        }

        public Builder permissionDiffSuppression(@Nullable Output<Boolean> permissionDiffSuppression) {
            $.permissionDiffSuppression = permissionDiffSuppression;
            return this;
        }

        public Builder permissionDiffSuppression(Boolean permissionDiffSuppression) {
            return permissionDiffSuppression(Output.of(permissionDiffSuppression));
        }

        public Builder repository(@Nullable Output<String> repository) {
            $.repository = repository;
            return this;
        }

        public Builder repository(String repository) {
            return repository(Output.of(repository));
        }

        public Builder username(@Nullable Output<String> username) {
            $.username = username;
            return this;
        }

        public Builder username(String username) {
            return username(Output.of(username));
        }

        public RepositoryCollaboratorState build() {
            return $;
        }
    }

}
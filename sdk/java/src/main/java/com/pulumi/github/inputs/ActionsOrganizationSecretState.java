// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ActionsOrganizationSecretState extends com.pulumi.resources.ResourceArgs {

    public static final ActionsOrganizationSecretState Empty = new ActionsOrganizationSecretState();

    @Import(name="createdAt")
    private @Nullable Output<String> createdAt;

    public Optional<Output<String>> createdAt() {
        return Optional.ofNullable(this.createdAt);
    }

    @Import(name="encryptedValue")
    private @Nullable Output<String> encryptedValue;

    public Optional<Output<String>> encryptedValue() {
        return Optional.ofNullable(this.encryptedValue);
    }

    @Import(name="plaintextValue")
    private @Nullable Output<String> plaintextValue;

    public Optional<Output<String>> plaintextValue() {
        return Optional.ofNullable(this.plaintextValue);
    }

    @Import(name="secretName")
    private @Nullable Output<String> secretName;

    public Optional<Output<String>> secretName() {
        return Optional.ofNullable(this.secretName);
    }

    @Import(name="selectedRepositoryIds")
    private @Nullable Output<List<Integer>> selectedRepositoryIds;

    public Optional<Output<List<Integer>>> selectedRepositoryIds() {
        return Optional.ofNullable(this.selectedRepositoryIds);
    }

    @Import(name="updatedAt")
    private @Nullable Output<String> updatedAt;

    public Optional<Output<String>> updatedAt() {
        return Optional.ofNullable(this.updatedAt);
    }

    @Import(name="visibility")
    private @Nullable Output<String> visibility;

    public Optional<Output<String>> visibility() {
        return Optional.ofNullable(this.visibility);
    }

    private ActionsOrganizationSecretState() {}

    private ActionsOrganizationSecretState(ActionsOrganizationSecretState $) {
        this.createdAt = $.createdAt;
        this.encryptedValue = $.encryptedValue;
        this.plaintextValue = $.plaintextValue;
        this.secretName = $.secretName;
        this.selectedRepositoryIds = $.selectedRepositoryIds;
        this.updatedAt = $.updatedAt;
        this.visibility = $.visibility;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ActionsOrganizationSecretState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ActionsOrganizationSecretState $;

        public Builder() {
            $ = new ActionsOrganizationSecretState();
        }

        public Builder(ActionsOrganizationSecretState defaults) {
            $ = new ActionsOrganizationSecretState(Objects.requireNonNull(defaults));
        }

        public Builder createdAt(@Nullable Output<String> createdAt) {
            $.createdAt = createdAt;
            return this;
        }

        public Builder createdAt(String createdAt) {
            return createdAt(Output.of(createdAt));
        }

        public Builder encryptedValue(@Nullable Output<String> encryptedValue) {
            $.encryptedValue = encryptedValue;
            return this;
        }

        public Builder encryptedValue(String encryptedValue) {
            return encryptedValue(Output.of(encryptedValue));
        }

        public Builder plaintextValue(@Nullable Output<String> plaintextValue) {
            $.plaintextValue = plaintextValue;
            return this;
        }

        public Builder plaintextValue(String plaintextValue) {
            return plaintextValue(Output.of(plaintextValue));
        }

        public Builder secretName(@Nullable Output<String> secretName) {
            $.secretName = secretName;
            return this;
        }

        public Builder secretName(String secretName) {
            return secretName(Output.of(secretName));
        }

        public Builder selectedRepositoryIds(@Nullable Output<List<Integer>> selectedRepositoryIds) {
            $.selectedRepositoryIds = selectedRepositoryIds;
            return this;
        }

        public Builder selectedRepositoryIds(List<Integer> selectedRepositoryIds) {
            return selectedRepositoryIds(Output.of(selectedRepositoryIds));
        }

        public Builder selectedRepositoryIds(Integer... selectedRepositoryIds) {
            return selectedRepositoryIds(List.of(selectedRepositoryIds));
        }

        public Builder updatedAt(@Nullable Output<String> updatedAt) {
            $.updatedAt = updatedAt;
            return this;
        }

        public Builder updatedAt(String updatedAt) {
            return updatedAt(Output.of(updatedAt));
        }

        public Builder visibility(@Nullable Output<String> visibility) {
            $.visibility = visibility;
            return this;
        }

        public Builder visibility(String visibility) {
            return visibility(Output.of(visibility));
        }

        public ActionsOrganizationSecretState build() {
            return $;
        }
    }

}

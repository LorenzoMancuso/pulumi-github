// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ActionsEnvironmentSecretArgs extends com.pulumi.resources.ResourceArgs {

    public static final ActionsEnvironmentSecretArgs Empty = new ActionsEnvironmentSecretArgs();

    @Import(name="encryptedValue")
    private @Nullable Output<String> encryptedValue;

    public Optional<Output<String>> encryptedValue() {
        return Optional.ofNullable(this.encryptedValue);
    }

    @Import(name="environment", required=true)
    private Output<String> environment;

    public Output<String> environment() {
        return this.environment;
    }

    @Import(name="plaintextValue")
    private @Nullable Output<String> plaintextValue;

    public Optional<Output<String>> plaintextValue() {
        return Optional.ofNullable(this.plaintextValue);
    }

    @Import(name="repository", required=true)
    private Output<String> repository;

    public Output<String> repository() {
        return this.repository;
    }

    @Import(name="secretName", required=true)
    private Output<String> secretName;

    public Output<String> secretName() {
        return this.secretName;
    }

    private ActionsEnvironmentSecretArgs() {}

    private ActionsEnvironmentSecretArgs(ActionsEnvironmentSecretArgs $) {
        this.encryptedValue = $.encryptedValue;
        this.environment = $.environment;
        this.plaintextValue = $.plaintextValue;
        this.repository = $.repository;
        this.secretName = $.secretName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ActionsEnvironmentSecretArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ActionsEnvironmentSecretArgs $;

        public Builder() {
            $ = new ActionsEnvironmentSecretArgs();
        }

        public Builder(ActionsEnvironmentSecretArgs defaults) {
            $ = new ActionsEnvironmentSecretArgs(Objects.requireNonNull(defaults));
        }

        public Builder encryptedValue(@Nullable Output<String> encryptedValue) {
            $.encryptedValue = encryptedValue;
            return this;
        }

        public Builder encryptedValue(String encryptedValue) {
            return encryptedValue(Output.of(encryptedValue));
        }

        public Builder environment(Output<String> environment) {
            $.environment = environment;
            return this;
        }

        public Builder environment(String environment) {
            return environment(Output.of(environment));
        }

        public Builder plaintextValue(@Nullable Output<String> plaintextValue) {
            $.plaintextValue = plaintextValue;
            return this;
        }

        public Builder plaintextValue(String plaintextValue) {
            return plaintextValue(Output.of(plaintextValue));
        }

        public Builder repository(Output<String> repository) {
            $.repository = repository;
            return this;
        }

        public Builder repository(String repository) {
            return repository(Output.of(repository));
        }

        public Builder secretName(Output<String> secretName) {
            $.secretName = secretName;
            return this;
        }

        public Builder secretName(String secretName) {
            return secretName(Output.of(secretName));
        }

        public ActionsEnvironmentSecretArgs build() {
            $.environment = Objects.requireNonNull($.environment, "expected parameter 'environment' to be non-null");
            $.repository = Objects.requireNonNull($.repository, "expected parameter 'repository' to be non-null");
            $.secretName = Objects.requireNonNull($.secretName, "expected parameter 'secretName' to be non-null");
            return $;
        }
    }

}

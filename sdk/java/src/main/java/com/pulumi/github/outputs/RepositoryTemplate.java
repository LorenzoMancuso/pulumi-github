// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class RepositoryTemplate {
    private @Nullable Boolean includeAllBranches;
    private String owner;
    private String repository;

    private RepositoryTemplate() {}
    public Optional<Boolean> includeAllBranches() {
        return Optional.ofNullable(this.includeAllBranches);
    }
    public String owner() {
        return this.owner;
    }
    public String repository() {
        return this.repository;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RepositoryTemplate defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean includeAllBranches;
        private String owner;
        private String repository;
        public Builder() {}
        public Builder(RepositoryTemplate defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.includeAllBranches = defaults.includeAllBranches;
    	      this.owner = defaults.owner;
    	      this.repository = defaults.repository;
        }

        @CustomType.Setter
        public Builder includeAllBranches(@Nullable Boolean includeAllBranches) {
            this.includeAllBranches = includeAllBranches;
            return this;
        }
        @CustomType.Setter
        public Builder owner(String owner) {
            this.owner = Objects.requireNonNull(owner);
            return this;
        }
        @CustomType.Setter
        public Builder repository(String repository) {
            this.repository = Objects.requireNonNull(repository);
            return this;
        }
        public RepositoryTemplate build() {
            final var o = new RepositoryTemplate();
            o.includeAllBranches = includeAllBranches;
            o.owner = owner;
            o.repository = repository;
            return o;
        }
    }
}
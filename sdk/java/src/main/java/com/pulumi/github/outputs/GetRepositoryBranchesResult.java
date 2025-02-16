// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.github.outputs.GetRepositoryBranchesBranch;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetRepositoryBranchesResult {
    private List<GetRepositoryBranchesBranch> branches;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private @Nullable Boolean onlyNonProtectedBranches;
    private @Nullable Boolean onlyProtectedBranches;
    private String repository;

    private GetRepositoryBranchesResult() {}
    public List<GetRepositoryBranchesBranch> branches() {
        return this.branches;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Optional<Boolean> onlyNonProtectedBranches() {
        return Optional.ofNullable(this.onlyNonProtectedBranches);
    }
    public Optional<Boolean> onlyProtectedBranches() {
        return Optional.ofNullable(this.onlyProtectedBranches);
    }
    public String repository() {
        return this.repository;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRepositoryBranchesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetRepositoryBranchesBranch> branches;
        private String id;
        private @Nullable Boolean onlyNonProtectedBranches;
        private @Nullable Boolean onlyProtectedBranches;
        private String repository;
        public Builder() {}
        public Builder(GetRepositoryBranchesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.branches = defaults.branches;
    	      this.id = defaults.id;
    	      this.onlyNonProtectedBranches = defaults.onlyNonProtectedBranches;
    	      this.onlyProtectedBranches = defaults.onlyProtectedBranches;
    	      this.repository = defaults.repository;
        }

        @CustomType.Setter
        public Builder branches(List<GetRepositoryBranchesBranch> branches) {
            this.branches = Objects.requireNonNull(branches);
            return this;
        }
        public Builder branches(GetRepositoryBranchesBranch... branches) {
            return branches(List.of(branches));
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder onlyNonProtectedBranches(@Nullable Boolean onlyNonProtectedBranches) {
            this.onlyNonProtectedBranches = onlyNonProtectedBranches;
            return this;
        }
        @CustomType.Setter
        public Builder onlyProtectedBranches(@Nullable Boolean onlyProtectedBranches) {
            this.onlyProtectedBranches = onlyProtectedBranches;
            return this;
        }
        @CustomType.Setter
        public Builder repository(String repository) {
            this.repository = Objects.requireNonNull(repository);
            return this;
        }
        public GetRepositoryBranchesResult build() {
            final var o = new GetRepositoryBranchesResult();
            o.branches = branches;
            o.id = id;
            o.onlyNonProtectedBranches = onlyNonProtectedBranches;
            o.onlyProtectedBranches = onlyProtectedBranches;
            o.repository = repository;
            return o;
        }
    }
}

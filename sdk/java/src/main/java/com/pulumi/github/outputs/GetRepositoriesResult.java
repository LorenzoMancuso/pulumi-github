// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetRepositoriesResult {
    private List<String> fullNames;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private @Nullable Boolean includeRepoId;
    private List<String> names;
    private String query;
    private List<Integer> repoIds;
    private @Nullable Integer resultsPerPage;
    private @Nullable String sort;

    private GetRepositoriesResult() {}
    public List<String> fullNames() {
        return this.fullNames;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Optional<Boolean> includeRepoId() {
        return Optional.ofNullable(this.includeRepoId);
    }
    public List<String> names() {
        return this.names;
    }
    public String query() {
        return this.query;
    }
    public List<Integer> repoIds() {
        return this.repoIds;
    }
    public Optional<Integer> resultsPerPage() {
        return Optional.ofNullable(this.resultsPerPage);
    }
    public Optional<String> sort() {
        return Optional.ofNullable(this.sort);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRepositoriesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> fullNames;
        private String id;
        private @Nullable Boolean includeRepoId;
        private List<String> names;
        private String query;
        private List<Integer> repoIds;
        private @Nullable Integer resultsPerPage;
        private @Nullable String sort;
        public Builder() {}
        public Builder(GetRepositoriesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.fullNames = defaults.fullNames;
    	      this.id = defaults.id;
    	      this.includeRepoId = defaults.includeRepoId;
    	      this.names = defaults.names;
    	      this.query = defaults.query;
    	      this.repoIds = defaults.repoIds;
    	      this.resultsPerPage = defaults.resultsPerPage;
    	      this.sort = defaults.sort;
        }

        @CustomType.Setter
        public Builder fullNames(List<String> fullNames) {
            this.fullNames = Objects.requireNonNull(fullNames);
            return this;
        }
        public Builder fullNames(String... fullNames) {
            return fullNames(List.of(fullNames));
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder includeRepoId(@Nullable Boolean includeRepoId) {
            this.includeRepoId = includeRepoId;
            return this;
        }
        @CustomType.Setter
        public Builder names(List<String> names) {
            this.names = Objects.requireNonNull(names);
            return this;
        }
        public Builder names(String... names) {
            return names(List.of(names));
        }
        @CustomType.Setter
        public Builder query(String query) {
            this.query = Objects.requireNonNull(query);
            return this;
        }
        @CustomType.Setter
        public Builder repoIds(List<Integer> repoIds) {
            this.repoIds = Objects.requireNonNull(repoIds);
            return this;
        }
        public Builder repoIds(Integer... repoIds) {
            return repoIds(List.of(repoIds));
        }
        @CustomType.Setter
        public Builder resultsPerPage(@Nullable Integer resultsPerPage) {
            this.resultsPerPage = resultsPerPage;
            return this;
        }
        @CustomType.Setter
        public Builder sort(@Nullable String sort) {
            this.sort = sort;
            return this;
        }
        public GetRepositoriesResult build() {
            final var o = new GetRepositoriesResult();
            o.fullNames = fullNames;
            o.id = id;
            o.includeRepoId = includeRepoId;
            o.names = names;
            o.query = query;
            o.repoIds = repoIds;
            o.resultsPerPage = resultsPerPage;
            o.sort = sort;
            return o;
        }
    }
}
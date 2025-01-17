// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetRepositoriesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetRepositoriesPlainArgs Empty = new GetRepositoriesPlainArgs();

    @Import(name="includeRepoId")
    private @Nullable Boolean includeRepoId;

    public Optional<Boolean> includeRepoId() {
        return Optional.ofNullable(this.includeRepoId);
    }

    @Import(name="query", required=true)
    private String query;

    public String query() {
        return this.query;
    }

    @Import(name="resultsPerPage")
    private @Nullable Integer resultsPerPage;

    public Optional<Integer> resultsPerPage() {
        return Optional.ofNullable(this.resultsPerPage);
    }

    @Import(name="sort")
    private @Nullable String sort;

    public Optional<String> sort() {
        return Optional.ofNullable(this.sort);
    }

    private GetRepositoriesPlainArgs() {}

    private GetRepositoriesPlainArgs(GetRepositoriesPlainArgs $) {
        this.includeRepoId = $.includeRepoId;
        this.query = $.query;
        this.resultsPerPage = $.resultsPerPage;
        this.sort = $.sort;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetRepositoriesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetRepositoriesPlainArgs $;

        public Builder() {
            $ = new GetRepositoriesPlainArgs();
        }

        public Builder(GetRepositoriesPlainArgs defaults) {
            $ = new GetRepositoriesPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder includeRepoId(@Nullable Boolean includeRepoId) {
            $.includeRepoId = includeRepoId;
            return this;
        }

        public Builder query(String query) {
            $.query = query;
            return this;
        }

        public Builder resultsPerPage(@Nullable Integer resultsPerPage) {
            $.resultsPerPage = resultsPerPage;
            return this;
        }

        public Builder sort(@Nullable String sort) {
            $.sort = sort;
            return this;
        }

        public GetRepositoriesPlainArgs build() {
            $.query = Objects.requireNonNull($.query, "expected parameter 'query' to be non-null");
            return $;
        }
    }

}

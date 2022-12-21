// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetRepositoriesArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetRepositoriesArgs Empty = new GetRepositoriesArgs();

    @Import(name="includeRepoId")
    private @Nullable Output<Boolean> includeRepoId;

    public Optional<Output<Boolean>> includeRepoId() {
        return Optional.ofNullable(this.includeRepoId);
    }

    @Import(name="query", required=true)
    private Output<String> query;

    public Output<String> query() {
        return this.query;
    }

    @Import(name="resultsPerPage")
    private @Nullable Output<Integer> resultsPerPage;

    public Optional<Output<Integer>> resultsPerPage() {
        return Optional.ofNullable(this.resultsPerPage);
    }

    @Import(name="sort")
    private @Nullable Output<String> sort;

    public Optional<Output<String>> sort() {
        return Optional.ofNullable(this.sort);
    }

    private GetRepositoriesArgs() {}

    private GetRepositoriesArgs(GetRepositoriesArgs $) {
        this.includeRepoId = $.includeRepoId;
        this.query = $.query;
        this.resultsPerPage = $.resultsPerPage;
        this.sort = $.sort;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetRepositoriesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetRepositoriesArgs $;

        public Builder() {
            $ = new GetRepositoriesArgs();
        }

        public Builder(GetRepositoriesArgs defaults) {
            $ = new GetRepositoriesArgs(Objects.requireNonNull(defaults));
        }

        public Builder includeRepoId(@Nullable Output<Boolean> includeRepoId) {
            $.includeRepoId = includeRepoId;
            return this;
        }

        public Builder includeRepoId(Boolean includeRepoId) {
            return includeRepoId(Output.of(includeRepoId));
        }

        public Builder query(Output<String> query) {
            $.query = query;
            return this;
        }

        public Builder query(String query) {
            return query(Output.of(query));
        }

        public Builder resultsPerPage(@Nullable Output<Integer> resultsPerPage) {
            $.resultsPerPage = resultsPerPage;
            return this;
        }

        public Builder resultsPerPage(Integer resultsPerPage) {
            return resultsPerPage(Output.of(resultsPerPage));
        }

        public Builder sort(@Nullable Output<String> sort) {
            $.sort = sort;
            return this;
        }

        public Builder sort(String sort) {
            return sort(Output.of(sort));
        }

        public GetRepositoriesArgs build() {
            $.query = Objects.requireNonNull($.query, "expected parameter 'query' to be non-null");
            return $;
        }
    }

}
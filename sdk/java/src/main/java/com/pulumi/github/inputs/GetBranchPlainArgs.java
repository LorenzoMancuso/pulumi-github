// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetBranchPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetBranchPlainArgs Empty = new GetBranchPlainArgs();

    @Import(name="branch", required=true)
    private String branch;

    public String branch() {
        return this.branch;
    }

    @Import(name="repository", required=true)
    private String repository;

    public String repository() {
        return this.repository;
    }

    private GetBranchPlainArgs() {}

    private GetBranchPlainArgs(GetBranchPlainArgs $) {
        this.branch = $.branch;
        this.repository = $.repository;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetBranchPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetBranchPlainArgs $;

        public Builder() {
            $ = new GetBranchPlainArgs();
        }

        public Builder(GetBranchPlainArgs defaults) {
            $ = new GetBranchPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder branch(String branch) {
            $.branch = branch;
            return this;
        }

        public Builder repository(String repository) {
            $.repository = repository;
            return this;
        }

        public GetBranchPlainArgs build() {
            $.branch = Objects.requireNonNull($.branch, "expected parameter 'branch' to be non-null");
            $.repository = Objects.requireNonNull($.repository, "expected parameter 'repository' to be non-null");
            return $;
        }
    }

}

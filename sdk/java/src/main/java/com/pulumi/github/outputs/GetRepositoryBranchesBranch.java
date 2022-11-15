// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetRepositoryBranchesBranch {
    private String name;
    private Boolean protected_;

    private GetRepositoryBranchesBranch() {}
    public String name() {
        return this.name;
    }
    public Boolean protected_() {
        return this.protected_;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRepositoryBranchesBranch defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String name;
        private Boolean protected_;
        public Builder() {}
        public Builder(GetRepositoryBranchesBranch defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.name = defaults.name;
    	      this.protected_ = defaults.protected_;
        }

        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter("protected")
        public Builder protected_(Boolean protected_) {
            this.protected_ = Objects.requireNonNull(protected_);
            return this;
        }
        public GetRepositoryBranchesBranch build() {
            final var o = new GetRepositoryBranchesBranch();
            o.name = name;
            o.protected_ = protected_;
            return o;
        }
    }
}

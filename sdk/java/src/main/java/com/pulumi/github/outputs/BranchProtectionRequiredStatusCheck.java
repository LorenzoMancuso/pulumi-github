// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class BranchProtectionRequiredStatusCheck {
    private @Nullable List<String> contexts;
    private @Nullable Boolean strict;

    private BranchProtectionRequiredStatusCheck() {}
    public List<String> contexts() {
        return this.contexts == null ? List.of() : this.contexts;
    }
    public Optional<Boolean> strict() {
        return Optional.ofNullable(this.strict);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(BranchProtectionRequiredStatusCheck defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<String> contexts;
        private @Nullable Boolean strict;
        public Builder() {}
        public Builder(BranchProtectionRequiredStatusCheck defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.contexts = defaults.contexts;
    	      this.strict = defaults.strict;
        }

        @CustomType.Setter
        public Builder contexts(@Nullable List<String> contexts) {
            this.contexts = contexts;
            return this;
        }
        public Builder contexts(String... contexts) {
            return contexts(List.of(contexts));
        }
        @CustomType.Setter
        public Builder strict(@Nullable Boolean strict) {
            this.strict = strict;
            return this;
        }
        public BranchProtectionRequiredStatusCheck build() {
            final var o = new BranchProtectionRequiredStatusCheck();
            o.contexts = contexts;
            o.strict = strict;
            return o;
        }
    }
}

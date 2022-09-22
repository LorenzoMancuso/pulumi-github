// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetRepositoryTeamsTeam {
    private String name;
    private String permission;
    private String slug;

    private GetRepositoryTeamsTeam() {}
    public String name() {
        return this.name;
    }
    public String permission() {
        return this.permission;
    }
    public String slug() {
        return this.slug;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRepositoryTeamsTeam defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String name;
        private String permission;
        private String slug;
        public Builder() {}
        public Builder(GetRepositoryTeamsTeam defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.name = defaults.name;
    	      this.permission = defaults.permission;
    	      this.slug = defaults.slug;
        }

        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder permission(String permission) {
            this.permission = Objects.requireNonNull(permission);
            return this;
        }
        @CustomType.Setter
        public Builder slug(String slug) {
            this.slug = Objects.requireNonNull(slug);
            return this;
        }
        public GetRepositoryTeamsTeam build() {
            final var o = new GetRepositoryTeamsTeam();
            o.name = name;
            o.permission = permission;
            o.slug = slug;
            return o;
        }
    }
}

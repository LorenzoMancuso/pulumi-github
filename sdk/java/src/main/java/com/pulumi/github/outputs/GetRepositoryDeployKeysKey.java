// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetRepositoryDeployKeysKey {
    private Integer id;
    private String key;
    private String title;
    private Boolean verified;

    private GetRepositoryDeployKeysKey() {}
    public Integer id() {
        return this.id;
    }
    public String key() {
        return this.key;
    }
    public String title() {
        return this.title;
    }
    public Boolean verified() {
        return this.verified;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRepositoryDeployKeysKey defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer id;
        private String key;
        private String title;
        private Boolean verified;
        public Builder() {}
        public Builder(GetRepositoryDeployKeysKey defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.key = defaults.key;
    	      this.title = defaults.title;
    	      this.verified = defaults.verified;
        }

        @CustomType.Setter
        public Builder id(Integer id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder key(String key) {
            this.key = Objects.requireNonNull(key);
            return this;
        }
        @CustomType.Setter
        public Builder title(String title) {
            this.title = Objects.requireNonNull(title);
            return this;
        }
        @CustomType.Setter
        public Builder verified(Boolean verified) {
            this.verified = Objects.requireNonNull(verified);
            return this;
        }
        public GetRepositoryDeployKeysKey build() {
            final var o = new GetRepositoryDeployKeysKey();
            o.id = id;
            o.key = key;
            o.title = title;
            o.verified = verified;
            return o;
        }
    }
}
// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.github.outputs.GetDependabotSecretsSecret;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetDependabotSecretsResult {
    private String fullName;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String name;
    private List<GetDependabotSecretsSecret> secrets;

    private GetDependabotSecretsResult() {}
    public String fullName() {
        return this.fullName;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public String name() {
        return this.name;
    }
    public List<GetDependabotSecretsSecret> secrets() {
        return this.secrets;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDependabotSecretsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String fullName;
        private String id;
        private String name;
        private List<GetDependabotSecretsSecret> secrets;
        public Builder() {}
        public Builder(GetDependabotSecretsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.fullName = defaults.fullName;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.secrets = defaults.secrets;
        }

        @CustomType.Setter
        public Builder fullName(String fullName) {
            this.fullName = Objects.requireNonNull(fullName);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder secrets(List<GetDependabotSecretsSecret> secrets) {
            this.secrets = Objects.requireNonNull(secrets);
            return this;
        }
        public Builder secrets(GetDependabotSecretsSecret... secrets) {
            return secrets(List.of(secrets));
        }
        public GetDependabotSecretsResult build() {
            final var o = new GetDependabotSecretsResult();
            o.fullName = fullName;
            o.id = id;
            o.name = name;
            o.secrets = secrets;
            return o;
        }
    }
}
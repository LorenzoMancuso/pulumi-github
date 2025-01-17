// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetRepositoryMilestoneResult {
    private String description;
    private String dueDate;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private Integer number;
    private String owner;
    private String repository;
    private String state;
    private String title;

    private GetRepositoryMilestoneResult() {}
    public String description() {
        return this.description;
    }
    public String dueDate() {
        return this.dueDate;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Integer number() {
        return this.number;
    }
    public String owner() {
        return this.owner;
    }
    public String repository() {
        return this.repository;
    }
    public String state() {
        return this.state;
    }
    public String title() {
        return this.title;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRepositoryMilestoneResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String description;
        private String dueDate;
        private String id;
        private Integer number;
        private String owner;
        private String repository;
        private String state;
        private String title;
        public Builder() {}
        public Builder(GetRepositoryMilestoneResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.dueDate = defaults.dueDate;
    	      this.id = defaults.id;
    	      this.number = defaults.number;
    	      this.owner = defaults.owner;
    	      this.repository = defaults.repository;
    	      this.state = defaults.state;
    	      this.title = defaults.title;
        }

        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder dueDate(String dueDate) {
            this.dueDate = Objects.requireNonNull(dueDate);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder number(Integer number) {
            this.number = Objects.requireNonNull(number);
            return this;
        }
        @CustomType.Setter
        public Builder owner(String owner) {
            this.owner = Objects.requireNonNull(owner);
            return this;
        }
        @CustomType.Setter
        public Builder repository(String repository) {
            this.repository = Objects.requireNonNull(repository);
            return this;
        }
        @CustomType.Setter
        public Builder state(String state) {
            this.state = Objects.requireNonNull(state);
            return this;
        }
        @CustomType.Setter
        public Builder title(String title) {
            this.title = Objects.requireNonNull(title);
            return this;
        }
        public GetRepositoryMilestoneResult build() {
            final var o = new GetRepositoryMilestoneResult();
            o.description = description;
            o.dueDate = dueDate;
            o.id = id;
            o.number = number;
            o.owner = owner;
            o.repository = repository;
            o.state = state;
            o.title = title;
            return o;
        }
    }
}

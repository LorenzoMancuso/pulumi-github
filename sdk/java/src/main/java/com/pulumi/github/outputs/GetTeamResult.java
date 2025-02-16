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
public final class GetTeamResult {
    private String description;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private List<String> members;
    private @Nullable String membershipType;
    private String name;
    private String nodeId;
    private String permission;
    private String privacy;
    private List<String> repositories;
    private @Nullable Integer resultsPerPage;
    private String slug;
    private @Nullable Boolean summaryOnly;

    private GetTeamResult() {}
    public String description() {
        return this.description;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public List<String> members() {
        return this.members;
    }
    public Optional<String> membershipType() {
        return Optional.ofNullable(this.membershipType);
    }
    public String name() {
        return this.name;
    }
    public String nodeId() {
        return this.nodeId;
    }
    public String permission() {
        return this.permission;
    }
    public String privacy() {
        return this.privacy;
    }
    public List<String> repositories() {
        return this.repositories;
    }
    public Optional<Integer> resultsPerPage() {
        return Optional.ofNullable(this.resultsPerPage);
    }
    public String slug() {
        return this.slug;
    }
    public Optional<Boolean> summaryOnly() {
        return Optional.ofNullable(this.summaryOnly);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetTeamResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String description;
        private String id;
        private List<String> members;
        private @Nullable String membershipType;
        private String name;
        private String nodeId;
        private String permission;
        private String privacy;
        private List<String> repositories;
        private @Nullable Integer resultsPerPage;
        private String slug;
        private @Nullable Boolean summaryOnly;
        public Builder() {}
        public Builder(GetTeamResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.id = defaults.id;
    	      this.members = defaults.members;
    	      this.membershipType = defaults.membershipType;
    	      this.name = defaults.name;
    	      this.nodeId = defaults.nodeId;
    	      this.permission = defaults.permission;
    	      this.privacy = defaults.privacy;
    	      this.repositories = defaults.repositories;
    	      this.resultsPerPage = defaults.resultsPerPage;
    	      this.slug = defaults.slug;
    	      this.summaryOnly = defaults.summaryOnly;
        }

        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder members(List<String> members) {
            this.members = Objects.requireNonNull(members);
            return this;
        }
        public Builder members(String... members) {
            return members(List.of(members));
        }
        @CustomType.Setter
        public Builder membershipType(@Nullable String membershipType) {
            this.membershipType = membershipType;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder nodeId(String nodeId) {
            this.nodeId = Objects.requireNonNull(nodeId);
            return this;
        }
        @CustomType.Setter
        public Builder permission(String permission) {
            this.permission = Objects.requireNonNull(permission);
            return this;
        }
        @CustomType.Setter
        public Builder privacy(String privacy) {
            this.privacy = Objects.requireNonNull(privacy);
            return this;
        }
        @CustomType.Setter
        public Builder repositories(List<String> repositories) {
            this.repositories = Objects.requireNonNull(repositories);
            return this;
        }
        public Builder repositories(String... repositories) {
            return repositories(List.of(repositories));
        }
        @CustomType.Setter
        public Builder resultsPerPage(@Nullable Integer resultsPerPage) {
            this.resultsPerPage = resultsPerPage;
            return this;
        }
        @CustomType.Setter
        public Builder slug(String slug) {
            this.slug = Objects.requireNonNull(slug);
            return this;
        }
        @CustomType.Setter
        public Builder summaryOnly(@Nullable Boolean summaryOnly) {
            this.summaryOnly = summaryOnly;
            return this;
        }
        public GetTeamResult build() {
            final var o = new GetTeamResult();
            o.description = description;
            o.id = id;
            o.members = members;
            o.membershipType = membershipType;
            o.name = name;
            o.nodeId = nodeId;
            o.permission = permission;
            o.privacy = privacy;
            o.repositories = repositories;
            o.resultsPerPage = resultsPerPage;
            o.slug = slug;
            o.summaryOnly = summaryOnly;
            return o;
        }
    }
}

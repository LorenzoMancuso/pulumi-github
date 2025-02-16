// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetOrganizationTeamSyncGroupsGroup {
    private String groupDescription;
    private String groupId;
    private String groupName;

    private GetOrganizationTeamSyncGroupsGroup() {}
    public String groupDescription() {
        return this.groupDescription;
    }
    public String groupId() {
        return this.groupId;
    }
    public String groupName() {
        return this.groupName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetOrganizationTeamSyncGroupsGroup defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String groupDescription;
        private String groupId;
        private String groupName;
        public Builder() {}
        public Builder(GetOrganizationTeamSyncGroupsGroup defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.groupDescription = defaults.groupDescription;
    	      this.groupId = defaults.groupId;
    	      this.groupName = defaults.groupName;
        }

        @CustomType.Setter
        public Builder groupDescription(String groupDescription) {
            this.groupDescription = Objects.requireNonNull(groupDescription);
            return this;
        }
        @CustomType.Setter
        public Builder groupId(String groupId) {
            this.groupId = Objects.requireNonNull(groupId);
            return this;
        }
        @CustomType.Setter
        public Builder groupName(String groupName) {
            this.groupName = Objects.requireNonNull(groupName);
            return this;
        }
        public GetOrganizationTeamSyncGroupsGroup build() {
            final var o = new GetOrganizationTeamSyncGroupsGroup();
            o.groupDescription = groupDescription;
            o.groupId = groupId;
            o.groupName = groupName;
            return o;
        }
    }
}

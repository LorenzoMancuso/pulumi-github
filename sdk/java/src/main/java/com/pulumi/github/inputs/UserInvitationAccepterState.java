// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class UserInvitationAccepterState extends com.pulumi.resources.ResourceArgs {

    public static final UserInvitationAccepterState Empty = new UserInvitationAccepterState();

    @Import(name="invitationId")
    private @Nullable Output<String> invitationId;

    public Optional<Output<String>> invitationId() {
        return Optional.ofNullable(this.invitationId);
    }

    private UserInvitationAccepterState() {}

    private UserInvitationAccepterState(UserInvitationAccepterState $) {
        this.invitationId = $.invitationId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(UserInvitationAccepterState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private UserInvitationAccepterState $;

        public Builder() {
            $ = new UserInvitationAccepterState();
        }

        public Builder(UserInvitationAccepterState defaults) {
            $ = new UserInvitationAccepterState(Objects.requireNonNull(defaults));
        }

        public Builder invitationId(@Nullable Output<String> invitationId) {
            $.invitationId = invitationId;
            return this;
        }

        public Builder invitationId(String invitationId) {
            return invitationId(Output.of(invitationId));
        }

        public UserInvitationAccepterState build() {
            return $;
        }
    }

}

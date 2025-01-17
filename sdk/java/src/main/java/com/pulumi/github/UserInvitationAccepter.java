// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.github.UserInvitationAccepterArgs;
import com.pulumi.github.Utilities;
import com.pulumi.github.inputs.UserInvitationAccepterState;
import java.lang.String;
import javax.annotation.Nullable;

@ResourceType(type="github:index/userInvitationAccepter:UserInvitationAccepter")
public class UserInvitationAccepter extends com.pulumi.resources.CustomResource {
    @Export(name="invitationId", type=String.class, parameters={})
    private Output<String> invitationId;

    public Output<String> invitationId() {
        return this.invitationId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public UserInvitationAccepter(String name) {
        this(name, UserInvitationAccepterArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public UserInvitationAccepter(String name, UserInvitationAccepterArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public UserInvitationAccepter(String name, UserInvitationAccepterArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/userInvitationAccepter:UserInvitationAccepter", name, args == null ? UserInvitationAccepterArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private UserInvitationAccepter(String name, Output<String> id, @Nullable UserInvitationAccepterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/userInvitationAccepter:UserInvitationAccepter", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static UserInvitationAccepter get(String name, Output<String> id, @Nullable UserInvitationAccepterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new UserInvitationAccepter(name, id, state, options);
    }
}

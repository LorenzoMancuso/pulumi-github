// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.github.RepositoryCollaboratorArgs;
import com.pulumi.github.Utilities;
import com.pulumi.github.inputs.RepositoryCollaboratorState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="github:index/repositoryCollaborator:RepositoryCollaborator")
public class RepositoryCollaborator extends com.pulumi.resources.CustomResource {
    @Export(name="invitationId", type=String.class, parameters={})
    private Output<String> invitationId;

    public Output<String> invitationId() {
        return this.invitationId;
    }
    @Export(name="permission", type=String.class, parameters={})
    private Output</* @Nullable */ String> permission;

    public Output<Optional<String>> permission() {
        return Codegen.optional(this.permission);
    }
    @Export(name="permissionDiffSuppression", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> permissionDiffSuppression;

    public Output<Optional<Boolean>> permissionDiffSuppression() {
        return Codegen.optional(this.permissionDiffSuppression);
    }
    @Export(name="repository", type=String.class, parameters={})
    private Output<String> repository;

    public Output<String> repository() {
        return this.repository;
    }
    @Export(name="username", type=String.class, parameters={})
    private Output<String> username;

    public Output<String> username() {
        return this.username;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RepositoryCollaborator(String name) {
        this(name, RepositoryCollaboratorArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RepositoryCollaborator(String name, RepositoryCollaboratorArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RepositoryCollaborator(String name, RepositoryCollaboratorArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/repositoryCollaborator:RepositoryCollaborator", name, args == null ? RepositoryCollaboratorArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RepositoryCollaborator(String name, Output<String> id, @Nullable RepositoryCollaboratorState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/repositoryCollaborator:RepositoryCollaborator", name, state, makeResourceOptions(options, id));
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
    public static RepositoryCollaborator get(String name, Output<String> id, @Nullable RepositoryCollaboratorState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RepositoryCollaborator(name, id, state, options);
    }
}

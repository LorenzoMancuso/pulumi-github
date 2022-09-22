// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.github.TeamRepositoryArgs;
import com.pulumi.github.Utilities;
import com.pulumi.github.inputs.TeamRepositoryState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="github:index/teamRepository:TeamRepository")
public class TeamRepository extends com.pulumi.resources.CustomResource {
    @Export(name="etag", type=String.class, parameters={})
    private Output<String> etag;

    public Output<String> etag() {
        return this.etag;
    }
    @Export(name="permission", type=String.class, parameters={})
    private Output</* @Nullable */ String> permission;

    public Output<Optional<String>> permission() {
        return Codegen.optional(this.permission);
    }
    @Export(name="repository", type=String.class, parameters={})
    private Output<String> repository;

    public Output<String> repository() {
        return this.repository;
    }
    /**
     * ID or slug of team
     * 
     */
    @Export(name="teamId", type=String.class, parameters={})
    private Output<String> teamId;

    /**
     * @return ID or slug of team
     * 
     */
    public Output<String> teamId() {
        return this.teamId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public TeamRepository(String name) {
        this(name, TeamRepositoryArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public TeamRepository(String name, TeamRepositoryArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public TeamRepository(String name, TeamRepositoryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/teamRepository:TeamRepository", name, args == null ? TeamRepositoryArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private TeamRepository(String name, Output<String> id, @Nullable TeamRepositoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/teamRepository:TeamRepository", name, state, makeResourceOptions(options, id));
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
    public static TeamRepository get(String name, Output<String> id, @Nullable TeamRepositoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new TeamRepository(name, id, state, options);
    }
}

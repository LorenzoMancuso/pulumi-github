// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.github.RepositoryAutolinkReferenceArgs;
import com.pulumi.github.Utilities;
import com.pulumi.github.inputs.RepositoryAutolinkReferenceState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="github:index/repositoryAutolinkReference:RepositoryAutolinkReference")
public class RepositoryAutolinkReference extends com.pulumi.resources.CustomResource {
    @Export(name="etag", type=String.class, parameters={})
    private Output<String> etag;

    public Output<String> etag() {
        return this.etag;
    }
    /**
     * Whether this autolink reference matches alphanumeric characters. If false, this autolink reference only matches numeric
     * characters.
     * 
     */
    @Export(name="isAlphanumeric", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> isAlphanumeric;

    /**
     * @return Whether this autolink reference matches alphanumeric characters. If false, this autolink reference only matches numeric
     * characters.
     * 
     */
    public Output<Optional<Boolean>> isAlphanumeric() {
        return Codegen.optional(this.isAlphanumeric);
    }
    /**
     * This prefix appended by a number will generate a link any time it is found in an issue, pull request, or commit
     * 
     */
    @Export(name="keyPrefix", type=String.class, parameters={})
    private Output<String> keyPrefix;

    /**
     * @return This prefix appended by a number will generate a link any time it is found in an issue, pull request, or commit
     * 
     */
    public Output<String> keyPrefix() {
        return this.keyPrefix;
    }
    /**
     * The repository name
     * 
     */
    @Export(name="repository", type=String.class, parameters={})
    private Output<String> repository;

    /**
     * @return The repository name
     * 
     */
    public Output<String> repository() {
        return this.repository;
    }
    /**
     * The template of the target URL used for the links; must be a valid URL and contain `&lt;num&gt;` for the reference number
     * 
     */
    @Export(name="targetUrlTemplate", type=String.class, parameters={})
    private Output<String> targetUrlTemplate;

    /**
     * @return The template of the target URL used for the links; must be a valid URL and contain `&lt;num&gt;` for the reference number
     * 
     */
    public Output<String> targetUrlTemplate() {
        return this.targetUrlTemplate;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RepositoryAutolinkReference(String name) {
        this(name, RepositoryAutolinkReferenceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RepositoryAutolinkReference(String name, RepositoryAutolinkReferenceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RepositoryAutolinkReference(String name, RepositoryAutolinkReferenceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/repositoryAutolinkReference:RepositoryAutolinkReference", name, args == null ? RepositoryAutolinkReferenceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RepositoryAutolinkReference(String name, Output<String> id, @Nullable RepositoryAutolinkReferenceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/repositoryAutolinkReference:RepositoryAutolinkReference", name, state, makeResourceOptions(options, id));
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
    public static RepositoryAutolinkReference get(String name, Output<String> id, @Nullable RepositoryAutolinkReferenceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RepositoryAutolinkReference(name, id, state, options);
    }
}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.github.ReleaseArgs;
import com.pulumi.github.Utilities;
import com.pulumi.github.inputs.ReleaseState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="github:index/release:Release")
public class Release extends com.pulumi.resources.CustomResource {
    @Export(name="body", type=String.class, parameters={})
    private Output</* @Nullable */ String> body;

    public Output<Optional<String>> body() {
        return Codegen.optional(this.body);
    }
    @Export(name="discussionCategoryName", type=String.class, parameters={})
    private Output</* @Nullable */ String> discussionCategoryName;

    public Output<Optional<String>> discussionCategoryName() {
        return Codegen.optional(this.discussionCategoryName);
    }
    @Export(name="draft", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> draft;

    public Output<Optional<Boolean>> draft() {
        return Codegen.optional(this.draft);
    }
    @Export(name="etag", type=String.class, parameters={})
    private Output<String> etag;

    public Output<String> etag() {
        return this.etag;
    }
    @Export(name="generateReleaseNotes", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> generateReleaseNotes;

    public Output<Optional<Boolean>> generateReleaseNotes() {
        return Codegen.optional(this.generateReleaseNotes);
    }
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    public Output<String> name() {
        return this.name;
    }
    @Export(name="prerelease", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> prerelease;

    public Output<Optional<Boolean>> prerelease() {
        return Codegen.optional(this.prerelease);
    }
    @Export(name="repository", type=String.class, parameters={})
    private Output<String> repository;

    public Output<String> repository() {
        return this.repository;
    }
    @Export(name="tagName", type=String.class, parameters={})
    private Output<String> tagName;

    public Output<String> tagName() {
        return this.tagName;
    }
    @Export(name="targetCommitish", type=String.class, parameters={})
    private Output</* @Nullable */ String> targetCommitish;

    public Output<Optional<String>> targetCommitish() {
        return Codegen.optional(this.targetCommitish);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Release(String name) {
        this(name, ReleaseArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Release(String name, ReleaseArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Release(String name, ReleaseArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/release:Release", name, args == null ? ReleaseArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Release(String name, Output<String> id, @Nullable ReleaseState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/release:Release", name, state, makeResourceOptions(options, id));
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
    public static Release get(String name, Output<String> id, @Nullable ReleaseState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Release(name, id, state, options);
    }
}
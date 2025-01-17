// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.github.ProviderArgs;
import com.pulumi.github.Utilities;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The provider type for the github package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 * 
 */
@ResourceType(type="pulumi:providers:github")
public class Provider extends com.pulumi.resources.ProviderResource {
    /**
     * The GitHub Base API URL
     * 
     */
    @Export(name="baseUrl", type=String.class, parameters={})
    private Output</* @Nullable */ String> baseUrl;

    /**
     * @return The GitHub Base API URL
     * 
     */
    public Output<Optional<String>> baseUrl() {
        return Codegen.optional(this.baseUrl);
    }
    /**
     * The GitHub organization name to manage. Use this field instead of `owner` when managing organization accounts.
     * 
     * @deprecated
     * Use owner (or GITHUB_OWNER) instead of organization (or GITHUB_ORGANIZATION)
     * 
     */
    @Deprecated /* Use owner (or GITHUB_OWNER) instead of organization (or GITHUB_ORGANIZATION) */
    @Export(name="organization", type=String.class, parameters={})
    private Output</* @Nullable */ String> organization;

    /**
     * @return The GitHub organization name to manage. Use this field instead of `owner` when managing organization accounts.
     * 
     */
    public Output<Optional<String>> organization() {
        return Codegen.optional(this.organization);
    }
    /**
     * The GitHub owner name to manage. Use this field instead of `organization` when managing individual accounts.
     * 
     */
    @Export(name="owner", type=String.class, parameters={})
    private Output</* @Nullable */ String> owner;

    /**
     * @return The GitHub owner name to manage. Use this field instead of `organization` when managing individual accounts.
     * 
     */
    public Output<Optional<String>> owner() {
        return Codegen.optional(this.owner);
    }
    /**
     * The OAuth token used to connect to GitHub. Anonymous mode is enabled if both `token` and `app_auth` are not set.
     * 
     */
    @Export(name="token", type=String.class, parameters={})
    private Output</* @Nullable */ String> token;

    /**
     * @return The OAuth token used to connect to GitHub. Anonymous mode is enabled if both `token` and `app_auth` are not set.
     * 
     */
    public Output<Optional<String>> token() {
        return Codegen.optional(this.token);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Provider(String name) {
        this(name, ProviderArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Provider(String name, @Nullable ProviderArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Provider(String name, @Nullable ProviderArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github", name, args == null ? ProviderArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

}

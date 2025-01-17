// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.github.UserGpgKeyArgs;
import com.pulumi.github.Utilities;
import com.pulumi.github.inputs.UserGpgKeyState;
import java.lang.String;
import javax.annotation.Nullable;

@ResourceType(type="github:index/userGpgKey:UserGpgKey")
public class UserGpgKey extends com.pulumi.resources.CustomResource {
    @Export(name="armoredPublicKey", type=String.class, parameters={})
    private Output<String> armoredPublicKey;

    public Output<String> armoredPublicKey() {
        return this.armoredPublicKey;
    }
    @Export(name="etag", type=String.class, parameters={})
    private Output<String> etag;

    public Output<String> etag() {
        return this.etag;
    }
    @Export(name="keyId", type=String.class, parameters={})
    private Output<String> keyId;

    public Output<String> keyId() {
        return this.keyId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public UserGpgKey(String name) {
        this(name, UserGpgKeyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public UserGpgKey(String name, UserGpgKeyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public UserGpgKey(String name, UserGpgKeyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/userGpgKey:UserGpgKey", name, args == null ? UserGpgKeyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private UserGpgKey(String name, Output<String> id, @Nullable UserGpgKeyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/userGpgKey:UserGpgKey", name, state, makeResourceOptions(options, id));
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
    public static UserGpgKey get(String name, Output<String> id, @Nullable UserGpgKeyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new UserGpgKey(name, id, state, options);
    }
}

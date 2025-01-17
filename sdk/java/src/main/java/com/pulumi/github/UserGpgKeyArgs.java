// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class UserGpgKeyArgs extends com.pulumi.resources.ResourceArgs {

    public static final UserGpgKeyArgs Empty = new UserGpgKeyArgs();

    @Import(name="armoredPublicKey", required=true)
    private Output<String> armoredPublicKey;

    public Output<String> armoredPublicKey() {
        return this.armoredPublicKey;
    }

    private UserGpgKeyArgs() {}

    private UserGpgKeyArgs(UserGpgKeyArgs $) {
        this.armoredPublicKey = $.armoredPublicKey;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(UserGpgKeyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private UserGpgKeyArgs $;

        public Builder() {
            $ = new UserGpgKeyArgs();
        }

        public Builder(UserGpgKeyArgs defaults) {
            $ = new UserGpgKeyArgs(Objects.requireNonNull(defaults));
        }

        public Builder armoredPublicKey(Output<String> armoredPublicKey) {
            $.armoredPublicKey = armoredPublicKey;
            return this;
        }

        public Builder armoredPublicKey(String armoredPublicKey) {
            return armoredPublicKey(Output.of(armoredPublicKey));
        }

        public UserGpgKeyArgs build() {
            $.armoredPublicKey = Objects.requireNonNull($.armoredPublicKey, "expected parameter 'armoredPublicKey' to be non-null");
            return $;
        }
    }

}

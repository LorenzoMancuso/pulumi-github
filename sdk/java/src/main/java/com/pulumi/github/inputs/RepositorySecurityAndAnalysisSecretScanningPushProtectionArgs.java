// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class RepositorySecurityAndAnalysisSecretScanningPushProtectionArgs extends com.pulumi.resources.ResourceArgs {

    public static final RepositorySecurityAndAnalysisSecretScanningPushProtectionArgs Empty = new RepositorySecurityAndAnalysisSecretScanningPushProtectionArgs();

    @Import(name="status", required=true)
    private Output<String> status;

    public Output<String> status() {
        return this.status;
    }

    private RepositorySecurityAndAnalysisSecretScanningPushProtectionArgs() {}

    private RepositorySecurityAndAnalysisSecretScanningPushProtectionArgs(RepositorySecurityAndAnalysisSecretScanningPushProtectionArgs $) {
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RepositorySecurityAndAnalysisSecretScanningPushProtectionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RepositorySecurityAndAnalysisSecretScanningPushProtectionArgs $;

        public Builder() {
            $ = new RepositorySecurityAndAnalysisSecretScanningPushProtectionArgs();
        }

        public Builder(RepositorySecurityAndAnalysisSecretScanningPushProtectionArgs defaults) {
            $ = new RepositorySecurityAndAnalysisSecretScanningPushProtectionArgs(Objects.requireNonNull(defaults));
        }

        public Builder status(Output<String> status) {
            $.status = status;
            return this;
        }

        public Builder status(String status) {
            return status(Output.of(status));
        }

        public RepositorySecurityAndAnalysisSecretScanningPushProtectionArgs build() {
            $.status = Objects.requireNonNull($.status, "expected parameter 'status' to be non-null");
            return $;
        }
    }

}
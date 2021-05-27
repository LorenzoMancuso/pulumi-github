// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github.Outputs
{

    [OutputType]
    public sealed class ProviderAppAuth
    {
        public readonly string Id;
        public readonly string InstallationId;
        public readonly string PemFile;

        [OutputConstructor]
        private ProviderAppAuth(
            string id,

            string installationId,

            string pemFile)
        {
            Id = id;
            InstallationId = installationId;
            PemFile = pemFile;
        }
    }
}

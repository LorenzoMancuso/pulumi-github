// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github.Inputs
{

    public sealed class ProviderAppAuthArgs : Pulumi.ResourceArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        [Input("installationId", required: true)]
        public Input<string> InstallationId { get; set; } = null!;

        [Input("pemFile", required: true)]
        public Input<string> PemFile { get; set; } = null!;

        public ProviderAppAuthArgs()
        {
        }
    }
}
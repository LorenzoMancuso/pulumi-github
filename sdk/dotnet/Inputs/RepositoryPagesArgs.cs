// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github.Inputs
{

    public sealed class RepositoryPagesArgs : global::Pulumi.ResourceArgs
    {
        [Input("cname")]
        public Input<string>? Cname { get; set; }

        [Input("custom404")]
        public Input<bool>? Custom404 { get; set; }

        [Input("htmlUrl")]
        public Input<string>? HtmlUrl { get; set; }

        [Input("source", required: true)]
        public Input<Inputs.RepositoryPagesSourceArgs> Source { get; set; } = null!;

        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("url")]
        public Input<string>? Url { get; set; }

        public RepositoryPagesArgs()
        {
        }
        public static new RepositoryPagesArgs Empty => new RepositoryPagesArgs();
    }
}

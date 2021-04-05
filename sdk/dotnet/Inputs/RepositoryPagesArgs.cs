// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github.Inputs
{

    public sealed class RepositoryPagesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The custom domain for the repository. This can only be set after the repository has been created.
        /// </summary>
        [Input("cname")]
        public Input<string>? Cname { get; set; }

        /// <summary>
        /// Whether the rendered GitHub Pages site has a custom 404 page.
        /// </summary>
        [Input("custom404")]
        public Input<bool>? Custom404 { get; set; }

        /// <summary>
        /// The absolute URL (including scheme) of the rendered GitHub Pages site e.g. `https://username.github.io`.
        /// </summary>
        [Input("htmlUrl")]
        public Input<string>? HtmlUrl { get; set; }

        /// <summary>
        /// The source branch and directory for the rendered Pages site. See GitHub Pages Source below for details.
        /// </summary>
        [Input("source", required: true)]
        public Input<Inputs.RepositoryPagesSourceArgs> Source { get; set; } = null!;

        /// <summary>
        /// The GitHub Pages site's build status e.g. `building` or `built`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("url")]
        public Input<string>? Url { get; set; }

        public RepositoryPagesArgs()
        {
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    public static class GetIpRanges
    {
        public static Task<GetIpRangesResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIpRangesResult>("github:index/getIpRanges:getIpRanges", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetIpRangesResult
    {
        public readonly ImmutableArray<string> Actions;
        public readonly ImmutableArray<string> ActionsIpv4s;
        public readonly ImmutableArray<string> ActionsIpv6s;
        public readonly ImmutableArray<string> ApiIpv4s;
        public readonly ImmutableArray<string> ApiIpv6s;
        public readonly ImmutableArray<string> Apis;
        public readonly ImmutableArray<string> DependabotIpv4s;
        public readonly ImmutableArray<string> DependabotIpv6s;
        public readonly ImmutableArray<string> Dependabots;
        public readonly ImmutableArray<string> GitIpv4s;
        public readonly ImmutableArray<string> GitIpv6s;
        public readonly ImmutableArray<string> Gits;
        public readonly ImmutableArray<string> Hooks;
        public readonly ImmutableArray<string> HooksIpv4s;
        public readonly ImmutableArray<string> HooksIpv6s;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> ImporterIpv4s;
        public readonly ImmutableArray<string> ImporterIpv6s;
        public readonly ImmutableArray<string> Importers;
        public readonly ImmutableArray<string> Pages;
        public readonly ImmutableArray<string> PagesIpv4s;
        public readonly ImmutableArray<string> PagesIpv6s;
        public readonly ImmutableArray<string> WebIpv4s;
        public readonly ImmutableArray<string> WebIpv6s;
        public readonly ImmutableArray<string> Webs;

        [OutputConstructor]
        private GetIpRangesResult(
            ImmutableArray<string> actions,

            ImmutableArray<string> actionsIpv4s,

            ImmutableArray<string> actionsIpv6s,

            ImmutableArray<string> apiIpv4s,

            ImmutableArray<string> apiIpv6s,

            ImmutableArray<string> apis,

            ImmutableArray<string> dependabotIpv4s,

            ImmutableArray<string> dependabotIpv6s,

            ImmutableArray<string> dependabots,

            ImmutableArray<string> gitIpv4s,

            ImmutableArray<string> gitIpv6s,

            ImmutableArray<string> gits,

            ImmutableArray<string> hooks,

            ImmutableArray<string> hooksIpv4s,

            ImmutableArray<string> hooksIpv6s,

            string id,

            ImmutableArray<string> importerIpv4s,

            ImmutableArray<string> importerIpv6s,

            ImmutableArray<string> importers,

            ImmutableArray<string> pages,

            ImmutableArray<string> pagesIpv4s,

            ImmutableArray<string> pagesIpv6s,

            ImmutableArray<string> webIpv4s,

            ImmutableArray<string> webIpv6s,

            ImmutableArray<string> webs)
        {
            Actions = actions;
            ActionsIpv4s = actionsIpv4s;
            ActionsIpv6s = actionsIpv6s;
            ApiIpv4s = apiIpv4s;
            ApiIpv6s = apiIpv6s;
            Apis = apis;
            DependabotIpv4s = dependabotIpv4s;
            DependabotIpv6s = dependabotIpv6s;
            Dependabots = dependabots;
            GitIpv4s = gitIpv4s;
            GitIpv6s = gitIpv6s;
            Gits = gits;
            Hooks = hooks;
            HooksIpv4s = hooksIpv4s;
            HooksIpv6s = hooksIpv6s;
            Id = id;
            ImporterIpv4s = importerIpv4s;
            ImporterIpv6s = importerIpv6s;
            Importers = importers;
            Pages = pages;
            PagesIpv4s = pagesIpv4s;
            PagesIpv6s = pagesIpv6s;
            WebIpv4s = webIpv4s;
            WebIpv6s = webIpv6s;
            Webs = webs;
        }
    }
}

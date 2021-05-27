// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;

namespace Pulumi.Github
{
    public static class Config
    {
        private static readonly Pulumi.Config __config = new Pulumi.Config("github");
        /// <summary>
        /// The GitHub App credentials used to connect to GitHub. Conflicts with `token`. Anonymous mode is enabled if both `token`
        /// and `app_auth` are not set.
        /// </summary>
        public static Pulumi.Github.Config.Types.AppAuth? AppAuth { get; set; } = __config.GetObject<Pulumi.Github.Config.Types.AppAuth>("appAuth");

        /// <summary>
        /// The GitHub Base API URL
        /// </summary>
        public static string? BaseUrl { get; set; } = __config.Get("baseUrl") ?? Utilities.GetEnv("GITHUB_BASE_URL") ?? "https://api.github.com/";

        /// <summary>
        /// Enable `insecure` mode for testing purposes
        /// </summary>
        public static bool? Insecure { get; set; } = __config.GetBoolean("insecure");

        /// <summary>
        /// The GitHub organization name to manage. Use this field instead of `owner` when managing organization accounts.
        /// </summary>
        public static string? Organization { get; set; } = __config.Get("organization");

        /// <summary>
        /// The GitHub owner name to manage. Use this field instead of `organization` when managing individual accounts.
        /// </summary>
        public static string? Owner { get; set; } = __config.Get("owner");

        /// <summary>
        /// The OAuth token used to connect to GitHub. Anonymous mode is enabled if both `token` and `app_auth` are not set.
        /// </summary>
        public static string? Token { get; set; } = __config.Get("token");

        public static class Types
        {

             public class AppAuth
             {
                public string Id { get; set; }
                public string InstallationId { get; set; }
                public string PemFile { get; set; }
            }
        }
    }
}

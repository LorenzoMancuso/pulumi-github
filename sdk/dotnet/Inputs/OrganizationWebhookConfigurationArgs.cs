// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github.Inputs
{

    public sealed class OrganizationWebhookConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("contentType")]
        public Input<string>? ContentType { get; set; }

        [Input("insecureSsl")]
        public Input<bool>? InsecureSsl { get; set; }

        [Input("secret")]
        private Input<string>? _secret;
        public Input<string>? Secret
        {
            get => _secret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("url", required: true)]
        private Input<string>? _url;
        public Input<string>? Url
        {
            get => _url;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _url = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public OrganizationWebhookConfigurationArgs()
        {
        }
        public static new OrganizationWebhookConfigurationArgs Empty => new OrganizationWebhookConfigurationArgs();
    }
}

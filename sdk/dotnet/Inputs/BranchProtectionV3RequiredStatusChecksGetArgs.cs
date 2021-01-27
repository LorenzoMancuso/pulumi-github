// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github.Inputs
{

    public sealed class BranchProtectionV3RequiredStatusChecksGetArgs : Pulumi.ResourceArgs
    {
        [Input("contexts")]
        private InputList<string>? _contexts;
        public InputList<string> Contexts
        {
            get => _contexts ?? (_contexts = new InputList<string>());
            set => _contexts = value;
        }

        [Input("includeAdmins")]
        public Input<bool>? IncludeAdmins { get; set; }

        [Input("strict")]
        public Input<bool>? Strict { get; set; }

        public BranchProtectionV3RequiredStatusChecksGetArgs()
        {
        }
    }
}
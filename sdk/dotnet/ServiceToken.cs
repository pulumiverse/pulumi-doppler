// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Doppler
{
    /// <summary>
    /// Manage a Doppler service token.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Doppler = Pulumiverse.Doppler;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var backendCiToken = new Doppler.ServiceToken("backend_ci_token", new()
    ///     {
    ///         Project = "backend",
    ///         Config = "ci",
    ///         Name = "Builder Token",
    ///         Access = "read",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [DopplerResourceType("doppler:index/serviceToken:ServiceToken")]
    public partial class ServiceToken : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The access level (read or read/write)
        /// </summary>
        [Output("access")]
        public Output<string?> Access { get; private set; } = null!;

        /// <summary>
        /// The name of the Doppler config where the service token is located
        /// </summary>
        [Output("config")]
        public Output<string> Config { get; private set; } = null!;

        /// <summary>
        /// The key for the Doppler service token
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        /// <summary>
        /// The name of the Doppler service token
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the Doppler project where the service token is located
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceToken resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceToken(string name, ServiceTokenArgs args, CustomResourceOptions? options = null)
            : base("doppler:index/serviceToken:ServiceToken", name, args ?? new ServiceTokenArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceToken(string name, Input<string> id, ServiceTokenState? state = null, CustomResourceOptions? options = null)
            : base("doppler:index/serviceToken:ServiceToken", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
                AdditionalSecretOutputs =
                {
                    "key",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ServiceToken resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceToken Get(string name, Input<string> id, ServiceTokenState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceToken(name, id, state, options);
        }
    }

    public sealed class ServiceTokenArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The access level (read or read/write)
        /// </summary>
        [Input("access")]
        public Input<string>? Access { get; set; }

        /// <summary>
        /// The name of the Doppler config where the service token is located
        /// </summary>
        [Input("config", required: true)]
        public Input<string> Config { get; set; } = null!;

        /// <summary>
        /// The name of the Doppler service token
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the Doppler project where the service token is located
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        public ServiceTokenArgs()
        {
        }
        public static new ServiceTokenArgs Empty => new ServiceTokenArgs();
    }

    public sealed class ServiceTokenState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The access level (read or read/write)
        /// </summary>
        [Input("access")]
        public Input<string>? Access { get; set; }

        /// <summary>
        /// The name of the Doppler config where the service token is located
        /// </summary>
        [Input("config")]
        public Input<string>? Config { get; set; }

        [Input("key")]
        private Input<string>? _key;

        /// <summary>
        /// The key for the Doppler service token
        /// </summary>
        public Input<string>? Key
        {
            get => _key;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _key = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The name of the Doppler service token
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the Doppler project where the service token is located
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public ServiceTokenState()
        {
        }
        public static new ServiceTokenState Empty => new ServiceTokenState();
    }
}

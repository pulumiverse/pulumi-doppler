// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Doppler.SecretsSync
{
    /// <summary>
    /// ## Example Usage
    /// </summary>
    [DopplerResourceType("doppler:secretsSync/terraformCloud:TerraformCloud")]
    public partial class TerraformCloud : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the Doppler config
        /// </summary>
        [Output("config")]
        public Output<string> Config { get; private set; } = null!;

        /// <summary>
        /// The behavior to be performed on the secrets in the sync target when this resource is deleted or recreated. Either `leave_in_target` (default) or `delete_from_target`.
        /// </summary>
        [Output("deleteBehavior")]
        public Output<string?> DeleteBehavior { get; private set; } = null!;

        /// <summary>
        /// The slug of the integration to use for this sync
        /// </summary>
        [Output("integration")]
        public Output<string> Integration { get; private set; } = null!;

        /// <summary>
        /// A name transform to apply before syncing secrets: "none" or "lowercase"
        /// </summary>
        [Output("nameTransform")]
        public Output<string> NameTransform { get; private set; } = null!;

        /// <summary>
        /// The name of the Doppler project
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Either "workspace" or "variableSet", based on the resource type to sync to
        /// </summary>
        [Output("syncTarget")]
        public Output<string> SyncTarget { get; private set; } = null!;

        [Output("variableSetId")]
        public Output<string?> VariableSetId { get; private set; } = null!;

        [Output("variableSyncType")]
        public Output<string> VariableSyncType { get; private set; } = null!;

        [Output("workspaceId")]
        public Output<string?> WorkspaceId { get; private set; } = null!;


        /// <summary>
        /// Create a TerraformCloud resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TerraformCloud(string name, TerraformCloudArgs args, CustomResourceOptions? options = null)
            : base("doppler:secretsSync/terraformCloud:TerraformCloud", name, args ?? new TerraformCloudArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TerraformCloud(string name, Input<string> id, TerraformCloudState? state = null, CustomResourceOptions? options = null)
            : base("doppler:secretsSync/terraformCloud:TerraformCloud", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing TerraformCloud resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TerraformCloud Get(string name, Input<string> id, TerraformCloudState? state = null, CustomResourceOptions? options = null)
        {
            return new TerraformCloud(name, id, state, options);
        }
    }

    public sealed class TerraformCloudArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Doppler config
        /// </summary>
        [Input("config", required: true)]
        public Input<string> Config { get; set; } = null!;

        /// <summary>
        /// The behavior to be performed on the secrets in the sync target when this resource is deleted or recreated. Either `leave_in_target` (default) or `delete_from_target`.
        /// </summary>
        [Input("deleteBehavior")]
        public Input<string>? DeleteBehavior { get; set; }

        /// <summary>
        /// The slug of the integration to use for this sync
        /// </summary>
        [Input("integration", required: true)]
        public Input<string> Integration { get; set; } = null!;

        /// <summary>
        /// A name transform to apply before syncing secrets: "none" or "lowercase"
        /// </summary>
        [Input("nameTransform", required: true)]
        public Input<string> NameTransform { get; set; } = null!;

        /// <summary>
        /// The name of the Doppler project
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// Either "workspace" or "variableSet", based on the resource type to sync to
        /// </summary>
        [Input("syncTarget", required: true)]
        public Input<string> SyncTarget { get; set; } = null!;

        [Input("variableSetId")]
        public Input<string>? VariableSetId { get; set; }

        [Input("variableSyncType", required: true)]
        public Input<string> VariableSyncType { get; set; } = null!;

        [Input("workspaceId")]
        public Input<string>? WorkspaceId { get; set; }

        public TerraformCloudArgs()
        {
        }
        public static new TerraformCloudArgs Empty => new TerraformCloudArgs();
    }

    public sealed class TerraformCloudState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Doppler config
        /// </summary>
        [Input("config")]
        public Input<string>? Config { get; set; }

        /// <summary>
        /// The behavior to be performed on the secrets in the sync target when this resource is deleted or recreated. Either `leave_in_target` (default) or `delete_from_target`.
        /// </summary>
        [Input("deleteBehavior")]
        public Input<string>? DeleteBehavior { get; set; }

        /// <summary>
        /// The slug of the integration to use for this sync
        /// </summary>
        [Input("integration")]
        public Input<string>? Integration { get; set; }

        /// <summary>
        /// A name transform to apply before syncing secrets: "none" or "lowercase"
        /// </summary>
        [Input("nameTransform")]
        public Input<string>? NameTransform { get; set; }

        /// <summary>
        /// The name of the Doppler project
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Either "workspace" or "variableSet", based on the resource type to sync to
        /// </summary>
        [Input("syncTarget")]
        public Input<string>? SyncTarget { get; set; }

        [Input("variableSetId")]
        public Input<string>? VariableSetId { get; set; }

        [Input("variableSyncType")]
        public Input<string>? VariableSyncType { get; set; }

        [Input("workspaceId")]
        public Input<string>? WorkspaceId { get; set; }

        public TerraformCloudState()
        {
        }
        public static new TerraformCloudState Empty => new TerraformCloudState();
    }
}

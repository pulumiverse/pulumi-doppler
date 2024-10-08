// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manage a CircleCI Doppler sync.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as doppler from "@pulumiverse/doppler";
 *
 * const prod = new doppler.integration.Circleci("prod", {
 *     name: "Production",
 *     apiToken: "my_api_token",
 * });
 * const backendProd = new doppler.secretssync.Circleci("backend_prod", {
 *     integration: prod.id,
 *     project: "backend",
 *     config: "prd",
 *     resourceType: "project",
 *     resourceId: "github/myorg/myproject",
 *     organizationSlug: "myorg",
 *     deleteBehavior: "leave_in_target",
 * });
 * ```
 */
export class Circleci extends pulumi.CustomResource {
    /**
     * Get an existing Circleci resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CircleciState, opts?: pulumi.CustomResourceOptions): Circleci {
        return new Circleci(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'doppler:secretsSync/circleci:Circleci';

    /**
     * Returns true if the given object is an instance of Circleci.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Circleci {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Circleci.__pulumiType;
    }

    /**
     * The name of the Doppler config
     */
    public readonly config!: pulumi.Output<string>;
    /**
     * The behavior to be performed on the secrets in the sync target when this resource is deleted or recreated. Either `leaveInTarget` (default) or `deleteFromTarget`.
     */
    public readonly deleteBehavior!: pulumi.Output<string | undefined>;
    /**
     * The slug of the integration to use for this sync
     */
    public readonly integration!: pulumi.Output<string>;
    /**
     * The organization slug where the resource is located
     */
    public readonly organizationSlug!: pulumi.Output<string>;
    /**
     * The name of the Doppler project
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The resource ID (either project or context) to sync to
     */
    public readonly resourceId!: pulumi.Output<string>;
    /**
     * Either "project" or "context", based on the resource type to sync to
     */
    public readonly resourceType!: pulumi.Output<string>;

    /**
     * Create a Circleci resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CircleciArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CircleciArgs | CircleciState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CircleciState | undefined;
            resourceInputs["config"] = state ? state.config : undefined;
            resourceInputs["deleteBehavior"] = state ? state.deleteBehavior : undefined;
            resourceInputs["integration"] = state ? state.integration : undefined;
            resourceInputs["organizationSlug"] = state ? state.organizationSlug : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["resourceId"] = state ? state.resourceId : undefined;
            resourceInputs["resourceType"] = state ? state.resourceType : undefined;
        } else {
            const args = argsOrState as CircleciArgs | undefined;
            if ((!args || args.config === undefined) && !opts.urn) {
                throw new Error("Missing required property 'config'");
            }
            if ((!args || args.integration === undefined) && !opts.urn) {
                throw new Error("Missing required property 'integration'");
            }
            if ((!args || args.organizationSlug === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organizationSlug'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.resourceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceId'");
            }
            if ((!args || args.resourceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceType'");
            }
            resourceInputs["config"] = args ? args.config : undefined;
            resourceInputs["deleteBehavior"] = args ? args.deleteBehavior : undefined;
            resourceInputs["integration"] = args ? args.integration : undefined;
            resourceInputs["organizationSlug"] = args ? args.organizationSlug : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["resourceId"] = args ? args.resourceId : undefined;
            resourceInputs["resourceType"] = args ? args.resourceType : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Circleci.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Circleci resources.
 */
export interface CircleciState {
    /**
     * The name of the Doppler config
     */
    config?: pulumi.Input<string>;
    /**
     * The behavior to be performed on the secrets in the sync target when this resource is deleted or recreated. Either `leaveInTarget` (default) or `deleteFromTarget`.
     */
    deleteBehavior?: pulumi.Input<string>;
    /**
     * The slug of the integration to use for this sync
     */
    integration?: pulumi.Input<string>;
    /**
     * The organization slug where the resource is located
     */
    organizationSlug?: pulumi.Input<string>;
    /**
     * The name of the Doppler project
     */
    project?: pulumi.Input<string>;
    /**
     * The resource ID (either project or context) to sync to
     */
    resourceId?: pulumi.Input<string>;
    /**
     * Either "project" or "context", based on the resource type to sync to
     */
    resourceType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Circleci resource.
 */
export interface CircleciArgs {
    /**
     * The name of the Doppler config
     */
    config: pulumi.Input<string>;
    /**
     * The behavior to be performed on the secrets in the sync target when this resource is deleted or recreated. Either `leaveInTarget` (default) or `deleteFromTarget`.
     */
    deleteBehavior?: pulumi.Input<string>;
    /**
     * The slug of the integration to use for this sync
     */
    integration: pulumi.Input<string>;
    /**
     * The organization slug where the resource is located
     */
    organizationSlug: pulumi.Input<string>;
    /**
     * The name of the Doppler project
     */
    project: pulumi.Input<string>;
    /**
     * The resource ID (either project or context) to sync to
     */
    resourceId: pulumi.Input<string>;
    /**
     * Either "project" or "context", based on the resource type to sync to
     */
    resourceType: pulumi.Input<string>;
}

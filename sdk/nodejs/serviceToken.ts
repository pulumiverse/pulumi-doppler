// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manage a Doppler service token.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as doppler from "@pulumiverse/doppler";
 *
 * const backendCiToken = new doppler.ServiceToken("backend_ci_token", {
 *     project: "backend",
 *     config: "ci",
 *     name: "Builder Token",
 *     access: "read",
 * });
 * ```
 */
export class ServiceToken extends pulumi.CustomResource {
    /**
     * Get an existing ServiceToken resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceTokenState, opts?: pulumi.CustomResourceOptions): ServiceToken {
        return new ServiceToken(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'doppler:index/serviceToken:ServiceToken';

    /**
     * Returns true if the given object is an instance of ServiceToken.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceToken {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceToken.__pulumiType;
    }

    /**
     * The access level (read or read/write)
     */
    public readonly access!: pulumi.Output<string | undefined>;
    /**
     * The name of the Doppler config where the service token is located
     */
    public readonly config!: pulumi.Output<string>;
    /**
     * The key for the Doppler service token
     */
    public /*out*/ readonly key!: pulumi.Output<string>;
    /**
     * The name of the Doppler service token
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the Doppler project where the service token is located
     */
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a ServiceToken resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceTokenArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceTokenArgs | ServiceTokenState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceTokenState | undefined;
            resourceInputs["access"] = state ? state.access : undefined;
            resourceInputs["config"] = state ? state.config : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as ServiceTokenArgs | undefined;
            if ((!args || args.config === undefined) && !opts.urn) {
                throw new Error("Missing required property 'config'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            resourceInputs["access"] = args ? args.access : undefined;
            resourceInputs["config"] = args ? args.config : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["key"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["key"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ServiceToken.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceToken resources.
 */
export interface ServiceTokenState {
    /**
     * The access level (read or read/write)
     */
    access?: pulumi.Input<string>;
    /**
     * The name of the Doppler config where the service token is located
     */
    config?: pulumi.Input<string>;
    /**
     * The key for the Doppler service token
     */
    key?: pulumi.Input<string>;
    /**
     * The name of the Doppler service token
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the Doppler project where the service token is located
     */
    project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServiceToken resource.
 */
export interface ServiceTokenArgs {
    /**
     * The access level (read or read/write)
     */
    access?: pulumi.Input<string>;
    /**
     * The name of the Doppler config where the service token is located
     */
    config: pulumi.Input<string>;
    /**
     * The name of the Doppler service token
     */
    name: pulumi.Input<string>;
    /**
     * The name of the Doppler project where the service token is located
     */
    project: pulumi.Input<string>;
}

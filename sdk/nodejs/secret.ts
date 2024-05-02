// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manage a single Doppler secret in a config.
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import doppler:index/secret:Secret default <project-name>.<config-name>.<secret-name>
 * ```
 */
export class Secret extends pulumi.CustomResource {
    /**
     * Get an existing Secret resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretState, opts?: pulumi.CustomResourceOptions): Secret {
        return new Secret(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'doppler:index/secret:Secret';

    /**
     * Returns true if the given object is an instance of Secret.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Secret {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Secret.__pulumiType;
    }

    /**
     * The computed secret value, after resolving secret references
     */
    public /*out*/ readonly computed!: pulumi.Output<string>;
    /**
     * The name of the Doppler config
     */
    public readonly config!: pulumi.Output<string>;
    /**
     * The name of the Doppler secret
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the Doppler project
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The raw secret value
     */
    public readonly value!: pulumi.Output<string>;
    /**
     * The visibility of the secret
     */
    public readonly visibility!: pulumi.Output<string | undefined>;

    /**
     * Create a Secret resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecretArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretArgs | SecretState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecretState | undefined;
            resourceInputs["computed"] = state ? state.computed : undefined;
            resourceInputs["config"] = state ? state.config : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["value"] = state ? state.value : undefined;
            resourceInputs["visibility"] = state ? state.visibility : undefined;
        } else {
            const args = argsOrState as SecretArgs | undefined;
            if ((!args || args.config === undefined) && !opts.urn) {
                throw new Error("Missing required property 'config'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.value === undefined) && !opts.urn) {
                throw new Error("Missing required property 'value'");
            }
            resourceInputs["config"] = args ? args.config : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["value"] = args?.value ? pulumi.secret(args.value) : undefined;
            resourceInputs["visibility"] = args ? args.visibility : undefined;
            resourceInputs["computed"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["computed", "value"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Secret.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Secret resources.
 */
export interface SecretState {
    /**
     * The computed secret value, after resolving secret references
     */
    computed?: pulumi.Input<string>;
    /**
     * The name of the Doppler config
     */
    config?: pulumi.Input<string>;
    /**
     * The name of the Doppler secret
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the Doppler project
     */
    project?: pulumi.Input<string>;
    /**
     * The raw secret value
     */
    value?: pulumi.Input<string>;
    /**
     * The visibility of the secret
     */
    visibility?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Secret resource.
 */
export interface SecretArgs {
    /**
     * The name of the Doppler config
     */
    config: pulumi.Input<string>;
    /**
     * The name of the Doppler secret
     */
    name: pulumi.Input<string>;
    /**
     * The name of the Doppler project
     */
    project: pulumi.Input<string>;
    /**
     * The raw secret value
     */
    value: pulumi.Input<string>;
    /**
     * The visibility of the secret
     */
    visibility?: pulumi.Input<string>;
}

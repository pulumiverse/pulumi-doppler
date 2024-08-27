// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { AwsParameterStoreArgs, AwsParameterStoreState } from "./awsParameterStore";
export type AwsParameterStore = import("./awsParameterStore").AwsParameterStore;
export const AwsParameterStore: typeof import("./awsParameterStore").AwsParameterStore = null as any;
utilities.lazyLoad(exports, ["AwsParameterStore"], () => require("./awsParameterStore"));

export { AwsSecretsManagerArgs, AwsSecretsManagerState } from "./awsSecretsManager";
export type AwsSecretsManager = import("./awsSecretsManager").AwsSecretsManager;
export const AwsSecretsManager: typeof import("./awsSecretsManager").AwsSecretsManager = null as any;
utilities.lazyLoad(exports, ["AwsSecretsManager"], () => require("./awsSecretsManager"));

export { CircleciArgs, CircleciState } from "./circleci";
export type Circleci = import("./circleci").Circleci;
export const Circleci: typeof import("./circleci").Circleci = null as any;
utilities.lazyLoad(exports, ["Circleci"], () => require("./circleci"));

export { FlyioArgs, FlyioState } from "./flyio";
export type Flyio = import("./flyio").Flyio;
export const Flyio: typeof import("./flyio").Flyio = null as any;
utilities.lazyLoad(exports, ["Flyio"], () => require("./flyio"));

export { GithubActionsArgs, GithubActionsState } from "./githubActions";
export type GithubActions = import("./githubActions").GithubActions;
export const GithubActions: typeof import("./githubActions").GithubActions = null as any;
utilities.lazyLoad(exports, ["GithubActions"], () => require("./githubActions"));

export { TerraformCloudArgs, TerraformCloudState } from "./terraformCloud";
export type TerraformCloud = import("./terraformCloud").TerraformCloud;
export const TerraformCloud: typeof import("./terraformCloud").TerraformCloud = null as any;
utilities.lazyLoad(exports, ["TerraformCloud"], () => require("./terraformCloud"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "doppler:secretsSync/awsParameterStore:AwsParameterStore":
                return new AwsParameterStore(name, <any>undefined, { urn })
            case "doppler:secretsSync/awsSecretsManager:AwsSecretsManager":
                return new AwsSecretsManager(name, <any>undefined, { urn })
            case "doppler:secretsSync/circleci:Circleci":
                return new Circleci(name, <any>undefined, { urn })
            case "doppler:secretsSync/flyio:Flyio":
                return new Flyio(name, <any>undefined, { urn })
            case "doppler:secretsSync/githubActions:GithubActions":
                return new GithubActions(name, <any>undefined, { urn })
            case "doppler:secretsSync/terraformCloud:TerraformCloud":
                return new TerraformCloud(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("doppler", "secretsSync/awsParameterStore", _module)
pulumi.runtime.registerResourceModule("doppler", "secretsSync/awsSecretsManager", _module)
pulumi.runtime.registerResourceModule("doppler", "secretsSync/circleci", _module)
pulumi.runtime.registerResourceModule("doppler", "secretsSync/flyio", _module)
pulumi.runtime.registerResourceModule("doppler", "secretsSync/githubActions", _module)
pulumi.runtime.registerResourceModule("doppler", "secretsSync/terraformCloud", _module)

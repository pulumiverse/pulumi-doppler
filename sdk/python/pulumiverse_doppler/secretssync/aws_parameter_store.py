# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['AwsParameterStoreArgs', 'AwsParameterStore']

@pulumi.input_type
class AwsParameterStoreArgs:
    def __init__(__self__, *,
                 config: pulumi.Input[str],
                 integration: pulumi.Input[str],
                 path: pulumi.Input[str],
                 project: pulumi.Input[str],
                 region: pulumi.Input[str],
                 delete_behavior: Optional[pulumi.Input[str]] = None,
                 kms_key_id: Optional[pulumi.Input[str]] = None,
                 secure_string: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a AwsParameterStore resource.
        :param pulumi.Input[str] config: The name of the Doppler config
        :param pulumi.Input[str] integration: The slug of the integration to use for this sync
        :param pulumi.Input[str] path: The path to the parameters in AWS
        :param pulumi.Input[str] project: The name of the Doppler project
        :param pulumi.Input[str] region: The AWS region
        :param pulumi.Input[str] delete_behavior: The behavior to be performed on the secrets in the sync target when this resource is deleted or recreated. Either `leave_in_target` (default) or `delete_from_target`.
        :param pulumi.Input[str] kms_key_id: The AWS KMS key used to encrypt the parameter (ID, Alias, or ARN)
        :param pulumi.Input[bool] secure_string: Whether or not the parameters are stored as a secure string
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: AWS tags to attach to the parameters
        """
        pulumi.set(__self__, "config", config)
        pulumi.set(__self__, "integration", integration)
        pulumi.set(__self__, "path", path)
        pulumi.set(__self__, "project", project)
        pulumi.set(__self__, "region", region)
        if delete_behavior is not None:
            pulumi.set(__self__, "delete_behavior", delete_behavior)
        if kms_key_id is not None:
            pulumi.set(__self__, "kms_key_id", kms_key_id)
        if secure_string is not None:
            pulumi.set(__self__, "secure_string", secure_string)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def config(self) -> pulumi.Input[str]:
        """
        The name of the Doppler config
        """
        return pulumi.get(self, "config")

    @config.setter
    def config(self, value: pulumi.Input[str]):
        pulumi.set(self, "config", value)

    @property
    @pulumi.getter
    def integration(self) -> pulumi.Input[str]:
        """
        The slug of the integration to use for this sync
        """
        return pulumi.get(self, "integration")

    @integration.setter
    def integration(self, value: pulumi.Input[str]):
        pulumi.set(self, "integration", value)

    @property
    @pulumi.getter
    def path(self) -> pulumi.Input[str]:
        """
        The path to the parameters in AWS
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: pulumi.Input[str]):
        pulumi.set(self, "path", value)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        """
        The name of the Doppler project
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def region(self) -> pulumi.Input[str]:
        """
        The AWS region
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input[str]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="deleteBehavior")
    def delete_behavior(self) -> Optional[pulumi.Input[str]]:
        """
        The behavior to be performed on the secrets in the sync target when this resource is deleted or recreated. Either `leave_in_target` (default) or `delete_from_target`.
        """
        return pulumi.get(self, "delete_behavior")

    @delete_behavior.setter
    def delete_behavior(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "delete_behavior", value)

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> Optional[pulumi.Input[str]]:
        """
        The AWS KMS key used to encrypt the parameter (ID, Alias, or ARN)
        """
        return pulumi.get(self, "kms_key_id")

    @kms_key_id.setter
    def kms_key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_key_id", value)

    @property
    @pulumi.getter(name="secureString")
    def secure_string(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not the parameters are stored as a secure string
        """
        return pulumi.get(self, "secure_string")

    @secure_string.setter
    def secure_string(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "secure_string", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        AWS tags to attach to the parameters
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _AwsParameterStoreState:
    def __init__(__self__, *,
                 config: Optional[pulumi.Input[str]] = None,
                 delete_behavior: Optional[pulumi.Input[str]] = None,
                 integration: Optional[pulumi.Input[str]] = None,
                 kms_key_id: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 secure_string: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering AwsParameterStore resources.
        :param pulumi.Input[str] config: The name of the Doppler config
        :param pulumi.Input[str] delete_behavior: The behavior to be performed on the secrets in the sync target when this resource is deleted or recreated. Either `leave_in_target` (default) or `delete_from_target`.
        :param pulumi.Input[str] integration: The slug of the integration to use for this sync
        :param pulumi.Input[str] kms_key_id: The AWS KMS key used to encrypt the parameter (ID, Alias, or ARN)
        :param pulumi.Input[str] path: The path to the parameters in AWS
        :param pulumi.Input[str] project: The name of the Doppler project
        :param pulumi.Input[str] region: The AWS region
        :param pulumi.Input[bool] secure_string: Whether or not the parameters are stored as a secure string
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: AWS tags to attach to the parameters
        """
        if config is not None:
            pulumi.set(__self__, "config", config)
        if delete_behavior is not None:
            pulumi.set(__self__, "delete_behavior", delete_behavior)
        if integration is not None:
            pulumi.set(__self__, "integration", integration)
        if kms_key_id is not None:
            pulumi.set(__self__, "kms_key_id", kms_key_id)
        if path is not None:
            pulumi.set(__self__, "path", path)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if secure_string is not None:
            pulumi.set(__self__, "secure_string", secure_string)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def config(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Doppler config
        """
        return pulumi.get(self, "config")

    @config.setter
    def config(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "config", value)

    @property
    @pulumi.getter(name="deleteBehavior")
    def delete_behavior(self) -> Optional[pulumi.Input[str]]:
        """
        The behavior to be performed on the secrets in the sync target when this resource is deleted or recreated. Either `leave_in_target` (default) or `delete_from_target`.
        """
        return pulumi.get(self, "delete_behavior")

    @delete_behavior.setter
    def delete_behavior(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "delete_behavior", value)

    @property
    @pulumi.getter
    def integration(self) -> Optional[pulumi.Input[str]]:
        """
        The slug of the integration to use for this sync
        """
        return pulumi.get(self, "integration")

    @integration.setter
    def integration(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "integration", value)

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> Optional[pulumi.Input[str]]:
        """
        The AWS KMS key used to encrypt the parameter (ID, Alias, or ARN)
        """
        return pulumi.get(self, "kms_key_id")

    @kms_key_id.setter
    def kms_key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_key_id", value)

    @property
    @pulumi.getter
    def path(self) -> Optional[pulumi.Input[str]]:
        """
        The path to the parameters in AWS
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "path", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Doppler project
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The AWS region
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="secureString")
    def secure_string(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not the parameters are stored as a secure string
        """
        return pulumi.get(self, "secure_string")

    @secure_string.setter
    def secure_string(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "secure_string", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        AWS tags to attach to the parameters
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


class AwsParameterStore(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config: Optional[pulumi.Input[str]] = None,
                 delete_behavior: Optional[pulumi.Input[str]] = None,
                 integration: Optional[pulumi.Input[str]] = None,
                 kms_key_id: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 secure_string: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Manage an AWS Parameter Store Doppler sync.

        ## Example Usage

        ```python
        import pulumi
        import json
        import pulumi_aws as aws
        import pulumiverse_doppler as doppler

        doppler_parameter_store = aws.iam.Role("doppler_parameter_store",
            name="doppler_parameter_store",
            assume_role_policy=json.dumps({
                "Version": "2012-10-17",
                "Statement": [{
                    "Effect": "Allow",
                    "Action": "sts:AssumeRole",
                    "Principal": {
                        "AWS": "arn:aws:iam::299900769157:user/doppler-integration-operator",
                    },
                    "Condition": {
                        "StringEquals": {
                            "sts:ExternalId": "<YOUR_WORKPLACE_SLUG>",
                        },
                    },
                }],
            }),
            inline_policies=[{
                "name": "doppler_secret_manager",
                "policy": json.dumps({
                    "version": "2012-10-17",
                    "statement": [{
                        "action": [
                            "ssm:PutParameter",
                            "ssm:LabelParameterVersion",
                            "ssm:DeleteParameter",
                            "ssm:RemoveTagsFromResource",
                            "ssm:GetParameterHistory",
                            "ssm:AddTagsToResource",
                            "ssm:GetParametersByPath",
                            "ssm:GetParameters",
                            "ssm:GetParameter",
                            "ssm:DeleteParameters",
                        ],
                        "effect": "Allow",
                        "resource": "*",
                    }],
                }),
            }])
        prod = doppler.integration.AwsParameterStore("prod",
            name="Production",
            assume_role_arn=doppler_parameter_store.arn)
        backend_prod = doppler.secrets_sync.AwsParameterStore("backend_prod",
            integration=prod.id,
            project="backend",
            config="prd",
            region="us-east-1",
            path="/backend/",
            secure_string=True,
            tags={
                "myTag": "enabled",
            },
            delete_behavior="leave_in_target")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] config: The name of the Doppler config
        :param pulumi.Input[str] delete_behavior: The behavior to be performed on the secrets in the sync target when this resource is deleted or recreated. Either `leave_in_target` (default) or `delete_from_target`.
        :param pulumi.Input[str] integration: The slug of the integration to use for this sync
        :param pulumi.Input[str] kms_key_id: The AWS KMS key used to encrypt the parameter (ID, Alias, or ARN)
        :param pulumi.Input[str] path: The path to the parameters in AWS
        :param pulumi.Input[str] project: The name of the Doppler project
        :param pulumi.Input[str] region: The AWS region
        :param pulumi.Input[bool] secure_string: Whether or not the parameters are stored as a secure string
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: AWS tags to attach to the parameters
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AwsParameterStoreArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manage an AWS Parameter Store Doppler sync.

        ## Example Usage

        ```python
        import pulumi
        import json
        import pulumi_aws as aws
        import pulumiverse_doppler as doppler

        doppler_parameter_store = aws.iam.Role("doppler_parameter_store",
            name="doppler_parameter_store",
            assume_role_policy=json.dumps({
                "Version": "2012-10-17",
                "Statement": [{
                    "Effect": "Allow",
                    "Action": "sts:AssumeRole",
                    "Principal": {
                        "AWS": "arn:aws:iam::299900769157:user/doppler-integration-operator",
                    },
                    "Condition": {
                        "StringEquals": {
                            "sts:ExternalId": "<YOUR_WORKPLACE_SLUG>",
                        },
                    },
                }],
            }),
            inline_policies=[{
                "name": "doppler_secret_manager",
                "policy": json.dumps({
                    "version": "2012-10-17",
                    "statement": [{
                        "action": [
                            "ssm:PutParameter",
                            "ssm:LabelParameterVersion",
                            "ssm:DeleteParameter",
                            "ssm:RemoveTagsFromResource",
                            "ssm:GetParameterHistory",
                            "ssm:AddTagsToResource",
                            "ssm:GetParametersByPath",
                            "ssm:GetParameters",
                            "ssm:GetParameter",
                            "ssm:DeleteParameters",
                        ],
                        "effect": "Allow",
                        "resource": "*",
                    }],
                }),
            }])
        prod = doppler.integration.AwsParameterStore("prod",
            name="Production",
            assume_role_arn=doppler_parameter_store.arn)
        backend_prod = doppler.secrets_sync.AwsParameterStore("backend_prod",
            integration=prod.id,
            project="backend",
            config="prd",
            region="us-east-1",
            path="/backend/",
            secure_string=True,
            tags={
                "myTag": "enabled",
            },
            delete_behavior="leave_in_target")
        ```

        :param str resource_name: The name of the resource.
        :param AwsParameterStoreArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AwsParameterStoreArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config: Optional[pulumi.Input[str]] = None,
                 delete_behavior: Optional[pulumi.Input[str]] = None,
                 integration: Optional[pulumi.Input[str]] = None,
                 kms_key_id: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 secure_string: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AwsParameterStoreArgs.__new__(AwsParameterStoreArgs)

            if config is None and not opts.urn:
                raise TypeError("Missing required property 'config'")
            __props__.__dict__["config"] = config
            __props__.__dict__["delete_behavior"] = delete_behavior
            if integration is None and not opts.urn:
                raise TypeError("Missing required property 'integration'")
            __props__.__dict__["integration"] = integration
            __props__.__dict__["kms_key_id"] = kms_key_id
            if path is None and not opts.urn:
                raise TypeError("Missing required property 'path'")
            __props__.__dict__["path"] = path
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            if region is None and not opts.urn:
                raise TypeError("Missing required property 'region'")
            __props__.__dict__["region"] = region
            __props__.__dict__["secure_string"] = secure_string
            __props__.__dict__["tags"] = tags
        super(AwsParameterStore, __self__).__init__(
            'doppler:secretsSync/awsParameterStore:AwsParameterStore',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            config: Optional[pulumi.Input[str]] = None,
            delete_behavior: Optional[pulumi.Input[str]] = None,
            integration: Optional[pulumi.Input[str]] = None,
            kms_key_id: Optional[pulumi.Input[str]] = None,
            path: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            secure_string: Optional[pulumi.Input[bool]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'AwsParameterStore':
        """
        Get an existing AwsParameterStore resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] config: The name of the Doppler config
        :param pulumi.Input[str] delete_behavior: The behavior to be performed on the secrets in the sync target when this resource is deleted or recreated. Either `leave_in_target` (default) or `delete_from_target`.
        :param pulumi.Input[str] integration: The slug of the integration to use for this sync
        :param pulumi.Input[str] kms_key_id: The AWS KMS key used to encrypt the parameter (ID, Alias, or ARN)
        :param pulumi.Input[str] path: The path to the parameters in AWS
        :param pulumi.Input[str] project: The name of the Doppler project
        :param pulumi.Input[str] region: The AWS region
        :param pulumi.Input[bool] secure_string: Whether or not the parameters are stored as a secure string
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: AWS tags to attach to the parameters
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AwsParameterStoreState.__new__(_AwsParameterStoreState)

        __props__.__dict__["config"] = config
        __props__.__dict__["delete_behavior"] = delete_behavior
        __props__.__dict__["integration"] = integration
        __props__.__dict__["kms_key_id"] = kms_key_id
        __props__.__dict__["path"] = path
        __props__.__dict__["project"] = project
        __props__.__dict__["region"] = region
        __props__.__dict__["secure_string"] = secure_string
        __props__.__dict__["tags"] = tags
        return AwsParameterStore(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def config(self) -> pulumi.Output[str]:
        """
        The name of the Doppler config
        """
        return pulumi.get(self, "config")

    @property
    @pulumi.getter(name="deleteBehavior")
    def delete_behavior(self) -> pulumi.Output[Optional[str]]:
        """
        The behavior to be performed on the secrets in the sync target when this resource is deleted or recreated. Either `leave_in_target` (default) or `delete_from_target`.
        """
        return pulumi.get(self, "delete_behavior")

    @property
    @pulumi.getter
    def integration(self) -> pulumi.Output[str]:
        """
        The slug of the integration to use for this sync
        """
        return pulumi.get(self, "integration")

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> pulumi.Output[Optional[str]]:
        """
        The AWS KMS key used to encrypt the parameter (ID, Alias, or ARN)
        """
        return pulumi.get(self, "kms_key_id")

    @property
    @pulumi.getter
    def path(self) -> pulumi.Output[str]:
        """
        The path to the parameters in AWS
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The name of the Doppler project
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The AWS region
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="secureString")
    def secure_string(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether or not the parameters are stored as a secure string
        """
        return pulumi.get(self, "secure_string")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        AWS tags to attach to the parameters
        """
        return pulumi.get(self, "tags")


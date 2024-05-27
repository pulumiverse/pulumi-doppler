# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['TerraformCloudArgs', 'TerraformCloud']

@pulumi.input_type
class TerraformCloudArgs:
    def __init__(__self__, *,
                 config: pulumi.Input[str],
                 integration: pulumi.Input[str],
                 name_transform: pulumi.Input[str],
                 project: pulumi.Input[str],
                 sync_target: pulumi.Input[str],
                 variable_sync_type: pulumi.Input[str],
                 variable_set_id: Optional[pulumi.Input[str]] = None,
                 workspace_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a TerraformCloud resource.
        :param pulumi.Input[str] config: The name of the Doppler config
        :param pulumi.Input[str] integration: The slug of the integration to use for this sync
        :param pulumi.Input[str] name_transform: A name transform to apply before syncing secrets: "none" or "lowercase"
        :param pulumi.Input[str] project: The name of the Doppler project
        :param pulumi.Input[str] sync_target: Either "workspace" or "variableSet", based on the resource type to sync to
        """
        pulumi.set(__self__, "config", config)
        pulumi.set(__self__, "integration", integration)
        pulumi.set(__self__, "name_transform", name_transform)
        pulumi.set(__self__, "project", project)
        pulumi.set(__self__, "sync_target", sync_target)
        pulumi.set(__self__, "variable_sync_type", variable_sync_type)
        if variable_set_id is not None:
            pulumi.set(__self__, "variable_set_id", variable_set_id)
        if workspace_id is not None:
            pulumi.set(__self__, "workspace_id", workspace_id)

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
    @pulumi.getter(name="nameTransform")
    def name_transform(self) -> pulumi.Input[str]:
        """
        A name transform to apply before syncing secrets: "none" or "lowercase"
        """
        return pulumi.get(self, "name_transform")

    @name_transform.setter
    def name_transform(self, value: pulumi.Input[str]):
        pulumi.set(self, "name_transform", value)

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
    @pulumi.getter(name="syncTarget")
    def sync_target(self) -> pulumi.Input[str]:
        """
        Either "workspace" or "variableSet", based on the resource type to sync to
        """
        return pulumi.get(self, "sync_target")

    @sync_target.setter
    def sync_target(self, value: pulumi.Input[str]):
        pulumi.set(self, "sync_target", value)

    @property
    @pulumi.getter(name="variableSyncType")
    def variable_sync_type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "variable_sync_type")

    @variable_sync_type.setter
    def variable_sync_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "variable_sync_type", value)

    @property
    @pulumi.getter(name="variableSetId")
    def variable_set_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "variable_set_id")

    @variable_set_id.setter
    def variable_set_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "variable_set_id", value)

    @property
    @pulumi.getter(name="workspaceId")
    def workspace_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "workspace_id")

    @workspace_id.setter
    def workspace_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "workspace_id", value)


@pulumi.input_type
class _TerraformCloudState:
    def __init__(__self__, *,
                 config: Optional[pulumi.Input[str]] = None,
                 integration: Optional[pulumi.Input[str]] = None,
                 name_transform: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 sync_target: Optional[pulumi.Input[str]] = None,
                 variable_set_id: Optional[pulumi.Input[str]] = None,
                 variable_sync_type: Optional[pulumi.Input[str]] = None,
                 workspace_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering TerraformCloud resources.
        :param pulumi.Input[str] config: The name of the Doppler config
        :param pulumi.Input[str] integration: The slug of the integration to use for this sync
        :param pulumi.Input[str] name_transform: A name transform to apply before syncing secrets: "none" or "lowercase"
        :param pulumi.Input[str] project: The name of the Doppler project
        :param pulumi.Input[str] sync_target: Either "workspace" or "variableSet", based on the resource type to sync to
        """
        if config is not None:
            pulumi.set(__self__, "config", config)
        if integration is not None:
            pulumi.set(__self__, "integration", integration)
        if name_transform is not None:
            pulumi.set(__self__, "name_transform", name_transform)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if sync_target is not None:
            pulumi.set(__self__, "sync_target", sync_target)
        if variable_set_id is not None:
            pulumi.set(__self__, "variable_set_id", variable_set_id)
        if variable_sync_type is not None:
            pulumi.set(__self__, "variable_sync_type", variable_sync_type)
        if workspace_id is not None:
            pulumi.set(__self__, "workspace_id", workspace_id)

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
    @pulumi.getter(name="nameTransform")
    def name_transform(self) -> Optional[pulumi.Input[str]]:
        """
        A name transform to apply before syncing secrets: "none" or "lowercase"
        """
        return pulumi.get(self, "name_transform")

    @name_transform.setter
    def name_transform(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name_transform", value)

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
    @pulumi.getter(name="syncTarget")
    def sync_target(self) -> Optional[pulumi.Input[str]]:
        """
        Either "workspace" or "variableSet", based on the resource type to sync to
        """
        return pulumi.get(self, "sync_target")

    @sync_target.setter
    def sync_target(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sync_target", value)

    @property
    @pulumi.getter(name="variableSetId")
    def variable_set_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "variable_set_id")

    @variable_set_id.setter
    def variable_set_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "variable_set_id", value)

    @property
    @pulumi.getter(name="variableSyncType")
    def variable_sync_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "variable_sync_type")

    @variable_sync_type.setter
    def variable_sync_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "variable_sync_type", value)

    @property
    @pulumi.getter(name="workspaceId")
    def workspace_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "workspace_id")

    @workspace_id.setter
    def workspace_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "workspace_id", value)


class TerraformCloud(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config: Optional[pulumi.Input[str]] = None,
                 integration: Optional[pulumi.Input[str]] = None,
                 name_transform: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 sync_target: Optional[pulumi.Input[str]] = None,
                 variable_set_id: Optional[pulumi.Input[str]] = None,
                 variable_sync_type: Optional[pulumi.Input[str]] = None,
                 workspace_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] config: The name of the Doppler config
        :param pulumi.Input[str] integration: The slug of the integration to use for this sync
        :param pulumi.Input[str] name_transform: A name transform to apply before syncing secrets: "none" or "lowercase"
        :param pulumi.Input[str] project: The name of the Doppler project
        :param pulumi.Input[str] sync_target: Either "workspace" or "variableSet", based on the resource type to sync to
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TerraformCloudArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        :param str resource_name: The name of the resource.
        :param TerraformCloudArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TerraformCloudArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config: Optional[pulumi.Input[str]] = None,
                 integration: Optional[pulumi.Input[str]] = None,
                 name_transform: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 sync_target: Optional[pulumi.Input[str]] = None,
                 variable_set_id: Optional[pulumi.Input[str]] = None,
                 variable_sync_type: Optional[pulumi.Input[str]] = None,
                 workspace_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TerraformCloudArgs.__new__(TerraformCloudArgs)

            if config is None and not opts.urn:
                raise TypeError("Missing required property 'config'")
            __props__.__dict__["config"] = config
            if integration is None and not opts.urn:
                raise TypeError("Missing required property 'integration'")
            __props__.__dict__["integration"] = integration
            if name_transform is None and not opts.urn:
                raise TypeError("Missing required property 'name_transform'")
            __props__.__dict__["name_transform"] = name_transform
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            if sync_target is None and not opts.urn:
                raise TypeError("Missing required property 'sync_target'")
            __props__.__dict__["sync_target"] = sync_target
            __props__.__dict__["variable_set_id"] = variable_set_id
            if variable_sync_type is None and not opts.urn:
                raise TypeError("Missing required property 'variable_sync_type'")
            __props__.__dict__["variable_sync_type"] = variable_sync_type
            __props__.__dict__["workspace_id"] = workspace_id
        super(TerraformCloud, __self__).__init__(
            'doppler:secretsSync/terraformCloud:TerraformCloud',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            config: Optional[pulumi.Input[str]] = None,
            integration: Optional[pulumi.Input[str]] = None,
            name_transform: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            sync_target: Optional[pulumi.Input[str]] = None,
            variable_set_id: Optional[pulumi.Input[str]] = None,
            variable_sync_type: Optional[pulumi.Input[str]] = None,
            workspace_id: Optional[pulumi.Input[str]] = None) -> 'TerraformCloud':
        """
        Get an existing TerraformCloud resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] config: The name of the Doppler config
        :param pulumi.Input[str] integration: The slug of the integration to use for this sync
        :param pulumi.Input[str] name_transform: A name transform to apply before syncing secrets: "none" or "lowercase"
        :param pulumi.Input[str] project: The name of the Doppler project
        :param pulumi.Input[str] sync_target: Either "workspace" or "variableSet", based on the resource type to sync to
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TerraformCloudState.__new__(_TerraformCloudState)

        __props__.__dict__["config"] = config
        __props__.__dict__["integration"] = integration
        __props__.__dict__["name_transform"] = name_transform
        __props__.__dict__["project"] = project
        __props__.__dict__["sync_target"] = sync_target
        __props__.__dict__["variable_set_id"] = variable_set_id
        __props__.__dict__["variable_sync_type"] = variable_sync_type
        __props__.__dict__["workspace_id"] = workspace_id
        return TerraformCloud(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def config(self) -> pulumi.Output[str]:
        """
        The name of the Doppler config
        """
        return pulumi.get(self, "config")

    @property
    @pulumi.getter
    def integration(self) -> pulumi.Output[str]:
        """
        The slug of the integration to use for this sync
        """
        return pulumi.get(self, "integration")

    @property
    @pulumi.getter(name="nameTransform")
    def name_transform(self) -> pulumi.Output[str]:
        """
        A name transform to apply before syncing secrets: "none" or "lowercase"
        """
        return pulumi.get(self, "name_transform")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The name of the Doppler project
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="syncTarget")
    def sync_target(self) -> pulumi.Output[str]:
        """
        Either "workspace" or "variableSet", based on the resource type to sync to
        """
        return pulumi.get(self, "sync_target")

    @property
    @pulumi.getter(name="variableSetId")
    def variable_set_id(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "variable_set_id")

    @property
    @pulumi.getter(name="variableSyncType")
    def variable_sync_type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "variable_sync_type")

    @property
    @pulumi.getter(name="workspaceId")
    def workspace_id(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "workspace_id")


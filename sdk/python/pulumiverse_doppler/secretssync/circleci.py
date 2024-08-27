# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['CircleciArgs', 'Circleci']

@pulumi.input_type
class CircleciArgs:
    def __init__(__self__, *,
                 config: pulumi.Input[str],
                 integration: pulumi.Input[str],
                 organization_slug: pulumi.Input[str],
                 project: pulumi.Input[str],
                 resource_id: pulumi.Input[str],
                 resource_type: pulumi.Input[str],
                 delete_behavior: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Circleci resource.
        :param pulumi.Input[str] config: The name of the Doppler config
        :param pulumi.Input[str] integration: The slug of the integration to use for this sync
        :param pulumi.Input[str] organization_slug: The organization slug where the resource is located
        :param pulumi.Input[str] project: The name of the Doppler project
        :param pulumi.Input[str] resource_id: The resource ID (either project or context) to sync to
        :param pulumi.Input[str] resource_type: Either "project" or "context", based on the resource type to sync to
        :param pulumi.Input[str] delete_behavior: The behavior to be performed on the secrets in the sync target when this resource is deleted or recreated. Either `leave_in_target` (default) or `delete_from_target`.
        """
        pulumi.set(__self__, "config", config)
        pulumi.set(__self__, "integration", integration)
        pulumi.set(__self__, "organization_slug", organization_slug)
        pulumi.set(__self__, "project", project)
        pulumi.set(__self__, "resource_id", resource_id)
        pulumi.set(__self__, "resource_type", resource_type)
        if delete_behavior is not None:
            pulumi.set(__self__, "delete_behavior", delete_behavior)

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
    @pulumi.getter(name="organizationSlug")
    def organization_slug(self) -> pulumi.Input[str]:
        """
        The organization slug where the resource is located
        """
        return pulumi.get(self, "organization_slug")

    @organization_slug.setter
    def organization_slug(self, value: pulumi.Input[str]):
        pulumi.set(self, "organization_slug", value)

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
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Input[str]:
        """
        The resource ID (either project or context) to sync to
        """
        return pulumi.get(self, "resource_id")

    @resource_id.setter
    def resource_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_id", value)

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> pulumi.Input[str]:
        """
        Either "project" or "context", based on the resource type to sync to
        """
        return pulumi.get(self, "resource_type")

    @resource_type.setter
    def resource_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_type", value)

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


@pulumi.input_type
class _CircleciState:
    def __init__(__self__, *,
                 config: Optional[pulumi.Input[str]] = None,
                 delete_behavior: Optional[pulumi.Input[str]] = None,
                 integration: Optional[pulumi.Input[str]] = None,
                 organization_slug: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 resource_type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Circleci resources.
        :param pulumi.Input[str] config: The name of the Doppler config
        :param pulumi.Input[str] delete_behavior: The behavior to be performed on the secrets in the sync target when this resource is deleted or recreated. Either `leave_in_target` (default) or `delete_from_target`.
        :param pulumi.Input[str] integration: The slug of the integration to use for this sync
        :param pulumi.Input[str] organization_slug: The organization slug where the resource is located
        :param pulumi.Input[str] project: The name of the Doppler project
        :param pulumi.Input[str] resource_id: The resource ID (either project or context) to sync to
        :param pulumi.Input[str] resource_type: Either "project" or "context", based on the resource type to sync to
        """
        if config is not None:
            pulumi.set(__self__, "config", config)
        if delete_behavior is not None:
            pulumi.set(__self__, "delete_behavior", delete_behavior)
        if integration is not None:
            pulumi.set(__self__, "integration", integration)
        if organization_slug is not None:
            pulumi.set(__self__, "organization_slug", organization_slug)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if resource_id is not None:
            pulumi.set(__self__, "resource_id", resource_id)
        if resource_type is not None:
            pulumi.set(__self__, "resource_type", resource_type)

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
    @pulumi.getter(name="organizationSlug")
    def organization_slug(self) -> Optional[pulumi.Input[str]]:
        """
        The organization slug where the resource is located
        """
        return pulumi.get(self, "organization_slug")

    @organization_slug.setter
    def organization_slug(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "organization_slug", value)

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
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> Optional[pulumi.Input[str]]:
        """
        The resource ID (either project or context) to sync to
        """
        return pulumi.get(self, "resource_id")

    @resource_id.setter
    def resource_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_id", value)

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> Optional[pulumi.Input[str]]:
        """
        Either "project" or "context", based on the resource type to sync to
        """
        return pulumi.get(self, "resource_type")

    @resource_type.setter
    def resource_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_type", value)


class Circleci(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config: Optional[pulumi.Input[str]] = None,
                 delete_behavior: Optional[pulumi.Input[str]] = None,
                 integration: Optional[pulumi.Input[str]] = None,
                 organization_slug: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 resource_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manage a CircleCI Doppler sync.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_doppler as doppler

        prod = doppler.integration.Circleci("prod",
            name="Production",
            api_token="my_api_token")
        backend_prod = doppler.secrets_sync.Circleci("backend_prod",
            integration=prod.id,
            project="backend",
            config="prd",
            resource_type="project",
            resource_id="github/myorg/myproject",
            organization_slug="myorg",
            delete_behavior="leave_in_target")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] config: The name of the Doppler config
        :param pulumi.Input[str] delete_behavior: The behavior to be performed on the secrets in the sync target when this resource is deleted or recreated. Either `leave_in_target` (default) or `delete_from_target`.
        :param pulumi.Input[str] integration: The slug of the integration to use for this sync
        :param pulumi.Input[str] organization_slug: The organization slug where the resource is located
        :param pulumi.Input[str] project: The name of the Doppler project
        :param pulumi.Input[str] resource_id: The resource ID (either project or context) to sync to
        :param pulumi.Input[str] resource_type: Either "project" or "context", based on the resource type to sync to
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CircleciArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manage a CircleCI Doppler sync.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_doppler as doppler

        prod = doppler.integration.Circleci("prod",
            name="Production",
            api_token="my_api_token")
        backend_prod = doppler.secrets_sync.Circleci("backend_prod",
            integration=prod.id,
            project="backend",
            config="prd",
            resource_type="project",
            resource_id="github/myorg/myproject",
            organization_slug="myorg",
            delete_behavior="leave_in_target")
        ```

        :param str resource_name: The name of the resource.
        :param CircleciArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CircleciArgs, pulumi.ResourceOptions, *args, **kwargs)
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
                 organization_slug: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 resource_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CircleciArgs.__new__(CircleciArgs)

            if config is None and not opts.urn:
                raise TypeError("Missing required property 'config'")
            __props__.__dict__["config"] = config
            __props__.__dict__["delete_behavior"] = delete_behavior
            if integration is None and not opts.urn:
                raise TypeError("Missing required property 'integration'")
            __props__.__dict__["integration"] = integration
            if organization_slug is None and not opts.urn:
                raise TypeError("Missing required property 'organization_slug'")
            __props__.__dict__["organization_slug"] = organization_slug
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            if resource_id is None and not opts.urn:
                raise TypeError("Missing required property 'resource_id'")
            __props__.__dict__["resource_id"] = resource_id
            if resource_type is None and not opts.urn:
                raise TypeError("Missing required property 'resource_type'")
            __props__.__dict__["resource_type"] = resource_type
        super(Circleci, __self__).__init__(
            'doppler:secretsSync/circleci:Circleci',
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
            organization_slug: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            resource_id: Optional[pulumi.Input[str]] = None,
            resource_type: Optional[pulumi.Input[str]] = None) -> 'Circleci':
        """
        Get an existing Circleci resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] config: The name of the Doppler config
        :param pulumi.Input[str] delete_behavior: The behavior to be performed on the secrets in the sync target when this resource is deleted or recreated. Either `leave_in_target` (default) or `delete_from_target`.
        :param pulumi.Input[str] integration: The slug of the integration to use for this sync
        :param pulumi.Input[str] organization_slug: The organization slug where the resource is located
        :param pulumi.Input[str] project: The name of the Doppler project
        :param pulumi.Input[str] resource_id: The resource ID (either project or context) to sync to
        :param pulumi.Input[str] resource_type: Either "project" or "context", based on the resource type to sync to
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CircleciState.__new__(_CircleciState)

        __props__.__dict__["config"] = config
        __props__.__dict__["delete_behavior"] = delete_behavior
        __props__.__dict__["integration"] = integration
        __props__.__dict__["organization_slug"] = organization_slug
        __props__.__dict__["project"] = project
        __props__.__dict__["resource_id"] = resource_id
        __props__.__dict__["resource_type"] = resource_type
        return Circleci(resource_name, opts=opts, __props__=__props__)

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
    @pulumi.getter(name="organizationSlug")
    def organization_slug(self) -> pulumi.Output[str]:
        """
        The organization slug where the resource is located
        """
        return pulumi.get(self, "organization_slug")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The name of the Doppler project
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Output[str]:
        """
        The resource ID (either project or context) to sync to
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> pulumi.Output[str]:
        """
        Either "project" or "context", based on the resource type to sync to
        """
        return pulumi.get(self, "resource_type")

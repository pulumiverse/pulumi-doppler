# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'WebhookAuthentication',
]

@pulumi.output_type
class WebhookAuthentication(dict):
    def __init__(__self__, *,
                 type: str,
                 password: Optional[str] = None,
                 token: Optional[str] = None,
                 username: Optional[str] = None):
        pulumi.set(__self__, "type", type)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if token is not None:
            pulumi.set(__self__, "token", token)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def password(self) -> Optional[str]:
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def token(self) -> Optional[str]:
        return pulumi.get(self, "token")

    @property
    @pulumi.getter
    def username(self) -> Optional[str]:
        return pulumi.get(self, "username")



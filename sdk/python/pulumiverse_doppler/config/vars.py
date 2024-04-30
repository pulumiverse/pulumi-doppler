# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

import types

__config__ = pulumi.Config('doppler')


class _ExportableConfig(types.ModuleType):
    @property
    def doppler_token(self) -> Optional[str]:
        """
        A Doppler token, either a personal or service token
        """
        return __config__.get('dopplerToken')

    @property
    def host(self) -> Optional[str]:
        """
        The Doppler API host (i.e. https://api.doppler.com)
        """
        return __config__.get('host')

    @property
    def verify_tls(self) -> Optional[bool]:
        """
        Whether or not to verify TLS
        """
        return __config__.get_bool('verifyTls')


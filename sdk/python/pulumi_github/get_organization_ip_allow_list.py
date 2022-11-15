# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetOrganizationIpAllowListResult',
    'AwaitableGetOrganizationIpAllowListResult',
    'get_organization_ip_allow_list',
]

@pulumi.output_type
class GetOrganizationIpAllowListResult:
    """
    A collection of values returned by getOrganizationIpAllowList.
    """
    def __init__(__self__, id=None, ip_allow_lists=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip_allow_lists and not isinstance(ip_allow_lists, list):
            raise TypeError("Expected argument 'ip_allow_lists' to be a list")
        pulumi.set(__self__, "ip_allow_lists", ip_allow_lists)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ipAllowLists")
    def ip_allow_lists(self) -> Sequence['outputs.GetOrganizationIpAllowListIpAllowListResult']:
        return pulumi.get(self, "ip_allow_lists")


class AwaitableGetOrganizationIpAllowListResult(GetOrganizationIpAllowListResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOrganizationIpAllowListResult(
            id=self.id,
            ip_allow_lists=self.ip_allow_lists)


def get_organization_ip_allow_list(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOrganizationIpAllowListResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('github:index/getOrganizationIpAllowList:getOrganizationIpAllowList', __args__, opts=opts, typ=GetOrganizationIpAllowListResult).value

    return AwaitableGetOrganizationIpAllowListResult(
        id=__ret__.id,
        ip_allow_lists=__ret__.ip_allow_lists)

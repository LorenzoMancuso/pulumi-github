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
    'GetExternalGroupsResult',
    'AwaitableGetExternalGroupsResult',
    'get_external_groups',
]

@pulumi.output_type
class GetExternalGroupsResult:
    """
    A collection of values returned by getExternalGroups.
    """
    def __init__(__self__, external_groups=None, id=None):
        if external_groups and not isinstance(external_groups, list):
            raise TypeError("Expected argument 'external_groups' to be a list")
        pulumi.set(__self__, "external_groups", external_groups)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter(name="externalGroups")
    def external_groups(self) -> Sequence['outputs.GetExternalGroupsExternalGroupResult']:
        return pulumi.get(self, "external_groups")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")


class AwaitableGetExternalGroupsResult(GetExternalGroupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetExternalGroupsResult(
            external_groups=self.external_groups,
            id=self.id)


def get_external_groups(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetExternalGroupsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('github:index/getExternalGroups:getExternalGroups', __args__, opts=opts, typ=GetExternalGroupsResult).value

    return AwaitableGetExternalGroupsResult(
        external_groups=__ret__.external_groups,
        id=__ret__.id)

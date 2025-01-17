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
    'GetOrganizationResult',
    'AwaitableGetOrganizationResult',
    'get_organization',
    'get_organization_output',
]

@pulumi.output_type
class GetOrganizationResult:
    """
    A collection of values returned by getOrganization.
    """
    def __init__(__self__, description=None, id=None, login=None, members=None, name=None, node_id=None, orgname=None, plan=None, repositories=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if login and not isinstance(login, str):
            raise TypeError("Expected argument 'login' to be a str")
        pulumi.set(__self__, "login", login)
        if members and not isinstance(members, list):
            raise TypeError("Expected argument 'members' to be a list")
        pulumi.set(__self__, "members", members)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if node_id and not isinstance(node_id, str):
            raise TypeError("Expected argument 'node_id' to be a str")
        pulumi.set(__self__, "node_id", node_id)
        if orgname and not isinstance(orgname, str):
            raise TypeError("Expected argument 'orgname' to be a str")
        pulumi.set(__self__, "orgname", orgname)
        if plan and not isinstance(plan, str):
            raise TypeError("Expected argument 'plan' to be a str")
        pulumi.set(__self__, "plan", plan)
        if repositories and not isinstance(repositories, list):
            raise TypeError("Expected argument 'repositories' to be a list")
        pulumi.set(__self__, "repositories", repositories)

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def login(self) -> str:
        return pulumi.get(self, "login")

    @property
    @pulumi.getter
    def members(self) -> Sequence[str]:
        return pulumi.get(self, "members")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nodeId")
    def node_id(self) -> str:
        return pulumi.get(self, "node_id")

    @property
    @pulumi.getter
    def orgname(self) -> str:
        return pulumi.get(self, "orgname")

    @property
    @pulumi.getter
    def plan(self) -> str:
        return pulumi.get(self, "plan")

    @property
    @pulumi.getter
    def repositories(self) -> Sequence[str]:
        return pulumi.get(self, "repositories")


class AwaitableGetOrganizationResult(GetOrganizationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOrganizationResult(
            description=self.description,
            id=self.id,
            login=self.login,
            members=self.members,
            name=self.name,
            node_id=self.node_id,
            orgname=self.orgname,
            plan=self.plan,
            repositories=self.repositories)


def get_organization(name: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOrganizationResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('github:index/getOrganization:getOrganization', __args__, opts=opts, typ=GetOrganizationResult).value

    return AwaitableGetOrganizationResult(
        description=__ret__.description,
        id=__ret__.id,
        login=__ret__.login,
        members=__ret__.members,
        name=__ret__.name,
        node_id=__ret__.node_id,
        orgname=__ret__.orgname,
        plan=__ret__.plan,
        repositories=__ret__.repositories)


@_utilities.lift_output_func(get_organization)
def get_organization_output(name: Optional[pulumi.Input[str]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetOrganizationResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...

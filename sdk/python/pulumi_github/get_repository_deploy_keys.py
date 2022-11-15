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
    'GetRepositoryDeployKeysResult',
    'AwaitableGetRepositoryDeployKeysResult',
    'get_repository_deploy_keys',
    'get_repository_deploy_keys_output',
]

@pulumi.output_type
class GetRepositoryDeployKeysResult:
    """
    A collection of values returned by getRepositoryDeployKeys.
    """
    def __init__(__self__, id=None, keys=None, repository=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if keys and not isinstance(keys, list):
            raise TypeError("Expected argument 'keys' to be a list")
        pulumi.set(__self__, "keys", keys)
        if repository and not isinstance(repository, str):
            raise TypeError("Expected argument 'repository' to be a str")
        pulumi.set(__self__, "repository", repository)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def keys(self) -> Sequence['outputs.GetRepositoryDeployKeysKeyResult']:
        return pulumi.get(self, "keys")

    @property
    @pulumi.getter
    def repository(self) -> str:
        return pulumi.get(self, "repository")


class AwaitableGetRepositoryDeployKeysResult(GetRepositoryDeployKeysResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRepositoryDeployKeysResult(
            id=self.id,
            keys=self.keys,
            repository=self.repository)


def get_repository_deploy_keys(repository: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRepositoryDeployKeysResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['repository'] = repository
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('github:index/getRepositoryDeployKeys:getRepositoryDeployKeys', __args__, opts=opts, typ=GetRepositoryDeployKeysResult).value

    return AwaitableGetRepositoryDeployKeysResult(
        id=__ret__.id,
        keys=__ret__.keys,
        repository=__ret__.repository)


@_utilities.lift_output_func(get_repository_deploy_keys)
def get_repository_deploy_keys_output(repository: Optional[pulumi.Input[str]] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRepositoryDeployKeysResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...

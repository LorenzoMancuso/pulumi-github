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
    'GetReleaseResult',
    'AwaitableGetReleaseResult',
    'get_release',
    'get_release_output',
]

@pulumi.output_type
class GetReleaseResult:
    """
    A collection of values returned by getRelease.
    """
    def __init__(__self__, asserts_url=None, assets=None, body=None, created_at=None, draft=None, html_url=None, id=None, name=None, owner=None, prerelease=None, published_at=None, release_id=None, release_tag=None, repository=None, retrieve_by=None, tarball_url=None, target_commitish=None, upload_url=None, url=None, zipball_url=None):
        if asserts_url and not isinstance(asserts_url, str):
            raise TypeError("Expected argument 'asserts_url' to be a str")
        pulumi.set(__self__, "asserts_url", asserts_url)
        if assets and not isinstance(assets, list):
            raise TypeError("Expected argument 'assets' to be a list")
        pulumi.set(__self__, "assets", assets)
        if body and not isinstance(body, str):
            raise TypeError("Expected argument 'body' to be a str")
        pulumi.set(__self__, "body", body)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if draft and not isinstance(draft, bool):
            raise TypeError("Expected argument 'draft' to be a bool")
        pulumi.set(__self__, "draft", draft)
        if html_url and not isinstance(html_url, str):
            raise TypeError("Expected argument 'html_url' to be a str")
        pulumi.set(__self__, "html_url", html_url)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if owner and not isinstance(owner, str):
            raise TypeError("Expected argument 'owner' to be a str")
        pulumi.set(__self__, "owner", owner)
        if prerelease and not isinstance(prerelease, bool):
            raise TypeError("Expected argument 'prerelease' to be a bool")
        pulumi.set(__self__, "prerelease", prerelease)
        if published_at and not isinstance(published_at, str):
            raise TypeError("Expected argument 'published_at' to be a str")
        pulumi.set(__self__, "published_at", published_at)
        if release_id and not isinstance(release_id, int):
            raise TypeError("Expected argument 'release_id' to be a int")
        pulumi.set(__self__, "release_id", release_id)
        if release_tag and not isinstance(release_tag, str):
            raise TypeError("Expected argument 'release_tag' to be a str")
        pulumi.set(__self__, "release_tag", release_tag)
        if repository and not isinstance(repository, str):
            raise TypeError("Expected argument 'repository' to be a str")
        pulumi.set(__self__, "repository", repository)
        if retrieve_by and not isinstance(retrieve_by, str):
            raise TypeError("Expected argument 'retrieve_by' to be a str")
        pulumi.set(__self__, "retrieve_by", retrieve_by)
        if tarball_url and not isinstance(tarball_url, str):
            raise TypeError("Expected argument 'tarball_url' to be a str")
        pulumi.set(__self__, "tarball_url", tarball_url)
        if target_commitish and not isinstance(target_commitish, str):
            raise TypeError("Expected argument 'target_commitish' to be a str")
        pulumi.set(__self__, "target_commitish", target_commitish)
        if upload_url and not isinstance(upload_url, str):
            raise TypeError("Expected argument 'upload_url' to be a str")
        pulumi.set(__self__, "upload_url", upload_url)
        if url and not isinstance(url, str):
            raise TypeError("Expected argument 'url' to be a str")
        pulumi.set(__self__, "url", url)
        if zipball_url and not isinstance(zipball_url, str):
            raise TypeError("Expected argument 'zipball_url' to be a str")
        pulumi.set(__self__, "zipball_url", zipball_url)

    @property
    @pulumi.getter(name="assertsUrl")
    def asserts_url(self) -> str:
        return pulumi.get(self, "asserts_url")

    @property
    @pulumi.getter
    def assets(self) -> Sequence['outputs.GetReleaseAssetResult']:
        return pulumi.get(self, "assets")

    @property
    @pulumi.getter
    def body(self) -> str:
        return pulumi.get(self, "body")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def draft(self) -> bool:
        return pulumi.get(self, "draft")

    @property
    @pulumi.getter(name="htmlUrl")
    def html_url(self) -> str:
        return pulumi.get(self, "html_url")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def owner(self) -> str:
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter
    def prerelease(self) -> bool:
        return pulumi.get(self, "prerelease")

    @property
    @pulumi.getter(name="publishedAt")
    def published_at(self) -> str:
        return pulumi.get(self, "published_at")

    @property
    @pulumi.getter(name="releaseId")
    def release_id(self) -> Optional[int]:
        return pulumi.get(self, "release_id")

    @property
    @pulumi.getter(name="releaseTag")
    def release_tag(self) -> Optional[str]:
        return pulumi.get(self, "release_tag")

    @property
    @pulumi.getter
    def repository(self) -> str:
        return pulumi.get(self, "repository")

    @property
    @pulumi.getter(name="retrieveBy")
    def retrieve_by(self) -> str:
        return pulumi.get(self, "retrieve_by")

    @property
    @pulumi.getter(name="tarballUrl")
    def tarball_url(self) -> str:
        return pulumi.get(self, "tarball_url")

    @property
    @pulumi.getter(name="targetCommitish")
    def target_commitish(self) -> str:
        return pulumi.get(self, "target_commitish")

    @property
    @pulumi.getter(name="uploadUrl")
    def upload_url(self) -> str:
        return pulumi.get(self, "upload_url")

    @property
    @pulumi.getter
    def url(self) -> str:
        return pulumi.get(self, "url")

    @property
    @pulumi.getter(name="zipballUrl")
    def zipball_url(self) -> str:
        return pulumi.get(self, "zipball_url")


class AwaitableGetReleaseResult(GetReleaseResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetReleaseResult(
            asserts_url=self.asserts_url,
            assets=self.assets,
            body=self.body,
            created_at=self.created_at,
            draft=self.draft,
            html_url=self.html_url,
            id=self.id,
            name=self.name,
            owner=self.owner,
            prerelease=self.prerelease,
            published_at=self.published_at,
            release_id=self.release_id,
            release_tag=self.release_tag,
            repository=self.repository,
            retrieve_by=self.retrieve_by,
            tarball_url=self.tarball_url,
            target_commitish=self.target_commitish,
            upload_url=self.upload_url,
            url=self.url,
            zipball_url=self.zipball_url)


def get_release(owner: Optional[str] = None,
                release_id: Optional[int] = None,
                release_tag: Optional[str] = None,
                repository: Optional[str] = None,
                retrieve_by: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetReleaseResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['owner'] = owner
    __args__['releaseId'] = release_id
    __args__['releaseTag'] = release_tag
    __args__['repository'] = repository
    __args__['retrieveBy'] = retrieve_by
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('github:index/getRelease:getRelease', __args__, opts=opts, typ=GetReleaseResult).value

    return AwaitableGetReleaseResult(
        asserts_url=__ret__.asserts_url,
        assets=__ret__.assets,
        body=__ret__.body,
        created_at=__ret__.created_at,
        draft=__ret__.draft,
        html_url=__ret__.html_url,
        id=__ret__.id,
        name=__ret__.name,
        owner=__ret__.owner,
        prerelease=__ret__.prerelease,
        published_at=__ret__.published_at,
        release_id=__ret__.release_id,
        release_tag=__ret__.release_tag,
        repository=__ret__.repository,
        retrieve_by=__ret__.retrieve_by,
        tarball_url=__ret__.tarball_url,
        target_commitish=__ret__.target_commitish,
        upload_url=__ret__.upload_url,
        url=__ret__.url,
        zipball_url=__ret__.zipball_url)


@_utilities.lift_output_func(get_release)
def get_release_output(owner: Optional[pulumi.Input[str]] = None,
                       release_id: Optional[pulumi.Input[Optional[int]]] = None,
                       release_tag: Optional[pulumi.Input[Optional[str]]] = None,
                       repository: Optional[pulumi.Input[str]] = None,
                       retrieve_by: Optional[pulumi.Input[str]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetReleaseResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...

# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = ['RepositoryMilestone']


class RepositoryMilestone(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 due_date: Optional[pulumi.Input[str]] = None,
                 owner: Optional[pulumi.Input[str]] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 title: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a GitHub repository milestone resource.

        This resource allows you to create and manage milestones for a Github Repository within an organization or user account.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_github as github

        # Create a milestone for a repository
        example = github.RepositoryMilestone("example",
            owner="example-owner",
            repository="example-repository",
            title="v1.1.0")
        ```

        ## Import

        A GitHub Repository Milestone can be imported using an ID made up of `owner/repository/number`, e.g.

        ```sh
         $ pulumi import github:index/repositoryMilestone:RepositoryMilestone example example-owner/example-repository/1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description of the milestone.
        :param pulumi.Input[str] due_date: The milestone due date. In `yyyy-mm-dd` format.
        :param pulumi.Input[str] owner: The owner of the Github Repository.
        :param pulumi.Input[str] repository: The name of the Github Repository.
        :param pulumi.Input[str] state: The state of the milestone. Either `open` or `closed`. Default: `open`
        :param pulumi.Input[str] title: The title of the milestone.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['description'] = description
            __props__['due_date'] = due_date
            if owner is None:
                raise TypeError("Missing required property 'owner'")
            __props__['owner'] = owner
            if repository is None:
                raise TypeError("Missing required property 'repository'")
            __props__['repository'] = repository
            __props__['state'] = state
            if title is None:
                raise TypeError("Missing required property 'title'")
            __props__['title'] = title
            __props__['number'] = None
        super(RepositoryMilestone, __self__).__init__(
            'github:index/repositoryMilestone:RepositoryMilestone',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            due_date: Optional[pulumi.Input[str]] = None,
            number: Optional[pulumi.Input[int]] = None,
            owner: Optional[pulumi.Input[str]] = None,
            repository: Optional[pulumi.Input[str]] = None,
            state: Optional[pulumi.Input[str]] = None,
            title: Optional[pulumi.Input[str]] = None) -> 'RepositoryMilestone':
        """
        Get an existing RepositoryMilestone resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description of the milestone.
        :param pulumi.Input[str] due_date: The milestone due date. In `yyyy-mm-dd` format.
        :param pulumi.Input[int] number: The number of the milestone.
        :param pulumi.Input[str] owner: The owner of the Github Repository.
        :param pulumi.Input[str] repository: The name of the Github Repository.
        :param pulumi.Input[str] state: The state of the milestone. Either `open` or `closed`. Default: `open`
        :param pulumi.Input[str] title: The title of the milestone.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["due_date"] = due_date
        __props__["number"] = number
        __props__["owner"] = owner
        __props__["repository"] = repository
        __props__["state"] = state
        __props__["title"] = title
        return RepositoryMilestone(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description of the milestone.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="dueDate")
    def due_date(self) -> pulumi.Output[Optional[str]]:
        """
        The milestone due date. In `yyyy-mm-dd` format.
        """
        return pulumi.get(self, "due_date")

    @property
    @pulumi.getter
    def number(self) -> pulumi.Output[int]:
        """
        The number of the milestone.
        """
        return pulumi.get(self, "number")

    @property
    @pulumi.getter
    def owner(self) -> pulumi.Output[str]:
        """
        The owner of the Github Repository.
        """
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter
    def repository(self) -> pulumi.Output[str]:
        """
        The name of the Github Repository.
        """
        return pulumi.get(self, "repository")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[Optional[str]]:
        """
        The state of the milestone. Either `open` or `closed`. Default: `open`
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def title(self) -> pulumi.Output[str]:
        """
        The title of the milestone.
        """
        return pulumi.get(self, "title")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


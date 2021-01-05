# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['BranchProtection']


class BranchProtection(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enforce_admins: Optional[pulumi.Input[bool]] = None,
                 pattern: Optional[pulumi.Input[str]] = None,
                 push_restrictions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 repository_id: Optional[pulumi.Input[str]] = None,
                 require_signed_commits: Optional[pulumi.Input[bool]] = None,
                 required_pull_request_reviews: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BranchProtectionRequiredPullRequestReviewArgs']]]]] = None,
                 required_status_checks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BranchProtectionRequiredStatusCheckArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Protects a GitHub branch.

        This resource allows you to configure branch protection for repositories in your organization. When applied, the branch will be protected from forced pushes and deletion. Additional constraints, such as required status checks or restrictions on users, teams, and apps, can also be configured.

        ## Import

        GitHub Branch Protection can be imported using an ID made up of `repository:pattern`, e.g.

        ```sh
         $ pulumi import github:index/branchProtection:BranchProtection terraform terraform:main
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enforce_admins: Boolean, setting this to `true` enforces status checks for repository administrators.
        :param pulumi.Input[str] pattern: Identifies the protection rule pattern.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] push_restrictions: The list of actor IDs that may push to the branch.
        :param pulumi.Input[str] repository_id: The name or node ID of the repository associated with this branch protection rule.
        :param pulumi.Input[bool] require_signed_commits: Boolean, setting this to `true` requires all commits to be signed with GPG.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BranchProtectionRequiredPullRequestReviewArgs']]]] required_pull_request_reviews: Enforce restrictions for pull request reviews. See Required Pull Request Reviews below for details.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BranchProtectionRequiredStatusCheckArgs']]]] required_status_checks: Enforce restrictions for required status checks. See Required Status Checks below for details.
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

            __props__['enforce_admins'] = enforce_admins
            if pattern is None:
                raise TypeError("Missing required property 'pattern'")
            __props__['pattern'] = pattern
            __props__['push_restrictions'] = push_restrictions
            if repository_id is None:
                raise TypeError("Missing required property 'repository_id'")
            __props__['repository_id'] = repository_id
            __props__['require_signed_commits'] = require_signed_commits
            __props__['required_pull_request_reviews'] = required_pull_request_reviews
            __props__['required_status_checks'] = required_status_checks
        super(BranchProtection, __self__).__init__(
            'github:index/branchProtection:BranchProtection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            enforce_admins: Optional[pulumi.Input[bool]] = None,
            pattern: Optional[pulumi.Input[str]] = None,
            push_restrictions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            repository_id: Optional[pulumi.Input[str]] = None,
            require_signed_commits: Optional[pulumi.Input[bool]] = None,
            required_pull_request_reviews: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BranchProtectionRequiredPullRequestReviewArgs']]]]] = None,
            required_status_checks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BranchProtectionRequiredStatusCheckArgs']]]]] = None) -> 'BranchProtection':
        """
        Get an existing BranchProtection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enforce_admins: Boolean, setting this to `true` enforces status checks for repository administrators.
        :param pulumi.Input[str] pattern: Identifies the protection rule pattern.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] push_restrictions: The list of actor IDs that may push to the branch.
        :param pulumi.Input[str] repository_id: The name or node ID of the repository associated with this branch protection rule.
        :param pulumi.Input[bool] require_signed_commits: Boolean, setting this to `true` requires all commits to be signed with GPG.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BranchProtectionRequiredPullRequestReviewArgs']]]] required_pull_request_reviews: Enforce restrictions for pull request reviews. See Required Pull Request Reviews below for details.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BranchProtectionRequiredStatusCheckArgs']]]] required_status_checks: Enforce restrictions for required status checks. See Required Status Checks below for details.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["enforce_admins"] = enforce_admins
        __props__["pattern"] = pattern
        __props__["push_restrictions"] = push_restrictions
        __props__["repository_id"] = repository_id
        __props__["require_signed_commits"] = require_signed_commits
        __props__["required_pull_request_reviews"] = required_pull_request_reviews
        __props__["required_status_checks"] = required_status_checks
        return BranchProtection(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="enforceAdmins")
    def enforce_admins(self) -> pulumi.Output[Optional[bool]]:
        """
        Boolean, setting this to `true` enforces status checks for repository administrators.
        """
        return pulumi.get(self, "enforce_admins")

    @property
    @pulumi.getter
    def pattern(self) -> pulumi.Output[str]:
        """
        Identifies the protection rule pattern.
        """
        return pulumi.get(self, "pattern")

    @property
    @pulumi.getter(name="pushRestrictions")
    def push_restrictions(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The list of actor IDs that may push to the branch.
        """
        return pulumi.get(self, "push_restrictions")

    @property
    @pulumi.getter(name="repositoryId")
    def repository_id(self) -> pulumi.Output[str]:
        """
        The name or node ID of the repository associated with this branch protection rule.
        """
        return pulumi.get(self, "repository_id")

    @property
    @pulumi.getter(name="requireSignedCommits")
    def require_signed_commits(self) -> pulumi.Output[Optional[bool]]:
        """
        Boolean, setting this to `true` requires all commits to be signed with GPG.
        """
        return pulumi.get(self, "require_signed_commits")

    @property
    @pulumi.getter(name="requiredPullRequestReviews")
    def required_pull_request_reviews(self) -> pulumi.Output[Optional[Sequence['outputs.BranchProtectionRequiredPullRequestReview']]]:
        """
        Enforce restrictions for pull request reviews. See Required Pull Request Reviews below for details.
        """
        return pulumi.get(self, "required_pull_request_reviews")

    @property
    @pulumi.getter(name="requiredStatusChecks")
    def required_status_checks(self) -> pulumi.Output[Optional[Sequence['outputs.BranchProtectionRequiredStatusCheck']]]:
        """
        Enforce restrictions for required status checks. See Required Status Checks below for details.
        """
        return pulumi.get(self, "required_status_checks")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


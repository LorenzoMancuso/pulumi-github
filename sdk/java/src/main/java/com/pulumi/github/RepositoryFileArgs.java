// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RepositoryFileArgs extends com.pulumi.resources.ResourceArgs {

    public static final RepositoryFileArgs Empty = new RepositoryFileArgs();

    /**
     * The branch name, defaults to &#34;main&#34;
     * 
     */
    @Import(name="branch")
    private @Nullable Output<String> branch;

    /**
     * @return The branch name, defaults to &#34;main&#34;
     * 
     */
    public Optional<Output<String>> branch() {
        return Optional.ofNullable(this.branch);
    }

    /**
     * The commit author name, defaults to the authenticated user&#39;s name. GitHub app users may omit author and email
     * information so GitHub can verify commits as the GitHub App.
     * 
     */
    @Import(name="commitAuthor")
    private @Nullable Output<String> commitAuthor;

    /**
     * @return The commit author name, defaults to the authenticated user&#39;s name. GitHub app users may omit author and email
     * information so GitHub can verify commits as the GitHub App.
     * 
     */
    public Optional<Output<String>> commitAuthor() {
        return Optional.ofNullable(this.commitAuthor);
    }

    /**
     * The commit author email address, defaults to the authenticated user&#39;s email address. GitHub app users may omit author
     * and email information so GitHub can verify commits as the GitHub App.
     * 
     */
    @Import(name="commitEmail")
    private @Nullable Output<String> commitEmail;

    /**
     * @return The commit author email address, defaults to the authenticated user&#39;s email address. GitHub app users may omit author
     * and email information so GitHub can verify commits as the GitHub App.
     * 
     */
    public Optional<Output<String>> commitEmail() {
        return Optional.ofNullable(this.commitEmail);
    }

    /**
     * The commit message when creating, updating or deleting the file
     * 
     */
    @Import(name="commitMessage")
    private @Nullable Output<String> commitMessage;

    /**
     * @return The commit message when creating, updating or deleting the file
     * 
     */
    public Optional<Output<String>> commitMessage() {
        return Optional.ofNullable(this.commitMessage);
    }

    /**
     * The file&#39;s content
     * 
     */
    @Import(name="content", required=true)
    private Output<String> content;

    /**
     * @return The file&#39;s content
     * 
     */
    public Output<String> content() {
        return this.content;
    }

    /**
     * The file path to manage
     * 
     */
    @Import(name="file", required=true)
    private Output<String> file;

    /**
     * @return The file path to manage
     * 
     */
    public Output<String> file() {
        return this.file;
    }

    /**
     * Enable overwriting existing files, defaults to &#34;false&#34;
     * 
     */
    @Import(name="overwriteOnCreate")
    private @Nullable Output<Boolean> overwriteOnCreate;

    /**
     * @return Enable overwriting existing files, defaults to &#34;false&#34;
     * 
     */
    public Optional<Output<Boolean>> overwriteOnCreate() {
        return Optional.ofNullable(this.overwriteOnCreate);
    }

    /**
     * The repository name
     * 
     */
    @Import(name="repository", required=true)
    private Output<String> repository;

    /**
     * @return The repository name
     * 
     */
    public Output<String> repository() {
        return this.repository;
    }

    private RepositoryFileArgs() {}

    private RepositoryFileArgs(RepositoryFileArgs $) {
        this.branch = $.branch;
        this.commitAuthor = $.commitAuthor;
        this.commitEmail = $.commitEmail;
        this.commitMessage = $.commitMessage;
        this.content = $.content;
        this.file = $.file;
        this.overwriteOnCreate = $.overwriteOnCreate;
        this.repository = $.repository;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RepositoryFileArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RepositoryFileArgs $;

        public Builder() {
            $ = new RepositoryFileArgs();
        }

        public Builder(RepositoryFileArgs defaults) {
            $ = new RepositoryFileArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param branch The branch name, defaults to &#34;main&#34;
         * 
         * @return builder
         * 
         */
        public Builder branch(@Nullable Output<String> branch) {
            $.branch = branch;
            return this;
        }

        /**
         * @param branch The branch name, defaults to &#34;main&#34;
         * 
         * @return builder
         * 
         */
        public Builder branch(String branch) {
            return branch(Output.of(branch));
        }

        /**
         * @param commitAuthor The commit author name, defaults to the authenticated user&#39;s name. GitHub app users may omit author and email
         * information so GitHub can verify commits as the GitHub App.
         * 
         * @return builder
         * 
         */
        public Builder commitAuthor(@Nullable Output<String> commitAuthor) {
            $.commitAuthor = commitAuthor;
            return this;
        }

        /**
         * @param commitAuthor The commit author name, defaults to the authenticated user&#39;s name. GitHub app users may omit author and email
         * information so GitHub can verify commits as the GitHub App.
         * 
         * @return builder
         * 
         */
        public Builder commitAuthor(String commitAuthor) {
            return commitAuthor(Output.of(commitAuthor));
        }

        /**
         * @param commitEmail The commit author email address, defaults to the authenticated user&#39;s email address. GitHub app users may omit author
         * and email information so GitHub can verify commits as the GitHub App.
         * 
         * @return builder
         * 
         */
        public Builder commitEmail(@Nullable Output<String> commitEmail) {
            $.commitEmail = commitEmail;
            return this;
        }

        /**
         * @param commitEmail The commit author email address, defaults to the authenticated user&#39;s email address. GitHub app users may omit author
         * and email information so GitHub can verify commits as the GitHub App.
         * 
         * @return builder
         * 
         */
        public Builder commitEmail(String commitEmail) {
            return commitEmail(Output.of(commitEmail));
        }

        /**
         * @param commitMessage The commit message when creating, updating or deleting the file
         * 
         * @return builder
         * 
         */
        public Builder commitMessage(@Nullable Output<String> commitMessage) {
            $.commitMessage = commitMessage;
            return this;
        }

        /**
         * @param commitMessage The commit message when creating, updating or deleting the file
         * 
         * @return builder
         * 
         */
        public Builder commitMessage(String commitMessage) {
            return commitMessage(Output.of(commitMessage));
        }

        /**
         * @param content The file&#39;s content
         * 
         * @return builder
         * 
         */
        public Builder content(Output<String> content) {
            $.content = content;
            return this;
        }

        /**
         * @param content The file&#39;s content
         * 
         * @return builder
         * 
         */
        public Builder content(String content) {
            return content(Output.of(content));
        }

        /**
         * @param file The file path to manage
         * 
         * @return builder
         * 
         */
        public Builder file(Output<String> file) {
            $.file = file;
            return this;
        }

        /**
         * @param file The file path to manage
         * 
         * @return builder
         * 
         */
        public Builder file(String file) {
            return file(Output.of(file));
        }

        /**
         * @param overwriteOnCreate Enable overwriting existing files, defaults to &#34;false&#34;
         * 
         * @return builder
         * 
         */
        public Builder overwriteOnCreate(@Nullable Output<Boolean> overwriteOnCreate) {
            $.overwriteOnCreate = overwriteOnCreate;
            return this;
        }

        /**
         * @param overwriteOnCreate Enable overwriting existing files, defaults to &#34;false&#34;
         * 
         * @return builder
         * 
         */
        public Builder overwriteOnCreate(Boolean overwriteOnCreate) {
            return overwriteOnCreate(Output.of(overwriteOnCreate));
        }

        /**
         * @param repository The repository name
         * 
         * @return builder
         * 
         */
        public Builder repository(Output<String> repository) {
            $.repository = repository;
            return this;
        }

        /**
         * @param repository The repository name
         * 
         * @return builder
         * 
         */
        public Builder repository(String repository) {
            return repository(Output.of(repository));
        }

        public RepositoryFileArgs build() {
            $.content = Objects.requireNonNull($.content, "expected parameter 'content' to be non-null");
            $.file = Objects.requireNonNull($.file, "expected parameter 'file' to be non-null");
            $.repository = Objects.requireNonNull($.repository, "expected parameter 'repository' to be non-null");
            return $;
        }
    }

}

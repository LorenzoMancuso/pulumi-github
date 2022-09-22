// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RepositoryMilestoneArgs extends com.pulumi.resources.ResourceArgs {

    public static final RepositoryMilestoneArgs Empty = new RepositoryMilestoneArgs();

    @Import(name="description")
    private @Nullable Output<String> description;

    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * in yyyy-mm-dd format
     * 
     */
    @Import(name="dueDate")
    private @Nullable Output<String> dueDate;

    /**
     * @return in yyyy-mm-dd format
     * 
     */
    public Optional<Output<String>> dueDate() {
        return Optional.ofNullable(this.dueDate);
    }

    @Import(name="owner", required=true)
    private Output<String> owner;

    public Output<String> owner() {
        return this.owner;
    }

    @Import(name="repository", required=true)
    private Output<String> repository;

    public Output<String> repository() {
        return this.repository;
    }

    @Import(name="state")
    private @Nullable Output<String> state;

    public Optional<Output<String>> state() {
        return Optional.ofNullable(this.state);
    }

    @Import(name="title", required=true)
    private Output<String> title;

    public Output<String> title() {
        return this.title;
    }

    private RepositoryMilestoneArgs() {}

    private RepositoryMilestoneArgs(RepositoryMilestoneArgs $) {
        this.description = $.description;
        this.dueDate = $.dueDate;
        this.owner = $.owner;
        this.repository = $.repository;
        this.state = $.state;
        this.title = $.title;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RepositoryMilestoneArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RepositoryMilestoneArgs $;

        public Builder() {
            $ = new RepositoryMilestoneArgs();
        }

        public Builder(RepositoryMilestoneArgs defaults) {
            $ = new RepositoryMilestoneArgs(Objects.requireNonNull(defaults));
        }

        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param dueDate in yyyy-mm-dd format
         * 
         * @return builder
         * 
         */
        public Builder dueDate(@Nullable Output<String> dueDate) {
            $.dueDate = dueDate;
            return this;
        }

        /**
         * @param dueDate in yyyy-mm-dd format
         * 
         * @return builder
         * 
         */
        public Builder dueDate(String dueDate) {
            return dueDate(Output.of(dueDate));
        }

        public Builder owner(Output<String> owner) {
            $.owner = owner;
            return this;
        }

        public Builder owner(String owner) {
            return owner(Output.of(owner));
        }

        public Builder repository(Output<String> repository) {
            $.repository = repository;
            return this;
        }

        public Builder repository(String repository) {
            return repository(Output.of(repository));
        }

        public Builder state(@Nullable Output<String> state) {
            $.state = state;
            return this;
        }

        public Builder state(String state) {
            return state(Output.of(state));
        }

        public Builder title(Output<String> title) {
            $.title = title;
            return this;
        }

        public Builder title(String title) {
            return title(Output.of(title));
        }

        public RepositoryMilestoneArgs build() {
            $.owner = Objects.requireNonNull($.owner, "expected parameter 'owner' to be non-null");
            $.repository = Objects.requireNonNull($.repository, "expected parameter 'repository' to be non-null");
            $.title = Objects.requireNonNull($.title, "expected parameter 'title' to be non-null");
            return $;
        }
    }

}

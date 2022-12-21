// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RepositoryMilestoneState extends com.pulumi.resources.ResourceArgs {

    public static final RepositoryMilestoneState Empty = new RepositoryMilestoneState();

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

    @Import(name="number")
    private @Nullable Output<Integer> number;

    public Optional<Output<Integer>> number() {
        return Optional.ofNullable(this.number);
    }

    @Import(name="owner")
    private @Nullable Output<String> owner;

    public Optional<Output<String>> owner() {
        return Optional.ofNullable(this.owner);
    }

    @Import(name="repository")
    private @Nullable Output<String> repository;

    public Optional<Output<String>> repository() {
        return Optional.ofNullable(this.repository);
    }

    @Import(name="state")
    private @Nullable Output<String> state;

    public Optional<Output<String>> state() {
        return Optional.ofNullable(this.state);
    }

    @Import(name="title")
    private @Nullable Output<String> title;

    public Optional<Output<String>> title() {
        return Optional.ofNullable(this.title);
    }

    private RepositoryMilestoneState() {}

    private RepositoryMilestoneState(RepositoryMilestoneState $) {
        this.description = $.description;
        this.dueDate = $.dueDate;
        this.number = $.number;
        this.owner = $.owner;
        this.repository = $.repository;
        this.state = $.state;
        this.title = $.title;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RepositoryMilestoneState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RepositoryMilestoneState $;

        public Builder() {
            $ = new RepositoryMilestoneState();
        }

        public Builder(RepositoryMilestoneState defaults) {
            $ = new RepositoryMilestoneState(Objects.requireNonNull(defaults));
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

        public Builder number(@Nullable Output<Integer> number) {
            $.number = number;
            return this;
        }

        public Builder number(Integer number) {
            return number(Output.of(number));
        }

        public Builder owner(@Nullable Output<String> owner) {
            $.owner = owner;
            return this;
        }

        public Builder owner(String owner) {
            return owner(Output.of(owner));
        }

        public Builder repository(@Nullable Output<String> repository) {
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

        public Builder title(@Nullable Output<String> title) {
            $.title = title;
            return this;
        }

        public Builder title(String title) {
            return title(Output.of(title));
        }

        public RepositoryMilestoneState build() {
            return $;
        }
    }

}
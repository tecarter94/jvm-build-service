package com.redhat.hacbs.resources.model.v1alpha1;

public class DependencyBuildStatus {

    private String state;

    public String getState() {
        return state;
    }

    public DependencyBuildStatus setState(String state) {
        this.state = state;
        return this;
    }
}
package com.redhat.hacbs.container.deploy.mavenrepository;

import com.amazonaws.services.codeartifact.AWSCodeArtifact;

public record CodeArtifactRepository(AWSCodeArtifact client, String domain, String repository) {
}

#!/usr/bin/env groovy

pipeline {
    agent any

    parameters {
        string(name: 'PROVIDER_VERSION',
            defaultValue: '0.0.1',
            description: 'The version of the terraform provider')
    }

    stages {
        stage("github-release") {
            environment {
                GITHUB_TOKEN = credentials('github-token')
                GPG_PRIVATE_KEY_FILE = credentials('gpg-private-key-file')
                GPG_FINGERPRINT = credentials("gpg-fingerprint")
            }
            steps {
                script {
                    assert params.PROVIDER_VERSION

                    terraformProviderRelease(releaseVersion: params.PROVIDER_VERSION, githubToken: "$GITHUB_TOKEN", 
                        gpgPrivateKeyFile: "$GPG_PRIVATE_KEY_FILE", gpgFingerprint: "$GPG_FINGERPRINT")
                }
            }
        }
    }

    post {
        failure {
            sh "git tag -d v${params.PROVIDER_VERSION}"
        }
    }
}

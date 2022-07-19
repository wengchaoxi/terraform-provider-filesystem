#!/usr/bin/env groovy

def withGo(String version, Closure body) {
    def goreleaserHome = tool name: "goreleaser-v${version}", type: 'com.cloudbees.jenkins.plugins.customtools.CustomTool'
    withEnv(["PATH+GORELEASER=${goreleaserHome}"], body)
}

def terraformProviderRelease(Map params = [:]) {
    def releaseVersion = params.get('releaseVersion', '')
    assert releaseVersion
    def githubToken = params.get('githubToken', '')
    assert githubToken
    def gpgPrivateKeyFile = params.get('gpgPrivateKeyFile', '')
    assert gpgPrivateKeyFile
    def gpgFingerprint = params.get('gpgFingerprint', '')
    assert gpgFingerprint

    withGo('1.18.4') {
        withGoReleaser('1.10.2') {
            sh """
                git tag v"${releaseVersion}"
                err=`gpg --armor --export "${gpgFingerprint}" 1>/dev/null`
                if [ -n "\$err" ]; then
                    gpg --import "${gpgPrivateKeyFile}"
                fi
                export "GPG_FINGERPRINT=${gpgFingerprint}"
                goreleaser release --rm-dist
            """
        }
    }
}

pipeline {
    agent any

    parameters {
        string(name: 'PROVIDER_VERSION',
            defaultValue: '0.0.1',
            description: 'The version of the terraform provider')
    }

    stages {
        stage("github release") {
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
}

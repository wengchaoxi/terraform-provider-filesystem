#!/usr/bin/env groovy

pipeline {
    agent any

    parameters {
        string(name: 'PROVIDER_VERSION',
            defaultValue: '0.0.1',
            description: 'The version of the terraform provider')
    }

    stages {
        stage("Release") {
            steps {
                script {
                    assert params.PROVIDER_VERSION
                    terraformProviderRelease(goVersion: '1.18.4', releaseVersion: params.PROVIDER_VERSION)
                }
            }
        }
    }
}

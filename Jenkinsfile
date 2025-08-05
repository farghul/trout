pipeline {
    agent { label "cactuar && deploy" }
    options {
        buildDiscarder logRotator(
            artifactDaysToKeepStr: "28",
            artifactNumToKeepStr: "5",
            daysToKeepStr: "56",
            numToKeepStr: "10"
        )
    }
    stages {
        stage('Clean WS') {
            steps {
                cleanWs()
            }
        }
        stage("Checkout Trout") {
            steps {
                checkout scmGit(
                    branches: [[name: 'main']],
                    userRemoteConfigs: [[url: 'https://github.com/farghul/trout.git']]
                )
            }
        }
        stage("Build Trout") {
            steps {
                script {
                    sh "/data/apps/go/bin/go build -o /data/automation/bin/trout"
                }
            }
        }
        stage("Checkout DAC") {
            steps {
                checkout scmGit(
                    branches: [[name: 'main']],
                    userRemoteConfigs: [[credentialsId: 'DES-Project', url: 'https://bitbucket.org/bc-gov/desso-automation-conf.git']]
                )
            }
        }
        stage('Run Trout') {
            steps {
                script {
                    sh './scripts/plugin/trout.sh'
                }
            }
        }
    }
}
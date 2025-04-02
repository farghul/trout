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
        stage("Sync") {
            steps {
                lock("satis-rebuild-resource") {
                    dir("/data/automation/github/trout") {
                        sh '''#!/bin/bash
                        source ~/.bashrc
                        git fetch --all
                        git pull
                        '''
                    }
                }
            }
        }
        stage("Build") {
            steps {
                lock("satis-rebuild-resource") {
                    dir("/data/automation/github/trout") {
                        sh "/data/apps/go/bin/go build -o /data/automation/bin/trout ."
                    }
                }
            }
        }
        stage("Run") {
            steps {
                lock("satis-rebuild-resource") {
                    timeout(time: 5, unit: "MINUTES") {
                        retry(2) {
                            sh "/data/automation/scripts/trout.sh ${release}"
                        }
                    }
                }
            }
        }
    }
}
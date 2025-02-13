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
                    dir("/data/scripts/automation/github/trout") {
                        sh "git pull"
                    }
                }
            }
        }
        stage("Build") {
            steps {
                lock("satis-rebuild-resource") {
                    dir("/data/scripts/automation/github/trout") {
                        sh "/data/apps/go/bin/go build -o /data/scripts/automation/programs/trout ."
                    }
                }
            }
        }
        stage("Run") {
            steps {
                lock("satis-rebuild-resource") {
                    timeout(time: 5, unit: "MINUTES") {
                        retry(2) {
                            sh "/data/scripts/automation/scripts/run_trout.sh ${release}"
                        }
                    }
                }
            }
        }
    }
}
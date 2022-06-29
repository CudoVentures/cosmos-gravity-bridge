pipeline {
    agent any
    tools {
        go 'Go 1.18.3'
    }
    environment {
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }
    stages {
        stage("build") {
            steps {
                dir('module'){
                    echo 'BUILD EXECUTION STARTED'
                    sh 'go version'
                    echo '$GOPATH'
                    sh 'make'         
                }
            }
        }
        stage("unit-test") {
            steps {
                dir('module'){
                    echo 'UNIT TEST EXECUTION STARTED'
                    sh 'make unit-tests'
                }
            } 
        }
        stage("functional-test") {
            steps {
                dir('module'){
                    echo 'FUNCTIONAL TEST EXECUTION STARTED'
                }
            }
        }
    }
}
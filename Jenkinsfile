pipeline {
    agent any
    tools {
        go 'Go 1.18.3'
    }
    stages {
        stage("build") {
            steps {
                dir('module'){
                    echo 'BUILD EXECUTION STARTED'
                    echo "WORKSPACE is ${WORKSPACE}"
                    sh 'printenv'
                    sh 'make'   
                }
            }
        }
        stage("unit-test") {
            steps {
                dir('module'){
                    echo 'UNIT TEST EXECUTION STARTED'
                    sh 'make test'
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
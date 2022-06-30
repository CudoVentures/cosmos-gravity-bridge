pipeline {
    agent any
    tools {
        go 'Go 1.18.3'
        nodejs 'NodeJs 16.15.1'
    }
    stages {
        // stage("build") {
        //     steps {
        //         dir('module'){
        //             echo 'BUILD EXECUTION STARTED'
        //             echo "WORKSPACE is ${WORKSPACE}"
        //             sh 'printenv'
        //             sh 'make'   
        //         }
        //     }
        // }
        // stage("unit-test") {
        //     steps {
        //         dir('module'){
        //             echo 'UNIT TEST EXECUTION STARTED'
        //             sh 'make test'
        //         }
        //     } 
        // }
        stage("solidity-test") {
            steps {
                dir('solidity'){
                    echo 'SOLIDITY TEST EXECUTION STARTED'
                    sh 'npm ci'
                    sh 'npm hardhat test'
                }
            }
        }
    }
}
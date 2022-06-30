pipeline {
    agent any
    tools {
        go 'Go 1.18.3'
        nodejs 'NodeJs 16.15.1'
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
        stage("solidity-test") {
            steps {
                dir('solidity'){
                    echo 'SOLIDITY TEST EXECUTION STARTED'
                    sh 'npm install'
                    sh 'npx hardhat typechain'
                    sh 'npx hardhat test'
                }
            }
        }
        stage('Rust test') {
            agent {
                docker {
                    image 'rust:latest'
                    reuseNode true
                }
            }
            steps {
                  dir('orchestrator'){
                    echo 'RUST TEST EXECUTION STARTED'
                    sh 'cargo check --all --verbose'
                    sh 'cargo test --all --verbose'
                    sh 'cargo fmt --all -- --check'
                    sh 'cargo clippy --all --all-targets --all-features -- -D warnings'
                }
            }
        }
        // stage('Store to GCS') { // not needed yet
        //     steps{
        //         sh '''
        //             env > build_environment.txt
        //         '''
        //         // If we name pattern build_environment.txt, this will upload the local file to our GCS bucket.
        //         step([$class: 'ClassicUploadStep', credentialsId: env
        //                 .CREDENTIALS_ID,  bucket: "gs://${env.BUCKET}",
        //                 pattern: env.PATTERN])
        //     }
        // }
    }
}
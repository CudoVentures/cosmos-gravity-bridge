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
                    // sh 'make'   
                }
            }
        }
        // stage("unit-test") {
        //     steps {
        //         dir('module'){
        //             echo 'UNIT TEST EXECUTION STARTED'
        //             sh 'make test'
        //         }
        //     } 
        // }
        // stage("solidity-test") {
        //     steps {
        //         dir('solidity'){
        //             echo 'SOLIDITY TEST EXECUTION STARTED'
        //             sh 'npm install'
        //             sh 'npx hardhat typechain'
        //             sh 'npx hardhat test'
        //         }
        //     }
        // }
        // stage('Rust test') {
        //     agent {
        //         docker {
        //             image 'rust:latest'
        //             reuseNode true
        //         }
        //     }
        //     steps {
        //           dir('orchestrator'){
        //             echo 'RUST TEST EXECUTION STARTED'
        //             sh 'rustup component add rustfmt'
        //             // sh 'rustup component add clippy'
        //             sh 'cargo check --all --verbose'
        //             sh 'cargo test --verbose'
        //             sh 'cargo fmt --all -- --check'
        //             // sh 'cargo clippy --all --all-targets --all-features -- -D warnings'  - clippy checks fails currently
        //         }
        //     }
        // }
        // stage('Store to GCS') { // not needed yet as we are not using the artifacts
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
    post {
        always {
            script {
                def userIds = slackUserIdsFromCommitters()
                echo "userIds is ${userIds}"
                def userIdsString = userIds.collect { "<@$it>" }.join(' ')
                echo "userIdsString is ${userIdsString}"
                slackSend(color: "good", message: "$userIdsString Message from Jenkins Pipeline")
            }
            cleanWs()
        }
    }
}


def color_slack_msg() {
    switch(currentBuild.currentResult) {
    case "SUCCESS":
        return "good"
        break
    case "FAILURE":
    case "UNSTABLE":
        return "danger"
        break
    default:
        return "warning"
        break
    }
}

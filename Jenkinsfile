pipeline {
    agent any
    stages {
        stage('Build') {
            agent {
                docker { image 'golang:1.18-bullseye' }
            }
            steps {
                sh 'go --version'
                echo '$GOPATH'
                dir('module') {
                     sh 'make'
                }
                          }
        }
        stage('Test') {
            agent {
                docker { image 'golang:1.18-bullseye' }
            }
            steps {
                sh 'go --version'
                echo '$GOPATH'
                dir('module') {
                     sh 'make test'
                }
                          }
        }
        stage('Run Solidity NodeJS tests') {
            agent {
                docker { image 'node:16.13.1-alpine' }
            }
            steps {
                sh 'node --version'
            }
        }
    }
}
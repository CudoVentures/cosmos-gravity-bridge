pipeline {
    agent none

    stages {
        stage('Build') {
            agent { label 'jenkins-node' }
            steps {
                echo 'Building..'
                sh '''
                '''
            }
        }
    }

    post {
        success {
            echo 'This will run only if successful'
        }
    }
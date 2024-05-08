pipeline {
    agent any

    stages {
        stage('Build binary') {
            agent {
                docker {
                    image "golang:1.22-bookworm"
                    reuseNode true
                }
            }
            steps {
                sh 'go build'
            }
        }
        stage("Deploy") {
            steps {
                sh "docker compose up --build -d"
            }
        }
    }
}

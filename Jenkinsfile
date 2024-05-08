pipeline {
    agent any

    stages {
        stage("Run & build") {
            steps {
                sh "docker compose up --build -d"
            }
        }
    }
}

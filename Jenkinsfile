
pipeline {
    agent { docker { image 'golang' } }
      environment {
    DOCKER_HOST = 'unix:///run/user/996/docker.sock'
  }

    stages {
        stage('build') {
            steps {
                sh 'go version'
            }
        }
    }
}

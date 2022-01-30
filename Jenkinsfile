
pipeline {
    agent { docker { image 'golang' } }
    

    stages {
        stage('test') {
            steps {
                sh 'make test'
            }
        }
        stage('build') {
            steps {
                sh 'make build'
            }
        }
        
    }
}

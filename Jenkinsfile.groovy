pipeline {
    agent {
        node {
            label 'bellex'
        }
    }
    options {
        disableConcurrentBuilds()
    }
    
    stages {
        stage('Pull Code') {
            steps {
                dir('C:/Users/Administrator/go/src/github.com/nomango/bellex') {
                    git url: 'git@github.com:nomango/bellex.git', branch: 'master', credentialsId: '53d39ac3-419f-48ce-ae60-f2cfcf32902d'
                }
            }
        }
        // stage('Build Bellex Back-end') {
        //     steps {
        //         dir('C:/Users/Administrator/go/src/github.com/nomango/bellex/server') {
        //             bat 'go build -o bellex main.go'
        //         }
        //     }
        // }
    }
    post {
        success {
            emailext body:
                '''${SCRIPT, template="build.template"}''',
                subject: "Job '${env.JOB_NAME}' (${env.BUILD_NUMBER}) is success :)",
                recipientProviders: [[$class: 'DevelopersRecipientProvider'], [$class: 'RequesterRecipientProvider']],
                to: "569629550@qq.com"
        }
        failure {
            emailext body:
                '''${SCRIPT, template="build.template"}''',
                subject: "Job '${env.JOB_NAME}' (${env.BUILD_NUMBER}) is failed :(",
                recipientProviders: [[$class: 'DevelopersRecipientProvider'], [$class: 'RequesterRecipientProvider']],
                to: "569629550@qq.com"
        }
    }
}
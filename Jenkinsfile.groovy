pipeline {
    agent any

    options {
        disableConcurrentBuilds()
    }
    
    stages {
        stage('Pull Code') {
            steps {
                dir('C:/Users/Administrator/go/src/github.com/nomango/bellex') {
                    git url: 'https://github.com/nomango/bellex/', branch: 'master'//, credentialsId: '818c7733-95ad-441a-9bb3-fee4accc6acc'
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
    // post {
    //     success {
    //         emailext body:
    //             '''${SCRIPT, template="build.template"}''',
    //             subject: "Job '${env.JOB_NAME}' (${env.BUILD_NUMBER}) is success :)",
    //             recipientProviders: [[$class: 'DevelopersRecipientProvider'], [$class: 'RequesterRecipientProvider']],
    //             to: "569629550@qq.com"
    //     }
    //     failure {
    //         emailext body:
    //             '''${SCRIPT, template="build.template"}''',
    //             subject: "Job '${env.JOB_NAME}' (${env.BUILD_NUMBER}) is failed :(",
    //             recipientProviders: [[$class: 'DevelopersRecipientProvider'], [$class: 'RequesterRecipientProvider']],
    //             to: "569629550@qq.com"
    //     }
    // }
}
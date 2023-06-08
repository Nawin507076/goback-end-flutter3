// Final to Use
pipeline {
    agent any
    environment {
        IMG_NAME = "507076/nawin-api"
        WORK_DIR = "nawin/nawin-backend/"
        PROD_SERVER = "ubuntu@13.229.66.4"
    }
    stages {

        stage('Build Img in Jenkins Serv') {
            steps {
                sh """
                  docker rmi -f $IMG_NAME
                  docker build --rm -t $IMG_NAME .
                """
            }
        }

        stage('Push Img to DockerHub') {
            steps {
                withCredentials([usernamePassword(credentialsId: 'dockerhub-nawin', passwordVariable: 'DOCKER_PASSWORD', usernameVariable: 'DOCKER_USERNAME')]) {
                sh """
                    docker login -u $DOCKER_USERNAME -p $DOCKER_PASSWORD
                    docker push $IMG_NAME
                """
                }
            }
        }

        stage('Deploy to Production') {
            steps {
                sshagent(credentials: ["nad-DEV-Server"]) {
                    sh """
                        ssh -o StrictHostKeyChecking=no $PROD_SERVER \"
                            mkdir -p $WORK_DIR
                            cd $WORK_DIR && pwd
                            docker compose down
                            docker rmi -f $IMG_NAME
                        \"
                        scp docker-compose.yml $PROD_SERVER:$WORK_DIR

                        ssh -o StrictHostKeyChecking=no $PROD_SERVER \"
                            cd $WORK_DIR && pwd && ls -la
                            docker compose up -d
                        \"
                    """
                }
            }
        }
    }
}
// sh ''' == string
// sh """ == มีตัวแปรข้างใน
// เครื่องหมาย / หลัง scp คือบอกให้รู้ว้าเป็น floder/ ไม่ใช้ file โดยจะโยนไฟล์ไปหลัง floder นั้น
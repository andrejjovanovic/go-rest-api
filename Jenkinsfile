#!/usr/bin/env groovy

pipeline {
    agent any

    stages{
        stage('Clean workspace') {
            steps{
                deleteDir()
            }
           
        }

        stage('Checkout'){
            steps{
                echo "Checkout"
                checkout scm
            }
        }

        stage('Build'){
            steps {
                echo "Building"
                sh "cat Dockerfile"
            }
        }

        stage('Deploy'){
            steps {
                echo "Deploy"
            } 
        }
    }
}

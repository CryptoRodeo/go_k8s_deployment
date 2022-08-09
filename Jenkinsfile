pipeline {
	agent any

	triggers {
		pollSCM('* * * * *')
	}

	environment = {
		dockerhub=credentials('DockerKey')
	}

	stages {
		stage("Build") {

			steps {
				script {
					env.CI = false //disable unecessary React warnings
				}

				sh "make"
				}
		}

		stage("Test") {
			steps { sh "echo 'WIP :)'" }
		}

		stage("Docker push") {
			//steps { sh "make publish_all" }
			steps {
				sh 'echo $dockerhub_PSW | docker login -u $dockerhub_USR --password-stdin'
				sh "make publish_all"
			}
		}
	}
}
pipeline {
	agent any

	triggers {
		pollSCM('* * * * *')
	}

	environment {
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
				sh "make publish_all"
			}
		}
	}
}
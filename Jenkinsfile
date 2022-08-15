pipeline {
	agent any

	triggers {
    // check repo for changes every minute
		pollSCM('* * * * *')
	}

	environment {
		dockerhub=credentials('DockerKey')
	}

	stages {
		stage("Build") {

			steps {
				script {
					env.CI = false //disable unnecessary React warnings
				}

				sh "make"
				}
		}

		stage("Test") {
      // WIP, still need to add tests.
			steps { sh "echo 'WIP :)'" }
		}

		stage("Docker push") {
			steps {
				sh "echo $dockerhub_PSW | docker login -u $dockerhub_USR --password-stdin"
				sh "make publish_all"
			}
		}
	}
}
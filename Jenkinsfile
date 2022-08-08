pipeline {
	agent any

	triggers {
		pollSCM('* * * * *')
	}

	stages {
		stage("Build") {
			steps { sh "make" }
		}

		stage("Test") {
			steps { sh "echo 'WIP :)'" }
		}

		stage("Docker push") {
			steps { sh "make publish_all" }
		}
	}
}
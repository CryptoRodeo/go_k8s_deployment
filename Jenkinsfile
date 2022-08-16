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

	post {
		always {
			SendEmailNotification(currentBuild.result)
		}
	}
}

def SendEmailNotification(String res) {
	// send the email to the one who triggered the build.
	// creds are the user's jenkins creds
	def to = emailextrecipients([
		requestor()
	])

	def subject = "${env.JOB_NAME} - Build #${env.BUILD_NUMBER} ${result}"
	def content = '${JELLY_SCRIPT,template="html"}'

	// send mail
	if(to != null && !to.isEmpty()) {
        emailext(body: content, mimeType: 'text/html',
        subject: subject,
        to: to, attachLog: true )
	}
}

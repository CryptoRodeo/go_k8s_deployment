const ENDPOINT = "http://localhost:8080/api/v1/users"

/**
 * 
 * @param {*} stateCb state callback
 * @param {*} errorCb  error callback
 */
async function getUsers(stateCb, errorCb) {
	await fetch(ENDPOINT,
		{
			mode: 'cors',
			headers: {
				'Content-Type': 'application/json',
			}
		})
		.then(res => res.json())
		.then((res) => stateCb(res.data), (err) => errorCb(err)) 
}

export default {
	getUsers,
}
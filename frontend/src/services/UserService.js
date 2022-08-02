
/**
 *
 * @param {*} stateCb React hook state callback
 * @param {*} errorCb  React hook error callback
 */
async function getUsers(stateCb, errorCb) {
	const users_endpoint = "http://localhost:8080/api/v1/users"
	await fetch(users_endpoint,
		{
			mode: 'cors',
			headers: {
				'Content-Type': 'application/json',
			}
		})
		.then(res => res.json())
		.then((res) => stateCb(res.data), (err) => errorCb(err))
}

/**
 *
 * @param {*} user_id int
 * @param {*} errorCb React hook error callback
 */
async function getUserTickets(user_id, errorCb) {
	const user_tickets_endpoint = `http://localhost:8080/api/v1/users/${user_id}/tickets`
	let tickets;

	await fetch(user_tickets_endpoint, {
		mode: 'cors',
		headers: {
			'Content-Type': 'application/json',
		}
	})
	.then(res => res.json())
	.then((res) => tickets = res)

	return tickets;
}

export default {
	getUsers,
	getUserTickets,
}
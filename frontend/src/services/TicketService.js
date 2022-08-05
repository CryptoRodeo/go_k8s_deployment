/**
 *
 * @param {*} stateCb React hook state callback
 * @param {*} errorCb  React hook error callback
 */
async function getTickets(stateCb, errorCb) {
	const ticket_endpoint = "http://localhost:8081/api/v1/tickets"
	await fetch(ticket_endpoint, {
		mode: 'cors',
		headers: {
			'Content-Type': 'application/json',
		}
	})
	.then(res => res.json())
	.then((res) => stateCb(res.data), (err) => errorCb(err))
}

const TicketService = {
	getTickets,
}

export default TicketService;

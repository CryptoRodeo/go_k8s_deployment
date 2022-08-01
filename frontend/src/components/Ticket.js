import React, { useState, useEffect } from 'react';
import './Ticket.scss';
export default function Ticket(props) {
	const [ticket, setTicket] = useState(props)

	useEffect(() => {
		setTicket(ticket);
	}, [props]);

	return (
		<div className="section ticket-section has-background-white">
			<div className="columns">
				<div className="column is-three-quarters">
					<h2 className="title is-3">Info:</h2> 
					<p>{ticket.description}</p>
				</div>
			</div>
		</div>
	)
}
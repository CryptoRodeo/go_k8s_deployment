import React, { useState, useEffect } from 'react';
import './Ticket.scss';
export default function Ticket(props) {
	const [ticket, setTicket] = useState(props)

	useEffect(() => {
		setTicket(ticket);
	}, [props]);

	return (
		<div className="card ticket-section has-background-white">
			<div className="card-content">
				<h2 className="title">{ticket.description}</h2>
				<h2 className="subtitle">ID: {ticket.id}</h2>
				<span className={ticket.is_urgent ? 'tag is-danger' : 'tag is-light'}>{ticket.is_urgent ? 'Urgent!' : 'Regular'}</span>
			</div>
		</div>
	)
}
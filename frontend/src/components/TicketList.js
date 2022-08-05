import './TicketList.scss';
import React, { useState, useEffect } from 'react';
import Ticket from './Ticket';
import TicketService from '../services/TicketService';

export default function TicketList(props) {
	const [error, setError] = useState(null)
	const [tickets, setTickets] = useState([]);

	useEffect(() => {
		// If we explicitly pass tickets, just render them
		if (props.tickets) {
			setTickets(props.tickets);
		}
		else {
			// else, use the service.
			TicketService.getTickets(setTickets, setError);
		}
	}, [props.tickets]);


	if (error) {
		console.error(error);
	}

	if (!tickets.length) {
		return (
			<div className="container ticket-list">Loading...</div>
		);
	}

	return (
		<div className="container ticket-list">
				{tickets.map(ticket => (<Ticket key={ticket.id} {...ticket} />))}
		</div>
	)
}
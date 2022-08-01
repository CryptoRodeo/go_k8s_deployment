import './TicketList.scss';
import React, { useState, useEffect } from 'react';
import Ticket from './Ticket';
import TicketService from '../services/TicketService';

export default function TicketList() {
	const [error, setError] = useState(null)
	const [tickets, setTickets] = useState([]);

	useEffect(() => {
		TicketService.getTickets(setTickets, setError)
	}, []);

	if (error) {
		console.error(error);
	}

	if (!tickets) {
		return (
			<div>Loading...</div>
		);
	}

	return (
		<div className="container ticket-list">
				{tickets.map(ticket => (<Ticket key={ticket.id} {...ticket} />))}
		</div>
	)
}
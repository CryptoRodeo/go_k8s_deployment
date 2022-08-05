import "./UserTicketList.scss";
import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import UserService from '../services/UserService';
import TicketList from './TicketList';

export default function UserTicketList() {
	const { id } = useParams(); // :id
	const [error, setError] = useState(null);
	const [userTickets, setUserTickets] = useState(null)
	const [user, setUser] = useState(null)

	useEffect(() => {
		UserService.getUserTickets(id, setUserTickets, setError)
		UserService.getUserByID(id, setUser, setError)
	}, [id]);

	if (error) {
		console.error(error)
	}

	if (!userTickets || !user) {
		return (
			<div className="container user-list">Loading...</div>
		);
	}

	return (
		<div>
			<div className="box has-background-dark user-name-box">
				<h1 className="title is-center has-text-white">Tickets for <span className="has-text-primary">{user.name}</span></h1>
			</div>
			<TicketList tickets={userTickets}/>
		</div>
	)
}
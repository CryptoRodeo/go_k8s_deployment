import React, { useState, useEffect } from 'react';
import "./User.scss";

export default function User(props) {
	const [user, setUser] = useState(props)
	const userTicketLink = `user/${user.id}/tickets`;

	useEffect(() => {
		setUser(user);
	}, [user]);

	return (
		<div className="card ticket-section has-background-white">
			<div className="card-content">
				<h2 className="title">{user.name}</h2>
				<h3 className="subtitle">Role: {user.role}</h3>
				<button className="button is-info">
					<a href={userTicketLink} className="has-text-white">
						View Tickets ({user.tickets.length})
					</a>
				</button>
			</div>

		</div>
	)
}
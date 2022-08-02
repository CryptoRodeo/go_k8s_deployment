import React, { useState, useEffect } from 'react';
import "./User.scss";
export default function User(props) {
	const [user, setUser] = useState(props)

	useEffect(() => {
		setUser(user);
	}, [props]);


	return (
		<div className="section ticket-section">
			<div className="columns">
				<div className="column is-one-quarter">
					<h2 className="title is-3">{user.name}</h2>
					<p className="subtitle is-6">Role: {user.role}</p>
					{/* <button className="button is-primary">View Tickets({user.tickets.length})</button>  */}
				</div>
			</div>
		</div>
	)
}
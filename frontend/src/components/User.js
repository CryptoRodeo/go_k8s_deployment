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
				<div className="column is-three-quarters">
					<h2 className="title is-3">Info:</h2> 
					<p>"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
				</div>
				<div className="column is-one-quarter">
					<h2 className="title is-3">{user.name}</h2> 
					<p className="subtitle is-6">Role: {user.role}</p>
					{/* <button className="button is-primary">View Tickets({user.tickets.length})</button>  */}
				</div>
			</div>
		</div>
	)
}
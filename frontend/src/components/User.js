import React, { useState, useEffect } from 'react';

export default function User(props) {
	const [user, setUser] = useState(props)
	
	useEffect(() => { 
		setUser(user);
	}, [props]);
		

	return (
		<li>
			<p>{user.name} - {user.role}</p>
		</li>
	)
}
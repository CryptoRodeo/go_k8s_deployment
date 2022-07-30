import './UserList.scss';
import React, { useState, useEffect } from 'react';
import User from './User';
import UserService from '../services/UserService';

export default function UserList() {
	const [error, setError] = useState(null);
	const [users, setUsers] = useState([]);

	useEffect(() => {
		UserService.getUsers(setUsers, setError)
	}, []);

	if (error) {
		console.error(error)
	}

	if (!users) {
		return <div>Loading...</div>
	}
	
	return (
		<div className="container">
			<div className="users-list">
				{users.map(user => (<User key={user.id} {...user} />))}
			</div>
		</div>

	)
}
import React from 'react';
import './NavBar.scss';
export default function NavBar(props) {
	return (
		<nav className="navbar" role="navigation" aria-label="main navigation">
			<div className="navbar-brand">
				<a role="button" className="navbar-burger" aria-label="menu" data-target="menu">
					<span aria-hidden="false" className="burger-line"></span>
					<span aria-hidden="false" className="burger-line"></span>
					<span aria-hidden="false" className="burger-line"></span>
   	 		</a>
				<a className="navbar-item" href="#">
					<h1 className="title has-text-primary">TicketMaster</h1>
				</a>
			</div>
		</nav>
	)
}
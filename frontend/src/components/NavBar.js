import React, { useState } from 'react';
import './NavBar.scss';
export default function NavBar(props) {
	let navburger_classes = `navbar-burger ${props.showMenu ? '' : 'is-active'}`;

	return (
		<nav className="navbar" role="navigation" aria-label="main navigation">
			<div className="navbar-brand">
				<a role="button" className={navburger_classes} onClick={() => props.toggleMenu(!props.showMenu)} aria-label="menu" data-target="#menuSection" data-toggle="menu">
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
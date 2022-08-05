import React from 'react';
import './Menu.scss';

export default function Menu(props) {
	return (
	<aside className="menu" id="menuSection">
 		<p className="menu-label">
 		  General
 		</p>
 		<ul className="menu-list">
 		  <li className="menu-item"><a href="/">Tickets</a></li>
 		  <li className="menu-item"><a href="/users">Users</a></li>
 		</ul>
 		<p className="menu-label">
 		  Other
 		</p>
 		<ul className="menu-list">
 		  <li className="menu-item"><a href="/">Source Code</a></li>
 		</ul>
	</aside>
	)
}
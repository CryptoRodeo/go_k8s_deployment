import React, {useState, useEffect} from 'react';
import './Menu.scss';

export default function Menu(props) {
	return (
	<aside className="menu">
 		<p className="menu-label">
 		  General
 		</p>
 		<ul className="menu-list">
 		  <li className="menu-item"><a>Tickets</a></li>
 		  <li className="menu-item"><a>Users</a></li>
 		</ul>
 		<p className="menu-label">
 		  Other
 		</p>
 		<ul className="menu-list">
 		  <li className="menu-item"><a>Source Code</a></li>
 		</ul>
	</aside>
	)
}
import './Main.scss'
import React, { useState, useEffect } from 'react';
import TicketList from './TicketList';
import Menu from './Menu';


export default function Main() {
	return (
		<div className="columns">
			<div className="column menu-column">
        <Menu />
			</div>
			<div className="column data-column">
				<TicketList/>
			</div>
		</div>
	)
}
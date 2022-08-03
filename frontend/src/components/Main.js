import './Main.scss'
import React from 'react';
import TicketList from './TicketList';
import Menu from './Menu';


export default function Main(props) {

	const menu_styles = {
		"display": (props.showMenu ? 'block' : 'none'),
	}
	return (
		<div className="columns">
			<div className="column menu-column" style={menu_styles}>
        <Menu />
			</div>
			<div className="column data-column">
				<TicketList/>
			</div>
		</div>
	)
}
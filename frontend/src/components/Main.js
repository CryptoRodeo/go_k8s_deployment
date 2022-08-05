import './Main.scss'
import React from 'react';
import TicketList from './TicketList';
import UserList from './UserList';
import Menu from './Menu';
import { BrowserRouter, Route, Routes} from 'react-router-dom';

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
				<BrowserRouter>
					<Routes>
						<Route path="/" element={<TicketList />} />
						<Route path="/users" element={<UserList/>} />
					</Routes>
				</BrowserRouter>
			</div>
		</div>
	)
}
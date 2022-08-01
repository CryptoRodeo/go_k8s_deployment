import React from 'react';

export default function NavBar(props) {
	return (
		<nav className="navbar is-black" role="navigation" aria-label="main navigation">
			<div className="navbar-brand">
				<a className="navbar-item" href="#">
					<h1 className="title has-text-primary">TicketMaster</h1>
				</a>
			</div>


			<div id="mainNavbar" className="navbar-menu has-text-white-ter">
				<div className="navbar-start">
					<a className="navbar-item">
						Tickets
					</a>
				</div>

				<div className="navbar-end">
					<div className="navbar-item">
						<div className="buttons">
							<a className="button is-primary">
								<strong>Sign Up</strong>
							</a>
							<a className="button is-light">
								Log in
							</a>
						</div>
					</div>
				</div>
			</div>

		</nav>
	)
}
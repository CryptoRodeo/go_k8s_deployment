import './App.scss';
import NavBar from './components/NavBar';
import Main from './components/Main';
import React, { useState } from 'react';

function App() {
  const [showMenu, toggleMenu] = useState(true);

  return (
    <div className="App container is-fullhd has-background-light">
      <NavBar showMenu={showMenu} toggleMenu={toggleMenu} />
      <Main showMenu={showMenu} />
    </div>
  );
}

export default App;

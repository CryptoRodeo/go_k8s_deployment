import logo from './logo.svg';
import './App.scss';
import TicketList from './components/TicketList';
import Header from './components/Header';
import NavBar from './components/NavBar';

function App() {
  return (
    <div className="App container is-fullhd has-background-light">
    <NavBar/>
    <TicketList/>
    </div>
  );
}

export default App;

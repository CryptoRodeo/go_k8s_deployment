import logo from './logo.svg';
import './App.scss';
import UserList from './components/UserList';
import Header from './components/Header';
import NavBar from './components/NavBar';

function App() {
  return (
    <div className="App container is-fullhd ">
    <NavBar/>
    <UserList />
    </div>
  );
}

export default App;

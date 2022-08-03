import logo from './logo.svg';
import './App.scss';
import NavBar from './components/NavBar';
import Main from './components/Main';

function App() {
  return (
    <div className="App container is-fullhd has-background-light">
      <NavBar/>
      <Main />
    </div>
  );
}

export default App;

import React from 'react';
import logo from '../assets/logo.svg';
import '../assets/App.css';
import swal from "sweetalert";


function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
          <Car />
        </a>
        <TestAlert />
      </header>
    </div>
  );
}

function Car() {
  return (
    <h1>N</h1>
  );
}

class TestAlert extends React.Component {
  constructor(props) {
    super(props)
    this.state = { msg : "Mafia"}
  }

  makeAlert = async () => {
    let result = await swal("Night")
    alert(result);
  }

  render() {
    return (
      <div>
        <button type = "button" onClick={this.makeAlert}>Check Career</button>
      </div>
    );
  }
}


export default App;
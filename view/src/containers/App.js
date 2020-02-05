import React from 'react';
import '../assets/App.css';
import Day from "../containers/Day"
import Night from "../containers/Night"
import Dead from "../containers/Dead"
import Start from "../containers/Start"
import { BrowserRouter as Router, Route } from 'react-router-dom'

function App() {
  return (
    <div className="App">
      <Router>
        <Route exact path="/" component={Start} />
        <Route path="/day" component={Day} />
        <Route path="/night" component={Night} />
        <Route path="/dead" component={Dead} />
      </Router>
    </div>
  );
}

export default App;
import React, { useEffect, useState } from 'react';
import './App.css';
import config from './config';
import logo from './logo.svg';

function App() {
  // proof of concept for talking to the golang api
  const [apiMessage, setApiMessage] = useState('nothing yet!');

  useEffect(() => {
    fetch(config.apiRoot)
      .then((response) => response.text())
      .then((text) => setApiMessage(text));
  }, []);

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>Hello, world!</p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
        <p>{apiMessage}</p>
      </header>
    </div>
  );
}

export default App;

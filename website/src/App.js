import './App.css';
import Button from 'react-bootstrap/Button';
import 'bootstrap/dist/css/bootstrap.min.css';
import React, { useEffect, useState } from 'react';


const API_URL = process.env.REACT_APP_API_URL;

function changeDeskMode(mode) {
  return fetch(API_URL, {
			method: 'POST',
			body: JSON.stringify({
				position: mode
			})
		});
}

function sit() {
  changeDeskMode('sit')
}

function stand() {
  changeDeskMode('stand')
}

function simulateNetworkRequest() {
  return new Promise((resolve) => setTimeout(resolve, 2000));
}

function App() {
  const [isLoading, setLoading] = useState(false);

  useEffect(() => {
    if (isLoading) {
      simulateNetworkRequest().then(() => {
        setLoading(false);
      });
    }
  }, [isLoading]);

  const handleStand = () => {
    setLoading(true);
    stand();
  }

  const handleSit = () => {
    setLoading(true);
    sit();
  }

  return (
    <div className="App">
      <link
        rel="stylesheet"
        href="https://cdn.jsdelivr.net/npm/bootstrap@5.2.0-beta1/dist/css/bootstrap.min.css"
        integrity="sha384-0evHe/X+R7YkIZDRvuzKMRqM+OrBnVFBL6DOitfPri4tjfHxaWutUpFmBp4vmVor"
        crossorigin="anonymous"
      />
      <header className="App-header">
        <Button onClick={!isLoading ? handleStand : null} disabled={isLoading} variant="success" size="lg">
          {isLoading ? 'Loading…' : 'Desk go up'}
        </Button>{' '}
        <Button onClick={!isLoading ? handleSit : null} disabled={isLoading} variant="warning" size="lg">
          {isLoading ? 'Loading…' : 'Desk go down'}
        </Button>
      </header>
    </div>
  );
}

export default App;

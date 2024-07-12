import React, { useEffect, useState } from 'react';
import {Header, GuideChart} from './components';
import './App.css';

function App() {
  const [message, setMessage] = useState('');
  const [frostDate, setfrostDate] = useState('05/01');


  useEffect(() => {
    fetch('/api/hello')
      .then(response => response.json())
      .then(data => setMessage(data.text));
  }, []);

  const sayGoodbye = () => {
    fetch('/api/goodbye')
      .then(response => response.json())
      .then(data => setMessage(data.text));
  }



  const handleChange = (event) => {
    setfrostDate(event.target.value);
  };

  const handleSubmit = (event) => {
    event.preventDefault();
    console.log({ lastFrostDate: frostDate });
  };
  return (
    <div className="appBody">     
        <Header />
        <form onSubmit={handleSubmit}>
          <div>
            <label htmlFor="lastFrostDate">Last Frost Date:</label>
            <input
              id="lastFrostDate"
              value={frostDate}
              onChange={handleChange}
            />
          </div>
          <button type="submit">Submit</button>
        </form>     
        <br></br>
        <GuideChart />
        <h3>Frost Date Selected: {frostDate}</h3>
        <p>message from main.go: {message}</p>
        <button onClick={sayGoodbye}>Say Goodbye</button>
    </div>
  );
}

export default App;


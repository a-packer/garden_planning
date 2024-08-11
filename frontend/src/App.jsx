import React, { useState } from 'react';
import {Header, GuideChart, Login, Register} from './components';
import './App.css';
import PlantList from './components/PlantList';

function App() {
  const [frostDate, setfrostDate] = useState('05/01');


  const handleChange = (event) => {
    setfrostDate(event.target.value);
  };

  const handleSubmit = (event) => {
    event.preventDefault();
    console.log('lastFrostDate:', frostDate );
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
        <Login />  
        <Register />   
        <br></br>
        <GuideChart />
        <PlantList />
    </div>
  );
}

export default App;


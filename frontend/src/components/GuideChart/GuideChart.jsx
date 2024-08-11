import React, {useState} from 'react'
import BarChart from './BarChart';

export const GuideChart = () => {
  const [plants, setPlants] = useState([])
  const displayPlantPlan = async (e) => {
    e.preventDefault();
    try {
      const response = await fetch('/gardens', {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
        }
      });
      if (response.ok) {
        setPlants(await response.json())
      }
    } catch (error) {
      console.log('Error with fetching plants');
    }
  };
  return (
    <>
      <button onClick={displayPlantPlan}>Set Plan for Plants</button>
      <BarChart data={plants} />
    </>
    
  )
}

export default GuideChart;

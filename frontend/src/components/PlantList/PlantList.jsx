import React, {useState} from 'react';

const PlantList = () => {
  const [userPlants, setUserPlants] = useState([])
  const displayPlants = async (e) => {
    e.preventDefault();
    try {
      const response = await fetch('/plants', {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
        }
      });
      if (response.ok) {
        setUserPlants(await response.json())
      }
    } catch (error) {
      console.log('Error with fetching plants');
    }
  };
  // console.log('userPlants', userPlants)
  // console.log('displayUserPlants', displayPlants)

  return (
    <>
    <br></br>
     <button onClick={displayPlants}>Display All Plants</button>
      <h2>PlantList</h2>
      {userPlants.map((plant) => 
        <>
          <label>{plant.plantName}</label>
          <input type="checkbox" key={plant.plantName}/>
          <br></br>
        </>
      )}
     
    </>

  )
}

export default PlantList
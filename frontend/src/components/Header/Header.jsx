import React from 'react'
import "./Header.css"

const Header = () => {
  return (
    <div className="header">
      <h1>Planting Schedule</h1>
      <div className="colorKey">
        <h2>Start Seeds Indoors</h2>
        <h2>Plant Seed/Transplant</h2>
        <h2>Harvest</h2>
      </div>
    </div>
  )
}

export default Header
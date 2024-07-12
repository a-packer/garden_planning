import React from 'react'
import BarChart from './BarChart';

export const GuideChart = () => {
  const plantingData = [
    {
      name: 'Basil',
      numWeeksToSproutIndoors: 5,
      weeksRelToFrostDateStartOutdoors: 1,
      totalGrowthMonths: 5,
    },
    {
      name: 'Beets',
      numWeeksToSproutIndoors: 5,
      weeksRelToFrostDateStartOutdoors: -2,
      totalGrowthMonths: 2.5,
    },
    {
      name: 'Broccoli',
      numWeeksToSproutIndoors: 5,
      weeksRelToFrostDateStartOutdoors: -2,
      totalGrowthMonths: 3,
    },
    {
      name: 'Carrots',
      numWeeksToSproutIndoors: 0,
      weeksRelToFrostDateStartOutdoors: -2,
      totalGrowthMonths: 2.5,
    },
    {
      name: 'Collards',
      numWeeksToSproutIndoors: 5,
      weeksRelToFrostDateStartOutdoors: -4,
      totalGrowthMonths: 3,
    } 
  ];
  return (
    <BarChart data={plantingData} />
  )
}

export default GuideChart;

import React, { useEffect, useRef } from 'react';
import * as d3 from 'd3';

const BarChart = ({ data }) => {

  console.log(data)
  const ref = useRef();

  useEffect(() => {
    const svg = d3.select(ref.current)
      .attr('width', 500)
      .attr('height', 300)
      .style('background-color', '#f0f0f0')
      .style('margin', '10px');

    // Define the gradient
    const defs = svg.append("defs");
    const gradient = defs.append("linearGradient")
      .attr("id", "gradient")
      .attr("x1", "0%")
      .attr("y1", "0%")
      .attr("x2", "100%")
      .attr("y2", "0%");
    gradient.append("stop")
      .attr("offset", "0%")
      .attr("stop-color", "green");
    gradient.append("stop")
      .attr("offset", "100%")
      .attr("stop-color", "salmon");  

    // Define the scales
    const xScale = d3.scaleLinear()
      .domain([0, d3.max(data, d => d.totalGrowth)])
      .range([0, 480]);

    const yScale = d3.scaleBand()
      .domain(data.map(d => d.plantName))
      .range([0, 280])
      .padding(0.1);

    // Draw bars
    svg.selectAll('.bar')
      .data(data)
      .enter()
      .append('rect')
      .attr('class', 'bar')
      .attr('x', 0)
      .attr('y', d => yScale(d.plantName))
      .attr('width', d => xScale(d.totalGrowth))
      .attr('height', yScale.bandwidth())
      .style('opacity', '0.8')
      .attr('fill', 'url(#gradient)');

    // Add axes
    svg.append('g')
      .attr('transform', 'translate(0,280)')
      .call(d3.axisBottom(xScale));

    svg.append('g')
      .call(d3.axisLeft(yScale));
  }, [data]);

  return (
    <svg ref={ref}></svg>
  );
};

export default BarChart;
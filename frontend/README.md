
# in garden-react-app root directory
`go run main.go`

# in src
`npm start`


# TODO:

## API 
Create a REST api with Go
Based on frost date and a list of crops (inputed by user on the frontend), create JSON data to send to D3 modal renderer.
https://medium.com/@Moesif/building-a-restful-api-with-go-dbd6e7aecf87

## Update the D3 modal
x axis need to be properly defined in months of the year
Add horizonal lines along the x-axis ticks to separate out months
Update how the bars are visualized
  left most point is based on the start date of seeds
  color change to yellow at point where one transplants to outdoors or starts seed if direct sow
  color gradient to peach at right most point

## React
Make the D3 modal interactive
Allow user to input a last frost date
Allow user to delete crops
Allow user to add crops (need to create another form for this - maybe a list of crops with checkmarks next to them?)

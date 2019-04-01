# Weather service
### Purpose
This service provides API that allows users to maintain favorite locations and generate for them weather statistics such as:
* Minimum temperature for each month
* Maximum temperature for each month
* Average temperature for each month
* Number of statistics data
* Overall weather conditions aggregated by days 
### Installation and Configuration
###### Download and build images
* git clone https://github.com/mieczyslaw1980/weather.git
* cd weather
* make containers
###### Start application
OPEN_WEATHER_MAP_TOKEN=<span style="color:red">[YOUR_OPEN_WEATHER_MAP_API_TOKEN]</span> docker-compose -f deployments/docker-compose.yml up
### Endpoints
1. Locations
* GET "/locations"
```$xslt
Get all user's locations
```
* GET "/locations/{id}"
```
Get one user's location
```
* DELETE "/locations/{id}"
```
Delete one user location
```
* POST "/locations"
```
Save new user's location by city name: 
   {"city_name": "London"}
Save new user's location by city name and country code:
   {"city_name": "London", "country_code": "GB"}
```
2. Weather
* GET "/weather/{id}"
```
Get current weather condition at the moment and save that for later statistis
```
* GET "/weather/{id}/statistics"
```
Calculate statistics for previous cumulated weather conditions
```

### API Documentation

https://github.com/mieczyslaw1980/weather/blob/master/api/swagger.json

### Unit tests
###### Start test
```
make test
```


### Examples
###### Create new location
* Create by city name

Request:
```
curl -X POST -H "content-type: application/json" --data '{"city_name": "Warsaw"}' localhost:8080/locations
```
Response:
```$xslt
{
 "city_name": "Warsaw",
 "country_code": "PL",
 "location_id": 756135,
 "latitude": 52.23,
 "longitude": 21.01
}
```
* Create by city name and country code

Request:
```
curl -vvv -X POST -H "content-type: application/json" --data '{"city_name": "London", "country_code": "GB"}' localhost:8080/locations
```

Response:
```$xslt
{
 "city_name": "London",
 "country_code": "GB",
 "location_id": 2643743,
 "latitude": 51.51,
 "longitude": -0.13
* Connection #0 to host localhost left intact
}
```

###### Delete location

Request:
```
curl -X DELETE  localhost:8080/locations/756135
```

Response:
```$xslt
HTTP Status: 200
```

###### Get locations

Request:
 ```
curl localhost:8080/locations
```
Response:
```
[
 {
  "city_name": "London",
  "country_code": "GB  ",
  "location_id": 2643743,
  "latitude": 51.51,
  "longitude": -0.13
 },
 {
  "city_name": "Warsaw",
  "country_code": "PL  ",
  "location_id": 756135,
  "latitude": 52.23,
  "longitude": 21.01
 }
]
```

###### Get location

Request:
```
curl localhost:8080/locations/2643743
```
Response:
```
{
 "city_name": "London",
 "country_code": "GB  ",
 "location_id": 2643743,
 "latitude": 51.51,
 "longitude": -0.13
}
```

###### Get weather conditions for location
Request:
```
curl localhost:8080/weather/2643743
```
Response:
```
{
 "temperature": 280.74,
 "LocationID": 2643743,
 "temp_min": 278.71,
 "temp_max": 282.59,
 "conditions": [
  {
   "statistic_id": 3,
   "type": "Clear"
  }
 ]
}
```

###### Get weather statistics for location
Request:
```
curl localhost:8080/weather/2643743/statistics
```
Response:
```
{
 "Count": 321,
 "MonthTemperature": [
  {
   "Min": 278.71,
   "Max": 283.15,
   "Avg": 280.80,
   "Month": "2019-03"
  },
  {
   "Min": 273.71,
   "Max": 278.15,
   "Avg": 285.85,
   "Month": "2018-03"
  }
 ],
 "DailyCondition": {
  "Clear": [
   "2019-03-31"
   "2019-03-30"
  ],
  "Clouds": [
   "2019-03-31"
  ],
  "Rain": [
   "2019-03-31",
   "2019-03-25",
   "2019-03-30",
   "2019-03-04",
  ]
 }
}
```


# weather-hybrid-work
**Description**
---
Weather hybrid work is a program that will help you to decide if you should work from home or not based on the weather forecast. It will ask you for your city and your schedule and will give you a ranking for each day of the week. The ranking is based on the precipitation forecast for the week and the apparent temperature. The ranking is from 1 to 5 where 1 is the best day to work from home. The program will also give you the forecast for the next week so you can plan ahead.

**Build**
---

1. Install go 1.19 or higher
2. Go to the project folder and run `go build`
3. For now you need to create a config.json file in the same folder as the executable with empty brackets `{}` inside but it will be handle in the future

**Usage**
---
Run the executable
1. Enter a city name and select the one corresponding in the list
2. Enter your schedule for when you are outside 
3. Select the forecast for the current week or the next one and the prediction appear with ranking where 1 is the best day to work at home
**Language**
---
Go 1.19

**Api used**
---
[Open Meteo geocoding](https://open-meteo.com/en/docs/geocoding-api)
[Open Meteo forecast](https://open-meteo.com/)

**TODO**
---
* Add exit to asktime
* Handle config.json
* Generalise ask methods
* Graphical interface
* Chart
* Better error handling
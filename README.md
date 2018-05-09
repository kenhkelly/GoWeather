# GoWeather 

Version 3.2

*Note: you will need an API key, you may register here: [https://developer.forecast.io/register](https://developer.forecast.io/register). Copy goweather.properties.sample to HOMEDIR/.goweather/goweather.properties and fill in your API key*

*Build it*: 

`./build.sh`

*Simple usage*

`./goweather`

*Note: you can copy the built goweather bin to /usr/local/bin/goweather and use it from any location*

*Options*

- -forecast=false: Show 8 day forecast  

*Examples*:

```
$ ./goweather
GoWeather 3.2 - @kenhkelly

      Found you in: Pompano Beach, Florida

      Current Weather: 09 May 18 17:02 EDT

        Summary     Partly Cloudy
        Temperature 83.059998
        Humidity    0.550000
        WindSpeed   13.720000
        WindBearing 54.000000

```

### Changelog

Version | Change
--------|----------
[v3.2]  | Add "found you in: {location}" and fixed panic bug related to bad API token
[v3.1]  | Add blank line after the output to help with spacing
[v3.0]  | Rewrite of the app to use forecast.io as the API source, because the last API no longer functioned
[v2.2]  | Add ability to use wind with a flag <br> Reorder the flags in the code and add descriptions to the flags <br> Remove the unnecessary help flag
[v2.1]  | Various bug fixes. 
[v2.0]  | Added the days and help flag. 
[v1.2]  | Added auto-detection of location when location is not provided
v1.1    | Refactor put in place
v1.0    | Initial version, provided the weather

[v3.2]: https://github.com/kenhkelly/GoWeather/tree/v3.2
[v3.1]: https://github.com/kenhkelly/GoWeather/tree/v3.1
[v3.0]: https://github.com/kenhkelly/GoWeather/tree/v3.0
[v2.2]: https://github.com/kenhkelly/GoWeather/tree/v2.2
[v2.1]: https://github.com/kenhkelly/GoWeather/tree/v2.1
[v2.0]: https://github.com/kenhkelly/GoWeather/tree/v2.0
[v1.2]: https://github.com/kenhkelly/GoWeather/tree/v1.2

#GoWeather

*Build it*: 

`go build goweather.go`

*Simple usage*

`./goweather 10001`

*Options*

- -days=1        "Shows forecasts for number of days (1-16)"
- -unit=imperial "Imperial, metric, or kelvin units of measurement"

*Examples*:

```
$ ./goweather 10001
Tuesday, September 22, 2015
 Current temp   Today's high   Today's low    Condition
 67.48          71.44          58.68          overcast clouds
```

```
$ ./goweather -days=7 10001
Tuesday, September 22, 2015
 Current temp   Today's high   Today's low    Condition
 67.48          71.44          58.68          overcast clouds

Wednesday, September 23, 2015
 Current temp   Today's high   Today's low    Condition
 69.58          75.90          52.97          few clouds

Thursday, September 24, 2015
 Current temp   Today's high   Today's low    Condition
 71.35          76.21          53.71          scattered clouds

Friday, September 25, 2015
 Current temp   Today's high   Today's low    Condition
 62.64          66.65          58.80          overcast clouds

Saturday, September 26, 2015
 Current temp   Today's high   Today's low    Condition
 68.47          68.47          59.13          light rain

Sunday, September 27, 2015
 Current temp   Today's high   Today's low    Condition
 66.33          66.33          58.46          light rain

Monday, September 28, 2015
 Current temp   Today's high   Today's low    Condition
 65.23          65.23          59.77          heavy intensity rain
```

```
$ ./goweather -days=3 -unit=metric 10001
Tuesday, September 22, 2015
 Current temp   Today's high   Today's low    Condition
 19.08          21.35          14.34          overcast clouds

Wednesday, September 23, 2015
 Current temp   Today's high   Today's low    Condition
 20.56          24.10          11.25          few clouds

Thursday, September 24, 2015
 Current temp   Today's high   Today's low    Condition
 21.84          24.58          11.96          scattered clouds
```

```
$ ./goweather -wind=true -days=1 10001
Friday, September 25, 2015
 Current temp   Today's high   Today's low    Condition           Wind speed     Wind direction
 72.18          72.18          60.15          broken clouds       3.43           East
```

### Changelog

Version | Change
--------|----------
[v2.1]  | Various bug fixes. 
[v2.0]  | Added the days and help flag. 
[v1.2]  | Added auto-detection of location when location is not provided
v1.1    | Refactor put in place
v1.0    | Initial version, provided the weather

[v2.1]: https://github.com/kenhkelly/GoWeather/tree/v2.1
[v2.0]: https://github.com/kenhkelly/GoWeather/tree/v2.0
[v1.2]: https://github.com/kenhkelly/GoWeather/tree/v1.2
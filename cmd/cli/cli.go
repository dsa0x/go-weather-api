package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli"
)

var owmAPI string = "https://api.openweathermap.org/data/2.5/weather?q="
var owmAPIID string = "&appid=4239f64721234295d28a661cf628c515"
var owmOthers string = "&units=metric"

// Clouds struct
type Clouds struct {
	All int `json:"all"`
}

// Coordinates struct
type Coordinates struct {
	Longitude float64 `json:"lon"`
	Latitude  float64 `json:"lat"`
}

// Main struct
type Main struct {
	Temp      float64 `json:"temp"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	FeelsLike float64 `json:"feels_like"`
	Pressure  float64 `json:"pressure"`
	SeaLevel  float64 `json:"sea_level"`
	GrndLevel float64 `json:"grnd_level"`
	Humidity  int     `json:"humidity"`
}

// Sys struct
type Sys struct {
	Type    int     `json:"type"`
	ID      int     `json:"id"`
	Message float64 `json:"message"`
	Country string  `json:"country"`
	Sunrise int     `json:"sunrise"`
	Sunset  int     `json:"sunset"`
}

// Weather struct
type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
}

// Wind struct
type Wind struct {
	Speed float64 `json:"speed"`
	Deg   float64 `json:"deg"`
}

// CurrentWeather struct
type CurrentWeather struct {
	Coord      Coordinates `json:"coord"`
	Sys        Sys         `json:"sys"`
	Base       string      `json:"base"`
	Weather    []Weather   `json:"weather"`
	Main       Main        `json:"main"`
	Wind       Wind        `json:"wind"`
	Clouds     Clouds      `json:"clouds"`
	Dt         int         `json:"dt"`
	Visibility int         `json:"visibility"`
	ID         int         `json:"id"`
	Name       string      `json:"name"`
	Cod        int         `json:"cod"`
	Timezone   int         `json:"timezone"`
}

func main() {
	app := cli.NewApp()
	app.Name = "Weather App CLI"
	app.Usage = "Fetch weather information for any city"

	myFlags := []cli.Flag{
		&cli.StringFlag{
			Name:  "host",
			Value: "dsa0x",
		},
		&cli.StringFlag{
			Name:  "city",
			Value: "Berlin",
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:  "owm",
			Usage: "Fetch weather from Open Weather Map",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				res, err := http.Get(owmAPI + c.String("city") + owmAPIID + owmOthers)
				if err != nil {
					return err
				}
				var r CurrentWeather

				if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
					fmt.Println("Please check your inputs.")
					return err
				}
				fmt.Println("City:", r.Name, ",", r.Sys.Country)
				fmt.Println("Temperature:", r.Main.Temp, "°C")
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

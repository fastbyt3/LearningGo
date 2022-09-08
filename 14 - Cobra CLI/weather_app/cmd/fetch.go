package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

type Response struct {
	Requests  Request `json:"request"`
	Locations Loc     `json:"location"`
	Currents  Current `json:"current"`
}

type Request struct {
	Type  string
	Query string
	Lang  string
	Unit  string
}

type Loc struct {
	Name           string
	Country        string
	Region         string
	Lat            string
	Lon            string
	TimeZoneId     string
	LocalTime      string
	LocalTimeEpoch int
	UTCOffset      string
}

type Current struct {
	Time         string
	Temp         int      `json:"temperature"`
	WeatherCode  int      `json:"weather_code"`
	WeatherIcons []string `json:"weather_icons"`
	WeatherDescs []string `json:"weather_descriptions"`
	WindSpeed    int      `json:"wind_speed"`
	WindDeg      int      `json:"wind_degree"`
	WindDir      string   `json:"wind_degree"`
	Pressure     int      `json:"pressure"`
	Precip       int      `json:"precip"`
	Humidity     int      `json:"humidity"`
	Cloudcover   int      `json:"cloudcover"`
	FeelsLike    int      `json:"feelslike"`
	UVIndex      int      `json:"uv_index"`
	Visibiltiy   int      `json:"visibility"`
	IsDay        string   `json:"is_day"`
}

var ApiKey string
var Location string

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch weather data",
	Long:  `Fetch weather data from weatherstack API using API key`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("fetch called")
		resp, err := http.Get("http://api.weatherstack.com/current?access_key=" + ApiKey + "&query=" + Location)

		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()

		var r Response
		err = json.NewDecoder(resp.Body).Decode(&r)
		if err != nil {
			panic(err)
		}
		fmt.Println(r.Currents)
		//var j interface{}
		//err = json.NewDecoder(resp.Body).Decode(&j)
		//if err != nil {
		//  panic(err)
		//}
		//fmt.Println(j)
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fetchCmd.PersistentFlags().String("foo", "", "A help for foo")
	fetchCmd.PersistentFlags().StringVarP(&ApiKey, "apikey", "a", "", "Provide API key to fetch data from weatherstack")
	fetchCmd.PersistentFlags().StringVarP(&Location, "location", "l", "", "Provide location")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fetchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

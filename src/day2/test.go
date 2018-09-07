package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/urfave/cli"
)

func main() {
	fmt.Println(request("http://www.sojson.com/open/api/weather/json.shtml?city=杭州"))
}

type Response struct {
	status   int // 'json:"status"'
	cityName string
	data     Data
	date     string
	message  string
	count    int
}
type Data struct {
	humidity string

	quality     string
	sensitivity string
	yestoday    Weather
	forecast    []Weather
}
type Weather struct {
	dateString      string
	sunriseTime     string
	sunsetTime      string
	highTemperature string
	lowTemperature  string
	aqi             float32
	windDerection   string
	windePower      string
	weatherType     string
	notice          string
}

func getWeatherInfo(url string, cityName string) (day string, r Response) {
	app := cli.NewApp()

}

func request(url string) (string, error) {

	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	body, b2 := ioutil.ReadAll(response.Body)
	response.Body.Close()
	return string(body), b2
}

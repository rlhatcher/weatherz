package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

/*
Copyright © 2019 Ronald Hatcher <ronaldhatcher@mac.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

// The format for Metoffice dates
const dateForm = "2006-01-02Z"

// MetOfficeFeed structure for weather feed from datapoint
type MetOfficeFeed struct {
	Site Site `json:"SiteRep"`
}

// Site of the forecast
type Site struct {
	Params Params `json:"Wx"`
	Data   Data   `json:"DV"`
}

// Params supplied in the report
type Params struct {
	Param []Param `json:"Param"`
}

// Param item for the report
type Param struct {
	Name  string `json:"name"`
	Units string `json:"units"`
	Label string `json:"$"`
}

// Data for the forecast
type Data struct {
	DataDate time.Time `json:"dataDate"`
	Type     string    `json:"type"`
	Location Location  `json:"Location"`
}

// Location of the report
type Location struct {
	I         string   `json:"i"`
	Lat       string   `json:"lat"`
	Lon       string   `json:"lon"`
	Name      string   `json:"name"`
	Country   string   `json:"country"`
	Continent string   `json:"continent"`
	Elevation string   `json:"elevation"`
	Periods   []Period `json:"Period"`
}

// Period of the report
type Period struct {
	Type    string   `json:"type"`
	Value   string   `json:"value"`
	Reports []Report `json:"Rep"`
}

// Report structure for a single report point
type Report struct {
	WindDirection string `json:"D"`
	FeelsLike     string `json:"F"`
	WindGustSpeed string `json:"G"`
	Humidity      string `json:"H"`
	PrecipProb    string `json:"Pp"`
	WindSpeed     string `json:"S"`
	Temperature   string `json:"T"`
	Visability    string `json:"V"`
	Weather       string `json:"W"`
	UvIndex       string `json:"U"`
	MinutesOfDay  string `json:"$"`
}

func metOfficeIcon(code string) (string, error) {

	icons := map[int]string{
		0:  "wsymbol_0008_clear_sky_night",
		1:  "wsymbol_0001_sunny",
		2:  "wsymbol_0041_partly_cloudy_night",
		3:  "wsymbol_0002_sunny_intervals",
		4:  "na",
		5:  "wsymbol_0006_mist",
		6:  "wsymbol_0007_fog",
		7:  "wsymbol_0003_white_cloud",
		8:  "wsymbol_0004_black_low_cloud",
		9:  "wsymbol_0025_light_rain_showers_night",
		10: "wsymbol_0009_light_rain_showers",
		11: "wsymbol_0048_drizzle",
		12: "wsymbol_0017_cloudy_with_light_rain",
		13: "wsymbol_0026_heavy_rain_showers_night",
		14: "wsymbol_0010_heavy_rain_showers",
		15: "wsymbol_0018_cloudy_with_heavy_rain",
		16: "wsymbol_0029_sleet_showers_night",
		17: "wsymbol_0013_sleet_showers",
		18: "wsymbol_0021_cloudy_with_sleet",
		19: "wsymbol_0030_light_hail_showers_night",
		20: "wsymbol_0014_light_hail_showers",
		21: "wsymbol_0022_cloudy_with_light_hail",
		22: "wsymbol_0027_light_snow_showers_night",
		23: "wsymbol_0011_light_snow_showers",
		24: "wsymbol_0019_cloudy_with_light_snow",
		25: "wsymbol_0028_heavy_snow_showers_night",
		26: "wsymbol_0012_heavy_snow_showers",
		27: "wsymbol_0020_cloudy_with_heavy_snow",
		28: "wsymbol_0032_thundery_showers_night",
		29: "wsymbol_0016_thundery_showers",
		30: "wsymbol_0024_thunderstorms",
	}

	i, err := strconv.Atoi(code)
	if err != nil {
		return code, err
	}

	if i < 0 || i > len(icons) {
		return code, errors.New("code out of range")
	}

	if colour {
		return "colour/" + icons[i], nil
	}
	return "bw/" + icons[i], nil
}

func metOffice() Summary {
	// Our JSON structure returned from the API call
	var feed MetOfficeFeed

	// Each part of the URL we need to call
	apiURL := "http://datapoint.metoffice.gov.uk"
	apiPath := "/public/data/val/wxfcs/all/json/322089"
	apiParams := "?res=3hourly"
	apiKey := "&key=" + apikey

	// Make the call
	response, err := http.Get(apiURL + apiPath + apiParams + apiKey)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		json.Unmarshal(data, &feed)
	}

	// Make sure we only return 5 days forecats even if we've been asked for more
	var reportdays int
	if numdays < 5 {
		reportdays = numdays
	} else {
		reportdays = 5
	}

	Forecasts := make([]Forecast, reportdays)

	Periods := feed.Site.Data.Location.Periods

	for i := 0; i < len(Periods) && i < reportdays; i++ {

		t, _ := time.Parse(dateForm, Periods[i].Value)

		Reports := Periods[i].Reports
		Forecasts[i].Date = t.String()
		Forecasts[i].High = Reports[0].Temperature
		Forecasts[i].Low = Reports[0].Temperature
		Forecasts[i].Description = t.Weekday().String()

		for j := 0; j < len(Reports); j++ {
			Forecasts[i].High, _ = strmax(Forecasts[i].High, Reports[j].Temperature)
			Forecasts[i].Low, _ = strmin(Forecasts[i].Low, Reports[j].Temperature)
			if Reports[j].MinutesOfDay == "720" {
				Forecasts[i].Icon, _ = metOfficeIcon(Reports[j].Weather)
			}
		}
		// If we check after mid-day, use the latest icon
		if Forecasts[i].Icon == "" {
			Forecasts[i].Icon, _ = metOfficeIcon(Reports[len(Reports)-1].Weather)
		}
	}

	var Summary Summary
	Summary.Forecasts = Forecasts

	return Summary
}

// Return the minimum of two string values representing integers
func strmin(a, b string) (string, error) {

	i, err := strconv.Atoi(a)
	if err != nil {
		return a, err
	}
	j, err := strconv.Atoi(b)
	if err != nil {
		return b, err
	}
	if i <= j {
		return a, nil
	}
	return b, nil
}

// Return the minimum of two string values representing integers
func strmax(a, b string) (string, error) {

	i, err := strconv.Atoi(a)
	if err != nil {
		return a, err
	}
	j, err := strconv.Atoi(b)
	if err != nil {
		return b, err
	}
	if i >= j {
		return a, nil
	}
	return b, nil
}

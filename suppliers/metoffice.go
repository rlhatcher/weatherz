package cmd

import "time"

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

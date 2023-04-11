package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"golang.org/x/text/encoding/charmap"
)

const (
	UserAgent string = "Golang_Spider_Bot/3.0"
	daysRange int    = 90
)

type CurrencyInfo struct {
	MaxValue ValueInfo
	MinValue ValueInfo
	AvgValue map[string]float64
}

type ValueInfo struct {
	Name  string
	Value float64
	Date  string
}

type ValCurs struct {
	Valute []Valute `xml:"Valute"`
	Date   string   `xml:"_Date"`
	Name   string   `xml:"_name"`
}

type Valute struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Nominal  string `xml:"Nominal"`
	Name     string `xml:"Name"`
	Value    string `xml:"Value"`
	ID       string `xml:"_ID"`
}

func main() {
	fmt.Println("Program started: Please wait..\n")
	var Courses []ValCurs
	getLast90Days(&Courses)

	var currInfo CurrencyInfo
	currInfo.AvgValue = make(map[string]float64)

	maxInfo := ValueInfo{}
	minInfo := ValueInfo{}

	commaReplacer := strings.NewReplacer(",", ".")

	for _, courses := range Courses {
		for _, currency := range courses.Valute {
			someVal, _ := strconv.ParseFloat(commaReplacer.Replace(currency.Value), 5)
			nominal, _ := strconv.ParseFloat(commaReplacer.Replace(currency.Nominal), 2)
			val := someVal / nominal

			if maxInfo.Value < val {
				maxInfo = ValueInfo{
					Name:  currency.Name,
					Value: val,
					Date:  courses.Date,
				}
			} else if minInfo.Name == "" || minInfo.Value > val {
				minInfo = ValueInfo{
					Name:  currency.Name,
					Value: val,
					Date:  courses.Date,
				}
			}
			currInfo.AvgValue[currency.Name] += val
		}
	}
	currInfo.MaxValue = maxInfo
	currInfo.MinValue = minInfo
	for key, _ := range currInfo.AvgValue {
		currInfo.AvgValue[key] /= float64(daysRange)
	}

	fmt.Printf("Max currency rate: %s, Value: %f, Date: %s\n", currInfo.MaxValue.Name, currInfo.MaxValue.Value, currInfo.MaxValue.Date)
	fmt.Printf("Min currency rate: %s, Value: %f, Date: %s\n", currInfo.MinValue.Name, currInfo.MinValue.Value, currInfo.MinValue.Date)
	fmt.Println("Average rates:")
	for key, value := range currInfo.AvgValue {
		fmt.Printf("%s: %f\n", key, value)
	}
}

// Iterate through (daysRange[default=90]) previous days and appending info to struct Courses
func getLast90Days(Courses *[]ValCurs) {
	t := time.Now()
	formatted := fmt.Sprintf("%02d/%02d/%d", t.Day(), t.Month(), t.Year())

	for i := 0; i < daysRange; i++ {
		if i == 0 { // for today
			var Prices ValCurs
			getResponse(&Prices, formatted)
			*Courses = append(*Courses, Prices)
			continue
		}

		newDate := t.AddDate(0, 0, -i)
		formatted = fmt.Sprintf("%02d/%02d/%d", newDate.Day(), newDate.Month(), newDate.Year())

		var Prices ValCurs

		getResponse(&Prices, formatted)
		*Courses = append(*Courses, Prices)
		// time.Sleep(time.Millisecond * 110) // Uncomment if needed
	}
}

// Getting page for certain date and writing info to struct Prices
func getResponse(Prices *ValCurs, date string) {
	url := fmt.Sprintf(
		"http://www.cbr.ru/scripts/XML_daily_eng.asp?date_req=%s",
		date,
	)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent", UserAgent)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("failed to fetch data: %d %s", resp.StatusCode, resp.Status)
	}

	d := xml.NewDecoder(resp.Body)
	d.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		switch charset {
		case "windows-1251":
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		default:
			return nil, fmt.Errorf("unknown charset: %s", charset)
		}
	}
	err = d.Decode(&Prices)
	if err != nil {
		log.Fatalf("can not decode to struct: %v", err)
	}
	Prices.Date = date
}

package main

import (
	"bytes"
	"encoding/csv"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type stockPrice struct {
	Date                   string
	Open, High, Low, Close float64
	Volume                 uint64
	AdjClose               float64
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	file, error := ioutil.ReadFile("table.csv")
	if error != nil {
		log.Fatalln(error)
	}

	f := bytes.NewReader(file)

	csv := csv.NewReader(f)

	records, error := csv.ReadAll()
	if error != nil {
		log.Fatalln(error)
	}

	var historicalStockPrices = []stockPrice{}

	for i, record := range records {
		if i == 0 {
			continue
		}

		dateString := strings.Split(record[0], "-")

		year, err := strconv.Atoi(dateString[0])
		if err != nil {
			log.Fatalln(err)
		}

		month, err := strconv.Atoi(dateString[1])
		if err != nil {
			log.Fatalln(err)
		}

		day, err := strconv.Atoi(dateString[2])
		if err != nil {
			log.Fatalln(err)
		}

		date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

		open, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Fatalln(err)
		}

		high, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			log.Fatalln(err)
		}

		low, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Fatalln(err)
		}

		close, err := strconv.ParseFloat(record[4], 64)
		if err != nil {
			log.Fatalln(err)
		}

		volume, err := strconv.ParseUint(record[5], 10, 64)
		if err != nil {
			log.Fatalln(err)
		}

		adjClose, err := strconv.ParseFloat(record[6], 64)
		if err != nil {
			log.Fatalln(err)
		}

		s := stockPrice{
			Date:     date.Format("Mon Jan _2, 2006"),
			Open:     open,
			High:     high,
			Low:      low,
			Close:    close,
			Volume:   volume,
			AdjClose: adjClose,
		}

		historicalStockPrices = append(historicalStockPrices, s)
	}

	error = tpl.Execute(os.Stdout, historicalStockPrices)
	if error != nil {
		log.Fatalln(error)
	}
}

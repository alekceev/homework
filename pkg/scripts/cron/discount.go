package main

import (
	"flag"
	"homework/pkg/database"
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

type Discount struct {
	Key      string `csv:"k"`
	Value    string `csv:"v"`
	Discount int    `csv:"discount"`
}

func main() {

	var fn string
	flag.StringVar(&fn, "file", "", "path/to/discount.csv")
	flag.Parse()
	if fn == "" {
		log.Panic("file not found")
	}

	file, err := os.Open(fn)
	if err != nil {
		log.Panicf("Error open file %s: %v", fn, err)
	}
	defer file.Close()

	discounts := []*Discount{}
	if err := gocsv.UnmarshalFile(file, &discounts); err != nil {
		log.Panic(err)
	}

	db := &database.DB{}
	if err := db.Open(); err != nil {
		log.Panicf("Db error: %v", err)
	}
	defer db.Close()

	// itemRepo := repositories.NewItemRepository(db.Dbh())

	// сброс цены со скидкой
	_, err = db.Dbh().Exec("update items set sale_price = price where price != sale_price")
	if err != nil {
		log.Panic(err)
	}

	for _, discount := range discounts {
		log.Printf("%#v", discount)
		switch discount.Key {
		case "category":
			//TODO update sale price for category
		case "items":
			//TODO update sale price for item by article
		case "-":
			//TODO update sale price for all
		}
	}
}

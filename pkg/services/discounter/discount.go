package discounter

import (
	"context"
	"fmt"
	"homework/pkg/interfaces"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gocarina/gocsv"
)

type DiscountService struct {
	repo          interfaces.Repository
	discountUrl   string
	nextStartTime time.Time
}

var _ interfaces.Service = &DiscountService{}

func NewDiscountService(repo interfaces.Repository, url string) *DiscountService {
	return &DiscountService{
		repo:          repo,
		discountUrl:   url,
		nextStartTime: jobStartedTime(),
	}
}

func (s *DiscountService) Start(ctx context.Context) error {
	for range time.Tick(1 * time.Minute) {
		now := time.Now()

		if now.Equal(s.nextStartTime) || now.After(s.nextStartTime) {
			s.nextStartTime = jobStartedTime()

			log.Println("Starting discount service...")
			go s.do()
		}
	}
	return nil
}

func (s *DiscountService) Stop(ctx context.Context) error {
	log.Println("Stoping discount service...")
	return nil
}

type Discount struct {
	Key      string `csv:"k"`
	Value    string `csv:"v"`
	Discount int    `csv:"discount"`
}

func (s *DiscountService) do() {
	discounts, err := fetchDiscounts(s.discountUrl)
	// discounts, err := getDiscounts("./web/discount.csv")
	if err != nil {
		log.Panic(err)
	}

	items, _ := s.repo.GetAll()
	log.Printf("Items: %d", len(items))

	// сброс цены со скидкой
	// _, err = s.repo.Db.Raw().Exec("update items set sale_price = price where price != sale_price")
	// if err != nil {
	// 	log.Panic(err)
	// }

	for _, discount := range discounts {
		log.Printf("Discount: %#v", discount)
		switch discount.Key {
		case "category":
			//TODO update sale price for category
		case "items":
			//TODO update sale price for item by number
		case "-":
			//TODO update sale price for all
		}
	}
}

// discounts from url
func fetchDiscounts(url string) ([]*Discount, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Could not fetch file %s: %s", url, err)
	}
	defer resp.Body.Close()

	discounts := []*Discount{}
	if err := gocsv.Unmarshal(resp.Body, &discounts); err != nil {
		return nil, err
	}

	return discounts, nil
}

// discounts from file
func getDiscounts(fn string) ([]*Discount, error) {
	file, err := os.Open(fn)
	if err != nil {
		log.Panicf("Error open file %s: %v", fn, err)
	}
	defer file.Close()

	discounts := []*Discount{}
	if err := gocsv.UnmarshalFile(file, &discounts); err != nil {
		return nil, err
	}

	return discounts, nil
}

// every day at 05:00
func jobStartedTime() time.Time {
	now := time.Now()
	// return now.Add(2 * time.Minute)
	if now.Hour() < 5 {
		return time.Date(now.Year(), now.Month(), now.Day(), 5, 0, 0, 0, time.Local)
	}

	return time.Date(now.Year(), now.Month(), now.Day(), 5, 0, 0, 0, time.Local).Add(24 * time.Hour)
}

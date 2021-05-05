package main

import (
	"encoding/json"
	"homework/pkg/app"
	"homework/pkg/models"
	"homework/pkg/repositories"
	"log"
)

func main() {
	app := app.NewApp()
	defer app.Close()

	itemRepo := repositories.NewItemRepository(app.DB)

	data := []byte(`[
		{"name":"Intel Core i3-8100", "description":"Процессор Intel", "number":"i3-8100", "category":"Процессоры", "price":7890.00},
		{"name":"Intel Core i5-7400", "description":"Процессор Intel", "number":"i5-7400", "category":"Процессоры", "price":12700.00},
		{"name":"AMD FX-8320E", "description":"Процессор AMD", "number":"fx-8320e", "category":"Процессоры", "price":4780.00},
		{"name":"AMD FX-8320", "description":"Процессор AMD", "number":"fx-8320", "category":"Процессоры", "price":7120.00},
		{"name":"ASUS ROG MAXIMUS X HERO", "description":"Z370, Socket 1151-V2, DDR4, ATX", "number":"aZ370", "category":"Мат.Платы", "price":19310.00},
		{"name":"Gigabyte H310M S2H", "description":"H310, Socket 1151-V2, DDR4, mATX", "number":"gH310", "category":"Мат.Платы", "price":4790.00},
		{"name":"MSI B250M GAMING PRO", "description":"B250, Socket 1151, DDR4, mATX", "number":"msiB250", "category":"Мат.Платы", "price":5060.00},
		{"name":"GeForce GTX 1060", "description":"Видеокарты Nvidia", "number":"gtx-1060", "category":"Видеокарты", "price":12600},
		{"name":"GeForce GTX 1070", "description":"Видеокарты Nvidia", "number":"gtx-1070", "category":"Видеокарты", "price":22100},
		{"name":"Radeon RX 580", "description":"Видеокарты Nvidia", "number":"rx-580", "category":"Видеокарты", "price":16000},
		{"name":"HyperX DDR4 4GB", "description":"Память", "number":"h-ddr4-4", "category":"Память", "price":2200},
		{"name":"Crystal DDR4 8GB", "description":"Память", "number":"c-ddr3-8", "category":"Память", "price":3900}
	]`)

	var items models.Items
	err := json.Unmarshal(data, &items)
	if err != nil {
		log.Panicf("Error decode: %v", err)
	}

	for _, item := range items {
		log.Printf("item %v\n", item)
		found_item, err := itemRepo.FindByName(item.Name)

		// update item
		if err == nil && found_item != nil {
			log.Printf("Found item %d\n", found_item.Id)
			item.Id = found_item.Id
			err := itemRepo.Update(item)
			if err != nil {
				log.Printf("Error update item %d %v\n", item.Id, err)
			}
		} else { // add item
			err = itemRepo.Save(item)
			if err != nil {
				log.Printf("Error save item: %v\n", err)
			} else {
				log.Printf("Add item: %d\n", item.Id)
			}
		}
	}
}

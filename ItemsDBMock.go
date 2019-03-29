package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

var itemsDB Items

func loadItems(vdmFileName string) {
	startTime := time.Now()
	itemFile, err := os.Open(vdmFileName)
	if err != nil {
		log.Panicf("Error when reading file: ", err.Error())
	}

	jsonParser := json.NewDecoder(itemFile)

	if err = jsonParser.Decode(&itemsDB); err != nil {
		log.Panicf("Error when decoding file: ", err.Error())
	}

	endTime := time.Now()
	fmt.Printf("Items loaded in %s s", endTime.Sub(startTime))
	return
}

func selectItem(id int) Item {
	return itemsDB[id]
}

func getRandomItem() Item {
	randIndex := rand.Intn(len(itemsDB) - 1)
	return itemsDB[randIndex]
}

func IncItemRating(itemID int) Item {
	itemsDB[itemID].Rating++
	return itemsDB[itemID]
}

func addItem(item Item) Item {
	itemsDB = append(itemsDB, item)
	log.Println("item added at index: ", len(itemsDB)-1)
	return item
}

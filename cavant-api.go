package main

import (
	cavantdb "cavant-api/cavant-db"
	"log"
)

func main() {
	print("Starting\n")
	database := cavantdb.DB{}
	database.InitDB()

	err := database.AddNewTable("FishTable", true)
	if err != nil {
		log.Fatal(err)
	}

	err = database.AddDataToTable("FishTable", "", "")
	if err != nil {
		log.Fatal(err)
	}

	data, err := database.GetInitialTable()
	if err != nil {
		log.Fatal(err)
	}

	for _, value := range data {
		print(value + ", ")
	}
	print("\nEnding\n")
}

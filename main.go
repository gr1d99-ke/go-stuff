package main

import (
	"C"
	"encoding/csv"
	"log"
	"os"
)

func toCsv(filename, data string) {
	_, err := os.Stat(filename)
	if err != nil {
		file, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	}

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var csvdata [][]string

	csvdata = append(csvdata, []string{data})

	w := csv.NewWriter(file)
	w.WriteAll(csvdata)

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}


func main()  {
	toCsv("admin.csv", "admins")
	toCsv("admin.csv", "admins")
	toCsv("admin.csv", "admins")
	toCsv("admin.csv", "admins")
	toCsv("admin.csv", "admins")
	toCsv("admin.csv", "admins")
	toCsv("admin.csv", "admins")
	toCsv("admin.csv", "admins")
	toCsv("admin.csv", "admins")
	toCsv("admin.csv", "admins")
}

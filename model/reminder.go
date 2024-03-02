package model

import (
	"encoding/csv"
	"os"
)

type Reminder struct {
	Date string
	Name string
	Desc string
}

func (r Reminder) Save() {
	records, err := r.Read()
	if err != nil {
		panic(err)
	}

	records = append(records, []string{r.Date, r.Name, r.Desc})

	r.Write(records)
}

func (r Reminder) Write(records [][]string) {
	file, err := os.OpenFile(App.File, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	writer := csv.NewWriter(file)

	if err := writer.WriteAll(records); err != nil {
		panic(err)
	}
}

func (r Reminder) Read() ([][]string, error) {
	file, err := os.OpenFile(App.File, os.O_CREATE|os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	return records, nil
}

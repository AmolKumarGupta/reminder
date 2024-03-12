package model

import (
	"encoding/csv"
	"os"
	"sort"
	"time"

	"github.com/AmolKumarGupta/reminder-cli/config"
)

type Reminder struct {
	Date string
	Name string
	Desc string
}

func (r Reminder) Save() error {
	records, err := Read()
	if err != nil {
		return err
	}

	records = append(records, []string{r.Date, r.Name, r.Desc})

	sort.Slice(records, func(i, j int) bool {
		iDate, _ := time.Parse("02-01", records[i][0])
		jDate, _ := time.Parse("02-01", records[j][0])
		return iDate.Before(jDate)
	})

	if err := r.Write(records); err != nil {
		return err
	}

	return nil
}

func (r Reminder) Write(records [][]string) error {
	file, err := os.OpenFile(config.App.File, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	writer := csv.NewWriter(file)

	if err := writer.WriteAll(records); err != nil {
		return err
	}

	return nil
}

func Read() ([][]string, error) {
	file, err := os.OpenFile(config.App.File, os.O_CREATE|os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

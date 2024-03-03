package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/AmolKumarGupta/reminder-cli/model"
	"github.com/AmolKumarGupta/reminder-cli/util"
	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "add reminder",
	Long:  "add reminder in the store",
	Run: func(cmd *cobra.Command, args []string) {
		date, err := getDate()
		if err != nil {
			fmt.Println(err)
			return
		}

		name, err := get(util.Bold("Enter Title: "))
		if err != nil {
			fmt.Println(err)
			return
		}

		if name == "" {
			fmt.Println("name cannot be empty")
			return
		}

		desc, err := get(util.Bold("Enter Description: "))
		if err != nil {
			fmt.Println(err)
			return
		}

		rem := model.Reminder{
			Date: date,
			Name: name,
			Desc: desc,
		}

		if err := rem.Save(); err != nil {
			fmt.Print(err)
		}
	},
}

func getDate() (string, error) {
	var date string

	fmt.Print(util.Bold("Enter day and month (25-03): "))

	if _, err := fmt.Scanln(&date); err != nil {
		return "", errors.New("something is wrong while fetching date")
	}

	_, err := time.Parse("02-01", date)
	if err != nil {
		return "", errors.New("it is not a valid date, it should be like DD-MM")
	}

	return date, nil
}

func get(label string) (string, error) {
	var line string

	fmt.Print(label)

	b := bufio.NewReader(os.Stdin)

	line, err := b.ReadString('\n')
	if err != nil {
		return "", errors.New("something went wrong")
	}

	return strings.TrimSuffix(line, "\n"), nil
}

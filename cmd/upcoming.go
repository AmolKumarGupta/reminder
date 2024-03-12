package cmd

import (
	"fmt"
	"time"

	"github.com/AmolKumarGupta/reminder-cli/model"
	"github.com/AmolKumarGupta/reminder-cli/tui/table"
	"github.com/AmolKumarGupta/reminder-cli/tui/text"
	"github.com/spf13/cobra"
)

var UpComingCmd = &cobra.Command{
	Use:   "upcoming",
	Short: "List upcoming 5 reminder",
	Long:  "List upcoming 5 reminder based on the dates",
	Run: func(cmd *cobra.Command, args []string) {
		global(cmd)

		now := time.Now()

		fmt.Printf("%s\n\n", text.New("Upcoming Events").Bold())

		records, err := model.Read()
		if err != nil {
			fmt.Println("Something went wrong")
		}

		size := len(records) - 1

		index := search(now, records, 0, size)

		if index < 0 {
			fmt.Println("No Event Found !!!")
			return
		}

		data := [][]string{}
		for i := index; i < index+7; i++ {
			if i > size {
				break
			}

			data = append(data, records[i])
		}

		table.New().
			Header([]string{"Date", "Name", "Description"}).
			Add(data).
			Render()
	},
}

func date(str string, c time.Time) time.Time {
	t, _ := time.Parse("02-01", str)

	return time.Date(
		c.Year(), t.Month(), t.Day(),
		c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.Location(),
	)
}

func search(x time.Time, data [][]string, min, max int) int {
	if min > max {
		return -1
	}
	mid := min + int((max-min)/2)

	cur := date(data[mid][0], x)

	if mid > min {
		prev := date(data[mid-1][0], x)
		if prev.Before(x) && x.Equal(cur) {
			return mid
		}

	} else if x.Equal(cur) {
		return mid

	} else if mid == min && x.Before(cur) {
		return mid
	}

	if mid < max {
		nxt := date(data[mid+1][0], x)
		if cur.Before(x) && (x.Before(nxt) || x.Equal(nxt)) {
			return mid + 1
		}

	}

	if cur.Before(x) {
		return search(x, data, mid+1, max)

	} else {
		return search(x, data, min, mid)
	}
}

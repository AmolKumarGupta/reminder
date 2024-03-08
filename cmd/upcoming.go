package cmd

import (
	"fmt"
	"time"

	"github.com/AmolKumarGupta/reminder-cli/model"
	"github.com/AmolKumarGupta/reminder-cli/tui/text"
	"github.com/spf13/cobra"
)

var UpComingCmd = &cobra.Command{
	Use:   "upcoming",
	Short: "List upcoming reminder in next 7 days",
	Long:  "List upcoming reminder in next 7 days based on the dates",
	Run: func(cmd *cobra.Command, args []string) {
		now := time.Now()

		fmt.Printf("%s - %d %s\n\n", text.New("Upcoming Events").Bold(), now.Day(), now.Month())

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

		for i := index; i <= size; i++ {
			data := records[i]
			fmt.Printf("%s: %s\n", text.New(data[0]).Bold(), data[1])
		}
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

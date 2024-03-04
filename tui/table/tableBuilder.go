package table

import (
	"fmt"
	"strings"
)

type TableBuilder struct {
	header []string
	data   [][]string
}

func NewTableBuilder() *TableBuilder {
	return &TableBuilder{}
}

func (t *TableBuilder) Header(header []string) *TableBuilder {
	t.header = append(t.header, header...)
	return t
}

func (t *TableBuilder) Add(data [][]string) *TableBuilder {
	t.data = append(t.data, data...)
	return t
}

func (t TableBuilder) Render() {
	colWidths := t.width()

	t.renderLine(colWidths, TopLeft, DownHorizontal, TopRight)

	fmt.Printf("%s", Vertical)
	for i, header := range t.header {
		fmt.Printf("%-*s%s", colWidths[i], header, Vertical)
	}
	fmt.Println()

	t.renderLine(colWidths, VerticalRight, VerticalHorizontal, VerticalLeft)

	rowCount := len(t.data)
	for z, row := range t.data {
		fmt.Printf("%s", Vertical)

		for i, value := range row {
			fmt.Printf("%-*s%s", colWidths[i], value, Vertical)
		}

		fmt.Println()

		if z+1 == rowCount {
			t.renderLine(colWidths, BottomLeft, UpHorizontal, BottomRight)

		} else {
			t.renderLine(colWidths, VerticalRight, VerticalHorizontal, VerticalLeft)
		}
	}
}

func (t TableBuilder) width() []int {
	colWidths := make([]int, len(t.header))
	for i, header := range t.header {
		colWidths[i] = len(header)
		for _, row := range t.data {
			if len(row[i]) > colWidths[i] {
				colWidths[i] = len(row[i])
			}
		}
	}

	for i := range colWidths {
		colWidths[i] += 1
	}

	return colWidths
}

func (t TableBuilder) renderLine(colWidths []int, left string, center string, right string) {
	totalCols := len(colWidths)

	for i, width := range colWidths {
		if i == 0 {
			fmt.Printf("%s", left)
		}

		fmt.Printf("%s", strings.Repeat(Horizontal, width))

		if i+1 == totalCols {
			fmt.Printf("%s\n", right)
		} else {
			fmt.Printf("%s", center)
		}
	}
}

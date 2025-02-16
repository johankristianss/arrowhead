package cli

import (
	"strings"

	goprettytable "github.com/jedib0t/go-pretty/v6/table"
	"github.com/johankristianss/arrowhead/internal/table"
	"github.com/muesli/termenv"
)

func printKeyValueTable(title string, theme table.Theme, data map[string]string) {
	t := table.NewTable(theme, table.TableOptions{
		Columns: []int{0, 1}, // Two columns: "Field" | "Value"
		SortBy:  0,
		Style:   goprettytable.StyleRounded,
	}, false)

	title = strings.ToUpper(title)
	t.SetTitle(title)
	t.SetCols([]table.Column{
		{ID: "field", Name: "Field", SortIndex: 0, Width: 0},
		{ID: "value", Name: "Value", SortIndex: 1, Width: 0},
	})

	for key, value := range data {
		valueTruncated := truncateString(value)
		t.AddRow([]interface{}{
			termenv.String(key).Foreground(theme.ColorCyan),
			termenv.String(valueTruncated).Foreground(theme.ColorMagenta),
		})
	}

	t.Render()
}

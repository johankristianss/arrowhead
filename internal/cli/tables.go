package cli

import (
	goprettytable "github.com/jedib0t/go-pretty/v6/table"
	"github.com/johankristianss/arrowhead/internal/table"
)

func createTable(sortCol int) (*table.Table, table.Theme) {
	style := goprettytable.StyleRounded

	theme, err := table.LoadTheme("solarized-dark")
	CheckError(err)

	return table.NewTable(theme, table.TableOptions{
		Columns: []int{1, 2},
		SortBy:  sortCol,
		Style:   style,
	}, ASCII), theme
}

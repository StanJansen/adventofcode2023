package main

import (
	"embed"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/StanJansen/adventofcode2023/solvers"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

//go:embed input
var f embed.FS

func main() {
	ss := solvers.Solvers
	day := 0
	if len(os.Args) > 1 {
		day, _ = strconv.Atoi(os.Args[1])
		ss = solvers.Solvers[day-1 : day]
	}

	t := setupTable()

	for i, s := range ss {
		d := day
		if 0 == d {
			d = i + 1
		}
		c, _ := f.ReadFile(fmt.Sprintf("input/%d", d))
		input := string(c)
		r1, t1 := solve(s[0], input)
		r2, t2 := solve(s[1], input)
		t.AppendRows([]table.Row{
			{d, r1, t1, r2, t2},
		})
	}

	fmt.Println(t.Render())
}

func solve(function func(string) int, input string) (int, time.Duration) {
	t := time.Now()
	r := function(input)
	return r, time.Since(t)
}

func setupTable() table.Writer {
	t := table.NewWriter()
	t.SetStyle(table.StyleDouble)
	t.Style().Title.Align = text.AlignCenter
	t.Style().Title.Colors = text.Colors{text.Bold}
	t.Style().Options.SeparateRows = true
	t.Style().Format.Header = text.FormatDefault
	t.SetTitle("🎄 Advent of Code 2023 🎄")
	t.AppendHeader(table.Row{"Day", "Part 1", "Part 1", "Part 2", "Part 2"}, table.RowConfig{AutoMerge: true})

	return t
}

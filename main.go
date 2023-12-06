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

	t := table.NewWriter()
	t.SetStyle(table.StyleDouble)
	t.Style().Title.Align = text.AlignCenter
	t.Style().Title.Colors = text.Colors{text.Bold}
	t.Style().Options.SeparateRows = true
	t.Style().Format.Header = text.FormatDefault
	t.Style().Format.Footer = text.FormatDefault
	t.SetTitle("🎄 Advent of Code 2023 🎄")

	rc := table.RowConfig{AutoMerge: true}
	t.AppendHeader(table.Row{"Day", "Part 1", "Part 1", "Part 2", "Part 2"}, rc)

	for i, s := range ss {
		d := day
		if 0 == d {
			d = i + 1
		}
		c, _ := f.ReadFile(fmt.Sprintf("input/%d", d))
		input := string(c)
		t1 := time.Now()
		r1 := s[0].Solve(input)
		t2 := time.Now()
		r2 := s[1].Solve(input)
		t.AppendRows([]table.Row{
			{d, r1, time.Since(t1), r2, time.Since(t2)},
		})
	}

	fmt.Println(t.Render())
}

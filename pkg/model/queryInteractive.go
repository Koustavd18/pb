package model

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Creates a TUI App for pb query run command with interactive flag
func QueryUI(query string, startTime time.Time, endTime time.Time) {
	app := tview.NewApplication()
	start := startTime.Format("02-Jan-2006 15:04:05")
	end := endTime.Format("02-Jan-2006 15:04:05")

	queryBox := tview.NewTextView()

	queryBox.SetText(query).
		SetTextColor(tcell.ColorOliveDrab).
		SetDynamicColors(true).
		SetBorder(true).
		SetBorderColor(tcell.ColorYellow).
		SetTitle(" Query ").
		SetBorderPadding(1, 0, 1, 1)

	dateTimeBox := tview.NewTextView()

	dateTimeBox.SetText(fmt.Sprintf("Start Time : %s\nEnd Time   : %s", start, end)).
		SetTextColor(tcell.ColorWhite).
		SetDynamicColors(true).
		SetBorder(true).
		SetTitle(" Date-Time Stamp ").
		SetBorderPadding(1, 0, 1, 1)

	queryInfo := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(
			tview.NewFlex().
				SetDirection(tview.FlexColumn).
				AddItem(queryBox, 0, 7, false).
				AddItem(dateTimeBox, 0, 3, false),
			0, 1, false)

	queryInfoRow := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(queryInfo, 0, 1, false)

	table := tview.NewTable()

	headers := []string{"Column 1", "Column 2", "Column 3", "Column 4", "Column 5"}
	for i, header := range headers {
		table.SetCell(0, i, tview.NewTableCell(header).
			SetTextColor(tcell.ColorYellow).
			SetAlign(tview.AlignCenter)).SetBorder(true).SetTitle(" Query Table ").SetBorderPadding(2, 2, 2, 2)
	}

	for row := 1; row <= 20; row++ {
		for col := 0; col < len(headers); col++ {
			cellText := tview.NewTableCell("Row " + strconv.Itoa(row) + " Col " + strconv.Itoa(col+1)).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft)

			table.SetCell(row, col, cellText)
		}
	}

	countBox := tview.NewTextView()
	countBox.SetText(fmt.Sprintf("Showing %d out of %d", 20, 500)).SetBorderPadding(0, 0, 1, 0)

	commandList := tview.NewList().
		AddItem("Exit", "Exit Query Table", 'q', func() {
			app.Stop()
		}).
		AddItem("Command 2", "Command 2", 'w', nil)

	commandBox := tview.NewFlex().AddItem(commandList, 0, 1, true)
	commandBox.SetBorder(true).SetTitle(" Available Commands ")

	finalLayout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(queryInfoRow, 6, 0, false).
		AddItem(table, 0, 3, false).
		AddItem(countBox, 2, 0, false).
		AddItem(commandBox, 6, 0, true)
	finalLayout.SetBorder(true)

	if err := app.SetRoot(finalLayout, true).Run(); err != nil {
		panic(err)
	}
}


// func fetchData(query string, start string, end string ){

//     return 
// }
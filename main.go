package main

import (
	"fmt"
	"image/color"
	"mazegame/maze"
	"mazegame/screens"
	"mazegame/status"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Maze Game")
	w.Resize(fyne.NewSize(600, 600))
	w.SetFixedSize(true)
	var level maze.Level

	labelGameName := canvas.NewText("MAZE GAME", color.NRGBA{R: 165, G: 200, B: 240, A: 205})
	labelGameName.Resize(fyne.NewSize(450, 50))
	labelGameName.TextSize = 50
	labelGameName.TextStyle = fyne.TextStyle{Bold: true}
	labelGameName.Move(fyne.NewPos(137, 175))

	labelWinner := canvas.NewText("", color.NRGBA{R: 0, G: 220, B: 100, A: 200})
	labelWinner.TextSize = 18
	labelWinner.Move(fyne.NewPos(175, 350))
	labelWinner.Hide()

	buttonColor := canvas.NewRectangle(color.NRGBA{R: 165, G: 200, B: 240, A: 205})
	buttonColor.SetMinSize(fyne.NewSize(100, 50))

	buttonStart := widget.NewButton("Start", func() {
		w.Hide()
		timer := time.Now()

		ch := make(chan status.Status, 1)
		screens.GameScreen(ch, level)
		switch <-ch {
		case status.Positive:
			labelWinner.Text = fmt.Sprintf("You won! It took %v seconds.", int(time.Since(timer).Seconds()))
		case status.Negative:
			labelWinner.Text = "Play again!!!"
		}
		labelWinner.Show()

		time.AfterFunc(time.Second*5, func() {
			labelWinner.Hide()
			labelWinner.Refresh()
		})
		w.Show()

	})
	buttonStart.Resize(fyne.NewSize(100, 50))
	buttonStartContent := container.NewWithoutLayout(buttonColor, buttonStart)
	buttonStartContent.Move(fyne.NewPos(200, 250))

	buttonQuit := widget.NewButton("Quit", func() {
		w.Close()
	})
	buttonQuit.Resize(fyne.NewSize(100, 50))
	buttonQuitContent := container.New(layout.NewMaxLayout(), buttonColor, buttonQuit)
	buttonQuitContent.Move(fyne.NewPos(300, 250))

	levelLable := canvas.NewText("Level:", color.White)
	levelLable.TextStyle = fyne.TextStyle{Bold: true, Italic: true}
	levelLable.TextSize = 16
	levelLable.Resize(fyne.NewSize(50, 50))
	levelLable.Move(fyne.NewPos(200, 300))

	levelList := widget.NewSelect([]string{
		"Easy",
		"Medium",
		"Hard",
		"Advanced",
		"Wizard!",
	}, func(s string) {
		switch s {
		case "Easy":
			level = maze.Easy
		case "Medium":
			level = maze.Medium
		case "Hard":
			level = maze.Hard
		case "Advanced":
			level = maze.Advanced
		case "Wizard!":
			level = maze.Wizard
		}
	})
	levelList.Resize(fyne.NewSize(150, 50))
	levelList.Move(fyne.NewPos(250, 300))
	levelList.SetSelectedIndex(0)

	mainBox := container.NewWithoutLayout(labelWinner, labelGameName, buttonStartContent, buttonQuitContent, levelLable, levelList)

	r, _ := fyne.LoadResourceFromPath("./mazeImg.svg")
	w.SetIcon(r)
	w.SetContent(mainBox)
	w.ShowAndRun()
}

package main

import (
	"fmt"
	"image/color"
	"mazegame/maze"

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

	labelGameName := canvas.NewText("MAZE GAME", color.NRGBA{R: 255, G: 200, B: 25, A: 230})
	labelGameName.Resize(fyne.NewSize(450, 50))
	labelGameName.TextSize = 50
	labelGameName.TextStyle = fyne.TextStyle{Bold: true}
	labelGameName.Move(fyne.NewPos(150, 175))

	buttonColor := canvas.NewRectangle(color.NRGBA{R: 255, G: 240, B: 25, A: 230})
	buttonColor.SetMinSize(fyne.NewSize(100, 50))

	buttonStart := widget.NewButton("Start", func() {
		fmt.Println(level)
		// w.Hide()
		// isActive := make(chan bool, 1)
		// screens.GameScreen(isActive, level)
		// for {
		// 	switch {
		// 	case <-isActive:
		// 		w.Show()
		// 	}
		// }

	})
	buttonStart.Resize(fyne.NewSize(100, 50))
	// buttonStart.Move(fyne.NewPos(200, 250))
	buttonStartContent := container.NewWithoutLayout(buttonColor, buttonStart)
	buttonStartContent.Move(fyne.NewPos(200, 250))

	buttonQuit := widget.NewButton("Quit", func() {
		w.Close()
	})
	buttonQuit.Resize(fyne.NewSize(100, 50))
	// buttonQuit.Move(fyne.NewPos(400, 250))
	buttonQuitContent := container.New(layout.NewMaxLayout(), buttonColor, buttonQuit)
	buttonQuitContent.Move(fyne.NewPos(300, 250))

	levelLable := canvas.NewText("Level:", color.White)
	levelLable.TextStyle = fyne.TextStyle{Bold: true, Italic: true}
	levelLable.TextSize = 32
	levelLable.Resize(fyne.NewSize(100, 50))
	levelLable.Move(fyne.NewPos(200, 300))

	levelList := widget.NewSelect([]string{
		"Light",
		"SemiLight",
		"Medium",
		"SemiHard",
		"Hard",
	}, func(s string) {
		switch s {
		case "Light":
			level = maze.Light
		case "SemiLight":
			level = maze.SemiLight
		case "Medium":
			level = maze.Medium
		case "SemiHard":
			level = maze.SemiHard
		case "Hard":
			level = maze.Hard
		}
	})
	levelList.Resize(fyne.NewSize(100, 50))
	levelList.Move(fyne.NewPos(300, 300))
	levelList.SetSelectedIndex(0)

	// gameNameBox := container.New(layout.NewCenterLayout(), labelGameName)
	// gameNameBox.Move(fyne.NewPos(0, 0))
	// buttonsBox := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), buttonStart, buttonQuit, layout.NewSpacer())
	mainBox := container.NewWithoutLayout(labelGameName, buttonStartContent, buttonQuitContent, levelLable, levelList)

	w.SetContent(mainBox)
	w.ShowAndRun()
}

package screens

import (
	"image/color"
	"mazegame/maze"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	//"fyne.io/fyne/v2/layout"
)

func GameScreen(ch chan bool, level maze.Level) {
	myApp := app.New()
	myWindow := myApp.NewWindow("Container")
	black := color.NRGBA{R: 0, G: 0, B: 0, A: 255}
	white := color.NRGBA{R: 255, G: 255, B: 255, A: 255}

	lightmaze := maze.NewMazeBuilder(maze.Light)
	lightmaze.CreateMaze()
	// blocks := make([]canvas.Rectangle, 0)
	var myColor color.Color
	mainContent := container.NewWithoutLayout()

	content := container.NewWithoutLayout()

	myMaze := lightmaze.CreateMaze()
	for v := 0; v < myMaze.Height; v++ {
		for h := 0; h < myMaze.Width; h++ {
			if myMaze.Plan[v][h] {
				myColor = white
			} else {
				myColor = black
			}
			if myMaze.StartPoint.X == h && myMaze.StartPoint.Y == v ||
				myMaze.EndPoint.X == h && myMaze.EndPoint.Y == v {
				myColor = color.NRGBA{R: 230, G: 70, B: 20, A: 255}
			}
			block := canvas.NewRectangle(myColor)
			block.Resize(fyne.NewSize(float32(myMaze.BlockSize), float32(myMaze.BlockSize)))
			block.Move(fyne.NewPos(float32(h*myMaze.BlockSize), float32(v*myMaze.BlockSize)))

			// blocks = append(blocks, *block)
			content.Add(block)
		}
	}
	player := canvas.NewImageFromFile("./man.png")
	player.Resize(fyne.NewSize(float32(myMaze.BlockSize), float32(myMaze.BlockSize)))
	player.Move(fyne.NewPos(float32(myMaze.StartPoint.X*myMaze.BlockSize), float32(myMaze.StartPoint.Y*myMaze.BlockSize)))

	content.Move(fyne.NewPos(0, 0))

	mainContent.Add(content)
	mainContent.Add(player)
	btn := widget.NewButton("ccc", func() {
		player.Move(fyne.NewPos(player.Position().X-float32(myMaze.BlockSize), player.Position().Y-float32(myMaze.BlockSize)))
		player.Refresh()
	})
	btn.Resize(fyne.NewSize(20, 20))
	btn.Move(fyne.NewPos(500, 500))
	mainContent.Add(btn)

	myWindow.SetContent(mainContent)
	myWindow.ShowAndRun()
}

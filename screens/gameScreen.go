package screens

import (
	"image/color"
	"mazegame/maze"
	"mazegame/status"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func GameScreen(ch chan status.Status, level maze.Level) {
	myWindow := fyne.CurrentApp().NewWindow("Maze Game")
	dark := color.NRGBA{R: 48, G: 48, B: 48, A: 190}
	light := color.NRGBA{R: 165, G: 200, B: 240, A: 205}

	mazer := maze.NewMazeBuilder(level)
	var myColor color.Color
	mainContent := container.NewWithoutLayout()

	content := container.NewWithoutLayout()

	myMaze := mazer.CreateMaze()
	for v := 0; v < myMaze.Height; v++ {
		for h := 0; h < myMaze.Width; h++ {
			if myMaze.Plan[v][h] {
				myColor = light
			} else {
				myColor = dark
			}
			if myMaze.StartPoint.X == h && myMaze.StartPoint.Y == v {
				myColor = color.NRGBA{R: 250, G: 250, B: 30, A: 190}
			}
			if myMaze.EndPoint.X == h && myMaze.EndPoint.Y == v {
				myColor = color.NRGBA{R: 45, G: 245, B: 125, A: 190}
			}
			block := canvas.NewRectangle(myColor)
			block.Resize(fyne.NewSize(float32(myMaze.BlockSize), float32(myMaze.BlockSize)))
			block.Move(fyne.NewPos(float32(h*myMaze.BlockSize), float32(v*myMaze.BlockSize)))
			content.Add(block)
		}
	}
	player := canvas.NewImageFromFile("dot.png")
	player.Resize(fyne.NewSize(float32(myMaze.BlockSize), float32(myMaze.BlockSize)))
	player.Move(fyne.NewPos(float32(myMaze.StartPoint.X*myMaze.BlockSize), float32(myMaze.StartPoint.Y*myMaze.BlockSize)))

	content.Move(fyne.NewPos(0, 0))

	mainContent.Add(content)
	mainContent.Add(player)

	myWindow.Canvas().SetOnTypedKey(func(key *fyne.KeyEvent) {
		switch key.Name {
		case "Up", "W":
			if h := myMaze.CurrentPoint.Y - 1; h >= 0 && myMaze.Plan[h][myMaze.CurrentPoint.X] {
				player.Move(fyne.NewPos(player.Position().X, float32(h*myMaze.BlockSize)))
				myMaze.CurrentPoint.Y = h
			}
		case "Right", "D":
			if w := myMaze.CurrentPoint.X + 1; w < myMaze.Width && myMaze.Plan[myMaze.CurrentPoint.Y][w] {
				player.Move(fyne.NewPos(float32(w*myMaze.BlockSize), player.Position().Y))
				myMaze.CurrentPoint.X = w
			}
		case "Down", "S":
			if h := myMaze.CurrentPoint.Y + 1; h < myMaze.Height && myMaze.Plan[h][myMaze.CurrentPoint.X] {
				player.Move(fyne.NewPos(player.Position().X, float32(h*myMaze.BlockSize)))
				myMaze.CurrentPoint.Y = h
			}
		case "Left", "A":
			if w := myMaze.CurrentPoint.X - 1; w >= 0 && myMaze.Plan[myMaze.CurrentPoint.Y][w] {
				player.Move(fyne.NewPos(float32(w*myMaze.BlockSize), player.Position().Y))
				myMaze.CurrentPoint.X = w
			}
		case "Q":
			ch <- status.Negative
			myWindow.Close()
		}
		player.Refresh()
		if myMaze.CurrentPoint.X == myMaze.EndPoint.X && myMaze.CurrentPoint.Y == myMaze.EndPoint.Y {
			ch <- status.Positive
			myWindow.Hide()
			myWindow.Close()
		}
	})

	myWindow.SetOnClosed(func() {
		ch <- status.Negative
	})

	myWindow.SetContent(mainContent)
	myWindow.Resize(fyne.NewSize(float32(myMaze.Width*myMaze.BlockSize), float32(myMaze.Height*myMaze.BlockSize)))

	r, _ := fyne.LoadResourceFromPath("mazeImg.svg")
	myWindow.SetIcon(r)
	myWindow.Show()
	myWindow.Content().Refresh()

	myWindow.SetFixedSize(true)

}

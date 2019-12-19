package main

import (
	"fmt"
	"github.com/EmilSabri/emiltris/tetris"
	//"image"
	//"os"
	"time"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var cellwidth int = 34

func drawBoard(imd *imdraw.IMDraw) {
	imd.Reset()
	imd.Color = colornames.Black
	imd.Push(pixel.V(float64(0), float64(0)), pixel.V(float64(10*cellwidth), float64(20*cellwidth)))
	imd.Rectangle(0)
	imd.Color = colornames.Blueviolet
	for i := 0; i < 20; i++ {
		for j := 0; j < 10; j++ {
			x := j * cellwidth // + int((gameCanv.Bounds().Max.X / float64(2))) - cellwidth*5
			y := i * cellwidth
			if tetris.Board[i][j] == true {
				imd.Push(pixel.V(float64(x), float64(y)), pixel.V(float64(x+cellwidth), float64(y+cellwidth)))
				imd.Rectangle(0)
			} else {
				imd.Push(pixel.V(float64(x), float64(y)), pixel.V(float64(x+cellwidth), float64(y)))
				imd.Line(3)
				imd.Push(pixel.V(float64(x+cellwidth), float64(y)), pixel.V(float64(x+cellwidth), float64(y+cellwidth)))
				imd.Line(3)
				imd.Push(pixel.V(float64(x), float64(y)), pixel.V(float64(x), float64(y+cellwidth)))
				imd.Line(3)
				imd.Push(pixel.V(float64(x), float64(y+cellwidth)), pixel.V(float64(x+cellwidth), float64(y+cellwidth)))
				imd.Line(3)
			}
		}
	}
}

func drawBlock(block tetris.Block, imd *imdraw.IMDraw) {
	imd.Reset()
	imd.Color = colornames.Red
	for _, point := range block.Piece {
		x := (point.X + block.X) * cellwidth
		y := (point.Y + block.Y) * cellwidth

		//fmt.Print(x, y)

		imd.Push(pixel.V(float64(x), float64(y)), pixel.V(float64(x+cellwidth), float64(y+cellwidth)))
		imd.Rectangle(0)
	}
}

func drawQueue(imd *imdraw.IMDraw) {
	imd.Reset()
	imd.Color = colornames.Pink
	for i := 0; i < 5; i++ {
		b := tetris.Blocks[(i+tetris.QueueHead)%5]
		b_struct := tetris.Block{b[0], 2, i*3 + 1, 0}
		drawBlock(b_struct, imd)
	}
}

/*
func drawQueue(imd *imdraw.IMDraw) {
	// Iterate over each block in the queue
	for i := 0; i < 5; i++ {
		b := tetris.Blocks[(i+tetris.QueueHead)%5]

		// Draw the block
		for _, point := range b[0] {
			x := (point.X + 1) * cellwidth
			y := (point.Y + (i * 3) + 1) * cellwidth

			imd.Push(pixel.V(float64(x), float64(y)), pixel.V(float64(x+cellwidth), float64(y+cellwidth)))
			imd.Rectangle(0)
		}
	}
}
*/
func run() {

	// Struct to setup the window configuration
	cfg := pixelgl.WindowConfig{
		Title:  "Emil-Tris",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	// Creates the window using the cfg struct
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	// Canvas setup
	// gameCanv fits the shift block, board, and blockQueue within it
	gameCanv := pixelgl.NewCanvas(pixel.R(0, 0, cfg.Bounds.Max.X*.625, cfg.Bounds.Max.Y*.90))
	gameCanv.Clear(colornames.Yellow)

	// Build board with grids
	var cellwidth int = int(gameCanv.Bounds().Max.Y) / 20
	imdBoard := imdraw.New(nil)
	drawBoard(imdBoard)

	// Block Queue
	xMin, yMin := gameCanv.Bounds().Min.X, gameCanv.Bounds().Min.Y
	xMax, yMax := gameCanv.Bounds().Max.X, gameCanv.Bounds().Max.Y
	fmt.Println("min", xMin, yMin)
	fmt.Println("max", xMax, yMax)
	queueCanv := pixelgl.NewCanvas(pixel.R(0, 0, float64(6*cellwidth), float64(15*cellwidth)))
	queueCanv.Clear(colornames.Teal)
	imdQueue := imdraw.New(nil)
	tetris.InitQueue()

	// ----------------------
	// Convert into function
	// Draw blocks
	curBlock := tetris.NewBlock(tetris.Queue[tetris.QueueHead], 5, 18, 0)
	curBlockType := tetris.Queue[tetris.QueueHead]
	// ------------------------------
	drop_tick := time.Tick(925 * time.Millisecond)
	for !win.Closed() {
		imdBlock := imdraw.New(nil)
		imdBlock.Color = colornames.Red

		// Drops block every drop_tick amount of time

		select {
		case <-drop_tick:
			curBlock.MoveDown(-1)
		default:
		}

		// Game Controls
		// Movement left/right
		if win.Pressed(pixelgl.KeyLeft) {
			curBlock.MoveLeft()
		} else if win.Pressed(pixelgl.KeyRight) {
			curBlock.MoveRight()
		}

		// Rotation
		if win.JustPressed(pixelgl.KeyUp) {
			curBlock.R = (curBlock.R + 1) % 4
			curBlock.Piece = tetris.Blocks[curBlockType][curBlock.R]
		}

		// Hard drop
		if win.JustPressed(pixelgl.KeySpace) {
			curBlock.HardDrop()
			curBlock = tetris.NewBlock(curBlock.Landed(), 5, 18, 0)
		} else if win.Pressed(pixelgl.KeyDown) { // Soft drop -- Change later to adjust for settings
			curBlock.MoveDown(-1)
		}

		// Swap block -- Only once per queue change
		if win.JustPressed(pixelgl.KeyLeftShift) {
			fmt.Println("Shift!")
		}

		drawBoard(imdBoard)
		imdBoard.Draw(gameCanv)
		drawBlock(*curBlock, imdBlock)
		imdBlock.Draw(gameCanv)
		drawQueue(imdQueue)
		imdQueue.Draw(queueCanv)
		queueCanv.Draw(gameCanv, pixel.IM.Moved(pixel.Vec{480, 520}))
		gameCanv.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}

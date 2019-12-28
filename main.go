package main

import (
	"fmt"
	"github.com/EmilSabri/emiltris/library"
	"github.com/EmilSabri/emiltris/tetris"
	//"image"
	//"os"
	"time"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"image/color"
)

var cellwidth int = 34

var blockColors [7]color.RGBA = [...]color.RGBA{colornames.Lightblue, colornames.Yellow, colornames.Purple, colornames.Green, colornames.Red, colornames.Blue, colornames.Orange}

type Cell struct {
	x, y  int
	color color.RGBA
}

func insertCells(b *tetris.Block, bType int, l *list.List) {
	color := blockColors[bType]
	for _, point := range b.Piece {
		x, y := point.X+b.X, point.Y+b.Y
		cell := list.Cell{x, y, color}
		l.Insert(cell)
	}
}

func drawBlock(block tetris.Block, imd *imdraw.IMDraw) {
	//imd.Reset()
	//imd.Color = colornames.Red
	for _, point := range block.Piece {
		x := (point.X + block.X) * cellwidth
		y := (point.Y + block.Y) * cellwidth

		//fmt.Print(x, y)

		imd.Push(pixel.V(float64(x), float64(y)), pixel.V(float64(x+cellwidth), float64(y+cellwidth)))
		imd.Rectangle(0)
	}
}

func drawCells(cells *list.List, imdCells *imdraw.IMDraw) {
	for cell := cells.Front(); cell != nil; cell = cell.Next() {
		imdCells.Color = cell.Color
		x := cell.X * cellwidth
		y := cell.Y * cellwidth
		imdCells.Push(pixel.V(float64(x), float64(y)), pixel.V(float64(x+cellwidth), float64(y+cellwidth)))
		imdCells.Rectangle(0)
	}
}

func drawBox(x, y int, imdBoard *imdraw.IMDraw) {
	x *= cellwidth
	y *= cellwidth

	imdBoard.Color = colornames.Black
	imdBoard.Push(pixel.V(float64(x), float64(y)), pixel.V(float64(x+cellwidth), float64(y+cellwidth)))
	imdBoard.Rectangle(0)

	imdBoard.Color = colornames.Blueviolet
	imdBoard.Push(pixel.V(float64(x), float64(y)), pixel.V(float64(x+cellwidth), float64(y)))
	imdBoard.Line(3)
	imdBoard.Push(pixel.V(float64(x+cellwidth), float64(y)), pixel.V(float64(x+cellwidth), float64(y+cellwidth)))
	imdBoard.Line(3)
	imdBoard.Push(pixel.V(float64(x), float64(y)), pixel.V(float64(x), float64(y+cellwidth)))
	imdBoard.Line(3)
	imdBoard.Push(pixel.V(float64(x), float64(y+cellwidth)), pixel.V(float64(x+cellwidth), float64(y+cellwidth)))
	imdBoard.Line(3)
}

func drawQueue(imd *imdraw.IMDraw) {
	imd.Reset()
	imd.Color = colornames.Black
	imd.Push(pixel.V(0, 0), pixel.V(float64(6*cellwidth), float64(15*cellwidth)))
	imd.Rectangle(0)
	imd.Color = colornames.Purple
	for i := 0; i < 5; i++ {
		b := tetris.Blocks[(i+tetris.QueueHead)%5]
		b_struct := tetris.Block{b[0], 2, i*3 + 1, 0}
		drawBlock(b_struct, imd)
	}
}

func drawSwapped(block *tetris.Block, imdSwapped *imdraw.IMDraw, xMax, yMax float64) {
	imdSwapped.Reset()
	imdSwapped.SetMatrix(pixel.IM.Moved(pixel.Vec{xMax - 200.0, 0.0}))
	imdSwapped.Color = colornames.Black
	imdSwapped.Push(pixel.V(0.0, 0.0), pixel.V(200.0, 200.0))
	imdSwapped.Rectangle(0)
	imdSwapped.Color = colornames.Green
	drawBlock(*block, imdSwapped)

}

func drawBoard(imdBoard *imdraw.IMDraw, row int) {
	imdBoard.Color = colornames.Black
	imdBoard.Push(pixel.V(float64(0), float64(0)), pixel.V(float64(10*cellwidth), float64(20*cellwidth)))
	imdBoard.Rectangle(0)
	imdBoard.Color = colornames.Blueviolet
	for i := 0; i < 11; i++ {
		x := i * cellwidth
		y := 20 * cellwidth
		imdBoard.Push(pixel.V(float64(x), float64(0)), pixel.V(float64(x), float64(y)))
		imdBoard.Line(3)
	}
	for j := 0; j < row; j++ {
		x := 10 * cellwidth
		y := j * cellwidth
		imdBoard.Push(pixel.V(float64(0), float64(y)), pixel.V(float64(x), float64(y)))
		imdBoard.Line(3)
	}
}

/*
func drawClear(imdBoard *imdraw.IMDraw, rows []int) {
	for _, y := range rows {
		for _, x :=
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
	gameCanv := pixelgl.NewCanvas(pixel.R(0, 0, cfg.Bounds.Max.X*.625+200, cfg.Bounds.Max.Y*.90))
	gameCanv.Clear(colornames.Yellow)

	// Build board with grids
	var cellwidth int = int(gameCanv.Bounds().Max.Y) / 20
	imdBoard := imdraw.New(nil)
	drawBoard(imdBoard, 21)

	imdCells := imdraw.New(nil)
	cells := list.New()

	// Block Queue
	xMin, yMin := gameCanv.Bounds().Min.X, gameCanv.Bounds().Min.Y
	xMax, yMax := gameCanv.Bounds().Max.X, gameCanv.Bounds().Max.Y
	fmt.Println("min", xMin, yMin)
	fmt.Println("max", xMax, yMax)
	queueCanv := pixelgl.NewCanvas(pixel.R(0, 0, float64(6*cellwidth), float64(15*cellwidth)))
	queueCanv.Clear(colornames.Teal)
	imdQueue := imdraw.New(nil)
	tetris.InitQueue()

	// Swapped Block
	imdSwapped := imdraw.New(nil)
	swappedBlock := &tetris.Block{R: -1}
	var swappedBlockType int
	// ----------------------
	// Convert into function
	// Draw blocks
	curBlockType := tetris.PopQueue()
	curBlock := tetris.NewBlock(curBlockType, 5, 18, 0)

	// ------------------------------
	drop_tick := time.Tick(925 * time.Millisecond)

	var (
		frames = 0
		second = time.Tick(time.Second)
	)

	drawQueue(imdQueue)
	imdQueue.Draw(queueCanv)
	queueCanv.Draw(gameCanv, pixel.IM.Moved(pixel.Vec{480, 520}))

	drawSwapped(swappedBlock, imdSwapped, xMax, 0)
	imdSwapped.Draw(gameCanv)

	for !win.Closed() {

		imdBlock := imdraw.New(nil)
		imdBlock.Reset()
		imdBlock.Color = blockColors[curBlockType]

		// Drops block every drop_tick amount of time

		frames++
		select {
		case <-drop_tick:
			curBlock.MoveDown(-1)
		case <-second:
			win.SetTitle(fmt.Sprintf("%s | FPS: %d", cfg.Title, frames))
			frames = 0
		default:
		}

		// Game Controls
		// Movement left/right
		if win.JustPressed(pixelgl.KeyLeft) {
			curBlock.MoveLeft()
		} else if win.JustPressed(pixelgl.KeyRight) {
			curBlock.MoveRight()
		}

		// Rotation
		if win.JustPressed(pixelgl.KeyUp) {
			curBlock.Rotate(curBlockType)
		}

		// Hard drop
		if win.JustPressed(pixelgl.KeySpace) {
			curBlock.HardDrop()
			insertCells(curBlock, curBlockType, &cells)

			rows := curBlock.ClearBoard() // make tetris.ClearBoard a method

			if rows != nil {
				imdCells.Reset()
				fmt.Println("Rows Cleared:", rows)
				//fmt.Println("Len Before:", cells.Len())
				// turn into function call
				for cell := cells.Front(); cell != nil; cell = cell.Next() {
					for _, r := range rows {
						if cell.Y == r { // Remove cells if they're in the same row as rows
							//imdCells.Color = colornames.Blueviolet
							//drawBox(cell.X, cell.Y, imdCells)
							cells.Remove(list.Cell{cell.X, cell.Y, cell.Color})
							break
						} else if cell.Y > r { // Drop the cells.y value if they are above the highest cleared point by len(rows)
							//fmt.Print(" Cell.Y: ", cell.Y, " ")
							cell.Y = cell.Y - len(rows)
							break
						}
					}

				}

				// Update the board above the cells with empty
				fmt.Println("Rows Cleared: ", rows)
				tetris.PrintBoard()
				fmt.Println("\n")
				drawBoard(imdCells, 21)
				// fmt.Println("Len After:", cells.Len())
				// Draw every point in cells to imdCells
				for cell := cells.Front(); cell != nil; cell = cell.Next() {
					fmt.Println("Cell: (", cell.X, ", ", cell.Y, ") ")
				}
			}

			drawCells(&cells, imdCells)
			curBlockType = curBlock.Landed()
			curBlock = tetris.NewBlock(curBlockType, 5, 18, 0)

			drawQueue(imdQueue)
			imdQueue.Draw(queueCanv)
			queueCanv.Draw(gameCanv, pixel.IM.Moved(pixel.Vec{480, 520}))
		} else if win.Pressed(pixelgl.KeyDown) { // Soft drop -- Change later to adjust for settings
			curBlock.MoveDown(-1)
		}

		// Swap block -- Only once per queue change
		if win.JustPressed(pixelgl.KeyLeftShift) {
			if swappedBlock.R == -1 {
				swappedBlockType = curBlockType
				swappedBlock = curBlock
				curBlockType = tetris.PopQueue()
				curBlock = tetris.NewBlock(curBlockType, 5, 18, 0)

			} else {
				swappedBlock, curBlock = curBlock, swappedBlock
				swappedBlockType, curBlockType = curBlockType, swappedBlockType

				curBlock.X, curBlock.Y = 5, 18
			}
			swappedBlock.X, swappedBlock.Y = 2, 2

			// Swap slows is bottle neck in performance
			drawSwapped(swappedBlock, imdSwapped, xMax, 0)
			imdSwapped.Draw(gameCanv)

			//fmt.Println("ActiveRow: ", tetris.GetActiveRow())
			//tetris.PrintBoard()

			fmt.Println("Shift!")
		}

		imdBoard.Draw(gameCanv)
		imdCells.Draw(gameCanv)        // The cells currently placed on the board
		drawBlock(*curBlock, imdBlock) // Used for drawing the current block as it drops/moves
		imdBlock.Draw(gameCanv)
		//drawQueue(imdQueue)
		//imdQueue.Draw(queueCanv)
		//queueCanv.Draw(gameCanv, pixel.IM.Moved(pixel.Vec{480, 520}))
		//drawSwapped(swappedBlock, imdSwapped, xMax, 0)
		//imdSwapped.Draw(gameCanv)

		gameCanv.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}

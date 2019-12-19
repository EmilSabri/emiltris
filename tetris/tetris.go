package tetris

import (
	"fmt"
	"image"
	"math/rand"
	"time"
)

const boardHeight = 20
const boardWidth = 10

var board = [boardHeight + 1][boardWidth + 1]bool{}
var Board = board
var activeRow int = boardHeight - 1 // Heighest row (smallest y) with a block in it
var Queue [5]int
var QueueHead int = 0
var random = rand.New((rand.NewSource(time.Now().UnixNano())))

var IBlockLeft = [4]image.Point{
	image.Pt(-1, 0),
	image.Pt(0, 0),
	image.Pt(1, 0),
	image.Pt(2, 0),
}

var IBlockUp = [4]image.Point{
	image.Pt(1, -1),
	image.Pt(1, 0),
	image.Pt(1, 1),
	image.Pt(1, 2),
}

var IBlockRight = [4]image.Point{
	image.Pt(2, 1),
	image.Pt(1, 1),
	image.Pt(0, 1),
	image.Pt(-1, 1),
}

var IBlockDown = [4]image.Point{
	image.Pt(0, 2),
	image.Pt(0, 1),
	image.Pt(0, 0),
	image.Pt(0, -1),
}

var OBlock0 = [4]image.Point{
	image.Pt(0, 0),
	image.Pt(1, 0),
	image.Pt(0, 1),
	image.Pt(1, 1),
}

var TBlock0 = [4]image.Point{
	image.Pt(0, 0),
	image.Pt(-1, 0),
	image.Pt(1, 0),
	image.Pt(0, 1),
}

var TBlock1 = [4]image.Point{
	image.Pt(0, 0),
	image.Pt(0, 1),
	image.Pt(1, 0),
	image.Pt(0, -1),
}

var TBlock2 = [4]image.Point{
	image.Pt(0, 0),
	image.Pt(-1, 0),
	image.Pt(1, 0),
	image.Pt(0, -1),
}

var TBlock3 = [4]image.Point{
	image.Pt(0, 0),
	image.Pt(0, 1),
	image.Pt(-1, 0),
	image.Pt(0, -1),
}

var SBlock0 = [4]image.Point{
	image.Pt(0, 0),
	image.Pt(-1, 0),
	image.Pt(0, 1),
	image.Pt(1, 1),
}

var SBlock1 = [4]image.Point{
	image.Pt(0, 0),
	image.Pt(0, 1),
	image.Pt(1, 0),
	image.Pt(1, -1),
}

var SBlock2 = [4]image.Point{
	image.Pt(0, 0),
	image.Pt(1, 0),
	image.Pt(0, -1),
	image.Pt(-1, -1),
}

var SBlock3 = [4]image.Point{
	image.Pt(0, 0),
	image.Pt(-1, 0),
	image.Pt(-1, 1),
	image.Pt(0, -1),
}

var ZBlock0 = [4]image.Point{
	image.Pt(0, 0),
	image.Pt(0, 1),
	image.Pt(-1, 1),
	image.Pt(1, 0),
}

var ZBlock1 = [4]image.Point{
	image.Pt(0, 0),
	image.Pt(1, 1),
	image.Pt(1, 0),
	image.Pt(0, -1),
}

var ZBlock2 = [4]image.Point{
	image.Pt(0, 0),
	image.Pt(0, -1),
	image.Pt(1, -1),
	image.Pt(-1, 0),
}

var ZBlock3 = [4]image.Point{
	image.Pt(-1, 0),
	image.Pt(0, 1),
	image.Pt(0, 0),
	image.Pt(-1, -1),
}

var JBlock0 = [4]image.Point{
	image.Pt(0, 0),
	image.Pt(-1, 0),
	image.Pt(-1, 1),
	image.Pt(1, 0),
}

var JBlock1 = [4]image.Point{
	image.Pt(0, 0),
	image.Pt(0, 1),
	image.Pt(1, 1),
	image.Pt(0, -1),
}

var JBlock2 = [4]image.Point{
	image.Pt(0, 0),
	image.Pt(-1, 0),
	image.Pt(1, 0),
	image.Pt(1, -1),
}

var JBlock3 = [4]image.Point{
	image.Pt(0, 0),
	image.Pt(0, -1),
	image.Pt(0, 1),
	image.Pt(-1, -1),
}

var LBlock0 = [4]image.Point{
	image.Pt(0, 0),
	image.Pt(-1, 0),
	image.Pt(1, 0),
	image.Pt(1, 1),
}

var LBlock1 = [4]image.Point{
	image.Pt(0, 0),
	image.Pt(1, 0),
	image.Pt(0, 1),
	image.Pt(0, 2),
}

var LBlock2 = [4]image.Point{
	image.Pt(-1, 0),
	image.Pt(-1, 1),
	image.Pt(0, 1),
	image.Pt(1, 1),
}

var LBlock3 = [4]image.Point{
	image.Pt(0, 0),
	image.Pt(0, 1),
	image.Pt(0, 2),
	image.Pt(-1, 2),
}

type Block struct {
	Piece [4]image.Point
	X     int // 0 <= xMovement <= 9
	Y     int // 0 <= yMovement <= 19
	R     int // rotation: used to access the arrays of XBlock
}

var IBlock = [4][4]image.Point{IBlockLeft, IBlockUp, IBlockRight, IBlockDown}
var OBlock = [4][4]image.Point{OBlock0, OBlock0, OBlock0, OBlock0}
var TBlock = [4][4]image.Point{TBlock0, TBlock1, TBlock2, TBlock3}
var SBlock = [4][4]image.Point{SBlock0, SBlock1, SBlock2, SBlock3}
var ZBlock = [4][4]image.Point{ZBlock0, ZBlock1, ZBlock2, ZBlock3}
var JBlock = [4][4]image.Point{JBlock0, JBlock1, JBlock2, JBlock3}
var LBlock = [4][4]image.Point{LBlock0, LBlock1, LBlock2, LBlock3}
var Blocks = [...][4][4]image.Point{IBlock, OBlock, TBlock, SBlock, ZBlock, JBlock, LBlock}

func InitQueue() {
	for i := 0; i < 5; i++ {
		Queue[i] = random.Int() % 7
	}
}

func randBlock() int {
	return random.Int() % 7
}

// Given a row number all rows above it will shift down by 1
// 1 <= row <= BoardHeight - 1
func clear(row int) {
	if row == 1 { // Very top most row will having nothing to shift. Set all cells to false
		for i := 0; i < boardWidth; i++ {
			board[row][i] = false
		}
	} else {
		for y := row; y >= activeRow; y-- {
			for x := 0; x < boardWidth; x++ {
				board[y][x] = board[y-1][x]
			}
		}
	}

	if activeRow != boardHeight {
		activeRow++
	}
}

func clearBoard(b *Block) {
	// Get rows where the block landed
	lower, higher := b.Piece[0].Y, b.Piece[0].Y
	for _, point := range b.Piece {
		if point.Y < lower {
			lower = point.Y
		}
		if point.Y > higher {
			higher = point.Y
		}
	}

	fmt.Println("lower", lower, "higher", higher)
	// Check which rows can be cleared
	for y := lower; y <= higher; y++ {
		cnt := 0
		for x := 0; x <= boardWidth; x++ {
			if board[y][x] == true {
				cnt++
			}
		}

		// Clears that row (sets everything above it down by 1)
		if cnt == boardWidth {
			// board[y][boardWidth] = true
			clear(y)
		}
	}

}

//------------------------------------------
// Checks if the given block can move
// left: -1
// right: 1
// down: 2
func checkMove(b *Block, move int) bool {
	if move == -1 {
		for _, point := range b.Piece {
			x, y := point.X+b.X, point.Y+b.Y
			if x <= 0 || board[y][x-1] == true {
				return false
			}
		}
	} else if move == 1 {
		for _, point := range b.Piece {
			x, y := point.X+b.X, point.Y+b.Y
			if x >= boardWidth || board[y][x+1] == true {
				return false
			}
		}
	} else {
		for _, point := range b.Piece {
			x, y := point.X+b.X, point.Y+b.Y
			if y <= 0 || board[y+1][x] == true {
				return false
			}
		}
	}
	return true
}

func (b *Block) MoveLeft() {
	if b.X > 0 && checkMove(b, -1) {
		b.X -= 1
	}
}

func (b *Block) MoveRight() {
	if b.X < boardWidth && checkMove(b, 1) {
		b.X += 1
	}
}

func (b *Block) MoveDown(speed int) {
	if b.Y > 0 && checkMove(b, 2) {
		b.Y += speed
	}
}

// Moves the block to the lowest level
func (b *Block) HardDrop() {
	// Get the max  y value out of all the points and drop the piece
	// by the delta Y
	maxY := 0
	deltaY := 0
	for _, point := range b.Piece {
		x, y := point.X+b.X, point.Y+b.Y

		// Given the point (x,y) go down the column to find lowest level
		for i := y; i >= 0; i-- {
			fmt.Print("i: ", i)
			if board[i][x] == true {
				if i >= maxY {
					maxY = i
					deltaY = y - i - 1
				}
			} else if i == 0 {
				if i >= maxY {
					maxY = i
					deltaY = y
				}
			}
		}
	}

	b.Y = b.Y - deltaY
}

// Does work when a block lands and returns a new block for the player
func (b *Block) Landed() int {
	// Paint the board true on the points where the block landed
	for _, point := range b.Piece {
		x, y := point.X+b.X, point.Y+b.Y
		board[y][x] = true
	}

	// Do stuff to queue
	newBlock := Queue[QueueHead]
	Queue[QueueHead] = randBlock()
	QueueHead = (QueueHead + 1) % 5
	return newBlock
}

// ---------------------------

func NewBlock(bType, x, y, r int) *Block {
	return &Block{Blocks[bType][0], x, y, r}
}

// Draws the pieces to the board state
func paintBoard(piece [4]image.Point, x_shift, y_shift int) {
	for _, point := range piece {
		x, y := point.X+x_shift, point.Y+y_shift
		board[y][x] = true
	}
}

func printBoard() {
	for _, row := range board {
		fmt.Println(row)
	}
}

func main() {

	// Set last row and very first row to be filled and cleared
	// testRow := boardHeight

	for _, b := range IBlock {
		fmt.Print(b)
	}

	/*
		activeRow = boardHeight - 3

		for y := boardHeight; y >= activeRow; y-- {
			for x := 0; x < boardWidth; x++ {
				board[y][x] = true
			}
		}
		board[20][8] = false
		board[18][8] = false

		keyBlock := block{}
		keyBlock.piece = iBlockDown
		keyBlock.piece[0].Y = 19
		keyBlock.piece[1].Y = 16
		keyBlock.piece[2].Y = 17
		keyBlock.piece[3].Y = 18

		printBoard()
		fmt.Println("\n\n\n")

		clearBoard(&keyBlock)
		printBoard()
	*/
}

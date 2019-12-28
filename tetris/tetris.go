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
var activeRow int = 0 // Heighest row (smallest y) with a block in it
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
	image.Pt(-1, 1),
	image.Pt(1, 0),
}

var LBlock1 = [4]image.Point{
	image.Pt(0, 0),
	image.Pt(0, 1),
	image.Pt(1, 1),
	image.Pt(0, -1),
}

var LBlock2 = [4]image.Point{
	image.Pt(0, 0),
	image.Pt(-1, 0),
	image.Pt(1, 0),
	image.Pt(1, -1),
}

var LBlock3 = [4]image.Point{
	image.Pt(0, 0),
	image.Pt(0, 1),
	image.Pt(0, -1),
	image.Pt(-1, -1),
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

//var Blocks = [...][4][4]image.Point{ZBlock, ZBlock, ZBlock, ZBlock, ZBlock, ZBlock, ZBlock}

func InitQueue() {
	for i := 0; i < 5; i++ {
		Queue[i] = random.Int() % 7
	}
}

func randBlock() int {
	return random.Int() % 7
}

func PopQueue() int {
	newBlock := Queue[QueueHead]
	Queue[QueueHead] = randBlock()
	QueueHead = (QueueHead + 1) % 5
	return newBlock
}

// Given a row number all rows above it will shift down by 1
// 1 <= row <= BoardHeight - 1
func clear(row int) {
	if row == 19 { // Very top most row will having nothing to shift. Set all cells to false
		for i := 0; i < boardWidth; i++ {
			board[row][i] = false
		}
	} else {
		for y := row; y <= 18; y++ {
			for x := 0; x < boardWidth; x++ {
				board[y][x] = board[y+1][x]
			}
		}
	}

	if activeRow > 0 {
		activeRow--
	}
}

func (b *Block) ClearBoard() []int {

	var cleared_rows []int

	// Get rows where the block landed
	lower := b.Piece[0].Y
	higher := lower
	for _, point := range b.Piece {
		if point.Y < lower {
			lower = point.Y
		}
		if point.Y > higher {
			higher = point.Y
		}
	}

	//fmt.Println("lower", lower, "higher", higher)
	// Check which rows can be cleared
	for y := lower; y <= higher; y++ {
		cnt := 0
		for x := 0; x < boardWidth; x++ {
			if board[y+b.Y][x] == true {
				cnt++
			}
		}
		//PrintBoard()
		// Clears that row (sets everything above it down by 1)
		if cnt == boardWidth {
			cleared_rows = append(cleared_rows, y)
			// board[y][boardWidth] = true
			clear(y)
		}
	}

	return cleared_rows
}

//------------------------------------------
// Checks if the given block can move based on it's points
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
			if x >= boardWidth-1 || board[y][x+1] == true {
				return false
			}
		}
	} else {
		for _, point := range b.Piece {
			x, y := point.X+b.X, point.Y+b.Y
			if y <= 0 || board[y-1][x] == true {
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

// Moves the block to the lowest row possible
func (b *Block) HardDrop() {
	prevY := b.Y
	b.MoveDown(-1)
	for prevY-b.Y != 0 {
		prevY = b.Y
		b.MoveDown(-1)
	}

	updateActiveRow(b)
	paintBoard(b)
}

/*
func (b *Block) HardDrop() {
	deltaY := b.Y + 5
	for _, point := range b.Piece {
		x, y := point.X+b.X, point.Y+b.Y

		for i := y - 1; i >= 0; i-- {
			if board[i][x] == true {
				if deltaY > y-i-1 {
					deltaY = y - i - 1
				}
				break
			} else if i == 0 {
				if deltaY > y {
					deltaY = b.Y
				}
			}
		}
	}

	//fmt.Println("HardDrop()")
	//fmt.Println("b.Y", b.Y)
	//fmt.Println("deltaY:", deltaY)

	if b.Y-deltaY < 0 {
		b.Y = 0
	} else {
		b.Y = b.Y - deltaY
	}

	//fmt.Println("b.Y", b.Y)
	//fmt.Println()

	updateActiveRow(b)

	paintBoard(b)
}
*/

func (b *Block) Rotate(bType int) {
	b.wallKick(bType)
}

var testRotate0 = [5]image.Point{
	image.Pt(0, 0),
	image.Pt(-1, 0),
	image.Pt(-1, 1),
	image.Pt(0, -2),
	image.Pt(-1, -2),
}

var testRotate1 = [5]image.Point{
	image.Pt(0, 0),
	image.Pt(1, 0),
	image.Pt(1, -1),
	image.Pt(0, 2),
	image.Pt(1, 2),
}

var testRotate2 = [5]image.Point{
	image.Pt(0, 0),
	image.Pt(1, 0),
	image.Pt(1, 1),
	image.Pt(0, -2),
	image.Pt(1, -2),
}

var testRotate3 = [5]image.Point{
	image.Pt(0, 0),
	image.Pt(-1, 0),
	image.Pt(-1, -1),
	image.Pt(0, 2),
	image.Pt(-1, 2),
}

var wallKickDataMany = [4][5]image.Point{testRotate0, testRotate1, testRotate2, testRotate3}

var testRotateI0 = [5]image.Point{
	image.Pt(0, 0),
	image.Pt(-2, 0),
	image.Pt(1, 0),
	image.Pt(-2, -1),
	image.Pt(1, 2),
}

var testRotateI1 = [5]image.Point{
	image.Pt(0, 0),
	image.Pt(-1, 0),
	image.Pt(2, 0),
	image.Pt(-1, 2),
	image.Pt(2, 1),
}

var testRotateI2 = [5]image.Point{
	image.Pt(0, 0),
	image.Pt(2, 0),
	image.Pt(-1, 0),
	image.Pt(2, 1),
	image.Pt(-1, -2),
}

var testRotateI3 = [5]image.Point{
	image.Pt(0, 0),
	image.Pt(1, 0),
	image.Pt(-2, 0),
	image.Pt(1, -2),
	image.Pt(-2, 1),
}

var wallKickDataI = [4][5]image.Point{testRotateI0, testRotateI1, testRotateI2, testRotateI3}

/*
The wall kick data used in wallKickDataI and wallKickDataMany was extracted from https://tetris.wiki/Super_Rotation_System.
The tables are comprised as 0 -> R, R->2, 2->L, L->0
*/

// Test rotation for possible wall kicks
func (b *Block) wallKick(bType int) {
	var wallKickData [4][5]image.Point
	switch bType {
	case 0:
		wallKickData = wallKickDataI
	default:
		wallKickData = wallKickDataMany // For blocks other than I (excluding O block)
	}

	// Temporary block is used to test the 5 rotation test cases
	tempBlock := NewBlock(bType, b.X, b.Y, rotateBlock(b.R))

	for _, testPt := range wallKickData[b.R] { // [b.R] is the initial rotation state
		tempBlock.X, tempBlock.Y = b.X+testPt.X, b.Y+testPt.Y
		if checkRotation(tempBlock) {
			b.Piece = tempBlock.Piece
			b.X, b.Y = tempBlock.X, tempBlock.Y
			b.R = tempBlock.R
			break
		}
	}
}

// Checks if the block's current rotation state and X/Y is out of
// bounds or colliding with other blocks.
func checkRotation(b *Block) bool {
	for _, point := range b.Piece {
		x, y := point.X+b.X, point.Y+b.Y

		if x < 0 || x >= boardWidth-1 { // Test x or y is out of the board
			return false
		}
		if y < 0 || y >= boardHeight {
			return false
		}
		if board[y][x] == true { // Tests collision with blocks
			return false
		}
	}

	return true
}

func rotateBlock(r int) int {
	return (r + 1) % 4
}

func updateActiveRow(b *Block) {
	for _, point := range b.Piece {
		y := point.Y + b.Y
		if y > activeRow {
			activeRow = y
		}
	}
}

// Does work when a block lands and returns a new block for the player
func (b *Block) Landed() int {
	// Paint the board true on the points where the block landed

	//paintBoard(b)

	// Do stuff to queue
	return PopQueue()
}

// ---------------------------

func NewBlock(bType, x, y, r int) *Block {
	return &Block{Blocks[bType][r], x, y, r}
}

// Draws the pieces to the board state
func paintBoard(b *Block) {
	for _, point := range b.Piece {
		x, y := point.X+b.X, point.Y+b.Y
		board[y][x] = true
	}
}

func GetActiveRow() int {
	return activeRow
}

func GetBoard() [21][11]bool {
	return board
}

func PrintBoard() {
	for i := 19; i >= 0; i-- {
		fmt.Print("[|")
		for j := 0; j < 10; j++ {
			cell := board[i][j]
			if cell == false {
				fmt.Print(" - |")
			} else {
				fmt.Print(" + |")
			}
		}
		fmt.Println("]")
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

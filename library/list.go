/*
	Implements a singly linked list for the cell struct
*/

package list

import (
	"fmt"
	"image/color"
)

type Cell struct {
	X, Y  int
	Color color.RGBA
}

type Node struct {
	Cell
	next *Node
	prev *Node
}

type List struct {
	head *Node
	len  int
}

func New() List {
	return List{nil, 0}
}

func (l *List) Front() *Node {
	return l.head
}

func (n *Node) Next() *Node {
	return n.next
}

func (l *List) Len() int {
	return l.len
}

func (l *List) Update(cell Cell) {
	l.head.Cell = cell
}

// Inserts at the front of the list
func (l *List) Insert(cell Cell) {

	node := &Node{Cell{cell.X, cell.Y, cell.Color}, l.head, nil}

	if l.len > 0 {
		l.head.prev = node
	}

	l.head = node
	l.len++
}

func (l *List) Remove(cell Cell) {
	if l.len == 1 && l.head.Cell == cell {
		l.head = nil
		return
	}

	for c := l.head; c != nil; c = c.next {
		if c.Cell == cell {

			if l.head == c { // Head
				l.head = c.next
				l.head.prev = nil
				return
			} else if c.next == nil { // Tail
				c.prev.next = nil
				return
			} else { // Middle
				c.prev.next = c.next
				c.next.prev = c.prev
				return
			}
		}
	}

	l.len--
}

func (l *List) PrintList() {
	for e := l.head; e != nil; e = e.next {
		fmt.Println(e.X, e.Y, e.Color)
	}
}

func main() {
	color := color.RGBA{0, 0, 0, 0}
	l := []Cell{Cell{0, 0, color}, Cell{1, 1, color}, Cell{2, 2, color}, Cell{3, 3, color}}
	link := New()

	fmt.Println("Testing on linked list len() == 1")
	link.Insert(Cell{-1, -1, color})
	fmt.Println(link.Front())

	fmt.Println("\nTesting Removal")
	link.Remove(Cell{-1, -1, color})
	fmt.Println(link.Front())

	link = New()
	fmt.Println("\nTesting linked list len() > 1")
	for _, v := range l {
		link.Insert(v)
	}

	link.PrintList()

	fmt.Println("\nTesting Removal front, middle, end")
	link.PrintList()
	link.Remove(Cell{0, 0, color})
	link.PrintList()

	fmt.Println("\nTesting removal while looping through List")
	link = New()
	l = append(l, Cell{4, 4, color}, Cell{5, 5, color}, Cell{6, 6, color})
	for _, v := range l {
		link.Insert(v)
	}

	for cell := link.Front(); cell != nil; cell = cell.Next() {
		if cell.X%2 != 0 {
			link.Remove(Cell{cell.X, cell.Y, cell.Color})
		}
	}

	link.PrintList()
}

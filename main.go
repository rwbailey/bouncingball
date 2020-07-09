package main

import (
	"fmt"
	"time"
)

func main() {
	board := newBoard(30, 80)
	buffer := make([]rune, 0, board.x*board.y)
	for {
		Clear()
		MoveTopLeft()
		board.cells = newCells(board.x, board.y)
		board.updateBallOnBoard()
		buffer = buffer[:0]
		buffer := board.loadBuffer(buffer)
		fmt.Println(string(buffer))
		board.nextBallPos()
		time.Sleep(time.Second / 6)
	}
}

type ball struct {
	pos []int
	vel []int
}

type board struct {
	x     int
	y     int
	cells [][]bool
	ball  *ball
}

func (b *board) loadBuffer(buffer []rune) []rune {
	for _, v := range b.cells {
		for _, u := range v {
			if u {
				buffer = append(buffer, rune(9917))
			} else {
				buffer = append(buffer, rune(32))
			}
		}
		buffer = append(buffer, rune(10))
	}
	return buffer
}

func (b *board) updateBallOnBoard() {
	b.cells[b.ball.pos[0]][b.ball.pos[1]] = true
}

func (b *board) nextBallPos() {
	if b.ball.vel[0] == 1 {
		if b.ball.pos[0] == b.x-1 {
			b.ball.vel[0] = -1
		}
	} else {
		if b.ball.pos[0] == 0 {
			b.ball.vel[0] = 1
		}
	}
	if b.ball.vel[1] == 1 {
		if b.ball.pos[1] == b.y-1 {
			b.ball.vel[1] = -1
		}
	} else {
		if b.ball.pos[1] == 0 {
			b.ball.vel[1] = 1
		}
	}

	b.ball.pos[0] += b.ball.vel[0]
	b.ball.pos[1] += b.ball.vel[1]
}

func newBall() *ball {
	return &ball{
		pos: []int{0, 0},
		vel: []int{1, 1},
	}
}

func newBoard(w, h int) *board {

	return &board{
		x:     w,
		y:     h,
		cells: newCells(w, h),
		ball:  newBall(),
	}
}

func newCells(w, h int) [][]bool {
	cells := make([][]bool, w)
	for i := range cells {
		cells[i] = make([]bool, h)
	}
	return cells
}

// Clear clears the screen
func Clear() {
	fmt.Print("\033[2J")
}

// MoveTopLeft moves the cursor to the top left position of the screen
func MoveTopLeft() {
	fmt.Print("\033[H")
}

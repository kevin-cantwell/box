package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"golang.org/x/crypto/ssh/terminal"
)

var (
	chars = []rune{'╭', '╮', '╯', '╰', '│', '─', ' '}

	aboves = []rune{'╭', '╮', '│'}
	belows = []rune{'╯', '╰', '│'}
	rights = []rune{'╮', '╯', '─'}
	lefts  = []rune{'╭', '╰', '─'}
)

func main() {
	run()
}

func run() {
	var (
		width, height = 40, 40
	)
	if terminal.IsTerminal(int(os.Stdout.Fd())) {
		w, h, err := terminal.GetSize(int(os.Stdout.Fd()))
		if err != nil {
			panic(err)
		}
		width, height = w, h
	}
	flag.IntVar(&width, "width", width, "Width")
	flag.IntVar(&height, "height", height, "Height")
	flag.Parse()

	height--

	t := make(table, height)

	for y := 0; y < height; y++ {
		t[y] = make([]rune, width)
		for x := 0; x < width; x++ {
			t[y][x] = chars[rand.Intn(len(chars)-1)]
		}
	}

	t.Print(false)

	for {
		x := rand.Intn(width)
		y := rand.Intn(height)

		setChar(x, y, t, 8)
		t.Print(true)
		time.Sleep(10 * time.Millisecond)

	}
}

func setChar(x, y int, t table, depth int) {
	if x < 0 || y < 0 || y >= len(t) || x >= len(t[0]) {
		return
	}
	switch {
	case t.up(x, y) && t.rt(x, y) && t.dn(x, y) && t.lt(x, y):
		t[y][x] = '┼'
	case t.up(x, y) && t.rt(x, y) && t.dn(x, y) && !t.lt(x, y):
		t[y][x] = '├'
	case t.up(x, y) && t.rt(x, y) && !t.dn(x, y) && t.lt(x, y):
		t[y][x] = '┴'
	case t.up(x, y) && !t.rt(x, y) && t.dn(x, y) && t.lt(x, y):
		t[y][x] = '┤'
	case !t.up(x, y) && t.rt(x, y) && t.dn(x, y) && t.lt(x, y):
		t[y][x] = '┬'
	case t.up(x, y) && !t.rt(x, y) && t.dn(x, y) && !t.lt(x, y):
		t[y][x] = '│'
	case !t.up(x, y) && t.rt(x, y) && !t.dn(x, y) && t.lt(x, y):
		t[y][x] = '─'
	case t.up(x, y) && t.rt(x, y) && !t.dn(x, y) && !t.lt(x, y):
		t[y][x] = '╰'
	case !t.up(x, y) && t.rt(x, y) && t.dn(x, y) && !t.lt(x, y):
		t[y][x] = '╭'
	case !t.up(x, y) && !t.rt(x, y) && t.dn(x, y) && t.lt(x, y):
		t[y][x] = '╮'
	case t.up(x, y) && !t.rt(x, y) && !t.dn(x, y) && t.lt(x, y):
		t[y][x] = '╯'
	case !t.up(x, y) && !t.rt(x, y) && !t.dn(x, y) && t.lt(x, y):
		t[y][x] = '◀'
	case !t.up(x, y) && !t.rt(x, y) && t.dn(x, y) && !t.lt(x, y):
		t[y][x] = '▼'
	case !t.up(x, y) && t.rt(x, y) && !t.dn(x, y) && !t.lt(x, y):
		t[y][x] = '▶'
	case t.up(x, y) && !t.rt(x, y) && !t.dn(x, y) && !t.lt(x, y):
		t[y][x] = '▲'
	default:
		t[y][x] = ' '
	}
	// for _, i := range []int{-1, 1} {
	// 	for _, j := range []int{-1, 1} {
	// 		depth--
	// 		if depth == 0 {
	// 			break
	// 		}
	// 		setChar(x+i, y+j, t, 1)
	// 	}
	// }
}

type table [][]rune

func (t table) up(x, y int) bool {
	if y == 0 {
		return false
	}
	r := t[y-1][x]
	return in(aboves, r)
}

func (t table) dn(x, y int) bool {
	if y == len(t)-1 {
		return false
	}
	r := t[y+1][x]
	return in(belows, r)
}

func (t table) lt(x, y int) bool {
	if x == 0 {
		return false
	}
	r := t[y][x-1]
	return in(lefts, r)
}

func (t table) rt(x, y int) bool {
	if x == len(t[0])-1 {
		return false
	}
	r := t[y][x+1]
	return in(rights, r)
}

func (t table) Print(clear bool) {
	if clear {
		fmt.Printf("\033[999D\033[%dA", len(t))
	}
	for _, row := range t {
		fmt.Println(string(row))
	}
}

func in(runes []rune, r rune) bool {
	for _, e := range runes {
		if e == r {
			return true
		}
	}
	return false
}

func uniq(runes []rune) []rune {
	m := map[rune]bool{}
	out := []rune{}
	for _, r := range runes {
		if m[r] {
			continue
		}
		m[r] = true
		out = append(out, r)
	}
	return out
}

package main

import (
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

const (
	width  = 20
	height = 20
)

type position struct {
	x int
	y int
}

type snake struct {
	body       []position
	directions []position
	direction  position
}

func (s *snake) move() {
	head := s.body[0]
	head.x += s.direction.x
	head.y += s.direction.y

	if head.x < 1 || head.x >= width-1 || head.y < 1 || head.y >= height-1 {
		panic("Game Over")
	}

	for i := 1; i < len(s.body); i++ {
		if head == s.body[i] {
			// the snake has collided with its body
			s.direction = position{0, 0} // stop the snake from moving
			panic("Self collide")
		}
	}

	s.body = append([]position{head}, s.body[:len(s.body)-1]...)
}

func (s *snake) changeDirection(newDirection position) {
	if newDirection.x != -s.direction.x || newDirection.y != -s.direction.y {
		s.direction = newDirection
	}
}

func (s *snake) collidesWith(pos position) bool {
	for _, p := range s.body {
		if p == pos {
			return true
		}
	}
	return false
}

func (s *snake) grow() {
	s.body = append(s.body, s.body[len(s.body)-1])
}

func createApple(s *snake) position {
	for {
		pos := position{rand.Intn(width-2) + 1, rand.Intn(height-2) + 1}
		if !s.collidesWith(pos) {
			return pos
		}
	}
}

func main() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	rand.Seed(time.Now().UnixNano())

	s := &snake{
		body: []position{{width / 2, height / 2}},
		directions: []position{
			{-1, 0},
			{1, 0},
			{0, -1},
			{0, 1},
		},
		direction: position{1, 0},
	}

	apple := createApple(s)

	// create event queue
	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	ticker := time.Tick(300 * time.Millisecond)

	for {
		// clear screen
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

		// redraw screen
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				pos := position{x, y}
				var ch rune
				if x == 0 || y == 0 || x == width-1 || y == height-1 {
					// draw a wall character for positions on the border
					ch = '█'
				} else if s.collidesWith(pos) {
					ch = '■'
				} else if pos == apple {
					ch = '★'
				} else {
					ch = ' '
				}
				termbox.SetCell(x, y, ch, termbox.ColorDefault, termbox.ColorDefault)
			}
		}

		// flush screen
		termbox.Flush()

		select {
		case event := <-eventQueue:
			if event.Type == termbox.EventKey {
				var dir position
				switch event.Key {
				case termbox.KeyArrowUp:
					dir = position{0, -1}
				case termbox.KeyArrowDown:
					dir = position{0, 1}
				case termbox.KeyArrowLeft:
					dir = position{-1, 0}
				case termbox.KeyArrowRight:
					dir = position{1, 0}
				case termbox.KeyEsc:
					return
				}
				s.changeDirection(dir)
			}
		case <-ticker:
			// move snake
			s.move()

			if s.collidesWith(apple) {
				s.grow()
				apple = createApple(s)
			}
		}
	}
}

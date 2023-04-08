package main

import "fmt"

func main() {

	p1 := Player{
		Name: "Parikxxit",
		Item: Item{
			X: 500,
			Y: 300,
		},
	}

	fmt.Printf("%#v\n", p1)
	fmt.Println(p1.X)
}

type Player struct {
	Name string
	Item // Item is embeded inside player
}

func (i *Item) Move(x, y int) {
	i.X = x
	i.Y = y
}

func NewItem(x, y int) (*Item, error) {
	if x < 0 || x > maxX || y < 0 || y > maxY {
		return nil, fmt.Errorf("%d/%d out of bounds %d/%d", x, y, maxX, maxY)
	}
	return &Item{
		X: x,
		Y: y,
	}, nil
}

const (
	maxX = 1000
	maxY = 600
)

type Item struct {
	X int
	Y int
}

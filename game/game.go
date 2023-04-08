package main

import (
	"fmt"
	"log"
)

func main() {

	p1 := Player{
		Name: "Parikxxit",
		Item: Item{
			X: 500,
			Y: 300,
		},
	}

	i1, err := NewItem(200, 300)
	if err != nil {
		log.Fatalf("got error while creating new Item err: %s", err)
	}
	fmt.Printf("%#v\n", p1)
	fmt.Println(p1.X)
	/*
		ms := []mover{
			p1,
			i1,
		}
		HERE it will give compile tinme error as Move implemented as pointer semantics and p1 is of type Player not *Player so we need to pass address of p1 as we are doing for i1
		as to group the interface we need to be explicit i.e if interface is implemented via pointer sematinc grouping will going to have pointer value if implemented via value semantics
		we need to group it as value not by referance
	*/
	ms := []mover{
		&p1,
		i1,
	}

	moveAll(ms, 0, 0) // move all ms player or item to starting pos i.e 0,0
}

func moveAll(ms []mover, x, y int) {
	for _, m := range ms {
		m.Move(x, y)
	}
}

type mover interface {
	Move(x, y int)
	// Move(int, int) is also fine
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

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	p := Payment{
		From:   "Wile. E. Coyote",
		To:     "ACME",
		Amount: 123.12,
	}
	p.Process()
	p.Process() // Payment will get process once as it is protected by .Once sig
}

func (p *Payment) process(t time.Time) {
	ts := t.Format(time.RFC3339)
	fmt.Printf("[%s] %s -> $%.2f -> %s", ts, p.From, p.Amount, p.To)
}

func (p *Payment) Process() {
	t := time.Now()
	p.once.Do(func() { p.process(t) })
}

type Payment struct {
	From   string
	To     string
	Amount float32 //USD
	once   sync.Once
}

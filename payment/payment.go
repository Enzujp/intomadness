package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	p :=
		Payment{
			From:   "Wile, E. Coyote",
			To:     "Enzu",
			Amount: 5000.00,
		}
	p.Process()
	p.Process()
	// We can use goroutines to ensure that payments occur just the once.
}

type Payment struct {
	From   string
	To     string
	Amount float64 // USD

	once sync.Once
}

func (p *Payment) Process() {
	t := time.Now()
	p.once.Do(func() {
		p.process(t)
	})
}

func (p *Payment) process(t time.Time) {
	ts := t.Format(time.RFC3339)
	fmt.Printf("[%s],%s -> $%.2F -> %s", ts, p.From, p.Amount, p.To)
}

/*
If we were to pass in a parameter into the Payment function, say time.Time, so we could return a timestamp perhaps, we would get an error
as the sync.once does not accept parameters, to remedy this however we could make use of closures( functions within a function)
Error Groupings also take functions with no arguments.
*/

package main

import "fmt"

func main() {
	playerJesse := CreatePlayer("jesse", Item{X: 4, Y: 5})
	fmt.Printf("These are the player's details: %#v\n", playerJesse)
	newItem := Item{
		X: 2,
		Y: 4,
	}

	// create two new players
	p1 := Player{
		Name: "Charles",
		Item: Item{X: 23, Y: 34},
	}

	p2 := Player{
		Name: "Lady Anne",
		Item: Item{X: 56, Y: 54},
	}

	ms := []mover{
		&p1,
		&p2,
		&newItem,
	}

	moveAll(ms, 100, 100)



	newItem.Move(5, 9) // taps into move function and changes values for X and Y from 2, 4 to 5, 9

	fmt.Println(newItem)
	fmt.Println(p1, p2)
}

type Player struct {
	Name string
	Item
	Keys []key
}

type Item struct {
	X int
	Y int
}

type key byte // alias for uint8 type, but usually designed to hold a smaller integer value

const (
	Jade key = iota + 1
	Kemi
	Paul
	invalid_key
)

// write a function that moves all players to the end of the game at Item{X: 100, Y: 100}
func moveAll(ms []mover, x, y int ){
	for _, m := range ms{
		m.Move(100, 100)
	}
}


// write a method that creates a new player

func CreatePlayer(name string, item Item) *Player{
	newPlayer := Player{
		Name: "Jesse",
		Item: Item{X: 4, Y: 5},
	}
	return &newPlayer
}

//using switch expression to find player with the key
func (k key) String() string{
	switch k{
	case Jade:
		return "jade"
	case Kemi:
		return "kemi"
	case Paul:
		return "paul"
		
	}
	return fmt.Sprintf("<Key %d>", k)
}


// To mutate values, use pointer receivers as they change the values in the memory and not create instances
func (i *Item) Move(x, y int) { 
	i.X = x
	i.Y = y
}

type mover interface {
	Move(x, y int)
}
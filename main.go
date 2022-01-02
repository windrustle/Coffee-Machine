package main

import (
	"errors"
	"fmt"
)

var (
	espresso   = state{water: 250, milk: 0, beans: 16, cups: 1, money: 4}
	latte      = state{water: 350, milk: 75, beans: 20, cups: 1, money: 5}
	cappuccino = state{water: 200, milk: 100, beans: 12, cups: 1, money: 6}
)

func main() {
	var s = new(state)

	for {
		if exit := s.printActions(); exit {
			break
		}
	}

	fmt.Println("Exiting...")
}

type state struct {
	water int
	milk  int
	beans int
	cups  int
	money int
}

func (s *state) printActions() bool {
	fmt.Println("Choose action:\n1. Show remaining resources\n2. Buy a cup of coffee\n3. Refill resources\n4. Take money\n5. Exit")

	var acton int
	read(&acton)

	switch acton {
	case 1:
		s.remaining()

	case 2:
		s.buy()

	case 3:
		s.fill()

	case 4:
		s.take()

	case 5:
		return true

	default:
		fmt.Println("Wrong action")
		fmt.Println()
	}

	return false
}

func (s *state) remaining() {
	fmt.Println("The coffee machine has:")
	fmt.Printf("%d ml of water\n", s.water)
	fmt.Printf("%d ml of milk\n", s.milk)
	fmt.Printf("%d g of coffee beans\n", s.beans)
	fmt.Printf("%d disposable cups\n", s.cups)
	fmt.Printf("$%d\n", s.money)
	fmt.Println()
}

func (s *state) buy() {
	for {
		fmt.Println("Choose an item:\n1. Espresso\n2. Latte\n3. Cappuccino\n4. Back to main menu")

		var (
			i    state
			item int
		)
		read(&item)

		switch item {
		case 1:
			i = espresso

		case 2:
			i = latte

		case 3:
			i = cappuccino

		case 4:
			return

		default:
			fmt.Println("Wrong item")
			fmt.Println()
			continue
		}

		if err := s.coffee(i); err != nil {
			fmt.Printf("Sorry, cant make a coffee: %v!\n", err)

		} else {
			fmt.Println("Please, take your coffee!")
		}

		fmt.Println()
		return
	}
}

func (s *state) fill() {
	var water, milk, beans, cups int

	fmt.Println("Write how many ml of water you want to add:")
	s.water += water

	fmt.Println("Write how many ml of milk you want to add:")
	read(&milk)
	s.milk += milk

	fmt.Println("Write how many grams of coffee beans you want to add:")
	read(&beans)
	s.beans += beans

	fmt.Println("Write how many disposable coffee cups you want to add:")
	read(&cups)
	s.cups += cups
}

func (s *state) take() {
	fmt.Printf("I gave you $%d\n", s.money)
	fmt.Println()
	s.money = 0
}

func (s *state) coffee(i state) error {
	if s.water < i.water {
		return errors.New("not enough water")
	}

	if s.milk < i.milk {
		return errors.New("not enough milk")
	}

	if s.beans < i.beans {
		return errors.New("not enough coffee beans")
	}

	if s.cups < i.cups {
		return errors.New("not enough disposable cups")
	}

	s.water -= i.water
	s.milk -= i.milk
	s.beans -= i.beans
	s.cups -= i.cups
	s.money += i.money

	return nil
}

func read(t *int) {
	fmt.Print("\n> ")
	if _, err := fmt.Scan(t); err != nil {
		panic(fmt.Errorf("failed to read from stdin: %v", err))
	}
	fmt.Println()
}

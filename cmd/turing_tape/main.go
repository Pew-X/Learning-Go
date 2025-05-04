package main

import (
	"fmt"
)

type Symbol int

const (
	Blank Symbol = -1
	Zero  Symbol = 0
	One   Symbol = 1
)

type Direction int

const (
	Left  Direction = -1
	Right Direction = +1
	Stay  Direction = 0
)

type State string

type Transiton struct {
	Write Symbol
	Move  Direction
	Next  State
}

type TuringMachine struct {
	Tape        []Symbol
	Head        int
	State       State
	HaltState   State
	Transitions map[State]map[Symbol]Transiton
}

func (tm *TuringMachine) Step() bool {
	if tm.State == tm.HaltState {
		return false
	}

	currentSymbol := tm.Tape[tm.Head]
	transitions := tm.Transitions[tm.State]

	transition, ok := transitions[currentSymbol]

	if !ok {
		fmt.Println("No valid transition found. Halting.")
		return false
	}

	tm.Tape[tm.Head] = transition.Write
	tm.Head += int(transition.Move)
	if tm.Head < 0 {
		// Expand tape to the left
		tm.Tape = append([]Symbol{Blank}, tm.Tape...)
		tm.Head = 0
	}

	if tm.Head >= len(tm.Tape) {
		// Expand tape to the right
		tm.Tape = append(tm.Tape, Blank)
	}

	tm.State = transition.Next
	return true

}

func (tm *TuringMachine) Run() {
	for tm.Step() {
	}
}

func printTape(tape []Symbol) {
	for _, s := range tape {
		if s == Blank {
			fmt.Print("_ ")
		} else {
			fmt.Printf("%d ", s)
		}
	}
	fmt.Println()

}

func main() {
	tape := []Symbol{One, Zero, One} // 101
	tm := TuringMachine{
		Tape:      tape,
		Head:      2, // Start at end of input
		State:     "check",
		HaltState: "halt",
		Transitions: map[State]map[Symbol]Transiton{
			"check": {
				One:   {Write: Zero, Move: Left, Next: "check"},
				Zero:  {Write: One, Move: Stay, Next: "halt"},
				Blank: {Write: One, Move: Stay, Next: "halt"}, // Handles carry-over
			},
		},
	}

	fmt.Print("Before: ")
	printTape(tm.Tape)
	tm.Run()
	fmt.Print("After:  ")
	printTape(tm.Tape)
}

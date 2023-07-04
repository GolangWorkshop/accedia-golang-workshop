package main

import (
	"fmt"
	"strings"
)

type Lamp struct {
	color string
	state string
}

func NewLamp(color, state string) *Lamp {
	return &Lamp{
		color: color,
		state: state,
	}
}

func (l *Lamp) GetColor() string {
	return l.color
}

func (l *Lamp) SetColor(color string) {
	switch color {
	case "red", "green", "orange", "yellow", "blue", "purple":
		l.color = color
	default:
		l.color = "black"
	}
}

func (l *Lamp) GetState() string {
	return l.state
}

func (l *Lamp) SetState(state string) {
	l.state = state
}

func (l *Lamp) SummonGenie() bool {
	return l.color == "red" && l.state == "on"
}

func Task_6_solution() {
	input := "blue on purple off red on"
	commands := strings.Fields(input)
	lamp := NewLamp("black", "off")
	for _, value := range commands {
		switch value {
		case "on", "off":
			lamp.SetState(value)
		default:
			lamp.SetColor(value)
		}
		fmt.Println(lamp.GetColor(), lamp.GetState())
	}

	if lamp.SummonGenie() {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}

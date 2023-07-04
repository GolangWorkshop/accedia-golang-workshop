package main

import (
	"fmt"
	"strings"
)

type programmer struct {
	name     string
	language string
}

type Programmer interface {
	Code() string
}

func (p *programmer) Code() string {
	return p.name + " is writing " + p.language
}

type singer struct {
	programmer
	songs []string
}

func (s *singer) Sing() string {
	return s.name + " is singing " + strings.Join(s.songs, ", ")
}

type Singer interface {
	Sing() string
}

func main() {
	i := &programmer{
		name:     "Ivan",
		language: "javascript",
	}

	m := &singer{
		programmer: programmer{
			name:     "Mayya",
			language: "golang",
		},

		songs: []string{"Ivane, Ivane...", "doko doko"},
	}

	forceToCode(i)
	forceToCode(m)
	forceToSing(m)
}

func forceToCode(p Programmer) {
	fmt.Println(p.Code())
}

func forceToSing(s Singer) {
	fmt.Println(s.Sing())
}
